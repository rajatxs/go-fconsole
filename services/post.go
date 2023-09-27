package services

import (
	"context"
	"log"
	"time"

	"github.com/rajatxs/go-fconsole/db"
	"github.com/rajatxs/go-fconsole/models"
	"github.com/rajatxs/go-fconsole/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostService struct {
	Ctx context.Context
}

// Creates new instance of PostService
func NewPostService() *PostService {
	return &PostService{
		Ctx: nil,
	}
}

// Returns Post metadata by given Raw ID
// By default this method will return public post
func (ps *PostService) GetPostMetadataById(rawid string, private bool) (*models.PostMetadataDocument, error) {
	var (
		doc      *models.PostMetadataDocument
		oid      primitive.ObjectID
		collName string
		filter   primitive.D
		err      error
	)

	if private {
		collName = "privatePostsMetadata"
	} else {
		collName = "publicPostsMetadata"
	}

	// Parse ObjectId
	if oid, err = primitive.ObjectIDFromHex(rawid); err != nil {
		return nil, err
	} else {
		filter = bson.D{{Key: "_id", Value: oid}}
	}

	// Find single post document by _id
	if err = db.
		MongoDb().
		Collection(collName).
		FindOne(ps.Ctx, filter).
		Decode(&doc); err != nil {
		return nil, err
	} else {
		return doc, err
	}
}

// Returns list of metadata of posts
// Need to provide result limit (default is 0)
func (ps *PostService) GetPostsMetadata(params *types.GetPostsMetadataOptions) ([]models.PostMetadataDocument, error) {
	var (
		posts     []models.PostMetadataDocument
		findOpts  = options.Find()
		collName  string
		sortProp  string
		sortOrder int
		filter    primitive.D = bson.D{}
	)

	if params.Private {
		collName = "privatePostsMetadata"
	} else {
		collName = "publicPostsMetadata"
	}

	// Use specific topic
	if params.Topic != "all" {
		filter = bson.D{{Key: "topic", Value: params.Topic}}
	}

	switch params.SortBy {
	case "title":
		sortProp = "title"
		sortOrder = 1

	case "topic":
		sortProp = "topic"
		sortOrder = 1

	case "newest":
		sortProp = "createdAt"
		sortOrder = -1

	case "oldest":
		sortProp = "createdAt"
		sortOrder = 1

	case "updated":
		sortProp = "updatedAt"
		sortOrder = -1
	}

	findOpts.SetSort(map[string]int{sortProp: sortOrder})
	findOpts.SetLimit(params.Limit)
	findOpts.SetSkip(params.Skip)

	if cur, err := db.
		MongoDb().
		Collection(collName).
		Find(ps.Ctx, filter, nil, findOpts); err != nil {
		log.Println(err)
	} else {
		defer cur.Close(ps.Ctx)

		for cur.Next(ps.Ctx) {
			var post models.PostMetadataDocument

			if err = cur.Decode(&post); err != nil {
				return nil, err
			}
			posts = append(posts, post)
		}

		if err = cur.Err(); err != nil {
			return nil, err
		}
	}

	return posts, nil
}

// Returns number of post (public + private) in posts collection
// By default the result will not include count of deleted posts
func (ps *PostService) GetPostCount(scope string, includeDeleted bool) (int64, error) {
	return db.
		MongoDb().
		Collection("posts").
		CountDocuments(
			ps.Ctx,
			bson.D{
				{Key: "public", Value: scope == "public"},
				{Key: "deleted", Value: includeDeleted},
			})
}

// Inserts new post document into posts collection
func (ps *PostService) CreatePost(payload *types.CreatePostPayload) (any, error) {
	var (
		authorId primitive.ObjectID
		related  = []string{}
		err      error
	)

	body := primitive.Binary{
		Subtype: bson.TypeBinaryGeneric,
		Data:    []byte(payload.Body),
	}

	if authorId, err = primitive.ObjectIDFromHex(payload.AuthorId); err != nil {
		return nil, err
	}

	newPost := bson.M{
		"title":    payload.Title,
		"slug":     payload.Slug,
		"desc":     payload.Desc,
		"topic":    payload.Topic,
		"tags":     payload.Tags,
		"body":     body,
		"stars":    0,
		"authorId": authorId,
		"public":   payload.Public,
		"deleted":  false,
		"coverImage": &models.PostCoverImage{
			Id:      payload.CoverImageId,
			Path:    payload.CoverImagePath,
			RefName: payload.CoverImageRefName,
			RefUrl:  payload.CoverImageRefUrl,
		},
		"related":   related,
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	}
	db.MongoDb().Collection("posts").InsertOne(ps.Ctx, newPost)

	return nil, err
}

// Set specified scope of post
func (ps *PostService) UpdatePostScope(rawid string, scope string) error {
	var (
		oid    primitive.ObjectID
		err    error
		filter bson.D
		update bson.M
		public bool = (scope == "public")
	)

	if oid, err = primitive.ObjectIDFromHex(rawid); err != nil {
		return err
	} else {
		filter = bson.D{{Key: "_id", Value: oid}}
		update = bson.M{"$set": bson.M{"public": public}}
	}

	// Update scope
	_, err = db.MongoDb().Collection("posts").UpdateOne(ps.Ctx, filter, update)
	return err
}

// Sets post delete flag by given post rawid
func (ps *PostService) SetPostDeleteFlag(rawid string, value bool) error {
	var (
		oid    primitive.ObjectID
		err    error
		filter bson.D
		update bson.M
	)

	if oid, err = primitive.ObjectIDFromHex(rawid); err != nil {
		return err
	} else {
		filter = bson.D{{Key: "_id", Value: oid}}
		update = bson.M{"$set": bson.M{"deleted": value}}
	}

	// Update document
	_, err = db.MongoDb().Collection("posts").UpdateOne(ps.Ctx, filter, update)
	return err
}

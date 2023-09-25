package services

import (
	"context"
	"log"

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
func (ps *PostService) CreatePost(payload map[string]interface{}) {
	post := bson.M{
		"title": payload["title"],
	}
	db.MongoDb().Collection("posts").InsertOne(ps.Ctx, post)
}

package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/rajatxs/go-fconsole/config"
	"github.com/rajatxs/go-fconsole/db"
	"github.com/rajatxs/go-fconsole/models"
	"github.com/rajatxs/go-fconsole/types"
	"github.com/rajatxs/go-fconsole/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostService struct {
	Ctx             context.Context
	TopicServiceRef *TopicService
}

// NewPostService creates new instance of PostService
func NewPostService() *PostService {
	return &PostService{
		Ctx: nil,
	}
}

// saveIndex writes new object to post search index
func (ps *PostService) saveIndex(id primitive.ObjectID) (res search.SaveObjectRes, err error) {
	var (
		metadata *models.PostMetadataDocument
		record   *models.PostIndex
		nilRes   = search.SaveObjectRes{}
	)

	if err = db.
		MongoDb().
		Collection("publicPostsMetadata").
		FindOne(ps.Ctx, &primitive.D{{Key: "_id", Value: id}}).
		Decode(&metadata); err != nil {
		return nilRes, err
	}

	record = &models.PostIndex{
		ObjectId:  id.Hex(),
		Name:      metadata.Title,
		Topic:     ps.TopicServiceRef.GetTopicNameById(metadata.Topic),
		Desc:      metadata.Desc,
		Tags:      metadata.Tags,
		Url:       fmt.Sprintf("%s/%s", config.ClientUrl(), metadata.Slug),
		Image:     util.GetPostCoverImageUrl(metadata.CoverImage.Path),
		CreatedAt: metadata.CreatedAt,
		UpdatedAt: metadata.UpdatedAt,
	}

	if res, err = util.PostIndex().SaveObject(record); err != nil {
		util.Log.Error(fmt.Sprintf("[PostService.saveIndex] %s", err.Error()))
		return nilRes, err
	} else {
		util.Log.Info(fmt.Sprintf("[PostService.saveIndex] Saved post index (id='%s')", id.Hex()))
		return res, nil
	}
}

// dropIndex removes object from post search index
func (ps *PostService) dropIndex(id primitive.ObjectID) (res search.DeleteTaskRes, err error) {
	if res, err = util.PostIndex().DeleteObject(id.Hex()); err != nil {
		util.Log.Error(fmt.Sprintf("[PostService.dropIndex] %s", err.Error()))
		return search.DeleteTaskRes{}, err
	} else {
		util.Log.Info(fmt.Sprintf("[PostService.saveIndex] Dropped post index (id='%s')", id.Hex()))
		return res, nil
	}
}

// updateIndex handles update/delete operation in post search index
func (ps *PostService) updateIndex(id primitive.ObjectID, public bool) (err error) {
	if config.IsProd() {
		if public {
			// add object to search index
			_, err = ps.saveIndex(id)
		} else {
			// remove object from search index
			_, err = ps.dropIndex(id)
		}

		return err
	} else {
		return nil
	}
}

// GetPostMetadataById returns Post metadata by given Raw ID
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

// GetPostById returns Post document by given Raw ID
func (ps *PostService) GetPostById(rawid string) (*models.PostObjectView, error) {
	var (
		doc    *models.PostObjectView
		oid    primitive.ObjectID
		filter primitive.D
		err    error
	)

	// Parse ObjectId
	if oid, err = primitive.ObjectIDFromHex(rawid); err != nil {
		return nil, err
	} else {
		filter = bson.D{{Key: "_id", Value: oid}}
	}

	// Find single post document by _id
	if err = db.
		MongoDb().
		Collection("postsObject").
		FindOne(ps.Ctx, filter).
		Decode(&doc); err != nil {
		return nil, err
	} else {
		return doc, err
	}
}

// GetPostsMetadata returns list of metadata of posts
// Need to provide result limit (default is 0)
func (ps *PostService) GetPostsMetadata(params *types.GetPostsMetadataOptions) ([]models.PostMetadataDocument, error) {
	var (
		posts     []models.PostMetadataDocument
		findOpts  = options.Find()
		collName  string
		sortProp  string
		sortOrder int
		filter    = bson.D{}
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

		if err = cur.Close(ps.Ctx); err != nil {
			return nil, err
		}
	}

	return posts, nil
}

// GetPostCount returns number of post (public + private) in posts collection
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

// CreatePost inserts new post document into posts collection
func (ps *PostService) CreatePost(payload *types.CreatePostPayload) (*mongo.InsertOneResult, error) {
	var (
		authorId primitive.ObjectID
		res      *mongo.InsertOneResult
		err      error
	)

	if authorId, err = primitive.ObjectIDFromHex(payload.AuthorId); err != nil {
		return nil, err
	}

	newPost := bson.M{
		"title":    payload.Title,
		"slug":     payload.Slug,
		"desc":     payload.Desc,
		"topic":    payload.Topic,
		"tags":     payload.Tags,
		"body":     payload.Body,
		"format":   payload.Format,
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
		"related":   []string{},
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	}

	if res, err = db.MongoDb().Collection("posts").InsertOne(ps.Ctx, newPost); err != nil {
		util.Log.Error(fmt.Sprintf("[PostService.CreatePost] %s", err.Error()))
		return nil, err
	} else {
		util.Log.Info(fmt.Sprintf("[PostService.CreatePost] Inserted post document (id='%s')", res.InsertedID.(primitive.ObjectID).Hex()))
	}

	// add object to search index
	if err = ps.updateIndex(res.InsertedID.(primitive.ObjectID), payload.Public); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdatePostById updates existing post document by given rawid
func (ps *PostService) UpdatePostById(rawid string, payload types.UpdatePostPayload) (*mongo.UpdateResult, error) {
	var (
		oid    primitive.ObjectID
		filter bson.D
		update bson.D
		res    *mongo.UpdateResult
		err    error
	)

	if oid, err = primitive.ObjectIDFromHex(rawid); err != nil {
		return nil, err
	} else {
		filter = bson.D{{Key: "_id", Value: oid}}
		update = bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "title", Value: payload.Title},
				{Key: "slug", Value: payload.Slug},
				{Key: "desc", Value: payload.Desc},
				{Key: "topic", Value: payload.Topic},
				{Key: "tags", Value: payload.Tags},
				{Key: "body", Value: payload.Body},
				{Key: "public", Value: payload.Public},
				{Key: "coverImage", Value: &models.PostCoverImage{
					Id:      payload.CoverImageId,
					Path:    payload.CoverImagePath,
					RefName: payload.CoverImageRefName,
					RefUrl:  payload.CoverImageRefUrl,
				}},
			}},
		}
	}

	// update document
	if res, err = db.MongoDb().Collection("posts").UpdateOne(ps.Ctx, filter, update); err != nil {
		util.Log.Error(fmt.Sprintf("[PostService.UpdatePostById] %s", err.Error()))
		return nil, err
	} else {
		util.Log.Info(fmt.Sprintf("[PostService.UpdatePostById] Updated post document (id='%s')", oid.Hex()))
	}

	// update search index
	if err = ps.updateIndex(oid, payload.Public); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdatePostScope set specified scope of post
func (ps *PostService) UpdatePostScope(rawid string, scope string) error {
	var (
		oid    primitive.ObjectID
		err    error
		filter bson.D
		update bson.M
		public = scope == "public"
	)

	if oid, err = primitive.ObjectIDFromHex(rawid); err != nil {
		return err
	} else {
		filter = bson.D{{Key: "_id", Value: oid}}
		update = bson.M{"$set": bson.M{"public": public}}
	}

	// update scope
	if _, err = db.MongoDb().Collection("posts").UpdateOne(ps.Ctx, filter, update); err != nil {
		util.Log.Error(fmt.Sprintf("[PostService.UpdatePostScope] %s", err.Error()))
		return err
	} else {
		util.Log.Info(fmt.Sprintf("[PostService.UpdatePostScope] Updated post scope (id='%s', scope='%s')", oid.Hex(), scope))
	}

	// update search index
	return ps.updateIndex(oid, public)
}

// SetPostDeleteFlag sets post delete flag by given post rawid
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

	// update document
	if _, err = db.MongoDb().Collection("posts").UpdateOne(ps.Ctx, filter, update); err != nil {
		util.Log.Error(fmt.Sprintf("[PostService.SetPostDeleteFlag] %s", err.Error()))
		return err
	} else {
		util.Log.Info(fmt.Sprintf("[PostService.SetPostDeleteFlag] Updated post delete flag (id='%s', value=%t)", oid.Hex(), value))
	}

	// update search index
	return ps.updateIndex(oid, !value)
}

// UploadPostCoverImage uploads cover image and returns uploaded file response
func (ps *PostService) UploadPostCoverImage(imageData []byte) (res *types.UploadedImageFile, err error) {
	return util.UploadImage(ps.Ctx, "fivemin-prod/post-cover-images", imageData)
}

// UploadPostEmbedImage uploads post embedded image and returns uploaded file response
func (ps *PostService) UploadPostEmbedImage(imageData []byte) (res *types.UploadedImageFile, err error) {
	return util.UploadImage(ps.Ctx, "fivemin-prod/post-images", imageData)
}

// DeletePostImage removes post related image from storage bucket
func (ps *PostService) DeletePostImage(publicId string) error {
	return util.DeleteImage(ps.Ctx, publicId)
}

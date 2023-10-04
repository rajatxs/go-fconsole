package services

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/rajatxs/go-fconsole/db"
	"github.com/rajatxs/go-fconsole/models"
	"github.com/rajatxs/go-fconsole/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostService struct {
	Ctx context.Context
}

// NewPostService creates new instance of PostService
func NewPostService() *PostService {
	return &PostService{
		Ctx: nil,
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
func (ps *PostService) GetPostById(rawid string) (*models.PostDocument, error) {
	var (
		doc    *models.PostDocument
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
		Collection("posts").
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
		err      error
		result   *mongo.InsertOneResult
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

	result, err = db.MongoDb().Collection("posts").InsertOne(ps.Ctx, newPost)
	return result, err
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

	// Update scope
	_, err = db.MongoDb().Collection("posts").UpdateOne(ps.Ctx, filter, update)
	return err
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

	// Update document
	_, err = db.MongoDb().Collection("posts").UpdateOne(ps.Ctx, filter, update)
	return err
}

// UploadPostImage uploads post related image and returns uploaded file response
func (ps *PostService) UploadPostImage(folderName string, imageData []byte) (res *types.PostImageFile, err error) {
	var (
		uploadResult *uploader.UploadResult
		file         = bytes.NewReader(imageData)
		params       = uploader.UploadParams{
			ResourceType: "image",
			Folder:       fmt.Sprintf("fivemin-prod/%s", folderName),
		}
	)

	if uploadResult, err = CloudinaryInstance().Upload.Upload(ps.Ctx, file, params); err != nil {
		return nil, err
	} else {
		res = &types.PostImageFile{
			PublicId: uploadResult.PublicID,
			AssetId:  uploadResult.AssetID,
			Format:   uploadResult.Format,
		}
		return res, nil
	}
}

// UploadPostCoverImage uploads cover image and returns uploaded file response
func (ps *PostService) UploadPostCoverImage(imageData []byte) (res *types.PostImageFile, err error) {
	return ps.UploadPostImage("post-cover-images", imageData)
}

// UploadPostEmbedImage uploads post embedded image and returns uploaded file response
func (ps *PostService) UploadPostEmbedImage(imageData []byte) (res *types.PostImageFile, err error) {
	return ps.UploadPostImage("post-images", imageData)
}

// DeletePostImage removes post related image from storage bucket
func (ps *PostService) DeletePostImage(publicId string) (res *admin.DeleteAssetsResult, err error) {
	return CloudinaryInstance().Admin.DeleteAssets(ps.Ctx, admin.DeleteAssetsParams{
		PublicIDs: []string{publicId},
	})
}

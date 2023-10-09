package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostCoverImage struct {
	Id      string `bson:"id" json:"id"`
	Path    string `bson:"path" json:"path"`
	RefName string `bson:"refName" json:"refName"`
	RefUrl  string `bson:"refUrl" json:"refUrl"`
}

type PostMetadataDocument struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	Title      string             `bson:"title" json:"title"`
	Slug       string             `bson:"slug" json:"slug"`
	Desc       string             `bson:"desc" json:"desc"`
	Tags       []string           `bson:"tags" json:"tags"`
	Topic      string             `bson:"topic" json:"topic"`
	Stars      int64              `bson:"stars" json:"stars"`
	Format     string             `bson:"format" json:"format"`
	AuthorId   primitive.ObjectID `bson:"authorId" json:"authorId"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
	CoverImage *PostCoverImage    `bson:"coverImage" json:"coverImage"`
}

type PostDocument struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	Title      string             `bson:"title" json:"title"`
	Slug       string             `bson:"slug" json:"slug"`
	Desc       string             `bson:"desc" json:"desc"`
	Tags       []string           `bson:"tags" json:"tags"`
	Topic      string             `bson:"topic" json:"topic"`
	Body       bson.M             `bson:"body" json:"body"`
	Format     string             `bson:"format" json:"format"`
	Stars      int64              `bson:"stars" json:"stars"`
	Public     bool               `bson:"public" json:"public"`
	Deleted    bool               `bson:"deleted" json:"deleted"`
	CoverImage *PostCoverImage    `bson:"coverImage" json:"coverImage"`
	AuthorId   primitive.ObjectID `bson:"authorId" json:"authorId"`
	Related    []string           `bson:"related" json:"related"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type PostIndex struct {
	ObjectId  string    `json:"objectID"`
	Name      string    `json:"name"`
	Topic     string    `json:"topic"`
	Desc      string    `json:"description"`
	Tags      []string  `json:"tags"`
	Url       string    `json:"url"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

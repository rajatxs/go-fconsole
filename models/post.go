package models

import (
	"time"

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
	AuthorId   primitive.ObjectID `bson:"authorId" json:"authorId"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
	CoverImage *PostCoverImage    `bson:"coverImage" json:"coverImage"`
}

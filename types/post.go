package types

import "go.mongodb.org/mongo-driver/bson"

type GetPostsMetadataOptions struct {
	Private bool   `json:"private"`
	Topic   string `json:"topic"`
	SortBy  string `json:"sortBy"`
	Limit   int64  `json:"limit"`
	Skip    int64  `json:"skip"`
}

type CreatePostPayload struct {
	Title             string   `json:"title"`
	Slug              string   `json:"slug"`
	Desc              string   `json:"desc"`
	Tags              []string `json:"tags"`
	Topic             string   `json:"topic"`
	Body              bson.M   `json:"body"`
	Format            string   `json:"format"`
	Public            bool     `json:"public"`
	CoverImageId      string   `json:"coverImageId"`
	CoverImagePath    string   `json:"coverImagePath"`
	CoverImageRefName string   `json:"coverImageRefName"`
	CoverImageRefUrl  string   `json:"coverImageRefUrl"`
	AuthorId          string   `json:"authorId"`
	License           string   `json:"license"`
	RelatedPosts      []string `json:"relatedPosts"`
}

type UpdatePostPayload struct {
	Title             string   `json:"title"`
	Slug              string   `json:"slug"`
	Desc              string   `json:"desc"`
	Tags              []string `json:"tags"`
	Topic             string   `json:"topic"`
	Body              bson.M   `json:"body"`
	Public            bool     `json:"public"`
	CoverImageId      string   `json:"coverImageId"`
	CoverImagePath    string   `json:"coverImagePath"`
	CoverImageRefName string   `json:"coverImageRefName"`
	CoverImageRefUrl  string   `json:"coverImageRefUrl"`
	License           string   `json:"license"`
	RelatedPosts      []string `json:"relatedPosts"`
}

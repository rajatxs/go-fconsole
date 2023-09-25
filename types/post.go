package types

type GetPostsMetadataOptions struct {
	Private bool   `json:"private"`
	Topic   string `json:"topic"`
	SortBy  string `json:"sortBy"`
	Limit   int64  `json:"limit"`
	Skip    int64  `json:"skip"`
}

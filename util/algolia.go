package util

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/rajatxs/go-fconsole/config"
)

var (
	client    *search.Client
	postIndex *search.Index
)

// PostIndex returns "posts" index reference
func PostIndex() *search.Index {
	return postIndex
}

// InitAlgolia creates new client instance with default config
func InitAlgolia() {
	client = search.NewClient(config.AlgoliaAppId(), config.AlgoliaApiKey())
	postIndex = client.InitIndex("posts")
}

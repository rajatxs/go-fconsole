package main

import (
	"context"
	"log"
	"os/user"
	"runtime"

	"github.com/rajatxs/go-fconsole/config"
	"github.com/rajatxs/go-fconsole/db"
	"github.com/rajatxs/go-fconsole/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	var err error

	a.ctx = ctx
	if err = db.ConnectMongoDb(ctx); err != nil {
		log.Fatal(err)
	}
}

// terminate is called when the app shutdown.
func (a *App) terminate(ctx context.Context) {
	var err error

	if err = db.DisconnectMongoDb(ctx); err != nil {
		log.Println(err)
	}
}

func (a *App) GetAppConfigVariables() map[string]string {
	var env = map[string]string{}

	env["ENV"] = config.Env()
	env["ADMIN_ID"] = config.AdminId()
	env["CLOUDINARY_ID"] = config.CloudinaryId()
	return env
}

func (a *App) GetVersions() map[string]string {
	var ver = map[string]string{}

	currentUser, err := user.Current()
	if err != nil {
		log.Println(err)
	}

	ver["app"] = "0.0.1"
	ver["date"] = "2023-09-20T07:53:14.501Z"
	ver["wails"] = "2.6.0"
	ver["go"] = "1.20.3"

	// TODO: Write "WebView2 Runtime Version" in UI
	ver["webview2"] = "111.0.2045.32"
	ver["os"] = runtime.GOOS
	ver["arch"] = runtime.GOARCH
	ver["uname"] = currentUser.Name
	ver["homedir"] = currentUser.HomeDir
	return ver
}

func (a *App) GetPostsMetadata(scope string, topic string, sortBy string, limit int64, skip int64) ([]models.PostMetadataDocument, error) {
	var posts []models.PostMetadataDocument
	var collName string
	var filter primitive.D = bson.D{}

	opts := options.Find()

	if scope == "public" {
		collName = "publicPostsMetadata"
	} else {
		collName = "privatePostsMetadata"
	}

	if topic != "all" {
		filter = bson.D{{Key: "topic", Value: topic}}
	}

	switch sortBy {
	case "title":
		opts.SetSort(map[string]int{"title": 1})

	case "topic":
		opts.SetSort(map[string]int{"topic": 1})

	case "newest":
		opts.SetSort(map[string]int{"createdAt": -1})

	case "oldest":
		opts.SetSort(map[string]int{"createdAt": 1})

	case "updated":
		opts.SetSort(map[string]int{"updatedAt": -1})
	}

	opts.SetLimit(limit)
	opts.SetSkip(skip)

	if cur, err := db.MongoDb().Collection(collName).Find(a.ctx, filter, nil, opts); err != nil {
		log.Println(err)
	} else {
		defer cur.Close(a.ctx)

		for cur.Next(a.ctx) {
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

func (a *App) GetPostCount(scope string, includeDeleted bool) (int64, error) {
	return db.MongoDb().Collection("posts").CountDocuments(
		a.ctx,
		bson.D{{Key: "public", Value: scope == "public"}, {Key: "deleted", Value: includeDeleted}})
}

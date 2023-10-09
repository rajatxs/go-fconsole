package config

import (
	"os"
)

func Env() string {
	if os.Getenv("FMC_ENV") == "production" {
		return "prod"
	} else {
		return "dev"
	}
}

func IsProd() bool {
	return Env() == "prod"
}

func MongoDbConnectionUrl() string {
	return os.Getenv("FMC_MONGODB_CONN_URL")
}

func MongoDbName() string {
	return os.Getenv("FMC_MONGODB_NAME")
}

func AdminId() string {
	return os.Getenv("FMC_ADMIN_ID")
}

func CloudinaryId() string {
	return os.Getenv("FMC_CLOUDINARY_ID")
}

func CloudinaryURL() string {
	return os.Getenv("CLOUDINARY_URL")
}

func AlgoliaAppId() string {
	return os.Getenv("FMC_ALGOLIA_APP_ID")
}

func AlgoliaApiKey() string {
	return os.Getenv("FMC_ALGOLIA_API_KEY")
}

func ClientUrl() string {
	return os.Getenv("FMC_CLIENT_URL")
}

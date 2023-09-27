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

package config

import (
	"os"
)

/* Platform environment */
func Env() string {
	if os.Getenv("FMC_ENV") == "production" {
		return "prod"
	} else {
		return "dev"
	}
}

/* MongoDB Connection URL */
func MongoDbConnectionUrl() string {
	return os.Getenv("FMC_MONGODB_CONN_URL")
}

/* MongoDB Database Name */
func MongoDbName() string {
	return os.Getenv("FMC_MONGODB_NAME")
}

/* Console Admin ID */
func AdminId() string {
	return os.Getenv("FMC_ADMIN_ID")
}

/* Cloudinary ID */
func CloudinaryId() string {
	return os.Getenv("FMC_CLOUDINARY_ID")
}

package services

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

var cld *cloudinary.Cloudinary

// InitCloudinary creates new instance of cloudinary with default configuration
func InitCloudinary() (err error) {
	if cld, err = cloudinary.New(); err != nil {
		return err
	} else {
		cld.Config.URL.Secure = true
		return nil
	}
}

// CloudinaryInstance returns active instance of cloudinary
func CloudinaryInstance() *cloudinary.Cloudinary {
	return cld
}

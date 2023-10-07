package util

import (
	"bytes"
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/rajatxs/go-fconsole/types"
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

// UploadImage uploads raw image and returns uploaded file response
func UploadImage(ctx context.Context, folderName string, imageData []byte) (res *types.UploadedImageFile, err error) {
	var (
		uploadResult *uploader.UploadResult
		file         = bytes.NewReader(imageData)
		params       = uploader.UploadParams{
			ResourceType: "image",
			Folder:       folderName,
		}
	)

	if uploadResult, err = CloudinaryInstance().Upload.Upload(ctx, file, params); err != nil {
		return nil, err
	} else {
		res = &types.UploadedImageFile{
			PublicId: uploadResult.PublicID,
			AssetId:  uploadResult.AssetID,
			Format:   uploadResult.Format,
		}
		return res, nil
	}
}

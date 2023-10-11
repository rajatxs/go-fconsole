package util

import (
	"bytes"
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
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

	Log.Info(fmt.Sprintf("[util.UploadImage] Uploading image (folder='%s')", folderName))

	if uploadResult, err = CloudinaryInstance().Upload.Upload(ctx, file, params); err != nil {
		Log.Error(fmt.Sprintf("[util.UploadImage] %s", err.Error()))
		return nil, err
	} else {
		res = &types.UploadedImageFile{
			PublicId: uploadResult.PublicID,
			AssetId:  uploadResult.AssetID,
			Format:   uploadResult.Format,
		}
		Log.Info(fmt.Sprintf(
			"[util.UploadImage] Image uploaded (format='%s', publicId='%s', assetId='%s')",
			res.Format,
			res.PublicId,
			res.AssetId))

		return res, nil
	}
}

// DeleteImage removes image from storage bucket
func DeleteImage(ctx context.Context, publicId string) (err error) {
	if _, err = CloudinaryInstance().Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
		PublicIDs: []string{publicId},
	}); err != nil {
		Log.Error(fmt.Sprintf("[util.DeleteImage] %s", err.Error()))
	} else {
		Log.Info(fmt.Sprintf("[util.DeleteImage] Image deleted (publicId='%s')", publicId))
	}

	return err
}

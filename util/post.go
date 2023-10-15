package util

import (
	"fmt"

	"github.com/rajatxs/go-fconsole/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetPostCoverImageUrl returns absolute url of post cover image
func GetPostCoverImageUrl(path string) string {
	return fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/c_scale,h_600/%s.webp", config.CloudinaryId(), path)
}

// ParsePostIds returns list of object id from raw post id
func ParsePostIds(ids []string) (oids []primitive.ObjectID, err error) {
	if len(ids) > 0 {
		var rpoid primitive.ObjectID

		for _, rpid := range ids {
			if rpoid, err = primitive.ObjectIDFromHex(rpid); err != nil {
				return nil, err
			} else {
				oids = append(oids, rpoid)
			}
		}
	} else {
		oids = []primitive.ObjectID{}
	}

	return oids, nil
}

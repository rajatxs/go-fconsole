package util

import (
	"fmt"

	"github.com/rajatxs/go-fconsole/config"
)

// GetPostCoverImageUrl returns absolute url of post cover image
func GetPostCoverImageUrl(path string) string {
	return fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/c_scale,h_600/%s.webp", config.CloudinaryId(), path)
}

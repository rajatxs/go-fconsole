package types

type AppPublicConfigVariables struct {
	ENV           string `json:"ENV"`
	ADMIN_ID      string `json:"ADMIN_ID"`
	CLOUDINARY_ID string `json:"CLOUDINARY_ID"`
}

type AppVersions struct {
	App      string `json:"app"`
	Date     string `json:"date"`
	Wails    string `json:"wails"`
	Go       string `json:"go"`
	WebView2 string `json:"webview2"`
	Os       string `json:"os"`
	Arch     string `json:"arch"`
	Username string `json:"uname"`
	HomeDir  string `json:"homedir"`
}

type UploadedImageFile struct {
	PublicId string `json:"publicId"`
	AssetId  string `json:"assetId"`
	Format   string `json:"format"`
}

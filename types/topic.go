package types

type Topic struct {
	Name      string `json:"name"`
	ThumbId   string `json:"thumbId"`
	ThumbPath string `json:"thumbPath"`
	Public    bool   `json:"public"`
}

type Topics = map[string]Topic

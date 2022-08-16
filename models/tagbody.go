package models

type TagBody struct {
	Tags *struct {
		Tag []Tag `json:"tag,omitempty"`
	} `json:"tags,omitempty"`
}

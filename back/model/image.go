package model

type ImageInfo struct {
	ImageURL string   `json:"pic"`
	Kinds    []string `json:"tags"`
}

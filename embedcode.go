package keezmovies

type KeezmoviesEmbedCode struct {
	Embed KeezmoviesCode `json:"video,omitempty"`
}

type KeezmoviesCode struct {
	Code string `json:"embed_code,omitempty"`
}

package keezmovies

type KeezmoviesVideo struct {
	ID          float64 `json:"video_id,omitempty"`
	Duration    string  `json:"duration,omitempty"`
	Views       float64 `json:"times_viewed,omitempty"`
	Rating      float64 `json:"rating,omitempty"`
	Ratings     float64 `json:"ratings,omitempty"`
	Title       string  `json:"title,omitempty"`
	URL         string  `json:"video_url,omitempty"`
	Thumb       string  `json:"image_url,omitempty"`
	PublishDate string  `json:"publish_date,omitempty"`
}

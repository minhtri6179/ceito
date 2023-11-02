package util

type Image struct {
	Id     int    `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Cloud  string `json:"cloud"`
}

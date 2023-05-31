package model

type Post struct{
	Id string `json:"id"`
	Usr string `json:"user"`
	Message string `json:"Message"`
	Url string `json:"url"`
	Type string `json:"type"`
}
package model

import "sync"

type Data struct{
	Id 		int `json:"id"`
	Title 	string `json:"title"`
	Author 	string `json:"author"`
}

type UpdateData struct{
	Title *string `json:"title"`
	Author *string `json:"author"`
}

var (
	Books []Data
	Mu sync.Mutex
)
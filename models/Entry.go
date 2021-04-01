package models

import "time"

type Entry struct {
	Name      string     `json:"name"`
	Date      time.Time  `json:"time"`
	Metadatas []Metadata `json:"metadatas"`
}

type Metadata struct {
	WindowName string    `json:"windowName"`
	Url        string    `json:"url"`
	Pid        int32     `json:"pid"`
	Date       time.Time `json:"date"`
	Duration   int32     `json:"duration"`
}

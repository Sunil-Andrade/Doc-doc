package models

import "time"

type Document struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	FilePath  string    `json:"filePath"`
	CreatedAt time.Time `json:"createdAt"`
}
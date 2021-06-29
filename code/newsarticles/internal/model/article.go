package model

import uuid "github.com/satori/go.uuid"

type Articles struct {
	UserID   uuid.UUID `json:"user_id"`
	Articles []Article `json:"articles"`
}

type Article struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Like     bool   `json:"like"`
}

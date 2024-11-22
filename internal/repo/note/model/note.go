package model

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        int64        `json:"id" bson:"id"`
	Title     string       `json:"title" bson:"title"`
	Body      string       `json:"body" bson:"body"`
	CreatedAt time.Time    `json:"created_at" bson:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" bson:"updated_at"`
	//    Image string `json: "image"`
}

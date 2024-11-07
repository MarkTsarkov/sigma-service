package model

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        int64        `json:"id"`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	//    Image string `json: "image"`
}

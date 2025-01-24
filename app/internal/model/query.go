package model

import "time"

type Query struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Response  string    `json:"response"`
	CreatedAt time.Time `json:"created_at"`
}

type QueryWithRawResponse struct {
	ID        int                    `json:"id"`
	Text      string                 `json:"text"`
	Response  map[string]interface{} `json:"response"`
	CreatedAt time.Time              `json:"created_at"`
}

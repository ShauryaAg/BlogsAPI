package models

import (
	"time"
)

type Blog struct {
	ID       string    `json:"ID"`
	Title    string    `json:"title"`
	Author   string    `json:"Author"`
	Content  string    `json:"Content"`
	DateTime time.Time `json:"DateTime"`
}

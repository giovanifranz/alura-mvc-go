package entities

import (
	"github.com/google/uuid"
)

type Tweet struct {
	ID          string `json:"id"`
	Description string `json:"description`
}

func NewTweet() *Tweet {
	return &Tweet{
		ID: uuid.New().String(),
	}
}

package tweet

import (
	"github.com/dghubble/go-twitter/twitter"
)

type Deleter struct {
	Client *twitter.Client
}

func NewDeleter(client *twitter.Client) *Deleter {
	return &Deleter{client}
}

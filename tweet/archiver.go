package tweet

import (
	"github.com/dghubble/go-twitter/twitter"
)

type Archiver struct {
	Client *twitter.Client
}

func NewArchiver(client *twitter.Client) *Archiver {
	return &Archiver{client}
}


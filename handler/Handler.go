package handler

import (
	"github.com/dghubble/go-twitter/twitter"
)

type Handler interface {
	HandleTweet(*twitter.Tweet)
}

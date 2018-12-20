package handler

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/adylanrff/Electics/model"
	"github.com/dghubble/go-twitter/twitter"
	kafka "github.com/segmentio/kafka-go"
)

type KafkaProducerHandler struct {
	KafkaProducer *kafka.Writer
}

func (kph *KafkaProducerHandler) HandleTweet(tweet *twitter.Tweet) {
	var text string
	var timestamp time.Time
	var coordinates [2]float64

	if tweet.ExtendedTweet != nil {
		text = tweet.ExtendedTweet.FullText
	} else {
		text = tweet.Text
	}

	timestamp, err := tweet.CreatedAtTime()
	if err != nil {
		log.Fatalln("Cannot parse tweet timestamp")
	}

	if tweet.Coordinates != nil {
		coordinates = tweet.Coordinates.Coordinates
	}

	tweetMsg := model.Tweet{
		TweetID:       tweet.ID,
		Timestamp:     timestamp,
		RetweetCount:  tweet.RetweetCount,
		FavoriteCount: tweet.FavoriteCount,
		Text:          text,
		Coordinates:   coordinates,
		Username:      tweet.User.Name,
	}

	tweetMsgJSON, err := json.Marshal(tweetMsg)

	if err != nil {
		log.Fatalln("Tweet JSON marshal failed")
	}

	kph.KafkaProducer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(string(tweetMsg.TweetID)),
			Value: tweetMsgJSON,
		})

	log.Printf("Sent tweet %d", tweetMsg.TweetID)
}

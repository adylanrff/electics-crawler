package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adylanrff/Electics/config"
	"github.com/adylanrff/Electics/handler"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/go-ini/ini"
	"github.com/segmentio/kafka-go"
)

func initTwitter(twitterConfig config.TwitterConfig, locationConfig config.LocationConfig) *twitter.Stream {
	oauthConfig := oauth1.NewConfig(twitterConfig.ConsumerKey, twitterConfig.ConsumerSecret)
	token := oauth1.NewToken(twitterConfig.AccessToken, twitterConfig.AccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	params := &twitter.StreamFilterParams{
		Locations:     locationConfig.Locations(),
		StallWarnings: twitter.Bool(true),
	}

	stream, err := client.Streams.Filter(params)

	if err != nil {
		log.Fatalf("Error : %s", err)
		os.Exit(2)
	}

	return stream
}

func main() {
	// Load config file
	if len(os.Args) < 2 {
		fmt.Println("Insufficient argument")
		fmt.Println("Usage <configpath>")
		os.Exit(3)
	}

	cfg, err := ini.Load(os.Args[1])
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	twitterConfig := config.NewTwitterConfig(cfg)
	locationConfig := config.NewLocationConfig(cfg)

	stream := initTwitter(twitterConfig, locationConfig)

	kafkaProducer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{cfg.Section("kafka").Key("broker").String(), "localhost:9093"},
		Topic:    cfg.Section("kafka").Key("topic").String(),
		Balancer: &kafka.LeastBytes{},
	})

	producerHandler := handler.KafkaProducerHandler{KafkaProducer: kafkaProducer}

	demux := twitter.NewSwitchDemux()
	demux.Tweet = producerHandler.HandleTweet
	for message := range stream.Messages {
		demux.Handle(message)
	}
}

package config

import (
	"github.com/go-ini/ini"
)

type TwitterConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func NewTwitterConfig(cfg *ini.File) TwitterConfig {
	// Load credentials
	consumerKey := cfg.Section("twitterAPI").Key("consumer_key").String()
	consumerSecret := cfg.Section("twitterAPI").Key("consumer_secret").String()
	accessToken := cfg.Section("twitterAPI").Key("access_token").String()
	accessSecret := cfg.Section("twitterAPI").Key("access_secret").String()

	twitterConfig := TwitterConfig{consumerKey, consumerSecret, accessToken, accessSecret}
	return twitterConfig
}

package main

import (
	"fmt"
	"io"

	"github.com/bunsenmcdubbs/tweeter-deleter/tweet"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		panic(fmt.Errorf("unable to parse appConfig: %w", err))
	}

	client, err := tweet.AuthorizeOOB(cfg.ConsumerKey, cfg.ConsumerSecret)
	if err != nil {
		panic(err)
	}

	path := "https://api.twitter.com/1.1/statuses/home_timeline.json?count=2"
	resp, _ := client.Get(path)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Raw Response Body:\n%v\n", string(body))
}

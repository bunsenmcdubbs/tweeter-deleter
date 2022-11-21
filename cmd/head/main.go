package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/bunsenmcdubbs/tweeter-deleter/cmd"
	"github.com/bunsenmcdubbs/tweeter-deleter/tweet"
)

func main() {
	cfg, err := cmd.ParseConfig()
	if err != nil {
		panic(fmt.Errorf("unable to parse appConfig: %w", err))
	}

	client, err := tweet.AuthorizeOOBInteractive(cfg.ConsumerKey, cfg.ConsumerSecret)
	if err != nil {
		panic(err)
	}

	// https://developer.twitter.com/en/docs/twitter-api/v1/tweets/timelines/api-reference/get-statuses-user_timeline
	path := "https://api.twitter.com/1.1/statuses/user_timeline.json?count=2&trim_user=1&include_rts=1"
	resp, _ := client.Get(path)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var pretty bytes.Buffer
	_ = json.Indent(&pretty, body, "", "  ")
	fmt.Printf("Raw Response Body:\n%s\n", pretty.String())
}

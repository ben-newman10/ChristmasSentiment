package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/marcusolsson/tui-go"
	"github.com/nlopes/slack"
	"github.com/chrisgeidi/go-twitter/twitter"
	"github.com/thoas/go-funk"
	"github.com/unifon/go-sentiment"
)

func main() {
	// Replace these with your own Twitter API keys and secrets
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	// Initialize the Twitter client
	client, err := twitter.NewClient(consumerKey, consumerSecret, accessToken, accessSecret)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the sentiment analysis engine
	engine, err := sentiment.New()
	if err != nil {
		log.Fatal(err)
	}

	// Search for tweets that contain the keywords "Christmas" and "Snow"
	searchParams := &twitter.SearchTweetParams{
		Query: "Christmas Snow",
	}
	searchResult, _, err := client.Search.Tweets(searchParams)
	if err != nil {
		log.Fatal(err)
	}

	// Go through the tweets and perform sentiment analysis on each one
	for _, tweet := range searchResult.Statuses {
		analysis, err := engine.Analyze(tweet.Text)
		if err != nil {
			log.Fatal(err)
		}

		// If the sentiment is negative, print the tweet to the console
		if analysis.Score < 0 {
			fmt.Println(tweet.Text)
		}
	}
}

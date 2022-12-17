package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cdipaolo/sentiment"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {

	// Set the timer to run every 10 seconds
	duration := time.Duration(10) * time.Second
	timer := time.NewTicker(duration)

	// Read API key and secret from environment variables
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	// Create an OAuth1 config
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	// Create an OAuth1 token
	token := oauth1.NewToken(accessToken, accessSecret)

	// Create a Twitter client
	client := twitter.NewClient(config.Client(oauth1.NoContext, token))

	for {
		// Search for tweets containing the keywords "Christmas" and "snow"
		search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
			Query: "Christmas snow",
		})
		if err != nil {
			fmt.Println("Error searching tweets:", err)
			return
		}

		// Create a new sentiment analyzer
		model, err := sentiment.Restore()
		if err != nil {
			fmt.Println("Error creating sentiment analyzer:", err)
			return
		}

		// Iterate through the search results and classify the sentiment of each tweet
		for _, tweet := range search.Statuses {

			// fmt.Println(tweet.Text)

			// Use the sentiment analyzer to classify the sentiment of the tweet
			score := model.SentimentAnalysis(tweet.Text, sentiment.English).Score

			// If the sentiment is negative, print the tweet to the console
			if score == 0 {
				fmt.Println(tweet.Text)
			}
		}

		// Wait for the timer to expire before searching for more tweets
		<-timer.C
	}
}

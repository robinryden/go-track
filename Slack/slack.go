package Slack

import (
	"log"
	"os"

	slackHandler "github.com/ashwanthkumar/slack-go-webhook"
	"github.com/joho/godotenv"
)

var (
	webhookURL = os.Getenv("SLACK_WEBHOOK_URL")
)

func init() {
	godotenv.Load()
}

func Post(message string) {
	payload := slackHandler.Payload{
		Text:      message,
		Username:  "go-track",
		Channel:   "#gotrack",
		IconEmoji: ":monkas:",
	}

	err := slackHandler.Send(webhookURL, "", payload)
	if err != nil {
		log.Fatal(err)
	}

}

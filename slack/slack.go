package Slack

import (
	"log"
	"os"

	slackHandler "github.com/ashwanthkumar/slack-go-webhook"
)

func Post(message string) {
	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
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

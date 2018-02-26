package Slack

import (
	"log"

	slackHandler "github.com/ashwanthkumar/slack-go-webhook"
)

var (
	webhookURL = "https://hooks.slack.com/services/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"
)

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

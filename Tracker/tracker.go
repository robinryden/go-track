package Tracker

import (
	"fmt"
	"go-track/Slack"
	"log"
	"net/http"
	"time"
)

type TrackedURL struct {
	name       string
	statusCode int
	timestamp  time.Time
}

var (
	Tracked = make(chan *TrackedURL)
)

func Start() {
	for {
		select {
		case trackedURL := <-Tracked:
			go HealthCheck(trackedURL)
			fmt.Printf("Tracked URL %v with statusCode %v\n", trackedURL.name, trackedURL.statusCode)
		}
	}
}

// TrackURL ...
func TrackURL(url string) {
	ping, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer ping.Body.Close()

	Tracked <- &TrackedURL{
		url,
		ping.StatusCode,
		time.Now(),
	}
}

func HealthCheck(url *TrackedURL) {
	fmt.Printf("Checking health on URL: %v \n", url.name)

	switch url.statusCode {
	case 200:
		go Slack.Post(fmt.Sprintf("%s resulted in a %d at %s\n", url.name, url.statusCode, url.timestamp.Format("2006-01-02 15:04:05")))
	case 403:
		go Slack.Post(fmt.Sprintf("%s resulted in a %d at %s\n", url.name, url.statusCode, url.timestamp.Format("2006-01-02 15:04:05")))
	case 404:
		go Slack.Post(fmt.Sprintf("%s resulted in a %d at %s\n", url.name, url.statusCode, url.timestamp.Format("2006-01-02 15:04:05")))
	case 500:
		go Slack.Post(fmt.Sprintf("%s resulted in a %d at %s\n", url.name, url.statusCode, url.timestamp.Format("2006-01-02 15:04:05")))
	}
}

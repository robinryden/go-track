package tracker

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/robinryden/go-track/redis"
	"github.com/robinryden/go-track/slack"
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
	// log url check to redis
	go redis.Logger(url.name, url.statusCode, url.timestamp)

	switch url.statusCode {
	case 403:
		fallthrough
	case 404:
		fallthrough
	case 500:
		go slack.Post(fmt.Sprintf("%s resulted in a %d at %s\n",
			url.name,
			url.statusCode,
			url.timestamp.Format("2006-01-02 15:04:05"),
		))
	}
}

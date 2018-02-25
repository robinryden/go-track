package Tracker

import (
	"fmt"
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
		fmt.Println(err)
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
		fmt.Printf("The site %v is alive", url.name)
	case 403:
		fmt.Printf("The site %v gives bad request", url.name)
	case 404:
		fmt.Printf("The site %v can not find the page", url.name)
	case 500:
		fmt.Printf("The site %v gives a internal server error", url.name)
	}
}

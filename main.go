package main

import (
	"fmt"
	"go-track/tracker"

	"github.com/robfig/cron"
	"github.com/subosito/gotenv"
)

var (
	urls = []string{"http://www.strateg.se/dajs.html", "http://www.google.se", "http://www.blocket.se"}
)

func init() {
	gotenv.Load()
}

func main() {
	go tracker.Start()

	cronWorker := cron.New()
	cronWorker.AddFunc("1/5 * * * * *", func() {
		go checkURLS()
	})

	cronWorker.Start()

	var input string
	fmt.Scanln(&input)
}

func checkURLS() {
	for _, url := range urls {
		go func(url string) {
			tracker.TrackURL(url)
		}(url)
	}
}

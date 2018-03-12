package main

import (
	"sync"

	"github.com/robinryden/go-track/tracker"

	"github.com/robfig/cron"
	"github.com/subosito/gotenv"
)

var (
	wg   sync.WaitGroup
	urls = []string{"https://www.reddit.com/", "https://www.reddit.com/r/golang/"}
)

func init() {
	gotenv.Load()
}

func main() {
	cronWorker := cron.New()

	go tracker.Start()

	wg.Add(1)
	cronWorker.AddFunc("1/5 * * * * *", func() {
		go checkURLS()
	})

	cronWorker.Start()
	wg.Wait()
}

func checkURLS() {
	for _, url := range urls {
		go func(url string) {
			wg.Add(1)
			tracker.TrackURL(url)
			wg.Done()
		}(url)
	}
}

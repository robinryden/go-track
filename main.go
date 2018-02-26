package main

import (
	"fmt"
	"go-track/Tracker"

	"github.com/subosito/gotenv"
)

var (
	urls = []string{"http://www.strateg.se/testsomintefinns.html", "http://www.google.se", "http://www.blocket.se"}
)

func init() {
	gotenv.Load()
}

func main() {
	go Tracker.Start()

	for _, url := range urls {
		go func(url string) {
			Tracker.TrackURL(url)
		}(url)
	}

	var input string
	fmt.Scanln(&input)
}

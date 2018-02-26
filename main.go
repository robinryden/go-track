package main

import (
	"fmt"
	"go-track/Tracker"
)

var (
	urls = []string{"http://www.strateg.se", "https://www.google.se", "https://www.blocket.se", "http://strateg.design", "http://www.strateg.se/monkasurl.html", "http://www.strateg.se/kontaktfelstavattyp.html"}
)

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

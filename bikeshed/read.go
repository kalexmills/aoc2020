package bikeshed

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

// Read reads the input from AOC for the provided day.
func Read(day int) io.ReadCloser {
	sessionID, ok := os.LookupEnv("SESSION_ID")
	if !ok {
		log.Fatalln("could not find SESSION_ID environment variable")
		return nil
	}
	client := http.DefaultClient
	url, err := url.Parse("https://adventofcode.com")
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln("could not add cookie to cookie jar")
		return nil
	}
	jar.SetCookies(url, []*http.Cookie{{Name: "session", Value: sessionID}})
	client.Jar = jar
	resp, err := client.Get(fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day))
	if err != nil {
		log.Fatalf("could not read day %d input", day)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("could not read day %d input: status %v", day, resp.Status)
		return nil
	}
	return resp.Body
}

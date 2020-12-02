package bikeshed

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

// Read reads the input from AOC for the provided day.
func Read(day int) (io.ReadCloser, error) {
	sessionID, ok := os.LookupEnv("SESSION_ID")
	if !ok {
		return nil, errors.New("could not find SESSION_ID environment variable")
	}
	client := http.DefaultClient
	url, err := url.Parse("https://adventofcode.com")
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	jar.SetCookies(url, []*http.Cookie{{Name: "session", Value: sessionID}})
	client.Jar = jar
	resp, err := client.Get(fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

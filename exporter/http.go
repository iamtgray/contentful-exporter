package exporter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/iamtgray/contentful-exporter/config"
	log "github.com/sirupsen/logrus"
)

// TargetResponse is a structure for a response and a target URL
type TargetResponse struct {
	Response  *Response
	TargetURL config.TargetURL
}

func asyncHTTPGets(targets []config.TargetUrl, token string) ([]*Response, error) {

	// Channels used to enable concurrent requests
	ch := make(chan *Response, len(targets))

	responses := []*Response{}

	for _, t := range targets {

		url := t.URL
		go func(url string) {
			err := getResponse(t, token, ch)
			if err != nil {
				ch <- &Response{url, nil, []byte{}, t, err}
			}
		}(url)

	}

	for {
		select {
		case r := <-ch:
			if r.err != nil {
				log.Errorf("Error scraping API, Error: %v", r.err)
				break
			}
			responses = append(responses, r)

			if len(responses) == len(targets) {
				return responses, nil
			}
		}

	}
}

// getResponse collects an individual http.response and returns a *Response
func getResponse(t config.TargetURL, token string, ch chan<- *Response) error {

	url := t.URL

	log.Infof("Fetching %s \n", url)

	resp, err := getHTTPResponse(url, token) // do this earlier
	if err != nil {
		return fmt.Errorf("Error fetching http response: %v", err)
	}
	defer resp.Body.Close()

	// Read the body to a byte array so it can be used elsewhere
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error converting body to byte array: %v", err)
	}

	// Triggers if a user specifies an invalid or not visible repository
	if resp.StatusCode == 404 {
		return fmt.Errorf("Error: Received 404 status from Github API, ensure the repsository URL is correct. If it's a privare repository, also check the oauth token is correct")
	}

	ch <- &Response{url, resp, body, t, err}

	return nil
}

// getHTTPResponse handles the http client creation, token setting and returns the *http.response
func getHTTPResponse(url string, token string) (*http.Response, error) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	// If a token is present, add it to the http.request
	if token != "" {
		req.Header.Add("Authorization", "token "+token)
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, err
}

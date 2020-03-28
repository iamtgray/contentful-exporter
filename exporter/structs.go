package exporter

import (
	"net/http"

	"github.com/infinityworks/github-exporter/config"
)

// Exporter is used to store Metrics data and embeds the config struct.
// This is done so that the relevant functions have easy access to the
// user defined runtime configuration when the Collect method is called.
type Exporter struct {
	APIMetrics map[string]PrometheusMetric
	config.Config
}

// Data is used to store an array of Datums.
// This is useful for the JSON array detection
type Data []Datum

// Datum is used to store data from all the relevant endpoints in the API
type Datum struct {
	SpaceID       string
	EnvironmentID string
	Total         string `json:"total"`
	Items         []struct {
		Name   string `json:name`
		Usage  int    `json:usage`
		Limits struct {
			Included int `json:included`
		} `json: limits`
	}
}

// Response struct is used to store http.Response and associated data
type Response struct {
	url       string
	response  *http.Response
	body      []byte
	targetURL config.TargetUrl
	err       error
}

package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"

	"os"

	cfg "github.com/infinityworks/go-common/config"
)

// TargetUrl is a structure containing a URL, space id & environment
type TargetUrl struct {
	SpaceID     string
	Environment string
	URL         string
}

// Config struct holds all of the runtime confgiguration for the application
type Config struct {
	*cfg.BaseConfig
	APIURL       string
	SpaceID      string
	APITokenEnv  string
	APITokenFile string
	APIToken     string
	Environments []string
	TargetURLs   []TargetUrl
}

// Init populates the Config struct based on environmental runtime configuration
func Init() Config {

	ac := cfg.Init()
	url := cfg.GetEnv("API_URL", "https://api.contentful.com")
	spaceID := os.Getenv("SPACE_ID")
	tokenEnv := os.Getenv("AUTH_TOKEN")
	tokenFile := os.Getenv("AUTH_TOKEN_FILE")
	token, err := getAuth(tokenEnv, tokenFile)
	environments, err := splitEnvironments(os.Getenv("ENVIRONMENTS"))
	scrapeUrls, err := getScrapeURLs(url, spaceID, environments)

	if err != nil {
		log.Errorf("Error initialising Configuration, Error: %v", err)
	}

	appConfig := Config{
		&ac,
		url,
		spaceID,
		tokenEnv,
		tokenFile,
		token,
		environments,
		scrapeUrls,
	}

	return appConfig
}

func splitEnvironments(environments string) ([]string, error) {
	var envList []string

	if environments == "" {
		return envList, fmt.Errorf("You must provide a space ID and at least one environment")
	}

	return strings.Split(environments, ","), nil
}

// Init populates the Config struct based on environmental runtime configuration
// All URL's are added to the TargetURL's string array
func getScrapeURLs(apiURL string, spaceID string, enviroments []string) ([]TargetUrl, error) {

	var urls []TargetUrl

	if spaceID == "" {
		return urls, fmt.Errorf("You must provide a space ID")
	}

	for _, e := range enviroments {
		urls = append(
			urls,
			TargetUrl{
				spaceID,
				e,
				fmt.Sprintf("/spaces/%s/environments/%s/resources", spaceID, e),
			},
		)
	}

	return urls, nil
}

// getAuth returns oauth2 token as string for usage in http.request
func getAuth(token string, tokenFile string) (string, error) {

	if token != "" {
		return token, nil
	} else if tokenFile != "" {
		b, err := ioutil.ReadFile(tokenFile)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(b)), err
	}

	return "", fmt.Errorf("No auth token provided")
}

package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/infinityworks/github-exporter/config"
	"github.com/infinityworks/github-exporter/exporter"
	web "github.com/infinityworks/github-exporter/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/steinfletcher/apitest"
)

func TestHomepage(t *testing.T) {
	test, collector := apiTest(withConfig("a/b"))
	defer prometheus.Unregister(&collector)

	test.Get("/").
		Expect(t).
		Assert(bodyContains("Contentful Prometheus Metrics Exporter")).
		Status(http.StatusOK).
		End()
}

func TestContentfulExporter(t *testing.T) {
	test, collector := apiTest(withConfig())
	defer prometheus.Unregister(&collector)

	test.Mocks(
		spacesResponse(),
	).
		Get("/metrics").
		Expect(t).
		Assert(bodyContains(`github_repo_release_downloads{created_at="2019-05-02T15:22:16Z",name="myRepo_2.0.0_windows_amd64.tar.gz",release="2.0.0",repo="myRepo",user="myOrg"} 55`)).
		Status(http.StatusOK).
		End()
}

func apiTest(conf config.Config) (*apitest.APITest, exporter.Exporter) {
	exp := exporter.Exporter{
		APIMetrics: exporter.AddMetrics(),
		Config:     conf,
	}
	server := web.NewServer(exp)

	return apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(server.Handler), exp
}

func withConfig() config.Config {
	_ = os.Setenv("SPACE_ID", "spaceid")
	_ = os.Setenv("ENVIRONMENTS", "environmentid")
	return config.Init()
}

func spacesResponse() *apitest.Mock {
	return apitest.NewMock().
		Get("https://api.contentful.com/spaces/spaceid/environments/environmentid/resources").
		Header("Authorization", "token 12345").
		RespondWith().
		Times(2).
		Body(readFile("testdata/environment_response.json")).
		Status(200).
		End()
}

func readFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func bodyContains(substr string) func(*http.Response, *http.Request) error {
	return func(res *http.Response, req *http.Request) error {
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		response := string(bytes)
		if !strings.Contains(response, substr) {
			return fmt.Errorf("response did not contain substring '%s'", substr)
		}
		return nil
	}
}

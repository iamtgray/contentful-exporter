package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

// PrometheusMetric is a structure to hold a metric specification and the contentful api name
type PrometheusMetric struct {
	UsageSpec *prometheus.Desc
	LimitSpec *prometheus.Desc
}

// AddMetrics - Add's all of the metrics to a map of strings, returns the map.
func AddMetrics() map[string]PrometheusMetric {

	APIMetrics := make(map[string]PrometheusMetric)

	// Space
	APIMetrics["Space"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "space", "usage"),
			"Total number of contentful spaces created",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "space", "limit"),
			"Limit for the number of spaces",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// Space membership
	APIMetrics["Space membership"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "spaceMembership", "usage"),
			"Total number of contentful space memberships created",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "spaceMembership", "limit"),
			"Limit for the number of space members",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// Entries
	APIMetrics["Entry"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "entry", "usage"),
			"Total number of contentful entries created",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "entry", "limit"),
			"Limit for the number of entries",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// Assets
	APIMetrics["Asset"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "asset", "usage"),
			"Total number of contentful assets created",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "asset", "limit"),
			"Limit for the number of assets",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// Records
	APIMetrics["Record"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "record", "usage"),
			"Total number of contentful records created",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "record", "limit"),
			"Limit for the number of records",
			[]string{"spaceId", "environment"}, nil,
		),
	}
	// Roles
	APIMetrics["Role"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "role", "usage"),
			"Total number of contentful roles created",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "role", "limit"),
			"Limit for the number of roles",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// Locales
	APIMetrics["Locale"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "locale", "usage"),
			"Total number of contentful locales created",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "locale", "limit"),
			"Limit for the number of locales",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// Content types
	APIMetrics["Content type"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "contentType", "usage"),
			"Total number of contentful content types created",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "contentType", "limit"),
			"Limit for the number of content types",
			[]string{"spaceId", "environment"}, nil,
		),
	}
	// API Key
	APIMetrics["Api key"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "apiKey", "usage"),
			"Total number of contentful api keys created",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "apiKey", "limit"),
			"Limit for the number of api keys",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// API Key
	APIMetrics["Environment"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "environment", "usage"),
			"Total number of contentful environments created",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "environment", "limit"),
			"Limit for the number of environments",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// Content delivery api request
	APIMetrics["Content delivery api request"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "deliveryApiRequests", "usage"),
			"Total number of contentful content delivery API requests",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "deliveryApiRequests", "limit"),
			"Limit for the number of contentful content delivery API requests",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// Asset bandwidth
	APIMetrics["Asset bandwidth"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "assetBandwidth", "usage"),
			"Total asset bandwidth usage",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "assetBandwidth", "limit"),
			"Limit for the asset bandwidth",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// Content management api request
	APIMetrics["Content management api request"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "managementApiRequests", "usage"),
			"Total management api requests",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "managementApiRequests", "limit"),
			"Limit for management api requests",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// Content preview api request
	APIMetrics["Content preview api request"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "previewApiRequests", "usage"),
			"Total preview api requests",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "previewApiRequests", "limit"),
			"Limit for preview api requests",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	// GraphQL api request
	APIMetrics["Graphql content delivery api request"] = PrometheusMetric{
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "graphQlApiRequests", "usage"),
			"Total graphql api requests",
			[]string{"spaceId", "environment"}, nil,
		),
		prometheus.NewDesc(
			prometheus.BuildFQName("contentful", "graphQlApiRequests", "limit"),
			"Limit for graphql api requests",
			[]string{"spaceId", "environment"}, nil,
		),
	}

	return APIMetrics
}

// processMetrics - processes the response data and sets the metrics using it as a source
func (e *Exporter) processMetrics(data []*Datum, ch chan<- prometheus.Metric) error {

	// APIMetrics - range through the data slice
	for _, d := range data {
		for _, i := range d.Items {
			if _, ok := e.APIMetrics[i.Name]; !ok {
				logrus.Errorf("Ignoring metric '%s' as it was not in my list of supported metrics", i.Name)
			}
			ch <- prometheus.MustNewConstMetric(e.APIMetrics[i.Name], prometheus.GaugeValue, float64(i.Usage), d.SpaceID, d.EnvironmentID)
		}
	}
	return nil
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/varshaprasad96/llamastack-go-client"
	"github.com/varshaprasad96/llamastack-go-client/internal/testutil"
	"github.com/varshaprasad96/llamastack-go-client/option"
)

func TestTelemetryLogEventWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Telemetry.LogEvent(context.TODO(), llamastackclient.TelemetryLogEventParams{
		Event: llamastackclient.TelemetryLogEventParamsEventUnion{
			OfUnstructuredLogEvent: &llamastackclient.TelemetryLogEventParamsEventUnstructuredLogEvent{
				Message:   "message",
				Severity:  "verbose",
				SpanID:    "span_id",
				Timestamp: time.Now(),
				TraceID:   "trace_id",
				Type:      llamastackclient.EventTypeUnstructuredLog,
				Attributes: map[string]llamastackclient.TelemetryLogEventParamsEventUnstructuredLogEventAttributeUnion{
					"foo": {
						OfString: llamastackclient.String("string"),
					},
				},
			},
		},
		TtlSeconds: 0,
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTelemetryQueryMetricWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Telemetry.QueryMetric(
		context.TODO(),
		"metric_name",
		llamastackclient.TelemetryQueryMetricParams{
			QueryType:   llamastackclient.TelemetryQueryMetricParamsQueryTypeRange,
			StartTime:   0,
			EndTime:     llamastackclient.Int(0),
			Granularity: llamastackclient.String("granularity"),
			LabelMatchers: []llamastackclient.TelemetryQueryMetricParamsLabelMatcher{{
				Name:     "name",
				Operator: "=",
				Value:    "value",
			}},
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

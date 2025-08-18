// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/llamastack-go-client-go"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/testutil"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Telemetry.LogEvent(context.TODO(), llamastackgoclient.TelemetryLogEventParams{
		Event: llamastackgoclient.TelemetryLogEventParamsEventUnion{
			OfUnstructuredLogEvent: &llamastackgoclient.TelemetryLogEventParamsEventUnstructuredLogEvent{
				Message:   "message",
				Severity:  "verbose",
				SpanID:    "span_id",
				Timestamp: time.Now(),
				TraceID:   "trace_id",
				Type:      llamastackgoclient.EventTypeUnstructuredLog,
				Attributes: map[string]llamastackgoclient.TelemetryLogEventParamsEventUnstructuredLogEventAttributeUnion{
					"foo": {
						OfString: llamastackgoclient.String("string"),
					},
				},
			},
		},
		TtlSeconds: 0,
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Telemetry.QueryMetric(
		context.TODO(),
		"metric_name",
		llamastackgoclient.TelemetryQueryMetricParams{
			QueryType:   llamastackgoclient.TelemetryQueryMetricParamsQueryTypeRange,
			StartTime:   0,
			EndTime:     llamastackgoclient.Int(0),
			Granularity: llamastackgoclient.String("granularity"),
			LabelMatchers: []llamastackgoclient.TelemetryQueryMetricParamsLabelMatcher{{
				Name:     "name",
				Operator: "=",
				Value:    "value",
			}},
		},
	)
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

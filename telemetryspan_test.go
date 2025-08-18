// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/llamastack-go-client-go"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/testutil"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
)

func TestTelemetrySpanExportWithOptionalParams(t *testing.T) {
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
	err := client.Telemetry.Spans.Export(context.TODO(), llamastackgoclient.TelemetrySpanExportParams{
		AttributeFilters: []llamastackgoclient.QueryConditionParam{{
			Key: "key",
			Op:  llamastackgoclient.QueryConditionOpEq,
			Value: llamastackgoclient.QueryConditionValueUnionParam{
				OfBool: llamastackgoclient.Bool(true),
			},
		}},
		AttributesToSave: []string{"string"},
		DatasetID:        "dataset_id",
		MaxDepth:         llamastackgoclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTelemetrySpanQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Telemetry.Spans.Query(context.TODO(), llamastackgoclient.TelemetrySpanQueryParams{
		AttributeFilters: []llamastackgoclient.QueryConditionParam{{
			Key: "key",
			Op:  llamastackgoclient.QueryConditionOpEq,
			Value: llamastackgoclient.QueryConditionValueUnionParam{
				OfBool: llamastackgoclient.Bool(true),
			},
		}},
		AttributesToReturn: []string{"string"},
		MaxDepth:           llamastackgoclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTelemetrySpanGetTreeWithOptionalParams(t *testing.T) {
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
	_, err := client.Telemetry.Spans.GetTree(
		context.TODO(),
		"span_id",
		llamastackgoclient.TelemetrySpanGetTreeParams{
			AttributesToReturn: []string{"string"},
			MaxDepth:           llamastackgoclient.Int(0),
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/varshaprasad96/llamastack-go-client"
	"github.com/varshaprasad96/llamastack-go-client/internal/testutil"
	"github.com/varshaprasad96/llamastack-go-client/option"
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Telemetry.Spans.Export(context.TODO(), llamastackclient.TelemetrySpanExportParams{
		AttributeFilters: []llamastackclient.QueryConditionParam{{
			Key: "key",
			Op:  llamastackclient.QueryConditionOpEq,
			Value: llamastackclient.QueryConditionValueUnionParam{
				OfBool: llamastackclient.Bool(true),
			},
		}},
		AttributesToSave: []string{"string"},
		DatasetID:        "dataset_id",
		MaxDepth:         llamastackclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Telemetry.Spans.Query(context.TODO(), llamastackclient.TelemetrySpanQueryParams{
		AttributeFilters: []llamastackclient.QueryConditionParam{{
			Key: "key",
			Op:  llamastackclient.QueryConditionOpEq,
			Value: llamastackclient.QueryConditionValueUnionParam{
				OfBool: llamastackclient.Bool(true),
			},
		}},
		AttributesToReturn: []string{"string"},
		MaxDepth:           llamastackclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Telemetry.Spans.GetTree(
		context.TODO(),
		"span_id",
		llamastackclient.TelemetrySpanGetTreeParams{
			AttributesToReturn: []string{"string"},
			MaxDepth:           llamastackclient.Int(0),
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

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

func TestOpenAIV1ResponseNewWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Responses.New(context.TODO(), llamastackclient.OpenAIV1ResponseNewParams{
		Input: llamastackclient.OpenAIV1ResponseNewParamsInputUnion{
			OfString: llamastackclient.String("string"),
		},
		Model:              "model",
		Include:            []string{"string"},
		Instructions:       llamastackclient.String("instructions"),
		MaxInferIters:      llamastackclient.Int(0),
		PreviousResponseID: llamastackclient.String("previous_response_id"),
		Store:              llamastackclient.Bool(true),
		Stream:             llamastackclient.Bool(true),
		Temperature:        llamastackclient.Float(0),
		Text: llamastackclient.ResponseTextParam{
			Format: llamastackclient.ResponseTextFormatParam{
				Type:        llamastackclient.ResponseTextFormatTypeText,
				Description: llamastackclient.String("description"),
				Name:        llamastackclient.String("name"),
				Schema: map[string]llamastackclient.ResponseTextFormatSchemaUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
				Strict: llamastackclient.Bool(true),
			},
		},
		Tools: []llamastackclient.OpenAIV1ResponseNewParamsToolUnion{{
			OfOpenAIResponseInputToolWebSearch: &llamastackclient.OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearch{
				Type:              llamastackclient.OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearch,
				SearchContextSize: llamastackclient.String("search_context_size"),
			},
		}},
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1ResponseGet(t *testing.T) {
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
	_, err := client.OpenAI.V1.Responses.Get(context.TODO(), "response_id")
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1ResponseListWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Responses.List(context.TODO(), llamastackclient.OpenAIV1ResponseListParams{
		After: llamastackclient.String("after"),
		Limit: llamastackclient.Int(0),
		Model: llamastackclient.String("model"),
		Order: llamastackclient.OrderAsc,
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1ResponseDelete(t *testing.T) {
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
	_, err := client.OpenAI.V1.Responses.Delete(context.TODO(), "response_id")
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1ResponseGetInputItemsWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Responses.GetInputItems(
		context.TODO(),
		"response_id",
		llamastackclient.OpenAIV1ResponseGetInputItemsParams{
			After:   llamastackclient.String("after"),
			Before:  llamastackclient.String("before"),
			Include: []string{"string"},
			Limit:   llamastackclient.Int(0),
			Order:   llamastackclient.OrderAsc,
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

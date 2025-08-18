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

func TestOpenAIV1ResponseNewWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Responses.New(context.TODO(), llamastackgoclient.OpenAIV1ResponseNewParams{
		Input: llamastackgoclient.OpenAIV1ResponseNewParamsInputUnion{
			OfString: llamastackgoclient.String("string"),
		},
		Model:              "model",
		Include:            []string{"string"},
		Instructions:       llamastackgoclient.String("instructions"),
		MaxInferIters:      llamastackgoclient.Int(0),
		PreviousResponseID: llamastackgoclient.String("previous_response_id"),
		Store:              llamastackgoclient.Bool(true),
		Stream:             llamastackgoclient.Bool(true),
		Temperature:        llamastackgoclient.Float(0),
		Text: llamastackgoclient.ResponseTextParam{
			Format: llamastackgoclient.ResponseTextFormatParam{
				Type:        llamastackgoclient.ResponseTextFormatTypeText,
				Description: llamastackgoclient.String("description"),
				Name:        llamastackgoclient.String("name"),
				Schema: map[string]llamastackgoclient.ResponseTextFormatSchemaUnionParam{
					"foo": {
						OfBool: llamastackgoclient.Bool(true),
					},
				},
				Strict: llamastackgoclient.Bool(true),
			},
		},
		Tools: []llamastackgoclient.OpenAIV1ResponseNewParamsToolUnion{{
			OfOpenAIResponseInputToolWebSearch: &llamastackgoclient.OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearch{
				Type:              llamastackgoclient.OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearch,
				SearchContextSize: llamastackgoclient.String("search_context_size"),
			},
		}},
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.Responses.Get(context.TODO(), "response_id")
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.Responses.List(context.TODO(), llamastackgoclient.OpenAIV1ResponseListParams{
		After: llamastackgoclient.String("after"),
		Limit: llamastackgoclient.Int(0),
		Model: llamastackgoclient.String("model"),
		Order: llamastackgoclient.OrderAsc,
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.Responses.Delete(context.TODO(), "response_id")
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.Responses.GetInputItems(
		context.TODO(),
		"response_id",
		llamastackgoclient.OpenAIV1ResponseGetInputItemsParams{
			After:   llamastackgoclient.String("after"),
			Before:  llamastackgoclient.String("before"),
			Include: []string{"string"},
			Limit:   llamastackgoclient.Int(0),
			Order:   llamastackgoclient.OrderAsc,
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

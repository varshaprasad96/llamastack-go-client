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

func TestOpenAIV1CompletionsWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Completions(context.TODO(), llamastackclient.OpenAIV1CompletionsParams{
		Model: "model",
		Prompt: llamastackclient.OpenAIV1CompletionsParamsPromptUnion{
			OfString: llamastackclient.String("string"),
		},
		BestOf:           llamastackclient.Int(0),
		Echo:             llamastackclient.Bool(true),
		FrequencyPenalty: llamastackclient.Float(0),
		GuidedChoice:     []string{"string"},
		LogitBias: map[string]float64{
			"foo": 0,
		},
		Logprobs:        llamastackclient.Bool(true),
		MaxTokens:       llamastackclient.Int(0),
		N:               llamastackclient.Int(0),
		PresencePenalty: llamastackclient.Float(0),
		PromptLogprobs:  llamastackclient.Int(0),
		Seed:            llamastackclient.Int(0),
		Stop: llamastackclient.OpenAIV1CompletionsParamsStopUnion{
			OfString: llamastackclient.String("string"),
		},
		Stream: llamastackclient.Bool(true),
		StreamOptions: map[string]llamastackclient.OpenAIV1CompletionsParamsStreamOptionUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		Suffix:      llamastackclient.String("suffix"),
		Temperature: llamastackclient.Float(0),
		TopP:        llamastackclient.Float(0),
		User:        llamastackclient.String("user"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1EmbeddingsWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Embeddings(context.TODO(), llamastackclient.OpenAIV1EmbeddingsParams{
		Input: llamastackclient.OpenAIV1EmbeddingsParamsInputUnion{
			OfString: llamastackclient.String("string"),
		},
		Model:          "model",
		Dimensions:     llamastackclient.Int(0),
		EncodingFormat: llamastackclient.String("encoding_format"),
		User:           llamastackclient.String("user"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1ModerationsWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Moderations(context.TODO(), llamastackclient.OpenAIV1ModerationsParams{
		Input: llamastackclient.OpenAIV1ModerationsParamsInputUnion{
			OfString: llamastackclient.String("string"),
		},
		Model: llamastackclient.String("model"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1GetModels(t *testing.T) {
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
	_, err := client.OpenAI.V1.GetModels(context.TODO())
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

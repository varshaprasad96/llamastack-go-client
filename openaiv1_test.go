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

func TestOpenAIV1CompletionsWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Completions(context.TODO(), llamastackgoclient.OpenAIV1CompletionsParams{
		Model: "model",
		Prompt: llamastackgoclient.OpenAIV1CompletionsParamsPromptUnion{
			OfString: llamastackgoclient.String("string"),
		},
		BestOf:           llamastackgoclient.Int(0),
		Echo:             llamastackgoclient.Bool(true),
		FrequencyPenalty: llamastackgoclient.Float(0),
		GuidedChoice:     []string{"string"},
		LogitBias: map[string]float64{
			"foo": 0,
		},
		Logprobs:        llamastackgoclient.Bool(true),
		MaxTokens:       llamastackgoclient.Int(0),
		N:               llamastackgoclient.Int(0),
		PresencePenalty: llamastackgoclient.Float(0),
		PromptLogprobs:  llamastackgoclient.Int(0),
		Seed:            llamastackgoclient.Int(0),
		Stop: llamastackgoclient.OpenAIV1CompletionsParamsStopUnion{
			OfString: llamastackgoclient.String("string"),
		},
		Stream: llamastackgoclient.Bool(true),
		StreamOptions: map[string]llamastackgoclient.OpenAIV1CompletionsParamsStreamOptionUnion{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		},
		Suffix:      llamastackgoclient.String("suffix"),
		Temperature: llamastackgoclient.Float(0),
		TopP:        llamastackgoclient.Float(0),
		User:        llamastackgoclient.String("user"),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.Embeddings(context.TODO(), llamastackgoclient.OpenAIV1EmbeddingsParams{
		Input: llamastackgoclient.OpenAIV1EmbeddingsParamsInputUnion{
			OfString: llamastackgoclient.String("string"),
		},
		Model:          "model",
		Dimensions:     llamastackgoclient.Int(0),
		EncodingFormat: llamastackgoclient.String("encoding_format"),
		User:           llamastackgoclient.String("user"),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1Moderations(t *testing.T) {
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
	_, err := client.OpenAI.V1.Moderations(context.TODO(), llamastackgoclient.OpenAIV1ModerationsParams{
		Input: llamastackgoclient.OpenAIV1ModerationsParamsInputUnion{
			OfString: llamastackgoclient.String("string"),
		},
		Model: "model",
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.GetModels(context.TODO())
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

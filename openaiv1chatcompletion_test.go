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

func TestOpenAIV1ChatCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Chat.Completions.New(context.TODO(), llamastackgoclient.OpenAIV1ChatCompletionNewParams{
		Messages: []llamastackgoclient.MessageParamUnion{{
			OfUser: &llamastackgoclient.MessageParamUser{
				Content: llamastackgoclient.MessageParamUserContentUnion{
					OfString: llamastackgoclient.String("string"),
				},
				Name: llamastackgoclient.String("name"),
			},
		}},
		Model:            "model",
		FrequencyPenalty: llamastackgoclient.Float(0),
		FunctionCall: llamastackgoclient.OpenAIV1ChatCompletionNewParamsFunctionCallUnion{
			OfString: llamastackgoclient.String("string"),
		},
		Functions: []map[string]llamastackgoclient.OpenAIV1ChatCompletionNewParamsFunctionUnion{{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		}},
		LogitBias: map[string]float64{
			"foo": 0,
		},
		Logprobs:            llamastackgoclient.Bool(true),
		MaxCompletionTokens: llamastackgoclient.Int(0),
		MaxTokens:           llamastackgoclient.Int(0),
		N:                   llamastackgoclient.Int(0),
		ParallelToolCalls:   llamastackgoclient.Bool(true),
		PresencePenalty:     llamastackgoclient.Float(0),
		ResponseFormat: llamastackgoclient.OpenAIV1ChatCompletionNewParamsResponseFormatUnion{
			OfText: &llamastackgoclient.OpenAIV1ChatCompletionNewParamsResponseFormatText{},
		},
		Seed: llamastackgoclient.Int(0),
		Stop: llamastackgoclient.OpenAIV1ChatCompletionNewParamsStopUnion{
			OfString: llamastackgoclient.String("string"),
		},
		Stream: llamastackgoclient.Bool(true),
		StreamOptions: map[string]llamastackgoclient.OpenAIV1ChatCompletionNewParamsStreamOptionUnion{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		},
		Temperature: llamastackgoclient.Float(0),
		ToolChoice: llamastackgoclient.OpenAIV1ChatCompletionNewParamsToolChoiceUnion{
			OfString: llamastackgoclient.String("string"),
		},
		Tools: []map[string]llamastackgoclient.OpenAIV1ChatCompletionNewParamsToolUnion{{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		}},
		TopLogprobs: llamastackgoclient.Int(0),
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

func TestOpenAIV1ChatCompletionGet(t *testing.T) {
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
	_, err := client.OpenAI.V1.Chat.Completions.Get(context.TODO(), "completion_id")
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1ChatCompletionListWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Chat.Completions.List(context.TODO(), llamastackgoclient.OpenAIV1ChatCompletionListParams{
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

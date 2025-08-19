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

func TestOpenAIV1ChatCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.Chat.Completions.New(context.TODO(), llamastackclient.OpenAIV1ChatCompletionNewParams{
		Messages: []llamastackclient.OpenAIV1ChatCompletionNewParamsMessageUnion{{
			OfOpenAIUserMessage: &llamastackclient.OpenAIUserMessageParam{
				Content: llamastackclient.OpenAIUserMessageParamContentUnion{
					OfString: llamastackclient.String("string"),
				},
				Role: "role",
				Name: llamastackclient.String("name"),
			},
		}},
		Model:            "model",
		FrequencyPenalty: llamastackclient.Float(0),
		FunctionCall: llamastackclient.OpenAIV1ChatCompletionNewParamsFunctionCallUnion{
			OfString: llamastackclient.String("string"),
		},
		Functions: []map[string]llamastackclient.OpenAIV1ChatCompletionNewParamsFunctionUnion{{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		}},
		LogitBias: map[string]float64{
			"foo": 0,
		},
		Logprobs:            llamastackclient.Bool(true),
		MaxCompletionTokens: llamastackclient.Int(0),
		MaxTokens:           llamastackclient.Int(0),
		N:                   llamastackclient.Int(0),
		ParallelToolCalls:   llamastackclient.Bool(true),
		PresencePenalty:     llamastackclient.Float(0),
		ResponseFormat: llamastackclient.OpenAIV1ChatCompletionNewParamsResponseFormatUnion{
			OfOpenAIResponseFormatText: &llamastackclient.OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatText{
				Type: "type",
			},
		},
		Seed: llamastackclient.Int(0),
		Stop: llamastackclient.OpenAIV1ChatCompletionNewParamsStopUnion{
			OfString: llamastackclient.String("string"),
		},
		Stream: llamastackclient.Bool(true),
		StreamOptions: map[string]llamastackclient.OpenAIV1ChatCompletionNewParamsStreamOptionUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		Temperature: llamastackclient.Float(0),
		ToolChoice: llamastackclient.OpenAIV1ChatCompletionNewParamsToolChoiceUnion{
			OfString: llamastackclient.String("string"),
		},
		Tools: []map[string]llamastackclient.OpenAIV1ChatCompletionNewParamsToolUnion{{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		}},
		TopLogprobs: llamastackclient.Int(0),
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

func TestOpenAIV1ChatCompletionGet(t *testing.T) {
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
	_, err := client.OpenAI.V1.Chat.Completions.Get(context.TODO(), "completion_id")
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.Chat.Completions.List(context.TODO(), llamastackclient.OpenAIV1ChatCompletionListParams{
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

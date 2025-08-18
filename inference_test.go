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

func TestInferenceBatchChatCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.BatchChatCompletion(context.TODO(), llamastackgoclient.InferenceBatchChatCompletionParams{
		MessagesBatch: [][]llamastackgoclient.MessageUnionParam{{{
			OfUser: &llamastackgoclient.UserMessageParam{
				Content: llamastackgoclient.InterleavedContentUnionParam{
					OfString: llamastackgoclient.String("string"),
				},
				Context: llamastackgoclient.InterleavedContentUnionParam{
					OfString: llamastackgoclient.String("string"),
				},
			},
		}}},
		ModelID: "model_id",
		Logprobs: llamastackgoclient.InferenceBatchChatCompletionParamsLogprobs{
			TopK: llamastackgoclient.Int(0),
		},
		ResponseFormat: llamastackgoclient.ResponseFormatUnionParam{
			OfJsonSchema: &llamastackgoclient.ResponseFormatJsonSchemaParam{
				JsonSchema: map[string]llamastackgoclient.ResponseFormatJsonSchemaJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackgoclient.Bool(true),
					},
				},
			},
		},
		SamplingParams: llamastackgoclient.SamplingParams{
			Strategy: llamastackgoclient.SamplingParamsStrategyUnion{
				OfGreedy: &llamastackgoclient.SamplingParamsStrategyGreedy{},
			},
			MaxTokens:         llamastackgoclient.Int(0),
			RepetitionPenalty: llamastackgoclient.Float(0),
			Stop:              []string{"string"},
		},
		ToolConfig: llamastackgoclient.ToolConfigParam{
			SystemMessageBehavior: llamastackgoclient.ToolConfigSystemMessageBehaviorAppend,
			ToolChoice:            llamastackgoclient.ToolConfigToolChoiceAuto,
			ToolPromptFormat:      llamastackgoclient.ToolConfigToolPromptFormatJson,
		},
		Tools: []llamastackgoclient.ToolDefinitionParam{{
			ToolName:    llamastackgoclient.ToolDefinitionToolNameBraveSearch,
			Description: llamastackgoclient.String("description"),
			Parameters: map[string]llamastackgoclient.ToolDefinitionParameterParam{
				"foo": {
					ParamType: "param_type",
					Default: llamastackgoclient.ToolDefinitionParameterDefaultUnionParam{
						OfBool: llamastackgoclient.Bool(true),
					},
					Description: llamastackgoclient.String("description"),
					Required:    llamastackgoclient.Bool(true),
				},
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

func TestInferenceBatchCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.BatchCompletion(context.TODO(), llamastackgoclient.InferenceBatchCompletionParams{
		ContentBatch: []llamastackgoclient.InterleavedContentUnionParam{{
			OfString: llamastackgoclient.String("string"),
		}},
		ModelID: "model_id",
		Logprobs: llamastackgoclient.InferenceBatchCompletionParamsLogprobs{
			TopK: llamastackgoclient.Int(0),
		},
		ResponseFormat: llamastackgoclient.ResponseFormatUnionParam{
			OfJsonSchema: &llamastackgoclient.ResponseFormatJsonSchemaParam{
				JsonSchema: map[string]llamastackgoclient.ResponseFormatJsonSchemaJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackgoclient.Bool(true),
					},
				},
			},
		},
		SamplingParams: llamastackgoclient.SamplingParams{
			Strategy: llamastackgoclient.SamplingParamsStrategyUnion{
				OfGreedy: &llamastackgoclient.SamplingParamsStrategyGreedy{},
			},
			MaxTokens:         llamastackgoclient.Int(0),
			RepetitionPenalty: llamastackgoclient.Float(0),
			Stop:              []string{"string"},
		},
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferenceChatCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.ChatCompletion(context.TODO(), llamastackgoclient.InferenceChatCompletionParams{
		Messages: []llamastackgoclient.MessageUnionParam{{
			OfUser: &llamastackgoclient.UserMessageParam{
				Content: llamastackgoclient.InterleavedContentUnionParam{
					OfString: llamastackgoclient.String("string"),
				},
				Context: llamastackgoclient.InterleavedContentUnionParam{
					OfString: llamastackgoclient.String("string"),
				},
			},
		}},
		ModelID: "model_id",
		Logprobs: llamastackgoclient.InferenceChatCompletionParamsLogprobs{
			TopK: llamastackgoclient.Int(0),
		},
		ResponseFormat: llamastackgoclient.ResponseFormatUnionParam{
			OfJsonSchema: &llamastackgoclient.ResponseFormatJsonSchemaParam{
				JsonSchema: map[string]llamastackgoclient.ResponseFormatJsonSchemaJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackgoclient.Bool(true),
					},
				},
			},
		},
		SamplingParams: llamastackgoclient.SamplingParams{
			Strategy: llamastackgoclient.SamplingParamsStrategyUnion{
				OfGreedy: &llamastackgoclient.SamplingParamsStrategyGreedy{},
			},
			MaxTokens:         llamastackgoclient.Int(0),
			RepetitionPenalty: llamastackgoclient.Float(0),
			Stop:              []string{"string"},
		},
		Stream:     llamastackgoclient.Bool(true),
		ToolChoice: llamastackgoclient.InferenceChatCompletionParamsToolChoiceAuto,
		ToolConfig: llamastackgoclient.ToolConfigParam{
			SystemMessageBehavior: llamastackgoclient.ToolConfigSystemMessageBehaviorAppend,
			ToolChoice:            llamastackgoclient.ToolConfigToolChoiceAuto,
			ToolPromptFormat:      llamastackgoclient.ToolConfigToolPromptFormatJson,
		},
		ToolPromptFormat: llamastackgoclient.InferenceChatCompletionParamsToolPromptFormatJson,
		Tools: []llamastackgoclient.ToolDefinitionParam{{
			ToolName:    llamastackgoclient.ToolDefinitionToolNameBraveSearch,
			Description: llamastackgoclient.String("description"),
			Parameters: map[string]llamastackgoclient.ToolDefinitionParameterParam{
				"foo": {
					ParamType: "param_type",
					Default: llamastackgoclient.ToolDefinitionParameterDefaultUnionParam{
						OfBool: llamastackgoclient.Bool(true),
					},
					Description: llamastackgoclient.String("description"),
					Required:    llamastackgoclient.Bool(true),
				},
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

func TestInferenceCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.Completion(context.TODO(), llamastackgoclient.InferenceCompletionParams{
		Content: llamastackgoclient.InterleavedContentUnionParam{
			OfString: llamastackgoclient.String("string"),
		},
		ModelID: "model_id",
		Logprobs: llamastackgoclient.InferenceCompletionParamsLogprobs{
			TopK: llamastackgoclient.Int(0),
		},
		ResponseFormat: llamastackgoclient.ResponseFormatUnionParam{
			OfJsonSchema: &llamastackgoclient.ResponseFormatJsonSchemaParam{
				JsonSchema: map[string]llamastackgoclient.ResponseFormatJsonSchemaJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackgoclient.Bool(true),
					},
				},
			},
		},
		SamplingParams: llamastackgoclient.SamplingParams{
			Strategy: llamastackgoclient.SamplingParamsStrategyUnion{
				OfGreedy: &llamastackgoclient.SamplingParamsStrategyGreedy{},
			},
			MaxTokens:         llamastackgoclient.Int(0),
			RepetitionPenalty: llamastackgoclient.Float(0),
			Stop:              []string{"string"},
		},
		Stream: llamastackgoclient.Bool(true),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferenceEmbeddingsWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.Embeddings(context.TODO(), llamastackgoclient.InferenceEmbeddingsParams{
		Contents: llamastackgoclient.InferenceEmbeddingsParamsContentsUnion{
			OfStringArray: []string{"string"},
		},
		ModelID:         "model_id",
		OutputDimension: llamastackgoclient.Int(0),
		TaskType:        llamastackgoclient.InferenceEmbeddingsParamsTaskTypeQuery,
		TextTruncation:  llamastackgoclient.InferenceEmbeddingsParamsTextTruncationNone,
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

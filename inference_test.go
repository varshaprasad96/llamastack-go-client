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

func TestInferenceBatchChatCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.BatchChatCompletion(context.TODO(), llamastackclient.InferenceBatchChatCompletionParams{
		MessagesBatch: [][]llamastackclient.MessageUnionParam{{{
			OfUserMessage: &llamastackclient.UserMessageParam{
				Content: llamastackclient.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
				Role: llamastackclient.UserMessageRoleSystem,
				Context: llamastackclient.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
			},
		}}},
		ModelID: "model_id",
		Logprobs: llamastackclient.InferenceBatchChatCompletionParamsLogprobs{
			TopK: llamastackclient.Int(0),
		},
		ResponseFormat: llamastackclient.ResponseFormatUnionParam{
			OfJsonSchemaResponseFormat: &llamastackclient.JsonSchemaResponseFormatParam{
				JsonSchema: map[string]llamastackclient.JsonSchemaResponseFormatJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
				Type: llamastackclient.JsonSchemaResponseFormatTypeJsonSchema,
			},
		},
		SamplingParams: llamastackclient.SamplingParams{
			Strategy: llamastackclient.SamplingParamsStrategyUnion{
				OfGreedySamplingStrategy: &llamastackclient.GreedySamplingStrategyParam{
					Type: llamastackclient.GreedySamplingStrategyTypeGreedy,
				},
			},
			MaxTokens:         llamastackclient.Int(0),
			RepetitionPenalty: llamastackclient.Float(0),
			Stop:              []string{"string"},
		},
		ToolConfig: llamastackclient.ToolConfigParam{
			SystemMessageBehavior: llamastackclient.ToolConfigSystemMessageBehaviorAppend,
			ToolChoice:            llamastackclient.ToolConfigToolChoiceAuto,
			ToolPromptFormat:      llamastackclient.ToolConfigToolPromptFormatJson,
		},
		Tools: []llamastackclient.ToolDefinitionParam{{
			ToolName:    llamastackclient.ToolDefinitionToolNameBraveSearch,
			Description: llamastackclient.String("description"),
			Parameters: map[string]llamastackclient.ToolDefinitionParameterParam{
				"foo": {
					ParamType: "param_type",
					Default: llamastackclient.ToolDefinitionParameterDefaultUnionParam{
						OfBool: llamastackclient.Bool(true),
					},
					Description: llamastackclient.String("description"),
					Required:    llamastackclient.Bool(true),
				},
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

func TestInferenceBatchCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.BatchCompletion(context.TODO(), llamastackclient.InferenceBatchCompletionParams{
		ContentBatch: []llamastackclient.InterleavedContentUnionParam{{
			OfString: llamastackclient.String("string"),
		}},
		ModelID: "model_id",
		Logprobs: llamastackclient.InferenceBatchCompletionParamsLogprobs{
			TopK: llamastackclient.Int(0),
		},
		ResponseFormat: llamastackclient.ResponseFormatUnionParam{
			OfJsonSchemaResponseFormat: &llamastackclient.JsonSchemaResponseFormatParam{
				JsonSchema: map[string]llamastackclient.JsonSchemaResponseFormatJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
				Type: llamastackclient.JsonSchemaResponseFormatTypeJsonSchema,
			},
		},
		SamplingParams: llamastackclient.SamplingParams{
			Strategy: llamastackclient.SamplingParamsStrategyUnion{
				OfGreedySamplingStrategy: &llamastackclient.GreedySamplingStrategyParam{
					Type: llamastackclient.GreedySamplingStrategyTypeGreedy,
				},
			},
			MaxTokens:         llamastackclient.Int(0),
			RepetitionPenalty: llamastackclient.Float(0),
			Stop:              []string{"string"},
		},
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Inference.ChatCompletion(context.TODO(), llamastackclient.InferenceChatCompletionParams{
		Messages: []llamastackclient.MessageUnionParam{{
			OfUserMessage: &llamastackclient.UserMessageParam{
				Content: llamastackclient.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
				Role: llamastackclient.UserMessageRoleSystem,
				Context: llamastackclient.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
			},
		}},
		ModelID: "model_id",
		Logprobs: llamastackclient.InferenceChatCompletionParamsLogprobs{
			TopK: llamastackclient.Int(0),
		},
		ResponseFormat: llamastackclient.ResponseFormatUnionParam{
			OfJsonSchemaResponseFormat: &llamastackclient.JsonSchemaResponseFormatParam{
				JsonSchema: map[string]llamastackclient.JsonSchemaResponseFormatJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
				Type: llamastackclient.JsonSchemaResponseFormatTypeJsonSchema,
			},
		},
		SamplingParams: llamastackclient.SamplingParams{
			Strategy: llamastackclient.SamplingParamsStrategyUnion{
				OfGreedySamplingStrategy: &llamastackclient.GreedySamplingStrategyParam{
					Type: llamastackclient.GreedySamplingStrategyTypeGreedy,
				},
			},
			MaxTokens:         llamastackclient.Int(0),
			RepetitionPenalty: llamastackclient.Float(0),
			Stop:              []string{"string"},
		},
		Stream:     llamastackclient.Bool(true),
		ToolChoice: llamastackclient.InferenceChatCompletionParamsToolChoiceAuto,
		ToolConfig: llamastackclient.ToolConfigParam{
			SystemMessageBehavior: llamastackclient.ToolConfigSystemMessageBehaviorAppend,
			ToolChoice:            llamastackclient.ToolConfigToolChoiceAuto,
			ToolPromptFormat:      llamastackclient.ToolConfigToolPromptFormatJson,
		},
		ToolPromptFormat: llamastackclient.InferenceChatCompletionParamsToolPromptFormatJson,
		Tools: []llamastackclient.ToolDefinitionParam{{
			ToolName:    llamastackclient.ToolDefinitionToolNameBraveSearch,
			Description: llamastackclient.String("description"),
			Parameters: map[string]llamastackclient.ToolDefinitionParameterParam{
				"foo": {
					ParamType: "param_type",
					Default: llamastackclient.ToolDefinitionParameterDefaultUnionParam{
						OfBool: llamastackclient.Bool(true),
					},
					Description: llamastackclient.String("description"),
					Required:    llamastackclient.Bool(true),
				},
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

func TestInferenceCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.Completion(context.TODO(), llamastackclient.InferenceCompletionParams{
		Content: llamastackclient.InterleavedContentUnionParam{
			OfString: llamastackclient.String("string"),
		},
		ModelID: "model_id",
		Logprobs: llamastackclient.InferenceCompletionParamsLogprobs{
			TopK: llamastackclient.Int(0),
		},
		ResponseFormat: llamastackclient.ResponseFormatUnionParam{
			OfJsonSchemaResponseFormat: &llamastackclient.JsonSchemaResponseFormatParam{
				JsonSchema: map[string]llamastackclient.JsonSchemaResponseFormatJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
				Type: llamastackclient.JsonSchemaResponseFormatTypeJsonSchema,
			},
		},
		SamplingParams: llamastackclient.SamplingParams{
			Strategy: llamastackclient.SamplingParamsStrategyUnion{
				OfGreedySamplingStrategy: &llamastackclient.GreedySamplingStrategyParam{
					Type: llamastackclient.GreedySamplingStrategyTypeGreedy,
				},
			},
			MaxTokens:         llamastackclient.Int(0),
			RepetitionPenalty: llamastackclient.Float(0),
			Stop:              []string{"string"},
		},
		Stream: llamastackclient.Bool(true),
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Inference.Embeddings(context.TODO(), llamastackclient.InferenceEmbeddingsParams{
		Contents: llamastackclient.InferenceEmbeddingsParamsContentsUnion{
			OfStringArray: []string{"string"},
		},
		ModelID:         "model_id",
		OutputDimension: llamastackclient.Int(0),
		TaskType:        llamastackclient.InferenceEmbeddingsParamsTaskTypeQuery,
		TextTruncation:  llamastackclient.InferenceEmbeddingsParamsTextTruncationNone,
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

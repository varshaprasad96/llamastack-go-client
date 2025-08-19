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
		Messages: []llamastackclient.MessageParam{{
			Content: llamastackclient.MessageContentUnionParam{
				OfString: llamastackclient.String("string"),
			},
			Role: llamastackclient.MessageRoleUser,
			Name: llamastackclient.String("name"),
			ToolCalls: []llamastackclient.ToolCallParam{{
				ID: "id",
				Function: llamastackclient.ToolCallFunctionParam{
					Arguments: "arguments",
					Name:      "name",
				},
			}},
		}},
		Model:  "model",
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
		Contents: []string{"string"},
		Model:    "model",
		Stream:   llamastackclient.Bool(true),
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
		Messages: []llamastackclient.MessageParam{{
			Content: llamastackclient.MessageContentUnionParam{
				OfString: llamastackclient.String("string"),
			},
			Role: llamastackclient.MessageRoleUser,
			Name: llamastackclient.String("name"),
			ToolCalls: []llamastackclient.ToolCallParam{{
				ID: "id",
				Function: llamastackclient.ToolCallFunctionParam{
					Arguments: "arguments",
					Name:      "name",
				},
			}},
		}},
		Model:  "model",
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
		Content: "content",
		Model:   "model",
		Stream:  llamastackclient.Bool(true),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferenceEmbeddings(t *testing.T) {
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
		Input: llamastackclient.InferenceEmbeddingsParamsInputUnion{
			OfString: llamastackclient.String("string"),
		},
		Model: "model",
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

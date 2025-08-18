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

func TestOpenAIV1VectorStoreFileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.Files.New(
		context.TODO(),
		"vector_store_id",
		llamastackgoclient.OpenAIV1VectorStoreFileNewParams{
			FileID: "file_id",
			Attributes: map[string]llamastackgoclient.OpenAIV1VectorStoreFileNewParamsAttributeUnion{
				"foo": {
					OfBool: llamastackgoclient.Bool(true),
				},
			},
			ChunkingStrategy: llamastackgoclient.ChunkingStrategyUnionParam{
				OfAuto: &llamastackgoclient.ChunkingStrategyAutoParam{},
			},
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

func TestOpenAIV1VectorStoreFileGet(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.Files.Get(
		context.TODO(),
		"file_id",
		llamastackgoclient.OpenAIV1VectorStoreFileGetParams{
			VectorStoreID: "vector_store_id",
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

func TestOpenAIV1VectorStoreFileUpdate(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.Files.Update(
		context.TODO(),
		"file_id",
		llamastackgoclient.OpenAIV1VectorStoreFileUpdateParams{
			VectorStoreID: "vector_store_id",
			Attributes: map[string]llamastackgoclient.OpenAIV1VectorStoreFileUpdateParamsAttributeUnion{
				"foo": {
					OfBool: llamastackgoclient.Bool(true),
				},
			},
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

func TestOpenAIV1VectorStoreFileListWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.Files.List(
		context.TODO(),
		"vector_store_id",
		llamastackgoclient.OpenAIV1VectorStoreFileListParams{
			After:  llamastackgoclient.String("after"),
			Before: llamastackgoclient.String("before"),
			Filter: llamastackgoclient.FileStatusCompleted,
			Limit:  llamastackgoclient.Int(0),
			Order:  llamastackgoclient.String("order"),
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

func TestOpenAIV1VectorStoreFileDelete(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.Files.Delete(
		context.TODO(),
		"file_id",
		llamastackgoclient.OpenAIV1VectorStoreFileDeleteParams{
			VectorStoreID: "vector_store_id",
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

func TestOpenAIV1VectorStoreFileGetContent(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.Files.GetContent(
		context.TODO(),
		"file_id",
		llamastackgoclient.OpenAIV1VectorStoreFileGetContentParams{
			VectorStoreID: "vector_store_id",
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

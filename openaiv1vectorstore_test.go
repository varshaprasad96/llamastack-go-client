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

func TestOpenAIV1VectorStoreNewWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.New(context.TODO(), llamastackgoclient.OpenAIV1VectorStoreNewParams{
		ChunkingStrategy: map[string]llamastackgoclient.OpenAIV1VectorStoreNewParamsChunkingStrategyUnion{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		},
		EmbeddingDimension: llamastackgoclient.Int(0),
		EmbeddingModel:     llamastackgoclient.String("embedding_model"),
		ExpiresAfter: map[string]llamastackgoclient.OpenAIV1VectorStoreNewParamsExpiresAfterUnion{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		},
		FileIDs: []string{"string"},
		Metadata: map[string]llamastackgoclient.OpenAIV1VectorStoreNewParamsMetadataUnion{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		},
		Name:       llamastackgoclient.String("name"),
		ProviderID: llamastackgoclient.String("provider_id"),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1VectorStoreGet(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.Get(context.TODO(), "vector_store_id")
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1VectorStoreUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.Update(
		context.TODO(),
		"vector_store_id",
		llamastackgoclient.OpenAIV1VectorStoreUpdateParams{
			ExpiresAfter: map[string]llamastackgoclient.OpenAIV1VectorStoreUpdateParamsExpiresAfterUnion{
				"foo": {
					OfBool: llamastackgoclient.Bool(true),
				},
			},
			Metadata: map[string]llamastackgoclient.OpenAIV1VectorStoreUpdateParamsMetadataUnion{
				"foo": {
					OfBool: llamastackgoclient.Bool(true),
				},
			},
			Name: llamastackgoclient.String("name"),
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

func TestOpenAIV1VectorStoreListWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.List(context.TODO(), llamastackgoclient.OpenAIV1VectorStoreListParams{
		After:  llamastackgoclient.String("after"),
		Before: llamastackgoclient.String("before"),
		Limit:  llamastackgoclient.Int(0),
		Order:  llamastackgoclient.String("order"),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1VectorStoreDelete(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.Delete(context.TODO(), "vector_store_id")
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOpenAIV1VectorStoreSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.Search(
		context.TODO(),
		"vector_store_id",
		llamastackgoclient.OpenAIV1VectorStoreSearchParams{
			Query: llamastackgoclient.OpenAIV1VectorStoreSearchParamsQueryUnion{
				OfString: llamastackgoclient.String("string"),
			},
			Filters: map[string]llamastackgoclient.OpenAIV1VectorStoreSearchParamsFilterUnion{
				"foo": {
					OfBool: llamastackgoclient.Bool(true),
				},
			},
			MaxNumResults: llamastackgoclient.Int(0),
			RankingOptions: llamastackgoclient.OpenAIV1VectorStoreSearchParamsRankingOptions{
				Ranker:         llamastackgoclient.String("ranker"),
				ScoreThreshold: llamastackgoclient.Float(0),
			},
			RewriteQuery: llamastackgoclient.Bool(true),
			SearchMode:   llamastackgoclient.String("search_mode"),
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

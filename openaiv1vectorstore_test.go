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

func TestOpenAIV1VectorStoreNewWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.V1.VectorStores.New(context.TODO(), llamastackclient.OpenAIV1VectorStoreNewParams{
		Name: "name",
		ChunkingStrategy: map[string]llamastackclient.OpenAIV1VectorStoreNewParamsChunkingStrategyUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		EmbeddingDimension: llamastackclient.Int(0),
		EmbeddingModel:     llamastackclient.String("embedding_model"),
		ExpiresAfter: map[string]llamastackclient.OpenAIV1VectorStoreNewParamsExpiresAfterUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		FileIDs: []string{"string"},
		Metadata: map[string]llamastackclient.OpenAIV1VectorStoreNewParamsMetadataUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		ProviderID:         llamastackclient.String("provider_id"),
		ProviderVectorDBID: llamastackclient.String("provider_vector_db_id"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.VectorStores.Get(context.TODO(), "vector_store_id")
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.VectorStores.Update(
		context.TODO(),
		"vector_store_id",
		llamastackclient.OpenAIV1VectorStoreUpdateParams{
			ExpiresAfter: map[string]llamastackclient.OpenAIV1VectorStoreUpdateParamsExpiresAfterUnion{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			},
			Metadata: map[string]llamastackclient.OpenAIV1VectorStoreUpdateParamsMetadataUnion{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			},
			Name: llamastackclient.String("name"),
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.VectorStores.List(context.TODO(), llamastackclient.OpenAIV1VectorStoreListParams{
		After:  llamastackclient.String("after"),
		Before: llamastackclient.String("before"),
		Limit:  llamastackclient.Int(0),
		Order:  llamastackclient.String("order"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.VectorStores.Delete(context.TODO(), "vector_store_id")
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.OpenAI.V1.VectorStores.Search(
		context.TODO(),
		"vector_store_id",
		llamastackclient.OpenAIV1VectorStoreSearchParams{
			Query: llamastackclient.OpenAIV1VectorStoreSearchParamsQueryUnion{
				OfString: llamastackclient.String("string"),
			},
			Filters: map[string]llamastackclient.OpenAIV1VectorStoreSearchParamsFilterUnion{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			},
			MaxNumResults: llamastackclient.Int(0),
			RankingOptions: llamastackclient.OpenAIV1VectorStoreSearchParamsRankingOptions{
				Ranker:         llamastackclient.String("ranker"),
				ScoreThreshold: llamastackclient.Float(0),
			},
			RewriteQuery: llamastackclient.Bool(true),
			SearchMode:   llamastackclient.String("search_mode"),
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

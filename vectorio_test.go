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

func TestVectorIoInsertWithOptionalParams(t *testing.T) {
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
	err := client.VectorIo.Insert(context.TODO(), llamastackgoclient.VectorIoInsertParams{
		Chunks: []llamastackgoclient.ChunkParam{{
			Content: llamastackgoclient.InterleavedContentUnionParam{
				OfString: llamastackgoclient.String("string"),
			},
			Metadata: map[string]llamastackgoclient.ChunkMetadataUnionParam{
				"foo": {
					OfBool: llamastackgoclient.Bool(true),
				},
			},
			ChunkMetadata: llamastackgoclient.ChunkChunkMetadataParam{
				ChunkEmbeddingDimension: llamastackgoclient.Int(0),
				ChunkEmbeddingModel:     llamastackgoclient.String("chunk_embedding_model"),
				ChunkID:                 llamastackgoclient.String("chunk_id"),
				ChunkTokenizer:          llamastackgoclient.String("chunk_tokenizer"),
				ChunkWindow:             llamastackgoclient.String("chunk_window"),
				ContentTokenCount:       llamastackgoclient.Int(0),
				CreatedTimestamp:        llamastackgoclient.Int(0),
				DocumentID:              llamastackgoclient.String("document_id"),
				MetadataTokenCount:      llamastackgoclient.Int(0),
				Source:                  llamastackgoclient.String("source"),
				UpdatedTimestamp:        llamastackgoclient.Int(0),
			},
			Embedding:     []float64{0},
			StoredChunkID: llamastackgoclient.String("stored_chunk_id"),
		}},
		VectorDBID: "vector_db_id",
		TtlSeconds: llamastackgoclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVectorIoQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorIo.Query(context.TODO(), llamastackgoclient.VectorIoQueryParams{
		Query: llamastackgoclient.InterleavedContentUnionParam{
			OfString: llamastackgoclient.String("string"),
		},
		VectorDBID: "vector_db_id",
		Params: map[string]llamastackgoclient.VectorIoQueryParamsParamUnion{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
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

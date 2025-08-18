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

func TestToolRuntimeRagToolInsert(t *testing.T) {
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
	err := client.ToolRuntime.RagTool.Insert(context.TODO(), llamastackgoclient.ToolRuntimeRagToolInsertParams{
		ChunkSizeInTokens: 0,
		Documents: []llamastackgoclient.ToolRuntimeRagToolInsertParamsDocument{{
			Content: llamastackgoclient.ToolRuntimeRagToolInsertParamsDocumentContentUnion{
				OfString: llamastackgoclient.String("string"),
			},
			DocumentID: "document_id",
			Metadata: map[string]llamastackgoclient.ToolRuntimeRagToolInsertParamsDocumentMetadataUnion{
				"foo": {
					OfBool: llamastackgoclient.Bool(true),
				},
			},
			MimeType: llamastackgoclient.String("mime_type"),
		}},
		VectorDBID: "vector_db_id",
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolRuntimeRagToolQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolRuntime.RagTool.Query(context.TODO(), llamastackgoclient.ToolRuntimeRagToolQueryParams{
		Content: llamastackgoclient.InterleavedContentUnionParam{
			OfString: llamastackgoclient.String("string"),
		},
		VectorDBIDs: []string{"string"},
		QueryConfig: llamastackgoclient.ToolRuntimeRagToolQueryParamsQueryConfig{
			ChunkTemplate:      "chunk_template",
			MaxChunks:          0,
			MaxTokensInContext: 0,
			QueryGeneratorConfig: llamastackgoclient.ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion{
				OfDefault: &llamastackgoclient.ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigDefault{
					Separator: "separator",
				},
			},
			Mode: "vector",
			Ranker: llamastackgoclient.ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion{
				OfRrf: &llamastackgoclient.ToolRuntimeRagToolQueryParamsQueryConfigRankerRrf{
					ImpactFactor: 0,
				},
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

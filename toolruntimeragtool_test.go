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

func TestToolRuntimeRagToolInsert(t *testing.T) {
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
	err := client.ToolRuntime.RagTool.Insert(context.TODO(), llamastackclient.ToolRuntimeRagToolInsertParams{
		ChunkSizeInTokens: 0,
		Documents: []llamastackclient.ToolRuntimeRagToolInsertParamsDocument{{
			Content: llamastackclient.ToolRuntimeRagToolInsertParamsDocumentContentUnion{
				OfString: llamastackclient.String("string"),
			},
			DocumentID: "document_id",
			Metadata: map[string]llamastackclient.ToolRuntimeRagToolInsertParamsDocumentMetadataUnion{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			},
			MimeType: llamastackclient.String("mime_type"),
		}},
		VectorDBID: "vector_db_id",
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolRuntime.RagTool.Query(context.TODO(), llamastackclient.ToolRuntimeRagToolQueryParams{
		Content: llamastackclient.InterleavedContentUnionParam{
			OfString: llamastackclient.String("string"),
		},
		VectorDBIDs: []string{"string"},
		QueryConfig: llamastackclient.ToolRuntimeRagToolQueryParamsQueryConfig{
			ChunkTemplate:      "chunk_template",
			MaxChunks:          0,
			MaxTokensInContext: 0,
			QueryGeneratorConfig: llamastackclient.ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion{
				OfDefault: &llamastackclient.ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigDefault{
					Separator: "separator",
				},
			},
			Mode: "vector",
			Ranker: llamastackclient.ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion{
				OfRrf: &llamastackclient.ToolRuntimeRagToolQueryParamsQueryConfigRankerRrf{
					ImpactFactor: 0,
				},
			},
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

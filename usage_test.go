// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/llamastack-go-client-go"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/testutil"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
)

func TestUsage(t *testing.T) {
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
	err := client.Datasetio.AppendRows(
		context.TODO(),
		"REPLACE_ME",
		llamastackgoclient.DatasetioAppendRowsParams{
			Rows: []map[string]llamastackgoclient.DatasetioAppendRowsParamsRowUnion{{
				"foo": {
					OfBool: llamastackgoclient.Bool(true),
				},
			}},
		},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

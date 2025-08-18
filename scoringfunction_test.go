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

func TestScoringFunctionNewWithOptionalParams(t *testing.T) {
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
	err := client.ScoringFunctions.New(context.TODO(), llamastackgoclient.ScoringFunctionNewParams{
		Description: "description",
		ReturnType: llamastackgoclient.ParamTypeUnion{
			OfString: &llamastackgoclient.ParamTypeString{},
		},
		ScoringFnID: "scoring_fn_id",
		Params: llamastackgoclient.ScoringFnParamsUnion{
			OfLlmAsJudgeScoringFns: &llamastackgoclient.ScoringFnParamsLlmAsJudgeScoringFnParams{
				AggregationFunctions: []llamastackgoclient.AggregationFunctionType{llamastackgoclient.AggregationFunctionTypeAverage},
				JudgeModel:           "judge_model",
				JudgeScoreRegexes:    []string{"string"},
				Type:                 llamastackgoclient.ScoringFnParamsTypeLlmAsJudge,
				PromptTemplate:       llamastackgoclient.String("prompt_template"),
			},
		},
		ProviderID:          llamastackgoclient.String("provider_id"),
		ProviderScoringFnID: llamastackgoclient.String("provider_scoring_fn_id"),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScoringFunctionGet(t *testing.T) {
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
	_, err := client.ScoringFunctions.Get(context.TODO(), "scoring_fn_id")
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScoringFunctionList(t *testing.T) {
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
	_, err := client.ScoringFunctions.List(context.TODO())
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

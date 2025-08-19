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

func TestScoringFunctionNewWithOptionalParams(t *testing.T) {
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
	err := client.ScoringFunctions.New(context.TODO(), llamastackclient.ScoringFunctionNewParams{
		Description: "description",
		ReturnType: llamastackclient.ParamTypeUnion{
			OfStringType: &llamastackclient.ParamTypeStringType{
				ParamTypeBase: llamastackclient.ParamTypeBase{
					Type: llamastackclient.ParamTypeBaseTypeString,
				},
				StringTypeParam: llamastackclient.NewStringTypeParam(),
			},
		},
		ScoringFnID: "scoring_fn_id",
		Params: llamastackclient.ScoringFnParamsUnion{
			OfLlmAsJudgeScoringFns: &llamastackclient.ScoringFnParamsLlmAsJudgeScoringFnParams{
				AggregationFunctions: []llamastackclient.AggregationFunctionType{llamastackclient.AggregationFunctionTypeAverage},
				JudgeModel:           "judge_model",
				JudgeScoreRegexes:    []string{"string"},
				Type:                 llamastackclient.ScoringFnParamsTypeLlmAsJudge,
				PromptTemplate:       llamastackclient.String("prompt_template"),
			},
		},
		ProviderID:          llamastackclient.String("provider_id"),
		ProviderScoringFnID: llamastackclient.String("provider_scoring_fn_id"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ScoringFunctions.Get(context.TODO(), "scoring_fn_id")
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ScoringFunctions.List(context.TODO())
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

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

func TestScoringScore(t *testing.T) {
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
	_, err := client.Scoring.Score(context.TODO(), llamastackclient.ScoringScoreParams{
		InputRows: []map[string]llamastackclient.ScoringScoreParamsInputRowUnion{{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		}},
		ScoringFunctions: map[string]llamastackclient.ScoringFnParamsUnion{
			"foo": {
				OfLlmAsJudgeScoringFns: &llamastackclient.ScoringFnParamsLlmAsJudgeScoringFnParams{
					AggregationFunctions: []llamastackclient.AggregationFunctionType{llamastackclient.AggregationFunctionTypeAverage},
					JudgeModel:           "judge_model",
					JudgeScoreRegexes:    []string{"string"},
					Type:                 llamastackclient.ScoringFnParamsTypeLlmAsJudge,
					PromptTemplate:       llamastackclient.String("prompt_template"),
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

func TestScoringScoreBatch(t *testing.T) {
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
	_, err := client.Scoring.ScoreBatch(context.TODO(), llamastackclient.ScoringScoreBatchParams{
		DatasetID:          "dataset_id",
		SaveResultsDataset: true,
		ScoringFunctions: map[string]llamastackclient.ScoringFnParamsUnion{
			"foo": {
				OfLlmAsJudgeScoringFns: &llamastackclient.ScoringFnParamsLlmAsJudgeScoringFnParams{
					AggregationFunctions: []llamastackclient.AggregationFunctionType{llamastackclient.AggregationFunctionTypeAverage},
					JudgeModel:           "judge_model",
					JudgeScoreRegexes:    []string{"string"},
					Type:                 llamastackclient.ScoringFnParamsTypeLlmAsJudge,
					PromptTemplate:       llamastackclient.String("prompt_template"),
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

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

func TestScoringScore(t *testing.T) {
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
	_, err := client.Scoring.Score(context.TODO(), llamastackgoclient.ScoringScoreParams{
		InputRows: []map[string]llamastackgoclient.ScoringScoreParamsInputRowUnion{{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		}},
		ScoringFunctions: map[string]llamastackgoclient.ScoringFnParamsUnion{
			"foo": {
				OfLlmAsJudgeScoringFns: &llamastackgoclient.ScoringFnParamsLlmAsJudgeScoringFnParams{
					AggregationFunctions: []llamastackgoclient.AggregationFunctionType{llamastackgoclient.AggregationFunctionTypeAverage},
					JudgeModel:           "judge_model",
					JudgeScoreRegexes:    []string{"string"},
					Type:                 llamastackgoclient.ScoringFnParamsTypeLlmAsJudge,
					PromptTemplate:       llamastackgoclient.String("prompt_template"),
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

func TestScoringScoreBatch(t *testing.T) {
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
	_, err := client.Scoring.ScoreBatch(context.TODO(), llamastackgoclient.ScoringScoreBatchParams{
		DatasetID:          "dataset_id",
		SaveResultsDataset: true,
		ScoringFunctions: map[string]llamastackgoclient.ScoringFnParamsUnion{
			"foo": {
				OfLlmAsJudgeScoringFns: &llamastackgoclient.ScoringFnParamsLlmAsJudgeScoringFnParams{
					AggregationFunctions: []llamastackgoclient.AggregationFunctionType{llamastackgoclient.AggregationFunctionTypeAverage},
					JudgeModel:           "judge_model",
					JudgeScoreRegexes:    []string{"string"},
					Type:                 llamastackgoclient.ScoringFnParamsTypeLlmAsJudge,
					PromptTemplate:       llamastackgoclient.String("prompt_template"),
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

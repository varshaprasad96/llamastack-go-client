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

func TestEvalBenchmarkNewWithOptionalParams(t *testing.T) {
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
	err := client.Eval.Benchmarks.New(context.TODO(), llamastackclient.EvalBenchmarkNewParams{
		BenchmarkID:      "benchmark_id",
		DatasetID:        "dataset_id",
		ScoringFunctions: []string{"string"},
		Metadata: map[string]llamastackclient.EvalBenchmarkNewParamsMetadataUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		ProviderBenchmarkID: llamastackclient.String("provider_benchmark_id"),
		ProviderID:          llamastackclient.String("provider_id"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvalBenchmarkGet(t *testing.T) {
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
	_, err := client.Eval.Benchmarks.Get(context.TODO(), "benchmark_id")
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvalBenchmarkList(t *testing.T) {
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
	_, err := client.Eval.Benchmarks.List(context.TODO())
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvalBenchmarkEvaluateWithOptionalParams(t *testing.T) {
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
	_, err := client.Eval.Benchmarks.Evaluate(
		context.TODO(),
		"benchmark_id",
		llamastackclient.EvalBenchmarkEvaluateParams{
			BenchmarkConfig: llamastackclient.BenchmarkConfigParam{
				EvalCandidate: llamastackclient.BenchmarkConfigEvalCandidateUnionParam{
					OfModel: &llamastackclient.ModelCandidateParam{
						Model: "model",
						SamplingParams: llamastackclient.SamplingParams{
							Strategy: llamastackclient.SamplingParamsStrategyUnion{
								OfGreedy: &llamastackclient.GreedySamplingStrategyParam{},
							},
							MaxTokens:         llamastackclient.Int(0),
							RepetitionPenalty: llamastackclient.Float(0),
							Stop:              []string{"string"},
						},
						SystemMessage: llamastackclient.SystemMessageParam{
							Content: llamastackclient.InterleavedContentUnionParam{
								OfString: llamastackclient.String("string"),
							},
						},
					},
				},
				ScoringParams: map[string]llamastackclient.ScoringFnParamsUnion{
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
				NumExamples: llamastackclient.Int(0),
			},
			InputRows: []map[string]llamastackclient.EvalBenchmarkEvaluateParamsInputRowUnion{{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			}},
			ScoringFunctions: []string{"string"},
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

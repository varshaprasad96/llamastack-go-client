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

func TestEvalBenchmarkJobGet(t *testing.T) {
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
	_, err := client.Eval.Benchmarks.Jobs.Get(
		context.TODO(),
		"job_id",
		llamastackclient.EvalBenchmarkJobGetParams{
			BenchmarkID: "benchmark_id",
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

func TestEvalBenchmarkJobCancel(t *testing.T) {
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
	err := client.Eval.Benchmarks.Jobs.Cancel(
		context.TODO(),
		"job_id",
		llamastackclient.EvalBenchmarkJobCancelParams{
			BenchmarkID: "benchmark_id",
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

func TestEvalBenchmarkJobResult(t *testing.T) {
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
	_, err := client.Eval.Benchmarks.Jobs.Result(
		context.TODO(),
		"job_id",
		llamastackclient.EvalBenchmarkJobResultParams{
			BenchmarkID: "benchmark_id",
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

func TestEvalBenchmarkJobRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Eval.Benchmarks.Jobs.Run(
		context.TODO(),
		"benchmark_id",
		llamastackclient.EvalBenchmarkJobRunParams{
			BenchmarkConfig: llamastackclient.BenchmarkConfigParam{
				EvalCandidate: llamastackclient.BenchmarkConfigEvalCandidateUnionParam{
					OfModel: &llamastackclient.BenchmarkConfigEvalCandidateModelParam{
						Model: "model",
						SamplingParams: llamastackclient.SamplingParams{
							Strategy: llamastackclient.SamplingParamsStrategyUnion{
								OfGreedy: &llamastackclient.SamplingParamsStrategyGreedy{},
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

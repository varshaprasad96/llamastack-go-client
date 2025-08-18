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

func TestEvalBenchmarkJobGet(t *testing.T) {
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
	_, err := client.Eval.Benchmarks.Jobs.Get(
		context.TODO(),
		"job_id",
		llamastackgoclient.EvalBenchmarkJobGetParams{
			BenchmarkID: "benchmark_id",
		},
	)
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Eval.Benchmarks.Jobs.Cancel(
		context.TODO(),
		"job_id",
		llamastackgoclient.EvalBenchmarkJobCancelParams{
			BenchmarkID: "benchmark_id",
		},
	)
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Eval.Benchmarks.Jobs.Result(
		context.TODO(),
		"job_id",
		llamastackgoclient.EvalBenchmarkJobResultParams{
			BenchmarkID: "benchmark_id",
		},
	)
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Eval.Benchmarks.Jobs.Run(
		context.TODO(),
		"benchmark_id",
		llamastackgoclient.EvalBenchmarkJobRunParams{
			BenchmarkConfig: llamastackgoclient.BenchmarkConfigParam{
				EvalCandidate: llamastackgoclient.BenchmarkConfigEvalCandidateUnionParam{
					OfModel: &llamastackgoclient.BenchmarkConfigEvalCandidateModelParam{
						Model: "model",
						SamplingParams: llamastackgoclient.SamplingParams{
							Strategy: llamastackgoclient.SamplingParamsStrategyUnion{
								OfGreedy: &llamastackgoclient.SamplingParamsStrategyGreedy{},
							},
							MaxTokens:         llamastackgoclient.Int(0),
							RepetitionPenalty: llamastackgoclient.Float(0),
							Stop:              []string{"string"},
						},
						SystemMessage: llamastackgoclient.SystemMessageParam{
							Content: llamastackgoclient.InterleavedContentUnionParam{
								OfString: llamastackgoclient.String("string"),
							},
						},
					},
				},
				ScoringParams: map[string]llamastackgoclient.ScoringFnParamsUnion{
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
				NumExamples: llamastackgoclient.Int(0),
			},
		},
	)
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

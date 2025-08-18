// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/llamastack-go-client-go/internal/apijson"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/param"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/respjson"
	"github.com/stainless-sdks/llamastack-go-client-go/shared/constant"
)

// EvalBenchmarkService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalBenchmarkService] method instead.
type EvalBenchmarkService struct {
	Options []option.RequestOption
	Jobs    EvalBenchmarkJobService
}

// NewEvalBenchmarkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvalBenchmarkService(opts ...option.RequestOption) (r EvalBenchmarkService) {
	r = EvalBenchmarkService{}
	r.Options = opts
	r.Jobs = NewEvalBenchmarkJobService(opts...)
	return
}

// Register a benchmark.
func (r *EvalBenchmarkService) New(ctx context.Context, body EvalBenchmarkNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/eval/benchmarks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get a benchmark by its ID.
func (r *EvalBenchmarkService) Get(ctx context.Context, benchmarkID string, opts ...option.RequestOption) (res *Benchmark, err error) {
	opts = append(r.Options[:], opts...)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all benchmarks.
func (r *EvalBenchmarkService) List(ctx context.Context, opts ...option.RequestOption) (res *EvalBenchmarkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/eval/benchmarks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Evaluate a list of rows on a benchmark.
func (r *EvalBenchmarkService) Evaluate(ctx context.Context, benchmarkID string, body EvalBenchmarkEvaluateParams, opts ...option.RequestOption) (res *EvaluateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/evaluations", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A benchmark resource for evaluating model performance.
type Benchmark struct {
	// Identifier of the dataset to use for the benchmark evaluation
	DatasetID  string `json:"dataset_id,required"`
	Identifier string `json:"identifier,required"`
	// Metadata for this evaluation task
	Metadata   map[string]BenchmarkMetadataUnion `json:"metadata,required"`
	ProviderID string                            `json:"provider_id,required"`
	// List of scoring function identifiers to apply during evaluation
	ScoringFunctions []string `json:"scoring_functions,required"`
	// The resource type, always benchmark
	Type               constant.Benchmark `json:"type,required"`
	ProviderResourceID string             `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DatasetID          respjson.Field
		Identifier         respjson.Field
		Metadata           respjson.Field
		ProviderID         respjson.Field
		ScoringFunctions   respjson.Field
		Type               respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Benchmark) RawJSON() string { return r.JSON.raw }
func (r *Benchmark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BenchmarkMetadataUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type BenchmarkMetadataUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u BenchmarkMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BenchmarkMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BenchmarkMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BenchmarkMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BenchmarkMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *BenchmarkMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A benchmark configuration for evaluation.
//
// The properties EvalCandidate, ScoringParams are required.
type BenchmarkConfigParam struct {
	// The candidate to evaluate.
	EvalCandidate BenchmarkConfigEvalCandidateUnionParam `json:"eval_candidate,omitzero,required"`
	// Map between scoring function id and parameters for each scoring function you
	// want to run
	ScoringParams map[string]ScoringFnParamsUnion `json:"scoring_params,omitzero,required"`
	// (Optional) The number of examples to evaluate. If not provided, all examples in
	// the dataset will be evaluated
	NumExamples param.Opt[int64] `json:"num_examples,omitzero"`
	paramObj
}

func (r BenchmarkConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BenchmarkConfigEvalCandidateUnionParam struct {
	OfModel *BenchmarkConfigEvalCandidateModelParam `json:",omitzero,inline"`
	OfAgent *BenchmarkConfigEvalCandidateAgentParam `json:",omitzero,inline"`
	paramUnion
}

func (u BenchmarkConfigEvalCandidateUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModel, u.OfAgent)
}
func (u *BenchmarkConfigEvalCandidateUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BenchmarkConfigEvalCandidateUnionParam) asAny() any {
	if !param.IsOmitted(u.OfModel) {
		return u.OfModel
	} else if !param.IsOmitted(u.OfAgent) {
		return u.OfAgent
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigEvalCandidateUnionParam) GetModel() *string {
	if vt := u.OfModel; vt != nil {
		return &vt.Model
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigEvalCandidateUnionParam) GetSamplingParams() *SamplingParams {
	if vt := u.OfModel; vt != nil {
		return &vt.SamplingParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigEvalCandidateUnionParam) GetSystemMessage() *SystemMessageParam {
	if vt := u.OfModel; vt != nil {
		return &vt.SystemMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigEvalCandidateUnionParam) GetConfig() *AgentConfigParam {
	if vt := u.OfAgent; vt != nil {
		return &vt.Config
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BenchmarkConfigEvalCandidateUnionParam) GetType() *string {
	if vt := u.OfModel; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAgent; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BenchmarkConfigEvalCandidateUnionParam](
		"type",
		apijson.Discriminator[BenchmarkConfigEvalCandidateModelParam]("model"),
		apijson.Discriminator[BenchmarkConfigEvalCandidateAgentParam]("agent"),
	)
}

// A model candidate for evaluation.
//
// The properties Model, SamplingParams, Type are required.
type BenchmarkConfigEvalCandidateModelParam struct {
	// The model ID to evaluate.
	Model string `json:"model,required"`
	// The sampling parameters for the model.
	SamplingParams SamplingParams `json:"sampling_params,omitzero,required"`
	// (Optional) The system message providing instructions or context to the model.
	SystemMessage SystemMessageParam `json:"system_message,omitzero"`
	// This field can be elided, and will marshal its zero value as "model".
	Type constant.Model `json:"type,required"`
	paramObj
}

func (r BenchmarkConfigEvalCandidateModelParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigEvalCandidateModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigEvalCandidateModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An agent candidate for evaluation.
//
// The properties Config, Type are required.
type BenchmarkConfigEvalCandidateAgentParam struct {
	// The configuration for the agent candidate.
	Config AgentConfigParam `json:"config,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "agent".
	Type constant.Agent `json:"type,required"`
	paramObj
}

func (r BenchmarkConfigEvalCandidateAgentParam) MarshalJSON() (data []byte, err error) {
	type shadow BenchmarkConfigEvalCandidateAgentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BenchmarkConfigEvalCandidateAgentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response from an evaluation.
type EvaluateResponse struct {
	// The generations from the evaluation.
	Generations []map[string]EvaluateResponseGenerationUnion `json:"generations,required"`
	// The scores from the evaluation.
	Scores map[string]EvaluateResponseScore `json:"scores,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Generations respjson.Field
		Scores      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluateResponse) RawJSON() string { return r.JSON.raw }
func (r *EvaluateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EvaluateResponseGenerationUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type EvaluateResponseGenerationUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u EvaluateResponseGenerationUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseGenerationUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseGenerationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseGenerationUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvaluateResponseGenerationUnion) RawJSON() string { return u.JSON.raw }

func (r *EvaluateResponseGenerationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A scoring result for a single row.
type EvaluateResponseScore struct {
	// Map of metric name to aggregated value
	AggregatedResults map[string]EvaluateResponseScoreAggregatedResultUnion `json:"aggregated_results,required"`
	// The scoring result for each row. Each row is a map of column name to value.
	ScoreRows []map[string]EvaluateResponseScoreScoreRowUnion `json:"score_rows,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregatedResults respjson.Field
		ScoreRows         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluateResponseScore) RawJSON() string { return r.JSON.raw }
func (r *EvaluateResponseScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EvaluateResponseScoreAggregatedResultUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type EvaluateResponseScoreAggregatedResultUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u EvaluateResponseScoreAggregatedResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseScoreAggregatedResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseScoreAggregatedResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseScoreAggregatedResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvaluateResponseScoreAggregatedResultUnion) RawJSON() string { return u.JSON.raw }

func (r *EvaluateResponseScoreAggregatedResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EvaluateResponseScoreScoreRowUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type EvaluateResponseScoreScoreRowUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u EvaluateResponseScoreScoreRowUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseScoreScoreRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseScoreScoreRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluateResponseScoreScoreRowUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvaluateResponseScoreScoreRowUnion) RawJSON() string { return u.JSON.raw }

func (r *EvaluateResponseScoreScoreRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalBenchmarkListResponse struct {
	Data []Benchmark `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalBenchmarkListResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalBenchmarkListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalBenchmarkNewParams struct {
	// The ID of the benchmark to register.
	BenchmarkID string `json:"benchmark_id,required"`
	// The ID of the dataset to use for the benchmark.
	DatasetID string `json:"dataset_id,required"`
	// The scoring functions to use for the benchmark.
	ScoringFunctions []string `json:"scoring_functions,omitzero,required"`
	// The ID of the provider benchmark to use for the benchmark.
	ProviderBenchmarkID param.Opt[string] `json:"provider_benchmark_id,omitzero"`
	// The ID of the provider to use for the benchmark.
	ProviderID param.Opt[string] `json:"provider_id,omitzero"`
	// The metadata to use for the benchmark.
	Metadata map[string]EvalBenchmarkNewParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r EvalBenchmarkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalBenchmarkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalBenchmarkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalBenchmarkNewParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u EvalBenchmarkNewParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *EvalBenchmarkNewParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalBenchmarkNewParamsMetadataUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

type EvalBenchmarkEvaluateParams struct {
	// The configuration for the benchmark.
	BenchmarkConfig BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	// The rows to evaluate.
	InputRows []map[string]EvalBenchmarkEvaluateParamsInputRowUnion `json:"input_rows,omitzero,required"`
	// The scoring functions to use for the evaluation.
	ScoringFunctions []string `json:"scoring_functions,omitzero,required"`
	paramObj
}

func (r EvalBenchmarkEvaluateParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalBenchmarkEvaluateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalBenchmarkEvaluateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalBenchmarkEvaluateParamsInputRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u EvalBenchmarkEvaluateParamsInputRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *EvalBenchmarkEvaluateParamsInputRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalBenchmarkEvaluateParamsInputRowUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

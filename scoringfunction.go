// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// ScoringFunctionService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScoringFunctionService] method instead.
type ScoringFunctionService struct {
	Options []option.RequestOption
}

// NewScoringFunctionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScoringFunctionService(opts ...option.RequestOption) (r ScoringFunctionService) {
	r = ScoringFunctionService{}
	r.Options = opts
	return
}

// Register a scoring function.
func (r *ScoringFunctionService) New(ctx context.Context, body ScoringFunctionNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/scoring-functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get a scoring function by its ID.
func (r *ScoringFunctionService) Get(ctx context.Context, scoringFnID string, opts ...option.RequestOption) (res *ScoringFn, err error) {
	opts = append(r.Options[:], opts...)
	if scoringFnID == "" {
		err = errors.New("missing required scoring_fn_id parameter")
		return
	}
	path := fmt.Sprintf("v1/scoring-functions/%s", scoringFnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all scoring functions.
func (r *ScoringFunctionService) List(ctx context.Context, opts ...option.RequestOption) (res *ScoringFunctionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/scoring-functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AggregationFunctionType string

const (
	AggregationFunctionTypeAverage          AggregationFunctionType = "average"
	AggregationFunctionTypeWeightedAverage  AggregationFunctionType = "weighted_average"
	AggregationFunctionTypeMedian           AggregationFunctionType = "median"
	AggregationFunctionTypeCategoricalCount AggregationFunctionType = "categorical_count"
	AggregationFunctionTypeAccuracy         AggregationFunctionType = "accuracy"
)

type ScoringFn struct {
	Identifier string                            `json:"identifier,required"`
	Metadata   map[string]ScoringFnMetadataUnion `json:"metadata,required"`
	ProviderID string                            `json:"provider_id,required"`
	ReturnType ScoringFnReturnType               `json:"return_type,required"`
	// Any of "model", "shield", "vector_db", "dataset", "scoring_function",
	// "benchmark", "tool", "tool_group".
	Type               ScoringFnType            `json:"type,required"`
	Description        string                   `json:"description"`
	Params             ScoringFnParamsUnionResp `json:"params"`
	ProviderResourceID string                   `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		Metadata           respjson.Field
		ProviderID         respjson.Field
		ReturnType         respjson.Field
		Type               respjson.Field
		Description        respjson.Field
		Params             respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFn) RawJSON() string { return r.JSON.raw }
func (r *ScoringFn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoringFnMetadataUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ScoringFnMetadataUnion struct {
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

func (u ScoringFnMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringFnMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ScoringFnMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringFnReturnType struct {
	// Any of "string", "number", "boolean", "array", "object", "json", "union",
	// "chat_completion_input", "completion_input", "agent_turn_input".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnReturnType) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnReturnType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringFnType string

const (
	ScoringFnTypeModel           ScoringFnType = "model"
	ScoringFnTypeShield          ScoringFnType = "shield"
	ScoringFnTypeVectorDB        ScoringFnType = "vector_db"
	ScoringFnTypeDataset         ScoringFnType = "dataset"
	ScoringFnTypeScoringFunction ScoringFnType = "scoring_function"
	ScoringFnTypeBenchmark       ScoringFnType = "benchmark"
	ScoringFnTypeTool            ScoringFnType = "tool"
	ScoringFnTypeToolGroup       ScoringFnType = "tool_group"
)

// ScoringFnParamsUnionResp contains all possible properties and values from
// [ScoringFnParamsLlmAsJudgeScoringFnParamsResp],
// [ScoringFnParamsRegexParserScoringFnParamsResp],
// [ScoringFnParamsBasicScoringFnParamsResp].
//
// Use the [ScoringFnParamsUnionResp.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ScoringFnParamsUnionResp struct {
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions"`
	// This field is from variant [ScoringFnParamsLlmAsJudgeScoringFnParamsResp].
	JudgeModel string `json:"judge_model"`
	// This field is from variant [ScoringFnParamsLlmAsJudgeScoringFnParamsResp].
	JudgeScoreRegexes []string `json:"judge_score_regexes"`
	// Any of nil, nil, nil.
	Type ScoringFnParamsType `json:"type"`
	// This field is from variant [ScoringFnParamsLlmAsJudgeScoringFnParamsResp].
	PromptTemplate string `json:"prompt_template"`
	// This field is from variant [ScoringFnParamsRegexParserScoringFnParamsResp].
	ParsingRegexes []string `json:"parsing_regexes"`
	JSON           struct {
		AggregationFunctions respjson.Field
		JudgeModel           respjson.Field
		JudgeScoreRegexes    respjson.Field
		Type                 respjson.Field
		PromptTemplate       respjson.Field
		ParsingRegexes       respjson.Field
		raw                  string
	} `json:"-"`
}

func (u ScoringFnParamsUnionResp) AsLlmAsJudgeScoringFnParams() (v ScoringFnParamsLlmAsJudgeScoringFnParamsResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnParamsUnionResp) AsRegexParserScoringFnParams() (v ScoringFnParamsRegexParserScoringFnParamsResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnParamsUnionResp) AsBasicScoringFnParams() (v ScoringFnParamsBasicScoringFnParamsResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringFnParamsUnionResp) RawJSON() string { return u.JSON.raw }

func (r *ScoringFnParamsUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ScoringFnParamsUnionResp to a ScoringFnParamsUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ScoringFnParamsUnion.Overrides()
func (r ScoringFnParamsUnionResp) ToParam() ScoringFnParamsUnion {
	return param.Override[ScoringFnParamsUnion](json.RawMessage(r.RawJSON()))
}

type ScoringFnParamsLlmAsJudgeScoringFnParamsResp struct {
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,required"`
	JudgeModel           string                    `json:"judge_model,required"`
	JudgeScoreRegexes    []string                  `json:"judge_score_regexes,required"`
	// Any of "llm_as_judge", "regex_parser", "basic".
	Type           ScoringFnParamsType `json:"type,required"`
	PromptTemplate string              `json:"prompt_template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregationFunctions respjson.Field
		JudgeModel           respjson.Field
		JudgeScoreRegexes    respjson.Field
		Type                 respjson.Field
		PromptTemplate       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnParamsLlmAsJudgeScoringFnParamsResp) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnParamsLlmAsJudgeScoringFnParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringFnParamsRegexParserScoringFnParamsResp struct {
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,required"`
	ParsingRegexes       []string                  `json:"parsing_regexes,required"`
	// Any of "llm_as_judge", "regex_parser", "basic".
	Type ScoringFnParamsType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregationFunctions respjson.Field
		ParsingRegexes       respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnParamsRegexParserScoringFnParamsResp) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnParamsRegexParserScoringFnParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringFnParamsBasicScoringFnParamsResp struct {
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,required"`
	// Any of "llm_as_judge", "regex_parser", "basic".
	Type ScoringFnParamsType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregationFunctions respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnParamsBasicScoringFnParamsResp) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnParamsBasicScoringFnParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ScoringFnParamsOfRegexParserScoringFns(aggregationFunctions []AggregationFunctionType, parsingRegexes []string, type_ ScoringFnParamsType) ScoringFnParamsUnion {
	var variant ScoringFnParamsRegexParserScoringFnParams
	variant.AggregationFunctions = aggregationFunctions
	variant.ParsingRegexes = parsingRegexes
	variant.Type = type_
	return ScoringFnParamsUnion{OfRegexParserScoringFns: &variant}
}

func ScoringFnParamsOfBasicScoringFns(aggregationFunctions []AggregationFunctionType, type_ ScoringFnParamsType) ScoringFnParamsUnion {
	var variant ScoringFnParamsBasicScoringFnParams
	variant.AggregationFunctions = aggregationFunctions
	variant.Type = type_
	return ScoringFnParamsUnion{OfBasicScoringFns: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScoringFnParamsUnion struct {
	OfLlmAsJudgeScoringFns  *ScoringFnParamsLlmAsJudgeScoringFnParams  `json:",omitzero,inline"`
	OfRegexParserScoringFns *ScoringFnParamsRegexParserScoringFnParams `json:",omitzero,inline"`
	OfBasicScoringFns       *ScoringFnParamsBasicScoringFnParams       `json:",omitzero,inline"`
	paramUnion
}

func (u ScoringFnParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLlmAsJudgeScoringFns, u.OfRegexParserScoringFns, u.OfBasicScoringFns)
}
func (u *ScoringFnParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScoringFnParamsUnion) asAny() any {
	if !param.IsOmitted(u.OfLlmAsJudgeScoringFns) {
		return u.OfLlmAsJudgeScoringFns
	} else if !param.IsOmitted(u.OfRegexParserScoringFns) {
		return u.OfRegexParserScoringFns
	} else if !param.IsOmitted(u.OfBasicScoringFns) {
		return u.OfBasicScoringFns
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringFnParamsUnion) GetJudgeModel() *string {
	if vt := u.OfLlmAsJudgeScoringFns; vt != nil {
		return &vt.JudgeModel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringFnParamsUnion) GetJudgeScoreRegexes() []string {
	if vt := u.OfLlmAsJudgeScoringFns; vt != nil {
		return vt.JudgeScoreRegexes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringFnParamsUnion) GetPromptTemplate() *string {
	if vt := u.OfLlmAsJudgeScoringFns; vt != nil && vt.PromptTemplate.Valid() {
		return &vt.PromptTemplate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringFnParamsUnion) GetParsingRegexes() []string {
	if vt := u.OfRegexParserScoringFns; vt != nil {
		return vt.ParsingRegexes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringFnParamsUnion) GetType() *string {
	if vt := u.OfLlmAsJudgeScoringFns; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRegexParserScoringFns; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBasicScoringFns; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's AggregationFunctions property, if
// present.
func (u ScoringFnParamsUnion) GetAggregationFunctions() []AggregationFunctionType {
	if vt := u.OfLlmAsJudgeScoringFns; vt != nil {
		return vt.AggregationFunctions
	} else if vt := u.OfRegexParserScoringFns; vt != nil {
		return vt.AggregationFunctions
	} else if vt := u.OfBasicScoringFns; vt != nil {
		return vt.AggregationFunctions
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ScoringFnParamsUnion](
		"type",
		apijson.Discriminator[ScoringFnParamsLlmAsJudgeScoringFnParams]("llm_as_judge"),
		apijson.Discriminator[ScoringFnParamsLlmAsJudgeScoringFnParams]("regex_parser"),
		apijson.Discriminator[ScoringFnParamsLlmAsJudgeScoringFnParams]("basic"),
		apijson.Discriminator[ScoringFnParamsRegexParserScoringFnParams]("llm_as_judge"),
		apijson.Discriminator[ScoringFnParamsRegexParserScoringFnParams]("regex_parser"),
		apijson.Discriminator[ScoringFnParamsRegexParserScoringFnParams]("basic"),
		apijson.Discriminator[ScoringFnParamsBasicScoringFnParams]("llm_as_judge"),
		apijson.Discriminator[ScoringFnParamsBasicScoringFnParams]("regex_parser"),
		apijson.Discriminator[ScoringFnParamsBasicScoringFnParams]("basic"),
	)
}

// The properties AggregationFunctions, JudgeModel, JudgeScoreRegexes, Type are
// required.
type ScoringFnParamsLlmAsJudgeScoringFnParams struct {
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,omitzero,required"`
	JudgeModel           string                    `json:"judge_model,required"`
	JudgeScoreRegexes    []string                  `json:"judge_score_regexes,omitzero,required"`
	// Any of "llm_as_judge", "regex_parser", "basic".
	Type           ScoringFnParamsType `json:"type,omitzero,required"`
	PromptTemplate param.Opt[string]   `json:"prompt_template,omitzero"`
	paramObj
}

func (r ScoringFnParamsLlmAsJudgeScoringFnParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFnParamsLlmAsJudgeScoringFnParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFnParamsLlmAsJudgeScoringFnParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AggregationFunctions, ParsingRegexes, Type are required.
type ScoringFnParamsRegexParserScoringFnParams struct {
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,omitzero,required"`
	ParsingRegexes       []string                  `json:"parsing_regexes,omitzero,required"`
	// Any of "llm_as_judge", "regex_parser", "basic".
	Type ScoringFnParamsType `json:"type,omitzero,required"`
	paramObj
}

func (r ScoringFnParamsRegexParserScoringFnParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFnParamsRegexParserScoringFnParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFnParamsRegexParserScoringFnParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AggregationFunctions, Type are required.
type ScoringFnParamsBasicScoringFnParams struct {
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,omitzero,required"`
	// Any of "llm_as_judge", "regex_parser", "basic".
	Type ScoringFnParamsType `json:"type,omitzero,required"`
	paramObj
}

func (r ScoringFnParamsBasicScoringFnParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFnParamsBasicScoringFnParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFnParamsBasicScoringFnParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringFnParamsType string

const (
	ScoringFnParamsTypeLlmAsJudge  ScoringFnParamsType = "llm_as_judge"
	ScoringFnParamsTypeRegexParser ScoringFnParamsType = "regex_parser"
	ScoringFnParamsTypeBasic       ScoringFnParamsType = "basic"
)

type ScoringFunctionListResponse struct {
	Data []ScoringFn `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFunctionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ScoringFunctionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringFunctionNewParams struct {
	// The description of the scoring function.
	Description string                             `json:"description,required"`
	ReturnType  ScoringFunctionNewParamsReturnType `json:"return_type,omitzero,required"`
	// The ID of the scoring function to register.
	ScoringFnID string `json:"scoring_fn_id,required"`
	// The ID of the provider to use for the scoring function.
	ProviderID param.Opt[string] `json:"provider_id,omitzero"`
	// The ID of the provider scoring function to use for the scoring function.
	ProviderScoringFnID param.Opt[string] `json:"provider_scoring_fn_id,omitzero"`
	// The parameters for the scoring function for benchmark eval, these can be
	// overridden for app eval.
	Params ScoringFnParamsUnion `json:"params,omitzero"`
	paramObj
}

func (r ScoringFunctionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFunctionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFunctionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ScoringFunctionNewParamsReturnType struct {
	// Any of "string", "number", "boolean", "array", "object", "json", "union",
	// "chat_completion_input", "completion_input", "agent_turn_input".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r ScoringFunctionNewParamsReturnType) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFunctionNewParamsReturnType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFunctionNewParamsReturnType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ScoringFunctionNewParamsReturnType](
		"type", "string", "number", "boolean", "array", "object", "json", "union", "chat_completion_input", "completion_input", "agent_turn_input",
	)
}

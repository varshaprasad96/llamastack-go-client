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

// Types of aggregation functions for scoring results.
type AggregationFunctionType string

const (
	AggregationFunctionTypeAverage          AggregationFunctionType = "average"
	AggregationFunctionTypeWeightedAverage  AggregationFunctionType = "weighted_average"
	AggregationFunctionTypeMedian           AggregationFunctionType = "median"
	AggregationFunctionTypeCategoricalCount AggregationFunctionType = "categorical_count"
	AggregationFunctionTypeAccuracy         AggregationFunctionType = "accuracy"
)

// ParamTypeUnionResp contains all possible properties and values from
// [ParamTypeStringResp], [ParamTypeNumberResp], [ParamTypeBooleanResp],
// [ParamTypeArrayResp], [ParamTypeObjectResp], [ParamTypeJsonResp],
// [ParamTypeUnionResp], [ParamTypeChatCompletionInputResp],
// [ParamTypeCompletionInputResp], [ParamTypeAgentTurnInputResp].
//
// Use the [ParamTypeUnionResp.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ParamTypeUnionResp struct {
	// Any of "string", "number", "boolean", "array", "object", "json", "union",
	// "chat_completion_input", "completion_input", "agent_turn_input".
	Type string `json:"type"`
	JSON struct {
		Type respjson.Field
		raw  string
	} `json:"-"`
}

// anyParamTypeResp is implemented by each variant of [ParamTypeUnionResp] to add
// type safety for the return type of [ParamTypeUnionResp.AsAny]
type anyParamTypeResp interface {
	implParamTypeUnionResp()
}

func (ParamTypeStringResp) implParamTypeUnionResp()              {}
func (ParamTypeNumberResp) implParamTypeUnionResp()              {}
func (ParamTypeBooleanResp) implParamTypeUnionResp()             {}
func (ParamTypeArrayResp) implParamTypeUnionResp()               {}
func (ParamTypeObjectResp) implParamTypeUnionResp()              {}
func (ParamTypeJsonResp) implParamTypeUnionResp()                {}
func (ParamTypeUnionResp) implParamTypeUnionResp()               {}
func (ParamTypeChatCompletionInputResp) implParamTypeUnionResp() {}
func (ParamTypeCompletionInputResp) implParamTypeUnionResp()     {}
func (ParamTypeAgentTurnInputResp) implParamTypeUnionResp()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ParamTypeUnionResp.AsAny().(type) {
//	case llamastackgoclient.ParamTypeStringResp:
//	case llamastackgoclient.ParamTypeNumberResp:
//	case llamastackgoclient.ParamTypeBooleanResp:
//	case llamastackgoclient.ParamTypeArrayResp:
//	case llamastackgoclient.ParamTypeObjectResp:
//	case llamastackgoclient.ParamTypeJsonResp:
//	case llamastackgoclient.ParamTypeUnionResp:
//	case llamastackgoclient.ParamTypeChatCompletionInputResp:
//	case llamastackgoclient.ParamTypeCompletionInputResp:
//	case llamastackgoclient.ParamTypeAgentTurnInputResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ParamTypeUnionResp) AsAny() anyParamTypeResp {
	switch u.Type {
	case "string":
		return u.AsString()
	case "number":
		return u.AsNumber()
	case "boolean":
		return u.AsBoolean()
	case "array":
		return u.AsArray()
	case "object":
		return u.AsObject()
	case "json":
		return u.AsJson()
	case "union":
		return u.AsUnion()
	case "chat_completion_input":
		return u.AsChatCompletionInput()
	case "completion_input":
		return u.AsCompletionInput()
	case "agent_turn_input":
		return u.AsAgentTurnInput()
	}
	return nil
}

func (u ParamTypeUnionResp) AsString() (v ParamTypeStringResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsNumber() (v ParamTypeNumberResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsBoolean() (v ParamTypeBooleanResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsArray() (v ParamTypeArrayResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsObject() (v ParamTypeObjectResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsJson() (v ParamTypeJsonResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsUnion() (v ParamTypeUnionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsChatCompletionInput() (v ParamTypeChatCompletionInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsCompletionInput() (v ParamTypeCompletionInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsAgentTurnInput() (v ParamTypeAgentTurnInputResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParamTypeUnionResp) RawJSON() string { return u.JSON.raw }

func (r *ParamTypeUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ParamTypeUnionResp to a ParamTypeUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ParamTypeUnion.Overrides()
func (r ParamTypeUnionResp) ToParam() ParamTypeUnion {
	return param.Override[ParamTypeUnion](json.RawMessage(r.RawJSON()))
}

// Parameter type for string values.
type ParamTypeStringResp struct {
	// Discriminator type. Always "string"
	Type constant.String `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParamTypeStringResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeStringResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for numeric values.
type ParamTypeNumberResp struct {
	// Discriminator type. Always "number"
	Type constant.Number `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParamTypeNumberResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeNumberResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for boolean values.
type ParamTypeBooleanResp struct {
	// Discriminator type. Always "boolean"
	Type constant.Boolean `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParamTypeBooleanResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeBooleanResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for array values.
type ParamTypeArrayResp struct {
	// Discriminator type. Always "array"
	Type constant.Array `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParamTypeArrayResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeArrayResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for object values.
type ParamTypeObjectResp struct {
	// Discriminator type. Always "object"
	Type constant.Object `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParamTypeObjectResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeObjectResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for JSON values.
type ParamTypeJsonResp struct {
	// Discriminator type. Always "json"
	Type constant.Json `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParamTypeJsonResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeJsonResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for chat completion input.
type ParamTypeChatCompletionInputResp struct {
	// Discriminator type. Always "chat_completion_input"
	Type constant.ChatCompletionInput `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParamTypeChatCompletionInputResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeChatCompletionInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for completion input.
type ParamTypeCompletionInputResp struct {
	// Discriminator type. Always "completion_input"
	Type constant.CompletionInput `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParamTypeCompletionInputResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeCompletionInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for agent turn input.
type ParamTypeAgentTurnInputResp struct {
	// Discriminator type. Always "agent_turn_input"
	Type constant.AgentTurnInput `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParamTypeAgentTurnInputResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeAgentTurnInputResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParamTypeUnion struct {
	OfString              *ParamTypeString              `json:",omitzero,inline"`
	OfNumber              *ParamTypeNumber              `json:",omitzero,inline"`
	OfBoolean             *ParamTypeBoolean             `json:",omitzero,inline"`
	OfArray               *ParamTypeArray               `json:",omitzero,inline"`
	OfObject              *ParamTypeObject              `json:",omitzero,inline"`
	OfJson                *ParamTypeJson                `json:",omitzero,inline"`
	OfUnion               *ParamTypeUnion               `json:",omitzero,inline"`
	OfChatCompletionInput *ParamTypeChatCompletionInput `json:",omitzero,inline"`
	OfCompletionInput     *ParamTypeCompletionInput     `json:",omitzero,inline"`
	OfAgentTurnInput      *ParamTypeAgentTurnInput      `json:",omitzero,inline"`
	paramUnion
}

func (u ParamTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfNumber,
		u.OfBoolean,
		u.OfArray,
		u.OfObject,
		u.OfJson,
		u.OfUnion,
		u.OfChatCompletionInput,
		u.OfCompletionInput,
		u.OfAgentTurnInput)
}
func (u *ParamTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ParamTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return u.OfString
	} else if !param.IsOmitted(u.OfNumber) {
		return u.OfNumber
	} else if !param.IsOmitted(u.OfBoolean) {
		return u.OfBoolean
	} else if !param.IsOmitted(u.OfArray) {
		return u.OfArray
	} else if !param.IsOmitted(u.OfObject) {
		return u.OfObject
	} else if !param.IsOmitted(u.OfJson) {
		return u.OfJson
	} else if !param.IsOmitted(u.OfUnion) {
		return u.OfUnion
	} else if !param.IsOmitted(u.OfChatCompletionInput) {
		return u.OfChatCompletionInput
	} else if !param.IsOmitted(u.OfCompletionInput) {
		return u.OfCompletionInput
	} else if !param.IsOmitted(u.OfAgentTurnInput) {
		return u.OfAgentTurnInput
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ParamTypeUnion) GetType() *string {
	if vt := u.OfString; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNumber; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBoolean; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfArray; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfObject; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJson; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfUnion; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfChatCompletionInput; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCompletionInput; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAgentTurnInput; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ParamTypeUnion](
		"type",
		apijson.Discriminator[ParamTypeString]("string"),
		apijson.Discriminator[ParamTypeNumber]("number"),
		apijson.Discriminator[ParamTypeBoolean]("boolean"),
		apijson.Discriminator[ParamTypeArray]("array"),
		apijson.Discriminator[ParamTypeObject]("object"),
		apijson.Discriminator[ParamTypeJson]("json"),
		apijson.Discriminator[ParamTypeUnion]("union"),
		apijson.Discriminator[ParamTypeChatCompletionInput]("chat_completion_input"),
		apijson.Discriminator[ParamTypeCompletionInput]("completion_input"),
		apijson.Discriminator[ParamTypeAgentTurnInput]("agent_turn_input"),
	)
}

func NewParamTypeString() ParamTypeString {
	return ParamTypeString{
		Type: "string",
	}
}

// Parameter type for string values.
//
// This struct has a constant value, construct it with [NewParamTypeString].
type ParamTypeString struct {
	// Discriminator type. Always "string"
	Type constant.String `json:"type,required"`
	paramObj
}

func (r ParamTypeString) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParamTypeString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewParamTypeNumber() ParamTypeNumber {
	return ParamTypeNumber{
		Type: "number",
	}
}

// Parameter type for numeric values.
//
// This struct has a constant value, construct it with [NewParamTypeNumber].
type ParamTypeNumber struct {
	// Discriminator type. Always "number"
	Type constant.Number `json:"type,required"`
	paramObj
}

func (r ParamTypeNumber) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeNumber
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParamTypeNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewParamTypeBoolean() ParamTypeBoolean {
	return ParamTypeBoolean{
		Type: "boolean",
	}
}

// Parameter type for boolean values.
//
// This struct has a constant value, construct it with [NewParamTypeBoolean].
type ParamTypeBoolean struct {
	// Discriminator type. Always "boolean"
	Type constant.Boolean `json:"type,required"`
	paramObj
}

func (r ParamTypeBoolean) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeBoolean
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParamTypeBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewParamTypeArray() ParamTypeArray {
	return ParamTypeArray{
		Type: "array",
	}
}

// Parameter type for array values.
//
// This struct has a constant value, construct it with [NewParamTypeArray].
type ParamTypeArray struct {
	// Discriminator type. Always "array"
	Type constant.Array `json:"type,required"`
	paramObj
}

func (r ParamTypeArray) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeArray
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParamTypeArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewParamTypeObject() ParamTypeObject {
	return ParamTypeObject{
		Type: "object",
	}
}

// Parameter type for object values.
//
// This struct has a constant value, construct it with [NewParamTypeObject].
type ParamTypeObject struct {
	// Discriminator type. Always "object"
	Type constant.Object `json:"type,required"`
	paramObj
}

func (r ParamTypeObject) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParamTypeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewParamTypeJson() ParamTypeJson {
	return ParamTypeJson{
		Type: "json",
	}
}

// Parameter type for JSON values.
//
// This struct has a constant value, construct it with [NewParamTypeJson].
type ParamTypeJson struct {
	// Discriminator type. Always "json"
	Type constant.Json `json:"type,required"`
	paramObj
}

func (r ParamTypeJson) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeJson
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParamTypeJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewParamTypeChatCompletionInput() ParamTypeChatCompletionInput {
	return ParamTypeChatCompletionInput{
		Type: "chat_completion_input",
	}
}

// Parameter type for chat completion input.
//
// This struct has a constant value, construct it with
// [NewParamTypeChatCompletionInput].
type ParamTypeChatCompletionInput struct {
	// Discriminator type. Always "chat_completion_input"
	Type constant.ChatCompletionInput `json:"type,required"`
	paramObj
}

func (r ParamTypeChatCompletionInput) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeChatCompletionInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParamTypeChatCompletionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewParamTypeCompletionInput() ParamTypeCompletionInput {
	return ParamTypeCompletionInput{
		Type: "completion_input",
	}
}

// Parameter type for completion input.
//
// This struct has a constant value, construct it with
// [NewParamTypeCompletionInput].
type ParamTypeCompletionInput struct {
	// Discriminator type. Always "completion_input"
	Type constant.CompletionInput `json:"type,required"`
	paramObj
}

func (r ParamTypeCompletionInput) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeCompletionInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParamTypeCompletionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewParamTypeAgentTurnInput() ParamTypeAgentTurnInput {
	return ParamTypeAgentTurnInput{
		Type: "agent_turn_input",
	}
}

// Parameter type for agent turn input.
//
// This struct has a constant value, construct it with
// [NewParamTypeAgentTurnInput].
type ParamTypeAgentTurnInput struct {
	// Discriminator type. Always "agent_turn_input"
	Type constant.AgentTurnInput `json:"type,required"`
	paramObj
}

func (r ParamTypeAgentTurnInput) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeAgentTurnInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParamTypeAgentTurnInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A scoring function resource for evaluating model outputs.
type ScoringFn struct {
	Identifier string                            `json:"identifier,required"`
	Metadata   map[string]ScoringFnMetadataUnion `json:"metadata,required"`
	ProviderID string                            `json:"provider_id,required"`
	// Parameter type for string values.
	ReturnType ParamTypeUnionResp `json:"return_type,required"`
	// The resource type, always scoring_function
	Type        constant.ScoringFunction `json:"type,required"`
	Description string                   `json:"description"`
	// Parameters for LLM-as-judge scoring function configuration.
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

// Parameters for LLM-as-judge scoring function configuration.
type ScoringFnParamsLlmAsJudgeScoringFnParamsResp struct {
	// Aggregation functions to apply to the scores of each row
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,required"`
	// Identifier of the LLM model to use as a judge for scoring
	JudgeModel string `json:"judge_model,required"`
	// Regexes to extract the answer from generated response
	JudgeScoreRegexes []string `json:"judge_score_regexes,required"`
	// The type of scoring function parameters, always llm_as_judge
	//
	// Any of "llm_as_judge", "regex_parser", "basic".
	Type ScoringFnParamsType `json:"type,required"`
	// (Optional) Custom prompt template for the judge model
	PromptTemplate string `json:"prompt_template"`
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

// Parameters for regex parser scoring function configuration.
type ScoringFnParamsRegexParserScoringFnParamsResp struct {
	// Aggregation functions to apply to the scores of each row
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,required"`
	// Regex to extract the answer from generated response
	ParsingRegexes []string `json:"parsing_regexes,required"`
	// The type of scoring function parameters, always regex_parser
	//
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

// Parameters for basic scoring function configuration.
type ScoringFnParamsBasicScoringFnParamsResp struct {
	// Aggregation functions to apply to the scores of each row
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,required"`
	// The type of scoring function parameters, always basic
	//
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

// Parameters for LLM-as-judge scoring function configuration.
//
// The properties AggregationFunctions, JudgeModel, JudgeScoreRegexes, Type are
// required.
type ScoringFnParamsLlmAsJudgeScoringFnParams struct {
	// Aggregation functions to apply to the scores of each row
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,omitzero,required"`
	// Identifier of the LLM model to use as a judge for scoring
	JudgeModel string `json:"judge_model,required"`
	// Regexes to extract the answer from generated response
	JudgeScoreRegexes []string `json:"judge_score_regexes,omitzero,required"`
	// The type of scoring function parameters, always llm_as_judge
	//
	// Any of "llm_as_judge", "regex_parser", "basic".
	Type ScoringFnParamsType `json:"type,omitzero,required"`
	// (Optional) Custom prompt template for the judge model
	PromptTemplate param.Opt[string] `json:"prompt_template,omitzero"`
	paramObj
}

func (r ScoringFnParamsLlmAsJudgeScoringFnParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFnParamsLlmAsJudgeScoringFnParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFnParamsLlmAsJudgeScoringFnParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for regex parser scoring function configuration.
//
// The properties AggregationFunctions, ParsingRegexes, Type are required.
type ScoringFnParamsRegexParserScoringFnParams struct {
	// Aggregation functions to apply to the scores of each row
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,omitzero,required"`
	// Regex to extract the answer from generated response
	ParsingRegexes []string `json:"parsing_regexes,omitzero,required"`
	// The type of scoring function parameters, always regex_parser
	//
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

// Parameters for basic scoring function configuration.
//
// The properties AggregationFunctions, Type are required.
type ScoringFnParamsBasicScoringFnParams struct {
	// Aggregation functions to apply to the scores of each row
	AggregationFunctions []AggregationFunctionType `json:"aggregation_functions,omitzero,required"`
	// The type of scoring function parameters, always basic
	//
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

// Types of scoring function parameter configurations.
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
	Description string `json:"description,required"`
	// The return type of the scoring function.
	ReturnType ParamTypeUnion `json:"return_type,omitzero,required"`
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

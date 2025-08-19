// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"encoding/json"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
	"github.com/varshaprasad96/llamastack-go-client/shared/constant"
)

// TypeService contains methods and other services that help with interacting with
// the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTypeService] method instead.
type TypeService struct {
	Options []option.RequestOption
}

// NewTypeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTypeService(opts ...option.RequestOption) (r TypeService) {
	r = TypeService{}
	r.Options = opts
	return
}

// Parameter type for agent turn input.
type AgentTurnInputType struct {
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
func (r AgentTurnInputType) RawJSON() string { return r.JSON.raw }
func (r *AgentTurnInputType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentTurnInputType to a AgentTurnInputTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentTurnInputTypeParam.Overrides()
func (r AgentTurnInputType) ToParam() AgentTurnInputTypeParam {
	return param.Override[AgentTurnInputTypeParam](json.RawMessage(r.RawJSON()))
}

func NewAgentTurnInputTypeParam() AgentTurnInputTypeParam {
	return AgentTurnInputTypeParam{
		Type: "agent_turn_input",
	}
}

// Parameter type for agent turn input.
//
// This struct has a constant value, construct it with
// [NewAgentTurnInputTypeParam].
type AgentTurnInputTypeParam struct {
	// Discriminator type. Always "agent_turn_input"
	Type constant.AgentTurnInput `json:"type,required"`
	paramObj
}

func (r AgentTurnInputTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnInputTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnInputTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for array values.
type ArrayType struct {
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
func (r ArrayType) RawJSON() string { return r.JSON.raw }
func (r *ArrayType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ArrayType to a ArrayTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ArrayTypeParam.Overrides()
func (r ArrayType) ToParam() ArrayTypeParam {
	return param.Override[ArrayTypeParam](json.RawMessage(r.RawJSON()))
}

func NewArrayTypeParam() ArrayTypeParam {
	return ArrayTypeParam{
		Type: "array",
	}
}

// Parameter type for array values.
//
// This struct has a constant value, construct it with [NewArrayTypeParam].
type ArrayTypeParam struct {
	// Discriminator type. Always "array"
	Type constant.Array `json:"type,required"`
	paramObj
}

func (r ArrayTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ArrayTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ArrayTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for boolean values.
type BooleanType struct {
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
func (r BooleanType) RawJSON() string { return r.JSON.raw }
func (r *BooleanType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BooleanType to a BooleanTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BooleanTypeParam.Overrides()
func (r BooleanType) ToParam() BooleanTypeParam {
	return param.Override[BooleanTypeParam](json.RawMessage(r.RawJSON()))
}

func NewBooleanTypeParam() BooleanTypeParam {
	return BooleanTypeParam{
		Type: "boolean",
	}
}

// Parameter type for boolean values.
//
// This struct has a constant value, construct it with [NewBooleanTypeParam].
type BooleanTypeParam struct {
	// Discriminator type. Always "boolean"
	Type constant.Boolean `json:"type,required"`
	paramObj
}

func (r BooleanTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow BooleanTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BooleanTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for chat completion input.
type ChatCompletionInputType struct {
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
func (r ChatCompletionInputType) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionInputType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChatCompletionInputType to a ChatCompletionInputTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChatCompletionInputTypeParam.Overrides()
func (r ChatCompletionInputType) ToParam() ChatCompletionInputTypeParam {
	return param.Override[ChatCompletionInputTypeParam](json.RawMessage(r.RawJSON()))
}

func NewChatCompletionInputTypeParam() ChatCompletionInputTypeParam {
	return ChatCompletionInputTypeParam{
		Type: "chat_completion_input",
	}
}

// Parameter type for chat completion input.
//
// This struct has a constant value, construct it with
// [NewChatCompletionInputTypeParam].
type ChatCompletionInputTypeParam struct {
	// Discriminator type. Always "chat_completion_input"
	Type constant.ChatCompletionInput `json:"type,required"`
	paramObj
}

func (r ChatCompletionInputTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionInputTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionInputTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for completion input.
type CompletionInputType struct {
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
func (r CompletionInputType) RawJSON() string { return r.JSON.raw }
func (r *CompletionInputType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CompletionInputType to a CompletionInputTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CompletionInputTypeParam.Overrides()
func (r CompletionInputType) ToParam() CompletionInputTypeParam {
	return param.Override[CompletionInputTypeParam](json.RawMessage(r.RawJSON()))
}

func NewCompletionInputTypeParam() CompletionInputTypeParam {
	return CompletionInputTypeParam{
		Type: "completion_input",
	}
}

// Parameter type for completion input.
//
// This struct has a constant value, construct it with
// [NewCompletionInputTypeParam].
type CompletionInputTypeParam struct {
	// Discriminator type. Always "completion_input"
	Type constant.CompletionInput `json:"type,required"`
	paramObj
}

func (r CompletionInputTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow CompletionInputTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletionInputTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from a completion request.
type CompletionResponse struct {
	// The generated completion text
	Content string `json:"content,required"`
	// Reason why generation stopped
	//
	// Any of "end_of_turn", "end_of_message", "out_of_tokens".
	StopReason CompletionResponseStopReason `json:"stop_reason,required"`
	// Optional log probabilities for generated tokens
	Logprobs []TokenLogProbs `json:"logprobs"`
	// (Optional) List of metrics associated with the API response
	Metrics []MetricInResponse `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		StopReason  respjson.Field
		Logprobs    respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *CompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reason why generation stopped
type CompletionResponseStopReason string

const (
	CompletionResponseStopReasonEndOfTurn    CompletionResponseStopReason = "end_of_turn"
	CompletionResponseStopReasonEndOfMessage CompletionResponseStopReason = "end_of_message"
	CompletionResponseStopReasonOutOfTokens  CompletionResponseStopReason = "out_of_tokens"
)

// Parameter type for JSON values.
type JsonType struct {
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
func (r JsonType) RawJSON() string { return r.JSON.raw }
func (r *JsonType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JsonType to a JsonTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JsonTypeParam.Overrides()
func (r JsonType) ToParam() JsonTypeParam {
	return param.Override[JsonTypeParam](json.RawMessage(r.RawJSON()))
}

func NewJsonTypeParam() JsonTypeParam {
	return JsonTypeParam{
		Type: "json",
	}
}

// Parameter type for JSON values.
//
// This struct has a constant value, construct it with [NewJsonTypeParam].
type JsonTypeParam struct {
	// Discriminator type. Always "json"
	Type constant.Json `json:"type,required"`
	paramObj
}

func (r JsonTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow JsonTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JsonTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for numeric values.
type NumberType struct {
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
func (r NumberType) RawJSON() string { return r.JSON.raw }
func (r *NumberType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NumberType to a NumberTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NumberTypeParam.Overrides()
func (r NumberType) ToParam() NumberTypeParam {
	return param.Override[NumberTypeParam](json.RawMessage(r.RawJSON()))
}

func NewNumberTypeParam() NumberTypeParam {
	return NumberTypeParam{
		Type: "number",
	}
}

// Parameter type for numeric values.
//
// This struct has a constant value, construct it with [NewNumberTypeParam].
type NumberTypeParam struct {
	// Discriminator type. Always "number"
	Type constant.Number `json:"type,required"`
	paramObj
}

func (r NumberTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow NumberTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for object values.
type ObjectType struct {
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
func (r ObjectType) RawJSON() string { return r.JSON.raw }
func (r *ObjectType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ObjectType to a ObjectTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ObjectTypeParam.Overrides()
func (r ObjectType) ToParam() ObjectTypeParam {
	return param.Override[ObjectTypeParam](json.RawMessage(r.RawJSON()))
}

func NewObjectTypeParam() ObjectTypeParam {
	return ObjectTypeParam{
		Type: "object",
	}
}

// Parameter type for object values.
//
// This struct has a constant value, construct it with [NewObjectTypeParam].
type ObjectTypeParam struct {
	// Discriminator type. Always "object"
	Type constant.Object `json:"type,required"`
	paramObj
}

func (r ObjectTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParamTypeUnionResp contains all possible properties and values from
// [ParamTypeStringTypeResp], [ParamTypeNumberTypeResp],
// [ParamTypeBooleanTypeResp], [ParamTypeArrayTypeResp], [ParamTypeObjectTypeResp],
// [ParamTypeJsonTypeResp], [ParamTypeUnionTypeResp],
// [ParamTypeChatCompletionInputTypeResp], [ParamTypeCompletionInputTypeResp],
// [ParamTypeAgentTurnInputTypeResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ParamTypeUnionResp struct {
	Type string `json:"type"`
	JSON struct {
		Type respjson.Field
		raw  string
	} `json:"-"`
}

func (u ParamTypeUnionResp) AsStringType() (v ParamTypeStringTypeResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsNumberType() (v ParamTypeNumberTypeResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsBooleanType() (v ParamTypeBooleanTypeResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsArrayType() (v ParamTypeArrayTypeResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsObjectType() (v ParamTypeObjectTypeResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsJsonType() (v ParamTypeJsonTypeResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsUnionType() (v ParamTypeUnionTypeResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsChatCompletionInputType() (v ParamTypeChatCompletionInputTypeResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsCompletionInputType() (v ParamTypeCompletionInputTypeResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParamTypeUnionResp) AsAgentTurnInputType() (v ParamTypeAgentTurnInputTypeResp) {
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
type ParamTypeStringTypeResp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ParamTypeBaseResp
	StringType
}

// Returns the unmodified JSON received from the API
func (r ParamTypeStringTypeResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeStringTypeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for numeric values.
type ParamTypeNumberTypeResp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ParamTypeBaseResp
	NumberType
}

// Returns the unmodified JSON received from the API
func (r ParamTypeNumberTypeResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeNumberTypeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for boolean values.
type ParamTypeBooleanTypeResp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ParamTypeBaseResp
	BooleanType
}

// Returns the unmodified JSON received from the API
func (r ParamTypeBooleanTypeResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeBooleanTypeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for array values.
type ParamTypeArrayTypeResp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ParamTypeBaseResp
	ArrayType
}

// Returns the unmodified JSON received from the API
func (r ParamTypeArrayTypeResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeArrayTypeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for object values.
type ParamTypeObjectTypeResp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ParamTypeBaseResp
	ObjectType
}

// Returns the unmodified JSON received from the API
func (r ParamTypeObjectTypeResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeObjectTypeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for JSON values.
type ParamTypeJsonTypeResp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ParamTypeBaseResp
	JsonType
}

// Returns the unmodified JSON received from the API
func (r ParamTypeJsonTypeResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeJsonTypeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for union values.
type ParamTypeUnionTypeResp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ParamTypeBaseResp
	UnionType
}

// Returns the unmodified JSON received from the API
func (r ParamTypeUnionTypeResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeUnionTypeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for chat completion input.
type ParamTypeChatCompletionInputTypeResp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ParamTypeBaseResp
	ChatCompletionInputType
}

// Returns the unmodified JSON received from the API
func (r ParamTypeChatCompletionInputTypeResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeChatCompletionInputTypeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for completion input.
type ParamTypeCompletionInputTypeResp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ParamTypeBaseResp
	CompletionInputType
}

// Returns the unmodified JSON received from the API
func (r ParamTypeCompletionInputTypeResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeCompletionInputTypeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for agent turn input.
type ParamTypeAgentTurnInputTypeResp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ParamTypeBaseResp
	AgentTurnInputType
}

// Returns the unmodified JSON received from the API
func (r ParamTypeAgentTurnInputTypeResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeAgentTurnInputTypeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ParamTypeOfStringType(type_ ParamTypeBaseType) ParamTypeUnion {
	var variant ParamTypeStringType
	variant.Type = type_
	return ParamTypeUnion{OfStringType: &variant}
}

func ParamTypeOfNumberType(type_ ParamTypeBaseType) ParamTypeUnion {
	var variant ParamTypeNumberType
	variant.Type = type_
	return ParamTypeUnion{OfNumberType: &variant}
}

func ParamTypeOfBooleanType(type_ ParamTypeBaseType) ParamTypeUnion {
	var variant ParamTypeBooleanType
	variant.Type = type_
	return ParamTypeUnion{OfBooleanType: &variant}
}

func ParamTypeOfArrayType(type_ ParamTypeBaseType) ParamTypeUnion {
	var variant ParamTypeArrayType
	variant.Type = type_
	return ParamTypeUnion{OfArrayType: &variant}
}

func ParamTypeOfObjectType(type_ ParamTypeBaseType) ParamTypeUnion {
	var variant ParamTypeObjectType
	variant.Type = type_
	return ParamTypeUnion{OfObjectType: &variant}
}

func ParamTypeOfJsonType(type_ ParamTypeBaseType) ParamTypeUnion {
	var variant ParamTypeJsonType
	variant.Type = type_
	return ParamTypeUnion{OfJsonType: &variant}
}

func ParamTypeOfUnionType(type_ ParamTypeBaseType) ParamTypeUnion {
	var variant ParamTypeUnionType
	variant.Type = type_
	return ParamTypeUnion{OfUnionType: &variant}
}

func ParamTypeOfChatCompletionInputType(type_ ParamTypeBaseType) ParamTypeUnion {
	var variant ParamTypeChatCompletionInputType
	variant.Type = type_
	return ParamTypeUnion{OfChatCompletionInputType: &variant}
}

func ParamTypeOfCompletionInputType(type_ ParamTypeBaseType) ParamTypeUnion {
	var variant ParamTypeCompletionInputType
	variant.Type = type_
	return ParamTypeUnion{OfCompletionInputType: &variant}
}

func ParamTypeOfAgentTurnInputType(type_ ParamTypeBaseType) ParamTypeUnion {
	var variant ParamTypeAgentTurnInputType
	variant.Type = type_
	return ParamTypeUnion{OfAgentTurnInputType: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParamTypeUnion struct {
	OfStringType              *ParamTypeStringType              `json:",omitzero,inline"`
	OfNumberType              *ParamTypeNumberType              `json:",omitzero,inline"`
	OfBooleanType             *ParamTypeBooleanType             `json:",omitzero,inline"`
	OfArrayType               *ParamTypeArrayType               `json:",omitzero,inline"`
	OfObjectType              *ParamTypeObjectType              `json:",omitzero,inline"`
	OfJsonType                *ParamTypeJsonType                `json:",omitzero,inline"`
	OfUnionType               *ParamTypeUnionType               `json:",omitzero,inline"`
	OfChatCompletionInputType *ParamTypeChatCompletionInputType `json:",omitzero,inline"`
	OfCompletionInputType     *ParamTypeCompletionInputType     `json:",omitzero,inline"`
	OfAgentTurnInputType      *ParamTypeAgentTurnInputType      `json:",omitzero,inline"`
	paramUnion
}

func (u ParamTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringType,
		u.OfNumberType,
		u.OfBooleanType,
		u.OfArrayType,
		u.OfObjectType,
		u.OfJsonType,
		u.OfUnionType,
		u.OfChatCompletionInputType,
		u.OfCompletionInputType,
		u.OfAgentTurnInputType)
}
func (u *ParamTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ParamTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfStringType) {
		return u.OfStringType
	} else if !param.IsOmitted(u.OfNumberType) {
		return u.OfNumberType
	} else if !param.IsOmitted(u.OfBooleanType) {
		return u.OfBooleanType
	} else if !param.IsOmitted(u.OfArrayType) {
		return u.OfArrayType
	} else if !param.IsOmitted(u.OfObjectType) {
		return u.OfObjectType
	} else if !param.IsOmitted(u.OfJsonType) {
		return u.OfJsonType
	} else if !param.IsOmitted(u.OfUnionType) {
		return u.OfUnionType
	} else if !param.IsOmitted(u.OfChatCompletionInputType) {
		return u.OfChatCompletionInputType
	} else if !param.IsOmitted(u.OfCompletionInputType) {
		return u.OfCompletionInputType
	} else if !param.IsOmitted(u.OfAgentTurnInputType) {
		return u.OfAgentTurnInputType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ParamTypeUnion) GetType() *string {
	if vt := u.OfStringType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStringType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNumberType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNumberType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBooleanType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBooleanType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfArrayType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfArrayType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfObjectType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfObjectType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJsonType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJsonType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfUnionType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfUnionType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfChatCompletionInputType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfChatCompletionInputType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCompletionInputType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCompletionInputType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAgentTurnInputType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAgentTurnInputType; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Parameter type for string values.
type ParamTypeStringType struct {
	ParamTypeBase
	StringTypeParam
}

func (r ParamTypeStringType) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeStringType
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameter type for numeric values.
type ParamTypeNumberType struct {
	ParamTypeBase
	NumberTypeParam
}

func (r ParamTypeNumberType) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeNumberType
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameter type for boolean values.
type ParamTypeBooleanType struct {
	ParamTypeBase
	BooleanTypeParam
}

func (r ParamTypeBooleanType) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeBooleanType
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameter type for array values.
type ParamTypeArrayType struct {
	ParamTypeBase
	ArrayTypeParam
}

func (r ParamTypeArrayType) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeArrayType
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameter type for object values.
type ParamTypeObjectType struct {
	ParamTypeBase
	ObjectTypeParam
}

func (r ParamTypeObjectType) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeObjectType
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameter type for JSON values.
type ParamTypeJsonType struct {
	ParamTypeBase
	JsonTypeParam
}

func (r ParamTypeJsonType) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeJsonType
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameter type for union values.
type ParamTypeUnionType struct {
	ParamTypeBase
	UnionTypeParam
}

func (r ParamTypeUnionType) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeUnionType
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameter type for chat completion input.
type ParamTypeChatCompletionInputType struct {
	ParamTypeBase
	ChatCompletionInputTypeParam
}

func (r ParamTypeChatCompletionInputType) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeChatCompletionInputType
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameter type for completion input.
type ParamTypeCompletionInputType struct {
	ParamTypeBase
	CompletionInputTypeParam
}

func (r ParamTypeCompletionInputType) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeCompletionInputType
	return param.MarshalObject(r, (*shadow)(&r))
}

// Parameter type for agent turn input.
type ParamTypeAgentTurnInputType struct {
	ParamTypeBase
	AgentTurnInputTypeParam
}

func (r ParamTypeAgentTurnInputType) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeAgentTurnInputType
	return param.MarshalObject(r, (*shadow)(&r))
}

type ParamTypeBaseResp struct {
	// Any of "string", "number", "boolean", "array", "object", "json", "union",
	// "chat_completion_input", "completion_input", "agent_turn_input".
	Type ParamTypeBaseType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParamTypeBaseResp) RawJSON() string { return r.JSON.raw }
func (r *ParamTypeBaseResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ParamTypeBaseResp to a ParamTypeBase.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ParamTypeBase.Overrides()
func (r ParamTypeBaseResp) ToParam() ParamTypeBase {
	return param.Override[ParamTypeBase](json.RawMessage(r.RawJSON()))
}

type ParamTypeBaseType string

const (
	ParamTypeBaseTypeString              ParamTypeBaseType = "string"
	ParamTypeBaseTypeNumber              ParamTypeBaseType = "number"
	ParamTypeBaseTypeBoolean             ParamTypeBaseType = "boolean"
	ParamTypeBaseTypeArray               ParamTypeBaseType = "array"
	ParamTypeBaseTypeObject              ParamTypeBaseType = "object"
	ParamTypeBaseTypeJson                ParamTypeBaseType = "json"
	ParamTypeBaseTypeUnion               ParamTypeBaseType = "union"
	ParamTypeBaseTypeChatCompletionInput ParamTypeBaseType = "chat_completion_input"
	ParamTypeBaseTypeCompletionInput     ParamTypeBaseType = "completion_input"
	ParamTypeBaseTypeAgentTurnInput      ParamTypeBaseType = "agent_turn_input"
)

// The property Type is required.
type ParamTypeBase struct {
	// Any of "string", "number", "boolean", "array", "object", "json", "union",
	// "chat_completion_input", "completion_input", "agent_turn_input".
	Type ParamTypeBaseType `json:"type,omitzero,required"`
	paramObj
}

func (r ParamTypeBase) MarshalJSON() (data []byte, err error) {
	type shadow ParamTypeBase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParamTypeBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter type for string values.
type StringType struct {
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
func (r StringType) RawJSON() string { return r.JSON.raw }
func (r *StringType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StringType to a StringTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StringTypeParam.Overrides()
func (r StringType) ToParam() StringTypeParam {
	return param.Override[StringTypeParam](json.RawMessage(r.RawJSON()))
}

func NewStringTypeParam() StringTypeParam {
	return StringTypeParam{
		Type: "string",
	}
}

// Parameter type for string values.
//
// This struct has a constant value, construct it with [NewStringTypeParam].
type StringTypeParam struct {
	// Discriminator type. Always "string"
	Type constant.String `json:"type,required"`
	paramObj
}

func (r StringTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow StringTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StringTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Log probabilities for generated tokens.
type TokenLogProbs struct {
	// Dictionary mapping tokens to their log probabilities
	LogprobsByToken map[string]float64 `json:"logprobs_by_token,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LogprobsByToken respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenLogProbs) RawJSON() string { return r.JSON.raw }
func (r *TokenLogProbs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition used in runtime contexts.
type ToolDef struct {
	// Name of the tool
	Name string `json:"name,required"`
	// (Optional) Human-readable description of what the tool does
	Description string `json:"description"`
	// (Optional) Additional metadata about the tool
	Metadata map[string]ToolDefMetadataUnion `json:"metadata"`
	// (Optional) List of parameters this tool accepts
	Parameters []ToolParameterResp `json:"parameters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Metadata    respjson.Field
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDef) RawJSON() string { return r.JSON.raw }
func (r *ToolDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolDef to a ToolDefParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolDefParam.Overrides()
func (r ToolDef) ToParam() ToolDefParam {
	return param.Override[ToolDefParam](json.RawMessage(r.RawJSON()))
}

// ToolDefMetadataUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolDefMetadataUnion struct {
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

func (u ToolDefMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolDefMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolDefMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolDefMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool definition used in runtime contexts.
//
// The property Name is required.
type ToolDefParam struct {
	// Name of the tool
	Name string `json:"name,required"`
	// (Optional) Human-readable description of what the tool does
	Description param.Opt[string] `json:"description,omitzero"`
	// (Optional) Additional metadata about the tool
	Metadata map[string]ToolDefMetadataUnionParam `json:"metadata,omitzero"`
	// (Optional) List of parameters this tool accepts
	Parameters []ToolParameter `json:"parameters,omitzero"`
	paramObj
}

func (r ToolDefParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolDefParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolDefParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolDefMetadataUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ToolDefMetadataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ToolDefMetadataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolDefMetadataUnionParam) asAny() any {
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

// Parameter type for union values.
type UnionType struct {
	// Discriminator type. Always "union"
	Type constant.Union `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnionType) RawJSON() string { return r.JSON.raw }
func (r *UnionType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UnionType to a UnionTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UnionTypeParam.Overrides()
func (r UnionType) ToParam() UnionTypeParam {
	return param.Override[UnionTypeParam](json.RawMessage(r.RawJSON()))
}

func NewUnionTypeParam() UnionTypeParam {
	return UnionTypeParam{
		Type: "union",
	}
}

// Parameter type for union values.
//
// This struct has a constant value, construct it with [NewUnionTypeParam].
type UnionTypeParam struct {
	// Discriminator type. Always "union"
	Type constant.Union `json:"type,required"`
	paramObj
}

func (r UnionTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow UnionTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnionTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

// SharedService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSharedService] method instead.
type SharedService struct {
	Options []option.RequestOption
}

// NewSharedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSharedService(opts ...option.RequestOption) (r SharedService) {
	r = SharedService{}
	r.Options = opts
	return
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
//	case llamastackclient.ParamTypeStringResp:
//	case llamastackclient.ParamTypeNumberResp:
//	case llamastackclient.ParamTypeBooleanResp:
//	case llamastackclient.ParamTypeArrayResp:
//	case llamastackclient.ParamTypeObjectResp:
//	case llamastackclient.ParamTypeJsonResp:
//	case llamastackclient.ParamTypeUnionResp:
//	case llamastackclient.ParamTypeChatCompletionInputResp:
//	case llamastackclient.ParamTypeCompletionInputResp:
//	case llamastackclient.ParamTypeAgentTurnInputResp:
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/apiquery"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
	"github.com/varshaprasad96/llamastack-go-client/shared/constant"
)

// OpenAIV1ChatCompletionService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIV1ChatCompletionService] method instead.
type OpenAIV1ChatCompletionService struct {
	Options []option.RequestOption
}

// NewOpenAIV1ChatCompletionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOpenAIV1ChatCompletionService(opts ...option.RequestOption) (r OpenAIV1ChatCompletionService) {
	r = OpenAIV1ChatCompletionService{}
	r.Options = opts
	return
}

// Generate an OpenAI-compatible chat completion for the given messages using the
// specified model.
func (r *OpenAIV1ChatCompletionService) New(ctx context.Context, body OpenAIV1ChatCompletionNewParams, opts ...option.RequestOption) (res *OpenAiv1ChatCompletionNewResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Describe a chat completion by its ID.
func (r *OpenAIV1ChatCompletionService) Get(ctx context.Context, completionID string, opts ...option.RequestOption) (res *OpenAiv1ChatCompletionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if completionID == "" {
		err = errors.New("missing required completion_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/chat/completions/%s", completionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all chat completions.
func (r *OpenAIV1ChatCompletionService) List(ctx context.Context, query OpenAIV1ChatCompletionListParams, opts ...option.RequestOption) (res *OpenAiv1ChatCompletionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Tool call specification for OpenAI-compatible chat completion responses.
type ChatCompletionToolCall struct {
	// Must be "function" to identify this as a function call
	Type constant.Function `json:"type,required"`
	// (Optional) Unique identifier for the tool call
	ID string `json:"id"`
	// (Optional) Function call details
	Function ChatCompletionToolCallFunction `json:"function"`
	// (Optional) Index of the tool call in the list
	Index int64 `json:"index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChatCompletionToolCall to a ChatCompletionToolCallParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChatCompletionToolCallParam.Overrides()
func (r ChatCompletionToolCall) ToParam() ChatCompletionToolCallParam {
	return param.Override[ChatCompletionToolCallParam](json.RawMessage(r.RawJSON()))
}

// (Optional) Function call details
type ChatCompletionToolCallFunction struct {
	// (Optional) Arguments to pass to the function as a JSON string
	Arguments string `json:"arguments"`
	// (Optional) Name of the function to call
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
//
// The property Type is required.
type ChatCompletionToolCallParam struct {
	// (Optional) Unique identifier for the tool call
	ID param.Opt[string] `json:"id,omitzero"`
	// (Optional) Index of the tool call in the list
	Index param.Opt[int64] `json:"index,omitzero"`
	// (Optional) Function call details
	Function ChatCompletionToolCallFunctionParam `json:"function,omitzero"`
	// Must be "function" to identify this as a function call
	//
	// This field can be elided, and will marshal its zero value as "function".
	Type constant.Function `json:"type,required"`
	paramObj
}

func (r ChatCompletionToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Function call details
type ChatCompletionToolCallFunctionParam struct {
	// (Optional) Arguments to pass to the function as a JSON string
	Arguments param.Opt[string] `json:"arguments,omitzero"`
	// (Optional) Name of the function to call
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ChatCompletionToolCallFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionToolCallFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionToolCallFunctionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A choice from an OpenAI-compatible chat completion response.
type Choice struct {
	// The reason the model stopped generating
	FinishReason string `json:"finish_reason,required"`
	// The index of the choice
	Index int64 `json:"index,required"`
	// The message from the model
	Message ChoiceMessageUnion `json:"message,required"`
	// (Optional) The log probabilities for the tokens in the message
	Logprobs ChoiceLogprobs `json:"logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Index        respjson.Field
		Message      respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Choice) RawJSON() string { return r.JSON.raw }
func (r *Choice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChoiceMessageUnion contains all possible properties and values from
// [OpenAIUserMessageParamResp], [OpenAISystemMessageParamResp],
// [OpenAIAssistantMessageParamResp], [OpenAIToolMessageParamResp],
// [OpenAIDeveloperMessageParamResp].
//
// Use the [ChoiceMessageUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChoiceMessageUnion struct {
	// This field is a union of [OpenAIUserMessageParamContentUnionResp],
	// [OpenAISystemMessageParamContentUnionResp],
	// [OpenAIAssistantMessageParamContentUnionResp],
	// [OpenAIToolMessageParamContentUnionResp],
	// [OpenAIDeveloperMessageParamContentUnionResp]
	Content ChoiceMessageUnionContent `json:"content"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant [OpenAIAssistantMessageParamResp].
	ToolCalls []ChatCompletionToolCall `json:"tool_calls"`
	// This field is from variant [OpenAIToolMessageParamResp].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Role       respjson.Field
		Name       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyChoiceMessage is implemented by each variant of [ChoiceMessageUnion] to add
// type safety for the return type of [ChoiceMessageUnion.AsAny]
type anyChoiceMessage interface {
	implChoiceMessageUnion()
}

func (OpenAIUserMessageParamResp) implChoiceMessageUnion()      {}
func (OpenAISystemMessageParamResp) implChoiceMessageUnion()    {}
func (OpenAIAssistantMessageParamResp) implChoiceMessageUnion() {}
func (OpenAIToolMessageParamResp) implChoiceMessageUnion()      {}
func (OpenAIDeveloperMessageParamResp) implChoiceMessageUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChoiceMessageUnion.AsAny().(type) {
//	case llamastackclient.OpenAIUserMessageParamResp:
//	case llamastackclient.OpenAISystemMessageParamResp:
//	case llamastackclient.OpenAIAssistantMessageParamResp:
//	case llamastackclient.OpenAIToolMessageParamResp:
//	case llamastackclient.OpenAIDeveloperMessageParamResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChoiceMessageUnion) AsAny() anyChoiceMessage {
	switch u.Role {
	case "user":
		return u.AsUser()
	case "system":
		return u.AsSystem()
	case "assistant":
		return u.AsAssistant()
	case "tool":
		return u.AsTool()
	case "developer":
		return u.AsDeveloper()
	}
	return nil
}

func (u ChoiceMessageUnion) AsUser() (v OpenAIUserMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChoiceMessageUnion) AsSystem() (v OpenAISystemMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChoiceMessageUnion) AsAssistant() (v OpenAIAssistantMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChoiceMessageUnion) AsTool() (v OpenAIToolMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChoiceMessageUnion) AsDeveloper() (v OpenAIDeveloperMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChoiceMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *ChoiceMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChoiceMessageUnionContent is an implicit subunion of [ChoiceMessageUnion].
// ChoiceMessageUnionContent provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [ChoiceMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOpenAIUserMessageContentArray OfContentPartTextArray]
type ChoiceMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIUserMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIUserMessageContentArray []OpenAIUserMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a [[]ContentPartTextParamResp]
	// instead of an object.
	OfContentPartTextArray []ContentPartTextParamResp `json:",inline"`
	JSON                   struct {
		OfString                        respjson.Field
		OfOpenAIUserMessageContentArray respjson.Field
		OfContentPartTextArray          respjson.Field
		raw                             string
	} `json:"-"`
}

func (r *ChoiceMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probabilities for the tokens in the message from an OpenAI-compatible
// chat completion response.
type ChoiceLogprobs struct {
	// (Optional) The log probabilities for the tokens in the message
	Content []TokenLogProb `json:"content"`
	// (Optional) The log probabilities for the tokens in the message
	Refusal []TokenLogProb `json:"refusal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *ChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part for OpenAI-compatible chat completion messages.
type ContentPartTextParamResp struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentPartTextParamResp) RawJSON() string { return r.JSON.raw }
func (r *ContentPartTextParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContentPartTextParamResp to a ContentPartTextParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContentPartTextParam.Overrides()
func (r ContentPartTextParamResp) ToParam() ContentPartTextParam {
	return param.Override[ContentPartTextParam](json.RawMessage(r.RawJSON()))
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type ContentPartTextParam struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ContentPartTextParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentPartTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentPartTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type OpenAIAssistantMessageParamResp struct {
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// The content of the model's response
	Content OpenAIAssistantMessageParamContentUnionResp `json:"content"`
	// (Optional) The name of the assistant message participant.
	Name string `json:"name"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Content     respjson.Field
		Name        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIAssistantMessageParamResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIAssistantMessageParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OpenAIAssistantMessageParamResp to a
// OpenAIAssistantMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OpenAIAssistantMessageParam.Overrides()
func (r OpenAIAssistantMessageParamResp) ToParam() OpenAIAssistantMessageParam {
	return param.Override[OpenAIAssistantMessageParam](json.RawMessage(r.RawJSON()))
}

// OpenAIAssistantMessageParamContentUnionResp contains all possible properties and
// values from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type OpenAIAssistantMessageParamContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]ContentPartTextParamResp]
	// instead of an object.
	OfContentPartTextArray []ContentPartTextParamResp `json:",inline"`
	JSON                   struct {
		OfString               respjson.Field
		OfContentPartTextArray respjson.Field
		raw                    string
	} `json:"-"`
}

func (u OpenAIAssistantMessageParamContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIAssistantMessageParamContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIAssistantMessageParamContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIAssistantMessageParamContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
//
// The property Role is required.
type OpenAIAssistantMessageParam struct {
	// (Optional) The name of the assistant message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// The content of the model's response
	Content OpenAIAssistantMessageParamContentUnion `json:"content,omitzero"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionToolCallParam `json:"tool_calls,omitzero"`
	// Must be "assistant" to identify this as the model's response
	//
	// This field can be elided, and will marshal its zero value as "assistant".
	Role constant.Assistant `json:"role,required"`
	paramObj
}

func (r OpenAIAssistantMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIAssistantMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIAssistantMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIAssistantMessageParamContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIAssistantMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *OpenAIAssistantMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIAssistantMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfContentPartTextArray) {
		return &u.OfContentPartTextArray
	}
	return nil
}

// A message from the developer in an OpenAI-compatible chat completion request.
type OpenAIDeveloperMessageParamResp struct {
	// The content of the developer message
	Content OpenAIDeveloperMessageParamContentUnionResp `json:"content,required"`
	// Must be "developer" to identify this as a developer message
	Role constant.Developer `json:"role,required"`
	// (Optional) The name of the developer message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIDeveloperMessageParamResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIDeveloperMessageParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OpenAIDeveloperMessageParamResp to a
// OpenAIDeveloperMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OpenAIDeveloperMessageParam.Overrides()
func (r OpenAIDeveloperMessageParamResp) ToParam() OpenAIDeveloperMessageParam {
	return param.Override[OpenAIDeveloperMessageParam](json.RawMessage(r.RawJSON()))
}

// OpenAIDeveloperMessageParamContentUnionResp contains all possible properties and
// values from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type OpenAIDeveloperMessageParamContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]ContentPartTextParamResp]
	// instead of an object.
	OfContentPartTextArray []ContentPartTextParamResp `json:",inline"`
	JSON                   struct {
		OfString               respjson.Field
		OfContentPartTextArray respjson.Field
		raw                    string
	} `json:"-"`
}

func (u OpenAIDeveloperMessageParamContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIDeveloperMessageParamContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIDeveloperMessageParamContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIDeveloperMessageParamContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
//
// The properties Content, Role are required.
type OpenAIDeveloperMessageParam struct {
	// The content of the developer message
	Content OpenAIDeveloperMessageParamContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the developer message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "developer" to identify this as a developer message
	//
	// This field can be elided, and will marshal its zero value as "developer".
	Role constant.Developer `json:"role,required"`
	paramObj
}

func (r OpenAIDeveloperMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIDeveloperMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIDeveloperMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIDeveloperMessageParamContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIDeveloperMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *OpenAIDeveloperMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIDeveloperMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfContentPartTextArray) {
		return &u.OfContentPartTextArray
	}
	return nil
}

// A system message providing instructions or context to the model.
type OpenAISystemMessageParamResp struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content OpenAISystemMessageParamContentUnionResp `json:"content,required"`
	// Must be "system" to identify this as a system message
	Role constant.System `json:"role,required"`
	// (Optional) The name of the system message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAISystemMessageParamResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAISystemMessageParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OpenAISystemMessageParamResp to a
// OpenAISystemMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OpenAISystemMessageParam.Overrides()
func (r OpenAISystemMessageParamResp) ToParam() OpenAISystemMessageParam {
	return param.Override[OpenAISystemMessageParam](json.RawMessage(r.RawJSON()))
}

// OpenAISystemMessageParamContentUnionResp contains all possible properties and
// values from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type OpenAISystemMessageParamContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]ContentPartTextParamResp]
	// instead of an object.
	OfContentPartTextArray []ContentPartTextParamResp `json:",inline"`
	JSON                   struct {
		OfString               respjson.Field
		OfContentPartTextArray respjson.Field
		raw                    string
	} `json:"-"`
}

func (u OpenAISystemMessageParamContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAISystemMessageParamContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAISystemMessageParamContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAISystemMessageParamContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
//
// The properties Content, Role are required.
type OpenAISystemMessageParam struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content OpenAISystemMessageParamContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the system message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "system" to identify this as a system message
	//
	// This field can be elided, and will marshal its zero value as "system".
	Role constant.System `json:"role,required"`
	paramObj
}

func (r OpenAISystemMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAISystemMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAISystemMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAISystemMessageParamContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAISystemMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *OpenAISystemMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAISystemMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfContentPartTextArray) {
		return &u.OfContentPartTextArray
	}
	return nil
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type OpenAIToolMessageParamResp struct {
	// The response content from the tool
	Content OpenAIToolMessageParamContentUnionResp `json:"content,required"`
	// Must be "tool" to identify this as a tool response
	Role constant.Tool `json:"role,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ToolCallID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIToolMessageParamResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIToolMessageParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OpenAIToolMessageParamResp to a OpenAIToolMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OpenAIToolMessageParam.Overrides()
func (r OpenAIToolMessageParamResp) ToParam() OpenAIToolMessageParam {
	return param.Override[OpenAIToolMessageParam](json.RawMessage(r.RawJSON()))
}

// OpenAIToolMessageParamContentUnionResp contains all possible properties and
// values from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type OpenAIToolMessageParamContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]ContentPartTextParamResp]
	// instead of an object.
	OfContentPartTextArray []ContentPartTextParamResp `json:",inline"`
	JSON                   struct {
		OfString               respjson.Field
		OfContentPartTextArray respjson.Field
		raw                    string
	} `json:"-"`
}

func (u OpenAIToolMessageParamContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIToolMessageParamContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIToolMessageParamContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIToolMessageParamContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
//
// The properties Content, Role, ToolCallID are required.
type OpenAIToolMessageParam struct {
	// The response content from the tool
	Content OpenAIToolMessageParamContentUnion `json:"content,omitzero,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// Must be "tool" to identify this as a tool response
	//
	// This field can be elided, and will marshal its zero value as "tool".
	Role constant.Tool `json:"role,required"`
	paramObj
}

func (r OpenAIToolMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIToolMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIToolMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIToolMessageParamContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIToolMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *OpenAIToolMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIToolMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfContentPartTextArray) {
		return &u.OfContentPartTextArray
	}
	return nil
}

// A message from the user in an OpenAI-compatible chat completion request.
type OpenAIUserMessageParamResp struct {
	// The content of the message, which can include text and other media
	Content OpenAIUserMessageParamContentUnionResp `json:"content,required"`
	// Must be "user" to identify this as a user message
	Role constant.User `json:"role,required"`
	// (Optional) The name of the user message participant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIUserMessageParamResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIUserMessageParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OpenAIUserMessageParamResp to a OpenAIUserMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OpenAIUserMessageParam.Overrides()
func (r OpenAIUserMessageParamResp) ToParam() OpenAIUserMessageParam {
	return param.Override[OpenAIUserMessageParam](json.RawMessage(r.RawJSON()))
}

// OpenAIUserMessageParamContentUnionResp contains all possible properties and
// values from [string], [[]OpenAIUserMessageParamContentArrayItemUnionResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOpenAIUserMessageContentArray]
type OpenAIUserMessageParamContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIUserMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIUserMessageContentArray []OpenAIUserMessageParamContentArrayItemUnionResp `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfOpenAIUserMessageContentArray respjson.Field
		raw                             string
	} `json:"-"`
}

func (u OpenAIUserMessageParamContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIUserMessageParamContentUnionResp) AsOpenAIUserMessageParamContentArrayRespArray() (v []OpenAIUserMessageParamContentArrayItemUnionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIUserMessageParamContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIUserMessageParamContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIUserMessageParamContentArrayItemUnionResp contains all possible properties
// and values from [ContentPartTextParamResp],
// [OpenAIUserMessageParamContentArrayItemImageURLResp],
// [OpenAIUserMessageParamContentArrayItemFileResp].
//
// Use the [OpenAIUserMessageParamContentArrayItemUnionResp.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAIUserMessageParamContentArrayItemUnionResp struct {
	// This field is from variant [ContentPartTextParamResp].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant [OpenAIUserMessageParamContentArrayItemImageURLResp].
	ImageURL OpenAIUserMessageParamContentArrayItemImageURLImageURLResp `json:"image_url"`
	// This field is from variant [OpenAIUserMessageParamContentArrayItemFileResp].
	File OpenAIUserMessageParamContentArrayItemFileFileResp `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyOpenAIUserMessageParamContentArrayItemResp is implemented by each variant of
// [OpenAIUserMessageParamContentArrayItemUnionResp] to add type safety for the
// return type of [OpenAIUserMessageParamContentArrayItemUnionResp.AsAny]
type anyOpenAIUserMessageParamContentArrayItemResp interface {
	implOpenAIUserMessageParamContentArrayItemUnionResp()
}

func (ContentPartTextParamResp) implOpenAIUserMessageParamContentArrayItemUnionResp() {}
func (OpenAIUserMessageParamContentArrayItemImageURLResp) implOpenAIUserMessageParamContentArrayItemUnionResp() {
}
func (OpenAIUserMessageParamContentArrayItemFileResp) implOpenAIUserMessageParamContentArrayItemUnionResp() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := OpenAIUserMessageParamContentArrayItemUnionResp.AsAny().(type) {
//	case llamastackclient.ContentPartTextParamResp:
//	case llamastackclient.OpenAIUserMessageParamContentArrayItemImageURLResp:
//	case llamastackclient.OpenAIUserMessageParamContentArrayItemFileResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u OpenAIUserMessageParamContentArrayItemUnionResp) AsAny() anyOpenAIUserMessageParamContentArrayItemResp {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image_url":
		return u.AsImageURL()
	case "file":
		return u.AsFile()
	}
	return nil
}

func (u OpenAIUserMessageParamContentArrayItemUnionResp) AsText() (v ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIUserMessageParamContentArrayItemUnionResp) AsImageURL() (v OpenAIUserMessageParamContentArrayItemImageURLResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIUserMessageParamContentArrayItemUnionResp) AsFile() (v OpenAIUserMessageParamContentArrayItemFileResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIUserMessageParamContentArrayItemUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIUserMessageParamContentArrayItemUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type OpenAIUserMessageParamContentArrayItemImageURLResp struct {
	// Image URL specification and processing details
	ImageURL OpenAIUserMessageParamContentArrayItemImageURLImageURLResp `json:"image_url,required"`
	// Must be "image_url" to identify this as image content
	Type constant.ImageURL `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIUserMessageParamContentArrayItemImageURLResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIUserMessageParamContentArrayItemImageURLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
type OpenAIUserMessageParamContentArrayItemImageURLImageURLResp struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail string `json:"detail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIUserMessageParamContentArrayItemImageURLImageURLResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAIUserMessageParamContentArrayItemImageURLImageURLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIUserMessageParamContentArrayItemFileResp struct {
	File OpenAIUserMessageParamContentArrayItemFileFileResp `json:"file,required"`
	Type constant.File                                      `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIUserMessageParamContentArrayItemFileResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIUserMessageParamContentArrayItemFileResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIUserMessageParamContentArrayItemFileFileResp struct {
	FileData string `json:"file_data"`
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileData    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIUserMessageParamContentArrayItemFileFileResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIUserMessageParamContentArrayItemFileFileResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
//
// The properties Content, Role are required.
type OpenAIUserMessageParam struct {
	// The content of the message, which can include text and other media
	Content OpenAIUserMessageParamContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the user message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "user" to identify this as a user message
	//
	// This field can be elided, and will marshal its zero value as "user".
	Role constant.User `json:"role,required"`
	paramObj
}

func (r OpenAIUserMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIUserMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIUserMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIUserMessageParamContentUnion struct {
	OfString                        param.Opt[string]                             `json:",omitzero,inline"`
	OfOpenAIUserMessageContentArray []OpenAIUserMessageParamContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIUserMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIUserMessageContentArray)
}
func (u *OpenAIUserMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIUserMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAIUserMessageContentArray) {
		return &u.OfOpenAIUserMessageContentArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIUserMessageParamContentArrayItemUnion struct {
	OfText     *ContentPartTextParam                           `json:",omitzero,inline"`
	OfImageURL *OpenAIUserMessageParamContentArrayItemImageURL `json:",omitzero,inline"`
	OfFile     *OpenAIUserMessageParamContentArrayItemFile     `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIUserMessageParamContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL, u.OfFile)
}
func (u *OpenAIUserMessageParamContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIUserMessageParamContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
	} else if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIUserMessageParamContentArrayItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIUserMessageParamContentArrayItemUnion) GetImageURL() *OpenAIUserMessageParamContentArrayItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIUserMessageParamContentArrayItemUnion) GetFile() *OpenAIUserMessageParamContentArrayItemFileFile {
	if vt := u.OfFile; vt != nil {
		return &vt.File
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIUserMessageParamContentArrayItemUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[OpenAIUserMessageParamContentArrayItemUnion](
		"type",
		apijson.Discriminator[ContentPartTextParam]("text"),
		apijson.Discriminator[OpenAIUserMessageParamContentArrayItemImageURL]("image_url"),
		apijson.Discriminator[OpenAIUserMessageParamContentArrayItemFile]("file"),
	)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type OpenAIUserMessageParamContentArrayItemImageURL struct {
	// Image URL specification and processing details
	ImageURL OpenAIUserMessageParamContentArrayItemImageURLImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r OpenAIUserMessageParamContentArrayItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIUserMessageParamContentArrayItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIUserMessageParamContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type OpenAIUserMessageParamContentArrayItemImageURLImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r OpenAIUserMessageParamContentArrayItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIUserMessageParamContentArrayItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIUserMessageParamContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties File, Type are required.
type OpenAIUserMessageParamContentArrayItemFile struct {
	File OpenAIUserMessageParamContentArrayItemFileFile `json:"file,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "file".
	Type constant.File `json:"type,required"`
	paramObj
}

func (r OpenAIUserMessageParamContentArrayItemFile) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIUserMessageParamContentArrayItemFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIUserMessageParamContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIUserMessageParamContentArrayItemFileFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r OpenAIUserMessageParamContentArrayItemFileFile) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIUserMessageParamContentArrayItemFileFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIUserMessageParamContentArrayItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The log probability for a token from an OpenAI-compatible chat completion
// response.
type TokenLogProb struct {
	Token       string                   `json:"token,required"`
	Logprob     float64                  `json:"logprob,required"`
	TopLogprobs []TokenLogProbTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                  `json:"bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenLogProb) RawJSON() string { return r.JSON.raw }
func (r *TokenLogProb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The top log probability for a token from an OpenAI-compatible chat completion
// response.
type TokenLogProbTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   []int64 `json:"bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenLogProbTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *TokenLogProbTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAiv1ChatCompletionNewResponseUnion contains all possible properties and
// values from [OpenAiv1ChatCompletionNewResponseOpenAIChatCompletion],
// [OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunk].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAiv1ChatCompletionNewResponseUnion struct {
	ID string `json:"id"`
	// This field is a union of [[]Choice],
	// [[]OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoice]
	Choices OpenAiv1ChatCompletionNewResponseUnionChoices `json:"choices"`
	Created int64                                         `json:"created"`
	Model   string                                        `json:"model"`
	Object  string                                        `json:"object"`
	JSON    struct {
		ID      respjson.Field
		Choices respjson.Field
		Created respjson.Field
		Model   respjson.Field
		Object  respjson.Field
		raw     string
	} `json:"-"`
}

func (u OpenAiv1ChatCompletionNewResponseUnion) AsOpenAIChatCompletion() (v OpenAiv1ChatCompletionNewResponseOpenAIChatCompletion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionNewResponseUnion) AsOpenAIChatCompletionChunk() (v OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAiv1ChatCompletionNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *OpenAiv1ChatCompletionNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAiv1ChatCompletionNewResponseUnionChoices is an implicit subunion of
// [OpenAiv1ChatCompletionNewResponseUnion].
// OpenAiv1ChatCompletionNewResponseUnionChoices provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OpenAiv1ChatCompletionNewResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChoiceArray
// OfOpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoices]
type OpenAiv1ChatCompletionNewResponseUnionChoices struct {
	// This field will be present if the value is a [[]Choice] instead of an object.
	OfChoiceArray []Choice `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoice] instead of
	// an object.
	OfOpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoices []OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoice `json:",inline"`
	JSON                                                                struct {
		OfChoiceArray                                                       respjson.Field
		OfOpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoices respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (r *OpenAiv1ChatCompletionNewResponseUnionChoices) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from an OpenAI-compatible chat completion request.
type OpenAiv1ChatCompletionNewResponseOpenAIChatCompletion struct {
	// The ID of the chat completion
	ID string `json:"id,required"`
	// List of choices
	Choices []Choice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created
	Created int64 `json:"created,required"`
	// The model that was used to generate the chat completion
	Model string `json:"model,required"`
	// The object type, which will be "chat.completion"
	Object constant.ChatCompletion `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Choices     respjson.Field
		Created     respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ChatCompletionNewResponseOpenAIChatCompletion) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1ChatCompletionNewResponseOpenAIChatCompletion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chunk from a streaming response to an OpenAI-compatible chat completion request.
type OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunk struct {
	// The ID of the chat completion
	ID string `json:"id,required"`
	// List of choices
	Choices []OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created
	Created int64 `json:"created,required"`
	// The model that was used to generate the chat completion
	Model string `json:"model,required"`
	// The object type, which will be "chat.completion.chunk"
	Object constant.ChatCompletionChunk `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Choices     respjson.Field
		Created     respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunk) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chunk choice from an OpenAI-compatible chat completion streaming response.
type OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoice struct {
	// The delta from the chunk
	Delta OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoiceDelta `json:"delta,required"`
	// The reason the model stopped generating
	FinishReason string `json:"finish_reason,required"`
	// The index of the choice
	Index int64 `json:"index,required"`
	// (Optional) The log probabilities for the tokens in the message
	Logprobs ChoiceLogprobs `json:"logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta        respjson.Field
		FinishReason respjson.Field
		Index        respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoice) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The delta from the chunk
type OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoiceDelta struct {
	// (Optional) The content of the delta
	Content string `json:"content"`
	// (Optional) The refusal of the delta
	Refusal string `json:"refusal"`
	// (Optional) The role of the delta
	Role string `json:"role"`
	// (Optional) The tool calls of the delta
	ToolCalls []ChatCompletionToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		Role        respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoiceDelta) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAiv1ChatCompletionNewResponseOpenAIChatCompletionChunkChoiceDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAiv1ChatCompletionGetResponse struct {
	// The ID of the chat completion
	ID string `json:"id,required"`
	// List of choices
	Choices []Choice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created
	Created       int64                                                `json:"created,required"`
	InputMessages []OpenAiv1ChatCompletionGetResponseInputMessageUnion `json:"input_messages,required"`
	// The model that was used to generate the chat completion
	Model string `json:"model,required"`
	// The object type, which will be "chat.completion"
	Object constant.ChatCompletion `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Choices       respjson.Field
		Created       respjson.Field
		InputMessages respjson.Field
		Model         respjson.Field
		Object        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ChatCompletionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1ChatCompletionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAiv1ChatCompletionGetResponseInputMessageUnion contains all possible
// properties and values from [OpenAIUserMessageParamResp],
// [OpenAISystemMessageParamResp], [OpenAIAssistantMessageParamResp],
// [OpenAIToolMessageParamResp], [OpenAIDeveloperMessageParamResp].
//
// Use the [OpenAiv1ChatCompletionGetResponseInputMessageUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAiv1ChatCompletionGetResponseInputMessageUnion struct {
	// This field is a union of [OpenAIUserMessageParamContentUnionResp],
	// [OpenAISystemMessageParamContentUnionResp],
	// [OpenAIAssistantMessageParamContentUnionResp],
	// [OpenAIToolMessageParamContentUnionResp],
	// [OpenAIDeveloperMessageParamContentUnionResp]
	Content OpenAiv1ChatCompletionGetResponseInputMessageUnionContent `json:"content"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant [OpenAIAssistantMessageParamResp].
	ToolCalls []ChatCompletionToolCall `json:"tool_calls"`
	// This field is from variant [OpenAIToolMessageParamResp].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Role       respjson.Field
		Name       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyOpenAiv1ChatCompletionGetResponseInputMessage is implemented by each variant
// of [OpenAiv1ChatCompletionGetResponseInputMessageUnion] to add type safety for
// the return type of [OpenAiv1ChatCompletionGetResponseInputMessageUnion.AsAny]
type anyOpenAiv1ChatCompletionGetResponseInputMessage interface {
	implOpenAiv1ChatCompletionGetResponseInputMessageUnion()
}

func (OpenAIUserMessageParamResp) implOpenAiv1ChatCompletionGetResponseInputMessageUnion()      {}
func (OpenAISystemMessageParamResp) implOpenAiv1ChatCompletionGetResponseInputMessageUnion()    {}
func (OpenAIAssistantMessageParamResp) implOpenAiv1ChatCompletionGetResponseInputMessageUnion() {}
func (OpenAIToolMessageParamResp) implOpenAiv1ChatCompletionGetResponseInputMessageUnion()      {}
func (OpenAIDeveloperMessageParamResp) implOpenAiv1ChatCompletionGetResponseInputMessageUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := OpenAiv1ChatCompletionGetResponseInputMessageUnion.AsAny().(type) {
//	case llamastackclient.OpenAIUserMessageParamResp:
//	case llamastackclient.OpenAISystemMessageParamResp:
//	case llamastackclient.OpenAIAssistantMessageParamResp:
//	case llamastackclient.OpenAIToolMessageParamResp:
//	case llamastackclient.OpenAIDeveloperMessageParamResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsAny() anyOpenAiv1ChatCompletionGetResponseInputMessage {
	switch u.Role {
	case "user":
		return u.AsUser()
	case "system":
		return u.AsSystem()
	case "assistant":
		return u.AsAssistant()
	case "tool":
		return u.AsTool()
	case "developer":
		return u.AsDeveloper()
	}
	return nil
}

func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsUser() (v OpenAIUserMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsSystem() (v OpenAISystemMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsAssistant() (v OpenAIAssistantMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsTool() (v OpenAIToolMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsDeveloper() (v OpenAIDeveloperMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *OpenAiv1ChatCompletionGetResponseInputMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAiv1ChatCompletionGetResponseInputMessageUnionContent is an implicit
// subunion of [OpenAiv1ChatCompletionGetResponseInputMessageUnion].
// OpenAiv1ChatCompletionGetResponseInputMessageUnionContent provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OpenAiv1ChatCompletionGetResponseInputMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOpenAIUserMessageContentArray OfContentPartTextArray]
type OpenAiv1ChatCompletionGetResponseInputMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIUserMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIUserMessageContentArray []OpenAIUserMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a [[]ContentPartTextParamResp]
	// instead of an object.
	OfContentPartTextArray []ContentPartTextParamResp `json:",inline"`
	JSON                   struct {
		OfString                        respjson.Field
		OfOpenAIUserMessageContentArray respjson.Field
		OfContentPartTextArray          respjson.Field
		raw                             string
	} `json:"-"`
}

func (r *OpenAiv1ChatCompletionGetResponseInputMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from listing OpenAI-compatible chat completions.
type OpenAiv1ChatCompletionListResponse struct {
	// List of chat completion objects with their input messages
	Data []OpenAiv1ChatCompletionListResponseData `json:"data,required"`
	// ID of the first completion in this list
	FirstID string `json:"first_id,required"`
	// Whether there are more completions available beyond this list
	HasMore bool `json:"has_more,required"`
	// ID of the last completion in this list
	LastID string `json:"last_id,required"`
	// Must be "list" to identify this as a list response
	Object constant.List `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		FirstID     respjson.Field
		HasMore     respjson.Field
		LastID      respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ChatCompletionListResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1ChatCompletionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAiv1ChatCompletionListResponseData struct {
	// The ID of the chat completion
	ID string `json:"id,required"`
	// List of choices
	Choices []Choice `json:"choices,required"`
	// The Unix timestamp in seconds when the chat completion was created
	Created       int64                                                     `json:"created,required"`
	InputMessages []OpenAiv1ChatCompletionListResponseDataInputMessageUnion `json:"input_messages,required"`
	// The model that was used to generate the chat completion
	Model string `json:"model,required"`
	// The object type, which will be "chat.completion"
	Object constant.ChatCompletion `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Choices       respjson.Field
		Created       respjson.Field
		InputMessages respjson.Field
		Model         respjson.Field
		Object        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ChatCompletionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1ChatCompletionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAiv1ChatCompletionListResponseDataInputMessageUnion contains all possible
// properties and values from [OpenAIUserMessageParamResp],
// [OpenAISystemMessageParamResp], [OpenAIAssistantMessageParamResp],
// [OpenAIToolMessageParamResp], [OpenAIDeveloperMessageParamResp].
//
// Use the [OpenAiv1ChatCompletionListResponseDataInputMessageUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAiv1ChatCompletionListResponseDataInputMessageUnion struct {
	// This field is a union of [OpenAIUserMessageParamContentUnionResp],
	// [OpenAISystemMessageParamContentUnionResp],
	// [OpenAIAssistantMessageParamContentUnionResp],
	// [OpenAIToolMessageParamContentUnionResp],
	// [OpenAIDeveloperMessageParamContentUnionResp]
	Content OpenAiv1ChatCompletionListResponseDataInputMessageUnionContent `json:"content"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant [OpenAIAssistantMessageParamResp].
	ToolCalls []ChatCompletionToolCall `json:"tool_calls"`
	// This field is from variant [OpenAIToolMessageParamResp].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Content    respjson.Field
		Role       respjson.Field
		Name       respjson.Field
		ToolCalls  respjson.Field
		ToolCallID respjson.Field
		raw        string
	} `json:"-"`
}

// anyOpenAiv1ChatCompletionListResponseDataInputMessage is implemented by each
// variant of [OpenAiv1ChatCompletionListResponseDataInputMessageUnion] to add type
// safety for the return type of
// [OpenAiv1ChatCompletionListResponseDataInputMessageUnion.AsAny]
type anyOpenAiv1ChatCompletionListResponseDataInputMessage interface {
	implOpenAiv1ChatCompletionListResponseDataInputMessageUnion()
}

func (OpenAIUserMessageParamResp) implOpenAiv1ChatCompletionListResponseDataInputMessageUnion()   {}
func (OpenAISystemMessageParamResp) implOpenAiv1ChatCompletionListResponseDataInputMessageUnion() {}
func (OpenAIAssistantMessageParamResp) implOpenAiv1ChatCompletionListResponseDataInputMessageUnion() {
}
func (OpenAIToolMessageParamResp) implOpenAiv1ChatCompletionListResponseDataInputMessageUnion() {}
func (OpenAIDeveloperMessageParamResp) implOpenAiv1ChatCompletionListResponseDataInputMessageUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := OpenAiv1ChatCompletionListResponseDataInputMessageUnion.AsAny().(type) {
//	case llamastackclient.OpenAIUserMessageParamResp:
//	case llamastackclient.OpenAISystemMessageParamResp:
//	case llamastackclient.OpenAIAssistantMessageParamResp:
//	case llamastackclient.OpenAIToolMessageParamResp:
//	case llamastackclient.OpenAIDeveloperMessageParamResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsAny() anyOpenAiv1ChatCompletionListResponseDataInputMessage {
	switch u.Role {
	case "user":
		return u.AsUser()
	case "system":
		return u.AsSystem()
	case "assistant":
		return u.AsAssistant()
	case "tool":
		return u.AsTool()
	case "developer":
		return u.AsDeveloper()
	}
	return nil
}

func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsUser() (v OpenAIUserMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsSystem() (v OpenAISystemMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsAssistant() (v OpenAIAssistantMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsTool() (v OpenAIToolMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsDeveloper() (v OpenAIDeveloperMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *OpenAiv1ChatCompletionListResponseDataInputMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAiv1ChatCompletionListResponseDataInputMessageUnionContent is an implicit
// subunion of [OpenAiv1ChatCompletionListResponseDataInputMessageUnion].
// OpenAiv1ChatCompletionListResponseDataInputMessageUnionContent provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OpenAiv1ChatCompletionListResponseDataInputMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOpenAIUserMessageContentArray OfContentPartTextArray]
type OpenAiv1ChatCompletionListResponseDataInputMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIUserMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIUserMessageContentArray []OpenAIUserMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a [[]ContentPartTextParamResp]
	// instead of an object.
	OfContentPartTextArray []ContentPartTextParamResp `json:",inline"`
	JSON                   struct {
		OfString                        respjson.Field
		OfOpenAIUserMessageContentArray respjson.Field
		OfContentPartTextArray          respjson.Field
		raw                             string
	} `json:"-"`
}

func (r *OpenAiv1ChatCompletionListResponseDataInputMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIV1ChatCompletionNewParams struct {
	// List of messages in the conversation.
	Messages []OpenAIV1ChatCompletionNewParamsMessageUnion `json:"messages,omitzero,required"`
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	Model string `json:"model,required"`
	// (Optional) The penalty for repeated tokens.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// (Optional) The log probabilities to use.
	Logprobs param.Opt[bool] `json:"logprobs,omitzero"`
	// (Optional) The maximum number of tokens to generate.
	MaxCompletionTokens param.Opt[int64] `json:"max_completion_tokens,omitzero"`
	// (Optional) The maximum number of tokens to generate.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// (Optional) The number of completions to generate.
	N param.Opt[int64] `json:"n,omitzero"`
	// (Optional) Whether to parallelize tool calls.
	ParallelToolCalls param.Opt[bool] `json:"parallel_tool_calls,omitzero"`
	// (Optional) The penalty for repeated tokens.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// (Optional) The seed to use.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// (Optional) Whether to stream the response.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// (Optional) The temperature to use.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// (Optional) The top log probabilities to use.
	TopLogprobs param.Opt[int64] `json:"top_logprobs,omitzero"`
	// (Optional) The top p to use.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// (Optional) The user to use.
	User param.Opt[string] `json:"user,omitzero"`
	// (Optional) The function call to use.
	FunctionCall OpenAIV1ChatCompletionNewParamsFunctionCallUnion `json:"function_call,omitzero"`
	// (Optional) List of functions to use.
	Functions []map[string]OpenAIV1ChatCompletionNewParamsFunctionUnion `json:"functions,omitzero"`
	// (Optional) The logit bias to use.
	LogitBias map[string]float64 `json:"logit_bias,omitzero"`
	// (Optional) The response format to use.
	ResponseFormat OpenAIV1ChatCompletionNewParamsResponseFormatUnion `json:"response_format,omitzero"`
	// (Optional) The stop tokens to use.
	Stop OpenAIV1ChatCompletionNewParamsStopUnion `json:"stop,omitzero"`
	// (Optional) The stream options to use.
	StreamOptions map[string]OpenAIV1ChatCompletionNewParamsStreamOptionUnion `json:"stream_options,omitzero"`
	// (Optional) The tool choice to use.
	ToolChoice OpenAIV1ChatCompletionNewParamsToolChoiceUnion `json:"tool_choice,omitzero"`
	// (Optional) The tools to use.
	Tools []map[string]OpenAIV1ChatCompletionNewParamsToolUnion `json:"tools,omitzero"`
	paramObj
}

func (r OpenAIV1ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ChatCompletionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ChatCompletionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsMessageUnion struct {
	OfUser      *OpenAIUserMessageParam      `json:",omitzero,inline"`
	OfSystem    *OpenAISystemMessageParam    `json:",omitzero,inline"`
	OfAssistant *OpenAIAssistantMessageParam `json:",omitzero,inline"`
	OfTool      *OpenAIToolMessageParam      `json:",omitzero,inline"`
	OfDeveloper *OpenAIDeveloperMessageParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsMessageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUser,
		u.OfSystem,
		u.OfAssistant,
		u.OfTool,
		u.OfDeveloper)
}
func (u *OpenAIV1ChatCompletionNewParamsMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsMessageUnion) asAny() any {
	if !param.IsOmitted(u.OfUser) {
		return u.OfUser
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfAssistant) {
		return u.OfAssistant
	} else if !param.IsOmitted(u.OfTool) {
		return u.OfTool
	} else if !param.IsOmitted(u.OfDeveloper) {
		return u.OfDeveloper
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsMessageUnion) GetToolCalls() []ChatCompletionToolCallParam {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsMessageUnion) GetToolCallID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.ToolCallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsMessageUnion) GetRole() *string {
	if vt := u.OfUser; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfDeveloper; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsMessageUnion) GetName() *string {
	if vt := u.OfUser; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfSystem; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfAssistant; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfDeveloper; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u OpenAIV1ChatCompletionNewParamsMessageUnion) GetContent() (res openAiv1ChatCompletionNewParamsMessageUnionContent) {
	if vt := u.OfUser; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfSystem; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfAssistant; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfTool; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfDeveloper; vt != nil {
		res.any = vt.Content.asAny()
	}
	return
}

// Can have the runtime types [*string],
// [_[]OpenAIUserMessageParamContentArrayItemUnion], [_[]ContentPartTextParam]
type openAiv1ChatCompletionNewParamsMessageUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]llamastackclient.OpenAIUserMessageParamContentArrayItemUnion:
//	case *[]llamastackclient.ContentPartTextParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u openAiv1ChatCompletionNewParamsMessageUnionContent) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[OpenAIV1ChatCompletionNewParamsMessageUnion](
		"role",
		apijson.Discriminator[OpenAIUserMessageParam]("user"),
		apijson.Discriminator[OpenAISystemMessageParam]("system"),
		apijson.Discriminator[OpenAIAssistantMessageParam]("assistant"),
		apijson.Discriminator[OpenAIToolMessageParam]("tool"),
		apijson.Discriminator[OpenAIDeveloperMessageParam]("developer"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsFunctionCallUnion struct {
	OfString                                       param.Opt[string]                                                  `json:",omitzero,inline"`
	OfOpenAiv1ChatCompletionNewsFunctionCallMapMap map[string]OpenAIV1ChatCompletionNewParamsFunctionCallMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsFunctionCallUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAiv1ChatCompletionNewsFunctionCallMapMap)
}
func (u *OpenAIV1ChatCompletionNewParamsFunctionCallUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsFunctionCallUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAiv1ChatCompletionNewsFunctionCallMapMap) {
		return &u.OfOpenAiv1ChatCompletionNewsFunctionCallMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsFunctionCallMapItemUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsFunctionCallMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1ChatCompletionNewParamsFunctionCallMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsFunctionCallMapItemUnion) asAny() any {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsFunctionUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsFunctionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1ChatCompletionNewParamsFunctionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsFunctionUnion) asAny() any {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsResponseFormatUnion struct {
	OfText       *OpenAIV1ChatCompletionNewParamsResponseFormatText       `json:",omitzero,inline"`
	OfJsonSchema *OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchema `json:",omitzero,inline"`
	OfJsonObject *OpenAIV1ChatCompletionNewParamsResponseFormatJsonObject `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsResponseFormatUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfJsonSchema, u.OfJsonObject)
}
func (u *OpenAIV1ChatCompletionNewParamsResponseFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsResponseFormatUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfJsonSchema) {
		return u.OfJsonSchema
	} else if !param.IsOmitted(u.OfJsonObject) {
		return u.OfJsonObject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsResponseFormatUnion) GetJsonSchema() *OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema {
	if vt := u.OfJsonSchema; vt != nil {
		return &vt.JsonSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsResponseFormatUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJsonSchema; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJsonObject; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[OpenAIV1ChatCompletionNewParamsResponseFormatUnion](
		"type",
		apijson.Discriminator[OpenAIV1ChatCompletionNewParamsResponseFormatText]("text"),
		apijson.Discriminator[OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchema]("json_schema"),
		apijson.Discriminator[OpenAIV1ChatCompletionNewParamsResponseFormatJsonObject]("json_object"),
	)
}

func NewOpenAIV1ChatCompletionNewParamsResponseFormatText() OpenAIV1ChatCompletionNewParamsResponseFormatText {
	return OpenAIV1ChatCompletionNewParamsResponseFormatText{
		Type: "text",
	}
}

// Text response format for OpenAI-compatible chat completion requests.
//
// This struct has a constant value, construct it with
// [NewOpenAIV1ChatCompletionNewParamsResponseFormatText].
type OpenAIV1ChatCompletionNewParamsResponseFormatText struct {
	// Must be "text" to indicate plain text response format
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r OpenAIV1ChatCompletionNewParamsResponseFormatText) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ChatCompletionNewParamsResponseFormatText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ChatCompletionNewParamsResponseFormatText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JSON schema response format for OpenAI-compatible chat completion requests.
//
// The properties JsonSchema, Type are required.
type OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchema struct {
	// The JSON schema specification for the response
	JsonSchema OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema `json:"json_schema,omitzero,required"`
	// Must be "json_schema" to indicate structured JSON response format
	//
	// This field can be elided, and will marshal its zero value as "json_schema".
	Type constant.JsonSchema `json:"type,required"`
	paramObj
}

func (r OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The JSON schema specification for the response
//
// The property Name is required.
type OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema struct {
	// Name of the schema
	Name string `json:"name,required"`
	// (Optional) Description of the schema
	Description param.Opt[string] `json:"description,omitzero"`
	// (Optional) Whether to enforce strict adherence to the schema
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// (Optional) The JSON schema definition
	Schema map[string]OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion `json:"schema,omitzero"`
	paramObj
}

func (r OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchemaSchemaUnion) asAny() any {
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

func NewOpenAIV1ChatCompletionNewParamsResponseFormatJsonObject() OpenAIV1ChatCompletionNewParamsResponseFormatJsonObject {
	return OpenAIV1ChatCompletionNewParamsResponseFormatJsonObject{
		Type: "json_object",
	}
}

// JSON object response format for OpenAI-compatible chat completion requests.
//
// This struct has a constant value, construct it with
// [NewOpenAIV1ChatCompletionNewParamsResponseFormatJsonObject].
type OpenAIV1ChatCompletionNewParamsResponseFormatJsonObject struct {
	// Must be "json_object" to indicate generic JSON object response format
	Type constant.JsonObject `json:"type,required"`
	paramObj
}

func (r OpenAIV1ChatCompletionNewParamsResponseFormatJsonObject) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ChatCompletionNewParamsResponseFormatJsonObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ChatCompletionNewParamsResponseFormatJsonObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsStopUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsStopUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpenAIV1ChatCompletionNewParamsStopUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsStopUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsStreamOptionUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsStreamOptionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1ChatCompletionNewParamsStreamOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsStreamOptionUnion) asAny() any {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsToolChoiceUnion struct {
	OfString                                     param.Opt[string]                                                `json:",omitzero,inline"`
	OfOpenAiv1ChatCompletionNewsToolChoiceMapMap map[string]OpenAIV1ChatCompletionNewParamsToolChoiceMapItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsToolChoiceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAiv1ChatCompletionNewsToolChoiceMapMap)
}
func (u *OpenAIV1ChatCompletionNewParamsToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsToolChoiceUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAiv1ChatCompletionNewsToolChoiceMapMap) {
		return &u.OfOpenAiv1ChatCompletionNewsToolChoiceMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsToolChoiceMapItemUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsToolChoiceMapItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1ChatCompletionNewParamsToolChoiceMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsToolChoiceMapItemUnion) asAny() any {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsToolUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsToolUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1ChatCompletionNewParamsToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsToolUnion) asAny() any {
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

type OpenAIV1ChatCompletionListParams struct {
	// The ID of the last chat completion to return.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of chat completions to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The model to filter by.
	Model param.Opt[string] `query:"model,omitzero" json:"-"`
	// The order to sort the chat completions by: "asc" or "desc". Defaults to "desc".
	//
	// Any of "asc", "desc".
	Order Order `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OpenAIV1ChatCompletionListParams]'s query parameters as
// `url.Values`.
func (r OpenAIV1ChatCompletionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

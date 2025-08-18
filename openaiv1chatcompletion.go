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
	Message MessageParamUnionResp `json:"message,required"`
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

// MessageParamUnionResp contains all possible properties and values from
// [MessageParamUserResp], [MessageParamSystemResp], [MessageParamAssistantResp],
// [MessageParamToolResp], [MessageParamDeveloperResp].
//
// Use the [MessageParamUnionResp.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageParamUnionResp struct {
	// This field is a union of [MessageParamUserContentUnionResp],
	// [MessageParamSystemContentUnionResp], [MessageParamAssistantContentUnionResp],
	// [MessageParamToolContentUnionResp], [MessageParamDeveloperContentUnionResp]
	Content MessageParamUnionRespContent `json:"content"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant [MessageParamAssistantResp].
	ToolCalls []ChatCompletionToolCall `json:"tool_calls"`
	// This field is from variant [MessageParamToolResp].
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

// anyMessageParamResp is implemented by each variant of [MessageParamUnionResp] to
// add type safety for the return type of [MessageParamUnionResp.AsAny]
type anyMessageParamResp interface {
	implMessageParamUnionResp()
}

func (MessageParamUserResp) implMessageParamUnionResp()      {}
func (MessageParamSystemResp) implMessageParamUnionResp()    {}
func (MessageParamAssistantResp) implMessageParamUnionResp() {}
func (MessageParamToolResp) implMessageParamUnionResp()      {}
func (MessageParamDeveloperResp) implMessageParamUnionResp() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MessageParamUnionResp.AsAny().(type) {
//	case llamastackclient.MessageParamUserResp:
//	case llamastackclient.MessageParamSystemResp:
//	case llamastackclient.MessageParamAssistantResp:
//	case llamastackclient.MessageParamToolResp:
//	case llamastackclient.MessageParamDeveloperResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MessageParamUnionResp) AsAny() anyMessageParamResp {
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

func (u MessageParamUnionResp) AsUser() (v MessageParamUserResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamUnionResp) AsSystem() (v MessageParamSystemResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamUnionResp) AsAssistant() (v MessageParamAssistantResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamUnionResp) AsTool() (v MessageParamToolResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamUnionResp) AsDeveloper() (v MessageParamDeveloperResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageParamUnionResp) RawJSON() string { return u.JSON.raw }

func (r *MessageParamUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageParamUnionRespContent is an implicit subunion of [MessageParamUnionResp].
// MessageParamUnionRespContent provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageParamUnionResp].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfMessageUserContentArray OfContentPartTextArray]
type MessageParamUnionRespContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]MessageParamUserContentArrayItemUnionResp] instead of an object.
	OfMessageUserContentArray []MessageParamUserContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a [[]ContentPartTextParamResp]
	// instead of an object.
	OfContentPartTextArray []ContentPartTextParamResp `json:",inline"`
	JSON                   struct {
		OfString                  respjson.Field
		OfMessageUserContentArray respjson.Field
		OfContentPartTextArray    respjson.Field
		raw                       string
	} `json:"-"`
}

func (r *MessageParamUnionRespContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageParamUnionResp to a MessageParamUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageParamUnion.Overrides()
func (r MessageParamUnionResp) ToParam() MessageParamUnion {
	return param.Override[MessageParamUnion](json.RawMessage(r.RawJSON()))
}

// A message from the user in an OpenAI-compatible chat completion request.
type MessageParamUserResp struct {
	// The content of the message, which can include text and other media
	Content MessageParamUserContentUnionResp `json:"content,required"`
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
func (r MessageParamUserResp) RawJSON() string { return r.JSON.raw }
func (r *MessageParamUserResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageParamUserContentUnionResp contains all possible properties and values
// from [string], [[]MessageParamUserContentArrayItemUnionResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfMessageUserContentArray]
type MessageParamUserContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]MessageParamUserContentArrayItemUnionResp] instead of an object.
	OfMessageUserContentArray []MessageParamUserContentArrayItemUnionResp `json:",inline"`
	JSON                      struct {
		OfString                  respjson.Field
		OfMessageUserContentArray respjson.Field
		raw                       string
	} `json:"-"`
}

func (u MessageParamUserContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamUserContentUnionResp) AsMessageParamUserContentArrayRespArray() (v []MessageParamUserContentArrayItemUnionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageParamUserContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *MessageParamUserContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageParamUserContentArrayItemUnionResp contains all possible properties and
// values from [ContentPartTextParamResp],
// [MessageParamUserContentArrayItemImageURLResp],
// [MessageParamUserContentArrayItemFileResp].
//
// Use the [MessageParamUserContentArrayItemUnionResp.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageParamUserContentArrayItemUnionResp struct {
	// This field is from variant [ContentPartTextParamResp].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant [MessageParamUserContentArrayItemImageURLResp].
	ImageURL MessageParamUserContentArrayItemImageURLImageURLResp `json:"image_url"`
	// This field is from variant [MessageParamUserContentArrayItemFileResp].
	File MessageParamUserContentArrayItemFileFileResp `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyMessageParamUserContentArrayItemResp is implemented by each variant of
// [MessageParamUserContentArrayItemUnionResp] to add type safety for the return
// type of [MessageParamUserContentArrayItemUnionResp.AsAny]
type anyMessageParamUserContentArrayItemResp interface {
	implMessageParamUserContentArrayItemUnionResp()
}

func (ContentPartTextParamResp) implMessageParamUserContentArrayItemUnionResp()                     {}
func (MessageParamUserContentArrayItemImageURLResp) implMessageParamUserContentArrayItemUnionResp() {}
func (MessageParamUserContentArrayItemFileResp) implMessageParamUserContentArrayItemUnionResp()     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MessageParamUserContentArrayItemUnionResp.AsAny().(type) {
//	case llamastackclient.ContentPartTextParamResp:
//	case llamastackclient.MessageParamUserContentArrayItemImageURLResp:
//	case llamastackclient.MessageParamUserContentArrayItemFileResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MessageParamUserContentArrayItemUnionResp) AsAny() anyMessageParamUserContentArrayItemResp {
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

func (u MessageParamUserContentArrayItemUnionResp) AsText() (v ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamUserContentArrayItemUnionResp) AsImageURL() (v MessageParamUserContentArrayItemImageURLResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamUserContentArrayItemUnionResp) AsFile() (v MessageParamUserContentArrayItemFileResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageParamUserContentArrayItemUnionResp) RawJSON() string { return u.JSON.raw }

func (r *MessageParamUserContentArrayItemUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type MessageParamUserContentArrayItemImageURLResp struct {
	// Image URL specification and processing details
	ImageURL MessageParamUserContentArrayItemImageURLImageURLResp `json:"image_url,required"`
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
func (r MessageParamUserContentArrayItemImageURLResp) RawJSON() string { return r.JSON.raw }
func (r *MessageParamUserContentArrayItemImageURLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
type MessageParamUserContentArrayItemImageURLImageURLResp struct {
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
func (r MessageParamUserContentArrayItemImageURLImageURLResp) RawJSON() string { return r.JSON.raw }
func (r *MessageParamUserContentArrayItemImageURLImageURLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageParamUserContentArrayItemFileResp struct {
	File MessageParamUserContentArrayItemFileFileResp `json:"file,required"`
	Type constant.File                                `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageParamUserContentArrayItemFileResp) RawJSON() string { return r.JSON.raw }
func (r *MessageParamUserContentArrayItemFileResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageParamUserContentArrayItemFileFileResp struct {
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
func (r MessageParamUserContentArrayItemFileFileResp) RawJSON() string { return r.JSON.raw }
func (r *MessageParamUserContentArrayItemFileFileResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type MessageParamSystemResp struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content MessageParamSystemContentUnionResp `json:"content,required"`
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
func (r MessageParamSystemResp) RawJSON() string { return r.JSON.raw }
func (r *MessageParamSystemResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageParamSystemContentUnionResp contains all possible properties and values
// from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type MessageParamSystemContentUnionResp struct {
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

func (u MessageParamSystemContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamSystemContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageParamSystemContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *MessageParamSystemContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type MessageParamAssistantResp struct {
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// The content of the model's response
	Content MessageParamAssistantContentUnionResp `json:"content"`
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
func (r MessageParamAssistantResp) RawJSON() string { return r.JSON.raw }
func (r *MessageParamAssistantResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageParamAssistantContentUnionResp contains all possible properties and
// values from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type MessageParamAssistantContentUnionResp struct {
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

func (u MessageParamAssistantContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamAssistantContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageParamAssistantContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *MessageParamAssistantContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type MessageParamToolResp struct {
	// The response content from the tool
	Content MessageParamToolContentUnionResp `json:"content,required"`
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
func (r MessageParamToolResp) RawJSON() string { return r.JSON.raw }
func (r *MessageParamToolResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageParamToolContentUnionResp contains all possible properties and values
// from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type MessageParamToolContentUnionResp struct {
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

func (u MessageParamToolContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamToolContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageParamToolContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *MessageParamToolContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type MessageParamDeveloperResp struct {
	// The content of the developer message
	Content MessageParamDeveloperContentUnionResp `json:"content,required"`
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
func (r MessageParamDeveloperResp) RawJSON() string { return r.JSON.raw }
func (r *MessageParamDeveloperResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageParamDeveloperContentUnionResp contains all possible properties and
// values from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type MessageParamDeveloperContentUnionResp struct {
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

func (u MessageParamDeveloperContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageParamDeveloperContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageParamDeveloperContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *MessageParamDeveloperContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func MessageParamOfUser[T string | []MessageParamUserContentArrayItemUnion](content T) MessageParamUnion {
	var user MessageParamUser
	switch v := any(content).(type) {
	case string:
		user.Content.OfString = param.NewOpt(v)
	case []MessageParamUserContentArrayItemUnion:
		user.Content.OfMessageUserContentArray = v
	}
	return MessageParamUnion{OfUser: &user}
}

func MessageParamOfSystem[T string | []ContentPartTextParam](content T) MessageParamUnion {
	var system MessageParamSystem
	switch v := any(content).(type) {
	case string:
		system.Content.OfString = param.NewOpt(v)
	case []ContentPartTextParam:
		system.Content.OfContentPartTextArray = v
	}
	return MessageParamUnion{OfSystem: &system}
}

func MessageParamOfTool[T string | []ContentPartTextParam](content T, toolCallID string) MessageParamUnion {
	var tool MessageParamTool
	switch v := any(content).(type) {
	case string:
		tool.Content.OfString = param.NewOpt(v)
	case []ContentPartTextParam:
		tool.Content.OfContentPartTextArray = v
	}
	tool.ToolCallID = toolCallID
	return MessageParamUnion{OfTool: &tool}
}

func MessageParamOfDeveloper[T string | []ContentPartTextParam](content T) MessageParamUnion {
	var developer MessageParamDeveloper
	switch v := any(content).(type) {
	case string:
		developer.Content.OfString = param.NewOpt(v)
	case []ContentPartTextParam:
		developer.Content.OfContentPartTextArray = v
	}
	return MessageParamUnion{OfDeveloper: &developer}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageParamUnion struct {
	OfUser      *MessageParamUser      `json:",omitzero,inline"`
	OfSystem    *MessageParamSystem    `json:",omitzero,inline"`
	OfAssistant *MessageParamAssistant `json:",omitzero,inline"`
	OfTool      *MessageParamTool      `json:",omitzero,inline"`
	OfDeveloper *MessageParamDeveloper `json:",omitzero,inline"`
	paramUnion
}

func (u MessageParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUser,
		u.OfSystem,
		u.OfAssistant,
		u.OfTool,
		u.OfDeveloper)
}
func (u *MessageParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageParamUnion) asAny() any {
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
func (u MessageParamUnion) GetToolCalls() []ChatCompletionToolCallParam {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageParamUnion) GetToolCallID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.ToolCallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageParamUnion) GetRole() *string {
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
func (u MessageParamUnion) GetName() *string {
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
func (u MessageParamUnion) GetContent() (res messageParamUnionContent) {
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
// [_[]MessageParamUserContentArrayItemUnion], [_[]ContentPartTextParam]
type messageParamUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]llamastackclient.MessageParamUserContentArrayItemUnion:
//	case *[]llamastackclient.ContentPartTextParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u messageParamUnionContent) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[MessageParamUnion](
		"role",
		apijson.Discriminator[MessageParamUser]("user"),
		apijson.Discriminator[MessageParamSystem]("system"),
		apijson.Discriminator[MessageParamAssistant]("assistant"),
		apijson.Discriminator[MessageParamTool]("tool"),
		apijson.Discriminator[MessageParamDeveloper]("developer"),
	)
}

// A message from the user in an OpenAI-compatible chat completion request.
//
// The properties Content, Role are required.
type MessageParamUser struct {
	// The content of the message, which can include text and other media
	Content MessageParamUserContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the user message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "user" to identify this as a user message
	//
	// This field can be elided, and will marshal its zero value as "user".
	Role constant.User `json:"role,required"`
	paramObj
}

func (r MessageParamUser) MarshalJSON() (data []byte, err error) {
	type shadow MessageParamUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageParamUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageParamUserContentUnion struct {
	OfString                  param.Opt[string]                       `json:",omitzero,inline"`
	OfMessageUserContentArray []MessageParamUserContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u MessageParamUserContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfMessageUserContentArray)
}
func (u *MessageParamUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageParamUserContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfMessageUserContentArray) {
		return &u.OfMessageUserContentArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageParamUserContentArrayItemUnion struct {
	OfText     *ContentPartTextParam                     `json:",omitzero,inline"`
	OfImageURL *MessageParamUserContentArrayItemImageURL `json:",omitzero,inline"`
	OfFile     *MessageParamUserContentArrayItemFile     `json:",omitzero,inline"`
	paramUnion
}

func (u MessageParamUserContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL, u.OfFile)
}
func (u *MessageParamUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageParamUserContentArrayItemUnion) asAny() any {
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
func (u MessageParamUserContentArrayItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageParamUserContentArrayItemUnion) GetImageURL() *MessageParamUserContentArrayItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageParamUserContentArrayItemUnion) GetFile() *MessageParamUserContentArrayItemFileFile {
	if vt := u.OfFile; vt != nil {
		return &vt.File
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageParamUserContentArrayItemUnion) GetType() *string {
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
	apijson.RegisterUnion[MessageParamUserContentArrayItemUnion](
		"type",
		apijson.Discriminator[ContentPartTextParam]("text"),
		apijson.Discriminator[MessageParamUserContentArrayItemImageURL]("image_url"),
		apijson.Discriminator[MessageParamUserContentArrayItemFile]("file"),
	)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type MessageParamUserContentArrayItemImageURL struct {
	// Image URL specification and processing details
	ImageURL MessageParamUserContentArrayItemImageURLImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r MessageParamUserContentArrayItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow MessageParamUserContentArrayItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageParamUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type MessageParamUserContentArrayItemImageURLImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r MessageParamUserContentArrayItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow MessageParamUserContentArrayItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageParamUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties File, Type are required.
type MessageParamUserContentArrayItemFile struct {
	File MessageParamUserContentArrayItemFileFile `json:"file,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "file".
	Type constant.File `json:"type,required"`
	paramObj
}

func (r MessageParamUserContentArrayItemFile) MarshalJSON() (data []byte, err error) {
	type shadow MessageParamUserContentArrayItemFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageParamUserContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageParamUserContentArrayItemFileFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r MessageParamUserContentArrayItemFileFile) MarshalJSON() (data []byte, err error) {
	type shadow MessageParamUserContentArrayItemFileFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageParamUserContentArrayItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
//
// The properties Content, Role are required.
type MessageParamSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content MessageParamSystemContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the system message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "system" to identify this as a system message
	//
	// This field can be elided, and will marshal its zero value as "system".
	Role constant.System `json:"role,required"`
	paramObj
}

func (r MessageParamSystem) MarshalJSON() (data []byte, err error) {
	type shadow MessageParamSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageParamSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageParamSystemContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u MessageParamSystemContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *MessageParamSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageParamSystemContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfContentPartTextArray) {
		return &u.OfContentPartTextArray
	}
	return nil
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
//
// The property Role is required.
type MessageParamAssistant struct {
	// (Optional) The name of the assistant message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// The content of the model's response
	Content MessageParamAssistantContentUnion `json:"content,omitzero"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionToolCallParam `json:"tool_calls,omitzero"`
	// Must be "assistant" to identify this as the model's response
	//
	// This field can be elided, and will marshal its zero value as "assistant".
	Role constant.Assistant `json:"role,required"`
	paramObj
}

func (r MessageParamAssistant) MarshalJSON() (data []byte, err error) {
	type shadow MessageParamAssistant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageParamAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageParamAssistantContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u MessageParamAssistantContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *MessageParamAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageParamAssistantContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfContentPartTextArray) {
		return &u.OfContentPartTextArray
	}
	return nil
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
//
// The properties Content, Role, ToolCallID are required.
type MessageParamTool struct {
	// The response content from the tool
	Content MessageParamToolContentUnion `json:"content,omitzero,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// Must be "tool" to identify this as a tool response
	//
	// This field can be elided, and will marshal its zero value as "tool".
	Role constant.Tool `json:"role,required"`
	paramObj
}

func (r MessageParamTool) MarshalJSON() (data []byte, err error) {
	type shadow MessageParamTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageParamTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageParamToolContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u MessageParamToolContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *MessageParamToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageParamToolContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfContentPartTextArray) {
		return &u.OfContentPartTextArray
	}
	return nil
}

// A message from the developer in an OpenAI-compatible chat completion request.
//
// The properties Content, Role are required.
type MessageParamDeveloper struct {
	// The content of the developer message
	Content MessageParamDeveloperContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the developer message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "developer" to identify this as a developer message
	//
	// This field can be elided, and will marshal its zero value as "developer".
	Role constant.Developer `json:"role,required"`
	paramObj
}

func (r MessageParamDeveloper) MarshalJSON() (data []byte, err error) {
	type shadow MessageParamDeveloper
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageParamDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageParamDeveloperContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u MessageParamDeveloperContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *MessageParamDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageParamDeveloperContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfContentPartTextArray) {
		return &u.OfContentPartTextArray
	}
	return nil
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
	Created       int64                   `json:"created,required"`
	InputMessages []MessageParamUnionResp `json:"input_messages,required"`
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
	Created       int64                   `json:"created,required"`
	InputMessages []MessageParamUnionResp `json:"input_messages,required"`
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

type OpenAIV1ChatCompletionNewParams struct {
	// List of messages in the conversation.
	Messages []MessageParamUnion `json:"messages,omitzero,required"`
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

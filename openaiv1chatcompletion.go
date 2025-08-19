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

type ChatCompletionToolCall struct {
	Type     string                         `json:"type,required"`
	ID       string                         `json:"id"`
	Function ChatCompletionToolCallFunction `json:"function"`
	Index    int64                          `json:"index"`
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

type ChatCompletionToolCallFunction struct {
	Arguments string `json:"arguments"`
	Name      string `json:"name"`
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

// The property Type is required.
type ChatCompletionToolCallParam struct {
	Type     string                              `json:"type,required"`
	ID       param.Opt[string]                   `json:"id,omitzero"`
	Index    param.Opt[int64]                    `json:"index,omitzero"`
	Function ChatCompletionToolCallFunctionParam `json:"function,omitzero"`
	paramObj
}

func (r ChatCompletionToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionToolCallFunctionParam struct {
	Arguments param.Opt[string] `json:"arguments,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
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
	// Any of nil, nil, nil, nil, nil.
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

func (u ChoiceMessageUnion) AsOpenAIUserMessageParam() (v OpenAIUserMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChoiceMessageUnion) AsOpenAISystemMessageParam() (v OpenAISystemMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChoiceMessageUnion) AsOpenAIAssistantMessageParam() (v OpenAIAssistantMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChoiceMessageUnion) AsOpenAIToolMessageParam() (v OpenAIToolMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChoiceMessageUnion) AsOpenAIDeveloperMessageParam() (v OpenAIDeveloperMessageParamResp) {
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
// will be valid: OfString OfOpenAIUserMessageContentArray
// OfOpenAISystemMessageContentArray OfOpenAIAssistantMessageContentArray
// OfOpenAIToolMessageContentArray OfOpenAIDeveloperMessageContentArray]
type ChoiceMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIUserMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIUserMessageContentArray []OpenAIUserMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAISystemMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAISystemMessageContentArray []OpenAISystemMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIAssistantMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIAssistantMessageContentArray []OpenAIAssistantMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIToolMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIToolMessageContentArray []OpenAIToolMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIDeveloperMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIDeveloperMessageContentArray []OpenAIDeveloperMessageParamContentArrayItemUnionResp `json:",inline"`
	JSON                                 struct {
		OfString                             respjson.Field
		OfOpenAIUserMessageContentArray      respjson.Field
		OfOpenAISystemMessageContentArray    respjson.Field
		OfOpenAIAssistantMessageContentArray respjson.Field
		OfOpenAIToolMessageContentArray      respjson.Field
		OfOpenAIDeveloperMessageContentArray respjson.Field
		raw                                  string
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

type ContentPartTextParamResp struct {
	Text string `json:"text,required"`
	Type string `json:"type,required"`
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

// The properties Text, Type are required.
type ContentPartTextParam struct {
	Text string `json:"text,required"`
	Type string `json:"type,required"`
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
	Role string `json:"role,required"`
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
// values from [string], [[]OpenAIAssistantMessageParamContentArrayItemUnionResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOpenAIAssistantMessageContentArray]
type OpenAIAssistantMessageParamContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIAssistantMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIAssistantMessageContentArray []OpenAIAssistantMessageParamContentArrayItemUnionResp `json:",inline"`
	JSON                                 struct {
		OfString                             respjson.Field
		OfOpenAIAssistantMessageContentArray respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u OpenAIAssistantMessageParamContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIAssistantMessageParamContentUnionResp) AsOpenAIAssistantMessageParamContentArrayRespArray() (v []OpenAIAssistantMessageParamContentArrayItemUnionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIAssistantMessageParamContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIAssistantMessageParamContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIAssistantMessageParamContentArrayItemUnionResp contains all possible
// properties and values from [ContentPartTextParamResp],
// [OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp].
//
// Use the [OpenAIAssistantMessageParamContentArrayItemUnionResp.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAIAssistantMessageParamContentArrayItemUnionResp struct {
	// This field is from variant [ContentPartTextParamResp].
	Text string `json:"text"`
	// Any of nil, nil.
	Type string `json:"type"`
	// This field is from variant
	// [OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp].
	ImageURL OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

func (u OpenAIAssistantMessageParamContentArrayItemUnionResp) AsOpenAIChatCompletionContentPartTextParam() (v ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIAssistantMessageParamContentArrayItemUnionResp) AsOpenAIChatCompletionContentPartImageParam() (v OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIAssistantMessageParamContentArrayItemUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIAssistantMessageParamContentArrayItemUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp struct {
	ImageURL OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp `json:"image_url,required"`
	Type     string                                                                                           `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp struct {
	URL    string `json:"url,required"`
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
func (r OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
//
// The property Role is required.
type OpenAIAssistantMessageParam struct {
	// Must be "assistant" to identify this as the model's response
	Role string `json:"role,required"`
	// (Optional) The name of the assistant message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// The content of the model's response
	Content OpenAIAssistantMessageParamContentUnion `json:"content,omitzero"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionToolCallParam `json:"tool_calls,omitzero"`
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
	OfString                             param.Opt[string]                                  `json:",omitzero,inline"`
	OfOpenAIAssistantMessageContentArray []OpenAIAssistantMessageParamContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIAssistantMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIAssistantMessageContentArray)
}
func (u *OpenAIAssistantMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIAssistantMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAIAssistantMessageContentArray) {
		return &u.OfOpenAIAssistantMessageContentArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIAssistantMessageParamContentArrayItemUnion struct {
	OfOpenAIChatCompletionContentPartText  *ContentPartTextParam                                                                 `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIAssistantMessageParamContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *OpenAIAssistantMessageParamContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIAssistantMessageParamContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartText) {
		return u.OfOpenAIChatCompletionContentPartText
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartImage) {
		return u.OfOpenAIChatCompletionContentPartImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIAssistantMessageParamContentArrayItemUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIAssistantMessageParamContentArrayItemUnion) GetImageURL() *OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIAssistantMessageParamContentArrayItemUnion) GetType() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[OpenAIAssistantMessageParamContentArrayItemUnion](
		"type",
		apijson.Discriminator[ContentPartTextParam](undefined),
		apijson.Discriminator[OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam](undefined),
	)
}

// The properties ImageURL, Type are required.
type OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam struct {
	ImageURL OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	Type     string                                                                                       `json:"type,required"`
	paramObj
}

func (r OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL struct {
	URL    string            `json:"url,required"`
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIAssistantMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type OpenAIDeveloperMessageParamResp struct {
	// The content of the developer message
	Content OpenAIDeveloperMessageParamContentUnionResp `json:"content,required"`
	// Must be "developer" to identify this as a developer message
	Role string `json:"role,required"`
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
// values from [string], [[]OpenAIDeveloperMessageParamContentArrayItemUnionResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOpenAIDeveloperMessageContentArray]
type OpenAIDeveloperMessageParamContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIDeveloperMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIDeveloperMessageContentArray []OpenAIDeveloperMessageParamContentArrayItemUnionResp `json:",inline"`
	JSON                                 struct {
		OfString                             respjson.Field
		OfOpenAIDeveloperMessageContentArray respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u OpenAIDeveloperMessageParamContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIDeveloperMessageParamContentUnionResp) AsOpenAIDeveloperMessageParamContentArrayRespArray() (v []OpenAIDeveloperMessageParamContentArrayItemUnionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIDeveloperMessageParamContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIDeveloperMessageParamContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIDeveloperMessageParamContentArrayItemUnionResp contains all possible
// properties and values from [ContentPartTextParamResp],
// [OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp].
//
// Use the [OpenAIDeveloperMessageParamContentArrayItemUnionResp.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAIDeveloperMessageParamContentArrayItemUnionResp struct {
	// This field is from variant [ContentPartTextParamResp].
	Text string `json:"text"`
	// Any of nil, nil.
	Type string `json:"type"`
	// This field is from variant
	// [OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp].
	ImageURL OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

func (u OpenAIDeveloperMessageParamContentArrayItemUnionResp) AsOpenAIChatCompletionContentPartTextParam() (v ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIDeveloperMessageParamContentArrayItemUnionResp) AsOpenAIChatCompletionContentPartImageParam() (v OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIDeveloperMessageParamContentArrayItemUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIDeveloperMessageParamContentArrayItemUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp struct {
	ImageURL OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp `json:"image_url,required"`
	Type     string                                                                                           `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp struct {
	URL    string `json:"url,required"`
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
func (r OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
//
// The properties Content, Role are required.
type OpenAIDeveloperMessageParam struct {
	// The content of the developer message
	Content OpenAIDeveloperMessageParamContentUnion `json:"content,omitzero,required"`
	// Must be "developer" to identify this as a developer message
	Role string `json:"role,required"`
	// (Optional) The name of the developer message participant.
	Name param.Opt[string] `json:"name,omitzero"`
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
	OfString                             param.Opt[string]                                  `json:",omitzero,inline"`
	OfOpenAIDeveloperMessageContentArray []OpenAIDeveloperMessageParamContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIDeveloperMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIDeveloperMessageContentArray)
}
func (u *OpenAIDeveloperMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIDeveloperMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAIDeveloperMessageContentArray) {
		return &u.OfOpenAIDeveloperMessageContentArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIDeveloperMessageParamContentArrayItemUnion struct {
	OfOpenAIChatCompletionContentPartText  *ContentPartTextParam                                                                 `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIDeveloperMessageParamContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *OpenAIDeveloperMessageParamContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIDeveloperMessageParamContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartText) {
		return u.OfOpenAIChatCompletionContentPartText
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartImage) {
		return u.OfOpenAIChatCompletionContentPartImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIDeveloperMessageParamContentArrayItemUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIDeveloperMessageParamContentArrayItemUnion) GetImageURL() *OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIDeveloperMessageParamContentArrayItemUnion) GetType() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[OpenAIDeveloperMessageParamContentArrayItemUnion](
		"type",
		apijson.Discriminator[ContentPartTextParam](undefined),
		apijson.Discriminator[OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam](undefined),
	)
}

// The properties ImageURL, Type are required.
type OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam struct {
	ImageURL OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	Type     string                                                                                       `json:"type,required"`
	paramObj
}

func (r OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL struct {
	URL    string            `json:"url,required"`
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIDeveloperMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type OpenAISystemMessageParamResp struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content OpenAISystemMessageParamContentUnionResp `json:"content,required"`
	// Must be "system" to identify this as a system message
	Role string `json:"role,required"`
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
// values from [string], [[]OpenAISystemMessageParamContentArrayItemUnionResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOpenAISystemMessageContentArray]
type OpenAISystemMessageParamContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAISystemMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAISystemMessageContentArray []OpenAISystemMessageParamContentArrayItemUnionResp `json:",inline"`
	JSON                              struct {
		OfString                          respjson.Field
		OfOpenAISystemMessageContentArray respjson.Field
		raw                               string
	} `json:"-"`
}

func (u OpenAISystemMessageParamContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAISystemMessageParamContentUnionResp) AsOpenAISystemMessageParamContentArrayRespArray() (v []OpenAISystemMessageParamContentArrayItemUnionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAISystemMessageParamContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAISystemMessageParamContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAISystemMessageParamContentArrayItemUnionResp contains all possible
// properties and values from [ContentPartTextParamResp],
// [OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp].
//
// Use the [OpenAISystemMessageParamContentArrayItemUnionResp.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAISystemMessageParamContentArrayItemUnionResp struct {
	// This field is from variant [ContentPartTextParamResp].
	Text string `json:"text"`
	// Any of nil, nil.
	Type string `json:"type"`
	// This field is from variant
	// [OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp].
	ImageURL OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

func (u OpenAISystemMessageParamContentArrayItemUnionResp) AsOpenAIChatCompletionContentPartTextParam() (v ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAISystemMessageParamContentArrayItemUnionResp) AsOpenAIChatCompletionContentPartImageParam() (v OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAISystemMessageParamContentArrayItemUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAISystemMessageParamContentArrayItemUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp struct {
	ImageURL OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp `json:"image_url,required"`
	Type     string                                                                                        `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp struct {
	URL    string `json:"url,required"`
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
func (r OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp) UnmarshalJSON(data []byte) error {
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
	// Must be "system" to identify this as a system message
	Role string `json:"role,required"`
	// (Optional) The name of the system message participant.
	Name param.Opt[string] `json:"name,omitzero"`
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
	OfString                          param.Opt[string]                               `json:",omitzero,inline"`
	OfOpenAISystemMessageContentArray []OpenAISystemMessageParamContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAISystemMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAISystemMessageContentArray)
}
func (u *OpenAISystemMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAISystemMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAISystemMessageContentArray) {
		return &u.OfOpenAISystemMessageContentArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAISystemMessageParamContentArrayItemUnion struct {
	OfOpenAIChatCompletionContentPartText  *ContentPartTextParam                                                              `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAISystemMessageParamContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *OpenAISystemMessageParamContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAISystemMessageParamContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartText) {
		return u.OfOpenAIChatCompletionContentPartText
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartImage) {
		return u.OfOpenAIChatCompletionContentPartImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAISystemMessageParamContentArrayItemUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAISystemMessageParamContentArrayItemUnion) GetImageURL() *OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAISystemMessageParamContentArrayItemUnion) GetType() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[OpenAISystemMessageParamContentArrayItemUnion](
		"type",
		apijson.Discriminator[ContentPartTextParam](undefined),
		apijson.Discriminator[OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam](undefined),
	)
}

// The properties ImageURL, Type are required.
type OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam struct {
	ImageURL OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	Type     string                                                                                    `json:"type,required"`
	paramObj
}

func (r OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL struct {
	URL    string            `json:"url,required"`
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAISystemMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type OpenAIToolMessageParamResp struct {
	// The response content from the tool
	Content OpenAIToolMessageParamContentUnionResp `json:"content,required"`
	// Must be "tool" to identify this as a tool response
	Role string `json:"role,required"`
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
// values from [string], [[]OpenAIToolMessageParamContentArrayItemUnionResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOpenAIToolMessageContentArray]
type OpenAIToolMessageParamContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIToolMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIToolMessageContentArray []OpenAIToolMessageParamContentArrayItemUnionResp `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfOpenAIToolMessageContentArray respjson.Field
		raw                             string
	} `json:"-"`
}

func (u OpenAIToolMessageParamContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIToolMessageParamContentUnionResp) AsOpenAIToolMessageParamContentArrayRespArray() (v []OpenAIToolMessageParamContentArrayItemUnionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIToolMessageParamContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIToolMessageParamContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIToolMessageParamContentArrayItemUnionResp contains all possible properties
// and values from [ContentPartTextParamResp],
// [OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp].
//
// Use the [OpenAIToolMessageParamContentArrayItemUnionResp.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAIToolMessageParamContentArrayItemUnionResp struct {
	// This field is from variant [ContentPartTextParamResp].
	Text string `json:"text"`
	// Any of nil, nil.
	Type string `json:"type"`
	// This field is from variant
	// [OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp].
	ImageURL OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

func (u OpenAIToolMessageParamContentArrayItemUnionResp) AsOpenAIChatCompletionContentPartTextParam() (v ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIToolMessageParamContentArrayItemUnionResp) AsOpenAIChatCompletionContentPartImageParam() (v OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIToolMessageParamContentArrayItemUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIToolMessageParamContentArrayItemUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp struct {
	ImageURL OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp `json:"image_url,required"`
	Type     string                                                                                      `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp struct {
	URL    string `json:"url,required"`
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
func (r OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
//
// The properties Content, Role, ToolCallID are required.
type OpenAIToolMessageParam struct {
	// The response content from the tool
	Content OpenAIToolMessageParamContentUnion `json:"content,omitzero,required"`
	// Must be "tool" to identify this as a tool response
	Role string `json:"role,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
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
	OfString                        param.Opt[string]                             `json:",omitzero,inline"`
	OfOpenAIToolMessageContentArray []OpenAIToolMessageParamContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIToolMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIToolMessageContentArray)
}
func (u *OpenAIToolMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIToolMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAIToolMessageContentArray) {
		return &u.OfOpenAIToolMessageContentArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIToolMessageParamContentArrayItemUnion struct {
	OfOpenAIChatCompletionContentPartText  *ContentPartTextParam                                                            `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIToolMessageParamContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *OpenAIToolMessageParamContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIToolMessageParamContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartText) {
		return u.OfOpenAIChatCompletionContentPartText
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartImage) {
		return u.OfOpenAIChatCompletionContentPartImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIToolMessageParamContentArrayItemUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIToolMessageParamContentArrayItemUnion) GetImageURL() *OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIToolMessageParamContentArrayItemUnion) GetType() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[OpenAIToolMessageParamContentArrayItemUnion](
		"type",
		apijson.Discriminator[ContentPartTextParam](undefined),
		apijson.Discriminator[OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam](undefined),
	)
}

// The properties ImageURL, Type are required.
type OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam struct {
	ImageURL OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	Type     string                                                                                  `json:"type,required"`
	paramObj
}

func (r OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL struct {
	URL    string            `json:"url,required"`
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIToolMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
type OpenAIUserMessageParamResp struct {
	// The content of the message, which can include text and other media
	Content OpenAIUserMessageParamContentUnionResp `json:"content,required"`
	// Must be "user" to identify this as a user message
	Role string `json:"role,required"`
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
// [OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp].
//
// Use the [OpenAIUserMessageParamContentArrayItemUnionResp.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAIUserMessageParamContentArrayItemUnionResp struct {
	// This field is from variant [ContentPartTextParamResp].
	Text string `json:"text"`
	// Any of nil, nil.
	Type string `json:"type"`
	// This field is from variant
	// [OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp].
	ImageURL OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

func (u OpenAIUserMessageParamContentArrayItemUnionResp) AsOpenAIChatCompletionContentPartTextParam() (v ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIUserMessageParamContentArrayItemUnionResp) AsOpenAIChatCompletionContentPartImageParam() (v OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIUserMessageParamContentArrayItemUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIUserMessageParamContentArrayItemUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp struct {
	ImageURL OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp `json:"image_url,required"`
	Type     string                                                                                      `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp struct {
	URL    string `json:"url,required"`
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
func (r OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in an OpenAI-compatible chat completion request.
//
// The properties Content, Role are required.
type OpenAIUserMessageParam struct {
	// The content of the message, which can include text and other media
	Content OpenAIUserMessageParamContentUnion `json:"content,omitzero,required"`
	// Must be "user" to identify this as a user message
	Role string `json:"role,required"`
	// (Optional) The name of the user message participant.
	Name param.Opt[string] `json:"name,omitzero"`
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
	OfOpenAIChatCompletionContentPartText  *ContentPartTextParam                                                            `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIUserMessageParamContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *OpenAIUserMessageParamContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIUserMessageParamContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartText) {
		return u.OfOpenAIChatCompletionContentPartText
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartImage) {
		return u.OfOpenAIChatCompletionContentPartImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIUserMessageParamContentArrayItemUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIUserMessageParamContentArrayItemUnion) GetImageURL() *OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIUserMessageParamContentArrayItemUnion) GetType() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[OpenAIUserMessageParamContentArrayItemUnion](
		"type",
		apijson.Discriminator[ContentPartTextParam](undefined),
		apijson.Discriminator[OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam](undefined),
	)
}

// The properties ImageURL, Type are required.
type OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam struct {
	ImageURL OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	Type     string                                                                                  `json:"type,required"`
	paramObj
}

func (r OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL struct {
	URL    string            `json:"url,required"`
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIUserMessageParamContentArrayItemOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
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
	Object string `json:"object,required"`
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
	// Any of nil, nil, nil, nil, nil.
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

func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsOpenAIUserMessageParam() (v OpenAIUserMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsOpenAISystemMessageParam() (v OpenAISystemMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsOpenAIAssistantMessageParam() (v OpenAIAssistantMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsOpenAIToolMessageParam() (v OpenAIToolMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionGetResponseInputMessageUnion) AsOpenAIDeveloperMessageParam() (v OpenAIDeveloperMessageParamResp) {
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
// will be valid: OfString OfOpenAIUserMessageContentArray
// OfOpenAISystemMessageContentArray OfOpenAIAssistantMessageContentArray
// OfOpenAIToolMessageContentArray OfOpenAIDeveloperMessageContentArray]
type OpenAiv1ChatCompletionGetResponseInputMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIUserMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIUserMessageContentArray []OpenAIUserMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAISystemMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAISystemMessageContentArray []OpenAISystemMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIAssistantMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIAssistantMessageContentArray []OpenAIAssistantMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIToolMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIToolMessageContentArray []OpenAIToolMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIDeveloperMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIDeveloperMessageContentArray []OpenAIDeveloperMessageParamContentArrayItemUnionResp `json:",inline"`
	JSON                                 struct {
		OfString                             respjson.Field
		OfOpenAIUserMessageContentArray      respjson.Field
		OfOpenAISystemMessageContentArray    respjson.Field
		OfOpenAIAssistantMessageContentArray respjson.Field
		OfOpenAIToolMessageContentArray      respjson.Field
		OfOpenAIDeveloperMessageContentArray respjson.Field
		raw                                  string
	} `json:"-"`
}

func (r *OpenAiv1ChatCompletionGetResponseInputMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAiv1ChatCompletionListResponse struct {
	Data    []OpenAiv1ChatCompletionListResponseData `json:"data,required"`
	FirstID string                                   `json:"first_id,required"`
	HasMore bool                                     `json:"has_more,required"`
	LastID  string                                   `json:"last_id,required"`
	Object  constant.List                            `json:"object,required"`
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
	// Any of nil, nil, nil, nil, nil.
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

func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsOpenAIUserMessageParam() (v OpenAIUserMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsOpenAISystemMessageParam() (v OpenAISystemMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsOpenAIAssistantMessageParam() (v OpenAIAssistantMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsOpenAIToolMessageParam() (v OpenAIToolMessageParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1ChatCompletionListResponseDataInputMessageUnion) AsOpenAIDeveloperMessageParam() (v OpenAIDeveloperMessageParamResp) {
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
// will be valid: OfString OfOpenAIUserMessageContentArray
// OfOpenAISystemMessageContentArray OfOpenAIAssistantMessageContentArray
// OfOpenAIToolMessageContentArray OfOpenAIDeveloperMessageContentArray]
type OpenAiv1ChatCompletionListResponseDataInputMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIUserMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIUserMessageContentArray []OpenAIUserMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAISystemMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAISystemMessageContentArray []OpenAISystemMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIAssistantMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIAssistantMessageContentArray []OpenAIAssistantMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIToolMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIToolMessageContentArray []OpenAIToolMessageParamContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIDeveloperMessageParamContentArrayItemUnionResp] instead of an object.
	OfOpenAIDeveloperMessageContentArray []OpenAIDeveloperMessageParamContentArrayItemUnionResp `json:",inline"`
	JSON                                 struct {
		OfString                             respjson.Field
		OfOpenAIUserMessageContentArray      respjson.Field
		OfOpenAISystemMessageContentArray    respjson.Field
		OfOpenAIAssistantMessageContentArray respjson.Field
		OfOpenAIToolMessageContentArray      respjson.Field
		OfOpenAIDeveloperMessageContentArray respjson.Field
		raw                                  string
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
	OfOpenAIUserMessage      *OpenAIUserMessageParam      `json:",omitzero,inline"`
	OfOpenAISystemMessage    *OpenAISystemMessageParam    `json:",omitzero,inline"`
	OfOpenAIAssistantMessage *OpenAIAssistantMessageParam `json:",omitzero,inline"`
	OfOpenAIToolMessage      *OpenAIToolMessageParam      `json:",omitzero,inline"`
	OfOpenAIDeveloperMessage *OpenAIDeveloperMessageParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsMessageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIUserMessage,
		u.OfOpenAISystemMessage,
		u.OfOpenAIAssistantMessage,
		u.OfOpenAIToolMessage,
		u.OfOpenAIDeveloperMessage)
}
func (u *OpenAIV1ChatCompletionNewParamsMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsMessageUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIUserMessage) {
		return u.OfOpenAIUserMessage
	} else if !param.IsOmitted(u.OfOpenAISystemMessage) {
		return u.OfOpenAISystemMessage
	} else if !param.IsOmitted(u.OfOpenAIAssistantMessage) {
		return u.OfOpenAIAssistantMessage
	} else if !param.IsOmitted(u.OfOpenAIToolMessage) {
		return u.OfOpenAIToolMessage
	} else if !param.IsOmitted(u.OfOpenAIDeveloperMessage) {
		return u.OfOpenAIDeveloperMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsMessageUnion) GetToolCalls() []ChatCompletionToolCallParam {
	if vt := u.OfOpenAIAssistantMessage; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsMessageUnion) GetToolCallID() *string {
	if vt := u.OfOpenAIToolMessage; vt != nil {
		return &vt.ToolCallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsMessageUnion) GetRole() *string {
	if vt := u.OfOpenAIUserMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfOpenAISystemMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfOpenAIAssistantMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfOpenAIToolMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfOpenAIDeveloperMessage; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsMessageUnion) GetName() *string {
	if vt := u.OfOpenAIUserMessage; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfOpenAISystemMessage; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfOpenAIAssistantMessage; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfOpenAIDeveloperMessage; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u OpenAIV1ChatCompletionNewParamsMessageUnion) GetContent() (res openAiv1ChatCompletionNewParamsMessageUnionContent) {
	if vt := u.OfOpenAIUserMessage; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfOpenAISystemMessage; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfOpenAIAssistantMessage; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfOpenAIToolMessage; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfOpenAIDeveloperMessage; vt != nil {
		res.any = vt.Content.asAny()
	}
	return
}

// Can have the runtime types [*string],
// [_[]OpenAIUserMessageParamContentArrayItemUnion],
// [_[]OpenAISystemMessageParamContentArrayItemUnion],
// [_[]OpenAIAssistantMessageParamContentArrayItemUnion],
// [_[]OpenAIToolMessageParamContentArrayItemUnion],
// [\*[]OpenAIDeveloperMessageParamContentArrayItemUnion]
type openAiv1ChatCompletionNewParamsMessageUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]llamastackclient.OpenAIUserMessageParamContentArrayItemUnion:
//	case *[]llamastackclient.OpenAISystemMessageParamContentArrayItemUnion:
//	case *[]llamastackclient.OpenAIAssistantMessageParamContentArrayItemUnion:
//	case *[]llamastackclient.OpenAIToolMessageParamContentArrayItemUnion:
//	case *[]llamastackclient.OpenAIDeveloperMessageParamContentArrayItemUnion:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u openAiv1ChatCompletionNewParamsMessageUnionContent) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[OpenAIV1ChatCompletionNewParamsMessageUnion](
		"role",
		apijson.Discriminator[OpenAIUserMessageParam](undefined),
		apijson.Discriminator[OpenAISystemMessageParam](undefined),
		apijson.Discriminator[OpenAIAssistantMessageParam](undefined),
		apijson.Discriminator[OpenAIToolMessageParam](undefined),
		apijson.Discriminator[OpenAIDeveloperMessageParam](undefined),
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
	OfOpenAIResponseFormatText       *OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatText       `json:",omitzero,inline"`
	OfOpenAIResponseFormatJsonSchema *OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchema `json:",omitzero,inline"`
	OfOpenAIResponseFormatJsonObject *OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonObject `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsResponseFormatUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIResponseFormatText, u.OfOpenAIResponseFormatJsonSchema, u.OfOpenAIResponseFormatJsonObject)
}
func (u *OpenAIV1ChatCompletionNewParamsResponseFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsResponseFormatUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIResponseFormatText) {
		return u.OfOpenAIResponseFormatText
	} else if !param.IsOmitted(u.OfOpenAIResponseFormatJsonSchema) {
		return u.OfOpenAIResponseFormatJsonSchema
	} else if !param.IsOmitted(u.OfOpenAIResponseFormatJsonObject) {
		return u.OfOpenAIResponseFormatJsonObject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsResponseFormatUnion) GetJsonSchema() *OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchema {
	if vt := u.OfOpenAIResponseFormatJsonSchema; vt != nil {
		return &vt.JsonSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ChatCompletionNewParamsResponseFormatUnion) GetType() *string {
	if vt := u.OfOpenAIResponseFormatText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseFormatJsonSchema; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseFormatJsonObject; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[OpenAIV1ChatCompletionNewParamsResponseFormatUnion](
		"type",
		apijson.Discriminator[OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatText](undefined),
		apijson.Discriminator[OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchema](undefined),
		apijson.Discriminator[OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonObject](undefined),
	)
}

// The property Type is required.
type OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatText struct {
	Type string `json:"type,required"`
	paramObj
}

func (r OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatText) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties JsonSchema, Type are required.
type OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchema struct {
	JsonSchema OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchema `json:"json_schema,omitzero,required"`
	Type       string                                                                                `json:"type,required"`
	paramObj
}

func (r OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchema struct {
	Name        string                                                                                                      `json:"name,required"`
	Description param.Opt[string]                                                                                           `json:"description,omitzero"`
	Strict      param.Opt[bool]                                                                                             `json:"strict,omitzero"`
	Schema      map[string]OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchemaSchemaUnion `json:"schema,omitzero"`
	paramObj
}

func (r OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchemaSchemaUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchemaSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchemaSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonSchemaJsonSchemaSchemaUnion) asAny() any {
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

// The property Type is required.
type OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonObject struct {
	Type string `json:"type,required"`
	paramObj
}

func (r OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonObject) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ChatCompletionNewParamsResponseFormatOpenAIResponseFormatJsonObject) UnmarshalJSON(data []byte) error {
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

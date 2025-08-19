// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// InferenceService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceService] method instead.
type InferenceService struct {
	Options []option.RequestOption
}

// NewInferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferenceService(opts ...option.RequestOption) (r InferenceService) {
	r = InferenceService{}
	r.Options = opts
	return
}

// Generate chat completions for a batch of messages using the specified model.
func (r *InferenceService) BatchChatCompletion(ctx context.Context, body InferenceBatchChatCompletionParams, opts ...option.RequestOption) (res *InferenceBatchChatCompletionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inference/batch-chat-completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate completions for a batch of content using the specified model.
func (r *InferenceService) BatchCompletion(ctx context.Context, body InferenceBatchCompletionParams, opts ...option.RequestOption) (res *InferenceBatchCompletionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inference/batch-completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate a chat completion for the given messages using the specified model.
func (r *InferenceService) ChatCompletion(ctx context.Context, body InferenceChatCompletionParams, opts ...option.RequestOption) (res *ChatCompletionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inference/chat-completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate a completion for the given content using the specified model.
func (r *InferenceService) Completion(ctx context.Context, body InferenceCompletionParams, opts ...option.RequestOption) (res *CompletionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inference/completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate embeddings for content pieces using the specified model.
func (r *InferenceService) Embeddings(ctx context.Context, body InferenceEmbeddingsParams, opts ...option.RequestOption) (res *InferenceEmbeddingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inference/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response from a chat completion request.
type ChatCompletionResponse struct {
	// The complete response message
	CompletionMessage CompletionMessage `json:"completion_message,required"`
	// Optional log probabilities for generated tokens
	Logprobs []TokenLogProbs    `json:"logprobs"`
	Metrics  []MetricInResponse `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionMessage respjson.Field
		Logprobs          respjson.Field
		Metrics           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in a chat conversation.
type CompletionMessage struct {
	// The content of the model's response
	Content InterleavedContentUnion `json:"content,required"`
	// Must be "assistant" to identify this as the model's response
	//
	// Any of "system", "user", "assistant", "tool".
	Role CompletionMessageRole `json:"role,required"`
	// Reason why the model stopped generating. Options are: -
	// `StopReason.end_of_turn`: The model finished generating the entire response. -
	// `StopReason.end_of_message`: The model finished generating but generated a
	// partial response -- usually, a tool call. The user may call the tool and
	// continue the conversation with the tool's response. -
	// `StopReason.out_of_tokens`: The model ran out of token budget.
	//
	// Any of "end_of_turn", "end_of_message", "out_of_tokens".
	StopReason CompletionMessageStopReason `json:"stop_reason,required"`
	// List of tool calls. Each tool call is a ToolCall object.
	ToolCalls []ToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		StopReason  respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionMessage) RawJSON() string { return r.JSON.raw }
func (r *CompletionMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CompletionMessage to a CompletionMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CompletionMessageParam.Overrides()
func (r CompletionMessage) ToParam() CompletionMessageParam {
	return param.Override[CompletionMessageParam](json.RawMessage(r.RawJSON()))
}

// Must be "assistant" to identify this as the model's response
type CompletionMessageRole string

const (
	CompletionMessageRoleSystem    CompletionMessageRole = "system"
	CompletionMessageRoleUser      CompletionMessageRole = "user"
	CompletionMessageRoleAssistant CompletionMessageRole = "assistant"
	CompletionMessageRoleTool      CompletionMessageRole = "tool"
)

// Reason why the model stopped generating. Options are: -
// `StopReason.end_of_turn`: The model finished generating the entire response. -
// `StopReason.end_of_message`: The model finished generating but generated a
// partial response -- usually, a tool call. The user may call the tool and
// continue the conversation with the tool's response. -
// `StopReason.out_of_tokens`: The model ran out of token budget.
type CompletionMessageStopReason string

const (
	CompletionMessageStopReasonEndOfTurn    CompletionMessageStopReason = "end_of_turn"
	CompletionMessageStopReasonEndOfMessage CompletionMessageStopReason = "end_of_message"
	CompletionMessageStopReasonOutOfTokens  CompletionMessageStopReason = "out_of_tokens"
)

// A message containing the model's (assistant) response in a chat conversation.
//
// The properties Content, Role, StopReason are required.
type CompletionMessageParam struct {
	// The content of the model's response
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Must be "assistant" to identify this as the model's response
	//
	// Any of "system", "user", "assistant", "tool".
	Role CompletionMessageRole `json:"role,omitzero,required"`
	// Reason why the model stopped generating. Options are: -
	// `StopReason.end_of_turn`: The model finished generating the entire response. -
	// `StopReason.end_of_message`: The model finished generating but generated a
	// partial response -- usually, a tool call. The user may call the tool and
	// continue the conversation with the tool's response. -
	// `StopReason.out_of_tokens`: The model ran out of token budget.
	//
	// Any of "end_of_turn", "end_of_message", "out_of_tokens".
	StopReason CompletionMessageStopReason `json:"stop_reason,omitzero,required"`
	// List of tool calls. Each tool call is a ToolCall object.
	ToolCalls []ToolCallParam `json:"tool_calls,omitzero"`
	paramObj
}

func (r CompletionMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow CompletionMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletionMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for grammar-guided response generation.
type GrammarResponseFormat struct {
	// The BNF grammar specification the response should conform to
	Bnf map[string]GrammarResponseFormatBnfUnion `json:"bnf,required"`
	// Must be "grammar" to identify this format type
	//
	// Any of "json_schema", "grammar".
	Type GrammarResponseFormatType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bnf         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GrammarResponseFormat) RawJSON() string { return r.JSON.raw }
func (r *GrammarResponseFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GrammarResponseFormat to a GrammarResponseFormatParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GrammarResponseFormatParam.Overrides()
func (r GrammarResponseFormat) ToParam() GrammarResponseFormatParam {
	return param.Override[GrammarResponseFormatParam](json.RawMessage(r.RawJSON()))
}

// GrammarResponseFormatBnfUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type GrammarResponseFormatBnfUnion struct {
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

func (u GrammarResponseFormatBnfUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GrammarResponseFormatBnfUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GrammarResponseFormatBnfUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GrammarResponseFormatBnfUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GrammarResponseFormatBnfUnion) RawJSON() string { return u.JSON.raw }

func (r *GrammarResponseFormatBnfUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be "grammar" to identify this format type
type GrammarResponseFormatType string

const (
	GrammarResponseFormatTypeJsonSchema GrammarResponseFormatType = "json_schema"
	GrammarResponseFormatTypeGrammar    GrammarResponseFormatType = "grammar"
)

// Configuration for grammar-guided response generation.
//
// The properties Bnf, Type are required.
type GrammarResponseFormatParam struct {
	// The BNF grammar specification the response should conform to
	Bnf map[string]GrammarResponseFormatBnfUnionParam `json:"bnf,omitzero,required"`
	// Must be "grammar" to identify this format type
	//
	// Any of "json_schema", "grammar".
	Type GrammarResponseFormatType `json:"type,omitzero,required"`
	paramObj
}

func (r GrammarResponseFormatParam) MarshalJSON() (data []byte, err error) {
	type shadow GrammarResponseFormatParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GrammarResponseFormatParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GrammarResponseFormatBnfUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u GrammarResponseFormatBnfUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *GrammarResponseFormatBnfUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GrammarResponseFormatBnfUnionParam) asAny() any {
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

type GreedySamplingStrategy struct {
	// The sampling method for text generation
	//
	// Any of "greedy", "top_p", "top_k".
	Type GreedySamplingStrategyType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GreedySamplingStrategy) RawJSON() string { return r.JSON.raw }
func (r *GreedySamplingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GreedySamplingStrategy to a GreedySamplingStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GreedySamplingStrategyParam.Overrides()
func (r GreedySamplingStrategy) ToParam() GreedySamplingStrategyParam {
	return param.Override[GreedySamplingStrategyParam](json.RawMessage(r.RawJSON()))
}

// The sampling method for text generation
type GreedySamplingStrategyType string

const (
	GreedySamplingStrategyTypeGreedy GreedySamplingStrategyType = "greedy"
	GreedySamplingStrategyTypeTopP   GreedySamplingStrategyType = "top_p"
	GreedySamplingStrategyTypeTopK   GreedySamplingStrategyType = "top_k"
)

// The property Type is required.
type GreedySamplingStrategyParam struct {
	// The sampling method for text generation
	//
	// Any of "greedy", "top_p", "top_k".
	Type GreedySamplingStrategyType `json:"type,omitzero,required"`
	paramObj
}

func (r GreedySamplingStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow GreedySamplingStrategyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GreedySamplingStrategyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InterleavedContentUnion contains all possible properties and values from
// [string], [InterleavedContentImageContentItem],
// [InterleavedContentTextContentItem], [[]InterleavedContentItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInterleavedContentItemArray]
type InterleavedContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]InterleavedContentItemUnion]
	// instead of an object.
	OfInterleavedContentItemArray []InterleavedContentItemUnion `json:",inline"`
	// This field is from variant [InterleavedContentImageContentItem].
	Image InterleavedContentImageContentItemImage `json:"image"`
	Type  string                                  `json:"type"`
	// This field is from variant [InterleavedContentTextContentItem].
	Text string `json:"text"`
	JSON struct {
		OfString                      respjson.Field
		OfInterleavedContentItemArray respjson.Field
		Image                         respjson.Field
		Type                          respjson.Field
		Text                          respjson.Field
		raw                           string
	} `json:"-"`
}

func (u InterleavedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentUnion) AsImageContentItem() (v InterleavedContentImageContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentUnion) AsTextContentItem() (v InterleavedContentTextContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentUnion) AsInterleavedContentItemArray() (v []InterleavedContentItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InterleavedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *InterleavedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InterleavedContentUnion to a InterleavedContentUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InterleavedContentUnionParam.Overrides()
func (r InterleavedContentUnion) ToParam() InterleavedContentUnionParam {
	return param.Override[InterleavedContentUnionParam](json.RawMessage(r.RawJSON()))
}

// A image content item
type InterleavedContentImageContentItem struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentImageContentItemImage `json:"image,required"`
	// Discriminator type of the content item. Always "image"
	//
	// Any of "text", "image", "tool_call".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentImageContentItem) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentImageContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type InterleavedContentImageContentItemImage struct {
	// base64 encoded image data as string
	Data string `json:"data"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL URL `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentImageContentItemImage) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentImageContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type InterleavedContentTextContentItem struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// Any of "text", "image", "tool_call".
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
func (r InterleavedContentTextContentItem) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func InterleavedContentParamOfImageContentItem(image InterleavedContentImageContentItemImageParam, type_ string) InterleavedContentUnionParam {
	var variant InterleavedContentImageContentItemParam
	variant.Image = image
	variant.Type = type_
	return InterleavedContentUnionParam{OfImageContentItem: &variant}
}

func InterleavedContentParamOfTextContentItem(text string, type_ string) InterleavedContentUnionParam {
	var variant InterleavedContentTextContentItemParam
	variant.Text = text
	variant.Type = type_
	return InterleavedContentUnionParam{OfTextContentItem: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InterleavedContentUnionParam struct {
	OfString                      param.Opt[string]                        `json:",omitzero,inline"`
	OfImageContentItem            *InterleavedContentImageContentItemParam `json:",omitzero,inline"`
	OfTextContentItem             *InterleavedContentTextContentItemParam  `json:",omitzero,inline"`
	OfInterleavedContentItemArray []InterleavedContentItemUnionParam       `json:",omitzero,inline"`
	paramUnion
}

func (u InterleavedContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfImageContentItem, u.OfTextContentItem, u.OfInterleavedContentItemArray)
}
func (u *InterleavedContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InterleavedContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfImageContentItem) {
		return u.OfImageContentItem
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	} else if !param.IsOmitted(u.OfInterleavedContentItemArray) {
		return &u.OfInterleavedContentItemArray
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentUnionParam) GetImage() *InterleavedContentImageContentItemImageParam {
	if vt := u.OfImageContentItem; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentUnionParam) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentUnionParam) GetType() *string {
	if vt := u.OfImageContentItem; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A image content item
//
// The properties Image, Type are required.
type InterleavedContentImageContentItemParam struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentImageContentItemImageParam `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// Any of "text", "image", "tool_call".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r InterleavedContentImageContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentImageContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentImageContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InterleavedContentImageContentItemParam](
		"type", "text", "image", "tool_call",
	)
}

// Image as a base64 encoded string or an URL
type InterleavedContentImageContentItemImageParam struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL URLParam `json:"url,omitzero"`
	paramObj
}

func (r InterleavedContentImageContentItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentImageContentItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentImageContentItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type InterleavedContentTextContentItemParam struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// Any of "text", "image", "tool_call".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r InterleavedContentTextContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentTextContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentTextContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InterleavedContentTextContentItemParam](
		"type", "text", "image", "tool_call",
	)
}

// InterleavedContentItemUnion contains all possible properties and values from
// [InterleavedContentItemImageContentItem],
// [InterleavedContentItemTextContentItem].
//
// Use the [InterleavedContentItemUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InterleavedContentItemUnion struct {
	// This field is from variant [InterleavedContentItemImageContentItem].
	Image InterleavedContentItemImageContentItemImage `json:"image"`
	// Any of nil, nil.
	Type string `json:"type"`
	// This field is from variant [InterleavedContentItemTextContentItem].
	Text string `json:"text"`
	JSON struct {
		Image respjson.Field
		Type  respjson.Field
		Text  respjson.Field
		raw   string
	} `json:"-"`
}

func (u InterleavedContentItemUnion) AsImageContentItem() (v InterleavedContentItemImageContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentItemUnion) AsTextContentItem() (v InterleavedContentItemTextContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InterleavedContentItemUnion) RawJSON() string { return u.JSON.raw }

func (r *InterleavedContentItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InterleavedContentItemUnion to a
// InterleavedContentItemUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InterleavedContentItemUnionParam.Overrides()
func (r InterleavedContentItemUnion) ToParam() InterleavedContentItemUnionParam {
	return param.Override[InterleavedContentItemUnionParam](json.RawMessage(r.RawJSON()))
}

// A image content item
type InterleavedContentItemImageContentItem struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentItemImageContentItemImage `json:"image,required"`
	// Discriminator type of the content item. Always "image"
	//
	// Any of "text", "image", "tool_call".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemImageContentItem) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemImageContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type InterleavedContentItemImageContentItemImage struct {
	// base64 encoded image data as string
	Data string `json:"data"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL URL `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemImageContentItemImage) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemImageContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type InterleavedContentItemTextContentItem struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// Any of "text", "image", "tool_call".
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
func (r InterleavedContentItemTextContentItem) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func InterleavedContentItemParamOfImageContentItem(image InterleavedContentItemImageContentItemImageParam, type_ string) InterleavedContentItemUnionParam {
	var variant InterleavedContentItemImageContentItemParam
	variant.Image = image
	variant.Type = type_
	return InterleavedContentItemUnionParam{OfImageContentItem: &variant}
}

func InterleavedContentItemParamOfTextContentItem(text string, type_ string) InterleavedContentItemUnionParam {
	var variant InterleavedContentItemTextContentItemParam
	variant.Text = text
	variant.Type = type_
	return InterleavedContentItemUnionParam{OfTextContentItem: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InterleavedContentItemUnionParam struct {
	OfImageContentItem *InterleavedContentItemImageContentItemParam `json:",omitzero,inline"`
	OfTextContentItem  *InterleavedContentItemTextContentItemParam  `json:",omitzero,inline"`
	paramUnion
}

func (u InterleavedContentItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfImageContentItem, u.OfTextContentItem)
}
func (u *InterleavedContentItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InterleavedContentItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfImageContentItem) {
		return u.OfImageContentItem
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetImage() *InterleavedContentItemImageContentItemImageParam {
	if vt := u.OfImageContentItem; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetType() *string {
	if vt := u.OfImageContentItem; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InterleavedContentItemUnionParam](
		"type",
		apijson.Discriminator[InterleavedContentItemImageContentItemParam]("text"),
		apijson.Discriminator[InterleavedContentItemImageContentItemParam]("image"),
		apijson.Discriminator[InterleavedContentItemImageContentItemParam]("tool_call"),
		apijson.Discriminator[InterleavedContentItemTextContentItemParam]("text"),
		apijson.Discriminator[InterleavedContentItemTextContentItemParam]("image"),
		apijson.Discriminator[InterleavedContentItemTextContentItemParam]("tool_call"),
	)
}

// A image content item
//
// The properties Image, Type are required.
type InterleavedContentItemImageContentItemParam struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentItemImageContentItemImageParam `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// Any of "text", "image", "tool_call".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r InterleavedContentItemImageContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemImageContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemImageContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InterleavedContentItemImageContentItemParam](
		"type", "text", "image", "tool_call",
	)
}

// Image as a base64 encoded string or an URL
type InterleavedContentItemImageContentItemImageParam struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL URLParam `json:"url,omitzero"`
	paramObj
}

func (r InterleavedContentItemImageContentItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemImageContentItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemImageContentItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type InterleavedContentItemTextContentItemParam struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// Any of "text", "image", "tool_call".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r InterleavedContentItemTextContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemTextContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemTextContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InterleavedContentItemTextContentItemParam](
		"type", "text", "image", "tool_call",
	)
}

// Configuration for JSON schema-guided response generation.
type JsonSchemaResponseFormat struct {
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model.
	JsonSchema map[string]JsonSchemaResponseFormatJsonSchemaUnion `json:"json_schema,required"`
	// Must be "json_schema" to identify this format type
	//
	// Any of "json_schema", "grammar".
	Type JsonSchemaResponseFormatType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JsonSchema  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JsonSchemaResponseFormat) RawJSON() string { return r.JSON.raw }
func (r *JsonSchemaResponseFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JsonSchemaResponseFormat to a
// JsonSchemaResponseFormatParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JsonSchemaResponseFormatParam.Overrides()
func (r JsonSchemaResponseFormat) ToParam() JsonSchemaResponseFormatParam {
	return param.Override[JsonSchemaResponseFormatParam](json.RawMessage(r.RawJSON()))
}

// JsonSchemaResponseFormatJsonSchemaUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type JsonSchemaResponseFormatJsonSchemaUnion struct {
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

func (u JsonSchemaResponseFormatJsonSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JsonSchemaResponseFormatJsonSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JsonSchemaResponseFormatJsonSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JsonSchemaResponseFormatJsonSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u JsonSchemaResponseFormatJsonSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *JsonSchemaResponseFormatJsonSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be "json_schema" to identify this format type
type JsonSchemaResponseFormatType string

const (
	JsonSchemaResponseFormatTypeJsonSchema JsonSchemaResponseFormatType = "json_schema"
	JsonSchemaResponseFormatTypeGrammar    JsonSchemaResponseFormatType = "grammar"
)

// Configuration for JSON schema-guided response generation.
//
// The properties JsonSchema, Type are required.
type JsonSchemaResponseFormatParam struct {
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model.
	JsonSchema map[string]JsonSchemaResponseFormatJsonSchemaUnionParam `json:"json_schema,omitzero,required"`
	// Must be "json_schema" to identify this format type
	//
	// Any of "json_schema", "grammar".
	Type JsonSchemaResponseFormatType `json:"type,omitzero,required"`
	paramObj
}

func (r JsonSchemaResponseFormatParam) MarshalJSON() (data []byte, err error) {
	type shadow JsonSchemaResponseFormatParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JsonSchemaResponseFormatParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type JsonSchemaResponseFormatJsonSchemaUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u JsonSchemaResponseFormatJsonSchemaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *JsonSchemaResponseFormatJsonSchemaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *JsonSchemaResponseFormatJsonSchemaUnionParam) asAny() any {
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

func MessageParamOfUserMessage[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T, role UserMessageRole) MessageUnionParam {
	var variant UserMessageParam
	switch v := any(content).(type) {
	case string:
		variant.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		variant.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		variant.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		variant.Content.OfInterleavedContentItemArray = v
	}
	variant.Role = role
	return MessageUnionParam{OfUserMessage: &variant}
}

func MessageParamOfSystemMessage[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T, role SystemMessageRole) MessageUnionParam {
	var variant SystemMessageParam
	switch v := any(content).(type) {
	case string:
		variant.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		variant.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		variant.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		variant.Content.OfInterleavedContentItemArray = v
	}
	variant.Role = role
	return MessageUnionParam{OfSystemMessage: &variant}
}

func MessageParamOfToolResponseMessage[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](callID string, content T, role ToolResponseMessageRole) MessageUnionParam {
	var variant ToolResponseMessageParam
	variant.CallID = callID
	switch v := any(content).(type) {
	case string:
		variant.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		variant.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		variant.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		variant.Content.OfInterleavedContentItemArray = v
	}
	variant.Role = role
	return MessageUnionParam{OfToolResponseMessage: &variant}
}

func MessageParamOfCompletionMessage[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T, role CompletionMessageRole, stopReason CompletionMessageStopReason) MessageUnionParam {
	var variant CompletionMessageParam
	switch v := any(content).(type) {
	case string:
		variant.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		variant.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		variant.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		variant.Content.OfInterleavedContentItemArray = v
	}
	variant.Role = role
	variant.StopReason = stopReason
	return MessageUnionParam{OfCompletionMessage: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageUnionParam struct {
	OfUserMessage         *UserMessageParam         `json:",omitzero,inline"`
	OfSystemMessage       *SystemMessageParam       `json:",omitzero,inline"`
	OfToolResponseMessage *ToolResponseMessageParam `json:",omitzero,inline"`
	OfCompletionMessage   *CompletionMessageParam   `json:",omitzero,inline"`
	paramUnion
}

func (u MessageUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUserMessage, u.OfSystemMessage, u.OfToolResponseMessage, u.OfCompletionMessage)
}
func (u *MessageUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageUnionParam) asAny() any {
	if !param.IsOmitted(u.OfUserMessage) {
		return u.OfUserMessage
	} else if !param.IsOmitted(u.OfSystemMessage) {
		return u.OfSystemMessage
	} else if !param.IsOmitted(u.OfToolResponseMessage) {
		return u.OfToolResponseMessage
	} else if !param.IsOmitted(u.OfCompletionMessage) {
		return u.OfCompletionMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetContext() *InterleavedContentUnionParam {
	if vt := u.OfUserMessage; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetCallID() *string {
	if vt := u.OfToolResponseMessage; vt != nil {
		return &vt.CallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetStopReason() *string {
	if vt := u.OfCompletionMessage; vt != nil {
		return (*string)(&vt.StopReason)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetToolCalls() []ToolCallParam {
	if vt := u.OfCompletionMessage; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetRole() *string {
	if vt := u.OfUserMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfSystemMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfToolResponseMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfCompletionMessage; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's Content property, if present.
func (u MessageUnionParam) GetContent() *InterleavedContentUnionParam {
	if vt := u.OfUserMessage; vt != nil {
		return &vt.Content
	} else if vt := u.OfSystemMessage; vt != nil {
		return &vt.Content
	} else if vt := u.OfToolResponseMessage; vt != nil {
		return &vt.Content
	} else if vt := u.OfCompletionMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

func init() {
	apijson.RegisterUnion[MessageUnionParam](
		"role",
		apijson.Discriminator[UserMessageParam]("system"),
		apijson.Discriminator[UserMessageParam]("user"),
		apijson.Discriminator[UserMessageParam]("assistant"),
		apijson.Discriminator[UserMessageParam]("tool"),
		apijson.Discriminator[SystemMessageParam]("system"),
		apijson.Discriminator[SystemMessageParam]("user"),
		apijson.Discriminator[SystemMessageParam]("assistant"),
		apijson.Discriminator[SystemMessageParam]("tool"),
		apijson.Discriminator[ToolResponseMessageParam]("system"),
		apijson.Discriminator[ToolResponseMessageParam]("user"),
		apijson.Discriminator[ToolResponseMessageParam]("assistant"),
		apijson.Discriminator[ToolResponseMessageParam]("tool"),
		apijson.Discriminator[CompletionMessageParam]("system"),
		apijson.Discriminator[CompletionMessageParam]("user"),
		apijson.Discriminator[CompletionMessageParam]("assistant"),
		apijson.Discriminator[CompletionMessageParam]("tool"),
	)
}

type MetricInResponse struct {
	Metric string  `json:"metric,required"`
	Value  float64 `json:"value,required"`
	Unit   string  `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metric      respjson.Field
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricInResponse) RawJSON() string { return r.JSON.raw }
func (r *MetricInResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseFormatUnion contains all possible properties and values from
// [JsonSchemaResponseFormat], [GrammarResponseFormat].
//
// Use the [ResponseFormatUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseFormatUnion struct {
	// This field is from variant [JsonSchemaResponseFormat].
	JsonSchema map[string]JsonSchemaResponseFormatJsonSchemaUnion `json:"json_schema"`
	// Any of nil, nil.
	Type string `json:"type"`
	// This field is from variant [GrammarResponseFormat].
	Bnf  map[string]GrammarResponseFormatBnfUnion `json:"bnf"`
	JSON struct {
		JsonSchema respjson.Field
		Type       respjson.Field
		Bnf        respjson.Field
		raw        string
	} `json:"-"`
}

func (u ResponseFormatUnion) AsJsonSchemaResponseFormat() (v JsonSchemaResponseFormat) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatUnion) AsGrammarResponseFormat() (v GrammarResponseFormat) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseFormatUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ResponseFormatUnion to a ResponseFormatUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ResponseFormatUnionParam.Overrides()
func (r ResponseFormatUnion) ToParam() ResponseFormatUnionParam {
	return param.Override[ResponseFormatUnionParam](json.RawMessage(r.RawJSON()))
}

func ResponseFormatParamOfJsonSchemaResponseFormat(jsonSchema map[string]JsonSchemaResponseFormatJsonSchemaUnionParam, type_ JsonSchemaResponseFormatType) ResponseFormatUnionParam {
	var variant JsonSchemaResponseFormatParam
	variant.JsonSchema = jsonSchema
	variant.Type = type_
	return ResponseFormatUnionParam{OfJsonSchemaResponseFormat: &variant}
}

func ResponseFormatParamOfGrammarResponseFormat(bnf map[string]GrammarResponseFormatBnfUnionParam, type_ GrammarResponseFormatType) ResponseFormatUnionParam {
	var variant GrammarResponseFormatParam
	variant.Bnf = bnf
	variant.Type = type_
	return ResponseFormatUnionParam{OfGrammarResponseFormat: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseFormatUnionParam struct {
	OfJsonSchemaResponseFormat *JsonSchemaResponseFormatParam `json:",omitzero,inline"`
	OfGrammarResponseFormat    *GrammarResponseFormatParam    `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseFormatUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfJsonSchemaResponseFormat, u.OfGrammarResponseFormat)
}
func (u *ResponseFormatUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseFormatUnionParam) asAny() any {
	if !param.IsOmitted(u.OfJsonSchemaResponseFormat) {
		return u.OfJsonSchemaResponseFormat
	} else if !param.IsOmitted(u.OfGrammarResponseFormat) {
		return u.OfGrammarResponseFormat
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseFormatUnionParam) GetJsonSchema() map[string]JsonSchemaResponseFormatJsonSchemaUnionParam {
	if vt := u.OfJsonSchemaResponseFormat; vt != nil {
		return vt.JsonSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseFormatUnionParam) GetBnf() map[string]GrammarResponseFormatBnfUnionParam {
	if vt := u.OfGrammarResponseFormat; vt != nil {
		return vt.Bnf
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseFormatUnionParam) GetType() *string {
	if vt := u.OfJsonSchemaResponseFormat; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGrammarResponseFormat; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseFormatUnionParam](
		"type",
		apijson.Discriminator[JsonSchemaResponseFormatParam]("json_schema"),
		apijson.Discriminator[JsonSchemaResponseFormatParam]("grammar"),
		apijson.Discriminator[GrammarResponseFormatParam]("json_schema"),
		apijson.Discriminator[GrammarResponseFormatParam]("grammar"),
	)
}

// Sampling parameters.
type SamplingParamsResp struct {
	// The sampling strategy.
	Strategy SamplingParamsStrategyUnionResp `json:"strategy,required"`
	// The maximum number of tokens that can be generated in the completion. The token
	// count of your prompt plus max_tokens cannot exceed the model's context length.
	MaxTokens int64 `json:"max_tokens"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on
	// whether they appear in the text so far, increasing the model's likelihood to
	// talk about new topics.
	RepetitionPenalty float64 `json:"repetition_penalty"`
	// Up to 4 sequences where the API will stop generating further tokens. The
	// returned text will not contain the stop sequence.
	Stop []string `json:"stop"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Strategy          respjson.Field
		MaxTokens         respjson.Field
		RepetitionPenalty respjson.Field
		Stop              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SamplingParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SamplingParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SamplingParamsResp to a SamplingParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SamplingParams.Overrides()
func (r SamplingParamsResp) ToParam() SamplingParams {
	return param.Override[SamplingParams](json.RawMessage(r.RawJSON()))
}

// SamplingParamsStrategyUnionResp contains all possible properties and values from
// [GreedySamplingStrategy], [TopPSamplingStrategy], [TopKSamplingStrategy].
//
// Use the [SamplingParamsStrategyUnionResp.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SamplingParamsStrategyUnionResp struct {
	// Any of nil, nil, nil.
	Type string `json:"type"`
	// This field is from variant [TopPSamplingStrategy].
	Temperature float64 `json:"temperature"`
	// This field is from variant [TopPSamplingStrategy].
	TopP float64 `json:"top_p"`
	// This field is from variant [TopKSamplingStrategy].
	TopK int64 `json:"top_k"`
	JSON struct {
		Type        respjson.Field
		Temperature respjson.Field
		TopP        respjson.Field
		TopK        respjson.Field
		raw         string
	} `json:"-"`
}

func (u SamplingParamsStrategyUnionResp) AsGreedySamplingStrategy() (v GreedySamplingStrategy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SamplingParamsStrategyUnionResp) AsTopPSamplingStrategy() (v TopPSamplingStrategy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SamplingParamsStrategyUnionResp) AsTopKSamplingStrategy() (v TopKSamplingStrategy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SamplingParamsStrategyUnionResp) RawJSON() string { return u.JSON.raw }

func (r *SamplingParamsStrategyUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sampling parameters.
//
// The property Strategy is required.
type SamplingParams struct {
	// The sampling strategy.
	Strategy SamplingParamsStrategyUnion `json:"strategy,omitzero,required"`
	// The maximum number of tokens that can be generated in the completion. The token
	// count of your prompt plus max_tokens cannot exceed the model's context length.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on
	// whether they appear in the text so far, increasing the model's likelihood to
	// talk about new topics.
	RepetitionPenalty param.Opt[float64] `json:"repetition_penalty,omitzero"`
	// Up to 4 sequences where the API will stop generating further tokens. The
	// returned text will not contain the stop sequence.
	Stop []string `json:"stop,omitzero"`
	paramObj
}

func (r SamplingParams) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SamplingParamsStrategyUnion struct {
	OfGreedySamplingStrategy *GreedySamplingStrategyParam `json:",omitzero,inline"`
	OfTopPSamplingStrategy   *TopPSamplingStrategyParam   `json:",omitzero,inline"`
	OfTopKSamplingStrategy   *TopKSamplingStrategyParam   `json:",omitzero,inline"`
	paramUnion
}

func (u SamplingParamsStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGreedySamplingStrategy, u.OfTopPSamplingStrategy, u.OfTopKSamplingStrategy)
}
func (u *SamplingParamsStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SamplingParamsStrategyUnion) asAny() any {
	if !param.IsOmitted(u.OfGreedySamplingStrategy) {
		return u.OfGreedySamplingStrategy
	} else if !param.IsOmitted(u.OfTopPSamplingStrategy) {
		return u.OfTopPSamplingStrategy
	} else if !param.IsOmitted(u.OfTopKSamplingStrategy) {
		return u.OfTopKSamplingStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTemperature() *float64 {
	if vt := u.OfTopPSamplingStrategy; vt != nil && vt.Temperature.Valid() {
		return &vt.Temperature.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTopP() *float64 {
	if vt := u.OfTopPSamplingStrategy; vt != nil && vt.TopP.Valid() {
		return &vt.TopP.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTopK() *int64 {
	if vt := u.OfTopKSamplingStrategy; vt != nil {
		return &vt.TopK
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetType() *string {
	if vt := u.OfGreedySamplingStrategy; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTopPSamplingStrategy; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTopKSamplingStrategy; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SamplingParamsStrategyUnion](
		"type",
		apijson.Discriminator[GreedySamplingStrategyParam]("greedy"),
		apijson.Discriminator[GreedySamplingStrategyParam]("top_p"),
		apijson.Discriminator[GreedySamplingStrategyParam]("top_k"),
		apijson.Discriminator[TopPSamplingStrategyParam]("greedy"),
		apijson.Discriminator[TopPSamplingStrategyParam]("top_p"),
		apijson.Discriminator[TopPSamplingStrategyParam]("top_k"),
		apijson.Discriminator[TopKSamplingStrategyParam]("greedy"),
		apijson.Discriminator[TopKSamplingStrategyParam]("top_p"),
		apijson.Discriminator[TopKSamplingStrategyParam]("top_k"),
	)
}

// A system message providing instructions or context to the model.
//
// The properties Content, Role are required.
type SystemMessageParam struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Must be "system" to identify this as a system message
	//
	// Any of "system", "user", "assistant", "tool".
	Role SystemMessageRole `json:"role,omitzero,required"`
	paramObj
}

func (r SystemMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be "system" to identify this as a system message
type SystemMessageRole string

const (
	SystemMessageRoleSystem    SystemMessageRole = "system"
	SystemMessageRoleUser      SystemMessageRole = "user"
	SystemMessageRoleAssistant SystemMessageRole = "assistant"
	SystemMessageRoleTool      SystemMessageRole = "tool"
)

type ToolCall struct {
	Arguments     ToolCallArgumentsUnion `json:"arguments,required"`
	CallID        string                 `json:"call_id,required"`
	ToolName      ToolCallToolName       `json:"tool_name,required"`
	ArgumentsJson string                 `json:"arguments_json"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments     respjson.Field
		CallID        respjson.Field
		ToolName      respjson.Field
		ArgumentsJson respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolCall) RawJSON() string { return r.JSON.raw }
func (r *ToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolCall to a ToolCallParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolCallParam.Overrides()
func (r ToolCall) ToParam() ToolCallParam {
	return param.Override[ToolCallParam](json.RawMessage(r.RawJSON()))
}

// ToolCallArgumentsUnion contains all possible properties and values from
// [string], [map[string]ToolCallArgumentsMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfToolCallArgumentsMapItemArray]
type ToolCallArgumentsUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ToolCallArgumentsMapItemArrayItemUnion] instead of an object.
	OfToolCallArgumentsMapItemArray []ToolCallArgumentsMapItemArrayItemUnion `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfFloat                         respjson.Field
		OfBool                          respjson.Field
		OfToolCallArgumentsMapItemArray respjson.Field
		raw                             string
	} `json:"-"`
}

func (u ToolCallArgumentsUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsUnion) AsToolCallArgumentsMapMap() (v map[string]ToolCallArgumentsMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolCallArgumentsUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolCallArgumentsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolCallArgumentsMapItemUnion contains all possible properties and values from
// [string], [float64], [bool], [[]ToolCallArgumentsMapItemArrayItemUnion],
// [map[string]ToolCallArgumentsMapItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfToolCallArgumentsMapItemArray]
type ToolCallArgumentsMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ToolCallArgumentsMapItemArrayItemUnion] instead of an object.
	OfToolCallArgumentsMapItemArray []ToolCallArgumentsMapItemArrayItemUnion `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfFloat                         respjson.Field
		OfBool                          respjson.Field
		OfToolCallArgumentsMapItemArray respjson.Field
		raw                             string
	} `json:"-"`
}

func (u ToolCallArgumentsMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemUnion) AsToolCallArgumentsMapItemArray() (v []ToolCallArgumentsMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemUnion) AsToolCallArgumentsMapItemMapMap() (v map[string]ToolCallArgumentsMapItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolCallArgumentsMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolCallArgumentsMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolCallArgumentsMapItemArrayItemUnion contains all possible properties and
// values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ToolCallArgumentsMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ToolCallArgumentsMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolCallArgumentsMapItemArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolCallArgumentsMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolCallArgumentsMapItemMapItemUnion contains all possible properties and values
// from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ToolCallArgumentsMapItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ToolCallArgumentsMapItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolCallArgumentsMapItemMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolCallArgumentsMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolCallToolName string

const (
	ToolCallToolNameBraveSearch     ToolCallToolName = "brave_search"
	ToolCallToolNameWolframAlpha    ToolCallToolName = "wolfram_alpha"
	ToolCallToolNamePhotogen        ToolCallToolName = "photogen"
	ToolCallToolNameCodeInterpreter ToolCallToolName = "code_interpreter"
)

// The properties Arguments, CallID, ToolName are required.
type ToolCallParam struct {
	Arguments     ToolCallArgumentsUnionParam `json:"arguments,omitzero,required"`
	CallID        string                      `json:"call_id,required"`
	ToolName      ToolCallToolName            `json:"tool_name,omitzero,required"`
	ArgumentsJson param.Opt[string]           `json:"arguments_json,omitzero"`
	paramObj
}

func (r ToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolCallArgumentsUnionParam struct {
	OfString                  param.Opt[string]                             `json:",omitzero,inline"`
	OfToolCallArgumentsMapMap map[string]ToolCallArgumentsMapItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ToolCallArgumentsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfToolCallArgumentsMapMap)
}
func (u *ToolCallArgumentsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolCallArgumentsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfToolCallArgumentsMapMap) {
		return &u.OfToolCallArgumentsMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolCallArgumentsMapItemUnionParam struct {
	OfString                         param.Opt[string]                                    `json:",omitzero,inline"`
	OfFloat                          param.Opt[float64]                                   `json:",omitzero,inline"`
	OfBool                           param.Opt[bool]                                      `json:",omitzero,inline"`
	OfToolCallArgumentsMapItemArray  []ToolCallArgumentsMapItemArrayItemUnionParam        `json:",omitzero,inline"`
	OfToolCallArgumentsMapItemMapMap map[string]ToolCallArgumentsMapItemMapItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ToolCallArgumentsMapItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfToolCallArgumentsMapItemArray,
		u.OfToolCallArgumentsMapItemMapMap)
}
func (u *ToolCallArgumentsMapItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolCallArgumentsMapItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfToolCallArgumentsMapItemArray) {
		return &u.OfToolCallArgumentsMapItemArray
	} else if !param.IsOmitted(u.OfToolCallArgumentsMapItemMapMap) {
		return &u.OfToolCallArgumentsMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolCallArgumentsMapItemArrayItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ToolCallArgumentsMapItemArrayItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ToolCallArgumentsMapItemArrayItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolCallArgumentsMapItemArrayItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolCallArgumentsMapItemMapItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ToolCallArgumentsMapItemMapItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ToolCallArgumentsMapItemMapItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolCallArgumentsMapItemMapItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Configuration for tool use.
type ToolConfig struct {
	// (Optional) Config for how to override the default system prompt. -
	// `SystemMessageBehavior.append`: Appends the provided system message to the
	// default system prompt. - `SystemMessageBehavior.replace`: Replaces the default
	// system prompt with the provided system message. The system message can include
	// the string '{{function_definitions}}' to indicate where the function definitions
	// should be inserted.
	//
	// Any of "append", "replace".
	SystemMessageBehavior ToolConfigSystemMessageBehavior `json:"system_message_behavior"`
	// (Optional) Whether tool use is automatic, required, or none. Can also specify a
	// tool name to use a specific tool. Defaults to ToolChoice.auto.
	ToolChoice ToolConfigToolChoice `json:"tool_choice"`
	// (Optional) Instructs the model how to format tool calls. By default, Llama Stack
	// will attempt to use a format that is best adapted to the model. -
	// `ToolPromptFormat.json`: The tool calls are formatted as a JSON object. -
	// `ToolPromptFormat.function_tag`: The tool calls are enclosed in a
	// <function=function_name> tag. - `ToolPromptFormat.python_list`: The tool calls
	// are output as Python syntax -- a list of function calls.
	//
	// Any of "json", "function_tag", "python_list".
	ToolPromptFormat ToolConfigToolPromptFormat `json:"tool_prompt_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SystemMessageBehavior respjson.Field
		ToolChoice            respjson.Field
		ToolPromptFormat      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolConfig) RawJSON() string { return r.JSON.raw }
func (r *ToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolConfig to a ToolConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolConfigParam.Overrides()
func (r ToolConfig) ToParam() ToolConfigParam {
	return param.Override[ToolConfigParam](json.RawMessage(r.RawJSON()))
}

// (Optional) Config for how to override the default system prompt. -
// `SystemMessageBehavior.append`: Appends the provided system message to the
// default system prompt. - `SystemMessageBehavior.replace`: Replaces the default
// system prompt with the provided system message. The system message can include
// the string '{{function_definitions}}' to indicate where the function definitions
// should be inserted.
type ToolConfigSystemMessageBehavior string

const (
	ToolConfigSystemMessageBehaviorAppend  ToolConfigSystemMessageBehavior = "append"
	ToolConfigSystemMessageBehaviorReplace ToolConfigSystemMessageBehavior = "replace"
)

// Whether tool use is required or automatic. This is a hint to the model which may
// not be followed. It depends on the Instruction Following capabilities of the
// model.
type ToolConfigToolChoice string

const (
	ToolConfigToolChoiceAuto     ToolConfigToolChoice = "auto"
	ToolConfigToolChoiceRequired ToolConfigToolChoice = "required"
	ToolConfigToolChoiceNone     ToolConfigToolChoice = "none"
)

// (Optional) Instructs the model how to format tool calls. By default, Llama Stack
// will attempt to use a format that is best adapted to the model. -
// `ToolPromptFormat.json`: The tool calls are formatted as a JSON object. -
// `ToolPromptFormat.function_tag`: The tool calls are enclosed in a
// <function=function_name> tag. - `ToolPromptFormat.python_list`: The tool calls
// are output as Python syntax -- a list of function calls.
type ToolConfigToolPromptFormat string

const (
	ToolConfigToolPromptFormatJson        ToolConfigToolPromptFormat = "json"
	ToolConfigToolPromptFormatFunctionTag ToolConfigToolPromptFormat = "function_tag"
	ToolConfigToolPromptFormatPythonList  ToolConfigToolPromptFormat = "python_list"
)

// Configuration for tool use.
type ToolConfigParam struct {
	// (Optional) Config for how to override the default system prompt. -
	// `SystemMessageBehavior.append`: Appends the provided system message to the
	// default system prompt. - `SystemMessageBehavior.replace`: Replaces the default
	// system prompt with the provided system message. The system message can include
	// the string '{{function_definitions}}' to indicate where the function definitions
	// should be inserted.
	//
	// Any of "append", "replace".
	SystemMessageBehavior ToolConfigSystemMessageBehavior `json:"system_message_behavior,omitzero"`
	// (Optional) Whether tool use is automatic, required, or none. Can also specify a
	// tool name to use a specific tool. Defaults to ToolChoice.auto.
	ToolChoice ToolConfigToolChoice `json:"tool_choice,omitzero"`
	// (Optional) Instructs the model how to format tool calls. By default, Llama Stack
	// will attempt to use a format that is best adapted to the model. -
	// `ToolPromptFormat.json`: The tool calls are formatted as a JSON object. -
	// `ToolPromptFormat.function_tag`: The tool calls are enclosed in a
	// <function=function_name> tag. - `ToolPromptFormat.python_list`: The tool calls
	// are output as Python syntax -- a list of function calls.
	//
	// Any of "json", "function_tag", "python_list".
	ToolPromptFormat ToolConfigToolPromptFormat `json:"tool_prompt_format,omitzero"`
	paramObj
}

func (r ToolConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ToolName is required.
type ToolDefinitionParam struct {
	ToolName    ToolDefinitionToolName                  `json:"tool_name,omitzero,required"`
	Description param.Opt[string]                       `json:"description,omitzero"`
	Parameters  map[string]ToolDefinitionParameterParam `json:"parameters,omitzero"`
	paramObj
}

func (r ToolDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDefinitionToolName string

const (
	ToolDefinitionToolNameBraveSearch     ToolDefinitionToolName = "brave_search"
	ToolDefinitionToolNameWolframAlpha    ToolDefinitionToolName = "wolfram_alpha"
	ToolDefinitionToolNamePhotogen        ToolDefinitionToolName = "photogen"
	ToolDefinitionToolNameCodeInterpreter ToolDefinitionToolName = "code_interpreter"
)

// The property ParamType is required.
type ToolDefinitionParameterParam struct {
	ParamType   string                                   `json:"param_type,required"`
	Description param.Opt[string]                        `json:"description,omitzero"`
	Required    param.Opt[bool]                          `json:"required,omitzero"`
	Default     ToolDefinitionParameterDefaultUnionParam `json:"default,omitzero"`
	paramObj
}

func (r ToolDefinitionParameterParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolDefinitionParameterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolDefinitionParameterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolDefinitionParameterDefaultUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ToolDefinitionParameterDefaultUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ToolDefinitionParameterDefaultUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolDefinitionParameterDefaultUnionParam) asAny() any {
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

type TopKSamplingStrategy struct {
	TopK int64 `json:"top_k,required"`
	// The sampling method for text generation
	//
	// Any of "greedy", "top_p", "top_k".
	Type TopKSamplingStrategyType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TopK        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TopKSamplingStrategy) RawJSON() string { return r.JSON.raw }
func (r *TopKSamplingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TopKSamplingStrategy to a TopKSamplingStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TopKSamplingStrategyParam.Overrides()
func (r TopKSamplingStrategy) ToParam() TopKSamplingStrategyParam {
	return param.Override[TopKSamplingStrategyParam](json.RawMessage(r.RawJSON()))
}

// The sampling method for text generation
type TopKSamplingStrategyType string

const (
	TopKSamplingStrategyTypeGreedy TopKSamplingStrategyType = "greedy"
	TopKSamplingStrategyTypeTopP   TopKSamplingStrategyType = "top_p"
	TopKSamplingStrategyTypeTopK   TopKSamplingStrategyType = "top_k"
)

// The properties TopK, Type are required.
type TopKSamplingStrategyParam struct {
	TopK int64 `json:"top_k,required"`
	// The sampling method for text generation
	//
	// Any of "greedy", "top_p", "top_k".
	Type TopKSamplingStrategyType `json:"type,omitzero,required"`
	paramObj
}

func (r TopKSamplingStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow TopKSamplingStrategyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TopKSamplingStrategyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TopPSamplingStrategy struct {
	// The sampling method for text generation
	//
	// Any of "greedy", "top_p", "top_k".
	Type        TopPSamplingStrategyType `json:"type,required"`
	Temperature float64                  `json:"temperature"`
	TopP        float64                  `json:"top_p"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Temperature respjson.Field
		TopP        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TopPSamplingStrategy) RawJSON() string { return r.JSON.raw }
func (r *TopPSamplingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TopPSamplingStrategy to a TopPSamplingStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TopPSamplingStrategyParam.Overrides()
func (r TopPSamplingStrategy) ToParam() TopPSamplingStrategyParam {
	return param.Override[TopPSamplingStrategyParam](json.RawMessage(r.RawJSON()))
}

// The sampling method for text generation
type TopPSamplingStrategyType string

const (
	TopPSamplingStrategyTypeGreedy TopPSamplingStrategyType = "greedy"
	TopPSamplingStrategyTypeTopP   TopPSamplingStrategyType = "top_p"
	TopPSamplingStrategyTypeTopK   TopPSamplingStrategyType = "top_k"
)

// The property Type is required.
type TopPSamplingStrategyParam struct {
	// The sampling method for text generation
	//
	// Any of "greedy", "top_p", "top_k".
	Type        TopPSamplingStrategyType `json:"type,omitzero,required"`
	Temperature param.Opt[float64]       `json:"temperature,omitzero"`
	TopP        param.Opt[float64]       `json:"top_p,omitzero"`
	paramObj
}

func (r TopPSamplingStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow TopPSamplingStrategyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TopPSamplingStrategyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceBatchChatCompletionResponse struct {
	Batch []ChatCompletionResponse `json:"batch,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Batch       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceBatchChatCompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *InferenceBatchChatCompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceBatchCompletionResponse struct {
	Batch []CompletionResponse `json:"batch,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Batch       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceBatchCompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *InferenceBatchCompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing generated embeddings.
type InferenceEmbeddingsResponse struct {
	// List of embedding vectors, one per input content. Each embedding is a list of
	// floats. The dimensionality of the embedding is model-specific; you can check
	// model metadata using /models/{model_id}
	Embeddings [][]float64 `json:"embeddings,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embeddings  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingsResponse) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceBatchChatCompletionParams struct {
	// The messages to generate completions for.
	MessagesBatch [][]MessageUnionParam `json:"messages_batch,omitzero,required"`
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	ModelID string `json:"model_id,required"`
	// (Optional) If specified, log probabilities for each token position will be
	// returned.
	Logprobs InferenceBatchChatCompletionParamsLogprobs `json:"logprobs,omitzero"`
	// (Optional) Grammar specification for guided (structured) decoding.
	ResponseFormat ResponseFormatUnionParam `json:"response_format,omitzero"`
	// (Optional) Parameters to control the sampling strategy.
	SamplingParams SamplingParams `json:"sampling_params,omitzero"`
	// (Optional) Configuration for tool use.
	ToolConfig ToolConfigParam `json:"tool_config,omitzero"`
	// (Optional) List of tool definitions available to the model.
	Tools []ToolDefinitionParam `json:"tools,omitzero"`
	paramObj
}

func (r InferenceBatchChatCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchChatCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchChatCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) If specified, log probabilities for each token position will be
// returned.
type InferenceBatchChatCompletionParamsLogprobs struct {
	// How many tokens (for each position) to return log probabilities for.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	paramObj
}

func (r InferenceBatchChatCompletionParamsLogprobs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchChatCompletionParamsLogprobs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchChatCompletionParamsLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceBatchCompletionParams struct {
	// The content to generate completions for.
	ContentBatch []InterleavedContentUnionParam `json:"content_batch,omitzero,required"`
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	ModelID string `json:"model_id,required"`
	// (Optional) If specified, log probabilities for each token position will be
	// returned.
	Logprobs InferenceBatchCompletionParamsLogprobs `json:"logprobs,omitzero"`
	// (Optional) Grammar specification for guided (structured) decoding.
	ResponseFormat ResponseFormatUnionParam `json:"response_format,omitzero"`
	// (Optional) Parameters to control the sampling strategy.
	SamplingParams SamplingParams `json:"sampling_params,omitzero"`
	paramObj
}

func (r InferenceBatchCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) If specified, log probabilities for each token position will be
// returned.
type InferenceBatchCompletionParamsLogprobs struct {
	// How many tokens (for each position) to return log probabilities for.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	paramObj
}

func (r InferenceBatchCompletionParamsLogprobs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchCompletionParamsLogprobs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchCompletionParamsLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceChatCompletionParams struct {
	// List of messages in the conversation.
	Messages []MessageUnionParam `json:"messages,omitzero,required"`
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	ModelID string `json:"model_id,required"`
	// (Optional) If True, generate an SSE event stream of the response. Defaults to
	// False.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// (Optional) If specified, log probabilities for each token position will be
	// returned.
	Logprobs InferenceChatCompletionParamsLogprobs `json:"logprobs,omitzero"`
	// (Optional) Grammar specification for guided (structured) decoding. There are two
	// options: - `ResponseFormat.json_schema`: The grammar is a JSON schema. Most
	// providers support this format. - `ResponseFormat.grammar`: The grammar is a BNF
	// grammar. This format is more flexible, but not all providers support it.
	ResponseFormat ResponseFormatUnionParam `json:"response_format,omitzero"`
	// Parameters to control the sampling strategy.
	SamplingParams SamplingParams `json:"sampling_params,omitzero"`
	// (Optional) Whether tool use is required or automatic. Defaults to
	// ToolChoice.auto. .. deprecated:: Use tool_config instead.
	//
	// Any of "auto", "required", "none".
	ToolChoice InferenceChatCompletionParamsToolChoice `json:"tool_choice,omitzero"`
	// (Optional) Configuration for tool use.
	ToolConfig ToolConfigParam `json:"tool_config,omitzero"`
	// (Optional) Instructs the model how to format tool calls. By default, Llama Stack
	// will attempt to use a format that is best adapted to the model. -
	// `ToolPromptFormat.json`: The tool calls are formatted as a JSON object. -
	// `ToolPromptFormat.function_tag`: The tool calls are enclosed in a
	// <function=function_name> tag. - `ToolPromptFormat.python_list`: The tool calls
	// are output as Python syntax -- a list of function calls. .. deprecated:: Use
	// tool_config instead.
	//
	// Any of "json", "function_tag", "python_list".
	ToolPromptFormat InferenceChatCompletionParamsToolPromptFormat `json:"tool_prompt_format,omitzero"`
	// (Optional) List of tool definitions available to the model.
	Tools []ToolDefinitionParam `json:"tools,omitzero"`
	paramObj
}

func (r InferenceChatCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceChatCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceChatCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) If specified, log probabilities for each token position will be
// returned.
type InferenceChatCompletionParamsLogprobs struct {
	// How many tokens (for each position) to return log probabilities for.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	paramObj
}

func (r InferenceChatCompletionParamsLogprobs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceChatCompletionParamsLogprobs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceChatCompletionParamsLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Whether tool use is required or automatic. Defaults to
// ToolChoice.auto. .. deprecated:: Use tool_config instead.
type InferenceChatCompletionParamsToolChoice string

const (
	InferenceChatCompletionParamsToolChoiceAuto     InferenceChatCompletionParamsToolChoice = "auto"
	InferenceChatCompletionParamsToolChoiceRequired InferenceChatCompletionParamsToolChoice = "required"
	InferenceChatCompletionParamsToolChoiceNone     InferenceChatCompletionParamsToolChoice = "none"
)

// (Optional) Instructs the model how to format tool calls. By default, Llama Stack
// will attempt to use a format that is best adapted to the model. -
// `ToolPromptFormat.json`: The tool calls are formatted as a JSON object. -
// `ToolPromptFormat.function_tag`: The tool calls are enclosed in a
// <function=function_name> tag. - `ToolPromptFormat.python_list`: The tool calls
// are output as Python syntax -- a list of function calls. .. deprecated:: Use
// tool_config instead.
type InferenceChatCompletionParamsToolPromptFormat string

const (
	InferenceChatCompletionParamsToolPromptFormatJson        InferenceChatCompletionParamsToolPromptFormat = "json"
	InferenceChatCompletionParamsToolPromptFormatFunctionTag InferenceChatCompletionParamsToolPromptFormat = "function_tag"
	InferenceChatCompletionParamsToolPromptFormatPythonList  InferenceChatCompletionParamsToolPromptFormat = "python_list"
)

type InferenceCompletionParams struct {
	// The content to generate a completion for.
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	ModelID string `json:"model_id,required"`
	// (Optional) If True, generate an SSE event stream of the response. Defaults to
	// False.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// (Optional) If specified, log probabilities for each token position will be
	// returned.
	Logprobs InferenceCompletionParamsLogprobs `json:"logprobs,omitzero"`
	// (Optional) Grammar specification for guided (structured) decoding.
	ResponseFormat ResponseFormatUnionParam `json:"response_format,omitzero"`
	// (Optional) Parameters to control the sampling strategy.
	SamplingParams SamplingParams `json:"sampling_params,omitzero"`
	paramObj
}

func (r InferenceCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) If specified, log probabilities for each token position will be
// returned.
type InferenceCompletionParamsLogprobs struct {
	// How many tokens (for each position) to return log probabilities for.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	paramObj
}

func (r InferenceCompletionParamsLogprobs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceCompletionParamsLogprobs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceCompletionParamsLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingsParams struct {
	// List of contents to generate embeddings for. Each content can be a string or an
	// InterleavedContentItem (and hence can be multimodal). The behavior depends on
	// the model and provider. Some models may only support text.
	Contents InferenceEmbeddingsParamsContentsUnion `json:"contents,omitzero,required"`
	// The identifier of the model to use. The model must be an embedding model
	// registered with Llama Stack and available via the /models endpoint.
	ModelID string `json:"model_id,required"`
	// (Optional) Output dimensionality for the embeddings. Only supported by
	// Matryoshka models.
	OutputDimension param.Opt[int64] `json:"output_dimension,omitzero"`
	// (Optional) How is the embedding being used? This is only supported by asymmetric
	// embedding models.
	//
	// Any of "query", "document".
	TaskType InferenceEmbeddingsParamsTaskType `json:"task_type,omitzero"`
	// (Optional) Config for how to truncate text for embedding when text is longer
	// than the model's max sequence length.
	//
	// Any of "none", "start", "end".
	TextTruncation InferenceEmbeddingsParamsTextTruncation `json:"text_truncation,omitzero"`
	paramObj
}

func (r InferenceEmbeddingsParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InferenceEmbeddingsParamsContentsUnion struct {
	OfStringArray                 []string                           `json:",omitzero,inline"`
	OfInterleavedContentItemArray []InterleavedContentItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u InferenceEmbeddingsParamsContentsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringArray, u.OfInterleavedContentItemArray)
}
func (u *InferenceEmbeddingsParamsContentsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InferenceEmbeddingsParamsContentsUnion) asAny() any {
	if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfInterleavedContentItemArray) {
		return &u.OfInterleavedContentItemArray
	}
	return nil
}

// (Optional) How is the embedding being used? This is only supported by asymmetric
// embedding models.
type InferenceEmbeddingsParamsTaskType string

const (
	InferenceEmbeddingsParamsTaskTypeQuery    InferenceEmbeddingsParamsTaskType = "query"
	InferenceEmbeddingsParamsTaskTypeDocument InferenceEmbeddingsParamsTaskType = "document"
)

// (Optional) Config for how to truncate text for embedding when text is longer
// than the model's max sequence length.
type InferenceEmbeddingsParamsTextTruncation string

const (
	InferenceEmbeddingsParamsTextTruncationNone  InferenceEmbeddingsParamsTextTruncation = "none"
	InferenceEmbeddingsParamsTextTruncationStart InferenceEmbeddingsParamsTextTruncation = "start"
	InferenceEmbeddingsParamsTextTruncationEnd   InferenceEmbeddingsParamsTextTruncation = "end"
)

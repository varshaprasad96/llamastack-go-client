// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stainless-sdks/llamastack-go-client-go/internal/apijson"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/param"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/respjson"
	"github.com/stainless-sdks/llamastack-go-client-go/shared/constant"
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
	Logprobs []TokenLogProbs `json:"logprobs"`
	// (Optional) List of metrics associated with the API response
	Metrics []MetricInResponse `json:"metrics"`
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
	Role constant.Assistant `json:"role,required"`
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
	// Must be "assistant" to identify this as the model's response
	//
	// This field can be elided, and will marshal its zero value as "assistant".
	Role constant.Assistant `json:"role,required"`
	paramObj
}

func (r CompletionMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow CompletionMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletionMessageParam) UnmarshalJSON(data []byte) error {
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
	Type constant.Image `json:"type,required"`
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
func (r InterleavedContentTextContentItem) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func InterleavedContentParamOfImageContentItem(image InterleavedContentImageContentItemImageParam) InterleavedContentUnionParam {
	var variant InterleavedContentImageContentItemParam
	variant.Image = image
	return InterleavedContentUnionParam{OfImageContentItem: &variant}
}

func InterleavedContentParamOfTextContentItem(text string) InterleavedContentUnionParam {
	var variant InterleavedContentTextContentItemParam
	variant.Text = text
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
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r InterleavedContentImageContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentImageContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentImageContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r InterleavedContentTextContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentTextContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentTextContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InterleavedContentItemUnion contains all possible properties and values from
// [InterleavedContentItemImage], [InterleavedContentItemText].
//
// Use the [InterleavedContentItemUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InterleavedContentItemUnion struct {
	// This field is from variant [InterleavedContentItemImage].
	Image InterleavedContentItemImageImage `json:"image"`
	// Any of "image", "text".
	Type string `json:"type"`
	// This field is from variant [InterleavedContentItemText].
	Text string `json:"text"`
	JSON struct {
		Image respjson.Field
		Type  respjson.Field
		Text  respjson.Field
		raw   string
	} `json:"-"`
}

// anyInterleavedContentItem is implemented by each variant of
// [InterleavedContentItemUnion] to add type safety for the return type of
// [InterleavedContentItemUnion.AsAny]
type anyInterleavedContentItem interface {
	implInterleavedContentItemUnion()
}

func (InterleavedContentItemImage) implInterleavedContentItemUnion() {}
func (InterleavedContentItemText) implInterleavedContentItemUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := InterleavedContentItemUnion.AsAny().(type) {
//	case llamastackgoclient.InterleavedContentItemImage:
//	case llamastackgoclient.InterleavedContentItemText:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u InterleavedContentItemUnion) AsAny() anyInterleavedContentItem {
	switch u.Type {
	case "image":
		return u.AsImage()
	case "text":
		return u.AsText()
	}
	return nil
}

func (u InterleavedContentItemUnion) AsImage() (v InterleavedContentItemImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentItemUnion) AsText() (v InterleavedContentItemText) {
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
type InterleavedContentItemImage struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentItemImageImage `json:"image,required"`
	// Discriminator type of the content item. Always "image"
	Type constant.Image `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemImage) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type InterleavedContentItemImageImage struct {
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
func (r InterleavedContentItemImageImage) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemImageImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type InterleavedContentItemText struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
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
func (r InterleavedContentItemText) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func InterleavedContentItemParamOfImage(image InterleavedContentItemImageImageParam) InterleavedContentItemUnionParam {
	var variant InterleavedContentItemImageParam
	variant.Image = image
	return InterleavedContentItemUnionParam{OfImage: &variant}
}

func InterleavedContentItemParamOfText(text string) InterleavedContentItemUnionParam {
	var variant InterleavedContentItemTextParam
	variant.Text = text
	return InterleavedContentItemUnionParam{OfText: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InterleavedContentItemUnionParam struct {
	OfImage *InterleavedContentItemImageParam `json:",omitzero,inline"`
	OfText  *InterleavedContentItemTextParam  `json:",omitzero,inline"`
	paramUnion
}

func (u InterleavedContentItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfImage, u.OfText)
}
func (u *InterleavedContentItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InterleavedContentItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfImage) {
		return u.OfImage
	} else if !param.IsOmitted(u.OfText) {
		return u.OfText
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetImage() *InterleavedContentItemImageImageParam {
	if vt := u.OfImage; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetType() *string {
	if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InterleavedContentItemUnionParam](
		"type",
		apijson.Discriminator[InterleavedContentItemImageParam]("image"),
		apijson.Discriminator[InterleavedContentItemTextParam]("text"),
	)
}

// A image content item
//
// The properties Image, Type are required.
type InterleavedContentItemImageParam struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentItemImageImageParam `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r InterleavedContentItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type InterleavedContentItemImageImageParam struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL URLParam `json:"url,omitzero"`
	paramObj
}

func (r InterleavedContentItemImageImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemImageImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemImageImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type InterleavedContentItemTextParam struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r InterleavedContentItemTextParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func MessageParamOfUser[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T) MessageUnionParam {
	var user UserMessageParam
	switch v := any(content).(type) {
	case string:
		user.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		user.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		user.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		user.Content.OfInterleavedContentItemArray = v
	}
	return MessageUnionParam{OfUser: &user}
}

func MessageParamOfSystem[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T) MessageUnionParam {
	var system SystemMessageParam
	switch v := any(content).(type) {
	case string:
		system.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		system.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		system.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		system.Content.OfInterleavedContentItemArray = v
	}
	return MessageUnionParam{OfSystem: &system}
}

func MessageParamOfTool[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](callID string, content T) MessageUnionParam {
	var tool ToolResponseMessageParam
	tool.CallID = callID
	switch v := any(content).(type) {
	case string:
		tool.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		tool.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		tool.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		tool.Content.OfInterleavedContentItemArray = v
	}
	return MessageUnionParam{OfTool: &tool}
}

func MessageParamOfAssistant[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T, stopReason CompletionMessageStopReason) MessageUnionParam {
	var assistant CompletionMessageParam
	switch v := any(content).(type) {
	case string:
		assistant.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		assistant.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		assistant.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		assistant.Content.OfInterleavedContentItemArray = v
	}
	assistant.StopReason = stopReason
	return MessageUnionParam{OfAssistant: &assistant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageUnionParam struct {
	OfUser      *UserMessageParam         `json:",omitzero,inline"`
	OfSystem    *SystemMessageParam       `json:",omitzero,inline"`
	OfTool      *ToolResponseMessageParam `json:",omitzero,inline"`
	OfAssistant *CompletionMessageParam   `json:",omitzero,inline"`
	paramUnion
}

func (u MessageUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUser, u.OfSystem, u.OfTool, u.OfAssistant)
}
func (u *MessageUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageUnionParam) asAny() any {
	if !param.IsOmitted(u.OfUser) {
		return u.OfUser
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfTool) {
		return u.OfTool
	} else if !param.IsOmitted(u.OfAssistant) {
		return u.OfAssistant
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetContext() *InterleavedContentUnionParam {
	if vt := u.OfUser; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetCallID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.CallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetStopReason() *string {
	if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.StopReason)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetToolCalls() []ToolCallParam {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetRole() *string {
	if vt := u.OfUser; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's Content property, if present.
func (u MessageUnionParam) GetContent() *InterleavedContentUnionParam {
	if vt := u.OfUser; vt != nil {
		return &vt.Content
	} else if vt := u.OfSystem; vt != nil {
		return &vt.Content
	} else if vt := u.OfTool; vt != nil {
		return &vt.Content
	} else if vt := u.OfAssistant; vt != nil {
		return &vt.Content
	}
	return nil
}

func init() {
	apijson.RegisterUnion[MessageUnionParam](
		"role",
		apijson.Discriminator[UserMessageParam]("user"),
		apijson.Discriminator[SystemMessageParam]("system"),
		apijson.Discriminator[ToolResponseMessageParam]("tool"),
		apijson.Discriminator[CompletionMessageParam]("assistant"),
	)
}

// A metric value included in API responses.
type MetricInResponse struct {
	// The name of the metric
	Metric string `json:"metric,required"`
	// The numeric value of the metric
	Value float64 `json:"value,required"`
	// (Optional) The unit of measurement for the metric value
	Unit string `json:"unit"`
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
// [ResponseFormatJsonSchema], [ResponseFormatGrammar].
//
// Use the [ResponseFormatUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseFormatUnion struct {
	// This field is from variant [ResponseFormatJsonSchema].
	JsonSchema map[string]ResponseFormatJsonSchemaJsonSchemaUnion `json:"json_schema"`
	// Any of "json_schema", "grammar".
	Type string `json:"type"`
	// This field is from variant [ResponseFormatGrammar].
	Bnf  map[string]ResponseFormatGrammarBnfUnion `json:"bnf"`
	JSON struct {
		JsonSchema respjson.Field
		Type       respjson.Field
		Bnf        respjson.Field
		raw        string
	} `json:"-"`
}

// anyResponseFormat is implemented by each variant of [ResponseFormatUnion] to add
// type safety for the return type of [ResponseFormatUnion.AsAny]
type anyResponseFormat interface {
	implResponseFormatUnion()
}

func (ResponseFormatJsonSchema) implResponseFormatUnion() {}
func (ResponseFormatGrammar) implResponseFormatUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseFormatUnion.AsAny().(type) {
//	case llamastackgoclient.ResponseFormatJsonSchema:
//	case llamastackgoclient.ResponseFormatGrammar:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseFormatUnion) AsAny() anyResponseFormat {
	switch u.Type {
	case "json_schema":
		return u.AsJsonSchema()
	case "grammar":
		return u.AsGrammar()
	}
	return nil
}

func (u ResponseFormatUnion) AsJsonSchema() (v ResponseFormatJsonSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatUnion) AsGrammar() (v ResponseFormatGrammar) {
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

// Configuration for JSON schema-guided response generation.
type ResponseFormatJsonSchema struct {
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model.
	JsonSchema map[string]ResponseFormatJsonSchemaJsonSchemaUnion `json:"json_schema,required"`
	// Must be "json_schema" to identify this format type
	Type constant.JsonSchema `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JsonSchema  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseFormatJsonSchema) RawJSON() string { return r.JSON.raw }
func (r *ResponseFormatJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseFormatJsonSchemaJsonSchemaUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseFormatJsonSchemaJsonSchemaUnion struct {
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

func (u ResponseFormatJsonSchemaJsonSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatJsonSchemaJsonSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatJsonSchemaJsonSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatJsonSchemaJsonSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseFormatJsonSchemaJsonSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseFormatJsonSchemaJsonSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for grammar-guided response generation.
type ResponseFormatGrammar struct {
	// The BNF grammar specification the response should conform to
	Bnf map[string]ResponseFormatGrammarBnfUnion `json:"bnf,required"`
	// Must be "grammar" to identify this format type
	Type constant.Grammar `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bnf         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseFormatGrammar) RawJSON() string { return r.JSON.raw }
func (r *ResponseFormatGrammar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseFormatGrammarBnfUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseFormatGrammarBnfUnion struct {
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

func (u ResponseFormatGrammarBnfUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatGrammarBnfUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatGrammarBnfUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatGrammarBnfUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseFormatGrammarBnfUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseFormatGrammarBnfUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ResponseFormatParamOfJsonSchema(jsonSchema map[string]ResponseFormatJsonSchemaJsonSchemaUnionParam) ResponseFormatUnionParam {
	var variant ResponseFormatJsonSchemaParam
	variant.JsonSchema = jsonSchema
	return ResponseFormatUnionParam{OfJsonSchema: &variant}
}

func ResponseFormatParamOfGrammar(bnf map[string]ResponseFormatGrammarBnfUnionParam) ResponseFormatUnionParam {
	var grammar ResponseFormatGrammarParam
	grammar.Bnf = bnf
	return ResponseFormatUnionParam{OfGrammar: &grammar}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseFormatUnionParam struct {
	OfJsonSchema *ResponseFormatJsonSchemaParam `json:",omitzero,inline"`
	OfGrammar    *ResponseFormatGrammarParam    `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseFormatUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfJsonSchema, u.OfGrammar)
}
func (u *ResponseFormatUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseFormatUnionParam) asAny() any {
	if !param.IsOmitted(u.OfJsonSchema) {
		return u.OfJsonSchema
	} else if !param.IsOmitted(u.OfGrammar) {
		return u.OfGrammar
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseFormatUnionParam) GetJsonSchema() map[string]ResponseFormatJsonSchemaJsonSchemaUnionParam {
	if vt := u.OfJsonSchema; vt != nil {
		return vt.JsonSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseFormatUnionParam) GetBnf() map[string]ResponseFormatGrammarBnfUnionParam {
	if vt := u.OfGrammar; vt != nil {
		return vt.Bnf
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseFormatUnionParam) GetType() *string {
	if vt := u.OfJsonSchema; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGrammar; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseFormatUnionParam](
		"type",
		apijson.Discriminator[ResponseFormatJsonSchemaParam]("json_schema"),
		apijson.Discriminator[ResponseFormatGrammarParam]("grammar"),
	)
}

// Configuration for JSON schema-guided response generation.
//
// The properties JsonSchema, Type are required.
type ResponseFormatJsonSchemaParam struct {
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model.
	JsonSchema map[string]ResponseFormatJsonSchemaJsonSchemaUnionParam `json:"json_schema,omitzero,required"`
	// Must be "json_schema" to identify this format type
	//
	// This field can be elided, and will marshal its zero value as "json_schema".
	Type constant.JsonSchema `json:"type,required"`
	paramObj
}

func (r ResponseFormatJsonSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseFormatJsonSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseFormatJsonSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseFormatJsonSchemaJsonSchemaUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseFormatJsonSchemaJsonSchemaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseFormatJsonSchemaJsonSchemaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseFormatJsonSchemaJsonSchemaUnionParam) asAny() any {
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

// Configuration for grammar-guided response generation.
//
// The properties Bnf, Type are required.
type ResponseFormatGrammarParam struct {
	// The BNF grammar specification the response should conform to
	Bnf map[string]ResponseFormatGrammarBnfUnionParam `json:"bnf,omitzero,required"`
	// Must be "grammar" to identify this format type
	//
	// This field can be elided, and will marshal its zero value as "grammar".
	Type constant.Grammar `json:"type,required"`
	paramObj
}

func (r ResponseFormatGrammarParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseFormatGrammarParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseFormatGrammarParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseFormatGrammarBnfUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseFormatGrammarBnfUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseFormatGrammarBnfUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseFormatGrammarBnfUnionParam) asAny() any {
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
// [SamplingParamsStrategyGreedyResp], [SamplingParamsStrategyTopPResp],
// [SamplingParamsStrategyTopKResp].
//
// Use the [SamplingParamsStrategyUnionResp.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SamplingParamsStrategyUnionResp struct {
	// Any of "greedy", "top_p", "top_k".
	Type string `json:"type"`
	// This field is from variant [SamplingParamsStrategyTopPResp].
	Temperature float64 `json:"temperature"`
	// This field is from variant [SamplingParamsStrategyTopPResp].
	TopP float64 `json:"top_p"`
	// This field is from variant [SamplingParamsStrategyTopKResp].
	TopK int64 `json:"top_k"`
	JSON struct {
		Type        respjson.Field
		Temperature respjson.Field
		TopP        respjson.Field
		TopK        respjson.Field
		raw         string
	} `json:"-"`
}

// anySamplingParamsStrategyResp is implemented by each variant of
// [SamplingParamsStrategyUnionResp] to add type safety for the return type of
// [SamplingParamsStrategyUnionResp.AsAny]
type anySamplingParamsStrategyResp interface {
	implSamplingParamsStrategyUnionResp()
}

func (SamplingParamsStrategyGreedyResp) implSamplingParamsStrategyUnionResp() {}
func (SamplingParamsStrategyTopPResp) implSamplingParamsStrategyUnionResp()   {}
func (SamplingParamsStrategyTopKResp) implSamplingParamsStrategyUnionResp()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := SamplingParamsStrategyUnionResp.AsAny().(type) {
//	case llamastackgoclient.SamplingParamsStrategyGreedyResp:
//	case llamastackgoclient.SamplingParamsStrategyTopPResp:
//	case llamastackgoclient.SamplingParamsStrategyTopKResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u SamplingParamsStrategyUnionResp) AsAny() anySamplingParamsStrategyResp {
	switch u.Type {
	case "greedy":
		return u.AsGreedy()
	case "top_p":
		return u.AsTopP()
	case "top_k":
		return u.AsTopK()
	}
	return nil
}

func (u SamplingParamsStrategyUnionResp) AsGreedy() (v SamplingParamsStrategyGreedyResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SamplingParamsStrategyUnionResp) AsTopP() (v SamplingParamsStrategyTopPResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SamplingParamsStrategyUnionResp) AsTopK() (v SamplingParamsStrategyTopKResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SamplingParamsStrategyUnionResp) RawJSON() string { return u.JSON.raw }

func (r *SamplingParamsStrategyUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Greedy sampling strategy that selects the highest probability token at each
// step.
type SamplingParamsStrategyGreedyResp struct {
	// Must be "greedy" to identify this sampling strategy
	Type constant.Greedy `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SamplingParamsStrategyGreedyResp) RawJSON() string { return r.JSON.raw }
func (r *SamplingParamsStrategyGreedyResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-p (nucleus) sampling strategy that samples from the smallest set of tokens
// with cumulative probability >= p.
type SamplingParamsStrategyTopPResp struct {
	// Must be "top_p" to identify this sampling strategy
	Type constant.TopP `json:"type,required"`
	// Controls randomness in sampling. Higher values increase randomness
	Temperature float64 `json:"temperature"`
	// Cumulative probability threshold for nucleus sampling. Defaults to 0.95
	TopP float64 `json:"top_p"`
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
func (r SamplingParamsStrategyTopPResp) RawJSON() string { return r.JSON.raw }
func (r *SamplingParamsStrategyTopPResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-k sampling strategy that restricts sampling to the k most likely tokens.
type SamplingParamsStrategyTopKResp struct {
	// Number of top tokens to consider for sampling. Must be at least 1
	TopK int64 `json:"top_k,required"`
	// Must be "top_k" to identify this sampling strategy
	Type constant.TopK `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TopK        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SamplingParamsStrategyTopKResp) RawJSON() string { return r.JSON.raw }
func (r *SamplingParamsStrategyTopKResp) UnmarshalJSON(data []byte) error {
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
	OfGreedy *SamplingParamsStrategyGreedy `json:",omitzero,inline"`
	OfTopP   *SamplingParamsStrategyTopP   `json:",omitzero,inline"`
	OfTopK   *SamplingParamsStrategyTopK   `json:",omitzero,inline"`
	paramUnion
}

func (u SamplingParamsStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGreedy, u.OfTopP, u.OfTopK)
}
func (u *SamplingParamsStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SamplingParamsStrategyUnion) asAny() any {
	if !param.IsOmitted(u.OfGreedy) {
		return u.OfGreedy
	} else if !param.IsOmitted(u.OfTopP) {
		return u.OfTopP
	} else if !param.IsOmitted(u.OfTopK) {
		return u.OfTopK
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTemperature() *float64 {
	if vt := u.OfTopP; vt != nil && vt.Temperature.Valid() {
		return &vt.Temperature.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTopP() *float64 {
	if vt := u.OfTopP; vt != nil && vt.TopP.Valid() {
		return &vt.TopP.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTopK() *int64 {
	if vt := u.OfTopK; vt != nil {
		return &vt.TopK
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetType() *string {
	if vt := u.OfGreedy; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTopP; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTopK; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SamplingParamsStrategyUnion](
		"type",
		apijson.Discriminator[SamplingParamsStrategyGreedy]("greedy"),
		apijson.Discriminator[SamplingParamsStrategyTopP]("top_p"),
		apijson.Discriminator[SamplingParamsStrategyTopK]("top_k"),
	)
}

func NewSamplingParamsStrategyGreedy() SamplingParamsStrategyGreedy {
	return SamplingParamsStrategyGreedy{
		Type: "greedy",
	}
}

// Greedy sampling strategy that selects the highest probability token at each
// step.
//
// This struct has a constant value, construct it with
// [NewSamplingParamsStrategyGreedy].
type SamplingParamsStrategyGreedy struct {
	// Must be "greedy" to identify this sampling strategy
	Type constant.Greedy `json:"type,required"`
	paramObj
}

func (r SamplingParamsStrategyGreedy) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyGreedy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyGreedy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-p (nucleus) sampling strategy that samples from the smallest set of tokens
// with cumulative probability >= p.
//
// The property Type is required.
type SamplingParamsStrategyTopP struct {
	// Controls randomness in sampling. Higher values increase randomness
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Cumulative probability threshold for nucleus sampling. Defaults to 0.95
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// Must be "top_p" to identify this sampling strategy
	//
	// This field can be elided, and will marshal its zero value as "top_p".
	Type constant.TopP `json:"type,required"`
	paramObj
}

func (r SamplingParamsStrategyTopP) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyTopP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyTopP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-k sampling strategy that restricts sampling to the k most likely tokens.
//
// The properties TopK, Type are required.
type SamplingParamsStrategyTopK struct {
	// Number of top tokens to consider for sampling. Must be at least 1
	TopK int64 `json:"top_k,required"`
	// Must be "top_k" to identify this sampling strategy
	//
	// This field can be elided, and will marshal its zero value as "top_k".
	Type constant.TopK `json:"type,required"`
	paramObj
}

func (r SamplingParamsStrategyTopK) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyTopK
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyTopK) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// This field can be elided, and will marshal its zero value as "system".
	Role constant.System `json:"role,required"`
	paramObj
}

func (r SystemMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageParam) UnmarshalJSON(data []byte) error {
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

// Response from a batch chat completion request.
type InferenceBatchChatCompletionResponse struct {
	// List of chat completion responses, one for each conversation in the batch
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

// Response from a batch completion request.
type InferenceBatchCompletionResponse struct {
	// List of completion responses, one for each input in the batch
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

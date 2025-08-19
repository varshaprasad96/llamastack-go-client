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
	"github.com/varshaprasad96/llamastack-go-client/shared/constant"
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

type ChatCompletionResponse struct {
	CompletionMessage CompletionMessage  `json:"completion_message,required"`
	Logprobs          []TokenLogProbs    `json:"logprobs"`
	Metrics           []MetricInResponse `json:"metrics"`
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

type CompletionMessage struct {
	Content CompletionMessageContentUnion `json:"content,required"`
	// The role of the message sender
	//
	// Any of "user", "assistant", "system", "tool".
	Role      CompletionMessageRole `json:"role,required"`
	Name      string                `json:"name"`
	ToolCalls []ToolCall            `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
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

// CompletionMessageContentUnion contains all possible properties and values from
// [string], [[]CompletionMessageContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCompletionMessageContentArray]
type CompletionMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]CompletionMessageContentArrayItemUnion] instead of an object.
	OfCompletionMessageContentArray []CompletionMessageContentArrayItemUnion `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfCompletionMessageContentArray respjson.Field
		raw                             string
	} `json:"-"`
}

func (u CompletionMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompletionMessageContentUnion) AsCompletionMessageContentArray() (v []CompletionMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CompletionMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *CompletionMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CompletionMessageContentArrayItemUnion contains all possible properties and
// values from [string], [string], [CompletionMessageContentArrayItemImageContent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCompletionMessageContentArrayItemString]
type CompletionMessageContentArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfCompletionMessageContentArrayItemString string `json:",inline"`
	// This field is from variant [CompletionMessageContentArrayItemImageContent].
	ImageURL CompletionMessageContentArrayItemImageContentImageURL `json:"image_url"`
	// This field is from variant [CompletionMessageContentArrayItemImageContent].
	Type constant.Image `json:"type"`
	JSON struct {
		OfCompletionMessageContentArrayItemString respjson.Field
		ImageURL                                  respjson.Field
		Type                                      respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u CompletionMessageContentArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompletionMessageContentArrayItemUnion) AsCompletionMessageContentArrayItemString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompletionMessageContentArrayItemUnion) AsCompletionMessageContentArrayItemImageContent() (v CompletionMessageContentArrayItemImageContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CompletionMessageContentArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *CompletionMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionMessageContentArrayItemImageContent struct {
	ImageURL CompletionMessageContentArrayItemImageContentImageURL `json:"image_url,required"`
	Type     constant.Image                                        `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionMessageContentArrayItemImageContent) RawJSON() string { return r.JSON.raw }
func (r *CompletionMessageContentArrayItemImageContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionMessageContentArrayItemImageContentImageURL struct {
	URL string `json:"url,required"`
	// Any of "low", "high", "auto".
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
func (r CompletionMessageContentArrayItemImageContentImageURL) RawJSON() string { return r.JSON.raw }
func (r *CompletionMessageContentArrayItemImageContentImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The role of the message sender
type CompletionMessageRole string

const (
	CompletionMessageRoleUser      CompletionMessageRole = "user"
	CompletionMessageRoleAssistant CompletionMessageRole = "assistant"
	CompletionMessageRoleSystem    CompletionMessageRole = "system"
	CompletionMessageRoleTool      CompletionMessageRole = "tool"
)

// The properties Content, Role are required.
type MessageParam struct {
	Content MessageContentUnionParam `json:"content,omitzero,required"`
	// The role of the message sender
	//
	// Any of "user", "assistant", "system", "tool".
	Role      MessageRole       `json:"role,omitzero,required"`
	Name      param.Opt[string] `json:"name,omitzero"`
	ToolCalls []ToolCallParam   `json:"tool_calls,omitzero"`
	paramObj
}

func (r MessageParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageContentUnionParam struct {
	OfString              param.Opt[string]                   `json:",omitzero,inline"`
	OfMessageContentArray []MessageContentArrayItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u MessageContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfMessageContentArray)
}
func (u *MessageContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfMessageContentArray) {
		return &u.OfMessageContentArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageContentArrayItemUnionParam struct {
	OfString                              param.Opt[string]                         `json:",omitzero,inline"`
	OfMessageContentArrayItemString       param.Opt[string]                         `json:",omitzero,inline"`
	OfMessageContentArrayItemImageContent *MessageContentArrayItemImageContentParam `json:",omitzero,inline"`
	paramUnion
}

func (u MessageContentArrayItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfMessageContentArrayItemString, u.OfMessageContentArrayItemImageContent)
}
func (u *MessageContentArrayItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageContentArrayItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfMessageContentArrayItemString) {
		return &u.OfMessageContentArrayItemString.Value
	} else if !param.IsOmitted(u.OfMessageContentArrayItemImageContent) {
		return u.OfMessageContentArrayItemImageContent
	}
	return nil
}

// The properties ImageURL, Type are required.
type MessageContentArrayItemImageContentParam struct {
	ImageURL MessageContentArrayItemImageContentImageURLParam `json:"image_url,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r MessageContentArrayItemImageContentParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentArrayItemImageContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContentArrayItemImageContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type MessageContentArrayItemImageContentImageURLParam struct {
	URL string `json:"url,required"`
	// Any of "low", "high", "auto".
	Detail string `json:"detail,omitzero"`
	paramObj
}

func (r MessageContentArrayItemImageContentImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentArrayItemImageContentImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContentArrayItemImageContentImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageContentArrayItemImageContentImageURLParam](
		"detail", "low", "high", "auto",
	)
}

// The role of the message sender
type MessageRole string

const (
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
	MessageRoleSystem    MessageRole = "system"
	MessageRoleTool      MessageRole = "tool"
)

type MetricInResponse struct {
	Name  string  `json:"name,required"`
	Value float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricInResponse) RawJSON() string { return r.JSON.raw }
func (r *MetricInResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolCall struct {
	ID       string            `json:"id,required"`
	Function ToolCallFunction  `json:"function,required"`
	Type     constant.Function `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

type ToolCallFunction struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *ToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Function, Type are required.
type ToolCallParam struct {
	ID       string                `json:"id,required"`
	Function ToolCallFunctionParam `json:"function,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "function".
	Type constant.Function `json:"type,required"`
	paramObj
}

func (r ToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Arguments, Name are required.
type ToolCallFunctionParam struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	paramObj
}

func (r ToolCallFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolCallFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolCallFunctionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceBatchChatCompletionResponse struct {
	Completions []ChatCompletionResponse `json:"completions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completions respjson.Field
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
	Completions []CompletionResponse `json:"completions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceBatchCompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *InferenceBatchCompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingsResponse struct {
	Embeddings [][]float64 `json:"embeddings,required"`
	Model      string      `json:"model,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embeddings  respjson.Field
		Model       respjson.Field
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
	Messages []MessageParam  `json:"messages,omitzero,required"`
	Model    string          `json:"model,required"`
	Stream   param.Opt[bool] `json:"stream,omitzero"`
	paramObj
}

func (r InferenceBatchChatCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchChatCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchChatCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceBatchCompletionParams struct {
	Contents []string        `json:"contents,omitzero,required"`
	Model    string          `json:"model,required"`
	Stream   param.Opt[bool] `json:"stream,omitzero"`
	paramObj
}

func (r InferenceBatchCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceChatCompletionParams struct {
	Messages []MessageParam  `json:"messages,omitzero,required"`
	Model    string          `json:"model,required"`
	Stream   param.Opt[bool] `json:"stream,omitzero"`
	paramObj
}

func (r InferenceChatCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceChatCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceChatCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceCompletionParams struct {
	Content string          `json:"content,required"`
	Model   string          `json:"model,required"`
	Stream  param.Opt[bool] `json:"stream,omitzero"`
	paramObj
}

func (r InferenceCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingsParams struct {
	Input InferenceEmbeddingsParamsInputUnion `json:"input,omitzero,required"`
	Model string                              `json:"model,required"`
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
type InferenceEmbeddingsParamsInputUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u InferenceEmbeddingsParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *InferenceEmbeddingsParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InferenceEmbeddingsParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

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

// OpenAIMessageParamUnionResp contains all possible properties and values from
// [OpenAIMessageParamUserResp], [OpenAIMessageParamSystemResp],
// [OpenAIMessageParamAssistantResp], [OpenAIMessageParamToolResp],
// [OpenAIMessageParamDeveloperResp].
//
// Use the [OpenAIMessageParamUnionResp.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAIMessageParamUnionResp struct {
	// This field is a union of [OpenAIMessageParamUserContentUnionResp],
	// [OpenAIMessageParamSystemContentUnionResp],
	// [OpenAIMessageParamAssistantContentUnionResp],
	// [OpenAIMessageParamToolContentUnionResp],
	// [OpenAIMessageParamDeveloperContentUnionResp]
	Content OpenAIMessageParamUnionRespContent `json:"content"`
	// Any of "user", "system", "assistant", "tool", "developer".
	Role string `json:"role"`
	Name string `json:"name"`
	// This field is from variant [OpenAIMessageParamAssistantResp].
	ToolCalls []ChatCompletionToolCall `json:"tool_calls"`
	// This field is from variant [OpenAIMessageParamToolResp].
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

// anyOpenAIMessageParamResp is implemented by each variant of
// [OpenAIMessageParamUnionResp] to add type safety for the return type of
// [OpenAIMessageParamUnionResp.AsAny]
type anyOpenAIMessageParamResp interface {
	implOpenAIMessageParamUnionResp()
}

func (OpenAIMessageParamUserResp) implOpenAIMessageParamUnionResp()      {}
func (OpenAIMessageParamSystemResp) implOpenAIMessageParamUnionResp()    {}
func (OpenAIMessageParamAssistantResp) implOpenAIMessageParamUnionResp() {}
func (OpenAIMessageParamToolResp) implOpenAIMessageParamUnionResp()      {}
func (OpenAIMessageParamDeveloperResp) implOpenAIMessageParamUnionResp() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := OpenAIMessageParamUnionResp.AsAny().(type) {
//	case llamastackclient.OpenAIMessageParamUserResp:
//	case llamastackclient.OpenAIMessageParamSystemResp:
//	case llamastackclient.OpenAIMessageParamAssistantResp:
//	case llamastackclient.OpenAIMessageParamToolResp:
//	case llamastackclient.OpenAIMessageParamDeveloperResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u OpenAIMessageParamUnionResp) AsAny() anyOpenAIMessageParamResp {
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

func (u OpenAIMessageParamUnionResp) AsUser() (v OpenAIMessageParamUserResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamUnionResp) AsSystem() (v OpenAIMessageParamSystemResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamUnionResp) AsAssistant() (v OpenAIMessageParamAssistantResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamUnionResp) AsTool() (v OpenAIMessageParamToolResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamUnionResp) AsDeveloper() (v OpenAIMessageParamDeveloperResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIMessageParamUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIMessageParamUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIMessageParamUnionRespContent is an implicit subunion of
// [OpenAIMessageParamUnionResp]. OpenAIMessageParamUnionRespContent provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OpenAIMessageParamUnionResp].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOpenAIMessageUserContentArray OfContentPartTextArray]
type OpenAIMessageParamUnionRespContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIMessageParamUserContentArrayItemUnionResp] instead of an object.
	OfOpenAIMessageUserContentArray []OpenAIMessageParamUserContentArrayItemUnionResp `json:",inline"`
	// This field will be present if the value is a [[]ContentPartTextParamResp]
	// instead of an object.
	OfContentPartTextArray []ContentPartTextParamResp `json:",inline"`
	JSON                   struct {
		OfString                        respjson.Field
		OfOpenAIMessageUserContentArray respjson.Field
		OfContentPartTextArray          respjson.Field
		raw                             string
	} `json:"-"`
}

func (r *OpenAIMessageParamUnionRespContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OpenAIMessageParamUnionResp to a OpenAIMessageParamUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OpenAIMessageParamUnion.Overrides()
func (r OpenAIMessageParamUnionResp) ToParam() OpenAIMessageParamUnion {
	return param.Override[OpenAIMessageParamUnion](json.RawMessage(r.RawJSON()))
}

// A message from the user in an OpenAI-compatible chat completion request.
type OpenAIMessageParamUserResp struct {
	// The content of the message, which can include text and other media
	Content OpenAIMessageParamUserContentUnionResp `json:"content,required"`
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
func (r OpenAIMessageParamUserResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIMessageParamUserResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIMessageParamUserContentUnionResp contains all possible properties and
// values from [string], [[]OpenAIMessageParamUserContentArrayItemUnionResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOpenAIMessageUserContentArray]
type OpenAIMessageParamUserContentUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]OpenAIMessageParamUserContentArrayItemUnionResp] instead of an object.
	OfOpenAIMessageUserContentArray []OpenAIMessageParamUserContentArrayItemUnionResp `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfOpenAIMessageUserContentArray respjson.Field
		raw                             string
	} `json:"-"`
}

func (u OpenAIMessageParamUserContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamUserContentUnionResp) AsOpenAIMessageParamUserContentArrayRespArray() (v []OpenAIMessageParamUserContentArrayItemUnionResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIMessageParamUserContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIMessageParamUserContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIMessageParamUserContentArrayItemUnionResp contains all possible properties
// and values from [ContentPartTextParamResp],
// [OpenAIMessageParamUserContentArrayItemImageURLResp],
// [OpenAIMessageParamUserContentArrayItemFileResp].
//
// Use the [OpenAIMessageParamUserContentArrayItemUnionResp.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OpenAIMessageParamUserContentArrayItemUnionResp struct {
	// This field is from variant [ContentPartTextParamResp].
	Text string `json:"text"`
	// Any of "text", "image_url", "file".
	Type string `json:"type"`
	// This field is from variant [OpenAIMessageParamUserContentArrayItemImageURLResp].
	ImageURL OpenAIMessageParamUserContentArrayItemImageURLImageURLResp `json:"image_url"`
	// This field is from variant [OpenAIMessageParamUserContentArrayItemFileResp].
	File OpenAIMessageParamUserContentArrayItemFileFileResp `json:"file"`
	JSON struct {
		Text     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		File     respjson.Field
		raw      string
	} `json:"-"`
}

// anyOpenAIMessageParamUserContentArrayItemResp is implemented by each variant of
// [OpenAIMessageParamUserContentArrayItemUnionResp] to add type safety for the
// return type of [OpenAIMessageParamUserContentArrayItemUnionResp.AsAny]
type anyOpenAIMessageParamUserContentArrayItemResp interface {
	implOpenAIMessageParamUserContentArrayItemUnionResp()
}

func (ContentPartTextParamResp) implOpenAIMessageParamUserContentArrayItemUnionResp() {}
func (OpenAIMessageParamUserContentArrayItemImageURLResp) implOpenAIMessageParamUserContentArrayItemUnionResp() {
}
func (OpenAIMessageParamUserContentArrayItemFileResp) implOpenAIMessageParamUserContentArrayItemUnionResp() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := OpenAIMessageParamUserContentArrayItemUnionResp.AsAny().(type) {
//	case llamastackclient.ContentPartTextParamResp:
//	case llamastackclient.OpenAIMessageParamUserContentArrayItemImageURLResp:
//	case llamastackclient.OpenAIMessageParamUserContentArrayItemFileResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u OpenAIMessageParamUserContentArrayItemUnionResp) AsAny() anyOpenAIMessageParamUserContentArrayItemResp {
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

func (u OpenAIMessageParamUserContentArrayItemUnionResp) AsText() (v ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamUserContentArrayItemUnionResp) AsImageURL() (v OpenAIMessageParamUserContentArrayItemImageURLResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamUserContentArrayItemUnionResp) AsFile() (v OpenAIMessageParamUserContentArrayItemFileResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIMessageParamUserContentArrayItemUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIMessageParamUserContentArrayItemUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
type OpenAIMessageParamUserContentArrayItemImageURLResp struct {
	// Image URL specification and processing details
	ImageURL OpenAIMessageParamUserContentArrayItemImageURLImageURLResp `json:"image_url,required"`
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
func (r OpenAIMessageParamUserContentArrayItemImageURLResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIMessageParamUserContentArrayItemImageURLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
type OpenAIMessageParamUserContentArrayItemImageURLImageURLResp struct {
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
func (r OpenAIMessageParamUserContentArrayItemImageURLImageURLResp) RawJSON() string {
	return r.JSON.raw
}
func (r *OpenAIMessageParamUserContentArrayItemImageURLImageURLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIMessageParamUserContentArrayItemFileResp struct {
	File OpenAIMessageParamUserContentArrayItemFileFileResp `json:"file,required"`
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
func (r OpenAIMessageParamUserContentArrayItemFileResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIMessageParamUserContentArrayItemFileResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIMessageParamUserContentArrayItemFileFileResp struct {
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
func (r OpenAIMessageParamUserContentArrayItemFileFileResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIMessageParamUserContentArrayItemFileFileResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
type OpenAIMessageParamSystemResp struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content OpenAIMessageParamSystemContentUnionResp `json:"content,required"`
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
func (r OpenAIMessageParamSystemResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIMessageParamSystemResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIMessageParamSystemContentUnionResp contains all possible properties and
// values from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type OpenAIMessageParamSystemContentUnionResp struct {
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

func (u OpenAIMessageParamSystemContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamSystemContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIMessageParamSystemContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIMessageParamSystemContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
type OpenAIMessageParamAssistantResp struct {
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// The content of the model's response
	Content OpenAIMessageParamAssistantContentUnionResp `json:"content"`
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
func (r OpenAIMessageParamAssistantResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIMessageParamAssistantResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIMessageParamAssistantContentUnionResp contains all possible properties and
// values from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type OpenAIMessageParamAssistantContentUnionResp struct {
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

func (u OpenAIMessageParamAssistantContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamAssistantContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIMessageParamAssistantContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIMessageParamAssistantContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
type OpenAIMessageParamToolResp struct {
	// The response content from the tool
	Content OpenAIMessageParamToolContentUnionResp `json:"content,required"`
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
func (r OpenAIMessageParamToolResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIMessageParamToolResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIMessageParamToolContentUnionResp contains all possible properties and
// values from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type OpenAIMessageParamToolContentUnionResp struct {
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

func (u OpenAIMessageParamToolContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamToolContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIMessageParamToolContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIMessageParamToolContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
type OpenAIMessageParamDeveloperResp struct {
	// The content of the developer message
	Content OpenAIMessageParamDeveloperContentUnionResp `json:"content,required"`
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
func (r OpenAIMessageParamDeveloperResp) RawJSON() string { return r.JSON.raw }
func (r *OpenAIMessageParamDeveloperResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIMessageParamDeveloperContentUnionResp contains all possible properties and
// values from [string], [[]ContentPartTextParamResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfContentPartTextArray]
type OpenAIMessageParamDeveloperContentUnionResp struct {
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

func (u OpenAIMessageParamDeveloperContentUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAIMessageParamDeveloperContentUnionResp) AsContentPartTextParamArray() (v []ContentPartTextParamResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAIMessageParamDeveloperContentUnionResp) RawJSON() string { return u.JSON.raw }

func (r *OpenAIMessageParamDeveloperContentUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func OpenAIMessageParamOfUser[
	T string | []OpenAIMessageParamUserContentArrayItemUnion,
](content T) OpenAIMessageParamUnion {
	var user OpenAIMessageParamUser
	switch v := any(content).(type) {
	case string:
		user.Content.OfString = param.NewOpt(v)
	case []OpenAIMessageParamUserContentArrayItemUnion:
		user.Content.OfOpenAIMessageUserContentArray = v
	}
	return OpenAIMessageParamUnion{OfUser: &user}
}

func OpenAIMessageParamOfSystem[T string | []ContentPartTextParam](content T) OpenAIMessageParamUnion {
	var system OpenAIMessageParamSystem
	switch v := any(content).(type) {
	case string:
		system.Content.OfString = param.NewOpt(v)
	case []ContentPartTextParam:
		system.Content.OfContentPartTextArray = v
	}
	return OpenAIMessageParamUnion{OfSystem: &system}
}

func OpenAIMessageParamOfTool[T string | []ContentPartTextParam](content T, toolCallID string) OpenAIMessageParamUnion {
	var tool OpenAIMessageParamTool
	switch v := any(content).(type) {
	case string:
		tool.Content.OfString = param.NewOpt(v)
	case []ContentPartTextParam:
		tool.Content.OfContentPartTextArray = v
	}
	tool.ToolCallID = toolCallID
	return OpenAIMessageParamUnion{OfTool: &tool}
}

func OpenAIMessageParamOfDeveloper[T string | []ContentPartTextParam](content T) OpenAIMessageParamUnion {
	var developer OpenAIMessageParamDeveloper
	switch v := any(content).(type) {
	case string:
		developer.Content.OfString = param.NewOpt(v)
	case []ContentPartTextParam:
		developer.Content.OfContentPartTextArray = v
	}
	return OpenAIMessageParamUnion{OfDeveloper: &developer}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIMessageParamUnion struct {
	OfUser      *OpenAIMessageParamUser      `json:",omitzero,inline"`
	OfSystem    *OpenAIMessageParamSystem    `json:",omitzero,inline"`
	OfAssistant *OpenAIMessageParamAssistant `json:",omitzero,inline"`
	OfTool      *OpenAIMessageParamTool      `json:",omitzero,inline"`
	OfDeveloper *OpenAIMessageParamDeveloper `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIMessageParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUser,
		u.OfSystem,
		u.OfAssistant,
		u.OfTool,
		u.OfDeveloper)
}
func (u *OpenAIMessageParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIMessageParamUnion) asAny() any {
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
func (u OpenAIMessageParamUnion) GetToolCalls() []ChatCompletionToolCallParam {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIMessageParamUnion) GetToolCallID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.ToolCallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIMessageParamUnion) GetRole() *string {
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
func (u OpenAIMessageParamUnion) GetName() *string {
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
func (u OpenAIMessageParamUnion) GetContent() (res openaiMessageParamUnionContent) {
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
// [_[]OpenAIMessageParamUserContentArrayItemUnion], [_[]ContentPartTextParam]
type openaiMessageParamUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]llamastackclient.OpenAIMessageParamUserContentArrayItemUnion:
//	case *[]llamastackclient.ContentPartTextParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u openaiMessageParamUnionContent) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[OpenAIMessageParamUnion](
		"role",
		apijson.Discriminator[OpenAIMessageParamUser]("user"),
		apijson.Discriminator[OpenAIMessageParamSystem]("system"),
		apijson.Discriminator[OpenAIMessageParamAssistant]("assistant"),
		apijson.Discriminator[OpenAIMessageParamTool]("tool"),
		apijson.Discriminator[OpenAIMessageParamDeveloper]("developer"),
	)
}

// A message from the user in an OpenAI-compatible chat completion request.
//
// The properties Content, Role are required.
type OpenAIMessageParamUser struct {
	// The content of the message, which can include text and other media
	Content OpenAIMessageParamUserContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the user message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "user" to identify this as a user message
	//
	// This field can be elided, and will marshal its zero value as "user".
	Role constant.User `json:"role,required"`
	paramObj
}

func (r OpenAIMessageParamUser) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIMessageParamUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIMessageParamUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIMessageParamUserContentUnion struct {
	OfString                        param.Opt[string]                             `json:",omitzero,inline"`
	OfOpenAIMessageUserContentArray []OpenAIMessageParamUserContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIMessageParamUserContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIMessageUserContentArray)
}
func (u *OpenAIMessageParamUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIMessageParamUserContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAIMessageUserContentArray) {
		return &u.OfOpenAIMessageUserContentArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIMessageParamUserContentArrayItemUnion struct {
	OfText     *ContentPartTextParam                           `json:",omitzero,inline"`
	OfImageURL *OpenAIMessageParamUserContentArrayItemImageURL `json:",omitzero,inline"`
	OfFile     *OpenAIMessageParamUserContentArrayItemFile     `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIMessageParamUserContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL, u.OfFile)
}
func (u *OpenAIMessageParamUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIMessageParamUserContentArrayItemUnion) asAny() any {
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
func (u OpenAIMessageParamUserContentArrayItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIMessageParamUserContentArrayItemUnion) GetImageURL() *OpenAIMessageParamUserContentArrayItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIMessageParamUserContentArrayItemUnion) GetFile() *OpenAIMessageParamUserContentArrayItemFileFile {
	if vt := u.OfFile; vt != nil {
		return &vt.File
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIMessageParamUserContentArrayItemUnion) GetType() *string {
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
	apijson.RegisterUnion[OpenAIMessageParamUserContentArrayItemUnion](
		"type",
		apijson.Discriminator[ContentPartTextParam]("text"),
		apijson.Discriminator[OpenAIMessageParamUserContentArrayItemImageURL]("image_url"),
		apijson.Discriminator[OpenAIMessageParamUserContentArrayItemFile]("file"),
	)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type OpenAIMessageParamUserContentArrayItemImageURL struct {
	// Image URL specification and processing details
	ImageURL OpenAIMessageParamUserContentArrayItemImageURLImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r OpenAIMessageParamUserContentArrayItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIMessageParamUserContentArrayItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIMessageParamUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type OpenAIMessageParamUserContentArrayItemImageURLImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r OpenAIMessageParamUserContentArrayItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIMessageParamUserContentArrayItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIMessageParamUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties File, Type are required.
type OpenAIMessageParamUserContentArrayItemFile struct {
	File OpenAIMessageParamUserContentArrayItemFileFile `json:"file,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "file".
	Type constant.File `json:"type,required"`
	paramObj
}

func (r OpenAIMessageParamUserContentArrayItemFile) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIMessageParamUserContentArrayItemFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIMessageParamUserContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIMessageParamUserContentArrayItemFileFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r OpenAIMessageParamUserContentArrayItemFileFile) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIMessageParamUserContentArrayItemFileFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIMessageParamUserContentArrayItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
//
// The properties Content, Role are required.
type OpenAIMessageParamSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content OpenAIMessageParamSystemContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the system message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "system" to identify this as a system message
	//
	// This field can be elided, and will marshal its zero value as "system".
	Role constant.System `json:"role,required"`
	paramObj
}

func (r OpenAIMessageParamSystem) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIMessageParamSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIMessageParamSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIMessageParamSystemContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIMessageParamSystemContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *OpenAIMessageParamSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIMessageParamSystemContentUnion) asAny() any {
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
type OpenAIMessageParamAssistant struct {
	// (Optional) The name of the assistant message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// The content of the model's response
	Content OpenAIMessageParamAssistantContentUnion `json:"content,omitzero"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []ChatCompletionToolCallParam `json:"tool_calls,omitzero"`
	// Must be "assistant" to identify this as the model's response
	//
	// This field can be elided, and will marshal its zero value as "assistant".
	Role constant.Assistant `json:"role,required"`
	paramObj
}

func (r OpenAIMessageParamAssistant) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIMessageParamAssistant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIMessageParamAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIMessageParamAssistantContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIMessageParamAssistantContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *OpenAIMessageParamAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIMessageParamAssistantContentUnion) asAny() any {
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
type OpenAIMessageParamTool struct {
	// The response content from the tool
	Content OpenAIMessageParamToolContentUnion `json:"content,omitzero,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// Must be "tool" to identify this as a tool response
	//
	// This field can be elided, and will marshal its zero value as "tool".
	Role constant.Tool `json:"role,required"`
	paramObj
}

func (r OpenAIMessageParamTool) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIMessageParamTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIMessageParamTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIMessageParamToolContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIMessageParamToolContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *OpenAIMessageParamToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIMessageParamToolContentUnion) asAny() any {
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
type OpenAIMessageParamDeveloper struct {
	// The content of the developer message
	Content OpenAIMessageParamDeveloperContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the developer message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "developer" to identify this as a developer message
	//
	// This field can be elided, and will marshal its zero value as "developer".
	Role constant.Developer `json:"role,required"`
	paramObj
}

func (r OpenAIMessageParamDeveloper) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIMessageParamDeveloper
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIMessageParamDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIMessageParamDeveloperContentUnion struct {
	OfString               param.Opt[string]      `json:",omitzero,inline"`
	OfContentPartTextArray []ContentPartTextParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIMessageParamDeveloperContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfContentPartTextArray)
}
func (u *OpenAIMessageParamDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIMessageParamDeveloperContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfContentPartTextArray) {
		return &u.OfContentPartTextArray
	}
	return nil
}

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

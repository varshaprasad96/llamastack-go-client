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

// OpenAIV1Service contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIV1Service] method instead.
type OpenAIV1Service struct {
	Options      []option.RequestOption
	Responses    OpenAIV1ResponseService
	Chat         OpenAIV1ChatService
	VectorStores OpenAIV1VectorStoreService
	Files        OpenAIV1FileService
}

// NewOpenAIV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOpenAIV1Service(opts ...option.RequestOption) (r OpenAIV1Service) {
	r = OpenAIV1Service{}
	r.Options = opts
	r.Responses = NewOpenAIV1ResponseService(opts...)
	r.Chat = NewOpenAIV1ChatService(opts...)
	r.VectorStores = NewOpenAIV1VectorStoreService(opts...)
	r.Files = NewOpenAIV1FileService(opts...)
	return
}

// Generate an OpenAI-compatible completion for the given prompt using the
// specified model.
func (r *OpenAIV1Service) Completions(ctx context.Context, body OpenAIV1CompletionsParams, opts ...option.RequestOption) (res *OpenAiv1CompletionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate OpenAI-compatible embeddings for the given input using the specified
// model.
func (r *OpenAIV1Service) Embeddings(ctx context.Context, body OpenAIV1EmbeddingsParams, opts ...option.RequestOption) (res *OpenAiv1EmbeddingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Classify if text violates OpenAI's Content Policy.
func (r *OpenAIV1Service) Moderations(ctx context.Context, body OpenAIV1ModerationsParams, opts ...option.RequestOption) (res *OpenAiv1ModerationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/moderations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List models using the OpenAI API.
func (r *OpenAIV1Service) GetModels(ctx context.Context, opts ...option.RequestOption) (res *OpenAiv1GetModelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response from an OpenAI-compatible completion request.
type OpenAiv1CompletionsResponse struct {
	ID      string                              `json:"id,required"`
	Choices []OpenAiv1CompletionsResponseChoice `json:"choices,required"`
	Created int64                               `json:"created,required"`
	Model   string                              `json:"model,required"`
	Object  constant.TextCompletion             `json:"object,required"`
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
func (r OpenAiv1CompletionsResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1CompletionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A choice from an OpenAI-compatible completion response.
type OpenAiv1CompletionsResponseChoice struct {
	FinishReason string `json:"finish_reason,required"`
	Index        int64  `json:"index,required"`
	Text         string `json:"text,required"`
	// The log probabilities for the tokens in the message from an OpenAI-compatible
	// chat completion response.
	Logprobs ChoiceLogprobs `json:"logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Index        respjson.Field
		Text         respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1CompletionsResponseChoice) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1CompletionsResponseChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from an OpenAI-compatible embeddings request.
type OpenAiv1EmbeddingsResponse struct {
	// List of embedding data objects
	Data []OpenAiv1EmbeddingsResponseData `json:"data,required"`
	// The model that was used to generate the embeddings
	Model string `json:"model,required"`
	// The object type, which will be "list"
	Object constant.List `json:"object,required"`
	// Usage information
	Usage OpenAiv1EmbeddingsResponseUsage `json:"usage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1EmbeddingsResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1EmbeddingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single embedding data object from an OpenAI-compatible embeddings response.
type OpenAiv1EmbeddingsResponseData struct {
	// The embedding vector as a list of floats (when encoding_format="float") or as a
	// base64-encoded string (when encoding_format="base64")
	Embedding OpenAiv1EmbeddingsResponseDataEmbeddingUnion `json:"embedding,required"`
	// The index of the embedding in the input list
	Index int64 `json:"index,required"`
	// The object type, which will be "embedding"
	Object constant.Embedding `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embedding   respjson.Field
		Index       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1EmbeddingsResponseData) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1EmbeddingsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAiv1EmbeddingsResponseDataEmbeddingUnion contains all possible properties
// and values from [[]float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloatArray OfString]
type OpenAiv1EmbeddingsResponseDataEmbeddingUnion struct {
	// This field will be present if the value is a [[]float64] instead of an object.
	OfFloatArray []float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloatArray respjson.Field
		OfString     respjson.Field
		raw          string
	} `json:"-"`
}

func (u OpenAiv1EmbeddingsResponseDataEmbeddingUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1EmbeddingsResponseDataEmbeddingUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAiv1EmbeddingsResponseDataEmbeddingUnion) RawJSON() string { return u.JSON.raw }

func (r *OpenAiv1EmbeddingsResponseDataEmbeddingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage information
type OpenAiv1EmbeddingsResponseUsage struct {
	// The number of tokens in the input
	PromptTokens int64 `json:"prompt_tokens,required"`
	// The total number of tokens used
	TotalTokens int64 `json:"total_tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PromptTokens respjson.Field
		TotalTokens  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1EmbeddingsResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1EmbeddingsResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAiv1ModerationsResponse struct {
	// The unique identifier for the moderation request.
	ID string `json:"id"`
	// The model used for moderation.
	Model   string                              `json:"model"`
	Results []OpenAiv1ModerationsResponseResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Model       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ModerationsResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1ModerationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAiv1ModerationsResponseResult struct {
	// Categories of content that were flagged.
	Categories any `json:"categories"`
	// Scores for each category.
	CategoryScores any `json:"category_scores"`
	// Whether the content was flagged.
	Flagged bool `json:"flagged"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories     respjson.Field
		CategoryScores respjson.Field
		Flagged        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ModerationsResponseResult) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1ModerationsResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAiv1GetModelsResponse struct {
	Data []OpenAiv1GetModelsResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1GetModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1GetModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A model from OpenAI.
type OpenAiv1GetModelsResponseData struct {
	ID      string         `json:"id,required"`
	Created int64          `json:"created,required"`
	Object  constant.Model `json:"object,required"`
	OwnedBy string         `json:"owned_by,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Object      respjson.Field
		OwnedBy     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1GetModelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1GetModelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIV1CompletionsParams struct {
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	Model string `json:"model,required"`
	// The prompt to generate a completion for.
	Prompt OpenAIV1CompletionsParamsPromptUnion `json:"prompt,omitzero,required"`
	// (Optional) The number of completions to generate.
	BestOf param.Opt[int64] `json:"best_of,omitzero"`
	// (Optional) Whether to echo the prompt.
	Echo param.Opt[bool] `json:"echo,omitzero"`
	// (Optional) The penalty for repeated tokens.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// (Optional) The log probabilities to use.
	Logprobs param.Opt[bool] `json:"logprobs,omitzero"`
	// (Optional) The maximum number of tokens to generate.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// (Optional) The number of completions to generate.
	N param.Opt[int64] `json:"n,omitzero"`
	// (Optional) The penalty for repeated tokens.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	PromptLogprobs  param.Opt[int64]   `json:"prompt_logprobs,omitzero"`
	// (Optional) The seed to use.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// (Optional) Whether to stream the response.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// (Optional) The suffix that should be appended to the completion.
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// (Optional) The temperature to use.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// (Optional) The top p to use.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// (Optional) The user to use.
	User         param.Opt[string] `json:"user,omitzero"`
	GuidedChoice []string          `json:"guided_choice,omitzero"`
	// (Optional) The logit bias to use.
	LogitBias map[string]float64 `json:"logit_bias,omitzero"`
	// (Optional) The stop tokens to use.
	Stop OpenAIV1CompletionsParamsStopUnion `json:"stop,omitzero"`
	// (Optional) The stream options to use.
	StreamOptions map[string]OpenAIV1CompletionsParamsStreamOptionUnion `json:"stream_options,omitzero"`
	paramObj
}

func (r OpenAIV1CompletionsParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1CompletionsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1CompletionsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1CompletionsParamsPromptUnion struct {
	OfString          param.Opt[string] `json:",omitzero,inline"`
	OfStringArray     []string          `json:",omitzero,inline"`
	OfIntArray        []int64           `json:",omitzero,inline"`
	OfArrayOfIntArray [][]int64         `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1CompletionsParamsPromptUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfIntArray, u.OfArrayOfIntArray)
}
func (u *OpenAIV1CompletionsParamsPromptUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1CompletionsParamsPromptUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfIntArray) {
		return &u.OfIntArray
	} else if !param.IsOmitted(u.OfArrayOfIntArray) {
		return &u.OfArrayOfIntArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1CompletionsParamsStopUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1CompletionsParamsStopUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpenAIV1CompletionsParamsStopUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1CompletionsParamsStopUnion) asAny() any {
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
type OpenAIV1CompletionsParamsStreamOptionUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1CompletionsParamsStreamOptionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1CompletionsParamsStreamOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1CompletionsParamsStreamOptionUnion) asAny() any {
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

type OpenAIV1EmbeddingsParams struct {
	// Input text to embed, encoded as a string or array of strings. To embed multiple
	// inputs in a single request, pass an array of strings.
	Input OpenAIV1EmbeddingsParamsInputUnion `json:"input,omitzero,required"`
	// The identifier of the model to use. The model must be an embedding model
	// registered with Llama Stack and available via the /models endpoint.
	Model string `json:"model,required"`
	// (Optional) The number of dimensions the resulting output embeddings should have.
	// Only supported in text-embedding-3 and later models.
	Dimensions param.Opt[int64] `json:"dimensions,omitzero"`
	// (Optional) The format to return the embeddings in. Can be either "float" or
	// "base64". Defaults to "float".
	EncodingFormat param.Opt[string] `json:"encoding_format,omitzero"`
	// (Optional) A unique identifier representing your end-user, which can help OpenAI
	// to monitor and detect abuse.
	User param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r OpenAIV1EmbeddingsParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1EmbeddingsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1EmbeddingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1EmbeddingsParamsInputUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1EmbeddingsParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpenAIV1EmbeddingsParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1EmbeddingsParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type OpenAIV1ModerationsParams struct {
	// The input text to classify.
	Input OpenAIV1ModerationsParamsInputUnion `json:"input,omitzero,required"`
	// The model to use for moderation.
	Model param.Opt[string] `json:"model,omitzero"`
	paramObj
}

func (r OpenAIV1ModerationsParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ModerationsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ModerationsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ModerationsParamsInputUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ModerationsParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpenAIV1ModerationsParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ModerationsParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

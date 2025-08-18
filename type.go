// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
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

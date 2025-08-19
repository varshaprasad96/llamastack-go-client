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

type CompletionResponse struct {
	Completion string `json:"completion,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completion  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *CompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenLogProbs struct {
	Token       string                    `json:"token,required"`
	Logprob     float64                   `json:"logprob,required"`
	TopLogprobs []TokenLogProbsTopLogprob `json:"top_logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenLogProbs) RawJSON() string { return r.JSON.raw }
func (r *TokenLogProbs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenLogProbsTopLogprob struct {
	Token   string  `json:"token"`
	Logprob float64 `json:"logprob"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenLogProbsTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *TokenLogProbsTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

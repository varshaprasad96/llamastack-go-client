// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/varshaprasad96/llamastack-go-client/option"
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

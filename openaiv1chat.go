// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient

import (
	"github.com/stainless-sdks/llamastack-go-client-go/option"
)

// OpenAIV1ChatService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIV1ChatService] method instead.
type OpenAIV1ChatService struct {
	Options     []option.RequestOption
	Completions OpenAIV1ChatCompletionService
}

// NewOpenAIV1ChatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOpenAIV1ChatService(opts ...option.RequestOption) (r OpenAIV1ChatService) {
	r = OpenAIV1ChatService{}
	r.Options = opts
	r.Completions = NewOpenAIV1ChatCompletionService(opts...)
	return
}

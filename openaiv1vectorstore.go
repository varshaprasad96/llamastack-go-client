// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/varshaprasad96/llamastack-go-client/option"
)

// OpenAIV1VectorStoreService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIV1VectorStoreService] method instead.
type OpenAIV1VectorStoreService struct {
	Options []option.RequestOption
	Files   OpenAIV1VectorStoreFileService
}

// NewOpenAIV1VectorStoreService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOpenAIV1VectorStoreService(opts ...option.RequestOption) (r OpenAIV1VectorStoreService) {
	r = OpenAIV1VectorStoreService{}
	r.Options = opts
	r.Files = NewOpenAIV1VectorStoreFileService(opts...)
	return
}

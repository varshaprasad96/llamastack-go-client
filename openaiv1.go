// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/varshaprasad96/llamastack-go-client/option"
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

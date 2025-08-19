// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"

	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
)

// OpenAIV1ResponseService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIV1ResponseService] method instead.
type OpenAIV1ResponseService struct {
	Options []option.RequestOption
}

// NewOpenAIV1ResponseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOpenAIV1ResponseService(opts ...option.RequestOption) (r OpenAIV1ResponseService) {
	r = OpenAIV1ResponseService{}
	r.Options = opts
	return
}

// List all OpenAI response objects.
func (r *OpenAIV1ResponseService) List(ctx context.Context, opts ...option.RequestOption) (res *OpenAiv1ResponseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OpenAiv1ResponseListResponse = any

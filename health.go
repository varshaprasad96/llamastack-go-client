// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/llamastack-go-client-go/internal/apijson"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/respjson"
)

// HealthService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.Options = opts
	return
}

// Get the current health status of the service.
func (r *HealthService) Get(ctx context.Context, opts ...option.RequestOption) (res *HealthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Health status information for the service.
type HealthGetResponse struct {
	// Current health status of the service
	//
	// Any of "OK", "Error", "Not Implemented".
	Status HealthGetResponseStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthGetResponse) RawJSON() string { return r.JSON.raw }
func (r *HealthGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current health status of the service
type HealthGetResponseStatus string

const (
	HealthGetResponseStatusOk             HealthGetResponseStatus = "OK"
	HealthGetResponseStatusError          HealthGetResponseStatus = "Error"
	HealthGetResponseStatusNotImplemented HealthGetResponseStatus = "Not Implemented"
)

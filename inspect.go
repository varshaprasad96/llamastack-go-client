// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// InspectService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInspectService] method instead.
type InspectService struct {
	Options []option.RequestOption
}

// NewInspectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInspectService(opts ...option.RequestOption) (r InspectService) {
	r = InspectService{}
	r.Options = opts
	return
}

// List all available API routes with their methods and implementing providers.
func (r *InspectService) ListRoutes(ctx context.Context, opts ...option.RequestOption) (res *InspectListRoutesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inspect/routes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response containing a list of all available API routes.
type InspectListRoutesResponse struct {
	// List of available route information objects
	Data []InspectListRoutesResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InspectListRoutesResponse) RawJSON() string { return r.JSON.raw }
func (r *InspectListRoutesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about an API route including its path, method, and implementing
// providers.
type InspectListRoutesResponseData struct {
	// HTTP method for the route
	Method string `json:"method,required"`
	// List of provider types that implement this route
	ProviderTypes []string `json:"provider_types,required"`
	// The API endpoint path
	Route string `json:"route,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method        respjson.Field
		ProviderTypes respjson.Field
		Route         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InspectListRoutesResponseData) RawJSON() string { return r.JSON.raw }
func (r *InspectListRoutesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

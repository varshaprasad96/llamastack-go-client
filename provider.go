// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// ProviderService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderService] method instead.
type ProviderService struct {
	Options []option.RequestOption
}

// NewProviderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderService(opts ...option.RequestOption) (r ProviderService) {
	r = ProviderService{}
	r.Options = opts
	return
}

// Get detailed information about a specific provider.
func (r *ProviderService) Get(ctx context.Context, providerID string, opts ...option.RequestOption) (res *ProviderInfo, err error) {
	opts = append(r.Options[:], opts...)
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/%s", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all available providers.
func (r *ProviderService) List(ctx context.Context, opts ...option.RequestOption) (res *ProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Information about a registered provider including its configuration and health
// status.
type ProviderInfo struct {
	// The API name this provider implements
	API string `json:"api,required"`
	// Configuration parameters for the provider
	Config map[string]ProviderInfoConfigUnion `json:"config,required"`
	// Current health status of the provider
	Health map[string]ProviderInfoHealthUnion `json:"health,required"`
	// Unique identifier for the provider
	ProviderID string `json:"provider_id,required"`
	// The type of provider implementation
	ProviderType string `json:"provider_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		API          respjson.Field
		Config       respjson.Field
		Health       respjson.Field
		ProviderID   respjson.Field
		ProviderType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderInfo) RawJSON() string { return r.JSON.raw }
func (r *ProviderInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProviderInfoConfigUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ProviderInfoConfigUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u ProviderInfoConfigUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoConfigUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoConfigUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoConfigUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProviderInfoConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *ProviderInfoConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProviderInfoHealthUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ProviderInfoHealthUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u ProviderInfoHealthUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoHealthUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoHealthUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoHealthUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProviderInfoHealthUnion) RawJSON() string { return u.JSON.raw }

func (r *ProviderInfoHealthUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a list of all available providers.
type ProviderListResponse struct {
	// List of provider information objects
	Data []ProviderInfo `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProviderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

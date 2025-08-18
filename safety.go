// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stainless-sdks/llamastack-go-client-go/internal/apijson"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/param"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/respjson"
)

// SafetyService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSafetyService] method instead.
type SafetyService struct {
	Options []option.RequestOption
}

// NewSafetyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSafetyService(opts ...option.RequestOption) (r SafetyService) {
	r = SafetyService{}
	r.Options = opts
	return
}

// Run a shield.
func (r *SafetyService) RunShield(ctx context.Context, body SafetyRunShieldParams, opts ...option.RequestOption) (res *SafetyRunShieldResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/safety/run-shield"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Details of a safety violation detected by content moderation.
type SafetyViolation struct {
	// Additional metadata including specific violation codes for debugging and
	// telemetry
	Metadata map[string]SafetyViolationMetadataUnion `json:"metadata,required"`
	// Severity level of the violation
	//
	// Any of "info", "warn", "error".
	ViolationLevel SafetyViolationViolationLevel `json:"violation_level,required"`
	// (Optional) Message to convey to the user about the violation
	UserMessage string `json:"user_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata       respjson.Field
		ViolationLevel respjson.Field
		UserMessage    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SafetyViolation) RawJSON() string { return r.JSON.raw }
func (r *SafetyViolation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SafetyViolationMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type SafetyViolationMetadataUnion struct {
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

func (u SafetyViolationMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SafetyViolationMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SafetyViolationMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SafetyViolationMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SafetyViolationMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *SafetyViolationMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Severity level of the violation
type SafetyViolationViolationLevel string

const (
	SafetyViolationViolationLevelInfo  SafetyViolationViolationLevel = "info"
	SafetyViolationViolationLevelWarn  SafetyViolationViolationLevel = "warn"
	SafetyViolationViolationLevelError SafetyViolationViolationLevel = "error"
)

// Response from running a safety shield.
type SafetyRunShieldResponse struct {
	// (Optional) Safety violation detected by the shield, if any
	Violation SafetyViolation `json:"violation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Violation   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SafetyRunShieldResponse) RawJSON() string { return r.JSON.raw }
func (r *SafetyRunShieldResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SafetyRunShieldParams struct {
	// The messages to run the shield on.
	Messages []MessageUnionParam `json:"messages,omitzero,required"`
	// The parameters of the shield.
	Params map[string]SafetyRunShieldParamsParamUnion `json:"params,omitzero,required"`
	// The identifier of the shield to run.
	ShieldID string `json:"shield_id,required"`
	paramObj
}

func (r SafetyRunShieldParams) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsParamUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *SafetyRunShieldParamsParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsParamUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/apiquery"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
	"github.com/varshaprasad96/llamastack-go-client/shared/constant"
)

// ToolService contains methods and other services that help with interacting with
// the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolService] method instead.
type ToolService struct {
	Options []option.RequestOption
}

// NewToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolService(opts ...option.RequestOption) (r ToolService) {
	r = ToolService{}
	r.Options = opts
	return
}

// Get a tool by its name.
func (r *ToolService) Get(ctx context.Context, toolName string, opts ...option.RequestOption) (res *Tool, err error) {
	opts = append(r.Options[:], opts...)
	if toolName == "" {
		err = errors.New("missing required tool_name parameter")
		return
	}
	path := fmt.Sprintf("v1/tools/%s", toolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List tools with optional tool group.
func (r *ToolService) List(ctx context.Context, query ToolListParams, opts ...option.RequestOption) (res *ToolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A tool that can be invoked by agents.
type Tool struct {
	// Human-readable description of what the tool does
	Description string `json:"description,required"`
	Identifier  string `json:"identifier,required"`
	// List of parameters this tool accepts
	Parameters []ToolParameterResp `json:"parameters,required"`
	ProviderID string              `json:"provider_id,required"`
	// ID of the tool group this tool belongs to
	ToolgroupID string `json:"toolgroup_id,required"`
	// Type of resource, always 'tool'
	Type constant.Tool `json:"type,required"`
	// (Optional) Additional metadata about the tool
	Metadata           map[string]ToolMetadataUnion `json:"metadata"`
	ProviderResourceID string                       `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description        respjson.Field
		Identifier         respjson.Field
		Parameters         respjson.Field
		ProviderID         respjson.Field
		ToolgroupID        respjson.Field
		Type               respjson.Field
		Metadata           respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tool) RawJSON() string { return r.JSON.raw }
func (r *Tool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolMetadataUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolMetadataUnion struct {
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

func (u ToolMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter definition for a tool.
type ToolParameterResp struct {
	// Human-readable description of what the parameter does
	Description string `json:"description,required"`
	// Name of the parameter
	Name string `json:"name,required"`
	// Type of the parameter (e.g., string, integer)
	ParameterType string `json:"parameter_type,required"`
	// Whether this parameter is required for tool invocation
	Required bool `json:"required,required"`
	// (Optional) Default value for the parameter if not provided
	Default ToolParameterDefaultUnionResp `json:"default,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description   respjson.Field
		Name          respjson.Field
		ParameterType respjson.Field
		Required      respjson.Field
		Default       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolParameterResp) RawJSON() string { return r.JSON.raw }
func (r *ToolParameterResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolParameterResp to a ToolParameter.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolParameter.Overrides()
func (r ToolParameterResp) ToParam() ToolParameter {
	return param.Override[ToolParameter](json.RawMessage(r.RawJSON()))
}

// ToolParameterDefaultUnionResp contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolParameterDefaultUnionResp struct {
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

func (u ToolParameterDefaultUnionResp) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolParameterDefaultUnionResp) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolParameterDefaultUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolParameterDefaultUnionResp) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolParameterDefaultUnionResp) RawJSON() string { return u.JSON.raw }

func (r *ToolParameterDefaultUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameter definition for a tool.
//
// The properties Description, Name, ParameterType, Required are required.
type ToolParameter struct {
	// Human-readable description of what the parameter does
	Description string `json:"description,required"`
	// Name of the parameter
	Name string `json:"name,required"`
	// Type of the parameter (e.g., string, integer)
	ParameterType string `json:"parameter_type,required"`
	// Whether this parameter is required for tool invocation
	Required bool `json:"required,required"`
	// (Optional) Default value for the parameter if not provided
	Default ToolParameterDefaultUnion `json:"default,omitzero"`
	paramObj
}

func (r ToolParameter) MarshalJSON() (data []byte, err error) {
	type shadow ToolParameter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolParameterDefaultUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ToolParameterDefaultUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ToolParameterDefaultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolParameterDefaultUnion) asAny() any {
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

// Response containing a list of tools.
type ToolListResponse struct {
	// List of tools
	Data []Tool `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolListResponse) RawJSON() string { return r.JSON.raw }
func (r *ToolListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolListParams struct {
	// The ID of the tool group to list tools for.
	ToolgroupID param.Opt[string] `query:"toolgroup_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolListParams]'s query parameters as `url.Values`.
func (r ToolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

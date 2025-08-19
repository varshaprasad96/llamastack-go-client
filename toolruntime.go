// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/apiquery"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// ToolRuntimeService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolRuntimeService] method instead.
type ToolRuntimeService struct {
	Options []option.RequestOption
	RagTool ToolRuntimeRagToolService
}

// NewToolRuntimeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolRuntimeService(opts ...option.RequestOption) (r ToolRuntimeService) {
	r = ToolRuntimeService{}
	r.Options = opts
	r.RagTool = NewToolRuntimeRagToolService(opts...)
	return
}

// Run a tool with the given arguments.
func (r *ToolRuntimeService) Invoke(ctx context.Context, body ToolRuntimeInvokeParams, opts ...option.RequestOption) (res *ToolRuntimeInvokeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tool-runtime/invoke"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all tools in the runtime.
func (r *ToolRuntimeService) ListTools(ctx context.Context, query ToolRuntimeListToolsParams, opts ...option.RequestOption) (res *ToolRuntimeListToolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tool-runtime/list-tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type URL struct {
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URL) RawJSON() string { return r.JSON.raw }
func (r *URL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this URL to a URLParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// URLParam.Overrides()
func (r URL) ToParam() URLParam {
	return param.Override[URLParam](json.RawMessage(r.RawJSON()))
}

// The property Uri is required.
type URLParam struct {
	Uri string `json:"uri,required"`
	paramObj
}

func (r URLParam) MarshalJSON() (data []byte, err error) {
	type shadow URLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolRuntimeInvokeResponse struct {
	// A image content item
	Content      InterleavedContentUnion                           `json:"content"`
	ErrorCode    int64                                             `json:"error_code"`
	ErrorMessage string                                            `json:"error_message"`
	Metadata     map[string]ToolRuntimeInvokeResponseMetadataUnion `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content      respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		Metadata     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolRuntimeInvokeResponse) RawJSON() string { return r.JSON.raw }
func (r *ToolRuntimeInvokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolRuntimeInvokeResponseMetadataUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolRuntimeInvokeResponseMetadataUnion struct {
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

func (u ToolRuntimeInvokeResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolRuntimeInvokeResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolRuntimeInvokeResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolRuntimeInvokeResponseMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolRuntimeInvokeResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolRuntimeInvokeResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolRuntimeListToolsResponse struct {
	Data []ToolDef `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolRuntimeListToolsResponse) RawJSON() string { return r.JSON.raw }
func (r *ToolRuntimeListToolsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolRuntimeInvokeParams struct {
	// A dictionary of arguments to pass to the tool.
	Kwargs map[string]ToolRuntimeInvokeParamsKwargUnion `json:"kwargs,omitzero,required"`
	// The name of the tool to invoke.
	ToolName string `json:"tool_name,required"`
	paramObj
}

func (r ToolRuntimeInvokeParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeInvokeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeInvokeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolRuntimeInvokeParamsKwargUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ToolRuntimeInvokeParamsKwargUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ToolRuntimeInvokeParamsKwargUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolRuntimeInvokeParamsKwargUnion) asAny() any {
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

type ToolRuntimeListToolsParams struct {
	// The ID of the tool group to list tools for.
	ToolGroupID param.Opt[string] `query:"tool_group_id,omitzero" json:"-"`
	// The MCP endpoint to use for the tool group.
	McpEndpoint URLParam `query:"mcp_endpoint,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolRuntimeListToolsParams]'s query parameters as
// `url.Values`.
func (r ToolRuntimeListToolsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

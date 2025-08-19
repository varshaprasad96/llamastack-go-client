// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// TelemetrySpanService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTelemetrySpanService] method instead.
type TelemetrySpanService struct {
	Options []option.RequestOption
}

// NewTelemetrySpanService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTelemetrySpanService(opts ...option.RequestOption) (r TelemetrySpanService) {
	r = TelemetrySpanService{}
	r.Options = opts
	return
}

// Save spans to a dataset.
func (r *TelemetrySpanService) Export(ctx context.Context, body TelemetrySpanExportParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/telemetry/spans/export"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Query spans.
func (r *TelemetrySpanService) Query(ctx context.Context, body TelemetrySpanQueryParams, opts ...option.RequestOption) (res *TelemetrySpanQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/telemetry/spans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a span tree by its ID.
func (r *TelemetrySpanService) GetTree(ctx context.Context, spanID string, body TelemetrySpanGetTreeParams, opts ...option.RequestOption) (res *TelemetrySpanGetTreeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if spanID == "" {
		err = errors.New("missing required span_id parameter")
		return
	}
	path := fmt.Sprintf("v1/telemetry/spans/%s/tree", spanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The properties Key, Op, Value are required.
type QueryConditionParam struct {
	Value QueryConditionValueUnionParam `json:"value,omitzero,required"`
	Key   string                        `json:"key,required"`
	// Any of "eq", "ne", "gt", "lt".
	Op QueryConditionOp `json:"op,omitzero,required"`
	paramObj
}

func (r QueryConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryConditionOp string

const (
	QueryConditionOpEq QueryConditionOp = "eq"
	QueryConditionOpNe QueryConditionOp = "ne"
	QueryConditionOpGt QueryConditionOp = "gt"
	QueryConditionOpLt QueryConditionOp = "lt"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type QueryConditionValueUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u QueryConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *QueryConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *QueryConditionValueUnionParam) asAny() any {
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

type TelemetrySpanQueryResponse struct {
	Data []Span `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetrySpanQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *TelemetrySpanQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetrySpanGetTreeResponse struct {
	Data map[string]TelemetrySpanGetTreeResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetrySpanGetTreeResponse) RawJSON() string { return r.JSON.raw }
func (r *TelemetrySpanGetTreeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetrySpanGetTreeResponseData struct {
	Name         string                                                    `json:"name,required"`
	SpanID       string                                                    `json:"span_id,required"`
	StartTime    time.Time                                                 `json:"start_time,required" format:"date-time"`
	TraceID      string                                                    `json:"trace_id,required"`
	Attributes   map[string]TelemetrySpanGetTreeResponseDataAttributeUnion `json:"attributes"`
	EndTime      time.Time                                                 `json:"end_time" format:"date-time"`
	ParentSpanID string                                                    `json:"parent_span_id"`
	// Any of "ok", "error".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		SpanID       respjson.Field
		StartTime    respjson.Field
		TraceID      respjson.Field
		Attributes   respjson.Field
		EndTime      respjson.Field
		ParentSpanID respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetrySpanGetTreeResponseData) RawJSON() string { return r.JSON.raw }
func (r *TelemetrySpanGetTreeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TelemetrySpanGetTreeResponseDataAttributeUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type TelemetrySpanGetTreeResponseDataAttributeUnion struct {
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

func (u TelemetrySpanGetTreeResponseDataAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TelemetrySpanGetTreeResponseDataAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TelemetrySpanGetTreeResponseDataAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TelemetrySpanGetTreeResponseDataAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TelemetrySpanGetTreeResponseDataAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *TelemetrySpanGetTreeResponseDataAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetrySpanExportParams struct {
	// The attribute filters to apply to the spans.
	AttributeFilters []QueryConditionParam `json:"attribute_filters,omitzero,required"`
	// The attributes to save to the dataset.
	AttributesToSave []string `json:"attributes_to_save,omitzero,required"`
	// The ID of the dataset to save the spans to.
	DatasetID string `json:"dataset_id,required"`
	// The maximum depth of the tree.
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	paramObj
}

func (r TelemetrySpanExportParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetrySpanExportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetrySpanExportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetrySpanQueryParams struct {
	// The attribute filters to apply to the spans.
	AttributeFilters []QueryConditionParam `json:"attribute_filters,omitzero,required"`
	// The attributes to return in the spans.
	AttributesToReturn []string `json:"attributes_to_return,omitzero,required"`
	// The maximum depth of the tree.
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	paramObj
}

func (r TelemetrySpanQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetrySpanQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetrySpanQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetrySpanGetTreeParams struct {
	// The maximum depth of the tree.
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	// The attributes to return in the tree.
	AttributesToReturn []string `json:"attributes_to_return,omitzero"`
	paramObj
}

func (r TelemetrySpanGetTreeParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetrySpanGetTreeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetrySpanGetTreeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

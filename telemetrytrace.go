// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/llamastack-go-client-go/internal/apijson"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/param"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/respjson"
)

// TelemetryTraceService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTelemetryTraceService] method instead.
type TelemetryTraceService struct {
	Options []option.RequestOption
}

// NewTelemetryTraceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTelemetryTraceService(opts ...option.RequestOption) (r TelemetryTraceService) {
	r = TelemetryTraceService{}
	r.Options = opts
	return
}

// Query traces.
func (r *TelemetryTraceService) Query(ctx context.Context, body TelemetryTraceQueryParams, opts ...option.RequestOption) (res *TelemetryTraceQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/telemetry/traces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a span by its ID.
func (r *TelemetryTraceService) GetSpan(ctx context.Context, spanID string, query TelemetryTraceGetSpanParams, opts ...option.RequestOption) (res *Span, err error) {
	opts = append(r.Options[:], opts...)
	if query.TraceID == "" {
		err = errors.New("missing required trace_id parameter")
		return
	}
	if spanID == "" {
		err = errors.New("missing required span_id parameter")
		return
	}
	path := fmt.Sprintf("v1/telemetry/traces/%s/spans/%s", query.TraceID, spanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a trace by its ID.
func (r *TelemetryTraceService) GetTrace(ctx context.Context, traceID string, opts ...option.RequestOption) (res *Trace, err error) {
	opts = append(r.Options[:], opts...)
	if traceID == "" {
		err = errors.New("missing required trace_id parameter")
		return
	}
	path := fmt.Sprintf("v1/telemetry/traces/%s", traceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A span representing a single operation within a trace.
type Span struct {
	// Human-readable name describing the operation this span represents
	Name string `json:"name,required"`
	// Unique identifier for the span
	SpanID string `json:"span_id,required"`
	// Timestamp when the operation began
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// Unique identifier for the trace this span belongs to
	TraceID string `json:"trace_id,required"`
	// (Optional) Key-value pairs containing additional metadata about the span
	Attributes map[string]SpanAttributeUnion `json:"attributes"`
	// (Optional) Timestamp when the operation finished, if completed
	EndTime time.Time `json:"end_time" format:"date-time"`
	// (Optional) Unique identifier for the parent span, if this is a child span
	ParentSpanID string `json:"parent_span_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		SpanID       respjson.Field
		StartTime    respjson.Field
		TraceID      respjson.Field
		Attributes   respjson.Field
		EndTime      respjson.Field
		ParentSpanID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Span) RawJSON() string { return r.JSON.raw }
func (r *Span) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SpanAttributeUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type SpanAttributeUnion struct {
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

func (u SpanAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SpanAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SpanAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SpanAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SpanAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *SpanAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A trace representing the complete execution path of a request across multiple
// operations.
type Trace struct {
	// Unique identifier for the root span that started this trace
	RootSpanID string `json:"root_span_id,required"`
	// Timestamp when the trace began
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// Unique identifier for the trace
	TraceID string `json:"trace_id,required"`
	// (Optional) Timestamp when the trace finished, if completed
	EndTime time.Time `json:"end_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RootSpanID  respjson.Field
		StartTime   respjson.Field
		TraceID     respjson.Field
		EndTime     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Trace) RawJSON() string { return r.JSON.raw }
func (r *Trace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a list of traces.
type TelemetryTraceQueryResponse struct {
	// List of traces matching the query criteria
	Data []Trace `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryTraceQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *TelemetryTraceQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetryTraceQueryParams struct {
	// The limit of traces to return.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The offset of the traces to return.
	Offset param.Opt[int64] `json:"offset,omitzero"`
	// The attribute filters to apply to the traces.
	AttributeFilters []QueryConditionParam `json:"attribute_filters,omitzero"`
	// The order by of the traces to return.
	OrderBy []string `json:"order_by,omitzero"`
	paramObj
}

func (r TelemetryTraceQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryTraceQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryTraceQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetryTraceGetSpanParams struct {
	TraceID string `path:"trace_id,required" json:"-"`
	paramObj
}

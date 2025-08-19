// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
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

// TelemetryService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTelemetryService] method instead.
type TelemetryService struct {
	Options []option.RequestOption
	Traces  TelemetryTraceService
	Spans   TelemetrySpanService
}

// NewTelemetryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTelemetryService(opts ...option.RequestOption) (r TelemetryService) {
	r = TelemetryService{}
	r.Options = opts
	r.Traces = NewTelemetryTraceService(opts...)
	r.Spans = NewTelemetrySpanService(opts...)
	return
}

// Log an event.
func (r *TelemetryService) LogEvent(ctx context.Context, body TelemetryLogEventParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/telemetry/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Query metrics.
func (r *TelemetryService) QueryMetric(ctx context.Context, metricName string, body TelemetryQueryMetricParams, opts ...option.RequestOption) (res *TelemetryQueryMetricResponse, err error) {
	opts = append(r.Options[:], opts...)
	if metricName == "" {
		err = errors.New("missing required metric_name parameter")
		return
	}
	path := fmt.Sprintf("v1/telemetry/metrics/%s", metricName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EventType string

const (
	EventTypeUnstructuredLog EventType = "unstructured_log"
	EventTypeStructuredLog   EventType = "structured_log"
	EventTypeMetric          EventType = "metric"
)

type StructuredLogType string

const (
	StructuredLogTypeSpanStart StructuredLogType = "span_start"
	StructuredLogTypeSpanEnd   StructuredLogType = "span_end"
)

type TelemetryQueryMetricResponse struct {
	Data []TelemetryQueryMetricResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryQueryMetricResponse) RawJSON() string { return r.JSON.raw }
func (r *TelemetryQueryMetricResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetryQueryMetricResponseData struct {
	Labels []TelemetryQueryMetricResponseDataLabel `json:"labels,required"`
	Metric string                                  `json:"metric,required"`
	Values []TelemetryQueryMetricResponseDataValue `json:"values,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Labels      respjson.Field
		Metric      respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryQueryMetricResponseData) RawJSON() string { return r.JSON.raw }
func (r *TelemetryQueryMetricResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetryQueryMetricResponseDataLabel struct {
	Name  string `json:"name,required"`
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryQueryMetricResponseDataLabel) RawJSON() string { return r.JSON.raw }
func (r *TelemetryQueryMetricResponseDataLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetryQueryMetricResponseDataValue struct {
	Timestamp int64   `json:"timestamp,required"`
	Value     float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryQueryMetricResponseDataValue) RawJSON() string { return r.JSON.raw }
func (r *TelemetryQueryMetricResponseDataValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetryLogEventParams struct {
	// The event to log.
	Event TelemetryLogEventParamsEventUnion `json:"event,omitzero,required"`
	// The time to live of the event.
	TtlSeconds int64 `json:"ttl_seconds,required"`
	paramObj
}

func (r TelemetryLogEventParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryLogEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryLogEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TelemetryLogEventParamsEventUnion struct {
	OfUnstructuredLogEvent *TelemetryLogEventParamsEventUnstructuredLogEvent `json:",omitzero,inline"`
	OfMetricEvent          *TelemetryLogEventParamsEventMetricEvent          `json:",omitzero,inline"`
	OfStructuredLogEvent   *TelemetryLogEventParamsEventStructuredLogEvent   `json:",omitzero,inline"`
	paramUnion
}

func (u TelemetryLogEventParamsEventUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUnstructuredLogEvent, u.OfMetricEvent, u.OfStructuredLogEvent)
}
func (u *TelemetryLogEventParamsEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TelemetryLogEventParamsEventUnion) asAny() any {
	if !param.IsOmitted(u.OfUnstructuredLogEvent) {
		return u.OfUnstructuredLogEvent
	} else if !param.IsOmitted(u.OfMetricEvent) {
		return u.OfMetricEvent
	} else if !param.IsOmitted(u.OfStructuredLogEvent) {
		return u.OfStructuredLogEvent
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventUnion) GetMessage() *string {
	if vt := u.OfUnstructuredLogEvent; vt != nil {
		return &vt.Message
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventUnion) GetSeverity() *string {
	if vt := u.OfUnstructuredLogEvent; vt != nil {
		return &vt.Severity
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventUnion) GetMetric() *string {
	if vt := u.OfMetricEvent; vt != nil {
		return &vt.Metric
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventUnion) GetUnit() *string {
	if vt := u.OfMetricEvent; vt != nil {
		return &vt.Unit
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventUnion) GetValue() *float64 {
	if vt := u.OfMetricEvent; vt != nil {
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventUnion) GetPayload() *TelemetryLogEventParamsEventStructuredLogEventPayloadUnion {
	if vt := u.OfStructuredLogEvent; vt != nil {
		return &vt.Payload
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventUnion) GetSpanID() *string {
	if vt := u.OfUnstructuredLogEvent; vt != nil {
		return (*string)(&vt.SpanID)
	} else if vt := u.OfMetricEvent; vt != nil {
		return (*string)(&vt.SpanID)
	} else if vt := u.OfStructuredLogEvent; vt != nil {
		return (*string)(&vt.SpanID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventUnion) GetTraceID() *string {
	if vt := u.OfUnstructuredLogEvent; vt != nil {
		return (*string)(&vt.TraceID)
	} else if vt := u.OfMetricEvent; vt != nil {
		return (*string)(&vt.TraceID)
	} else if vt := u.OfStructuredLogEvent; vt != nil {
		return (*string)(&vt.TraceID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventUnion) GetType() *string {
	if vt := u.OfUnstructuredLogEvent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMetricEvent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStructuredLogEvent; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's Timestamp property, if present.
func (u TelemetryLogEventParamsEventUnion) GetTimestamp() *time.Time {
	if vt := u.OfUnstructuredLogEvent; vt != nil {
		return &vt.Timestamp
	} else if vt := u.OfMetricEvent; vt != nil {
		return &vt.Timestamp
	} else if vt := u.OfStructuredLogEvent; vt != nil {
		return &vt.Timestamp
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u TelemetryLogEventParamsEventUnion) GetAttributes() (res telemetryLogEventParamsEventUnionAttributes) {
	if vt := u.OfUnstructuredLogEvent; vt != nil {
		res.any = &vt.Attributes
	} else if vt := u.OfMetricEvent; vt != nil {
		res.any = &vt.Attributes
	} else if vt := u.OfStructuredLogEvent; vt != nil {
		res.any = &vt.Attributes
	}
	return
}

// Can have the runtime types
// [*map[string]TelemetryLogEventParamsEventUnstructuredLogEventAttributeUnion],
// [*map[string]TelemetryLogEventParamsEventMetricEventAttributeUnion],
// [\*map[string]TelemetryLogEventParamsEventStructuredLogEventAttributeUnion]
type telemetryLogEventParamsEventUnionAttributes struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *map[string]llamastackclient.TelemetryLogEventParamsEventUnstructuredLogEventAttributeUnion:
//	case *map[string]llamastackclient.TelemetryLogEventParamsEventMetricEventAttributeUnion:
//	case *map[string]llamastackclient.TelemetryLogEventParamsEventStructuredLogEventAttributeUnion:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u telemetryLogEventParamsEventUnionAttributes) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[TelemetryLogEventParamsEventUnion](
		"type",
		apijson.Discriminator[TelemetryLogEventParamsEventUnstructuredLogEvent]("unstructured_log"),
		apijson.Discriminator[TelemetryLogEventParamsEventUnstructuredLogEvent]("structured_log"),
		apijson.Discriminator[TelemetryLogEventParamsEventUnstructuredLogEvent]("metric"),
		apijson.Discriminator[TelemetryLogEventParamsEventMetricEvent]("unstructured_log"),
		apijson.Discriminator[TelemetryLogEventParamsEventMetricEvent]("structured_log"),
		apijson.Discriminator[TelemetryLogEventParamsEventMetricEvent]("metric"),
		apijson.Discriminator[TelemetryLogEventParamsEventStructuredLogEvent]("unstructured_log"),
		apijson.Discriminator[TelemetryLogEventParamsEventStructuredLogEvent]("structured_log"),
		apijson.Discriminator[TelemetryLogEventParamsEventStructuredLogEvent]("metric"),
	)
}

// The properties Message, Severity, SpanID, Timestamp, TraceID, Type are required.
type TelemetryLogEventParamsEventUnstructuredLogEvent struct {
	Message string `json:"message,required"`
	// Any of "verbose", "debug", "info", "warn", "error", "critical".
	Severity  string    `json:"severity,omitzero,required"`
	SpanID    string    `json:"span_id,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	TraceID   string    `json:"trace_id,required"`
	// Any of "unstructured_log", "structured_log", "metric".
	Type       EventType                                                                 `json:"type,omitzero,required"`
	Attributes map[string]TelemetryLogEventParamsEventUnstructuredLogEventAttributeUnion `json:"attributes,omitzero"`
	paramObj
}

func (r TelemetryLogEventParamsEventUnstructuredLogEvent) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryLogEventParamsEventUnstructuredLogEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryLogEventParamsEventUnstructuredLogEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TelemetryLogEventParamsEventUnstructuredLogEvent](
		"severity", "verbose", "debug", "info", "warn", "error", "critical",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TelemetryLogEventParamsEventUnstructuredLogEventAttributeUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u TelemetryLogEventParamsEventUnstructuredLogEventAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *TelemetryLogEventParamsEventUnstructuredLogEventAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TelemetryLogEventParamsEventUnstructuredLogEventAttributeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The properties Metric, SpanID, Timestamp, TraceID, Type, Unit, Value are
// required.
type TelemetryLogEventParamsEventMetricEvent struct {
	Metric    string    `json:"metric,required"`
	SpanID    string    `json:"span_id,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	TraceID   string    `json:"trace_id,required"`
	// Any of "unstructured_log", "structured_log", "metric".
	Type       EventType                                                        `json:"type,omitzero,required"`
	Unit       string                                                           `json:"unit,required"`
	Value      float64                                                          `json:"value,required"`
	Attributes map[string]TelemetryLogEventParamsEventMetricEventAttributeUnion `json:"attributes,omitzero"`
	paramObj
}

func (r TelemetryLogEventParamsEventMetricEvent) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryLogEventParamsEventMetricEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryLogEventParamsEventMetricEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TelemetryLogEventParamsEventMetricEventAttributeUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u TelemetryLogEventParamsEventMetricEventAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *TelemetryLogEventParamsEventMetricEventAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TelemetryLogEventParamsEventMetricEventAttributeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The properties Payload, SpanID, Timestamp, TraceID, Type are required.
type TelemetryLogEventParamsEventStructuredLogEvent struct {
	Payload   TelemetryLogEventParamsEventStructuredLogEventPayloadUnion `json:"payload,omitzero,required"`
	SpanID    string                                                     `json:"span_id,required"`
	Timestamp time.Time                                                  `json:"timestamp,required" format:"date-time"`
	TraceID   string                                                     `json:"trace_id,required"`
	// Any of "unstructured_log", "structured_log", "metric".
	Type       EventType                                                               `json:"type,omitzero,required"`
	Attributes map[string]TelemetryLogEventParamsEventStructuredLogEventAttributeUnion `json:"attributes,omitzero"`
	paramObj
}

func (r TelemetryLogEventParamsEventStructuredLogEvent) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryLogEventParamsEventStructuredLogEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryLogEventParamsEventStructuredLogEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TelemetryLogEventParamsEventStructuredLogEventPayloadUnion struct {
	OfSpanStartPayload *TelemetryLogEventParamsEventStructuredLogEventPayloadSpanStartPayload `json:",omitzero,inline"`
	OfSpanEndPayload   *TelemetryLogEventParamsEventStructuredLogEventPayloadSpanEndPayload   `json:",omitzero,inline"`
	paramUnion
}

func (u TelemetryLogEventParamsEventStructuredLogEventPayloadUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSpanStartPayload, u.OfSpanEndPayload)
}
func (u *TelemetryLogEventParamsEventStructuredLogEventPayloadUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TelemetryLogEventParamsEventStructuredLogEventPayloadUnion) asAny() any {
	if !param.IsOmitted(u.OfSpanStartPayload) {
		return u.OfSpanStartPayload
	} else if !param.IsOmitted(u.OfSpanEndPayload) {
		return u.OfSpanEndPayload
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventStructuredLogEventPayloadUnion) GetName() *string {
	if vt := u.OfSpanStartPayload; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventStructuredLogEventPayloadUnion) GetParentSpanID() *string {
	if vt := u.OfSpanStartPayload; vt != nil && vt.ParentSpanID.Valid() {
		return &vt.ParentSpanID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventStructuredLogEventPayloadUnion) GetStatus() *string {
	if vt := u.OfSpanEndPayload; vt != nil {
		return &vt.Status
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TelemetryLogEventParamsEventStructuredLogEventPayloadUnion) GetType() *string {
	if vt := u.OfSpanStartPayload; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSpanEndPayload; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[TelemetryLogEventParamsEventStructuredLogEventPayloadUnion](
		"type",
		apijson.Discriminator[TelemetryLogEventParamsEventStructuredLogEventPayloadSpanStartPayload]("span_start"),
		apijson.Discriminator[TelemetryLogEventParamsEventStructuredLogEventPayloadSpanStartPayload]("span_end"),
		apijson.Discriminator[TelemetryLogEventParamsEventStructuredLogEventPayloadSpanEndPayload]("span_start"),
		apijson.Discriminator[TelemetryLogEventParamsEventStructuredLogEventPayloadSpanEndPayload]("span_end"),
	)
}

// The properties Name, Type are required.
type TelemetryLogEventParamsEventStructuredLogEventPayloadSpanStartPayload struct {
	Name string `json:"name,required"`
	// Any of "span_start", "span_end".
	Type         StructuredLogType `json:"type,omitzero,required"`
	ParentSpanID param.Opt[string] `json:"parent_span_id,omitzero"`
	paramObj
}

func (r TelemetryLogEventParamsEventStructuredLogEventPayloadSpanStartPayload) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryLogEventParamsEventStructuredLogEventPayloadSpanStartPayload
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryLogEventParamsEventStructuredLogEventPayloadSpanStartPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Status, Type are required.
type TelemetryLogEventParamsEventStructuredLogEventPayloadSpanEndPayload struct {
	// Any of "ok", "error".
	Status string `json:"status,omitzero,required"`
	// Any of "span_start", "span_end".
	Type StructuredLogType `json:"type,omitzero,required"`
	paramObj
}

func (r TelemetryLogEventParamsEventStructuredLogEventPayloadSpanEndPayload) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryLogEventParamsEventStructuredLogEventPayloadSpanEndPayload
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryLogEventParamsEventStructuredLogEventPayloadSpanEndPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TelemetryLogEventParamsEventStructuredLogEventPayloadSpanEndPayload](
		"status", "ok", "error",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TelemetryLogEventParamsEventStructuredLogEventAttributeUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u TelemetryLogEventParamsEventStructuredLogEventAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *TelemetryLogEventParamsEventStructuredLogEventAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TelemetryLogEventParamsEventStructuredLogEventAttributeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type TelemetryQueryMetricParams struct {
	// The type of query to perform.
	//
	// Any of "range", "instant".
	QueryType TelemetryQueryMetricParamsQueryType `json:"query_type,omitzero,required"`
	// The start time of the metric to query.
	StartTime int64 `json:"start_time,required"`
	// The end time of the metric to query.
	EndTime param.Opt[int64] `json:"end_time,omitzero"`
	// The granularity of the metric to query.
	Granularity param.Opt[string] `json:"granularity,omitzero"`
	// The label matchers to apply to the metric.
	LabelMatchers []TelemetryQueryMetricParamsLabelMatcher `json:"label_matchers,omitzero"`
	paramObj
}

func (r TelemetryQueryMetricParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryQueryMetricParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryQueryMetricParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of query to perform.
type TelemetryQueryMetricParamsQueryType string

const (
	TelemetryQueryMetricParamsQueryTypeRange   TelemetryQueryMetricParamsQueryType = "range"
	TelemetryQueryMetricParamsQueryTypeInstant TelemetryQueryMetricParamsQueryType = "instant"
)

// The properties Name, Operator, Value are required.
type TelemetryQueryMetricParamsLabelMatcher struct {
	Name string `json:"name,required"`
	// Any of "=", "!=", "=~", "!~".
	Operator string `json:"operator,omitzero,required"`
	Value    string `json:"value,required"`
	paramObj
}

func (r TelemetryQueryMetricParamsLabelMatcher) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryQueryMetricParamsLabelMatcher
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryQueryMetricParamsLabelMatcher) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TelemetryQueryMetricParamsLabelMatcher](
		"operator", "=", "!=", "=~", "!~",
	)
}

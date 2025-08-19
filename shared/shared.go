// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/varshaprasad96/llamastack-go-client"
	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// A image content item
type ImageContentItem struct {
	// Image as a base64 encoded string or an URL
	Image ImageContentItemImage `json:"image,required"`
	// Discriminator type of the content item. Always "image"
	//
	// Any of "text", "image", "tool_call".
	Type ImageContentItemType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageContentItem) RawJSON() string { return r.JSON.raw }
func (r *ImageContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ImageContentItem to a ImageContentItemParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ImageContentItemParam.Overrides()
func (r ImageContentItem) ToParam() ImageContentItemParam {
	return param.Override[ImageContentItemParam](json.RawMessage(r.RawJSON()))
}

// Image as a base64 encoded string or an URL
type ImageContentItemImage struct {
	// base64 encoded image data as string
	Data string `json:"data"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL llamastackclient.URL `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageContentItemImage) RawJSON() string { return r.JSON.raw }
func (r *ImageContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discriminator type of the content item. Always "image"
type ImageContentItemType string

const (
	ImageContentItemTypeText     ImageContentItemType = "text"
	ImageContentItemTypeImage    ImageContentItemType = "image"
	ImageContentItemTypeToolCall ImageContentItemType = "tool_call"
)

// A image content item
//
// The properties Image, Type are required.
type ImageContentItemParam struct {
	// Image as a base64 encoded string or an URL
	Image ImageContentItemImageParam `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// Any of "text", "image", "tool_call".
	Type ImageContentItemType `json:"type,omitzero,required"`
	paramObj
}

func (r ImageContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow ImageContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImageContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type ImageContentItemImageParam struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL llamastackclient.URLParam `json:"url,omitzero"`
	paramObj
}

func (r ImageContentItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow ImageContentItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImageContentItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A scoring result for a single row.
type ScoringResult struct {
	// Map of metric name to aggregated value
	AggregatedResults map[string]ScoringResultAggregatedResultUnion `json:"aggregated_results,required"`
	// The scoring result for each row. Each row is a map of column name to value.
	ScoreRows []map[string]ScoringResultScoreRowUnion `json:"score_rows,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregatedResults respjson.Field
		ScoreRows         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringResult) RawJSON() string { return r.JSON.raw }
func (r *ScoringResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoringResultAggregatedResultUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ScoringResultAggregatedResultUnion struct {
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

func (u ScoringResultAggregatedResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultAggregatedResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultAggregatedResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultAggregatedResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringResultAggregatedResultUnion) RawJSON() string { return u.JSON.raw }

func (r *ScoringResultAggregatedResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoringResultScoreRowUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ScoringResultScoreRowUnion struct {
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

func (u ScoringResultScoreRowUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultScoreRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultScoreRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultScoreRowUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringResultScoreRowUnion) RawJSON() string { return u.JSON.raw }

func (r *ScoringResultScoreRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type TextContentItem struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// Any of "text", "image", "tool_call".
	Type TextContentItemType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextContentItem) RawJSON() string { return r.JSON.raw }
func (r *TextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TextContentItem to a TextContentItemParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TextContentItemParam.Overrides()
func (r TextContentItem) ToParam() TextContentItemParam {
	return param.Override[TextContentItemParam](json.RawMessage(r.RawJSON()))
}

// Discriminator type of the content item. Always "text"
type TextContentItemType string

const (
	TextContentItemTypeText     TextContentItemType = "text"
	TextContentItemTypeImage    TextContentItemType = "image"
	TextContentItemTypeToolCall TextContentItemType = "tool_call"
)

// A text content item
//
// The properties Text, Type are required.
type TextContentItemParam struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// Any of "text", "image", "tool_call".
	Type TextContentItemType `json:"type,omitzero,required"`
	paramObj
}

func (r TextContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow TextContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

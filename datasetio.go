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
)

// DatasetioService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetioService] method instead.
type DatasetioService struct {
	Options []option.RequestOption
}

// NewDatasetioService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatasetioService(opts ...option.RequestOption) (r DatasetioService) {
	r = DatasetioService{}
	r.Options = opts
	return
}

// Append rows to a dataset.
func (r *DatasetioService) AppendRows(ctx context.Context, datasetID string, body DatasetioAppendRowsParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/datasetio/append-rows/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get a paginated list of rows from a dataset. Uses offset-based pagination where:
//
// - start_index: The starting index (0-based). If None, starts from beginning.
// - limit: Number of items to return. If None or -1, returns all items.
//
// The response includes:
//
// - data: List of items for the current page.
// - has_more: Whether there are more items available after this set.
func (r *DatasetioService) IterateRows(ctx context.Context, datasetID string, query DatasetioIterateRowsParams, opts ...option.RequestOption) (res *DatasetioIterateRowsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/datasetio/iterrows/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A generic paginated response that follows a simple format.
type DatasetioIterateRowsResponse struct {
	// The list of items for the current page
	Data []map[string]DatasetioIterateRowsResponseDataUnion `json:"data,required"`
	// Whether there are more items available after this set
	HasMore bool `json:"has_more,required"`
	// The URL for accessing this list
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatasetioIterateRowsResponse) RawJSON() string { return r.JSON.raw }
func (r *DatasetioIterateRowsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DatasetioIterateRowsResponseDataUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type DatasetioIterateRowsResponseDataUnion struct {
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

func (u DatasetioIterateRowsResponseDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetioIterateRowsResponseDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetioIterateRowsResponseDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetioIterateRowsResponseDataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DatasetioIterateRowsResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *DatasetioIterateRowsResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetioAppendRowsParams struct {
	// The rows to append to the dataset.
	Rows []map[string]DatasetioAppendRowsParamsRowUnion `json:"rows,omitzero,required"`
	paramObj
}

func (r DatasetioAppendRowsParams) MarshalJSON() (data []byte, err error) {
	type shadow DatasetioAppendRowsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetioAppendRowsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DatasetioAppendRowsParamsRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u DatasetioAppendRowsParamsRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *DatasetioAppendRowsParamsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DatasetioAppendRowsParamsRowUnion) asAny() any {
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

type DatasetioIterateRowsParams struct {
	// The number of rows to get.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Index into dataset for the first row to get. Get all rows if None.
	StartIndex param.Opt[int64] `query:"start_index,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DatasetioIterateRowsParams]'s query parameters as
// `url.Values`.
func (r DatasetioIterateRowsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

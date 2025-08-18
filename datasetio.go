// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/llamastack-go-client-go/internal/apijson"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/apiquery"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/param"
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
func (r *DatasetioService) IterateRows(ctx context.Context, datasetID string, query DatasetioIterateRowsParams, opts ...option.RequestOption) (res *PaginatedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/datasetio/iterrows/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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

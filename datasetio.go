// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/apiquery"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
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

// Iterate over rows in a dataset.
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

type DatasetioIterateRowsResponse = any

type DatasetioAppendRowsParams struct {
	Rows []map[string]any `json:"rows,omitzero,required"`
	paramObj
}

func (r DatasetioAppendRowsParams) MarshalJSON() (data []byte, err error) {
	type shadow DatasetioAppendRowsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetioAppendRowsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetioIterateRowsParams struct {
	// The number of rows to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index to start the iteration from.
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

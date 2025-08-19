// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
	"github.com/varshaprasad96/llamastack-go-client/shared"
)

// ScoringService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScoringService] method instead.
type ScoringService struct {
	Options []option.RequestOption
}

// NewScoringService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScoringService(opts ...option.RequestOption) (r ScoringService) {
	r = ScoringService{}
	r.Options = opts
	return
}

// Score a list of rows.
func (r *ScoringService) Score(ctx context.Context, body ScoringScoreParams, opts ...option.RequestOption) (res *ScoringScoreResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/scoring/score"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Score a batch of rows.
func (r *ScoringService) ScoreBatch(ctx context.Context, body ScoringScoreBatchParams, opts ...option.RequestOption) (res *ScoringScoreBatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/scoring/score-batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The response from scoring.
type ScoringScoreResponse struct {
	// A map of scoring function name to ScoringResult.
	Results map[string]shared.ScoringResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringScoreResponse) RawJSON() string { return r.JSON.raw }
func (r *ScoringScoreResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from batch scoring operations on datasets.
type ScoringScoreBatchResponse struct {
	// A map of scoring function name to ScoringResult
	Results map[string]shared.ScoringResult `json:"results,required"`
	// (Optional) The identifier of the dataset that was scored
	DatasetID string `json:"dataset_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		DatasetID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringScoreBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *ScoringScoreBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringScoreParams struct {
	// The rows to score.
	InputRows []map[string]ScoringScoreParamsInputRowUnion `json:"input_rows,omitzero,required"`
	// The scoring functions to use for the scoring.
	ScoringFunctions map[string]ScoringFnParamsUnion `json:"scoring_functions,omitzero,required"`
	paramObj
}

func (r ScoringScoreParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringScoreParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringScoreParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScoringScoreParamsInputRowUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ScoringScoreParamsInputRowUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ScoringScoreParamsInputRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScoringScoreParamsInputRowUnion) asAny() any {
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

type ScoringScoreBatchParams struct {
	// The ID of the dataset to score.
	DatasetID string `json:"dataset_id,required"`
	// Whether to save the results to a dataset.
	SaveResultsDataset bool `json:"save_results_dataset,required"`
	// The scoring functions to use for the scoring.
	ScoringFunctions map[string]ScoringFnParamsUnion `json:"scoring_functions,omitzero,required"`
	paramObj
}

func (r ScoringScoreBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringScoreBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringScoreBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

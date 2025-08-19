// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// VectorDBService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVectorDBService] method instead.
type VectorDBService struct {
	Options []option.RequestOption
}

// NewVectorDBService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVectorDBService(opts ...option.RequestOption) (r VectorDBService) {
	r = VectorDBService{}
	r.Options = opts
	return
}

// Register a vector database.
func (r *VectorDBService) New(ctx context.Context, body VectorDBNewParams, opts ...option.RequestOption) (res *VectorDB, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vector-dbs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a vector database by its identifier.
func (r *VectorDBService) Get(ctx context.Context, vectorDBID string, opts ...option.RequestOption) (res *VectorDB, err error) {
	opts = append(r.Options[:], opts...)
	if vectorDBID == "" {
		err = errors.New("missing required vector_db_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector-dbs/%s", vectorDBID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all vector databases.
func (r *VectorDBService) List(ctx context.Context, opts ...option.RequestOption) (res *VectorDBListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vector-dbs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Unregister a vector database.
func (r *VectorDBService) Delete(ctx context.Context, vectorDBID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if vectorDBID == "" {
		err = errors.New("missing required vector_db_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vector-dbs/%s", vectorDBID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type VectorDB struct {
	EmbeddingDimension int64  `json:"embedding_dimension,required"`
	EmbeddingModel     string `json:"embedding_model,required"`
	Identifier         string `json:"identifier,required"`
	ProviderID         string `json:"provider_id,required"`
	// Any of "model", "shield", "vector_db", "dataset", "scoring_function",
	// "benchmark", "tool", "tool_group".
	Type               VectorDBType `json:"type,required"`
	ProviderResourceID string       `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmbeddingDimension respjson.Field
		EmbeddingModel     respjson.Field
		Identifier         respjson.Field
		ProviderID         respjson.Field
		Type               respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorDB) RawJSON() string { return r.JSON.raw }
func (r *VectorDB) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorDBType string

const (
	VectorDBTypeModel           VectorDBType = "model"
	VectorDBTypeShield          VectorDBType = "shield"
	VectorDBTypeVectorDB        VectorDBType = "vector_db"
	VectorDBTypeDataset         VectorDBType = "dataset"
	VectorDBTypeScoringFunction VectorDBType = "scoring_function"
	VectorDBTypeBenchmark       VectorDBType = "benchmark"
	VectorDBTypeTool            VectorDBType = "tool"
	VectorDBTypeToolGroup       VectorDBType = "tool_group"
)

type VectorDBListResponse struct {
	Data []VectorDB `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorDBListResponse) RawJSON() string { return r.JSON.raw }
func (r *VectorDBListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorDBNewParams struct {
	// The embedding model to use.
	EmbeddingModel string `json:"embedding_model,required"`
	// The identifier of the vector database to register.
	VectorDBID string `json:"vector_db_id,required"`
	// The dimension of the embedding model.
	EmbeddingDimension param.Opt[int64] `json:"embedding_dimension,omitzero"`
	// The identifier of the provider.
	ProviderID param.Opt[string] `json:"provider_id,omitzero"`
	// The identifier of the vector database in the provider.
	ProviderVectorDBID param.Opt[string] `json:"provider_vector_db_id,omitzero"`
	paramObj
}

func (r VectorDBNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorDBNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorDBNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

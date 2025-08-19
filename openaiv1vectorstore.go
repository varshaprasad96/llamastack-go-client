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

// OpenAIV1VectorStoreService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIV1VectorStoreService] method instead.
type OpenAIV1VectorStoreService struct {
	Options []option.RequestOption
	Files   OpenAIV1VectorStoreFileService
}

// NewOpenAIV1VectorStoreService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOpenAIV1VectorStoreService(opts ...option.RequestOption) (r OpenAIV1VectorStoreService) {
	r = OpenAIV1VectorStoreService{}
	r.Options = opts
	r.Files = NewOpenAIV1VectorStoreFileService(opts...)
	return
}

// Creates a vector store.
func (r *OpenAIV1VectorStoreService) New(ctx context.Context, body OpenAIV1VectorStoreNewParams, opts ...option.RequestOption) (res *VectorStoreObject, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/vector_stores"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a vector store.
func (r *OpenAIV1VectorStoreService) Get(ctx context.Context, vectorStoreID string, opts ...option.RequestOption) (res *VectorStoreObject, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a vector store.
func (r *OpenAIV1VectorStoreService) Update(ctx context.Context, vectorStoreID string, body OpenAIV1VectorStoreUpdateParams, opts ...option.RequestOption) (res *VectorStoreObject, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of vector stores.
func (r *OpenAIV1VectorStoreService) List(ctx context.Context, query OpenAIV1VectorStoreListParams, opts ...option.RequestOption) (res *OpenAiv1VectorStoreListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/vector_stores"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a vector store.
func (r *OpenAIV1VectorStoreService) Delete(ctx context.Context, vectorStoreID string, opts ...option.RequestOption) (res *OpenAiv1VectorStoreDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Search for chunks in a vector store. Searches a vector store for relevant chunks
// based on a query and optional file attribute filters.
func (r *OpenAIV1VectorStoreService) Search(ctx context.Context, vectorStoreID string, body OpenAIV1VectorStoreSearchParams, opts ...option.RequestOption) (res *OpenAiv1VectorStoreSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/search", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// OpenAI Vector Store object.
type VectorStoreObject struct {
	ID           string                                        `json:"id,required"`
	CreatedAt    int64                                         `json:"created_at,required"`
	FileCounts   VectorStoreObjectFileCounts                   `json:"file_counts,required"`
	Metadata     map[string]VectorStoreObjectMetadataUnion     `json:"metadata,required"`
	Object       string                                        `json:"object,required"`
	Status       string                                        `json:"status,required"`
	UsageBytes   int64                                         `json:"usage_bytes,required"`
	ExpiresAfter map[string]VectorStoreObjectExpiresAfterUnion `json:"expires_after"`
	ExpiresAt    int64                                         `json:"expires_at"`
	LastActiveAt int64                                         `json:"last_active_at"`
	Name         string                                        `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		FileCounts   respjson.Field
		Metadata     respjson.Field
		Object       respjson.Field
		Status       respjson.Field
		UsageBytes   respjson.Field
		ExpiresAfter respjson.Field
		ExpiresAt    respjson.Field
		LastActiveAt respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreObject) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreObjectFileCounts struct {
	Cancelled  int64 `json:"cancelled,required"`
	Completed  int64 `json:"completed,required"`
	Failed     int64 `json:"failed,required"`
	InProgress int64 `json:"in_progress,required"`
	Total      int64 `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled   respjson.Field
		Completed   respjson.Field
		Failed      respjson.Field
		InProgress  respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreObjectFileCounts) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreObjectFileCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VectorStoreObjectMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type VectorStoreObjectMetadataUnion struct {
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

func (u VectorStoreObjectMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreObjectMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreObjectMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreObjectMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorStoreObjectMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreObjectMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VectorStoreObjectExpiresAfterUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type VectorStoreObjectExpiresAfterUnion struct {
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

func (u VectorStoreObjectExpiresAfterUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreObjectExpiresAfterUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreObjectExpiresAfterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreObjectExpiresAfterUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorStoreObjectExpiresAfterUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreObjectExpiresAfterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from listing vector stores.
type OpenAiv1VectorStoreListResponse struct {
	Data    []VectorStoreObject `json:"data,required"`
	HasMore bool                `json:"has_more,required"`
	Object  string              `json:"object,required"`
	FirstID string              `json:"first_id"`
	LastID  string              `json:"last_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Object      respjson.Field
		FirstID     respjson.Field
		LastID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1VectorStoreListResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1VectorStoreListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from deleting a vector store.
type OpenAiv1VectorStoreDeleteResponse struct {
	ID      string `json:"id,required"`
	Deleted bool   `json:"deleted,required"`
	Object  string `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1VectorStoreDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1VectorStoreDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from searching a vector store.
type OpenAiv1VectorStoreSearchResponse struct {
	Data        []OpenAiv1VectorStoreSearchResponseData `json:"data,required"`
	HasMore     bool                                    `json:"has_more,required"`
	Object      string                                  `json:"object,required"`
	SearchQuery string                                  `json:"search_query,required"`
	NextPage    string                                  `json:"next_page"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Object      respjson.Field
		SearchQuery respjson.Field
		NextPage    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1VectorStoreSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1VectorStoreSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from searching a vector store.
type OpenAiv1VectorStoreSearchResponseData struct {
	Content    []Content                                                      `json:"content,required"`
	FileID     string                                                         `json:"file_id,required"`
	Filename   string                                                         `json:"filename,required"`
	Score      float64                                                        `json:"score,required"`
	Attributes map[string]OpenAiv1VectorStoreSearchResponseDataAttributeUnion `json:"attributes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		Score       respjson.Field
		Attributes  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1VectorStoreSearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1VectorStoreSearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAiv1VectorStoreSearchResponseDataAttributeUnion contains all possible
// properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type OpenAiv1VectorStoreSearchResponseDataAttributeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u OpenAiv1VectorStoreSearchResponseDataAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1VectorStoreSearchResponseDataAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1VectorStoreSearchResponseDataAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAiv1VectorStoreSearchResponseDataAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *OpenAiv1VectorStoreSearchResponseDataAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIV1VectorStoreNewParams struct {
	// A name for the vector store.
	Name string `json:"name,required"`
	// The dimension of the embedding vectors (default: 384).
	EmbeddingDimension param.Opt[int64] `json:"embedding_dimension,omitzero"`
	// The embedding model to use for this vector store.
	EmbeddingModel param.Opt[string] `json:"embedding_model,omitzero"`
	// The ID of the provider to use for this vector store.
	ProviderID param.Opt[string] `json:"provider_id,omitzero"`
	// The provider-specific vector database ID.
	ProviderVectorDBID param.Opt[string] `json:"provider_vector_db_id,omitzero"`
	// The chunking strategy used to chunk the file(s). If not set, will use the `auto`
	// strategy.
	ChunkingStrategy map[string]OpenAIV1VectorStoreNewParamsChunkingStrategyUnion `json:"chunking_strategy,omitzero"`
	// The expiration policy for a vector store.
	ExpiresAfter map[string]OpenAIV1VectorStoreNewParamsExpiresAfterUnion `json:"expires_after,omitzero"`
	// A list of File IDs that the vector store should use. Useful for tools like
	// `file_search` that can access files.
	FileIDs []string `json:"file_ids,omitzero"`
	// Set of 16 key-value pairs that can be attached to an object.
	Metadata map[string]OpenAIV1VectorStoreNewParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r OpenAIV1VectorStoreNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1VectorStoreNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1VectorStoreNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1VectorStoreNewParamsChunkingStrategyUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1VectorStoreNewParamsChunkingStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1VectorStoreNewParamsChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1VectorStoreNewParamsChunkingStrategyUnion) asAny() any {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1VectorStoreNewParamsExpiresAfterUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1VectorStoreNewParamsExpiresAfterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1VectorStoreNewParamsExpiresAfterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1VectorStoreNewParamsExpiresAfterUnion) asAny() any {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1VectorStoreNewParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1VectorStoreNewParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1VectorStoreNewParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1VectorStoreNewParamsMetadataUnion) asAny() any {
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

type OpenAIV1VectorStoreUpdateParams struct {
	// The name of the vector store.
	Name param.Opt[string] `json:"name,omitzero"`
	// The expiration policy for a vector store.
	ExpiresAfter map[string]OpenAIV1VectorStoreUpdateParamsExpiresAfterUnion `json:"expires_after,omitzero"`
	// Set of 16 key-value pairs that can be attached to an object.
	Metadata map[string]OpenAIV1VectorStoreUpdateParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r OpenAIV1VectorStoreUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1VectorStoreUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1VectorStoreUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1VectorStoreUpdateParamsExpiresAfterUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1VectorStoreUpdateParamsExpiresAfterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1VectorStoreUpdateParamsExpiresAfterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1VectorStoreUpdateParamsExpiresAfterUnion) asAny() any {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1VectorStoreUpdateParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1VectorStoreUpdateParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1VectorStoreUpdateParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1VectorStoreUpdateParamsMetadataUnion) asAny() any {
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

type OpenAIV1VectorStoreListParams struct {
	// A cursor for use in pagination. `after` is an object ID that defines your place
	// in the list.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// A cursor for use in pagination. `before` is an object ID that defines your place
	// in the list.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// A limit on the number of objects to be returned. Limit can range between 1 and
	// 100, and the default is 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order by the `created_at` timestamp of the objects. `asc` for ascending
	// order and `desc` for descending order.
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OpenAIV1VectorStoreListParams]'s query parameters as
// `url.Values`.
func (r OpenAIV1VectorStoreListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OpenAIV1VectorStoreSearchParams struct {
	// The query string or array for performing the search.
	Query OpenAIV1VectorStoreSearchParamsQueryUnion `json:"query,omitzero,required"`
	// Maximum number of results to return (1 to 50 inclusive, default 10).
	MaxNumResults param.Opt[int64] `json:"max_num_results,omitzero"`
	// Whether to rewrite the natural language query for vector search (default false)
	RewriteQuery param.Opt[bool] `json:"rewrite_query,omitzero"`
	// The search mode to use - "keyword", "vector", or "hybrid" (default "vector")
	SearchMode param.Opt[string] `json:"search_mode,omitzero"`
	// Filters based on file attributes to narrow the search results.
	Filters map[string]OpenAIV1VectorStoreSearchParamsFilterUnion `json:"filters,omitzero"`
	// Ranking options for fine-tuning the search results.
	RankingOptions OpenAIV1VectorStoreSearchParamsRankingOptions `json:"ranking_options,omitzero"`
	paramObj
}

func (r OpenAIV1VectorStoreSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1VectorStoreSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1VectorStoreSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1VectorStoreSearchParamsQueryUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1VectorStoreSearchParamsQueryUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *OpenAIV1VectorStoreSearchParamsQueryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1VectorStoreSearchParamsQueryUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1VectorStoreSearchParamsFilterUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1VectorStoreSearchParamsFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1VectorStoreSearchParamsFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1VectorStoreSearchParamsFilterUnion) asAny() any {
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

// Ranking options for fine-tuning the search results.
type OpenAIV1VectorStoreSearchParamsRankingOptions struct {
	Ranker         param.Opt[string]  `json:"ranker,omitzero"`
	ScoreThreshold param.Opt[float64] `json:"score_threshold,omitzero"`
	paramObj
}

func (r OpenAIV1VectorStoreSearchParamsRankingOptions) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1VectorStoreSearchParamsRankingOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1VectorStoreSearchParamsRankingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

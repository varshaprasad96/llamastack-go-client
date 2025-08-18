// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stainless-sdks/llamastack-go-client-go/internal/apijson"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/param"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/respjson"
)

// VectorIoService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVectorIoService] method instead.
type VectorIoService struct {
	Options []option.RequestOption
}

// NewVectorIoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVectorIoService(opts ...option.RequestOption) (r VectorIoService) {
	r = VectorIoService{}
	r.Options = opts
	return
}

// Insert chunks into a vector database.
func (r *VectorIoService) Insert(ctx context.Context, body VectorIoInsertParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/vector-io/insert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Query chunks from a vector database.
func (r *VectorIoService) Query(ctx context.Context, body VectorIoQueryParams, opts ...option.RequestOption) (res *VectorIoQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vector-io/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A chunk of content that can be inserted into a vector database.
type Chunk struct {
	// The content of the chunk, which can be interleaved text, images, or other types.
	Content InterleavedContentUnion `json:"content,required"`
	// Metadata associated with the chunk that will be used in the model context during
	// inference.
	Metadata map[string]ChunkMetadataUnion `json:"metadata,required"`
	// Metadata for the chunk that will NOT be used in the context during inference.
	// The `chunk_metadata` is required backend functionality.
	ChunkMetadata ChunkChunkMetadata `json:"chunk_metadata"`
	// Optional embedding for the chunk. If not provided, it will be computed later.
	Embedding []float64 `json:"embedding"`
	// The chunk ID that is stored in the vector database. Used for backend
	// functionality.
	StoredChunkID string `json:"stored_chunk_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content       respjson.Field
		Metadata      respjson.Field
		ChunkMetadata respjson.Field
		Embedding     respjson.Field
		StoredChunkID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Chunk) RawJSON() string { return r.JSON.raw }
func (r *Chunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Chunk to a ChunkParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChunkParam.Overrides()
func (r Chunk) ToParam() ChunkParam {
	return param.Override[ChunkParam](json.RawMessage(r.RawJSON()))
}

// ChunkMetadataUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ChunkMetadataUnion struct {
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

func (u ChunkMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChunkMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ChunkMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for the chunk that will NOT be used in the context during inference.
// The `chunk_metadata` is required backend functionality.
type ChunkChunkMetadata struct {
	// The dimension of the embedding vector for the chunk.
	ChunkEmbeddingDimension int64 `json:"chunk_embedding_dimension"`
	// The embedding model used to create the chunk's embedding.
	ChunkEmbeddingModel string `json:"chunk_embedding_model"`
	// The ID of the chunk. If not set, it will be generated based on the document ID
	// and content.
	ChunkID string `json:"chunk_id"`
	// The tokenizer used to create the chunk. Default is Tiktoken.
	ChunkTokenizer string `json:"chunk_tokenizer"`
	// The window of the chunk, which can be used to group related chunks together.
	ChunkWindow string `json:"chunk_window"`
	// The number of tokens in the content of the chunk.
	ContentTokenCount int64 `json:"content_token_count"`
	// An optional timestamp indicating when the chunk was created.
	CreatedTimestamp int64 `json:"created_timestamp"`
	// The ID of the document this chunk belongs to.
	DocumentID string `json:"document_id"`
	// The number of tokens in the metadata of the chunk.
	MetadataTokenCount int64 `json:"metadata_token_count"`
	// The source of the content, such as a URL, file path, or other identifier.
	Source string `json:"source"`
	// An optional timestamp indicating when the chunk was last updated.
	UpdatedTimestamp int64 `json:"updated_timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkEmbeddingDimension respjson.Field
		ChunkEmbeddingModel     respjson.Field
		ChunkID                 respjson.Field
		ChunkTokenizer          respjson.Field
		ChunkWindow             respjson.Field
		ContentTokenCount       respjson.Field
		CreatedTimestamp        respjson.Field
		DocumentID              respjson.Field
		MetadataTokenCount      respjson.Field
		Source                  respjson.Field
		UpdatedTimestamp        respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChunkChunkMetadata) RawJSON() string { return r.JSON.raw }
func (r *ChunkChunkMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chunk of content that can be inserted into a vector database.
//
// The properties Content, Metadata are required.
type ChunkParam struct {
	// The content of the chunk, which can be interleaved text, images, or other types.
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Metadata associated with the chunk that will be used in the model context during
	// inference.
	Metadata map[string]ChunkMetadataUnionParam `json:"metadata,omitzero,required"`
	// The chunk ID that is stored in the vector database. Used for backend
	// functionality.
	StoredChunkID param.Opt[string] `json:"stored_chunk_id,omitzero"`
	// Metadata for the chunk that will NOT be used in the context during inference.
	// The `chunk_metadata` is required backend functionality.
	ChunkMetadata ChunkChunkMetadataParam `json:"chunk_metadata,omitzero"`
	// Optional embedding for the chunk. If not provided, it will be computed later.
	Embedding []float64 `json:"embedding,omitzero"`
	paramObj
}

func (r ChunkParam) MarshalJSON() (data []byte, err error) {
	type shadow ChunkParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChunkParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChunkMetadataUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ChunkMetadataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ChunkMetadataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChunkMetadataUnionParam) asAny() any {
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

// Metadata for the chunk that will NOT be used in the context during inference.
// The `chunk_metadata` is required backend functionality.
type ChunkChunkMetadataParam struct {
	// The dimension of the embedding vector for the chunk.
	ChunkEmbeddingDimension param.Opt[int64] `json:"chunk_embedding_dimension,omitzero"`
	// The embedding model used to create the chunk's embedding.
	ChunkEmbeddingModel param.Opt[string] `json:"chunk_embedding_model,omitzero"`
	// The ID of the chunk. If not set, it will be generated based on the document ID
	// and content.
	ChunkID param.Opt[string] `json:"chunk_id,omitzero"`
	// The tokenizer used to create the chunk. Default is Tiktoken.
	ChunkTokenizer param.Opt[string] `json:"chunk_tokenizer,omitzero"`
	// The window of the chunk, which can be used to group related chunks together.
	ChunkWindow param.Opt[string] `json:"chunk_window,omitzero"`
	// The number of tokens in the content of the chunk.
	ContentTokenCount param.Opt[int64] `json:"content_token_count,omitzero"`
	// An optional timestamp indicating when the chunk was created.
	CreatedTimestamp param.Opt[int64] `json:"created_timestamp,omitzero"`
	// The ID of the document this chunk belongs to.
	DocumentID param.Opt[string] `json:"document_id,omitzero"`
	// The number of tokens in the metadata of the chunk.
	MetadataTokenCount param.Opt[int64] `json:"metadata_token_count,omitzero"`
	// The source of the content, such as a URL, file path, or other identifier.
	Source param.Opt[string] `json:"source,omitzero"`
	// An optional timestamp indicating when the chunk was last updated.
	UpdatedTimestamp param.Opt[int64] `json:"updated_timestamp,omitzero"`
	paramObj
}

func (r ChunkChunkMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ChunkChunkMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChunkChunkMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from querying chunks in a vector database.
type VectorIoQueryResponse struct {
	// List of content chunks returned from the query
	Chunks []Chunk `json:"chunks,required"`
	// Relevance scores corresponding to each returned chunk
	Scores []float64 `json:"scores,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chunks      respjson.Field
		Scores      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorIoQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *VectorIoQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorIoInsertParams struct {
	// The chunks to insert. Each `Chunk` should contain content which can be
	// interleaved text, images, or other types. `metadata`: `dict[str, Any]` and
	// `embedding`: `List[float]` are optional. If `metadata` is provided, you
	// configure how Llama Stack formats the chunk during generation. If `embedding` is
	// not provided, it will be computed later.
	Chunks []ChunkParam `json:"chunks,omitzero,required"`
	// The identifier of the vector database to insert the chunks into.
	VectorDBID string `json:"vector_db_id,required"`
	// The time to live of the chunks.
	TtlSeconds param.Opt[int64] `json:"ttl_seconds,omitzero"`
	paramObj
}

func (r VectorIoInsertParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoInsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoInsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorIoQueryParams struct {
	// The query to search for.
	Query InterleavedContentUnionParam `json:"query,omitzero,required"`
	// The identifier of the vector database to query.
	VectorDBID string `json:"vector_db_id,required"`
	// The parameters of the query.
	Params map[string]VectorIoQueryParamsParamUnion `json:"params,omitzero"`
	paramObj
}

func (r VectorIoQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorIoQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorIoQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VectorIoQueryParamsParamUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u VectorIoQueryParamsParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *VectorIoQueryParamsParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VectorIoQueryParamsParamUnion) asAny() any {
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

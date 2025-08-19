// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
	"github.com/varshaprasad96/llamastack-go-client/shared"
	"github.com/varshaprasad96/llamastack-go-client/shared/constant"
)

// ToolRuntimeRagToolService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolRuntimeRagToolService] method instead.
type ToolRuntimeRagToolService struct {
	Options []option.RequestOption
}

// NewToolRuntimeRagToolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewToolRuntimeRagToolService(opts ...option.RequestOption) (r ToolRuntimeRagToolService) {
	r = ToolRuntimeRagToolService{}
	r.Options = opts
	return
}

// Index documents so they can be used by the RAG system.
func (r *ToolRuntimeRagToolService) Insert(ctx context.Context, body ToolRuntimeRagToolInsertParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/tool-runtime/rag-tool/insert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Query the RAG system for context; typically invoked by the agent.
func (r *ToolRuntimeRagToolService) Query(ctx context.Context, body ToolRuntimeRagToolQueryParams, opts ...option.RequestOption) (res *ToolRuntimeRagToolQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tool-runtime/rag-tool/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Result of a RAG query containing retrieved content and metadata.
type ToolRuntimeRagToolQueryResponse struct {
	// Additional metadata about the query result
	Metadata map[string]ToolRuntimeRagToolQueryResponseMetadataUnion `json:"metadata,required"`
	// (Optional) The retrieved content from the query
	Content InterleavedContentUnion `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolRuntimeRagToolQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *ToolRuntimeRagToolQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolRuntimeRagToolQueryResponseMetadataUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolRuntimeRagToolQueryResponseMetadataUnion struct {
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

func (u ToolRuntimeRagToolQueryResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolRuntimeRagToolQueryResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolRuntimeRagToolQueryResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolRuntimeRagToolQueryResponseMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolRuntimeRagToolQueryResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolRuntimeRagToolQueryResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolRuntimeRagToolInsertParams struct {
	// (Optional) Size in tokens for document chunking during indexing
	ChunkSizeInTokens int64 `json:"chunk_size_in_tokens,required"`
	// List of documents to index in the RAG system
	Documents []ToolRuntimeRagToolInsertParamsDocument `json:"documents,omitzero,required"`
	// ID of the vector database to store the document embeddings
	VectorDBID string `json:"vector_db_id,required"`
	paramObj
}

func (r ToolRuntimeRagToolInsertParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeRagToolInsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeRagToolInsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A document to be used for document ingestion in the RAG Tool.
//
// The properties Content, DocumentID, Metadata are required.
type ToolRuntimeRagToolInsertParamsDocument struct {
	// The content of the document.
	Content ToolRuntimeRagToolInsertParamsDocumentContentUnion `json:"content,omitzero,required"`
	// The unique identifier for the document.
	DocumentID string `json:"document_id,required"`
	// Additional metadata for the document.
	Metadata map[string]ToolRuntimeRagToolInsertParamsDocumentMetadataUnion `json:"metadata,omitzero,required"`
	// The MIME type of the document.
	MimeType param.Opt[string] `json:"mime_type,omitzero"`
	paramObj
}

func (r ToolRuntimeRagToolInsertParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeRagToolInsertParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeRagToolInsertParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolRuntimeRagToolInsertParamsDocumentContentUnion struct {
	OfString                      param.Opt[string]                  `json:",omitzero,inline"`
	OfImageContentItem            *shared.ImageContentItemParam      `json:",omitzero,inline"`
	OfTextContentItem             *shared.TextContentItemParam       `json:",omitzero,inline"`
	OfInterleavedContentItemArray []InterleavedContentItemUnionParam `json:",omitzero,inline"`
	OfURL                         *URLParam                          `json:",omitzero,inline"`
	paramUnion
}

func (u ToolRuntimeRagToolInsertParamsDocumentContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfImageContentItem,
		u.OfTextContentItem,
		u.OfInterleavedContentItemArray,
		u.OfURL)
}
func (u *ToolRuntimeRagToolInsertParamsDocumentContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolRuntimeRagToolInsertParamsDocumentContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfImageContentItem) {
		return u.OfImageContentItem
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	} else if !param.IsOmitted(u.OfInterleavedContentItemArray) {
		return &u.OfInterleavedContentItemArray
	} else if !param.IsOmitted(u.OfURL) {
		return u.OfURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolInsertParamsDocumentContentUnion) GetImage() *shared.ImageContentItemImageParam {
	if vt := u.OfImageContentItem; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolInsertParamsDocumentContentUnion) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolInsertParamsDocumentContentUnion) GetUri() *string {
	if vt := u.OfURL; vt != nil {
		return &vt.Uri
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolInsertParamsDocumentContentUnion) GetType() *string {
	if vt := u.OfImageContentItem; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolRuntimeRagToolInsertParamsDocumentMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ToolRuntimeRagToolInsertParamsDocumentMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ToolRuntimeRagToolInsertParamsDocumentMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolRuntimeRagToolInsertParamsDocumentMetadataUnion) asAny() any {
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

type ToolRuntimeRagToolQueryParams struct {
	// The query content to search for in the indexed documents
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// List of vector database IDs to search within
	VectorDBIDs []string `json:"vector_db_ids,omitzero,required"`
	// (Optional) Configuration parameters for the query operation
	QueryConfig ToolRuntimeRagToolQueryParamsQueryConfig `json:"query_config,omitzero"`
	paramObj
}

func (r ToolRuntimeRagToolQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeRagToolQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeRagToolQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Configuration parameters for the query operation
//
// The properties ChunkTemplate, MaxChunks, MaxTokensInContext,
// QueryGeneratorConfig are required.
type ToolRuntimeRagToolQueryParamsQueryConfig struct {
	// Template for formatting each retrieved chunk in the context. Available
	// placeholders: {index} (1-based chunk ordinal), {chunk.content} (chunk content
	// string), {metadata} (chunk metadata dict). Default: "Result {index}\nContent:
	// {chunk.content}\nMetadata: {metadata}\n"
	ChunkTemplate string `json:"chunk_template,required"`
	// Maximum number of chunks to retrieve.
	MaxChunks int64 `json:"max_chunks,required"`
	// Maximum number of tokens in the context.
	MaxTokensInContext int64 `json:"max_tokens_in_context,required"`
	// Configuration for the query generator.
	QueryGeneratorConfig ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion `json:"query_generator_config,omitzero,required"`
	// Search mode for retrieval—either "vector", "keyword", or "hybrid". Default
	// "vector".
	//
	// Any of "vector", "keyword", "hybrid".
	Mode string `json:"mode,omitzero"`
	// Configuration for the ranker to use in hybrid search. Defaults to RRF ranker.
	Ranker ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion `json:"ranker,omitzero"`
	paramObj
}

func (r ToolRuntimeRagToolQueryParamsQueryConfig) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeRagToolQueryParamsQueryConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeRagToolQueryParamsQueryConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ToolRuntimeRagToolQueryParamsQueryConfig](
		"mode", "vector", "keyword", "hybrid",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion struct {
	OfDefault *ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigDefault `json:",omitzero,inline"`
	OfLlm     *ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigLlm     `json:",omitzero,inline"`
	paramUnion
}

func (u ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDefault, u.OfLlm)
}
func (u *ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfDefault) {
		return u.OfDefault
	} else if !param.IsOmitted(u.OfLlm) {
		return u.OfLlm
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion) GetSeparator() *string {
	if vt := u.OfDefault; vt != nil {
		return &vt.Separator
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion) GetModel() *string {
	if vt := u.OfLlm; vt != nil {
		return &vt.Model
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion) GetTemplate() *string {
	if vt := u.OfLlm; vt != nil {
		return &vt.Template
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion) GetType() *string {
	if vt := u.OfDefault; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLlm; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigUnion](
		"type",
		apijson.Discriminator[ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigDefault]("default"),
		apijson.Discriminator[ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigLlm]("llm"),
	)
}

// Configuration for the default RAG query generator.
//
// The properties Separator, Type are required.
type ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigDefault struct {
	// String separator used to join query terms
	Separator string `json:"separator,required"`
	// Type of query generator, always 'default'
	//
	// This field can be elided, and will marshal its zero value as "default".
	Type constant.Default `json:"type,required"`
	paramObj
}

func (r ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigDefault) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigDefault
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the LLM-based RAG query generator.
//
// The properties Model, Template, Type are required.
type ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigLlm struct {
	// Name of the language model to use for query generation
	Model string `json:"model,required"`
	// Template string for formatting the query generation prompt
	Template string `json:"template,required"`
	// Type of query generator, always 'llm'
	//
	// This field can be elided, and will marshal its zero value as "llm".
	Type constant.Llm `json:"type,required"`
	paramObj
}

func (r ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigLlm) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigLlm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeRagToolQueryParamsQueryConfigQueryGeneratorConfigLlm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion struct {
	OfRrf      *ToolRuntimeRagToolQueryParamsQueryConfigRankerRrf      `json:",omitzero,inline"`
	OfWeighted *ToolRuntimeRagToolQueryParamsQueryConfigRankerWeighted `json:",omitzero,inline"`
	paramUnion
}

func (u ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRrf, u.OfWeighted)
}
func (u *ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion) asAny() any {
	if !param.IsOmitted(u.OfRrf) {
		return u.OfRrf
	} else if !param.IsOmitted(u.OfWeighted) {
		return u.OfWeighted
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion) GetImpactFactor() *float64 {
	if vt := u.OfRrf; vt != nil {
		return &vt.ImpactFactor
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion) GetAlpha() *float64 {
	if vt := u.OfWeighted; vt != nil {
		return &vt.Alpha
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion) GetType() *string {
	if vt := u.OfRrf; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWeighted; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ToolRuntimeRagToolQueryParamsQueryConfigRankerUnion](
		"type",
		apijson.Discriminator[ToolRuntimeRagToolQueryParamsQueryConfigRankerRrf]("rrf"),
		apijson.Discriminator[ToolRuntimeRagToolQueryParamsQueryConfigRankerWeighted]("weighted"),
	)
}

// Reciprocal Rank Fusion (RRF) ranker configuration.
//
// The properties ImpactFactor, Type are required.
type ToolRuntimeRagToolQueryParamsQueryConfigRankerRrf struct {
	// The impact factor for RRF scoring. Higher values give more weight to
	// higher-ranked results. Must be greater than 0
	ImpactFactor float64 `json:"impact_factor,required"`
	// The type of ranker, always "rrf"
	//
	// This field can be elided, and will marshal its zero value as "rrf".
	Type constant.Rrf `json:"type,required"`
	paramObj
}

func (r ToolRuntimeRagToolQueryParamsQueryConfigRankerRrf) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeRagToolQueryParamsQueryConfigRankerRrf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeRagToolQueryParamsQueryConfigRankerRrf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weighted ranker configuration that combines vector and keyword scores.
//
// The properties Alpha, Type are required.
type ToolRuntimeRagToolQueryParamsQueryConfigRankerWeighted struct {
	// Weight factor between 0 and 1. 0 means only use keyword scores, 1 means only use
	// vector scores, values in between blend both scores.
	Alpha float64 `json:"alpha,required"`
	// The type of ranker, always "weighted"
	//
	// This field can be elided, and will marshal its zero value as "weighted".
	Type constant.Weighted `json:"type,required"`
	paramObj
}

func (r ToolRuntimeRagToolQueryParamsQueryConfigRankerWeighted) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeRagToolQueryParamsQueryConfigRankerWeighted
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeRagToolQueryParamsQueryConfigRankerWeighted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

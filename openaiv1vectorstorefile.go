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
	"github.com/varshaprasad96/llamastack-go-client/shared/constant"
)

// OpenAIV1VectorStoreFileService contains methods and other services that help
// with interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIV1VectorStoreFileService] method instead.
type OpenAIV1VectorStoreFileService struct {
	Options []option.RequestOption
}

// NewOpenAIV1VectorStoreFileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOpenAIV1VectorStoreFileService(opts ...option.RequestOption) (r OpenAIV1VectorStoreFileService) {
	r = OpenAIV1VectorStoreFileService{}
	r.Options = opts
	return
}

// Attach a file to a vector store.
func (r *OpenAIV1VectorStoreFileService) New(ctx context.Context, vectorStoreID string, body OpenAIV1VectorStoreFileNewParams, opts ...option.RequestOption) (res *VectorStoreFile, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a vector store file.
func (r *OpenAIV1VectorStoreFileService) Get(ctx context.Context, fileID string, query OpenAIV1VectorStoreFileGetParams, opts ...option.RequestOption) (res *VectorStoreFile, err error) {
	opts = append(r.Options[:], opts...)
	if query.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files/%s", query.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a vector store file.
func (r *OpenAIV1VectorStoreFileService) Update(ctx context.Context, fileID string, params OpenAIV1VectorStoreFileUpdateParams, opts ...option.RequestOption) (res *VectorStoreFile, err error) {
	opts = append(r.Options[:], opts...)
	if params.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files/%s", params.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List files in a vector store.
func (r *OpenAIV1VectorStoreFileService) List(ctx context.Context, vectorStoreID string, query OpenAIV1VectorStoreFileListParams, opts ...option.RequestOption) (res *OpenAiv1VectorStoreFileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files", vectorStoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a vector store file.
func (r *OpenAIV1VectorStoreFileService) Delete(ctx context.Context, fileID string, body OpenAIV1VectorStoreFileDeleteParams, opts ...option.RequestOption) (res *OpenAiv1VectorStoreFileDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files/%s", body.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieves the contents of a vector store file.
func (r *OpenAIV1VectorStoreFileService) GetContent(ctx context.Context, fileID string, query OpenAIV1VectorStoreFileGetContentParams, opts ...option.RequestOption) (res *OpenAiv1VectorStoreFileGetContentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.VectorStoreID == "" {
		err = errors.New("missing required vector_store_id parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/vector_stores/%s/files/%s/content", query.VectorStoreID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// ChunkingStrategyUnion contains all possible properties and values from
// [ChunkingStrategyAuto], [ChunkingStrategyStatic].
//
// Use the [ChunkingStrategyUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChunkingStrategyUnion struct {
	// Any of "auto", "static".
	Type string `json:"type"`
	// This field is from variant [ChunkingStrategyStatic].
	Static ChunkingStrategyStaticStatic `json:"static"`
	JSON   struct {
		Type   respjson.Field
		Static respjson.Field
		raw    string
	} `json:"-"`
}

// anyChunkingStrategy is implemented by each variant of [ChunkingStrategyUnion] to
// add type safety for the return type of [ChunkingStrategyUnion.AsAny]
type anyChunkingStrategy interface {
	implChunkingStrategyUnion()
}

func (ChunkingStrategyAuto) implChunkingStrategyUnion()   {}
func (ChunkingStrategyStatic) implChunkingStrategyUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChunkingStrategyUnion.AsAny().(type) {
//	case llamastackclient.ChunkingStrategyAuto:
//	case llamastackclient.ChunkingStrategyStatic:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChunkingStrategyUnion) AsAny() anyChunkingStrategy {
	switch u.Type {
	case "auto":
		return u.AsAuto()
	case "static":
		return u.AsStatic()
	}
	return nil
}

func (u ChunkingStrategyUnion) AsAuto() (v ChunkingStrategyAuto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkingStrategyUnion) AsStatic() (v ChunkingStrategyStatic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChunkingStrategyUnion) RawJSON() string { return u.JSON.raw }

func (r *ChunkingStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChunkingStrategyUnion to a ChunkingStrategyUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChunkingStrategyUnionParam.Overrides()
func (r ChunkingStrategyUnion) ToParam() ChunkingStrategyUnionParam {
	return param.Override[ChunkingStrategyUnionParam](json.RawMessage(r.RawJSON()))
}

type ChunkingStrategyAuto struct {
	Type constant.Auto `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChunkingStrategyAuto) RawJSON() string { return r.JSON.raw }
func (r *ChunkingStrategyAuto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChunkingStrategyStatic struct {
	Static ChunkingStrategyStaticStatic `json:"static,required"`
	Type   constant.Static              `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Static      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChunkingStrategyStatic) RawJSON() string { return r.JSON.raw }
func (r *ChunkingStrategyStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChunkingStrategyStaticStatic struct {
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens,required"`
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkOverlapTokens respjson.Field
		MaxChunkSizeTokens respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChunkingStrategyStaticStatic) RawJSON() string { return r.JSON.raw }
func (r *ChunkingStrategyStaticStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ChunkingStrategyParamOfStatic(static ChunkingStrategyStaticStaticParam) ChunkingStrategyUnionParam {
	var variant ChunkingStrategyStaticParam
	variant.Static = static
	return ChunkingStrategyUnionParam{OfStatic: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChunkingStrategyUnionParam struct {
	OfAuto   *ChunkingStrategyAutoParam   `json:",omitzero,inline"`
	OfStatic *ChunkingStrategyStaticParam `json:",omitzero,inline"`
	paramUnion
}

func (u ChunkingStrategyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAuto, u.OfStatic)
}
func (u *ChunkingStrategyUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChunkingStrategyUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAuto) {
		return u.OfAuto
	} else if !param.IsOmitted(u.OfStatic) {
		return u.OfStatic
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChunkingStrategyUnionParam) GetStatic() *ChunkingStrategyStaticStaticParam {
	if vt := u.OfStatic; vt != nil {
		return &vt.Static
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChunkingStrategyUnionParam) GetType() *string {
	if vt := u.OfAuto; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStatic; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChunkingStrategyUnionParam](
		"type",
		apijson.Discriminator[ChunkingStrategyAutoParam]("auto"),
		apijson.Discriminator[ChunkingStrategyStaticParam]("static"),
	)
}

func NewChunkingStrategyAutoParam() ChunkingStrategyAutoParam {
	return ChunkingStrategyAutoParam{
		Type: "auto",
	}
}

// This struct has a constant value, construct it with
// [NewChunkingStrategyAutoParam].
type ChunkingStrategyAutoParam struct {
	Type constant.Auto `json:"type,required"`
	paramObj
}

func (r ChunkingStrategyAutoParam) MarshalJSON() (data []byte, err error) {
	type shadow ChunkingStrategyAutoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChunkingStrategyAutoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Static, Type are required.
type ChunkingStrategyStaticParam struct {
	Static ChunkingStrategyStaticStaticParam `json:"static,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "static".
	Type constant.Static `json:"type,required"`
	paramObj
}

func (r ChunkingStrategyStaticParam) MarshalJSON() (data []byte, err error) {
	type shadow ChunkingStrategyStaticParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChunkingStrategyStaticParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChunkOverlapTokens, MaxChunkSizeTokens are required.
type ChunkingStrategyStaticStaticParam struct {
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens,required"`
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens,required"`
	paramObj
}

func (r ChunkingStrategyStaticStaticParam) MarshalJSON() (data []byte, err error) {
	type shadow ChunkingStrategyStaticStaticParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChunkingStrategyStaticStaticParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Content struct {
	Text string        `json:"text,required"`
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Content) RawJSON() string { return r.JSON.raw }
func (r *Content) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileStatus string

const (
	FileStatusCompleted  FileStatus = "completed"
	FileStatusInProgress FileStatus = "in_progress"
	FileStatusCancelled  FileStatus = "cancelled"
	FileStatusFailed     FileStatus = "failed"
)

// OpenAI Vector Store File object.
type VectorStoreFile struct {
	ID               string                                   `json:"id,required"`
	Attributes       map[string]VectorStoreFileAttributeUnion `json:"attributes,required"`
	ChunkingStrategy ChunkingStrategyUnion                    `json:"chunking_strategy,required"`
	CreatedAt        int64                                    `json:"created_at,required"`
	Object           string                                   `json:"object,required"`
	// Any of "completed", "in_progress", "cancelled", "failed".
	Status        FileStatus               `json:"status,required"`
	UsageBytes    int64                    `json:"usage_bytes,required"`
	VectorStoreID string                   `json:"vector_store_id,required"`
	LastError     VectorStoreFileLastError `json:"last_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Attributes       respjson.Field
		ChunkingStrategy respjson.Field
		CreatedAt        respjson.Field
		Object           respjson.Field
		Status           respjson.Field
		UsageBytes       respjson.Field
		VectorStoreID    respjson.Field
		LastError        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFile) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VectorStoreFileAttributeUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type VectorStoreFileAttributeUnion struct {
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

func (u VectorStoreFileAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VectorStoreFileAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VectorStoreFileAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *VectorStoreFileAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileLastError struct {
	// Any of "server_error", "rate_limit_exceeded".
	Code    VectorStoreFileLastErrorCode `json:"code,required"`
	Message string                       `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorStoreFileLastError) RawJSON() string { return r.JSON.raw }
func (r *VectorStoreFileLastError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorStoreFileLastErrorCode string

const (
	VectorStoreFileLastErrorCodeServerError       VectorStoreFileLastErrorCode = "server_error"
	VectorStoreFileLastErrorCodeRateLimitExceeded VectorStoreFileLastErrorCode = "rate_limit_exceeded"
)

// Response from listing vector stores.
type OpenAiv1VectorStoreFileListResponse struct {
	Data    []VectorStoreFile `json:"data,required"`
	HasMore bool              `json:"has_more,required"`
	Object  string            `json:"object,required"`
	FirstID string            `json:"first_id"`
	LastID  string            `json:"last_id"`
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
func (r OpenAiv1VectorStoreFileListResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1VectorStoreFileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from deleting a vector store file.
type OpenAiv1VectorStoreFileDeleteResponse struct {
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
func (r OpenAiv1VectorStoreFileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1VectorStoreFileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from retrieving the contents of a vector store file.
type OpenAiv1VectorStoreFileGetContentResponse struct {
	Attributes map[string]OpenAiv1VectorStoreFileGetContentResponseAttributeUnion `json:"attributes,required"`
	Content    []Content                                                          `json:"content,required"`
	FileID     string                                                             `json:"file_id,required"`
	Filename   string                                                             `json:"filename,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		Content     respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1VectorStoreFileGetContentResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1VectorStoreFileGetContentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAiv1VectorStoreFileGetContentResponseAttributeUnion contains all possible
// properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type OpenAiv1VectorStoreFileGetContentResponseAttributeUnion struct {
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

func (u OpenAiv1VectorStoreFileGetContentResponseAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1VectorStoreFileGetContentResponseAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1VectorStoreFileGetContentResponseAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OpenAiv1VectorStoreFileGetContentResponseAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OpenAiv1VectorStoreFileGetContentResponseAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *OpenAiv1VectorStoreFileGetContentResponseAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIV1VectorStoreFileNewParams struct {
	// The ID of the file to attach to the vector store.
	FileID string `json:"file_id,required"`
	// The key-value attributes stored with the file, which can be used for filtering.
	Attributes map[string]OpenAIV1VectorStoreFileNewParamsAttributeUnion `json:"attributes,omitzero"`
	// The chunking strategy to use for the file.
	ChunkingStrategy ChunkingStrategyUnionParam `json:"chunking_strategy,omitzero"`
	paramObj
}

func (r OpenAIV1VectorStoreFileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1VectorStoreFileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1VectorStoreFileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1VectorStoreFileNewParamsAttributeUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1VectorStoreFileNewParamsAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1VectorStoreFileNewParamsAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1VectorStoreFileNewParamsAttributeUnion) asAny() any {
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

type OpenAIV1VectorStoreFileGetParams struct {
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	paramObj
}

type OpenAIV1VectorStoreFileUpdateParams struct {
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	// The updated key-value attributes to store with the file.
	Attributes map[string]OpenAIV1VectorStoreFileUpdateParamsAttributeUnion `json:"attributes,omitzero,required"`
	paramObj
}

func (r OpenAIV1VectorStoreFileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1VectorStoreFileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1VectorStoreFileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1VectorStoreFileUpdateParamsAttributeUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1VectorStoreFileUpdateParamsAttributeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1VectorStoreFileUpdateParamsAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1VectorStoreFileUpdateParamsAttributeUnion) asAny() any {
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

type OpenAIV1VectorStoreFileListParams struct {
	After  param.Opt[string] `query:"after,omitzero" json:"-"`
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Order  param.Opt[string] `query:"order,omitzero" json:"-"`
	// Any of "completed", "in_progress", "cancelled", "failed".
	Filter FileStatus `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OpenAIV1VectorStoreFileListParams]'s query parameters as
// `url.Values`.
func (r OpenAIV1VectorStoreFileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OpenAIV1VectorStoreFileDeleteParams struct {
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	paramObj
}

type OpenAIV1VectorStoreFileGetContentParams struct {
	VectorStoreID string `path:"vector_store_id,required" json:"-"`
	paramObj
}

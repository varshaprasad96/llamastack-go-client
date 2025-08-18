// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/varshaprasad96/llamastack-go-client/internal/apiform"
	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/apiquery"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
	"github.com/varshaprasad96/llamastack-go-client/shared/constant"
)

// OpenAIV1FileService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIV1FileService] method instead.
type OpenAIV1FileService struct {
	Options []option.RequestOption
}

// NewOpenAIV1FileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOpenAIV1FileService(opts ...option.RequestOption) (r OpenAIV1FileService) {
	r = OpenAIV1FileService{}
	r.Options = opts
	return
}

// Upload a file that can be used across various endpoints. The file upload should
// be a multipart form request with:
//
// - file: The File object (not file name) to be uploaded.
// - purpose: The intended purpose of the uploaded file.
func (r *OpenAIV1FileService) New(ctx context.Context, body OpenAIV1FileNewParams, opts ...option.RequestOption) (res *OpenAIFile, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns information about a specific file.
func (r *OpenAIV1FileService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *OpenAIFile, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of files that belong to the user's organization.
func (r *OpenAIV1FileService) List(ctx context.Context, query OpenAIV1FileListParams, opts ...option.RequestOption) (res *OpenAiv1FileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a file.
func (r *OpenAIV1FileService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (res *OpenAiv1FileDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns the contents of the specified file.
func (r *OpenAIV1FileService) GetContent(ctx context.Context, fileID string, opts ...option.RequestOption) (res *OpenAiv1FileGetContentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/files/%s/content", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Valid purpose values for OpenAI Files API.
type FilePurpose string

const (
	FilePurposeAssistants FilePurpose = "assistants"
)

// OpenAI File object as defined in the OpenAI Files API.
type OpenAIFile struct {
	// The file identifier, which can be referenced in the API endpoints
	ID string `json:"id,required"`
	// The size of the file, in bytes
	Bytes int64 `json:"bytes,required"`
	// The Unix timestamp (in seconds) for when the file was created
	CreatedAt int64 `json:"created_at,required"`
	// The Unix timestamp (in seconds) for when the file expires
	ExpiresAt int64 `json:"expires_at,required"`
	// The name of the file
	Filename string `json:"filename,required"`
	// The object type, which is always "file"
	Object constant.File `json:"object,required"`
	// The intended purpose of the file
	//
	// Any of "assistants".
	Purpose OpenAIFilePurpose `json:"purpose,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Bytes       respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Filename    respjson.Field
		Object      respjson.Field
		Purpose     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIFile) RawJSON() string { return r.JSON.raw }
func (r *OpenAIFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The intended purpose of the file
type OpenAIFilePurpose string

const (
	OpenAIFilePurposeAssistants OpenAIFilePurpose = "assistants"
)

// Response for listing files in OpenAI Files API.
type OpenAiv1FileListResponse struct {
	// List of file objects
	Data []OpenAIFile `json:"data,required"`
	// ID of the first file in the list for pagination
	FirstID string `json:"first_id,required"`
	// Whether there are more files available beyond this page
	HasMore bool `json:"has_more,required"`
	// ID of the last file in the list for pagination
	LastID string `json:"last_id,required"`
	// The object type, which is always "list"
	Object constant.List `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		FirstID     respjson.Field
		HasMore     respjson.Field
		LastID      respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1FileListResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1FileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for deleting a file in OpenAI Files API.
type OpenAiv1FileDeleteResponse struct {
	// The file identifier that was deleted
	ID string `json:"id,required"`
	// Whether the file was successfully deleted
	Deleted bool `json:"deleted,required"`
	// The object type, which is always "file"
	Object constant.File `json:"object,required"`
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
func (r OpenAiv1FileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1FileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAiv1FileGetContentResponse = any

type OpenAIV1FileNewParams struct {
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Valid purpose values for OpenAI Files API.
	//
	// Any of "assistants".
	Purpose FilePurpose `json:"purpose,omitzero,required"`
	paramObj
}

func (r OpenAIV1FileNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type OpenAIV1FileListParams struct {
	// A cursor for use in pagination. `after` is an object ID that defines your place
	// in the list. For instance, if you make a list request and receive 100 objects,
	// ending with obj_foo, your subsequent call can include after=obj_foo in order to
	// fetch the next page of the list.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// A limit on the number of objects to be returned. Limit can range between 1 and
	// 10,000, and the default is 10,000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order by the `created_at` timestamp of the objects. `asc` for ascending
	// order and `desc` for descending order.
	//
	// Any of "asc", "desc".
	Order Order `query:"order,omitzero" json:"-"`
	// Only return files with the given purpose.
	//
	// Any of "assistants".
	Purpose FilePurpose `query:"purpose,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OpenAIV1FileListParams]'s query parameters as `url.Values`.
func (r OpenAIV1FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

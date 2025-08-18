// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/llamastack-go-client-go/internal/apijson"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/param"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/respjson"
	"github.com/stainless-sdks/llamastack-go-client-go/shared/constant"
)

// DatasetService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetService] method instead.
type DatasetService struct {
	Options []option.RequestOption
}

// NewDatasetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDatasetService(opts ...option.RequestOption) (r DatasetService) {
	r = DatasetService{}
	r.Options = opts
	return
}

// Register a new dataset.
func (r *DatasetService) New(ctx context.Context, body DatasetNewParams, opts ...option.RequestOption) (res *Dataset, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a dataset by its ID.
func (r *DatasetService) Get(ctx context.Context, datasetID string, opts ...option.RequestOption) (res *Dataset, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/datasets/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all datasets.
func (r *DatasetService) List(ctx context.Context, opts ...option.RequestOption) (res *DatasetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Unregister a dataset by its ID.
func (r *DatasetService) Delete(ctx context.Context, datasetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/datasets/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// DataSourceUnion contains all possible properties and values from
// [DataSourceUri], [DataSourceRows].
//
// Use the [DataSourceUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DataSourceUnion struct {
	// Any of "uri", "rows".
	Type string `json:"type"`
	// This field is from variant [DataSourceUri].
	Uri string `json:"uri"`
	// This field is from variant [DataSourceRows].
	Rows []map[string]DataSourceRowsRowUnion `json:"rows"`
	JSON struct {
		Type respjson.Field
		Uri  respjson.Field
		Rows respjson.Field
		raw  string
	} `json:"-"`
}

// anyDataSource is implemented by each variant of [DataSourceUnion] to add type
// safety for the return type of [DataSourceUnion.AsAny]
type anyDataSource interface {
	implDataSourceUnion()
}

func (DataSourceUri) implDataSourceUnion()  {}
func (DataSourceRows) implDataSourceUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := DataSourceUnion.AsAny().(type) {
//	case llamastackgoclient.DataSourceUri:
//	case llamastackgoclient.DataSourceRows:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u DataSourceUnion) AsAny() anyDataSource {
	switch u.Type {
	case "uri":
		return u.AsUri()
	case "rows":
		return u.AsRows()
	}
	return nil
}

func (u DataSourceUnion) AsUri() (v DataSourceUri) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceUnion) AsRows() (v DataSourceRows) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *DataSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DataSourceUnion to a DataSourceUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DataSourceUnionParam.Overrides()
func (r DataSourceUnion) ToParam() DataSourceUnionParam {
	return param.Override[DataSourceUnionParam](json.RawMessage(r.RawJSON()))
}

// A dataset that can be obtained from a URI.
type DataSourceUri struct {
	Type constant.Uri `json:"type,required"`
	// The dataset can be obtained from a URI. E.g. -
	// "https://mywebsite.com/mydata.jsonl" - "lsfs://mydata.jsonl" -
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataSourceUri) RawJSON() string { return r.JSON.raw }
func (r *DataSourceUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset stored in rows.
type DataSourceRows struct {
	// The dataset is stored in rows. E.g. - [ {"messages": [{"role": "user",
	// "content": "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}
	// ]
	Rows []map[string]DataSourceRowsRowUnion `json:"rows,required"`
	Type constant.Rows                       `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataSourceRows) RawJSON() string { return r.JSON.raw }
func (r *DataSourceRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataSourceRowsRowUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type DataSourceRowsRowUnion struct {
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

func (u DataSourceRowsRowUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceRowsRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceRowsRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceRowsRowUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataSourceRowsRowUnion) RawJSON() string { return u.JSON.raw }

func (r *DataSourceRowsRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func DataSourceParamOfUri(uri string) DataSourceUnionParam {
	var variant DataSourceUriParam
	variant.Uri = uri
	return DataSourceUnionParam{OfUri: &variant}
}

func DataSourceParamOfRows(rows []map[string]DataSourceRowsRowUnionParam) DataSourceUnionParam {
	var variant DataSourceRowsParam
	variant.Rows = rows
	return DataSourceUnionParam{OfRows: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataSourceUnionParam struct {
	OfUri  *DataSourceUriParam  `json:",omitzero,inline"`
	OfRows *DataSourceRowsParam `json:",omitzero,inline"`
	paramUnion
}

func (u DataSourceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUri, u.OfRows)
}
func (u *DataSourceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DataSourceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfUri) {
		return u.OfUri
	} else if !param.IsOmitted(u.OfRows) {
		return u.OfRows
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DataSourceUnionParam) GetUri() *string {
	if vt := u.OfUri; vt != nil {
		return &vt.Uri
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DataSourceUnionParam) GetRows() []map[string]DataSourceRowsRowUnionParam {
	if vt := u.OfRows; vt != nil {
		return vt.Rows
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DataSourceUnionParam) GetType() *string {
	if vt := u.OfUri; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRows; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[DataSourceUnionParam](
		"type",
		apijson.Discriminator[DataSourceUriParam]("uri"),
		apijson.Discriminator[DataSourceRowsParam]("rows"),
	)
}

// A dataset that can be obtained from a URI.
//
// The properties Type, Uri are required.
type DataSourceUriParam struct {
	// The dataset can be obtained from a URI. E.g. -
	// "https://mywebsite.com/mydata.jsonl" - "lsfs://mydata.jsonl" -
	// "data:csv;base64,{base64_content}"
	Uri string `json:"uri,required"`
	// This field can be elided, and will marshal its zero value as "uri".
	Type constant.Uri `json:"type,required"`
	paramObj
}

func (r DataSourceUriParam) MarshalJSON() (data []byte, err error) {
	type shadow DataSourceUriParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataSourceUriParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset stored in rows.
//
// The properties Rows, Type are required.
type DataSourceRowsParam struct {
	// The dataset is stored in rows. E.g. - [ {"messages": [{"role": "user",
	// "content": "Hello, world!"}, {"role": "assistant", "content": "Hello, world!"}]}
	// ]
	Rows []map[string]DataSourceRowsRowUnionParam `json:"rows,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "rows".
	Type constant.Rows `json:"type,required"`
	paramObj
}

func (r DataSourceRowsParam) MarshalJSON() (data []byte, err error) {
	type shadow DataSourceRowsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataSourceRowsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataSourceRowsRowUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u DataSourceRowsRowUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *DataSourceRowsRowUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DataSourceRowsRowUnionParam) asAny() any {
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

// Dataset resource for storing and accessing training or evaluation data.
type Dataset struct {
	Identifier string `json:"identifier,required"`
	// Additional metadata for the dataset
	Metadata   map[string]DatasetMetadataUnion `json:"metadata,required"`
	ProviderID string                          `json:"provider_id,required"`
	// Purpose of the dataset indicating its intended use
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose DatasetPurpose `json:"purpose,required"`
	// Data source configuration for the dataset
	Source DataSourceUnion `json:"source,required"`
	// Type of resource, always 'dataset' for datasets
	Type               constant.Dataset `json:"type,required"`
	ProviderResourceID string           `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		Metadata           respjson.Field
		ProviderID         respjson.Field
		Purpose            respjson.Field
		Source             respjson.Field
		Type               respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Dataset) RawJSON() string { return r.JSON.raw }
func (r *Dataset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DatasetMetadataUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type DatasetMetadataUnion struct {
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

func (u DatasetMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DatasetMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DatasetMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *DatasetMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Purpose of the dataset indicating its intended use
type DatasetPurpose string

const (
	DatasetPurposePostTrainingMessages DatasetPurpose = "post-training/messages"
	DatasetPurposeEvalQuestionAnswer   DatasetPurpose = "eval/question-answer"
	DatasetPurposeEvalMessagesAnswer   DatasetPurpose = "eval/messages-answer"
)

// Response from listing datasets.
type DatasetListResponse struct {
	// List of datasets
	Data []Dataset `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatasetListResponse) RawJSON() string { return r.JSON.raw }
func (r *DatasetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetNewParams struct {
	// The purpose of the dataset. One of: - "post-training/messages": The dataset
	// contains a messages column with list of messages for post-training. {
	// "messages": [ {"role": "user", "content": "Hello, world!"}, {"role":
	// "assistant", "content": "Hello, world!"}, ] } - "eval/question-answer": The
	// dataset contains a question column and an answer column for evaluation. {
	// "question": "What is the capital of France?", "answer": "Paris" } -
	// "eval/messages-answer": The dataset contains a messages column with list of
	// messages and an answer column for evaluation. { "messages": [ {"role": "user",
	// "content": "Hello, my name is John Doe."}, {"role": "assistant", "content":
	// "Hello, John Doe. How can I help you today?"}, {"role": "user", "content":
	// "What's my name?"}, ], "answer": "John Doe" }
	//
	// Any of "post-training/messages", "eval/question-answer", "eval/messages-answer".
	Purpose DatasetNewParamsPurpose `json:"purpose,omitzero,required"`
	// The data source of the dataset. Ensure that the data source schema is compatible
	// with the purpose of the dataset. Examples: - { "type": "uri", "uri":
	// "https://mywebsite.com/mydata.jsonl" } - { "type": "uri", "uri":
	// "lsfs://mydata.jsonl" } - { "type": "uri", "uri":
	// "data:csv;base64,{base64_content}" } - { "type": "uri", "uri":
	// "huggingface://llamastack/simpleqa?split=train" } - { "type": "rows", "rows": [
	// { "messages": [ {"role": "user", "content": "Hello, world!"}, {"role":
	// "assistant", "content": "Hello, world!"}, ] } ] }
	Source DataSourceUnionParam `json:"source,omitzero,required"`
	// The ID of the dataset. If not provided, an ID will be generated.
	DatasetID param.Opt[string] `json:"dataset_id,omitzero"`
	// The metadata for the dataset. - E.g. {"description": "My dataset"}.
	Metadata map[string]DatasetNewParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r DatasetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DatasetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The purpose of the dataset. One of: - "post-training/messages": The dataset
// contains a messages column with list of messages for post-training. {
// "messages": [ {"role": "user", "content": "Hello, world!"}, {"role":
// "assistant", "content": "Hello, world!"}, ] } - "eval/question-answer": The
// dataset contains a question column and an answer column for evaluation. {
// "question": "What is the capital of France?", "answer": "Paris" } -
// "eval/messages-answer": The dataset contains a messages column with list of
// messages and an answer column for evaluation. { "messages": [ {"role": "user",
// "content": "Hello, my name is John Doe."}, {"role": "assistant", "content":
// "Hello, John Doe. How can I help you today?"}, {"role": "user", "content":
// "What's my name?"}, ], "answer": "John Doe" }
type DatasetNewParamsPurpose string

const (
	DatasetNewParamsPurposePostTrainingMessages DatasetNewParamsPurpose = "post-training/messages"
	DatasetNewParamsPurposeEvalQuestionAnswer   DatasetNewParamsPurpose = "eval/question-answer"
	DatasetNewParamsPurposeEvalMessagesAnswer   DatasetNewParamsPurpose = "eval/messages-answer"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DatasetNewParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u DatasetNewParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *DatasetNewParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DatasetNewParamsMetadataUnion) asAny() any {
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

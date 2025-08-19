// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// ModelService contains methods and other services that help with interacting with
// the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.Options = opts
	return
}

// Register a model.
func (r *ModelService) New(ctx context.Context, body ModelNewParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a model by its identifier.
func (r *ModelService) Get(ctx context.Context, modelID string, opts ...option.RequestOption) (res *Model, err error) {
	opts = append(r.Options[:], opts...)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("v1/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all models.
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *ModelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Unregister a model.
func (r *ModelService) Delete(ctx context.Context, modelID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("v1/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Model struct {
	Identifier string                        `json:"identifier,required"`
	Metadata   map[string]ModelMetadataUnion `json:"metadata,required"`
	// Any of "llm", "embedding".
	ModelType  ModelType `json:"model_type,required"`
	ProviderID string    `json:"provider_id,required"`
	// Any of "model", "shield", "vector_db", "dataset", "scoring_function",
	// "benchmark", "tool", "tool_group".
	Type               ModelType `json:"type,required"`
	ProviderResourceID string    `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		Metadata           respjson.Field
		ModelType          respjson.Field
		ProviderID         respjson.Field
		Type               respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ModelMetadataUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ModelMetadataUnion struct {
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

func (u ModelMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModelMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ModelMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelType string

const (
	ModelTypeModel           ModelType = "model"
	ModelTypeShield          ModelType = "shield"
	ModelTypeVectorDB        ModelType = "vector_db"
	ModelTypeDataset         ModelType = "dataset"
	ModelTypeScoringFunction ModelType = "scoring_function"
	ModelTypeBenchmark       ModelType = "benchmark"
	ModelTypeTool            ModelType = "tool"
	ModelTypeToolGroup       ModelType = "tool_group"
)

type ModelListResponse struct {
	Data []Model `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelNewParams struct {
	// The identifier of the model to register.
	ModelID string `json:"model_id,required"`
	// The identifier of the provider.
	ProviderID param.Opt[string] `json:"provider_id,omitzero"`
	// The identifier of the model in the provider.
	ProviderModelID param.Opt[string] `json:"provider_model_id,omitzero"`
	// Any additional metadata for this model.
	Metadata map[string]ModelNewParamsMetadataUnion `json:"metadata,omitzero"`
	// The type of model to register.
	//
	// Any of "llm", "embedding".
	ModelType ModelType `json:"model_type,omitzero"`
	paramObj
}

func (r ModelNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ModelNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ModelNewParamsMetadataUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ModelNewParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ModelNewParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ModelNewParamsMetadataUnion) asAny() any {
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

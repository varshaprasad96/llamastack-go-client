// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/apiquery"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// PostTrainingJobService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostTrainingJobService] method instead.
type PostTrainingJobService struct {
	Options []option.RequestOption
}

// NewPostTrainingJobService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPostTrainingJobService(opts ...option.RequestOption) (r PostTrainingJobService) {
	r = PostTrainingJobService{}
	r.Options = opts
	return
}

// Cancel a training job.
func (r *PostTrainingJobService) Cancel(ctx context.Context, body PostTrainingJobCancelParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/post-training/job/cancel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get the artifacts of a training job.
func (r *PostTrainingJobService) GetArtifacts(ctx context.Context, query PostTrainingJobGetArtifactsParams, opts ...option.RequestOption) (res *PostTrainingJobGetArtifactsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/post-training/job/artifacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the status of a training job.
func (r *PostTrainingJobService) GetStatus(ctx context.Context, query PostTrainingJobGetStatusParams, opts ...option.RequestOption) (res *PostTrainingJobGetStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/post-training/job/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Checkpoint created during training runs.
type Checkpoint struct {
	// Timestamp when the checkpoint was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Training epoch when the checkpoint was saved
	Epoch int64 `json:"epoch,required"`
	// Unique identifier for the checkpoint
	Identifier string `json:"identifier,required"`
	// File system path where the checkpoint is stored
	Path string `json:"path,required"`
	// Identifier of the training job that created this checkpoint
	PostTrainingJobID string `json:"post_training_job_id,required"`
	// (Optional) Training metrics associated with this checkpoint
	TrainingMetrics CheckpointTrainingMetrics `json:"training_metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt         respjson.Field
		Epoch             respjson.Field
		Identifier        respjson.Field
		Path              respjson.Field
		PostTrainingJobID respjson.Field
		TrainingMetrics   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Checkpoint) RawJSON() string { return r.JSON.raw }
func (r *Checkpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Training metrics associated with this checkpoint
type CheckpointTrainingMetrics struct {
	// Training epoch number
	Epoch int64 `json:"epoch,required"`
	// Perplexity metric indicating model confidence
	Perplexity float64 `json:"perplexity,required"`
	// Loss value on the training dataset
	TrainLoss float64 `json:"train_loss,required"`
	// Loss value on the validation dataset
	ValidationLoss float64 `json:"validation_loss,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Epoch          respjson.Field
		Perplexity     respjson.Field
		TrainLoss      respjson.Field
		ValidationLoss respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckpointTrainingMetrics) RawJSON() string { return r.JSON.raw }
func (r *CheckpointTrainingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Artifacts of a finetuning job.
type PostTrainingJobGetArtifactsResponse struct {
	// List of model checkpoints created during training
	Checkpoints []Checkpoint `json:"checkpoints,required"`
	// Unique identifier for the training job
	JobUuid string `json:"job_uuid,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checkpoints respjson.Field
		JobUuid     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostTrainingJobGetArtifactsResponse) RawJSON() string { return r.JSON.raw }
func (r *PostTrainingJobGetArtifactsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a finetuning job.
type PostTrainingJobGetStatusResponse struct {
	// List of model checkpoints created during training
	Checkpoints []Checkpoint `json:"checkpoints,required"`
	// Unique identifier for the training job
	JobUuid string `json:"job_uuid,required"`
	// Current status of the training job
	//
	// Any of "completed", "in_progress", "failed", "scheduled", "cancelled".
	Status PostTrainingJobGetStatusResponseStatus `json:"status,required"`
	// (Optional) Timestamp when the job finished, if completed
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// (Optional) Information about computational resources allocated to the job
	ResourcesAllocated map[string]PostTrainingJobGetStatusResponseResourcesAllocatedUnion `json:"resources_allocated"`
	// (Optional) Timestamp when the job was scheduled
	ScheduledAt time.Time `json:"scheduled_at" format:"date-time"`
	// (Optional) Timestamp when the job execution began
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checkpoints        respjson.Field
		JobUuid            respjson.Field
		Status             respjson.Field
		CompletedAt        respjson.Field
		ResourcesAllocated respjson.Field
		ScheduledAt        respjson.Field
		StartedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostTrainingJobGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *PostTrainingJobGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the training job
type PostTrainingJobGetStatusResponseStatus string

const (
	PostTrainingJobGetStatusResponseStatusCompleted  PostTrainingJobGetStatusResponseStatus = "completed"
	PostTrainingJobGetStatusResponseStatusInProgress PostTrainingJobGetStatusResponseStatus = "in_progress"
	PostTrainingJobGetStatusResponseStatusFailed     PostTrainingJobGetStatusResponseStatus = "failed"
	PostTrainingJobGetStatusResponseStatusScheduled  PostTrainingJobGetStatusResponseStatus = "scheduled"
	PostTrainingJobGetStatusResponseStatusCancelled  PostTrainingJobGetStatusResponseStatus = "cancelled"
)

// PostTrainingJobGetStatusResponseResourcesAllocatedUnion contains all possible
// properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type PostTrainingJobGetStatusResponseResourcesAllocatedUnion struct {
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

func (u PostTrainingJobGetStatusResponseResourcesAllocatedUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PostTrainingJobGetStatusResponseResourcesAllocatedUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PostTrainingJobGetStatusResponseResourcesAllocatedUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PostTrainingJobGetStatusResponseResourcesAllocatedUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PostTrainingJobGetStatusResponseResourcesAllocatedUnion) RawJSON() string { return u.JSON.raw }

func (r *PostTrainingJobGetStatusResponseResourcesAllocatedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostTrainingJobCancelParams struct {
	// The UUID of the job to cancel.
	JobUuid string `json:"job_uuid,required"`
	paramObj
}

func (r PostTrainingJobCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingJobCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingJobCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostTrainingJobGetArtifactsParams struct {
	// The UUID of the job to get the artifacts of.
	JobUuid string `query:"job_uuid,required" json:"-"`
	paramObj
}

// URLQuery serializes [PostTrainingJobGetArtifactsParams]'s query parameters as
// `url.Values`.
func (r PostTrainingJobGetArtifactsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PostTrainingJobGetStatusParams struct {
	// The UUID of the job to get the status of.
	JobUuid string `query:"job_uuid,required" json:"-"`
	paramObj
}

// URLQuery serializes [PostTrainingJobGetStatusParams]'s query parameters as
// `url.Values`.
func (r PostTrainingJobGetStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

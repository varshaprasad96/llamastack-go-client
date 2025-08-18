// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"time"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/option"
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
	"github.com/varshaprasad96/llamastack-go-client/shared/constant"
)

// PostTrainingService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostTrainingService] method instead.
type PostTrainingService struct {
	Options []option.RequestOption
	Job     PostTrainingJobService
}

// NewPostTrainingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPostTrainingService(opts ...option.RequestOption) (r PostTrainingService) {
	r = PostTrainingService{}
	r.Options = opts
	r.Job = NewPostTrainingJobService(opts...)
	return
}

// Run supervised fine-tuning of a model.
func (r *PostTrainingService) FineTuneSupervised(ctx context.Context, body PostTrainingFineTuneSupervisedParams, opts ...option.RequestOption) (res *PostTrainingJob, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/post-training/supervised-fine-tune"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get all training jobs.
func (r *PostTrainingService) ListJobs(ctx context.Context, opts ...option.RequestOption) (res *PostTrainingListJobsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/post-training/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Run preference optimization of a model.
func (r *PostTrainingService) OptimizePreferences(ctx context.Context, body PostTrainingOptimizePreferencesParams, opts ...option.RequestOption) (res *PostTrainingJob, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/post-training/preference-optimize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The properties BatchSize, DataFormat, DatasetID, Shuffle are required.
type DataConfigParam struct {
	BatchSize int64 `json:"batch_size,required"`
	// Any of "instruct", "dialog".
	DataFormat          DataConfigDataFormat `json:"data_format,omitzero,required"`
	DatasetID           string               `json:"dataset_id,required"`
	Shuffle             bool                 `json:"shuffle,required"`
	Packed              param.Opt[bool]      `json:"packed,omitzero"`
	TrainOnInput        param.Opt[bool]      `json:"train_on_input,omitzero"`
	ValidationDatasetID param.Opt[string]    `json:"validation_dataset_id,omitzero"`
	paramObj
}

func (r DataConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow DataConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataConfigDataFormat string

const (
	DataConfigDataFormatInstruct DataConfigDataFormat = "instruct"
	DataConfigDataFormatDialog   DataConfigDataFormat = "dialog"
)

type EfficiencyConfigParam struct {
	EnableActivationCheckpointing param.Opt[bool] `json:"enable_activation_checkpointing,omitzero"`
	EnableActivationOffloading    param.Opt[bool] `json:"enable_activation_offloading,omitzero"`
	FsdpCPUOffload                param.Opt[bool] `json:"fsdp_cpu_offload,omitzero"`
	MemoryEfficientFsdpWrap       param.Opt[bool] `json:"memory_efficient_fsdp_wrap,omitzero"`
	paramObj
}

func (r EfficiencyConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow EfficiencyConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EfficiencyConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Lr, NumWarmupSteps, OptimizerType, WeightDecay are required.
type OptimizerConfigParam struct {
	Lr             float64 `json:"lr,required"`
	NumWarmupSteps int64   `json:"num_warmup_steps,required"`
	// Any of "adam", "adamw", "sgd".
	OptimizerType OptimizerConfigOptimizerType `json:"optimizer_type,omitzero,required"`
	WeightDecay   float64                      `json:"weight_decay,required"`
	paramObj
}

func (r OptimizerConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow OptimizerConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizerConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptimizerConfigOptimizerType string

const (
	OptimizerConfigOptimizerTypeAdam  OptimizerConfigOptimizerType = "adam"
	OptimizerConfigOptimizerTypeAdamw OptimizerConfigOptimizerType = "adamw"
	OptimizerConfigOptimizerTypeSgd   OptimizerConfigOptimizerType = "sgd"
)

type PostTrainingJob struct {
	JobUuid string `json:"job_uuid,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobUuid     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostTrainingJob) RawJSON() string { return r.JSON.raw }
func (r *PostTrainingJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties GradientAccumulationSteps, MaxStepsPerEpoch, NEpochs are
// required.
type TrainingConfigParam struct {
	GradientAccumulationSteps int64                 `json:"gradient_accumulation_steps,required"`
	MaxStepsPerEpoch          int64                 `json:"max_steps_per_epoch,required"`
	NEpochs                   int64                 `json:"n_epochs,required"`
	Dtype                     param.Opt[string]     `json:"dtype,omitzero"`
	MaxValidationSteps        param.Opt[int64]      `json:"max_validation_steps,omitzero"`
	DataConfig                DataConfigParam       `json:"data_config,omitzero"`
	EfficiencyConfig          EfficiencyConfigParam `json:"efficiency_config,omitzero"`
	OptimizerConfig           OptimizerConfigParam  `json:"optimizer_config,omitzero"`
	paramObj
}

func (r TrainingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow TrainingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrainingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostTrainingListJobsResponse struct {
	Data []PostTrainingListJobsResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostTrainingListJobsResponse) RawJSON() string { return r.JSON.raw }
func (r *PostTrainingListJobsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostTrainingListJobsResponseData struct {
	JobUuid string `json:"job_uuid,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobUuid     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostTrainingListJobsResponseData) RawJSON() string { return r.JSON.raw }
func (r *PostTrainingListJobsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostTrainingFineTuneSupervisedParams struct {
	// The hyperparam search configuration.
	HyperparamSearchConfig map[string]PostTrainingFineTuneSupervisedParamsHyperparamSearchConfigUnion `json:"hyperparam_search_config,omitzero,required"`
	// The UUID of the job to create.
	JobUuid string `json:"job_uuid,required"`
	// The logger configuration.
	LoggerConfig map[string]PostTrainingFineTuneSupervisedParamsLoggerConfigUnion `json:"logger_config,omitzero,required"`
	// The training configuration.
	TrainingConfig TrainingConfigParam `json:"training_config,omitzero,required"`
	// The directory to save checkpoint(s) to.
	CheckpointDir param.Opt[string] `json:"checkpoint_dir,omitzero"`
	// The model to fine-tune.
	Model param.Opt[string] `json:"model,omitzero"`
	// The algorithm configuration.
	AlgorithmConfig PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion `json:"algorithm_config,omitzero"`
	paramObj
}

func (r PostTrainingFineTuneSupervisedParams) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingFineTuneSupervisedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingFineTuneSupervisedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PostTrainingFineTuneSupervisedParamsHyperparamSearchConfigUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u PostTrainingFineTuneSupervisedParamsHyperparamSearchConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *PostTrainingFineTuneSupervisedParamsHyperparamSearchConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PostTrainingFineTuneSupervisedParamsHyperparamSearchConfigUnion) asAny() any {
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
type PostTrainingFineTuneSupervisedParamsLoggerConfigUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u PostTrainingFineTuneSupervisedParamsLoggerConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *PostTrainingFineTuneSupervisedParamsLoggerConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PostTrainingFineTuneSupervisedParamsLoggerConfigUnion) asAny() any {
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
type PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion struct {
	OfLoRa *PostTrainingFineTuneSupervisedParamsAlgorithmConfigLoRa `json:",omitzero,inline"`
	OfQat  *PostTrainingFineTuneSupervisedParamsAlgorithmConfigQat  `json:",omitzero,inline"`
	paramUnion
}

func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLoRa, u.OfQat)
}
func (u *PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfLoRa) {
		return u.OfLoRa
	} else if !param.IsOmitted(u.OfQat) {
		return u.OfQat
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) GetAlpha() *int64 {
	if vt := u.OfLoRa; vt != nil {
		return &vt.Alpha
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) GetApplyLoraToMlp() *bool {
	if vt := u.OfLoRa; vt != nil {
		return &vt.ApplyLoraToMlp
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) GetApplyLoraToOutput() *bool {
	if vt := u.OfLoRa; vt != nil {
		return &vt.ApplyLoraToOutput
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) GetLoraAttnModules() []string {
	if vt := u.OfLoRa; vt != nil {
		return vt.LoraAttnModules
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) GetRank() *int64 {
	if vt := u.OfLoRa; vt != nil {
		return &vt.Rank
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) GetQuantizeBase() *bool {
	if vt := u.OfLoRa; vt != nil && vt.QuantizeBase.Valid() {
		return &vt.QuantizeBase.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) GetUseDora() *bool {
	if vt := u.OfLoRa; vt != nil && vt.UseDora.Valid() {
		return &vt.UseDora.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) GetGroupSize() *int64 {
	if vt := u.OfQat; vt != nil {
		return &vt.GroupSize
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) GetQuantizerName() *string {
	if vt := u.OfQat; vt != nil {
		return &vt.QuantizerName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion) GetType() *string {
	if vt := u.OfLoRa; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfQat; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion](
		"type",
		apijson.Discriminator[PostTrainingFineTuneSupervisedParamsAlgorithmConfigLoRa]("LoRA"),
		apijson.Discriminator[PostTrainingFineTuneSupervisedParamsAlgorithmConfigQat]("QAT"),
	)
}

// The properties Alpha, ApplyLoraToMlp, ApplyLoraToOutput, LoraAttnModules, Rank,
// Type are required.
type PostTrainingFineTuneSupervisedParamsAlgorithmConfigLoRa struct {
	Alpha             int64           `json:"alpha,required"`
	ApplyLoraToMlp    bool            `json:"apply_lora_to_mlp,required"`
	ApplyLoraToOutput bool            `json:"apply_lora_to_output,required"`
	LoraAttnModules   []string        `json:"lora_attn_modules,omitzero,required"`
	Rank              int64           `json:"rank,required"`
	QuantizeBase      param.Opt[bool] `json:"quantize_base,omitzero"`
	UseDora           param.Opt[bool] `json:"use_dora,omitzero"`
	// This field can be elided, and will marshal its zero value as "LoRA".
	Type constant.LoRa `json:"type,required"`
	paramObj
}

func (r PostTrainingFineTuneSupervisedParamsAlgorithmConfigLoRa) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingFineTuneSupervisedParamsAlgorithmConfigLoRa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingFineTuneSupervisedParamsAlgorithmConfigLoRa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties GroupSize, QuantizerName, Type are required.
type PostTrainingFineTuneSupervisedParamsAlgorithmConfigQat struct {
	GroupSize     int64  `json:"group_size,required"`
	QuantizerName string `json:"quantizer_name,required"`
	// This field can be elided, and will marshal its zero value as "QAT".
	Type constant.Qat `json:"type,required"`
	paramObj
}

func (r PostTrainingFineTuneSupervisedParamsAlgorithmConfigQat) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingFineTuneSupervisedParamsAlgorithmConfigQat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingFineTuneSupervisedParamsAlgorithmConfigQat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostTrainingOptimizePreferencesParams struct {
	// The algorithm configuration.
	AlgorithmConfig PostTrainingOptimizePreferencesParamsAlgorithmConfig `json:"algorithm_config,omitzero,required"`
	// The model to fine-tune.
	FinetunedModel string `json:"finetuned_model,required"`
	// The hyperparam search configuration.
	HyperparamSearchConfig map[string]PostTrainingOptimizePreferencesParamsHyperparamSearchConfigUnion `json:"hyperparam_search_config,omitzero,required"`
	// The UUID of the job to create.
	JobUuid string `json:"job_uuid,required"`
	// The logger configuration.
	LoggerConfig map[string]PostTrainingOptimizePreferencesParamsLoggerConfigUnion `json:"logger_config,omitzero,required"`
	// The training configuration.
	TrainingConfig TrainingConfigParam `json:"training_config,omitzero,required"`
	paramObj
}

func (r PostTrainingOptimizePreferencesParams) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingOptimizePreferencesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingOptimizePreferencesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The algorithm configuration.
//
// The properties Epsilon, Gamma, RewardClip, RewardScale are required.
type PostTrainingOptimizePreferencesParamsAlgorithmConfig struct {
	Epsilon     float64 `json:"epsilon,required"`
	Gamma       float64 `json:"gamma,required"`
	RewardClip  float64 `json:"reward_clip,required"`
	RewardScale float64 `json:"reward_scale,required"`
	paramObj
}

func (r PostTrainingOptimizePreferencesParamsAlgorithmConfig) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingOptimizePreferencesParamsAlgorithmConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingOptimizePreferencesParamsAlgorithmConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PostTrainingOptimizePreferencesParamsHyperparamSearchConfigUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u PostTrainingOptimizePreferencesParamsHyperparamSearchConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *PostTrainingOptimizePreferencesParamsHyperparamSearchConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PostTrainingOptimizePreferencesParamsHyperparamSearchConfigUnion) asAny() any {
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
type PostTrainingOptimizePreferencesParamsLoggerConfigUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u PostTrainingOptimizePreferencesParamsLoggerConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *PostTrainingOptimizePreferencesParamsLoggerConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PostTrainingOptimizePreferencesParamsLoggerConfigUnion) asAny() any {
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

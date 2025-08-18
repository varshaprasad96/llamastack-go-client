// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/llamastack-go-client-go/internal/apijson"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/param"
	"github.com/stainless-sdks/llamastack-go-client-go/packages/respjson"
)

// EvalBenchmarkJobService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalBenchmarkJobService] method instead.
type EvalBenchmarkJobService struct {
	Options []option.RequestOption
}

// NewEvalBenchmarkJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEvalBenchmarkJobService(opts ...option.RequestOption) (r EvalBenchmarkJobService) {
	r = EvalBenchmarkJobService{}
	r.Options = opts
	return
}

// Get the status of a job.
func (r *EvalBenchmarkJobService) Get(ctx context.Context, jobID string, query EvalBenchmarkJobGetParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = append(r.Options[:], opts...)
	if query.BenchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/jobs/%s", query.BenchmarkID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel a job.
func (r *EvalBenchmarkJobService) Cancel(ctx context.Context, jobID string, body EvalBenchmarkJobCancelParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.BenchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/jobs/%s", body.BenchmarkID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get the result of a job.
func (r *EvalBenchmarkJobService) Result(ctx context.Context, jobID string, query EvalBenchmarkJobResultParams, opts ...option.RequestOption) (res *EvaluateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.BenchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/jobs/%s/result", query.BenchmarkID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Run an evaluation on a benchmark.
func (r *EvalBenchmarkJobService) Run(ctx context.Context, benchmarkID string, body EvalBenchmarkJobRunParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = append(r.Options[:], opts...)
	if benchmarkID == "" {
		err = errors.New("missing required benchmark_id parameter")
		return
	}
	path := fmt.Sprintf("v1/eval/benchmarks/%s/jobs", benchmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A job execution instance with status tracking.
type Job struct {
	// Unique identifier for the job
	JobID string `json:"job_id,required"`
	// Current execution status of the job
	//
	// Any of "completed", "in_progress", "failed", "scheduled", "cancelled".
	Status JobStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Job) RawJSON() string { return r.JSON.raw }
func (r *Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current execution status of the job
type JobStatus string

const (
	JobStatusCompleted  JobStatus = "completed"
	JobStatusInProgress JobStatus = "in_progress"
	JobStatusFailed     JobStatus = "failed"
	JobStatusScheduled  JobStatus = "scheduled"
	JobStatusCancelled  JobStatus = "cancelled"
)

type EvalBenchmarkJobGetParams struct {
	BenchmarkID string `path:"benchmark_id,required" json:"-"`
	paramObj
}

type EvalBenchmarkJobCancelParams struct {
	BenchmarkID string `path:"benchmark_id,required" json:"-"`
	paramObj
}

type EvalBenchmarkJobResultParams struct {
	BenchmarkID string `path:"benchmark_id,required" json:"-"`
	paramObj
}

type EvalBenchmarkJobRunParams struct {
	// The configuration for the benchmark.
	BenchmarkConfig BenchmarkConfigParam `json:"benchmark_config,omitzero,required"`
	paramObj
}

func (r EvalBenchmarkJobRunParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalBenchmarkJobRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalBenchmarkJobRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

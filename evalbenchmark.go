// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/varshaprasad96/llamastack-go-client/option"
)

// EvalBenchmarkService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalBenchmarkService] method instead.
type EvalBenchmarkService struct {
	Options []option.RequestOption
	Jobs    EvalBenchmarkJobService
}

// NewEvalBenchmarkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvalBenchmarkService(opts ...option.RequestOption) (r EvalBenchmarkService) {
	r = EvalBenchmarkService{}
	r.Options = opts
	r.Jobs = NewEvalBenchmarkJobService(opts...)
	return
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/varshaprasad96/llamastack-go-client/option"
)

// ShieldService contains methods and other services that help with interacting
// with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewShieldService] method instead.
type ShieldService struct {
	Options []option.RequestOption
}

// NewShieldService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewShieldService(opts ...option.RequestOption) (r ShieldService) {
	r = ShieldService{}
	r.Options = opts
	return
}

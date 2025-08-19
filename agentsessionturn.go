// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	shimjson "github.com/varshaprasad96/llamastack-go-client/internal/encoding/json"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
)

// AgentSessionTurnService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentSessionTurnService] method instead.
type AgentSessionTurnService struct {
	Options []option.RequestOption
	Step    AgentSessionTurnStepService
}

// NewAgentSessionTurnService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgentSessionTurnService(opts ...option.RequestOption) (r AgentSessionTurnService) {
	r = AgentSessionTurnService{}
	r.Options = opts
	r.Step = NewAgentSessionTurnStepService(opts...)
	return
}

// Create a new turn for an agent.
func (r *AgentSessionTurnService) New(ctx context.Context, sessionID string, params AgentSessionTurnNewParams, opts ...option.RequestOption) (res *AgentSessionTurnNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s/turn", params.AgentID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AgentSessionTurnNewResponse = any

type AgentSessionTurnNewParams struct {
	AgentID string `path:"agent_id,required" json:"-"`
	Body    any
	paramObj
}

func (r AgentSessionTurnNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AgentSessionTurnNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

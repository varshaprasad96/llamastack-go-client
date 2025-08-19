// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/varshaprasad96/llamastack-go-client/internal/apiquery"
	shimjson "github.com/varshaprasad96/llamastack-go-client/internal/encoding/json"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
)

// AgentService contains methods and other services that help with interacting with
// the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
	Session AgentSessionService
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.Options = opts
	r.Session = NewAgentSessionService(opts...)
	return
}

// Create an agent with the given configuration.
func (r *AgentService) New(ctx context.Context, body AgentNewParams, opts ...option.RequestOption) (res *AgentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all agents.
func (r *AgentService) List(ctx context.Context, query AgentListParams, opts ...option.RequestOption) (res *AgentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create a new session for an agent.
func (r *AgentService) NewSession(ctx context.Context, agentID string, body AgentNewSessionParams, opts ...option.RequestOption) (res *AgentNewSessionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AgentNewResponse = any

type AgentListResponse = any

type AgentNewSessionResponse = any

type AgentNewParams struct {
	Body any
	paramObj
}

func (r AgentNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AgentNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type AgentListParams struct {
	// The number of agents to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index to start the pagination from.
	StartIndex param.Opt[int64] `query:"start_index,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentListParams]'s query parameters as `url.Values`.
func (r AgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentNewSessionParams struct {
	Body any
	paramObj
}

func (r AgentNewSessionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AgentNewSessionParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

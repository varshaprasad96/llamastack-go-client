// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/apiquery"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// AgentSessionService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentSessionService] method instead.
type AgentSessionService struct {
	Options []option.RequestOption
	Turn    AgentSessionTurnService
}

// NewAgentSessionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentSessionService(opts ...option.RequestOption) (r AgentSessionService) {
	r = AgentSessionService{}
	r.Options = opts
	r.Turn = NewAgentSessionTurnService(opts...)
	return
}

// Retrieve an agent session by its ID.
func (r *AgentSessionService) Get(ctx context.Context, sessionID string, params AgentSessionGetParams, opts ...option.RequestOption) (res *AgentSessionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s", params.AgentID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete an agent session by its ID and its associated turns.
func (r *AgentSessionService) Delete(ctx context.Context, sessionID string, body AgentSessionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s", body.AgentID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// A single session of an interaction with an Agentic System.
type AgentSessionGetResponse struct {
	// Unique identifier for the conversation session
	SessionID string `json:"session_id,required"`
	// Human-readable name for the session
	SessionName string `json:"session_name,required"`
	// Timestamp when the session was created
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// List of all turns that have occurred in this session
	Turns []Turn `json:"turns,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SessionID   respjson.Field
		SessionName respjson.Field
		StartedAt   respjson.Field
		Turns       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentSessionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentSessionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentSessionGetParams struct {
	AgentID string `path:"agent_id,required" json:"-"`
	// (Optional) List of turn IDs to filter the session by.
	TurnIDs []string `query:"turn_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentSessionGetParams]'s query parameters as `url.Values`.
func (r AgentSessionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentSessionDeleteParams struct {
	AgentID string `path:"agent_id,required" json:"-"`
	paramObj
}

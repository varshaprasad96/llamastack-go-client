// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
)

// AgentSessionTurnStepService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentSessionTurnStepService] method instead.
type AgentSessionTurnStepService struct {
	Options []option.RequestOption
}

// NewAgentSessionTurnStepService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgentSessionTurnStepService(opts ...option.RequestOption) (r AgentSessionTurnStepService) {
	r = AgentSessionTurnStepService{}
	r.Options = opts
	return
}

// Retrieve an agent step by its ID.
func (r *AgentSessionTurnStepService) Get(ctx context.Context, stepID string, query AgentSessionTurnStepGetParams, opts ...option.RequestOption) (res *AgentSessionTurnStepGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	if query.TurnID == "" {
		err = errors.New("missing required turn_id parameter")
		return
	}
	if stepID == "" {
		err = errors.New("missing required step_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s/turn/%s/step/%s", query.AgentID, query.SessionID, query.TurnID, stepID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AgentSessionTurnStepGetResponse struct {
	// An inference step in an agent turn.
	Step AgentSessionTurnStepGetResponseStepUnion `json:"step,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Step        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentSessionTurnStepGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentSessionTurnStepGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentSessionTurnStepGetResponseStepUnion contains all possible properties and
// values from [InferenceStep], [ToolExecutionStep], [ShieldCallStep],
// [MemoryRetrievalStep].
//
// Use the [AgentSessionTurnStepGetResponseStepUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AgentSessionTurnStepGetResponseStepUnion struct {
	// This field is from variant [InferenceStep].
	ModelResponse CompletionMessage `json:"model_response"`
	StepID        string            `json:"step_id"`
	// Any of nil, nil, nil, nil.
	StepType    string    `json:"step_type"`
	TurnID      string    `json:"turn_id"`
	CompletedAt time.Time `json:"completed_at"`
	StartedAt   time.Time `json:"started_at"`
	// This field is from variant [ToolExecutionStep].
	ToolCalls []ToolCall `json:"tool_calls"`
	// This field is from variant [ToolExecutionStep].
	ToolResponses []ToolResponse `json:"tool_responses"`
	// This field is from variant [ShieldCallStep].
	Violation SafetyViolation `json:"violation"`
	// This field is from variant [MemoryRetrievalStep].
	InsertedContext InterleavedContentUnion `json:"inserted_context"`
	// This field is from variant [MemoryRetrievalStep].
	VectorDBIDs string `json:"vector_db_ids"`
	JSON        struct {
		ModelResponse   respjson.Field
		StepID          respjson.Field
		StepType        respjson.Field
		TurnID          respjson.Field
		CompletedAt     respjson.Field
		StartedAt       respjson.Field
		ToolCalls       respjson.Field
		ToolResponses   respjson.Field
		Violation       respjson.Field
		InsertedContext respjson.Field
		VectorDBIDs     respjson.Field
		raw             string
	} `json:"-"`
}

func (u AgentSessionTurnStepGetResponseStepUnion) AsInferenceStep() (v InferenceStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentSessionTurnStepGetResponseStepUnion) AsToolExecutionStep() (v ToolExecutionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentSessionTurnStepGetResponseStepUnion) AsShieldCallStep() (v ShieldCallStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentSessionTurnStepGetResponseStepUnion) AsMemoryRetrievalStep() (v MemoryRetrievalStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentSessionTurnStepGetResponseStepUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentSessionTurnStepGetResponseStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentSessionTurnStepGetParams struct {
	AgentID   string `path:"agent_id,required" json:"-"`
	SessionID string `path:"session_id,required" json:"-"`
	TurnID    string `path:"turn_id,required" json:"-"`
	paramObj
}

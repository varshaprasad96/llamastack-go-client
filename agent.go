// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

// Describe an agent by its ID.
func (r *AgentService) Get(ctx context.Context, agentID string, opts ...option.RequestOption) (res *AgentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all agents.
func (r *AgentService) List(ctx context.Context, query AgentListParams, opts ...option.RequestOption) (res *PaginatedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an agent by its ID and its associated sessions and turns.
func (r *AgentService) Delete(ctx context.Context, agentID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
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

// List all session(s) of a given agent.
func (r *AgentService) GetSessions(ctx context.Context, agentID string, query AgentGetSessionsParams, opts ...option.RequestOption) (res *PaginatedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/sessions", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Configuration for an agent.
type AgentConfig struct {
	// The system instructions for the agent
	Instructions string `json:"instructions,required"`
	// The model identifier to use for the agent
	Model       string    `json:"model,required"`
	ClientTools []ToolDef `json:"client_tools"`
	// Optional flag indicating whether session data has to be persisted
	EnableSessionPersistence bool     `json:"enable_session_persistence"`
	InputShields             []string `json:"input_shields"`
	MaxInferIters            int64    `json:"max_infer_iters"`
	// Optional name for the agent, used in telemetry and identification
	Name          string   `json:"name"`
	OutputShields []string `json:"output_shields"`
	// Optional response format configuration
	ResponseFormat ResponseFormatUnion `json:"response_format"`
	// Sampling parameters.
	SamplingParams SamplingParamsResp `json:"sampling_params"`
	// Whether tool use is required or automatic. This is a hint to the model which may
	// not be followed. It depends on the Instruction Following capabilities of the
	// model.
	//
	// Any of "auto", "required", "none".
	//
	// Deprecated: deprecated
	ToolChoice AgentConfigToolChoice `json:"tool_choice"`
	// Configuration for tool use.
	ToolConfig ToolConfig `json:"tool_config"`
	// Prompt format for calling custom / zero shot tools.
	//
	// Any of "json", "function_tag", "python_list".
	//
	// Deprecated: deprecated
	ToolPromptFormat AgentConfigToolPromptFormat `json:"tool_prompt_format"`
	Toolgroups       []AgentToolUnion            `json:"toolgroups"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instructions             respjson.Field
		Model                    respjson.Field
		ClientTools              respjson.Field
		EnableSessionPersistence respjson.Field
		InputShields             respjson.Field
		MaxInferIters            respjson.Field
		Name                     respjson.Field
		OutputShields            respjson.Field
		ResponseFormat           respjson.Field
		SamplingParams           respjson.Field
		ToolChoice               respjson.Field
		ToolConfig               respjson.Field
		ToolPromptFormat         respjson.Field
		Toolgroups               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConfig) RawJSON() string { return r.JSON.raw }
func (r *AgentConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentConfig to a AgentConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentConfigParam.Overrides()
func (r AgentConfig) ToParam() AgentConfigParam {
	return param.Override[AgentConfigParam](json.RawMessage(r.RawJSON()))
}

// Whether tool use is required or automatic. This is a hint to the model which may
// not be followed. It depends on the Instruction Following capabilities of the
// model.
type AgentConfigToolChoice string

const (
	AgentConfigToolChoiceAuto     AgentConfigToolChoice = "auto"
	AgentConfigToolChoiceRequired AgentConfigToolChoice = "required"
	AgentConfigToolChoiceNone     AgentConfigToolChoice = "none"
)

// Prompt format for calling custom / zero shot tools.
type AgentConfigToolPromptFormat string

const (
	AgentConfigToolPromptFormatJson        AgentConfigToolPromptFormat = "json"
	AgentConfigToolPromptFormatFunctionTag AgentConfigToolPromptFormat = "function_tag"
	AgentConfigToolPromptFormatPythonList  AgentConfigToolPromptFormat = "python_list"
)

// Configuration for an agent.
//
// The properties Instructions, Model are required.
type AgentConfigParam struct {
	// The system instructions for the agent
	Instructions string `json:"instructions,required"`
	// The model identifier to use for the agent
	Model string `json:"model,required"`
	// Optional flag indicating whether session data has to be persisted
	EnableSessionPersistence param.Opt[bool]  `json:"enable_session_persistence,omitzero"`
	MaxInferIters            param.Opt[int64] `json:"max_infer_iters,omitzero"`
	// Optional name for the agent, used in telemetry and identification
	Name          param.Opt[string] `json:"name,omitzero"`
	ClientTools   []ToolDefParam    `json:"client_tools,omitzero"`
	InputShields  []string          `json:"input_shields,omitzero"`
	OutputShields []string          `json:"output_shields,omitzero"`
	// Optional response format configuration
	ResponseFormat ResponseFormatUnionParam `json:"response_format,omitzero"`
	// Sampling parameters.
	SamplingParams SamplingParams `json:"sampling_params,omitzero"`
	// Whether tool use is required or automatic. This is a hint to the model which may
	// not be followed. It depends on the Instruction Following capabilities of the
	// model.
	//
	// Any of "auto", "required", "none".
	//
	// Deprecated: deprecated
	ToolChoice AgentConfigToolChoice `json:"tool_choice,omitzero"`
	// Configuration for tool use.
	ToolConfig ToolConfigParam `json:"tool_config,omitzero"`
	// Prompt format for calling custom / zero shot tools.
	//
	// Any of "json", "function_tag", "python_list".
	//
	// Deprecated: deprecated
	ToolPromptFormat AgentConfigToolPromptFormat `json:"tool_prompt_format,omitzero"`
	Toolgroups       []AgentToolUnionParam       `json:"toolgroups,omitzero"`
	paramObj
}

func (r AgentConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A generic paginated response that follows a simple format.
type PaginatedResponse struct {
	// The list of items for the current page
	Data []map[string]PaginatedResponseDataUnion `json:"data,required"`
	// Whether there are more items available after this set
	HasMore bool `json:"has_more,required"`
	// The URL for accessing this list
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginatedResponse) RawJSON() string { return r.JSON.raw }
func (r *PaginatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaginatedResponseDataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type PaginatedResponseDataUnion struct {
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

func (u PaginatedResponseDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaginatedResponseDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaginatedResponseDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaginatedResponseDataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaginatedResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *PaginatedResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response returned when creating a new agent.
type AgentNewResponse struct {
	// Unique identifier for the created agent
	AgentID string `json:"agent_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An agent instance with configuration and metadata.
type AgentGetResponse struct {
	// Configuration settings for the agent
	AgentConfig AgentConfig `json:"agent_config,required"`
	// Unique identifier for the agent
	AgentID string `json:"agent_id,required"`
	// Timestamp when the agent was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentConfig respjson.Field
		AgentID     respjson.Field
		CreatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response returned when creating a new agent session.
type AgentNewSessionResponse struct {
	// Unique identifier for the created session
	SessionID string `json:"session_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentNewSessionResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentNewSessionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentNewParams struct {
	// The configuration for the agent.
	AgentConfig AgentConfigParam `json:"agent_config,omitzero,required"`
	paramObj
}

func (r AgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// The name of the session to create.
	SessionName string `json:"session_name,required"`
	paramObj
}

func (r AgentNewSessionParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewSessionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewSessionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetSessionsParams struct {
	// The number of sessions to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index to start the pagination from.
	StartIndex param.Opt[int64] `query:"start_index,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentGetSessionsParams]'s query parameters as `url.Values`.
func (r AgentGetSessionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

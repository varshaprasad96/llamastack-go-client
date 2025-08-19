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
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
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
func (r *AgentSessionTurnService) New(ctx context.Context, sessionID string, params AgentSessionTurnNewParams, opts ...option.RequestOption) (res *Turn, err error) {
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

// Retrieve an agent turn by its ID.
func (r *AgentSessionTurnService) Get(ctx context.Context, turnID string, query AgentSessionTurnGetParams, opts ...option.RequestOption) (res *Turn, err error) {
	opts = append(r.Options[:], opts...)
	if query.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	if turnID == "" {
		err = errors.New("missing required turn_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s/turn/%s", query.AgentID, query.SessionID, turnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resume an agent turn with executed tool call responses. When a Turn has the
// status `awaiting_input` due to pending input from client side tool calls, this
// endpoint can be used to submit the outputs from the tool calls once they are
// ready.
func (r *AgentSessionTurnService) Resume(ctx context.Context, turnID string, params AgentSessionTurnResumeParams, opts ...option.RequestOption) (res *Turn, err error) {
	opts = append(r.Options[:], opts...)
	if params.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if params.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	if turnID == "" {
		err = errors.New("missing required turn_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s/turn/%s/resume", params.AgentID, params.SessionID, turnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// AgentToolUnion contains all possible properties and values from [string],
// [AgentToolAgentToolGroupWithArgs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type AgentToolUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [AgentToolAgentToolGroupWithArgs].
	Args map[string]AgentToolAgentToolGroupWithArgsArgUnion `json:"args"`
	// This field is from variant [AgentToolAgentToolGroupWithArgs].
	Name string `json:"name"`
	JSON struct {
		OfString respjson.Field
		Args     respjson.Field
		Name     respjson.Field
		raw      string
	} `json:"-"`
}

func (u AgentToolUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentToolUnion) AsAgentToolGroupWithArgs() (v AgentToolAgentToolGroupWithArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentToolUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentToolUnion to a AgentToolUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentToolUnionParam.Overrides()
func (r AgentToolUnion) ToParam() AgentToolUnionParam {
	return param.Override[AgentToolUnionParam](json.RawMessage(r.RawJSON()))
}

type AgentToolAgentToolGroupWithArgs struct {
	Args map[string]AgentToolAgentToolGroupWithArgsArgUnion `json:"args,required"`
	Name string                                             `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentToolAgentToolGroupWithArgs) RawJSON() string { return r.JSON.raw }
func (r *AgentToolAgentToolGroupWithArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentToolAgentToolGroupWithArgsArgUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type AgentToolAgentToolGroupWithArgsArgUnion struct {
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

func (u AgentToolAgentToolGroupWithArgsArgUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentToolAgentToolGroupWithArgsArgUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentToolAgentToolGroupWithArgsArgUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentToolAgentToolGroupWithArgsArgUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentToolAgentToolGroupWithArgsArgUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentToolAgentToolGroupWithArgsArgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func AgentToolParamOfAgentToolGroupWithArgs(args map[string]AgentToolAgentToolGroupWithArgsArgUnionParam, name string) AgentToolUnionParam {
	var variant AgentToolAgentToolGroupWithArgsParam
	variant.Args = args
	variant.Name = name
	return AgentToolUnionParam{OfAgentToolGroupWithArgs: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentToolUnionParam struct {
	OfString                 param.Opt[string]                     `json:",omitzero,inline"`
	OfAgentToolGroupWithArgs *AgentToolAgentToolGroupWithArgsParam `json:",omitzero,inline"`
	paramUnion
}

func (u AgentToolUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAgentToolGroupWithArgs)
}
func (u *AgentToolUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentToolUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAgentToolGroupWithArgs) {
		return u.OfAgentToolGroupWithArgs
	}
	return nil
}

// The properties Args, Name are required.
type AgentToolAgentToolGroupWithArgsParam struct {
	Args map[string]AgentToolAgentToolGroupWithArgsArgUnionParam `json:"args,omitzero,required"`
	Name string                                                  `json:"name,required"`
	paramObj
}

func (r AgentToolAgentToolGroupWithArgsParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentToolAgentToolGroupWithArgsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentToolAgentToolGroupWithArgsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentToolAgentToolGroupWithArgsArgUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u AgentToolAgentToolGroupWithArgsArgUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *AgentToolAgentToolGroupWithArgsArgUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentToolAgentToolGroupWithArgsArgUnionParam) asAny() any {
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

// An inference step in an agent turn.
type InferenceStep struct {
	// The response from the LLM.
	ModelResponse CompletionMessage `json:"model_response,required"`
	// The ID of the step.
	StepID string `json:"step_id,required"`
	// Type of the step in an agent turn.
	//
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType InferenceStepStepType `json:"step_type,required"`
	// The ID of the turn.
	TurnID string `json:"turn_id,required"`
	// The time the step completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// The time the step started.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelResponse respjson.Field
		StepID        respjson.Field
		StepType      respjson.Field
		TurnID        respjson.Field
		CompletedAt   respjson.Field
		StartedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceStep) RawJSON() string { return r.JSON.raw }
func (r *InferenceStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the step in an agent turn.
type InferenceStepStepType string

const (
	InferenceStepStepTypeInference       InferenceStepStepType = "inference"
	InferenceStepStepTypeToolExecution   InferenceStepStepType = "tool_execution"
	InferenceStepStepTypeShieldCall      InferenceStepStepType = "shield_call"
	InferenceStepStepTypeMemoryRetrieval InferenceStepStepType = "memory_retrieval"
)

// A memory retrieval step in an agent turn.
type MemoryRetrievalStep struct {
	// The context retrieved from the vector databases.
	InsertedContext InterleavedContentUnion `json:"inserted_context,required"`
	// The ID of the step.
	StepID string `json:"step_id,required"`
	// Type of the step in an agent turn.
	//
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType MemoryRetrievalStepStepType `json:"step_type,required"`
	// The ID of the turn.
	TurnID string `json:"turn_id,required"`
	// The IDs of the vector databases to retrieve context from.
	VectorDBIDs string `json:"vector_db_ids,required"`
	// The time the step completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// The time the step started.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InsertedContext respjson.Field
		StepID          respjson.Field
		StepType        respjson.Field
		TurnID          respjson.Field
		VectorDBIDs     respjson.Field
		CompletedAt     respjson.Field
		StartedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryRetrievalStep) RawJSON() string { return r.JSON.raw }
func (r *MemoryRetrievalStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the step in an agent turn.
type MemoryRetrievalStepStepType string

const (
	MemoryRetrievalStepStepTypeInference       MemoryRetrievalStepStepType = "inference"
	MemoryRetrievalStepStepTypeToolExecution   MemoryRetrievalStepStepType = "tool_execution"
	MemoryRetrievalStepStepTypeShieldCall      MemoryRetrievalStepStepType = "shield_call"
	MemoryRetrievalStepStepTypeMemoryRetrieval MemoryRetrievalStepStepType = "memory_retrieval"
)

// A shield call step in an agent turn.
type ShieldCallStep struct {
	// The ID of the step.
	StepID string `json:"step_id,required"`
	// Type of the step in an agent turn.
	//
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType ShieldCallStepStepType `json:"step_type,required"`
	// The ID of the turn.
	TurnID string `json:"turn_id,required"`
	// The time the step completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// The time the step started.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// The violation from the shield call.
	Violation SafetyViolation `json:"violation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StepID      respjson.Field
		StepType    respjson.Field
		TurnID      respjson.Field
		CompletedAt respjson.Field
		StartedAt   respjson.Field
		Violation   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShieldCallStep) RawJSON() string { return r.JSON.raw }
func (r *ShieldCallStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the step in an agent turn.
type ShieldCallStepStepType string

const (
	ShieldCallStepStepTypeInference       ShieldCallStepStepType = "inference"
	ShieldCallStepStepTypeToolExecution   ShieldCallStepStepType = "tool_execution"
	ShieldCallStepStepTypeShieldCall      ShieldCallStepStepType = "shield_call"
	ShieldCallStepStepTypeMemoryRetrieval ShieldCallStepStepType = "memory_retrieval"
)

// A tool execution step in an agent turn.
type ToolExecutionStep struct {
	// The ID of the step.
	StepID string `json:"step_id,required"`
	// Type of the step in an agent turn.
	//
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType ToolExecutionStepStepType `json:"step_type,required"`
	// The tool calls to execute.
	ToolCalls []ToolCall `json:"tool_calls,required"`
	// The tool responses from the tool calls.
	ToolResponses []ToolResponse `json:"tool_responses,required"`
	// The ID of the turn.
	TurnID string `json:"turn_id,required"`
	// The time the step completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// The time the step started.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StepID        respjson.Field
		StepType      respjson.Field
		ToolCalls     respjson.Field
		ToolResponses respjson.Field
		TurnID        respjson.Field
		CompletedAt   respjson.Field
		StartedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolExecutionStep) RawJSON() string { return r.JSON.raw }
func (r *ToolExecutionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the step in an agent turn.
type ToolExecutionStepStepType string

const (
	ToolExecutionStepStepTypeInference       ToolExecutionStepStepType = "inference"
	ToolExecutionStepStepTypeToolExecution   ToolExecutionStepStepType = "tool_execution"
	ToolExecutionStepStepTypeShieldCall      ToolExecutionStepStepType = "shield_call"
	ToolExecutionStepStepTypeMemoryRetrieval ToolExecutionStepStepType = "memory_retrieval"
)

type ToolResponse struct {
	CallID string `json:"call_id,required"`
	// A image content item
	Content  InterleavedContentUnion              `json:"content,required"`
	ToolName ToolResponseToolName                 `json:"tool_name,required"`
	Metadata map[string]ToolResponseMetadataUnion `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID      respjson.Field
		Content     respjson.Field
		ToolName    respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolResponse) RawJSON() string { return r.JSON.raw }
func (r *ToolResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolResponse to a ToolResponseParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolResponseParam.Overrides()
func (r ToolResponse) ToParam() ToolResponseParam {
	return param.Override[ToolResponseParam](json.RawMessage(r.RawJSON()))
}

type ToolResponseToolName string

const (
	ToolResponseToolNameBraveSearch     ToolResponseToolName = "brave_search"
	ToolResponseToolNameWolframAlpha    ToolResponseToolName = "wolfram_alpha"
	ToolResponseToolNamePhotogen        ToolResponseToolName = "photogen"
	ToolResponseToolNameCodeInterpreter ToolResponseToolName = "code_interpreter"
)

// ToolResponseMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolResponseMetadataUnion struct {
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

func (u ToolResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolResponseMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CallID, Content, ToolName are required.
type ToolResponseParam struct {
	CallID string `json:"call_id,required"`
	// A image content item
	Content  InterleavedContentUnionParam              `json:"content,omitzero,required"`
	ToolName ToolResponseToolName                      `json:"tool_name,omitzero,required"`
	Metadata map[string]ToolResponseMetadataUnionParam `json:"metadata,omitzero"`
	paramObj
}

func (r ToolResponseParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolResponseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolResponseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolResponseMetadataUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ToolResponseMetadataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ToolResponseMetadataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolResponseMetadataUnionParam) asAny() any {
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

// A message representing the result of a tool invocation.
type ToolResponseMessage struct {
	// Unique identifier for the tool call this response is for
	CallID string `json:"call_id,required"`
	// The response content from the tool
	Content InterleavedContentUnion `json:"content,required"`
	// Must be "tool" to identify this as a tool response
	//
	// Any of "system", "user", "assistant", "tool".
	Role ToolResponseMessageRole `json:"role,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID      respjson.Field
		Content     respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ToolResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolResponseMessage to a ToolResponseMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolResponseMessageParam.Overrides()
func (r ToolResponseMessage) ToParam() ToolResponseMessageParam {
	return param.Override[ToolResponseMessageParam](json.RawMessage(r.RawJSON()))
}

// Must be "tool" to identify this as a tool response
type ToolResponseMessageRole string

const (
	ToolResponseMessageRoleSystem    ToolResponseMessageRole = "system"
	ToolResponseMessageRoleUser      ToolResponseMessageRole = "user"
	ToolResponseMessageRoleAssistant ToolResponseMessageRole = "assistant"
	ToolResponseMessageRoleTool      ToolResponseMessageRole = "tool"
)

// A message representing the result of a tool invocation.
//
// The properties CallID, Content, Role are required.
type ToolResponseMessageParam struct {
	// Unique identifier for the tool call this response is for
	CallID string `json:"call_id,required"`
	// The response content from the tool
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Must be "tool" to identify this as a tool response
	//
	// Any of "system", "user", "assistant", "tool".
	Role ToolResponseMessageRole `json:"role,omitzero,required"`
	paramObj
}

func (r ToolResponseMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolResponseMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolResponseMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single turn in an interaction with an Agentic System.
type Turn struct {
	InputMessages []TurnInputMessageUnion `json:"input_messages,required"`
	// A message containing the model's (assistant) response in a chat conversation.
	OutputMessage     CompletionMessage      `json:"output_message,required"`
	SessionID         string                 `json:"session_id,required"`
	StartedAt         time.Time              `json:"started_at,required" format:"date-time"`
	Steps             []TurnStepUnion        `json:"steps,required"`
	TurnID            string                 `json:"turn_id,required"`
	CompletedAt       time.Time              `json:"completed_at" format:"date-time"`
	OutputAttachments []TurnOutputAttachment `json:"output_attachments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputMessages     respjson.Field
		OutputMessage     respjson.Field
		SessionID         respjson.Field
		StartedAt         respjson.Field
		Steps             respjson.Field
		TurnID            respjson.Field
		CompletedAt       respjson.Field
		OutputAttachments respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Turn) RawJSON() string { return r.JSON.raw }
func (r *Turn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnInputMessageUnion contains all possible properties and values from
// [UserMessage], [ToolResponseMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TurnInputMessageUnion struct {
	// This field is from variant [UserMessage].
	Content InterleavedContentUnion `json:"content"`
	Role    string                  `json:"role"`
	// This field is from variant [UserMessage].
	Context InterleavedContentUnion `json:"context"`
	// This field is from variant [ToolResponseMessage].
	CallID string `json:"call_id"`
	JSON   struct {
		Content respjson.Field
		Role    respjson.Field
		Context respjson.Field
		CallID  respjson.Field
		raw     string
	} `json:"-"`
}

func (u TurnInputMessageUnion) AsUserMessage() (v UserMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnInputMessageUnion) AsToolResponseMessage() (v ToolResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnInputMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *TurnInputMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnStepUnion contains all possible properties and values from [InferenceStep],
// [ToolExecutionStep], [ShieldCallStep], [MemoryRetrievalStep].
//
// Use the [TurnStepUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TurnStepUnion struct {
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

func (u TurnStepUnion) AsInferenceStep() (v InferenceStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnStepUnion) AsToolExecutionStep() (v ToolExecutionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnStepUnion) AsShieldCallStep() (v ShieldCallStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnStepUnion) AsMemoryRetrievalStep() (v MemoryRetrievalStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnStepUnion) RawJSON() string { return u.JSON.raw }

func (r *TurnStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An attachment to an agent turn.
type TurnOutputAttachment struct {
	// The content of the attachment.
	Content TurnOutputAttachmentContentUnion `json:"content,required"`
	// The MIME type of the attachment.
	MimeType string `json:"mime_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		MimeType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnOutputAttachment) RawJSON() string { return r.JSON.raw }
func (r *TurnOutputAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnOutputAttachmentContentUnion contains all possible properties and values
// from [string], [TurnOutputAttachmentContentImageContentItem],
// [TurnOutputAttachmentContentTextContentItem], [[]InterleavedContentItemUnion],
// [URL].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInterleavedContentItemArray]
type TurnOutputAttachmentContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]InterleavedContentItemUnion]
	// instead of an object.
	OfInterleavedContentItemArray []InterleavedContentItemUnion `json:",inline"`
	// This field is from variant [TurnOutputAttachmentContentImageContentItem].
	Image TurnOutputAttachmentContentImageContentItemImage `json:"image"`
	Type  string                                           `json:"type"`
	// This field is from variant [TurnOutputAttachmentContentTextContentItem].
	Text string `json:"text"`
	// This field is from variant [URL].
	Uri  string `json:"uri"`
	JSON struct {
		OfString                      respjson.Field
		OfInterleavedContentItemArray respjson.Field
		Image                         respjson.Field
		Type                          respjson.Field
		Text                          respjson.Field
		Uri                           respjson.Field
		raw                           string
	} `json:"-"`
}

func (u TurnOutputAttachmentContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnOutputAttachmentContentUnion) AsImageContentItem() (v TurnOutputAttachmentContentImageContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnOutputAttachmentContentUnion) AsTextContentItem() (v TurnOutputAttachmentContentTextContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnOutputAttachmentContentUnion) AsInterleavedContentItemArray() (v []InterleavedContentItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnOutputAttachmentContentUnion) AsURL() (v URL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnOutputAttachmentContentUnion) RawJSON() string { return u.JSON.raw }

func (r *TurnOutputAttachmentContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A image content item
type TurnOutputAttachmentContentImageContentItem struct {
	// Image as a base64 encoded string or an URL
	Image TurnOutputAttachmentContentImageContentItemImage `json:"image,required"`
	// Discriminator type of the content item. Always "image"
	//
	// Any of "text", "image", "tool_call".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnOutputAttachmentContentImageContentItem) RawJSON() string { return r.JSON.raw }
func (r *TurnOutputAttachmentContentImageContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type TurnOutputAttachmentContentImageContentItemImage struct {
	// base64 encoded image data as string
	Data string `json:"data"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL URL `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnOutputAttachmentContentImageContentItemImage) RawJSON() string { return r.JSON.raw }
func (r *TurnOutputAttachmentContentImageContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type TurnOutputAttachmentContentTextContentItem struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// Any of "text", "image", "tool_call".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnOutputAttachmentContentTextContentItem) RawJSON() string { return r.JSON.raw }
func (r *TurnOutputAttachmentContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in a chat conversation.
type UserMessage struct {
	// The content of the message, which can include text and other media
	Content InterleavedContentUnion `json:"content,required"`
	// Must be "user" to identify this as a user message
	//
	// Any of "system", "user", "assistant", "tool".
	Role UserMessageRole `json:"role,required"`
	// (Optional) This field is used internally by Llama Stack to pass RAG context.
	// This field may be removed in the API in the future.
	Context InterleavedContentUnion `json:"context"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Context     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserMessage) RawJSON() string { return r.JSON.raw }
func (r *UserMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UserMessage to a UserMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserMessageParam.Overrides()
func (r UserMessage) ToParam() UserMessageParam {
	return param.Override[UserMessageParam](json.RawMessage(r.RawJSON()))
}

// Must be "user" to identify this as a user message
type UserMessageRole string

const (
	UserMessageRoleSystem    UserMessageRole = "system"
	UserMessageRoleUser      UserMessageRole = "user"
	UserMessageRoleAssistant UserMessageRole = "assistant"
	UserMessageRoleTool      UserMessageRole = "tool"
)

// A message from the user in a chat conversation.
//
// The properties Content, Role are required.
type UserMessageParam struct {
	// The content of the message, which can include text and other media
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Must be "user" to identify this as a user message
	//
	// Any of "system", "user", "assistant", "tool".
	Role UserMessageRole `json:"role,omitzero,required"`
	// (Optional) This field is used internally by Llama Stack to pass RAG context.
	// This field may be removed in the API in the future.
	Context InterleavedContentUnionParam `json:"context,omitzero"`
	paramObj
}

func (r UserMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow UserMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentSessionTurnNewParams struct {
	AgentID string `path:"agent_id,required" json:"-"`
	// List of messages to start the turn with.
	Messages []AgentSessionTurnNewParamsMessageUnion `json:"messages,omitzero,required"`
	// (Optional) If True, generate an SSE event stream of the response. Defaults to
	// False.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// (Optional) List of documents to create the turn with.
	Documents []AgentSessionTurnNewParamsDocument `json:"documents,omitzero"`
	// (Optional) The tool configuration to create the turn with, will be used to
	// override the agent's tool_config.
	ToolConfig ToolConfigParam `json:"tool_config,omitzero"`
	// (Optional) List of toolgroups to create the turn with, will be used in addition
	// to the agent's config toolgroups for the request.
	Toolgroups []AgentToolUnionParam `json:"toolgroups,omitzero"`
	paramObj
}

func (r AgentSessionTurnNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentSessionTurnNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSessionTurnNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentSessionTurnNewParamsMessageUnion struct {
	OfUserMessage         *UserMessageParam         `json:",omitzero,inline"`
	OfToolResponseMessage *ToolResponseMessageParam `json:",omitzero,inline"`
	paramUnion
}

func (u AgentSessionTurnNewParamsMessageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUserMessage, u.OfToolResponseMessage)
}
func (u *AgentSessionTurnNewParamsMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentSessionTurnNewParamsMessageUnion) asAny() any {
	if !param.IsOmitted(u.OfUserMessage) {
		return u.OfUserMessage
	} else if !param.IsOmitted(u.OfToolResponseMessage) {
		return u.OfToolResponseMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentSessionTurnNewParamsMessageUnion) GetContext() *InterleavedContentUnionParam {
	if vt := u.OfUserMessage; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentSessionTurnNewParamsMessageUnion) GetCallID() *string {
	if vt := u.OfToolResponseMessage; vt != nil {
		return &vt.CallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentSessionTurnNewParamsMessageUnion) GetRole() *string {
	if vt := u.OfUserMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfToolResponseMessage; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's Content property, if present.
func (u AgentSessionTurnNewParamsMessageUnion) GetContent() *InterleavedContentUnionParam {
	if vt := u.OfUserMessage; vt != nil {
		return &vt.Content
	} else if vt := u.OfToolResponseMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// A document to be used by an agent.
//
// The properties Content, MimeType are required.
type AgentSessionTurnNewParamsDocument struct {
	// The content of the document.
	Content AgentSessionTurnNewParamsDocumentContentUnion `json:"content,omitzero,required"`
	// The MIME type of the document.
	MimeType string `json:"mime_type,required"`
	paramObj
}

func (r AgentSessionTurnNewParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow AgentSessionTurnNewParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSessionTurnNewParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentSessionTurnNewParamsDocumentContentUnion struct {
	OfString                      param.Opt[string]                                         `json:",omitzero,inline"`
	OfImageContentItem            *AgentSessionTurnNewParamsDocumentContentImageContentItem `json:",omitzero,inline"`
	OfTextContentItem             *AgentSessionTurnNewParamsDocumentContentTextContentItem  `json:",omitzero,inline"`
	OfInterleavedContentItemArray []InterleavedContentItemUnionParam                        `json:",omitzero,inline"`
	OfURL                         *URLParam                                                 `json:",omitzero,inline"`
	paramUnion
}

func (u AgentSessionTurnNewParamsDocumentContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfImageContentItem,
		u.OfTextContentItem,
		u.OfInterleavedContentItemArray,
		u.OfURL)
}
func (u *AgentSessionTurnNewParamsDocumentContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentSessionTurnNewParamsDocumentContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfImageContentItem) {
		return u.OfImageContentItem
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	} else if !param.IsOmitted(u.OfInterleavedContentItemArray) {
		return &u.OfInterleavedContentItemArray
	} else if !param.IsOmitted(u.OfURL) {
		return u.OfURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentSessionTurnNewParamsDocumentContentUnion) GetImage() *AgentSessionTurnNewParamsDocumentContentImageContentItemImage {
	if vt := u.OfImageContentItem; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentSessionTurnNewParamsDocumentContentUnion) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentSessionTurnNewParamsDocumentContentUnion) GetUri() *string {
	if vt := u.OfURL; vt != nil {
		return &vt.Uri
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentSessionTurnNewParamsDocumentContentUnion) GetType() *string {
	if vt := u.OfImageContentItem; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A image content item
//
// The properties Image, Type are required.
type AgentSessionTurnNewParamsDocumentContentImageContentItem struct {
	// Image as a base64 encoded string or an URL
	Image AgentSessionTurnNewParamsDocumentContentImageContentItemImage `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// Any of "text", "image", "tool_call".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r AgentSessionTurnNewParamsDocumentContentImageContentItem) MarshalJSON() (data []byte, err error) {
	type shadow AgentSessionTurnNewParamsDocumentContentImageContentItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSessionTurnNewParamsDocumentContentImageContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentSessionTurnNewParamsDocumentContentImageContentItem](
		"type", "text", "image", "tool_call",
	)
}

// Image as a base64 encoded string or an URL
type AgentSessionTurnNewParamsDocumentContentImageContentItemImage struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL URLParam `json:"url,omitzero"`
	paramObj
}

func (r AgentSessionTurnNewParamsDocumentContentImageContentItemImage) MarshalJSON() (data []byte, err error) {
	type shadow AgentSessionTurnNewParamsDocumentContentImageContentItemImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSessionTurnNewParamsDocumentContentImageContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type AgentSessionTurnNewParamsDocumentContentTextContentItem struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// Any of "text", "image", "tool_call".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r AgentSessionTurnNewParamsDocumentContentTextContentItem) MarshalJSON() (data []byte, err error) {
	type shadow AgentSessionTurnNewParamsDocumentContentTextContentItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSessionTurnNewParamsDocumentContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentSessionTurnNewParamsDocumentContentTextContentItem](
		"type", "text", "image", "tool_call",
	)
}

type AgentSessionTurnGetParams struct {
	AgentID   string `path:"agent_id,required" json:"-"`
	SessionID string `path:"session_id,required" json:"-"`
	paramObj
}

type AgentSessionTurnResumeParams struct {
	AgentID   string `path:"agent_id,required" json:"-"`
	SessionID string `path:"session_id,required" json:"-"`
	// The tool call responses to resume the turn with.
	ToolResponses []ToolResponseParam `json:"tool_responses,omitzero,required"`
	// Whether to stream the response.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	paramObj
}

func (r AgentSessionTurnResumeParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentSessionTurnResumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSessionTurnResumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

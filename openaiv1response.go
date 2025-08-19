// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/varshaprasad96/llamastack-go-client/internal/apijson"
	"github.com/varshaprasad96/llamastack-go-client/internal/apiquery"
	"github.com/varshaprasad96/llamastack-go-client/internal/requestconfig"
	"github.com/varshaprasad96/llamastack-go-client/option"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/packages/respjson"
	"github.com/varshaprasad96/llamastack-go-client/shared/constant"
)

// OpenAIV1ResponseService contains methods and other services that help with
// interacting with the llamastack-go-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIV1ResponseService] method instead.
type OpenAIV1ResponseService struct {
	Options []option.RequestOption
}

// NewOpenAIV1ResponseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOpenAIV1ResponseService(opts ...option.RequestOption) (r OpenAIV1ResponseService) {
	r = OpenAIV1ResponseService{}
	r.Options = opts
	return
}

// Create a new OpenAI response.
func (r *OpenAIV1ResponseService) New(ctx context.Context, body OpenAIV1ResponseNewParams, opts ...option.RequestOption) (res *ResponseObject, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an OpenAI response by its ID.
func (r *OpenAIV1ResponseService) Get(ctx context.Context, responseID string, opts ...option.RequestOption) (res *ResponseObject, err error) {
	opts = append(r.Options[:], opts...)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all OpenAI responses.
func (r *OpenAIV1ResponseService) List(ctx context.Context, query OpenAIV1ResponseListParams, opts ...option.RequestOption) (res *OpenAiv1ResponseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/openai/v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an OpenAI response by its ID.
func (r *OpenAIV1ResponseService) Delete(ctx context.Context, responseID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// List input items for a given OpenAI response.
func (r *OpenAIV1ResponseService) GetInputItems(ctx context.Context, responseID string, query OpenAIV1ResponseGetInputItemsParams, opts ...option.RequestOption) (res *OpenAiv1ResponseGetInputItemsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/openai/v1/responses/%s/input_items", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

// OutputUnion contains all possible properties and values from [ResponseMessage],
// [OutputMessageWebSearchToolCall], [OutputMessageFileSearchToolCall],
// [OutputMessageFunctionToolCall], [OutputMcpCall], [OutputMcpListTools].
//
// Use the [OutputUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OutputUnion struct {
	// This field is from variant [ResponseMessage].
	Content ResponseMessageContentUnion `json:"content"`
	// This field is from variant [ResponseMessage].
	Role ResponseMessageRole `json:"role"`
	// Any of nil, "web_search_call", nil, nil, "mcp_call", "mcp_list_tools".
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
	// This field is from variant [OutputMessageFileSearchToolCall].
	Queries []string `json:"queries"`
	// This field is from variant [OutputMessageFileSearchToolCall].
	Results   []map[string]OutputMessageFileSearchToolCallResultUnion `json:"results"`
	Arguments string                                                  `json:"arguments"`
	// This field is from variant [OutputMessageFunctionToolCall].
	CallID      string `json:"call_id"`
	Name        string `json:"name"`
	ServerLabel string `json:"server_label"`
	// This field is from variant [OutputMcpCall].
	Error string `json:"error"`
	// This field is from variant [OutputMcpCall].
	Output string `json:"output"`
	// This field is from variant [OutputMcpListTools].
	Tools []OutputMcpListToolsTool `json:"tools"`
	JSON  struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		Queries     respjson.Field
		Results     respjson.Field
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		Tools       respjson.Field
		raw         string
	} `json:"-"`
}

func (u OutputUnion) AsOpenAIResponseMessage() (v ResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputUnion) AsWebSearchCall() (v OutputMessageWebSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputUnion) AsOpenAIResponseOutputMessageFileSearchToolCall() (v OutputMessageFileSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputUnion) AsOpenAIResponseOutputMessageFunctionToolCall() (v OutputMessageFunctionToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputUnion) AsMcpCall() (v OutputMcpCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputUnion) AsMcpListTools() (v OutputMcpListTools) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OutputUnion) RawJSON() string { return u.JSON.raw }

func (r *OutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputMcpCall struct {
	ID          string           `json:"id,required"`
	Arguments   string           `json:"arguments,required"`
	Name        string           `json:"name,required"`
	ServerLabel string           `json:"server_label,required"`
	Type        constant.McpCall `json:"type,required"`
	Error       string           `json:"error"`
	Output      string           `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Arguments   respjson.Field
		Name        respjson.Field
		ServerLabel respjson.Field
		Type        respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputMcpCall) RawJSON() string { return r.JSON.raw }
func (r *OutputMcpCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputMcpListTools struct {
	ID          string                   `json:"id,required"`
	ServerLabel string                   `json:"server_label,required"`
	Tools       []OutputMcpListToolsTool `json:"tools,required"`
	Type        constant.McpListTools    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ServerLabel respjson.Field
		Tools       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputMcpListTools) RawJSON() string { return r.JSON.raw }
func (r *OutputMcpListTools) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputMcpListToolsTool struct {
	InputSchema map[string]OutputMcpListToolsToolInputSchemaUnion `json:"input_schema,required"`
	Name        string                                            `json:"name,required"`
	Description string                                            `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputSchema respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputMcpListToolsTool) RawJSON() string { return r.JSON.raw }
func (r *OutputMcpListToolsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputMcpListToolsToolInputSchemaUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type OutputMcpListToolsToolInputSchemaUnion struct {
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

func (u OutputMcpListToolsToolInputSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputMcpListToolsToolInputSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputMcpListToolsToolInputSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputMcpListToolsToolInputSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OutputMcpListToolsToolInputSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *OutputMcpListToolsToolInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputRole string

const (
	OutputRoleSystem    OutputRole = "system"
	OutputRoleDeveloper OutputRole = "developer"
	OutputRoleUser      OutputRole = "user"
	OutputRoleAssistant OutputRole = "assistant"
)

type OutputMessageFileSearchToolCall struct {
	ID      string                                                  `json:"id,required"`
	Queries []string                                                `json:"queries,required"`
	Status  string                                                  `json:"status,required"`
	Type    string                                                  `json:"type,required"`
	Results []map[string]OutputMessageFileSearchToolCallResultUnion `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Queries     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputMessageFileSearchToolCall) RawJSON() string { return r.JSON.raw }
func (r *OutputMessageFileSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OutputMessageFileSearchToolCall to a
// OutputMessageFileSearchToolCallParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OutputMessageFileSearchToolCallParam.Overrides()
func (r OutputMessageFileSearchToolCall) ToParam() OutputMessageFileSearchToolCallParam {
	return param.Override[OutputMessageFileSearchToolCallParam](json.RawMessage(r.RawJSON()))
}

// OutputMessageFileSearchToolCallResultUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type OutputMessageFileSearchToolCallResultUnion struct {
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

func (u OutputMessageFileSearchToolCallResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputMessageFileSearchToolCallResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputMessageFileSearchToolCallResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputMessageFileSearchToolCallResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OutputMessageFileSearchToolCallResultUnion) RawJSON() string { return u.JSON.raw }

func (r *OutputMessageFileSearchToolCallResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Queries, Status, Type are required.
type OutputMessageFileSearchToolCallParam struct {
	ID      string                                                       `json:"id,required"`
	Queries []string                                                     `json:"queries,omitzero,required"`
	Status  string                                                       `json:"status,required"`
	Type    string                                                       `json:"type,required"`
	Results []map[string]OutputMessageFileSearchToolCallResultUnionParam `json:"results,omitzero"`
	paramObj
}

func (r OutputMessageFileSearchToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow OutputMessageFileSearchToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutputMessageFileSearchToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OutputMessageFileSearchToolCallResultUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OutputMessageFileSearchToolCallResultUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OutputMessageFileSearchToolCallResultUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OutputMessageFileSearchToolCallResultUnionParam) asAny() any {
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

type OutputMessageFunctionToolCall struct {
	Arguments string `json:"arguments,required"`
	CallID    string `json:"call_id,required"`
	Name      string `json:"name,required"`
	Type      string `json:"type,required"`
	ID        string `json:"id"`
	Status    string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputMessageFunctionToolCall) RawJSON() string { return r.JSON.raw }
func (r *OutputMessageFunctionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OutputMessageFunctionToolCall to a
// OutputMessageFunctionToolCallParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OutputMessageFunctionToolCallParam.Overrides()
func (r OutputMessageFunctionToolCall) ToParam() OutputMessageFunctionToolCallParam {
	return param.Override[OutputMessageFunctionToolCallParam](json.RawMessage(r.RawJSON()))
}

// The properties Arguments, CallID, Name, Type are required.
type OutputMessageFunctionToolCallParam struct {
	Arguments string            `json:"arguments,required"`
	CallID    string            `json:"call_id,required"`
	Name      string            `json:"name,required"`
	Type      string            `json:"type,required"`
	ID        param.Opt[string] `json:"id,omitzero"`
	Status    param.Opt[string] `json:"status,omitzero"`
	paramObj
}

func (r OutputMessageFunctionToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow OutputMessageFunctionToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutputMessageFunctionToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputMessageWebSearchToolCall struct {
	ID     string                 `json:"id,required"`
	Status string                 `json:"status,required"`
	Type   constant.WebSearchCall `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputMessageWebSearchToolCall) RawJSON() string { return r.JSON.raw }
func (r *OutputMessageWebSearchToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OutputMessageWebSearchToolCall to a
// OutputMessageWebSearchToolCallParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OutputMessageWebSearchToolCallParam.Overrides()
func (r OutputMessageWebSearchToolCall) ToParam() OutputMessageWebSearchToolCallParam {
	return param.Override[OutputMessageWebSearchToolCallParam](json.RawMessage(r.RawJSON()))
}

// The properties ID, Status, Type are required.
type OutputMessageWebSearchToolCallParam struct {
	ID     string `json:"id,required"`
	Status string `json:"status,required"`
	// This field can be elided, and will marshal its zero value as "web_search_call".
	Type constant.WebSearchCall `json:"type,required"`
	paramObj
}

func (r OutputMessageWebSearchToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow OutputMessageWebSearchToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutputMessageWebSearchToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseError struct {
	Code    string `json:"code,required"`
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseError) RawJSON() string { return r.JSON.raw }
func (r *ResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseInputUnion contains all possible properties and values from
// [OutputMessageWebSearchToolCall], [OutputMessageFileSearchToolCall],
// [OutputMessageFunctionToolCall],
// [ResponseInputOpenAIResponseInputFunctionToolCallOutput], [ResponseMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseInputUnion struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	// This field is from variant [OutputMessageFileSearchToolCall].
	Queries []string `json:"queries"`
	// This field is from variant [OutputMessageFileSearchToolCall].
	Results []map[string]OutputMessageFileSearchToolCallResultUnion `json:"results"`
	// This field is from variant [OutputMessageFunctionToolCall].
	Arguments string `json:"arguments"`
	CallID    string `json:"call_id"`
	// This field is from variant [OutputMessageFunctionToolCall].
	Name string `json:"name"`
	// This field is from variant
	// [ResponseInputOpenAIResponseInputFunctionToolCallOutput].
	Output string `json:"output"`
	// This field is from variant [ResponseMessage].
	Content ResponseMessageContentUnion `json:"content"`
	// This field is from variant [ResponseMessage].
	Role ResponseMessageRole `json:"role"`
	JSON struct {
		ID        respjson.Field
		Status    respjson.Field
		Type      respjson.Field
		Queries   respjson.Field
		Results   respjson.Field
		Arguments respjson.Field
		CallID    respjson.Field
		Name      respjson.Field
		Output    respjson.Field
		Content   respjson.Field
		Role      respjson.Field
		raw       string
	} `json:"-"`
}

func (u ResponseInputUnion) AsOpenAIResponseOutputMessageWebSearchToolCall() (v OutputMessageWebSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputUnion) AsOpenAIResponseOutputMessageFileSearchToolCall() (v OutputMessageFileSearchToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputUnion) AsOpenAIResponseOutputMessageFunctionToolCall() (v OutputMessageFunctionToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputUnion) AsOpenAIResponseInputFunctionToolCallOutput() (v ResponseInputOpenAIResponseInputFunctionToolCallOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseInputUnion) AsOpenAIResponseMessage() (v ResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseInputUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ResponseInputUnion to a ResponseInputUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ResponseInputUnionParam.Overrides()
func (r ResponseInputUnion) ToParam() ResponseInputUnionParam {
	return param.Override[ResponseInputUnionParam](json.RawMessage(r.RawJSON()))
}

// This represents the output of a function call that gets passed back to the
// model.
type ResponseInputOpenAIResponseInputFunctionToolCallOutput struct {
	CallID string                      `json:"call_id,required"`
	Output string                      `json:"output,required"`
	Type   constant.FunctionCallOutput `json:"type,required"`
	ID     string                      `json:"id"`
	Status string                      `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID      respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInputOpenAIResponseInputFunctionToolCallOutput) RawJSON() string { return r.JSON.raw }
func (r *ResponseInputOpenAIResponseInputFunctionToolCallOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseInputRole string

const (
	ResponseInputRoleSystem    ResponseInputRole = "system"
	ResponseInputRoleDeveloper ResponseInputRole = "developer"
	ResponseInputRoleUser      ResponseInputRole = "user"
	ResponseInputRoleAssistant ResponseInputRole = "assistant"
)

func ResponseInputParamOfOpenAIResponseOutputMessageWebSearchToolCall(id string, status string) ResponseInputUnionParam {
	var variant OutputMessageWebSearchToolCallParam
	variant.ID = id
	variant.Status = status
	return ResponseInputUnionParam{OfOpenAIResponseOutputMessageWebSearchToolCall: &variant}
}

func ResponseInputParamOfOpenAIResponseInputFunctionToolCallOutput(callID string, output string) ResponseInputUnionParam {
	var variant ResponseInputOpenAIResponseInputFunctionToolCallOutputParam
	variant.CallID = callID
	variant.Output = output
	return ResponseInputUnionParam{OfOpenAIResponseInputFunctionToolCallOutput: &variant}
}

func ResponseInputParamOfOpenAIResponseMessage[
	T string | []ResponseMessageContentArrayItemUnionParam | []ResponseMessageContentArrayItemParam,
](content T, role ResponseMessageRole, type_ string) ResponseInputUnionParam {
	var variant ResponseMessageParam
	switch v := any(content).(type) {
	case string:
		variant.Content.OfString = param.NewOpt(v)
	case []ResponseMessageContentArrayItemUnionParam:
		variant.Content.OfResponseMessageContentArray = v
	case []ResponseMessageContentArrayItemParam:
		variant.Content.OfVariant2 = v
	}
	variant.Role = role
	variant.Type = type_
	return ResponseInputUnionParam{OfOpenAIResponseMessage: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseInputUnionParam struct {
	OfOpenAIResponseOutputMessageWebSearchToolCall  *OutputMessageWebSearchToolCallParam                         `json:",omitzero,inline"`
	OfOpenAIResponseOutputMessageFileSearchToolCall *OutputMessageFileSearchToolCallParam                        `json:",omitzero,inline"`
	OfOpenAIResponseOutputMessageFunctionToolCall   *OutputMessageFunctionToolCallParam                          `json:",omitzero,inline"`
	OfOpenAIResponseInputFunctionToolCallOutput     *ResponseInputOpenAIResponseInputFunctionToolCallOutputParam `json:",omitzero,inline"`
	OfOpenAIResponseMessage                         *ResponseMessageParam                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseInputUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIResponseOutputMessageWebSearchToolCall,
		u.OfOpenAIResponseOutputMessageFileSearchToolCall,
		u.OfOpenAIResponseOutputMessageFunctionToolCall,
		u.OfOpenAIResponseInputFunctionToolCallOutput,
		u.OfOpenAIResponseMessage)
}
func (u *ResponseInputUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseInputUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOpenAIResponseOutputMessageWebSearchToolCall) {
		return u.OfOpenAIResponseOutputMessageWebSearchToolCall
	} else if !param.IsOmitted(u.OfOpenAIResponseOutputMessageFileSearchToolCall) {
		return u.OfOpenAIResponseOutputMessageFileSearchToolCall
	} else if !param.IsOmitted(u.OfOpenAIResponseOutputMessageFunctionToolCall) {
		return u.OfOpenAIResponseOutputMessageFunctionToolCall
	} else if !param.IsOmitted(u.OfOpenAIResponseInputFunctionToolCallOutput) {
		return u.OfOpenAIResponseInputFunctionToolCallOutput
	} else if !param.IsOmitted(u.OfOpenAIResponseMessage) {
		return u.OfOpenAIResponseMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetQueries() []string {
	if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return vt.Queries
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetResults() []map[string]OutputMessageFileSearchToolCallResultUnionParam {
	if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return vt.Results
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetArguments() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return &vt.Arguments
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetName() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetOutput() *string {
	if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil {
		return &vt.Output
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetContent() *ResponseMessageContentUnionParam {
	if vt := u.OfOpenAIResponseMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetRole() *string {
	if vt := u.OfOpenAIResponseMessage; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetID() *string {
	if vt := u.OfOpenAIResponseOutputMessageWebSearchToolCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfOpenAIResponseMessage; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetStatus() *string {
	if vt := u.OfOpenAIResponseOutputMessageWebSearchToolCall; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	} else if vt := u.OfOpenAIResponseMessage; vt != nil && vt.Status.Valid() {
		return &vt.Status.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetType() *string {
	if vt := u.OfOpenAIResponseOutputMessageWebSearchToolCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseOutputMessageFileSearchToolCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIResponseMessage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseInputUnionParam) GetCallID() *string {
	if vt := u.OfOpenAIResponseOutputMessageFunctionToolCall; vt != nil {
		return (*string)(&vt.CallID)
	} else if vt := u.OfOpenAIResponseInputFunctionToolCallOutput; vt != nil {
		return (*string)(&vt.CallID)
	}
	return nil
}

// This represents the output of a function call that gets passed back to the
// model.
//
// The properties CallID, Output, Type are required.
type ResponseInputOpenAIResponseInputFunctionToolCallOutputParam struct {
	CallID string            `json:"call_id,required"`
	Output string            `json:"output,required"`
	ID     param.Opt[string] `json:"id,omitzero"`
	Status param.Opt[string] `json:"status,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "function_call_output".
	Type constant.FunctionCallOutput `json:"type,required"`
	paramObj
}

func (r ResponseInputOpenAIResponseInputFunctionToolCallOutputParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseInputOpenAIResponseInputFunctionToolCallOutputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseInputOpenAIResponseInputFunctionToolCallOutputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
type ResponseMessage struct {
	Content ResponseMessageContentUnion `json:"content,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseMessageRole `json:"role,required"`
	Type   string              `json:"type,required"`
	ID     string              `json:"id"`
	Status string              `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ResponseMessage to a ResponseMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ResponseMessageParam.Overrides()
func (r ResponseMessage) ToParam() ResponseMessageParam {
	return param.Override[ResponseMessageParam](json.RawMessage(r.RawJSON()))
}

// ResponseMessageContentUnion contains all possible properties and values from
// [string], [[]ResponseMessageContentArrayItemUnion],
// [[]ResponseMessageContentArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfResponseMessageContentArray OfVariant2]
type ResponseMessageContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ResponseMessageContentArrayItemUnion] instead of an object.
	OfResponseMessageContentArray []ResponseMessageContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [[]ResponseMessageContentArrayItem]
	// instead of an object.
	OfVariant2 []ResponseMessageContentArrayItem `json:",inline"`
	JSON       struct {
		OfString                      respjson.Field
		OfResponseMessageContentArray respjson.Field
		OfVariant2                    respjson.Field
		raw                           string
	} `json:"-"`
}

func (u ResponseMessageContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseMessageContentUnion) AsResponseMessageContentArray() (v []ResponseMessageContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseMessageContentUnion) AsVariant2() (v []ResponseMessageContentArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseMessageContentArrayItemUnion contains all possible properties and values
// from [ResponseMessageContentArrayItemInputText],
// [ResponseMessageContentArrayItemInputImage].
//
// Use the [ResponseMessageContentArrayItemUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseMessageContentArrayItemUnion struct {
	// This field is from variant [ResponseMessageContentArrayItemInputText].
	Text string `json:"text"`
	// Any of "input_text", "input_image".
	Type string `json:"type"`
	// This field is from variant [ResponseMessageContentArrayItemInputImage].
	Detail ResponseMessageContentArrayItemInputImageDetail `json:"detail"`
	// This field is from variant [ResponseMessageContentArrayItemInputImage].
	ImageURL string `json:"image_url"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Detail   respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

// anyResponseMessageContentArrayItem is implemented by each variant of
// [ResponseMessageContentArrayItemUnion] to add type safety for the return type of
// [ResponseMessageContentArrayItemUnion.AsAny]
type anyResponseMessageContentArrayItem interface {
	implResponseMessageContentArrayItemUnion()
}

func (ResponseMessageContentArrayItemInputText) implResponseMessageContentArrayItemUnion()  {}
func (ResponseMessageContentArrayItemInputImage) implResponseMessageContentArrayItemUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseMessageContentArrayItemUnion.AsAny().(type) {
//	case llamastackclient.ResponseMessageContentArrayItemInputText:
//	case llamastackclient.ResponseMessageContentArrayItemInputImage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseMessageContentArrayItemUnion) AsAny() anyResponseMessageContentArrayItem {
	switch u.Type {
	case "input_text":
		return u.AsInputText()
	case "input_image":
		return u.AsInputImage()
	}
	return nil
}

func (u ResponseMessageContentArrayItemUnion) AsInputText() (v ResponseMessageContentArrayItemInputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseMessageContentArrayItemUnion) AsInputImage() (v ResponseMessageContentArrayItemInputImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseMessageContentArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseMessageContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseMessageContentArrayItemInputText struct {
	Text string             `json:"text,required"`
	Type constant.InputText `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseMessageContentArrayItemInputText) RawJSON() string { return r.JSON.raw }
func (r *ResponseMessageContentArrayItemInputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseMessageContentArrayItemInputImage struct {
	// Any of "low", "high", "auto".
	Detail   ResponseMessageContentArrayItemInputImageDetail `json:"detail,required"`
	Type     constant.InputImage                             `json:"type,required"`
	ImageURL string                                          `json:"image_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		Type        respjson.Field
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseMessageContentArrayItemInputImage) RawJSON() string { return r.JSON.raw }
func (r *ResponseMessageContentArrayItemInputImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseMessageContentArrayItemInputImageDetail string

const (
	ResponseMessageContentArrayItemInputImageDetailLow  ResponseMessageContentArrayItemInputImageDetail = "low"
	ResponseMessageContentArrayItemInputImageDetailHigh ResponseMessageContentArrayItemInputImageDetail = "high"
	ResponseMessageContentArrayItemInputImageDetailAuto ResponseMessageContentArrayItemInputImageDetail = "auto"
)

type ResponseMessageContentArrayItemDetail string

const (
	ResponseMessageContentArrayItemDetailLow  ResponseMessageContentArrayItemDetail = "low"
	ResponseMessageContentArrayItemDetailHigh ResponseMessageContentArrayItemDetail = "high"
	ResponseMessageContentArrayItemDetailAuto ResponseMessageContentArrayItemDetail = "auto"
)

type ResponseMessageContentArrayItem struct {
	Annotations []ResponseMessageContentArrayItemAnnotationUnion `json:"annotations,required"`
	Text        string                                           `json:"text,required"`
	Type        string                                           `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Annotations respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseMessageContentArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ResponseMessageContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseMessageContentArrayItemAnnotationUnion contains all possible properties
// and values from [ResponseMessageContentArrayItemAnnotationFileCitation],
// [ResponseMessageContentArrayItemAnnotationURLCitation],
// [ResponseMessageContentArrayItemAnnotationContainerFileCitation],
// [ResponseMessageContentArrayItemAnnotationFilePath].
//
// Use the [ResponseMessageContentArrayItemAnnotationUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseMessageContentArrayItemAnnotationUnion struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	Index    int64  `json:"index"`
	// Any of "file_citation", "url_citation", "container_file_citation", "file_path".
	Type       string `json:"type"`
	EndIndex   int64  `json:"end_index"`
	StartIndex int64  `json:"start_index"`
	// This field is from variant
	// [ResponseMessageContentArrayItemAnnotationURLCitation].
	Title string `json:"title"`
	// This field is from variant
	// [ResponseMessageContentArrayItemAnnotationURLCitation].
	URL string `json:"url"`
	// This field is from variant
	// [ResponseMessageContentArrayItemAnnotationContainerFileCitation].
	ContainerID string `json:"container_id"`
	JSON        struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ContainerID respjson.Field
		raw         string
	} `json:"-"`
}

// anyResponseMessageContentArrayItemAnnotation is implemented by each variant of
// [ResponseMessageContentArrayItemAnnotationUnion] to add type safety for the
// return type of [ResponseMessageContentArrayItemAnnotationUnion.AsAny]
type anyResponseMessageContentArrayItemAnnotation interface {
	implResponseMessageContentArrayItemAnnotationUnion()
}

func (ResponseMessageContentArrayItemAnnotationFileCitation) implResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseMessageContentArrayItemAnnotationURLCitation) implResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseMessageContentArrayItemAnnotationContainerFileCitation) implResponseMessageContentArrayItemAnnotationUnion() {
}
func (ResponseMessageContentArrayItemAnnotationFilePath) implResponseMessageContentArrayItemAnnotationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseMessageContentArrayItemAnnotationUnion.AsAny().(type) {
//	case llamastackclient.ResponseMessageContentArrayItemAnnotationFileCitation:
//	case llamastackclient.ResponseMessageContentArrayItemAnnotationURLCitation:
//	case llamastackclient.ResponseMessageContentArrayItemAnnotationContainerFileCitation:
//	case llamastackclient.ResponseMessageContentArrayItemAnnotationFilePath:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseMessageContentArrayItemAnnotationUnion) AsAny() anyResponseMessageContentArrayItemAnnotation {
	switch u.Type {
	case "file_citation":
		return u.AsFileCitation()
	case "url_citation":
		return u.AsURLCitation()
	case "container_file_citation":
		return u.AsContainerFileCitation()
	case "file_path":
		return u.AsFilePath()
	}
	return nil
}

func (u ResponseMessageContentArrayItemAnnotationUnion) AsFileCitation() (v ResponseMessageContentArrayItemAnnotationFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseMessageContentArrayItemAnnotationUnion) AsURLCitation() (v ResponseMessageContentArrayItemAnnotationURLCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseMessageContentArrayItemAnnotationUnion) AsContainerFileCitation() (v ResponseMessageContentArrayItemAnnotationContainerFileCitation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseMessageContentArrayItemAnnotationUnion) AsFilePath() (v ResponseMessageContentArrayItemAnnotationFilePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseMessageContentArrayItemAnnotationUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseMessageContentArrayItemAnnotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseMessageContentArrayItemAnnotationFileCitation struct {
	FileID   string                `json:"file_id,required"`
	Filename string                `json:"filename,required"`
	Index    int64                 `json:"index,required"`
	Type     constant.FileCitation `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Filename    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseMessageContentArrayItemAnnotationFileCitation) RawJSON() string { return r.JSON.raw }
func (r *ResponseMessageContentArrayItemAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseMessageContentArrayItemAnnotationURLCitation struct {
	EndIndex   int64                `json:"end_index,required"`
	StartIndex int64                `json:"start_index,required"`
	Title      string               `json:"title,required"`
	Type       constant.URLCitation `json:"type,required"`
	URL        string               `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseMessageContentArrayItemAnnotationURLCitation) RawJSON() string { return r.JSON.raw }
func (r *ResponseMessageContentArrayItemAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseMessageContentArrayItemAnnotationContainerFileCitation struct {
	ContainerID string                         `json:"container_id,required"`
	EndIndex    int64                          `json:"end_index,required"`
	FileID      string                         `json:"file_id,required"`
	Filename    string                         `json:"filename,required"`
	StartIndex  int64                          `json:"start_index,required"`
	Type        constant.ContainerFileCitation `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerID respjson.Field
		EndIndex    respjson.Field
		FileID      respjson.Field
		Filename    respjson.Field
		StartIndex  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseMessageContentArrayItemAnnotationContainerFileCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *ResponseMessageContentArrayItemAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseMessageContentArrayItemAnnotationFilePath struct {
	FileID string            `json:"file_id,required"`
	Index  int64             `json:"index,required"`
	Type   constant.FilePath `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseMessageContentArrayItemAnnotationFilePath) RawJSON() string { return r.JSON.raw }
func (r *ResponseMessageContentArrayItemAnnotationFilePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseMessageRole string

const (
	ResponseMessageRoleSystem    ResponseMessageRole = "system"
	ResponseMessageRoleDeveloper ResponseMessageRole = "developer"
	ResponseMessageRoleUser      ResponseMessageRole = "user"
	ResponseMessageRoleAssistant ResponseMessageRole = "assistant"
)

// Corresponds to the various Message types in the Responses API. They are all
// under one type because the Responses API gives them all the same "type" value,
// and there is no way to tell them apart in certain scenarios.
//
// The properties Content, Role, Type are required.
type ResponseMessageParam struct {
	Content ResponseMessageContentUnionParam `json:"content,omitzero,required"`
	// Any of "system", "developer", "user", "assistant".
	Role   ResponseMessageRole `json:"role,omitzero,required"`
	Type   string              `json:"type,required"`
	ID     param.Opt[string]   `json:"id,omitzero"`
	Status param.Opt[string]   `json:"status,omitzero"`
	paramObj
}

func (r ResponseMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseMessageContentUnionParam struct {
	OfString                      param.Opt[string]                           `json:",omitzero,inline"`
	OfResponseMessageContentArray []ResponseMessageContentArrayItemUnionParam `json:",omitzero,inline"`
	OfVariant2                    []ResponseMessageContentArrayItemParam      `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseMessageContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfResponseMessageContentArray, u.OfVariant2)
}
func (u *ResponseMessageContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseMessageContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfResponseMessageContentArray) {
		return &u.OfResponseMessageContentArray
	} else if !param.IsOmitted(u.OfVariant2) {
		return &u.OfVariant2
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseMessageContentArrayItemUnionParam struct {
	OfInputText  *ResponseMessageContentArrayItemInputTextParam  `json:",omitzero,inline"`
	OfInputImage *ResponseMessageContentArrayItemInputImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseMessageContentArrayItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInputText, u.OfInputImage)
}
func (u *ResponseMessageContentArrayItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseMessageContentArrayItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfInputText) {
		return u.OfInputText
	} else if !param.IsOmitted(u.OfInputImage) {
		return u.OfInputImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemUnionParam) GetText() *string {
	if vt := u.OfInputText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemUnionParam) GetDetail() *string {
	if vt := u.OfInputImage; vt != nil {
		return (*string)(&vt.Detail)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemUnionParam) GetImageURL() *string {
	if vt := u.OfInputImage; vt != nil && vt.ImageURL.Valid() {
		return &vt.ImageURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemUnionParam) GetType() *string {
	if vt := u.OfInputText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInputImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseMessageContentArrayItemUnionParam](
		"type",
		apijson.Discriminator[ResponseMessageContentArrayItemInputTextParam]("input_text"),
		apijson.Discriminator[ResponseMessageContentArrayItemInputImageParam]("input_image"),
	)
}

// The properties Text, Type are required.
type ResponseMessageContentArrayItemInputTextParam struct {
	Text string `json:"text,required"`
	// This field can be elided, and will marshal its zero value as "input_text".
	Type constant.InputText `json:"type,required"`
	paramObj
}

func (r ResponseMessageContentArrayItemInputTextParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseMessageContentArrayItemInputTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseMessageContentArrayItemInputTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Detail, Type are required.
type ResponseMessageContentArrayItemInputImageParam struct {
	// Any of "low", "high", "auto".
	Detail   ResponseMessageContentArrayItemInputImageDetail `json:"detail,omitzero,required"`
	ImageURL param.Opt[string]                               `json:"image_url,omitzero"`
	// This field can be elided, and will marshal its zero value as "input_image".
	Type constant.InputImage `json:"type,required"`
	paramObj
}

func (r ResponseMessageContentArrayItemInputImageParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseMessageContentArrayItemInputImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseMessageContentArrayItemInputImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Annotations, Text, Type are required.
type ResponseMessageContentArrayItemParam struct {
	Annotations []ResponseMessageContentArrayItemAnnotationUnionParam `json:"annotations,omitzero,required"`
	Text        string                                                `json:"text,required"`
	Type        string                                                `json:"type,required"`
	paramObj
}

func (r ResponseMessageContentArrayItemParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseMessageContentArrayItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseMessageContentArrayItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseMessageContentArrayItemAnnotationUnionParam struct {
	OfFileCitation          *ResponseMessageContentArrayItemAnnotationFileCitationParam          `json:",omitzero,inline"`
	OfURLCitation           *ResponseMessageContentArrayItemAnnotationURLCitationParam           `json:",omitzero,inline"`
	OfContainerFileCitation *ResponseMessageContentArrayItemAnnotationContainerFileCitationParam `json:",omitzero,inline"`
	OfFilePath              *ResponseMessageContentArrayItemAnnotationFilePathParam              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseMessageContentArrayItemAnnotationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFileCitation, u.OfURLCitation, u.OfContainerFileCitation, u.OfFilePath)
}
func (u *ResponseMessageContentArrayItemAnnotationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseMessageContentArrayItemAnnotationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFileCitation) {
		return u.OfFileCitation
	} else if !param.IsOmitted(u.OfURLCitation) {
		return u.OfURLCitation
	} else if !param.IsOmitted(u.OfContainerFileCitation) {
		return u.OfContainerFileCitation
	} else if !param.IsOmitted(u.OfFilePath) {
		return u.OfFilePath
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemAnnotationUnionParam) GetTitle() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemAnnotationUnionParam) GetURL() *string {
	if vt := u.OfURLCitation; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemAnnotationUnionParam) GetContainerID() *string {
	if vt := u.OfContainerFileCitation; vt != nil {
		return &vt.ContainerID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemAnnotationUnionParam) GetFileID() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.FileID)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.FileID)
	} else if vt := u.OfFilePath; vt != nil {
		return (*string)(&vt.FileID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemAnnotationUnionParam) GetFilename() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.Filename)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemAnnotationUnionParam) GetIndex() *int64 {
	if vt := u.OfFileCitation; vt != nil {
		return (*int64)(&vt.Index)
	} else if vt := u.OfFilePath; vt != nil {
		return (*int64)(&vt.Index)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemAnnotationUnionParam) GetType() *string {
	if vt := u.OfFileCitation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfURLCitation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFilePath; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemAnnotationUnionParam) GetEndIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.EndIndex)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseMessageContentArrayItemAnnotationUnionParam) GetStartIndex() *int64 {
	if vt := u.OfURLCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	} else if vt := u.OfContainerFileCitation; vt != nil {
		return (*int64)(&vt.StartIndex)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseMessageContentArrayItemAnnotationUnionParam](
		"type",
		apijson.Discriminator[ResponseMessageContentArrayItemAnnotationFileCitationParam]("file_citation"),
		apijson.Discriminator[ResponseMessageContentArrayItemAnnotationURLCitationParam]("url_citation"),
		apijson.Discriminator[ResponseMessageContentArrayItemAnnotationContainerFileCitationParam]("container_file_citation"),
		apijson.Discriminator[ResponseMessageContentArrayItemAnnotationFilePathParam]("file_path"),
	)
}

// The properties FileID, Filename, Index, Type are required.
type ResponseMessageContentArrayItemAnnotationFileCitationParam struct {
	FileID   string `json:"file_id,required"`
	Filename string `json:"filename,required"`
	Index    int64  `json:"index,required"`
	// This field can be elided, and will marshal its zero value as "file_citation".
	Type constant.FileCitation `json:"type,required"`
	paramObj
}

func (r ResponseMessageContentArrayItemAnnotationFileCitationParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseMessageContentArrayItemAnnotationFileCitationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseMessageContentArrayItemAnnotationFileCitationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EndIndex, StartIndex, Title, Type, URL are required.
type ResponseMessageContentArrayItemAnnotationURLCitationParam struct {
	EndIndex   int64  `json:"end_index,required"`
	StartIndex int64  `json:"start_index,required"`
	Title      string `json:"title,required"`
	URL        string `json:"url,required"`
	// This field can be elided, and will marshal its zero value as "url_citation".
	Type constant.URLCitation `json:"type,required"`
	paramObj
}

func (r ResponseMessageContentArrayItemAnnotationURLCitationParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseMessageContentArrayItemAnnotationURLCitationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseMessageContentArrayItemAnnotationURLCitationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ContainerID, EndIndex, FileID, Filename, StartIndex, Type are
// required.
type ResponseMessageContentArrayItemAnnotationContainerFileCitationParam struct {
	ContainerID string `json:"container_id,required"`
	EndIndex    int64  `json:"end_index,required"`
	FileID      string `json:"file_id,required"`
	Filename    string `json:"filename,required"`
	StartIndex  int64  `json:"start_index,required"`
	// This field can be elided, and will marshal its zero value as
	// "container_file_citation".
	Type constant.ContainerFileCitation `json:"type,required"`
	paramObj
}

func (r ResponseMessageContentArrayItemAnnotationContainerFileCitationParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseMessageContentArrayItemAnnotationContainerFileCitationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseMessageContentArrayItemAnnotationContainerFileCitationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FileID, Index, Type are required.
type ResponseMessageContentArrayItemAnnotationFilePathParam struct {
	FileID string `json:"file_id,required"`
	Index  int64  `json:"index,required"`
	// This field can be elided, and will marshal its zero value as "file_path".
	Type constant.FilePath `json:"type,required"`
	paramObj
}

func (r ResponseMessageContentArrayItemAnnotationFilePathParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseMessageContentArrayItemAnnotationFilePathParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseMessageContentArrayItemAnnotationFilePathParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseObject struct {
	ID                 string            `json:"id,required"`
	CreatedAt          int64             `json:"created_at,required"`
	Model              string            `json:"model,required"`
	Object             constant.Response `json:"object,required"`
	Output             []OutputUnion     `json:"output,required"`
	ParallelToolCalls  bool              `json:"parallel_tool_calls,required"`
	Status             string            `json:"status,required"`
	Text               ResponseText      `json:"text,required"`
	Error              ResponseError     `json:"error"`
	PreviousResponseID string            `json:"previous_response_id"`
	Temperature        float64           `json:"temperature"`
	TopP               float64           `json:"top_p"`
	Truncation         string            `json:"truncation"`
	User               string            `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Model              respjson.Field
		Object             respjson.Field
		Output             respjson.Field
		ParallelToolCalls  respjson.Field
		Status             respjson.Field
		Text               respjson.Field
		Error              respjson.Field
		PreviousResponseID respjson.Field
		Temperature        respjson.Field
		TopP               respjson.Field
		Truncation         respjson.Field
		User               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseObject) RawJSON() string { return r.JSON.raw }
func (r *ResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseText struct {
	// Configuration for Responses API text format.
	Format ResponseTextFormat `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseText) RawJSON() string { return r.JSON.raw }
func (r *ResponseText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ResponseText to a ResponseTextParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ResponseTextParam.Overrides()
func (r ResponseText) ToParam() ResponseTextParam {
	return param.Override[ResponseTextParam](json.RawMessage(r.RawJSON()))
}

// Configuration for Responses API text format.
type ResponseTextFormat struct {
	// Must be "text", "json_schema", or "json_object" to identify the format type
	//
	// Any of "text", "json_schema", "json_object".
	Type ResponseTextFormatType `json:"type,required"`
	// (Optional) A description of the response format. Only used for json_schema.
	Description string `json:"description"`
	// The name of the response format. Only used for json_schema.
	Name string `json:"name"`
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model. Only used for json_schema.
	Schema map[string]ResponseTextFormatSchemaUnion `json:"schema"`
	// (Optional) Whether to strictly enforce the JSON schema. If true, the response
	// must match the schema exactly. Only used for json_schema.
	Strict bool `json:"strict"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Schema      respjson.Field
		Strict      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseTextFormat) RawJSON() string { return r.JSON.raw }
func (r *ResponseTextFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be "text", "json_schema", or "json_object" to identify the format type
type ResponseTextFormatType string

const (
	ResponseTextFormatTypeText       ResponseTextFormatType = "text"
	ResponseTextFormatTypeJsonSchema ResponseTextFormatType = "json_schema"
	ResponseTextFormatTypeJsonObject ResponseTextFormatType = "json_object"
)

// ResponseTextFormatSchemaUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseTextFormatSchemaUnion struct {
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

func (u ResponseTextFormatSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseTextFormatSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseTextFormatSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseTextFormatSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseTextFormatSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseTextFormatSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseTextParam struct {
	// Configuration for Responses API text format.
	Format ResponseTextFormatParam `json:"format,omitzero"`
	paramObj
}

func (r ResponseTextParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for Responses API text format.
//
// The property Type is required.
type ResponseTextFormatParam struct {
	// Must be "text", "json_schema", or "json_object" to identify the format type
	//
	// Any of "text", "json_schema", "json_object".
	Type ResponseTextFormatType `json:"type,omitzero,required"`
	// (Optional) A description of the response format. Only used for json_schema.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the response format. Only used for json_schema.
	Name param.Opt[string] `json:"name,omitzero"`
	// (Optional) Whether to strictly enforce the JSON schema. If true, the response
	// must match the schema exactly. Only used for json_schema.
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model. Only used for json_schema.
	Schema map[string]ResponseTextFormatSchemaUnionParam `json:"schema,omitzero"`
	paramObj
}

func (r ResponseTextFormatParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseTextFormatParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseTextFormatParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseTextFormatSchemaUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseTextFormatSchemaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseTextFormatSchemaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseTextFormatSchemaUnionParam) asAny() any {
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

type OpenAiv1ResponseListResponse struct {
	Data    []OpenAiv1ResponseListResponseData `json:"data,required"`
	FirstID string                             `json:"first_id,required"`
	HasMore bool                               `json:"has_more,required"`
	LastID  string                             `json:"last_id,required"`
	Object  constant.List                      `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		FirstID     respjson.Field
		HasMore     respjson.Field
		LastID      respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ResponseListResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1ResponseListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAiv1ResponseListResponseData struct {
	ID                 string               `json:"id,required"`
	CreatedAt          int64                `json:"created_at,required"`
	Input              []ResponseInputUnion `json:"input,required"`
	Model              string               `json:"model,required"`
	Object             constant.Response    `json:"object,required"`
	Output             []OutputUnion        `json:"output,required"`
	ParallelToolCalls  bool                 `json:"parallel_tool_calls,required"`
	Status             string               `json:"status,required"`
	Text               ResponseText         `json:"text,required"`
	Error              ResponseError        `json:"error"`
	PreviousResponseID string               `json:"previous_response_id"`
	Temperature        float64              `json:"temperature"`
	TopP               float64              `json:"top_p"`
	Truncation         string               `json:"truncation"`
	User               string               `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Input              respjson.Field
		Model              respjson.Field
		Object             respjson.Field
		Output             respjson.Field
		ParallelToolCalls  respjson.Field
		Status             respjson.Field
		Text               respjson.Field
		Error              respjson.Field
		PreviousResponseID respjson.Field
		Temperature        respjson.Field
		TopP               respjson.Field
		Truncation         respjson.Field
		User               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ResponseListResponseData) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1ResponseListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAiv1ResponseGetInputItemsResponse struct {
	Data   []ResponseInputUnion `json:"data,required"`
	Object constant.List        `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAiv1ResponseGetInputItemsResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenAiv1ResponseGetInputItemsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIV1ResponseNewParams struct {
	// Input message(s) to create the response.
	Input OpenAIV1ResponseNewParamsInputUnion `json:"input,omitzero,required"`
	// The underlying LLM used for completions.
	Model         string            `json:"model,required"`
	Instructions  param.Opt[string] `json:"instructions,omitzero"`
	MaxInferIters param.Opt[int64]  `json:"max_infer_iters,omitzero"`
	// (Optional) if specified, the new response will be a continuation of the previous
	// response. This can be used to easily fork-off new responses from existing
	// responses.
	PreviousResponseID param.Opt[string]                    `json:"previous_response_id,omitzero"`
	Store              param.Opt[bool]                      `json:"store,omitzero"`
	Stream             param.Opt[bool]                      `json:"stream,omitzero"`
	Temperature        param.Opt[float64]                   `json:"temperature,omitzero"`
	Text               ResponseTextParam                    `json:"text,omitzero"`
	Tools              []OpenAIV1ResponseNewParamsToolUnion `json:"tools,omitzero"`
	paramObj
}

func (r OpenAIV1ResponseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ResponseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ResponseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ResponseNewParamsInputUnion struct {
	OfString             param.Opt[string]         `json:",omitzero,inline"`
	OfResponseInputArray []ResponseInputUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ResponseNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfResponseInputArray)
}
func (u *OpenAIV1ResponseNewParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ResponseNewParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfResponseInputArray) {
		return &u.OfResponseInputArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ResponseNewParamsToolUnion struct {
	OfOpenAIResponseInputToolWebSearch *OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearch `json:",omitzero,inline"`
	OfFileSearch                       *OpenAIV1ResponseNewParamsToolFileSearch                       `json:",omitzero,inline"`
	OfFunction                         *OpenAIV1ResponseNewParamsToolFunction                         `json:",omitzero,inline"`
	OfMcp                              *OpenAIV1ResponseNewParamsToolMcp                              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ResponseNewParamsToolUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIResponseInputToolWebSearch, u.OfFileSearch, u.OfFunction, u.OfMcp)
}
func (u *OpenAIV1ResponseNewParamsToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ResponseNewParamsToolUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIResponseInputToolWebSearch) {
		return u.OfOpenAIResponseInputToolWebSearch
	} else if !param.IsOmitted(u.OfFileSearch) {
		return u.OfFileSearch
	} else if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	} else if !param.IsOmitted(u.OfMcp) {
		return u.OfMcp
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetSearchContextSize() *string {
	if vt := u.OfOpenAIResponseInputToolWebSearch; vt != nil && vt.SearchContextSize.Valid() {
		return &vt.SearchContextSize.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetVectorStoreIDs() []string {
	if vt := u.OfFileSearch; vt != nil {
		return vt.VectorStoreIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetFilters() map[string]OpenAIV1ResponseNewParamsToolFileSearchFilterUnion {
	if vt := u.OfFileSearch; vt != nil {
		return vt.Filters
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetMaxNumResults() *int64 {
	if vt := u.OfFileSearch; vt != nil && vt.MaxNumResults.Valid() {
		return &vt.MaxNumResults.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetRankingOptions() *OpenAIV1ResponseNewParamsToolFileSearchRankingOptions {
	if vt := u.OfFileSearch; vt != nil {
		return &vt.RankingOptions
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetName() *string {
	if vt := u.OfFunction; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetDescription() *string {
	if vt := u.OfFunction; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetParameters() map[string]OpenAIV1ResponseNewParamsToolFunctionParameterUnion {
	if vt := u.OfFunction; vt != nil {
		return vt.Parameters
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetStrict() *bool {
	if vt := u.OfFunction; vt != nil && vt.Strict.Valid() {
		return &vt.Strict.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetRequireApproval() *OpenAIV1ResponseNewParamsToolMcpRequireApprovalUnion {
	if vt := u.OfMcp; vt != nil {
		return &vt.RequireApproval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetServerLabel() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.ServerLabel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetServerURL() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.ServerURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetAllowedTools() *OpenAIV1ResponseNewParamsToolMcpAllowedToolsUnion {
	if vt := u.OfMcp; vt != nil {
		return &vt.AllowedTools
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetHeaders() map[string]OpenAIV1ResponseNewParamsToolMcpHeaderUnion {
	if vt := u.OfMcp; vt != nil {
		return vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OpenAIV1ResponseNewParamsToolUnion) GetType() *string {
	if vt := u.OfOpenAIResponseInputToolWebSearch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFileSearch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcp; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[OpenAIV1ResponseNewParamsToolUnion](
		"type",
		apijson.Discriminator[OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearch]("web_search"),
		apijson.Discriminator[OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearch]("web_search_preview"),
		apijson.Discriminator[OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearch]("web_search_preview_2025_03_11"),
		apijson.Discriminator[OpenAIV1ResponseNewParamsToolFileSearch]("file_search"),
		apijson.Discriminator[OpenAIV1ResponseNewParamsToolFunction]("function"),
		apijson.Discriminator[OpenAIV1ResponseNewParamsToolMcp]("mcp"),
	)
}

// The property Type is required.
type OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearch struct {
	// Any of "web_search", "web_search_preview", "web_search_preview_2025_03_11".
	Type              OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearchType `json:"type,omitzero,required"`
	SearchContextSize param.Opt[string]                                                 `json:"search_context_size,omitzero"`
	paramObj
}

func (r OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearch) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearchType string

const (
	OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearch                  OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearchType = "web_search"
	OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearchPreview           OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearchType = "web_search_preview"
	OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearchPreview2025_03_11 OpenAIV1ResponseNewParamsToolOpenAIResponseInputToolWebSearchType = "web_search_preview_2025_03_11"
)

// The properties Type, VectorStoreIDs are required.
type OpenAIV1ResponseNewParamsToolFileSearch struct {
	VectorStoreIDs []string                                                      `json:"vector_store_ids,omitzero,required"`
	MaxNumResults  param.Opt[int64]                                              `json:"max_num_results,omitzero"`
	Filters        map[string]OpenAIV1ResponseNewParamsToolFileSearchFilterUnion `json:"filters,omitzero"`
	RankingOptions OpenAIV1ResponseNewParamsToolFileSearchRankingOptions         `json:"ranking_options,omitzero"`
	// This field can be elided, and will marshal its zero value as "file_search".
	Type constant.FileSearch `json:"type,required"`
	paramObj
}

func (r OpenAIV1ResponseNewParamsToolFileSearch) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ResponseNewParamsToolFileSearch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ResponseNewParamsToolFileSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ResponseNewParamsToolFileSearchFilterUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ResponseNewParamsToolFileSearchFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1ResponseNewParamsToolFileSearchFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ResponseNewParamsToolFileSearchFilterUnion) asAny() any {
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

type OpenAIV1ResponseNewParamsToolFileSearchRankingOptions struct {
	Ranker         param.Opt[string]  `json:"ranker,omitzero"`
	ScoreThreshold param.Opt[float64] `json:"score_threshold,omitzero"`
	paramObj
}

func (r OpenAIV1ResponseNewParamsToolFileSearchRankingOptions) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ResponseNewParamsToolFileSearchRankingOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ResponseNewParamsToolFileSearchRankingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type OpenAIV1ResponseNewParamsToolFunction struct {
	Name        string                                                         `json:"name,required"`
	Description param.Opt[string]                                              `json:"description,omitzero"`
	Strict      param.Opt[bool]                                                `json:"strict,omitzero"`
	Parameters  map[string]OpenAIV1ResponseNewParamsToolFunctionParameterUnion `json:"parameters,omitzero"`
	// This field can be elided, and will marshal its zero value as "function".
	Type constant.Function `json:"type,required"`
	paramObj
}

func (r OpenAIV1ResponseNewParamsToolFunction) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ResponseNewParamsToolFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ResponseNewParamsToolFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ResponseNewParamsToolFunctionParameterUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ResponseNewParamsToolFunctionParameterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1ResponseNewParamsToolFunctionParameterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ResponseNewParamsToolFunctionParameterUnion) asAny() any {
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

// The properties RequireApproval, ServerLabel, ServerURL, Type are required.
type OpenAIV1ResponseNewParamsToolMcp struct {
	RequireApproval OpenAIV1ResponseNewParamsToolMcpRequireApprovalUnion   `json:"require_approval,omitzero,required"`
	ServerLabel     string                                                 `json:"server_label,required"`
	ServerURL       string                                                 `json:"server_url,required"`
	AllowedTools    OpenAIV1ResponseNewParamsToolMcpAllowedToolsUnion      `json:"allowed_tools,omitzero"`
	Headers         map[string]OpenAIV1ResponseNewParamsToolMcpHeaderUnion `json:"headers,omitzero"`
	// This field can be elided, and will marshal its zero value as "mcp".
	Type constant.Mcp `json:"type,required"`
	paramObj
}

func (r OpenAIV1ResponseNewParamsToolMcp) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ResponseNewParamsToolMcp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ResponseNewParamsToolMcp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ResponseNewParamsToolMcpRequireApprovalUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfOpenAIV1ResponseNewsToolMcpRequireApprovalString)
	OfOpenAIV1ResponseNewsToolMcpRequireApprovalString param.Opt[OpenAIV1ResponseNewParamsToolMcpRequireApprovalString] `json:",omitzero,inline"`
	OfApprovalFilter                                   *OpenAIV1ResponseNewParamsToolMcpRequireApprovalApprovalFilter   `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ResponseNewParamsToolMcpRequireApprovalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenAIV1ResponseNewsToolMcpRequireApprovalString, u.OfApprovalFilter)
}
func (u *OpenAIV1ResponseNewParamsToolMcpRequireApprovalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ResponseNewParamsToolMcpRequireApprovalUnion) asAny() any {
	if !param.IsOmitted(u.OfOpenAIV1ResponseNewsToolMcpRequireApprovalString) {
		return &u.OfOpenAIV1ResponseNewsToolMcpRequireApprovalString
	} else if !param.IsOmitted(u.OfApprovalFilter) {
		return u.OfApprovalFilter
	}
	return nil
}

type OpenAIV1ResponseNewParamsToolMcpRequireApprovalString string

const (
	OpenAIV1ResponseNewParamsToolMcpRequireApprovalStringAlways OpenAIV1ResponseNewParamsToolMcpRequireApprovalString = "always"
	OpenAIV1ResponseNewParamsToolMcpRequireApprovalStringNever  OpenAIV1ResponseNewParamsToolMcpRequireApprovalString = "never"
)

type OpenAIV1ResponseNewParamsToolMcpRequireApprovalApprovalFilter struct {
	Always []string `json:"always,omitzero"`
	Never  []string `json:"never,omitzero"`
	paramObj
}

func (r OpenAIV1ResponseNewParamsToolMcpRequireApprovalApprovalFilter) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ResponseNewParamsToolMcpRequireApprovalApprovalFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ResponseNewParamsToolMcpRequireApprovalApprovalFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ResponseNewParamsToolMcpAllowedToolsUnion struct {
	OfStringArray        []string                                                        `json:",omitzero,inline"`
	OfAllowedToolsFilter *OpenAIV1ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ResponseNewParamsToolMcpAllowedToolsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringArray, u.OfAllowedToolsFilter)
}
func (u *OpenAIV1ResponseNewParamsToolMcpAllowedToolsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ResponseNewParamsToolMcpAllowedToolsUnion) asAny() any {
	if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfAllowedToolsFilter) {
		return u.OfAllowedToolsFilter
	}
	return nil
}

type OpenAIV1ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter struct {
	ToolNames []string `json:"tool_names,omitzero"`
	paramObj
}

func (r OpenAIV1ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIV1ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIV1ResponseNewParamsToolMcpAllowedToolsAllowedToolsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OpenAIV1ResponseNewParamsToolMcpHeaderUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u OpenAIV1ResponseNewParamsToolMcpHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *OpenAIV1ResponseNewParamsToolMcpHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OpenAIV1ResponseNewParamsToolMcpHeaderUnion) asAny() any {
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

type OpenAIV1ResponseListParams struct {
	// The ID of the last response to return.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The number of responses to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The model to filter responses by.
	Model param.Opt[string] `query:"model,omitzero" json:"-"`
	// The order to sort responses by when sorted by created_at ('asc' or 'desc').
	//
	// Any of "asc", "desc".
	Order Order `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OpenAIV1ResponseListParams]'s query parameters as
// `url.Values`.
func (r OpenAIV1ResponseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OpenAIV1ResponseGetInputItemsParams struct {
	// An item ID to list items after, used for pagination.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// An item ID to list items before, used for pagination.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// A limit on the number of objects to be returned. Limit can range between 1 and
	// 100, and the default is 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Additional fields to include in the response.
	Include []string `query:"include,omitzero" json:"-"`
	// The order to return the input items in. Default is desc.
	//
	// Any of "asc", "desc".
	Order Order `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OpenAIV1ResponseGetInputItemsParams]'s query parameters as
// `url.Values`.
func (r OpenAIV1ResponseGetInputItemsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

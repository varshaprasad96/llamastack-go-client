// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/varshaprasad96/llamastack-go-client/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type AgentTurnInput string      // Always "agent_turn_input"
type Array string               // Always "array"
type Assistant string           // Always "assistant"
type Boolean string             // Always "boolean"
type ChatCompletionInput string // Always "chat_completion_input"
type CompletionInput string     // Always "completion_input"
type Grammar string             // Always "grammar"
type Greedy string              // Always "greedy"
type Image string               // Always "image"
type Json string                // Always "json"
type JsonSchema string          // Always "json_schema"
type Number string              // Always "number"
type Object string              // Always "object"
type String string              // Always "string"
type System string              // Always "system"
type Text string                // Always "text"
type Tool string                // Always "tool"
type TopK string                // Always "top_k"
type TopP string                // Always "top_p"
type Union string               // Always "union"
type User string                // Always "user"

func (c AgentTurnInput) Default() AgentTurnInput           { return "agent_turn_input" }
func (c Array) Default() Array                             { return "array" }
func (c Assistant) Default() Assistant                     { return "assistant" }
func (c Boolean) Default() Boolean                         { return "boolean" }
func (c ChatCompletionInput) Default() ChatCompletionInput { return "chat_completion_input" }
func (c CompletionInput) Default() CompletionInput         { return "completion_input" }
func (c Grammar) Default() Grammar                         { return "grammar" }
func (c Greedy) Default() Greedy                           { return "greedy" }
func (c Image) Default() Image                             { return "image" }
func (c Json) Default() Json                               { return "json" }
func (c JsonSchema) Default() JsonSchema                   { return "json_schema" }
func (c Number) Default() Number                           { return "number" }
func (c Object) Default() Object                           { return "object" }
func (c String) Default() String                           { return "string" }
func (c System) Default() System                           { return "system" }
func (c Text) Default() Text                               { return "text" }
func (c Tool) Default() Tool                               { return "tool" }
func (c TopK) Default() TopK                               { return "top_k" }
func (c TopP) Default() TopP                               { return "top_p" }
func (c Union) Default() Union                             { return "union" }
func (c User) Default() User                               { return "user" }

func (c AgentTurnInput) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Array) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Assistant) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Boolean) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c ChatCompletionInput) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c CompletionInput) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Grammar) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Greedy) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Image) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Json) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c JsonSchema) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Number) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Object) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c String) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c System) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Tool) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c TopK) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c TopP) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Union) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c User) MarshalJSON() ([]byte, error)                { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}

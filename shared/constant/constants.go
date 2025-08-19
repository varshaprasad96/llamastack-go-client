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

type Agent string                 // Always "agent"
type ChatCompletion string        // Always "chat.completion"
type ChatCompletionChunk string   // Always "chat.completion.chunk"
type ContainerFileCitation string // Always "container_file_citation"
type Default string               // Always "default"
type File string                  // Always "file"
type FileCitation string          // Always "file_citation"
type FilePath string              // Always "file_path"
type FileSearch string            // Always "file_search"
type Function string              // Always "function"
type FunctionCallOutput string    // Always "function_call_output"
type InputImage string            // Always "input_image"
type InputText string             // Always "input_text"
type Json string                  // Always "json"
type List string                  // Always "list"
type Llm string                   // Always "llm"
type LoRa string                  // Always "LoRA"
type Mcp string                   // Always "mcp"
type McpCall string               // Always "mcp_call"
type McpListTools string          // Always "mcp_list_tools"
type Model string                 // Always "model"
type Qat string                   // Always "QAT"
type Response string              // Always "response"
type Rrf string                   // Always "rrf"
type Text string                  // Always "text"
type TextCompletion string        // Always "text_completion"
type URLCitation string           // Always "url_citation"
type WebSearchCall string         // Always "web_search_call"
type Weighted string              // Always "weighted"

func (c Agent) Default() Agent                                 { return "agent" }
func (c ChatCompletion) Default() ChatCompletion               { return "chat.completion" }
func (c ChatCompletionChunk) Default() ChatCompletionChunk     { return "chat.completion.chunk" }
func (c ContainerFileCitation) Default() ContainerFileCitation { return "container_file_citation" }
func (c Default) Default() Default                             { return "default" }
func (c File) Default() File                                   { return "file" }
func (c FileCitation) Default() FileCitation                   { return "file_citation" }
func (c FilePath) Default() FilePath                           { return "file_path" }
func (c FileSearch) Default() FileSearch                       { return "file_search" }
func (c Function) Default() Function                           { return "function" }
func (c FunctionCallOutput) Default() FunctionCallOutput       { return "function_call_output" }
func (c InputImage) Default() InputImage                       { return "input_image" }
func (c InputText) Default() InputText                         { return "input_text" }
func (c Json) Default() Json                                   { return "json" }
func (c List) Default() List                                   { return "list" }
func (c Llm) Default() Llm                                     { return "llm" }
func (c LoRa) Default() LoRa                                   { return "LoRA" }
func (c Mcp) Default() Mcp                                     { return "mcp" }
func (c McpCall) Default() McpCall                             { return "mcp_call" }
func (c McpListTools) Default() McpListTools                   { return "mcp_list_tools" }
func (c Model) Default() Model                                 { return "model" }
func (c Qat) Default() Qat                                     { return "QAT" }
func (c Response) Default() Response                           { return "response" }
func (c Rrf) Default() Rrf                                     { return "rrf" }
func (c Text) Default() Text                                   { return "text" }
func (c TextCompletion) Default() TextCompletion               { return "text_completion" }
func (c URLCitation) Default() URLCitation                     { return "url_citation" }
func (c WebSearchCall) Default() WebSearchCall                 { return "web_search_call" }
func (c Weighted) Default() Weighted                           { return "weighted" }

func (c Agent) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c ChatCompletion) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ChatCompletionChunk) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c ContainerFileCitation) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Default) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c FileCitation) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c FilePath) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c FileSearch) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Function) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c FunctionCallOutput) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c InputImage) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c InputText) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Json) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c List) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Llm) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c LoRa) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Mcp) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c McpCall) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c McpListTools) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Model) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Qat) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Response) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Rrf) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c TextCompletion) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c URLCitation) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c WebSearchCall) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Weighted) MarshalJSON() ([]byte, error)              { return marshalString(c) }

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

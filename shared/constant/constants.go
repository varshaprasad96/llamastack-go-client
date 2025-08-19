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
type AgentTurnInput string        // Always "agent_turn_input"
type Array string                 // Always "array"
type Assistant string             // Always "assistant"
type Auto string                  // Always "auto"
type Benchmark string             // Always "benchmark"
type Boolean string               // Always "boolean"
type ChatCompletionInput string   // Always "chat_completion_input"
type ChatCompletion string        // Always "chat.completion"
type ChatCompletionChunk string   // Always "chat.completion.chunk"
type CompletionInput string       // Always "completion_input"
type ContainerFileCitation string // Always "container_file_citation"
type Dataset string               // Always "dataset"
type Default string               // Always "default"
type Developer string             // Always "developer"
type Embedding string             // Always "embedding"
type File string                  // Always "file"
type FileCitation string          // Always "file_citation"
type FilePath string              // Always "file_path"
type FileSearch string            // Always "file_search"
type FileSearchCall string        // Always "file_search_call"
type Function string              // Always "function"
type FunctionCall string          // Always "function_call"
type FunctionCallOutput string    // Always "function_call_output"
type ImageURL string              // Always "image_url"
type Inference string             // Always "inference"
type InputImage string            // Always "input_image"
type InputText string             // Always "input_text"
type Json string                  // Always "json"
type JsonObject string            // Always "json_object"
type JsonSchema string            // Always "json_schema"
type List string                  // Always "list"
type Llm string                   // Always "llm"
type LoRa string                  // Always "LoRA"
type Mcp string                   // Always "mcp"
type McpCall string               // Always "mcp_call"
type McpListTools string          // Always "mcp_list_tools"
type MemoryRetrieval string       // Always "memory_retrieval"
type Message string               // Always "message"
type Model string                 // Always "model"
type Number string                // Always "number"
type Object string                // Always "object"
type OutputText string            // Always "output_text"
type Qat string                   // Always "QAT"
type Response string              // Always "response"
type Rows string                  // Always "rows"
type Rrf string                   // Always "rrf"
type ScoringFunction string       // Always "scoring_function"
type Shield string                // Always "shield"
type ShieldCall string            // Always "shield_call"
type Static string                // Always "static"
type String string                // Always "string"
type System string                // Always "system"
type Text string                  // Always "text"
type TextCompletion string        // Always "text_completion"
type Tool string                  // Always "tool"
type ToolExecution string         // Always "tool_execution"
type ToolGroup string             // Always "tool_group"
type Union string                 // Always "union"
type Uri string                   // Always "uri"
type URLCitation string           // Always "url_citation"
type User string                  // Always "user"
type VectorDB string              // Always "vector_db"
type WebSearchCall string         // Always "web_search_call"
type Weighted string              // Always "weighted"

func (c Agent) Default() Agent                                 { return "agent" }
func (c AgentTurnInput) Default() AgentTurnInput               { return "agent_turn_input" }
func (c Array) Default() Array                                 { return "array" }
func (c Assistant) Default() Assistant                         { return "assistant" }
func (c Auto) Default() Auto                                   { return "auto" }
func (c Benchmark) Default() Benchmark                         { return "benchmark" }
func (c Boolean) Default() Boolean                             { return "boolean" }
func (c ChatCompletionInput) Default() ChatCompletionInput     { return "chat_completion_input" }
func (c ChatCompletion) Default() ChatCompletion               { return "chat.completion" }
func (c ChatCompletionChunk) Default() ChatCompletionChunk     { return "chat.completion.chunk" }
func (c CompletionInput) Default() CompletionInput             { return "completion_input" }
func (c ContainerFileCitation) Default() ContainerFileCitation { return "container_file_citation" }
func (c Dataset) Default() Dataset                             { return "dataset" }
func (c Default) Default() Default                             { return "default" }
func (c Developer) Default() Developer                         { return "developer" }
func (c Embedding) Default() Embedding                         { return "embedding" }
func (c File) Default() File                                   { return "file" }
func (c FileCitation) Default() FileCitation                   { return "file_citation" }
func (c FilePath) Default() FilePath                           { return "file_path" }
func (c FileSearch) Default() FileSearch                       { return "file_search" }
func (c FileSearchCall) Default() FileSearchCall               { return "file_search_call" }
func (c Function) Default() Function                           { return "function" }
func (c FunctionCall) Default() FunctionCall                   { return "function_call" }
func (c FunctionCallOutput) Default() FunctionCallOutput       { return "function_call_output" }
func (c ImageURL) Default() ImageURL                           { return "image_url" }
func (c Inference) Default() Inference                         { return "inference" }
func (c InputImage) Default() InputImage                       { return "input_image" }
func (c InputText) Default() InputText                         { return "input_text" }
func (c Json) Default() Json                                   { return "json" }
func (c JsonObject) Default() JsonObject                       { return "json_object" }
func (c JsonSchema) Default() JsonSchema                       { return "json_schema" }
func (c List) Default() List                                   { return "list" }
func (c Llm) Default() Llm                                     { return "llm" }
func (c LoRa) Default() LoRa                                   { return "LoRA" }
func (c Mcp) Default() Mcp                                     { return "mcp" }
func (c McpCall) Default() McpCall                             { return "mcp_call" }
func (c McpListTools) Default() McpListTools                   { return "mcp_list_tools" }
func (c MemoryRetrieval) Default() MemoryRetrieval             { return "memory_retrieval" }
func (c Message) Default() Message                             { return "message" }
func (c Model) Default() Model                                 { return "model" }
func (c Number) Default() Number                               { return "number" }
func (c Object) Default() Object                               { return "object" }
func (c OutputText) Default() OutputText                       { return "output_text" }
func (c Qat) Default() Qat                                     { return "QAT" }
func (c Response) Default() Response                           { return "response" }
func (c Rows) Default() Rows                                   { return "rows" }
func (c Rrf) Default() Rrf                                     { return "rrf" }
func (c ScoringFunction) Default() ScoringFunction             { return "scoring_function" }
func (c Shield) Default() Shield                               { return "shield" }
func (c ShieldCall) Default() ShieldCall                       { return "shield_call" }
func (c Static) Default() Static                               { return "static" }
func (c String) Default() String                               { return "string" }
func (c System) Default() System                               { return "system" }
func (c Text) Default() Text                                   { return "text" }
func (c TextCompletion) Default() TextCompletion               { return "text_completion" }
func (c Tool) Default() Tool                                   { return "tool" }
func (c ToolExecution) Default() ToolExecution                 { return "tool_execution" }
func (c ToolGroup) Default() ToolGroup                         { return "tool_group" }
func (c Union) Default() Union                                 { return "union" }
func (c Uri) Default() Uri                                     { return "uri" }
func (c URLCitation) Default() URLCitation                     { return "url_citation" }
func (c User) Default() User                                   { return "user" }
func (c VectorDB) Default() VectorDB                           { return "vector_db" }
func (c WebSearchCall) Default() WebSearchCall                 { return "web_search_call" }
func (c Weighted) Default() Weighted                           { return "weighted" }

func (c Agent) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c AgentTurnInput) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Array) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Assistant) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Auto) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Benchmark) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Boolean) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c ChatCompletionInput) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c ChatCompletion) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ChatCompletionChunk) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c CompletionInput) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ContainerFileCitation) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Dataset) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Default) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Developer) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Embedding) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c FileCitation) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c FilePath) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c FileSearch) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c FileSearchCall) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Function) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c FunctionCall) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c FunctionCallOutput) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c ImageURL) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Inference) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c InputImage) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c InputText) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Json) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c JsonObject) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c JsonSchema) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c List) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Llm) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c LoRa) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Mcp) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c McpCall) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c McpListTools) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c MemoryRetrieval) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Message) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Model) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Number) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Object) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c OutputText) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Qat) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Response) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Rows) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Rrf) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c ScoringFunction) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Shield) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ShieldCall) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Static) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c String) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c System) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c TextCompletion) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Tool) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c ToolExecution) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c ToolGroup) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Union) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Uri) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c URLCitation) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c User) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c VectorDB) MarshalJSON() ([]byte, error)              { return marshalString(c) }
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

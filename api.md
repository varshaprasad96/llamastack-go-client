# Types

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#CompletionResponse">CompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TokenLogProbs">TokenLogProbs</a>

# Datasetio

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetioIterateRowsResponse">DatasetioIterateRowsResponse</a>

Methods:

- <code title="post /v1/datasetio/append-rows/{dataset_id}">client.Datasetio.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetioService.AppendRows">AppendRows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetioAppendRowsParams">DatasetioAppendRowsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/datasetio/iterrows/{dataset_id}">client.Datasetio.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetioService.IterateRows">IterateRows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetioIterateRowsParams">DatasetioIterateRowsParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetioIterateRowsResponse">DatasetioIterateRowsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Inference

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#MessageParam">MessageParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolCallParam">ToolCallParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ChatCompletionResponse">ChatCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#CompletionMessage">CompletionMessage</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#MetricInResponse">MetricInResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolCall">ToolCall</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceBatchChatCompletionResponse">InferenceBatchChatCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceBatchCompletionResponse">InferenceBatchCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceEmbeddingsResponse">InferenceEmbeddingsResponse</a>

Methods:

- <code title="post /v1/inference/batch-chat-completion">client.Inference.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceService.BatchChatCompletion">BatchChatCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceBatchChatCompletionParams">InferenceBatchChatCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceBatchChatCompletionResponse">InferenceBatchChatCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/inference/batch-completion">client.Inference.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceService.BatchCompletion">BatchCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceBatchCompletionParams">InferenceBatchCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceBatchCompletionResponse">InferenceBatchCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/inference/chat-completion">client.Inference.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceService.ChatCompletion">ChatCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceChatCompletionParams">InferenceChatCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ChatCompletionResponse">ChatCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/inference/completion">client.Inference.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceService.Completion">Completion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceCompletionParams">InferenceCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#CompletionResponse">CompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/inference/embeddings">client.Inference.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceService.Embeddings">Embeddings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceEmbeddingsParams">InferenceEmbeddingsParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceEmbeddingsResponse">InferenceEmbeddingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PostTraining

## Job

Methods:

- <code title="post /v1/post-training/job/cancel">client.PostTraining.Job.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobCancelParams">PostTrainingJobCancelParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Agents

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewResponse">AgentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentListResponse">AgentListResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewSessionResponse">AgentNewSessionResponse</a>

Methods:

- <code title="post /v1/agents">client.Agents.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewParams">AgentNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewResponse">AgentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents">client.Agents.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentListParams">AgentListParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentListResponse">AgentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/agents/{agent_id}/session">client.Agents.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentService.NewSession">NewSession</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewSessionParams">AgentNewSessionParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewSessionResponse">AgentNewSessionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Session

### Turn

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnNewResponse">AgentSessionTurnNewResponse</a>

Methods:

- <code title="post /v1/agents/{agent_id}/session/{session_id}/turn">client.Agents.Session.Turn.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sessionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnNewParams">AgentSessionTurnNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnNewResponse">AgentSessionTurnNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Step

# OpenAI

## V1

### Responses

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ResponseListResponse">OpenAiv1ResponseListResponse</a>

Methods:

- <code title="get /v1/openai/v1/responses">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ResponseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ResponseListResponse">OpenAiv1ResponseListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Chat

#### Completions

### VectorStores

#### Files

### Files

# Eval

## Benchmarks

### Jobs

# Datasets

# Models

# ScoringFunctions

# Shields

# Telemetry

## Traces

## Spans

# Tools

# Toolgroups

# VectorDBs

# Health

# ToolRuntime

## RagTool

# VectorIo

# Providers

# Inspect

# Safety

# Scoring

# SyntheticDataGeneration

# Version

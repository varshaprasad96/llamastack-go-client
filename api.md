# Shared

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolDefParam">ToolDefParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#CompletionResponse">CompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TokenLogProbs">TokenLogProbs</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolDef">ToolDef</a>

# Datasetio

Methods:

- <code title="post /v1/datasetio/append-rows/{dataset_id}">client.Datasetio.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetioService.AppendRows">AppendRows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetioAppendRowsParams">DatasetioAppendRowsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/datasetio/iterrows/{dataset_id}">client.Datasetio.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetioService.IterateRows">IterateRows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetioIterateRowsParams">DatasetioIterateRowsParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PaginatedResponse">PaginatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Inference

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#CompletionMessageParam">CompletionMessageParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InterleavedContentUnionParam">InterleavedContentUnionParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InterleavedContentItemUnionParam">InterleavedContentItemUnionParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#MessageUnionParam">MessageUnionParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseFormatUnionParam">ResponseFormatUnionParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SamplingParams">SamplingParams</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SystemMessageParam">SystemMessageParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolCallParam">ToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolConfigParam">ToolConfigParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolDefinitionParam">ToolDefinitionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ChatCompletionResponse">ChatCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#CompletionMessage">CompletionMessage</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InterleavedContentUnion">InterleavedContentUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InterleavedContentItemUnion">InterleavedContentItemUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#MetricInResponse">MetricInResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseFormatUnion">ResponseFormatUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SamplingParamsResp">SamplingParamsResp</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolCall">ToolCall</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolConfig">ToolConfig</a>
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

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TrainingConfigParam">TrainingConfigParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJob">PostTrainingJob</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingListJobsResponse">PostTrainingListJobsResponse</a>

Methods:

- <code title="post /v1/post-training/supervised-fine-tune">client.PostTraining.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingService.FineTuneSupervised">FineTuneSupervised</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingFineTuneSupervisedParams">PostTrainingFineTuneSupervisedParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJob">PostTrainingJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/post-training/jobs">client.PostTraining.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingService.ListJobs">ListJobs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingListJobsResponse">PostTrainingListJobsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/post-training/preference-optimize">client.PostTraining.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingService.OptimizePreferences">OptimizePreferences</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingOptimizePreferencesParams">PostTrainingOptimizePreferencesParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJob">PostTrainingJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Job

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Checkpoint">Checkpoint</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobGetArtifactsResponse">PostTrainingJobGetArtifactsResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobGetStatusResponse">PostTrainingJobGetStatusResponse</a>

Methods:

- <code title="post /v1/post-training/job/cancel">client.PostTraining.Job.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobCancelParams">PostTrainingJobCancelParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/post-training/job/artifacts">client.PostTraining.Job.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobService.GetArtifacts">GetArtifacts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobGetArtifactsParams">PostTrainingJobGetArtifactsParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobGetArtifactsResponse">PostTrainingJobGetArtifactsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/post-training/job/status">client.PostTraining.Job.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobGetStatusParams">PostTrainingJobGetStatusParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PostTrainingJobGetStatusResponse">PostTrainingJobGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Agents

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentConfigParam">AgentConfigParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentConfig">AgentConfig</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PaginatedResponse">PaginatedResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewResponse">AgentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentGetResponse">AgentGetResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewSessionResponse">AgentNewSessionResponse</a>

Methods:

- <code title="post /v1/agents">client.Agents.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewParams">AgentNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewResponse">AgentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agent_id}">client.Agents.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentGetResponse">AgentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents">client.Agents.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentListParams">AgentListParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PaginatedResponse">PaginatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/agents/{agent_id}">client.Agents.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/agents/{agent_id}/session">client.Agents.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentService.NewSession">NewSession</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewSessionParams">AgentNewSessionParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentNewSessionResponse">AgentNewSessionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agent_id}/sessions">client.Agents.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentService.GetSessions">GetSessions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentGetSessionsParams">AgentGetSessionsParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#PaginatedResponse">PaginatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Session

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionGetResponse">AgentSessionGetResponse</a>

Methods:

- <code title="get /v1/agents/{agent_id}/session/{session_id}">client.Agents.Session.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sessionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionGetParams">AgentSessionGetParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionGetResponse">AgentSessionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/agents/{agent_id}/session/{session_id}">client.Agents.Session.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sessionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionDeleteParams">AgentSessionDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Turn

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentToolUnionParam">AgentToolUnionParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolResponseParam">ToolResponseParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolResponseMessageParam">ToolResponseMessageParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#UserMessageParam">UserMessageParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentToolUnion">AgentToolUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InferenceStep">InferenceStep</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#MemoryRetrievalStep">MemoryRetrievalStep</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ShieldCallStep">ShieldCallStep</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolExecutionStep">ToolExecutionStep</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolResponse">ToolResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolResponseMessage">ToolResponseMessage</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Turn">Turn</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#UserMessage">UserMessage</a>

Methods:

- <code title="post /v1/agents/{agent_id}/session/{session_id}/turn">client.Agents.Session.Turn.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sessionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnNewParams">AgentSessionTurnNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Turn">Turn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agent_id}/session/{session_id}/turn/{turn_id}">client.Agents.Session.Turn.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, turnID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnGetParams">AgentSessionTurnGetParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Turn">Turn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/agents/{agent_id}/session/{session_id}/turn/{turn_id}/resume">client.Agents.Session.Turn.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnService.Resume">Resume</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, turnID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnResumeParams">AgentSessionTurnResumeParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Turn">Turn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Step

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnStepGetResponse">AgentSessionTurnStepGetResponse</a>

Methods:

- <code title="get /v1/agents/{agent_id}/session/{session_id}/turn/{turn_id}/step/{step_id}">client.Agents.Session.Turn.Step.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnStepService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnStepGetParams">AgentSessionTurnStepGetParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AgentSessionTurnStepGetResponse">AgentSessionTurnStepGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OpenAI

## V1

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1CompletionsResponse">OpenAiv1CompletionsResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1EmbeddingsResponse">OpenAiv1EmbeddingsResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ModerationsResponse">OpenAiv1ModerationsResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1GetModelsResponse">OpenAiv1GetModelsResponse</a>

Methods:

- <code title="post /v1/openai/v1/completions">client.OpenAI.V1.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1Service.Completions">Completions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1CompletionsParams">OpenAIV1CompletionsParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1CompletionsResponse">OpenAiv1CompletionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/openai/v1/embeddings">client.OpenAI.V1.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1Service.Embeddings">Embeddings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1EmbeddingsParams">OpenAIV1EmbeddingsParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1EmbeddingsResponse">OpenAiv1EmbeddingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/openai/v1/moderations">client.OpenAI.V1.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1Service.Moderations">Moderations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ModerationsParams">OpenAIV1ModerationsParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ModerationsResponse">OpenAiv1ModerationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/models">client.OpenAI.V1.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1Service.GetModels">GetModels</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1GetModelsResponse">OpenAiv1GetModelsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Responses

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OutputMessageFileSearchToolCallParam">OutputMessageFileSearchToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OutputMessageFunctionToolCallParam">OutputMessageFunctionToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OutputMessageWebSearchToolCallParam">OutputMessageWebSearchToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseInputUnionParam">ResponseInputUnionParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseMessageParam">ResponseMessageParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseTextParam">ResponseTextParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OutputUnion">OutputUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OutputMessageFileSearchToolCall">OutputMessageFileSearchToolCall</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OutputMessageFunctionToolCall">OutputMessageFunctionToolCall</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OutputMessageWebSearchToolCall">OutputMessageWebSearchToolCall</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseError">ResponseError</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseInputUnion">ResponseInputUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseMessage">ResponseMessage</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseObject">ResponseObject</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseText">ResponseText</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ResponseListResponse">OpenAiv1ResponseListResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ResponseDeleteResponse">OpenAiv1ResponseDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ResponseGetInputItemsResponse">OpenAiv1ResponseGetInputItemsResponse</a>

Methods:

- <code title="post /v1/openai/v1/responses">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ResponseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ResponseNewParams">OpenAIV1ResponseNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseObject">ResponseObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/responses/{response_id}">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ResponseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ResponseObject">ResponseObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/responses">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ResponseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ResponseListParams">OpenAIV1ResponseListParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ResponseListResponse">OpenAiv1ResponseListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/openai/v1/responses/{response_id}">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ResponseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ResponseDeleteResponse">OpenAiv1ResponseDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/responses/{response_id}/input_items">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ResponseService.GetInputItems">GetInputItems</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ResponseGetInputItemsParams">OpenAIV1ResponseGetInputItemsParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ResponseGetInputItemsResponse">OpenAiv1ResponseGetInputItemsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Chat

#### Completions

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ChatCompletionToolCallParam">ChatCompletionToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ContentPartTextParam">ContentPartTextParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#MessageParamUnion">MessageParamUnion</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ChatCompletionToolCall">ChatCompletionToolCall</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Choice">Choice</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ChoiceLogprobs">ChoiceLogprobs</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ContentPartTextParamResp">ContentPartTextParamResp</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#MessageParamUnionResp">MessageParamUnionResp</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TokenLogProb">TokenLogProb</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ChatCompletionNewResponseUnion">OpenAiv1ChatCompletionNewResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ChatCompletionGetResponse">OpenAiv1ChatCompletionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ChatCompletionListResponse">OpenAiv1ChatCompletionListResponse</a>

Methods:

- <code title="post /v1/openai/v1/chat/completions">client.OpenAI.V1.Chat.Completions.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ChatCompletionNewParams">OpenAIV1ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ChatCompletionNewResponseUnion">OpenAiv1ChatCompletionNewResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/chat/completions/{completion_id}">client.OpenAI.V1.Chat.Completions.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ChatCompletionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, completionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ChatCompletionGetResponse">OpenAiv1ChatCompletionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/chat/completions">client.OpenAI.V1.Chat.Completions.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ChatCompletionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1ChatCompletionListParams">OpenAIV1ChatCompletionListParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1ChatCompletionListResponse">OpenAiv1ChatCompletionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### VectorStores

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorStoreObject">VectorStoreObject</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreListResponse">OpenAiv1VectorStoreListResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreDeleteResponse">OpenAiv1VectorStoreDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreSearchResponse">OpenAiv1VectorStoreSearchResponse</a>

Methods:

- <code title="post /v1/openai/v1/vector_stores">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreNewParams">OpenAIV1VectorStoreNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorStoreObject">VectorStoreObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/vector_stores/{vector_store_id}">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorStoreObject">VectorStoreObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/openai/v1/vector_stores/{vector_store_id}">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreUpdateParams">OpenAIV1VectorStoreUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorStoreObject">VectorStoreObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/vector_stores">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreListParams">OpenAIV1VectorStoreListParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreListResponse">OpenAiv1VectorStoreListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/openai/v1/vector_stores/{vector_store_id}">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreDeleteResponse">OpenAiv1VectorStoreDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/openai/v1/vector_stores/{vector_store_id}/search">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreSearchParams">OpenAIV1VectorStoreSearchParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreSearchResponse">OpenAiv1VectorStoreSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Files

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ChunkingStrategyUnionParam">ChunkingStrategyUnionParam</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#FileStatus">FileStatus</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ChunkingStrategyUnion">ChunkingStrategyUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Content">Content</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#FileStatus">FileStatus</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorStoreFile">VectorStoreFile</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreFileListResponse">OpenAiv1VectorStoreFileListResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreFileDeleteResponse">OpenAiv1VectorStoreFileDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreFileGetContentResponse">OpenAiv1VectorStoreFileGetContentResponse</a>

Methods:

- <code title="post /v1/openai/v1/vector_stores/{vector_store_id}/files">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileNewParams">OpenAIV1VectorStoreFileNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/vector_stores/{vector_store_id}/files/{file_id}">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileGetParams">OpenAIV1VectorStoreFileGetParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/openai/v1/vector_stores/{vector_store_id}/files/{file_id}">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileUpdateParams">OpenAIV1VectorStoreFileUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/vector_stores/{vector_store_id}/files">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileListParams">OpenAIV1VectorStoreFileListParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreFileListResponse">OpenAiv1VectorStoreFileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/openai/v1/vector_stores/{vector_store_id}/files/{file_id}">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileDeleteParams">OpenAIV1VectorStoreFileDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreFileDeleteResponse">OpenAiv1VectorStoreFileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/vector_stores/{vector_store_id}/files/{file_id}/content">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileService.GetContent">GetContent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1VectorStoreFileGetContentParams">OpenAIV1VectorStoreFileGetContentParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1VectorStoreFileGetContentResponse">OpenAiv1VectorStoreFileGetContentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Files

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#FilePurpose">FilePurpose</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIFile">OpenAIFile</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1FileListResponse">OpenAiv1FileListResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1FileDeleteResponse">OpenAiv1FileDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1FileGetContentResponse">OpenAiv1FileGetContentResponse</a>

Methods:

- <code title="post /v1/openai/v1/files">client.OpenAI.V1.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1FileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1FileNewParams">OpenAIV1FileNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIFile">OpenAIFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/files/{file_id}">client.OpenAI.V1.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIFile">OpenAIFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/files">client.OpenAI.V1.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1FileListParams">OpenAIV1FileListParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1FileListResponse">OpenAiv1FileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/openai/v1/files/{file_id}">client.OpenAI.V1.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1FileDeleteResponse">OpenAiv1FileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/files/{file_id}/content">client.OpenAI.V1.Files.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAIV1FileService.GetContent">GetContent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#OpenAiv1FileGetContentResponse">OpenAiv1FileGetContentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Eval

## Benchmarks

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#BenchmarkConfigParam">BenchmarkConfigParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Benchmark">Benchmark</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvaluateResponse">EvaluateResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkListResponse">EvalBenchmarkListResponse</a>

Methods:

- <code title="post /v1/eval/benchmarks">client.Eval.Benchmarks.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkNewParams">EvalBenchmarkNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/eval/benchmarks/{benchmark_id}">client.Eval.Benchmarks.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Benchmark">Benchmark</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/eval/benchmarks">client.Eval.Benchmarks.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkListResponse">EvalBenchmarkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/eval/benchmarks/{benchmark_id}/evaluations">client.Eval.Benchmarks.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkService.Evaluate">Evaluate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkEvaluateParams">EvalBenchmarkEvaluateParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvaluateResponse">EvaluateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Job">Job</a>

Methods:

- <code title="get /v1/eval/benchmarks/{benchmark_id}/jobs/{job_id}">client.Eval.Benchmarks.Jobs.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkJobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkJobGetParams">EvalBenchmarkJobGetParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/eval/benchmarks/{benchmark_id}/jobs/{job_id}">client.Eval.Benchmarks.Jobs.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkJobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkJobCancelParams">EvalBenchmarkJobCancelParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/eval/benchmarks/{benchmark_id}/jobs/{job_id}/result">client.Eval.Benchmarks.Jobs.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkJobService.Result">Result</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkJobResultParams">EvalBenchmarkJobResultParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvaluateResponse">EvaluateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/eval/benchmarks/{benchmark_id}/jobs">client.Eval.Benchmarks.Jobs.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkJobService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EvalBenchmarkJobRunParams">EvalBenchmarkJobRunParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Datasets

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DataSourceUnionParam">DataSourceUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DataSourceUnion">DataSourceUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Dataset">Dataset</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetListResponse">DatasetListResponse</a>

Methods:

- <code title="post /v1/datasets">client.Datasets.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetNewParams">DatasetNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/datasets/{dataset_id}">client.Datasets.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/datasets">client.Datasets.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetListResponse">DatasetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/datasets/{dataset_id}">client.Datasets.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#DatasetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Models

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ModelType">ModelType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ModelType">ModelType</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="post /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ModelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ModelNewParams">ModelNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ModelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# ScoringFunctions

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AggregationFunctionType">AggregationFunctionType</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ParamTypeUnion">ParamTypeUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFnParamsUnion">ScoringFnParamsUnion</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFnParamsType">ScoringFnParamsType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#AggregationFunctionType">AggregationFunctionType</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ParamTypeUnionResp">ParamTypeUnionResp</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFn">ScoringFn</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFnParamsUnionResp">ScoringFnParamsUnionResp</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFnParamsType">ScoringFnParamsType</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFunctionListResponse">ScoringFunctionListResponse</a>

Methods:

- <code title="post /v1/scoring-functions">client.ScoringFunctions.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFunctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFunctionNewParams">ScoringFunctionNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/scoring-functions/{scoring_fn_id}">client.ScoringFunctions.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFunctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, scoringFnID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFn">ScoringFn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/scoring-functions">client.ScoringFunctions.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringFunctionListResponse">ScoringFunctionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Shields

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Shield">Shield</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ShieldListResponse">ShieldListResponse</a>

Methods:

- <code title="post /v1/shields">client.Shields.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ShieldService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ShieldNewParams">ShieldNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Shield">Shield</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/shields/{identifier}">client.Shields.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ShieldService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Shield">Shield</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/shields">client.Shields.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ShieldService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ShieldListResponse">ShieldListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/shields/{identifier}">client.Shields.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ShieldService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Telemetry

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#EventType">EventType</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#StructuredLogType">StructuredLogType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryQueryMetricResponse">TelemetryQueryMetricResponse</a>

Methods:

- <code title="post /v1/telemetry/events">client.Telemetry.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryService.LogEvent">LogEvent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryLogEventParams">TelemetryLogEventParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/telemetry/metrics/{metric_name}">client.Telemetry.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryService.QueryMetric">QueryMetric</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, metricName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryQueryMetricParams">TelemetryQueryMetricParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryQueryMetricResponse">TelemetryQueryMetricResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Traces

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Span">Span</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Trace">Trace</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryTraceQueryResponse">TelemetryTraceQueryResponse</a>

Methods:

- <code title="post /v1/telemetry/traces">client.Telemetry.Traces.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryTraceService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryTraceQueryParams">TelemetryTraceQueryParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryTraceQueryResponse">TelemetryTraceQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/telemetry/traces/{trace_id}/spans/{span_id}">client.Telemetry.Traces.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryTraceService.GetSpan">GetSpan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, spanID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryTraceGetSpanParams">TelemetryTraceGetSpanParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Span">Span</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/telemetry/traces/{trace_id}">client.Telemetry.Traces.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetryTraceService.GetTrace">GetTrace</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, traceID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Trace">Trace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Spans

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#QueryConditionParam">QueryConditionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetrySpanQueryResponse">TelemetrySpanQueryResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetrySpanGetTreeResponse">TelemetrySpanGetTreeResponse</a>

Methods:

- <code title="post /v1/telemetry/spans/export">client.Telemetry.Spans.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetrySpanService.Export">Export</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetrySpanExportParams">TelemetrySpanExportParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/telemetry/spans">client.Telemetry.Spans.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetrySpanService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetrySpanQueryParams">TelemetrySpanQueryParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetrySpanQueryResponse">TelemetrySpanQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/telemetry/spans/{span_id}/tree">client.Telemetry.Spans.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetrySpanService.GetTree">GetTree</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, spanID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetrySpanGetTreeParams">TelemetrySpanGetTreeParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#TelemetrySpanGetTreeResponse">TelemetrySpanGetTreeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tools

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolParameter">ToolParameter</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Tool">Tool</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolParameterResp">ToolParameterResp</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolListResponse">ToolListResponse</a>

Methods:

- <code title="get /v1/tools/{tool_name}">client.Tools.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Tool">Tool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tools">client.Tools.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolListParams">ToolListParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolListResponse">ToolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Toolgroups

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolGroup">ToolGroup</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolgroupListResponse">ToolgroupListResponse</a>

Methods:

- <code title="post /v1/toolgroups">client.Toolgroups.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolgroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolgroupNewParams">ToolgroupNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/toolgroups/{toolgroup_id}">client.Toolgroups.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolgroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolgroupID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolGroup">ToolGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/toolgroups">client.Toolgroups.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolgroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolgroupListResponse">ToolgroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/toolgroups/{toolgroup_id}">client.Toolgroups.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolgroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolgroupID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# VectorDBs

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorDB">VectorDB</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorDBListResponse">VectorDBListResponse</a>

Methods:

- <code title="post /v1/vector-dbs">client.VectorDBs.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorDBService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorDBNewParams">VectorDBNewParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorDB">VectorDB</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector-dbs/{vector_db_id}">client.VectorDBs.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorDBService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorDBID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorDB">VectorDB</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector-dbs">client.VectorDBs.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorDBService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorDBListResponse">VectorDBListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/vector-dbs/{vector_db_id}">client.VectorDBs.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorDBService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorDBID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#HealthGetResponse">HealthGetResponse</a>

Methods:

- <code title="get /v1/health">client.Health.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#HealthService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#HealthGetResponse">HealthGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ToolRuntime

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#URLParam">URLParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#URL">URL</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeInvokeResponse">ToolRuntimeInvokeResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeListToolsResponse">ToolRuntimeListToolsResponse</a>

Methods:

- <code title="post /v1/tool-runtime/invoke">client.ToolRuntime.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeInvokeParams">ToolRuntimeInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeInvokeResponse">ToolRuntimeInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tool-runtime/list-tools">client.ToolRuntime.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeService.ListTools">ListTools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeListToolsParams">ToolRuntimeListToolsParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeListToolsResponse">ToolRuntimeListToolsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RagTool

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeRagToolQueryResponse">ToolRuntimeRagToolQueryResponse</a>

Methods:

- <code title="post /v1/tool-runtime/rag-tool/insert">client.ToolRuntime.RagTool.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeRagToolService.Insert">Insert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeRagToolInsertParams">ToolRuntimeRagToolInsertParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/tool-runtime/rag-tool/query">client.ToolRuntime.RagTool.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeRagToolService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeRagToolQueryParams">ToolRuntimeRagToolQueryParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ToolRuntimeRagToolQueryResponse">ToolRuntimeRagToolQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VectorIo

Params Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ChunkParam">ChunkParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#Chunk">Chunk</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorIoQueryResponse">VectorIoQueryResponse</a>

Methods:

- <code title="post /v1/vector-io/insert">client.VectorIo.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorIoService.Insert">Insert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorIoInsertParams">VectorIoInsertParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/vector-io/query">client.VectorIo.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorIoService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorIoQueryParams">VectorIoQueryParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VectorIoQueryResponse">VectorIoQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Providers

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ProviderInfo">ProviderInfo</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ProviderListResponse">ProviderListResponse</a>

Methods:

- <code title="get /v1/providers/{provider_id}">client.Providers.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ProviderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, providerID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ProviderInfo">ProviderInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/providers">client.Providers.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ProviderListResponse">ProviderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Inspect

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InspectListRoutesResponse">InspectListRoutesResponse</a>

Methods:

- <code title="get /v1/inspect/routes">client.Inspect.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InspectService.ListRoutes">ListRoutes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#InspectListRoutesResponse">InspectListRoutesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Safety

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SafetyViolation">SafetyViolation</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SafetyRunShieldResponse">SafetyRunShieldResponse</a>

Methods:

- <code title="post /v1/safety/run-shield">client.Safety.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SafetyService.RunShield">RunShield</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SafetyRunShieldParams">SafetyRunShieldParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SafetyRunShieldResponse">SafetyRunShieldResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Scoring

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringScoreResponse">ScoringScoreResponse</a>
- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringScoreBatchResponse">ScoringScoreBatchResponse</a>

Methods:

- <code title="post /v1/scoring/score">client.Scoring.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringService.Score">Score</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringScoreParams">ScoringScoreParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringScoreResponse">ScoringScoreResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/scoring/score-batch">client.Scoring.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringService.ScoreBatch">ScoreBatch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringScoreBatchParams">ScoringScoreBatchParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#ScoringScoreBatchResponse">ScoringScoreBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SyntheticDataGeneration

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SyntheticDataGenerationGenerateResponse">SyntheticDataGenerationGenerateResponse</a>

Methods:

- <code title="post /v1/synthetic-data-generation/generate">client.SyntheticDataGeneration.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SyntheticDataGenerationService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SyntheticDataGenerationGenerateParams">SyntheticDataGenerationGenerateParams</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#SyntheticDataGenerationGenerateResponse">SyntheticDataGenerationGenerateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Version

Response Types:

- <a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VersionGetResponse">VersionGetResponse</a>

Methods:

- <code title="get /v1/version">client.Version.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/varshaprasad96/llamastack-go-client#VersionGetResponse">VersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

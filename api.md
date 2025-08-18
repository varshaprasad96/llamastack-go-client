# Datasetio

Methods:

- <code title="post /v1/datasetio/append-rows/{dataset_id}">client.Datasetio.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetioService.AppendRows">AppendRows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetioAppendRowsParams">DatasetioAppendRowsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/datasetio/iterrows/{dataset_id}">client.Datasetio.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetioService.IterateRows">IterateRows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetioIterateRowsParams">DatasetioIterateRowsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PaginatedResponse">PaginatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Inference

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#CompletionMessageParam">CompletionMessageParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InterleavedContentUnionParam">InterleavedContentUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InterleavedContentItemUnionParam">InterleavedContentItemUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#MessageUnionParam">MessageUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseFormatUnionParam">ResponseFormatUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SamplingParams">SamplingParams</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SystemMessageParam">SystemMessageParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolCallParam">ToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolConfigParam">ToolConfigParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolDefinitionParam">ToolDefinitionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ChatCompletionResponse">ChatCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#CompletionMessage">CompletionMessage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#CompletionResponse">CompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InterleavedContentUnion">InterleavedContentUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InterleavedContentItemUnion">InterleavedContentItemUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#MetricInResponse">MetricInResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseFormatUnion">ResponseFormatUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SamplingParamsResp">SamplingParamsResp</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TokenLogProbs">TokenLogProbs</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolCall">ToolCall</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolConfig">ToolConfig</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceBatchChatCompletionResponse">InferenceBatchChatCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceBatchCompletionResponse">InferenceBatchCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceEmbeddingsResponse">InferenceEmbeddingsResponse</a>

Methods:

- <code title="post /v1/inference/batch-chat-completion">client.Inference.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceService.BatchChatCompletion">BatchChatCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceBatchChatCompletionParams">InferenceBatchChatCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceBatchChatCompletionResponse">InferenceBatchChatCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/inference/batch-completion">client.Inference.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceService.BatchCompletion">BatchCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceBatchCompletionParams">InferenceBatchCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceBatchCompletionResponse">InferenceBatchCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/inference/chat-completion">client.Inference.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceService.ChatCompletion">ChatCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceChatCompletionParams">InferenceChatCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ChatCompletionResponse">ChatCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/inference/completion">client.Inference.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceService.Completion">Completion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceCompletionParams">InferenceCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#CompletionResponse">CompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/inference/embeddings">client.Inference.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceService.Embeddings">Embeddings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceEmbeddingsParams">InferenceEmbeddingsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceEmbeddingsResponse">InferenceEmbeddingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PostTraining

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TrainingConfigParam">TrainingConfigParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJob">PostTrainingJob</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingListJobsResponse">PostTrainingListJobsResponse</a>

Methods:

- <code title="post /v1/post-training/supervised-fine-tune">client.PostTraining.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingService.FineTuneSupervised">FineTuneSupervised</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingFineTuneSupervisedParams">PostTrainingFineTuneSupervisedParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJob">PostTrainingJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/post-training/jobs">client.PostTraining.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingService.ListJobs">ListJobs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingListJobsResponse">PostTrainingListJobsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/post-training/preference-optimize">client.PostTraining.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingService.OptimizePreferences">OptimizePreferences</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingOptimizePreferencesParams">PostTrainingOptimizePreferencesParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJob">PostTrainingJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Job

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Checkpoint">Checkpoint</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJobGetArtifactsResponse">PostTrainingJobGetArtifactsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJobGetStatusResponse">PostTrainingJobGetStatusResponse</a>

Methods:

- <code title="post /v1/post-training/job/cancel">client.PostTraining.Job.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJobCancelParams">PostTrainingJobCancelParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/post-training/job/artifacts">client.PostTraining.Job.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJobService.GetArtifacts">GetArtifacts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJobGetArtifactsParams">PostTrainingJobGetArtifactsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJobGetArtifactsResponse">PostTrainingJobGetArtifactsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/post-training/job/status">client.PostTraining.Job.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJobService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJobGetStatusParams">PostTrainingJobGetStatusParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PostTrainingJobGetStatusResponse">PostTrainingJobGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Agents

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentConfigParam">AgentConfigParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentConfig">AgentConfig</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PaginatedResponse">PaginatedResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentNewResponse">AgentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentGetResponse">AgentGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentNewSessionResponse">AgentNewSessionResponse</a>

Methods:

- <code title="post /v1/agents">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentNewParams">AgentNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentNewResponse">AgentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agent_id}">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentGetResponse">AgentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentListParams">AgentListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PaginatedResponse">PaginatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/agents/{agent_id}">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/agents/{agent_id}/session">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentService.NewSession">NewSession</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentNewSessionParams">AgentNewSessionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentNewSessionResponse">AgentNewSessionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agent_id}/sessions">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentService.GetSessions">GetSessions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentGetSessionsParams">AgentGetSessionsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#PaginatedResponse">PaginatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Session

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionGetResponse">AgentSessionGetResponse</a>

Methods:

- <code title="get /v1/agents/{agent_id}/session/{session_id}">client.Agents.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sessionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionGetParams">AgentSessionGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionGetResponse">AgentSessionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/agents/{agent_id}/session/{session_id}">client.Agents.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sessionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionDeleteParams">AgentSessionDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Turn

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentToolUnionParam">AgentToolUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolResponseParam">ToolResponseParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolResponseMessageParam">ToolResponseMessageParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#UserMessageParam">UserMessageParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentToolUnion">AgentToolUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InferenceStep">InferenceStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#MemoryRetrievalStep">MemoryRetrievalStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ShieldCallStep">ShieldCallStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolExecutionStep">ToolExecutionStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolResponse">ToolResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolResponseMessage">ToolResponseMessage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Turn">Turn</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#UserMessage">UserMessage</a>

Methods:

- <code title="post /v1/agents/{agent_id}/session/{session_id}/turn">client.Agents.Session.Turn.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionTurnService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sessionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionTurnNewParams">AgentSessionTurnNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Turn">Turn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{agent_id}/session/{session_id}/turn/{turn_id}">client.Agents.Session.Turn.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionTurnService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, turnID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionTurnGetParams">AgentSessionTurnGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Turn">Turn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/agents/{agent_id}/session/{session_id}/turn/{turn_id}/resume">client.Agents.Session.Turn.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionTurnService.Resume">Resume</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, turnID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionTurnResumeParams">AgentSessionTurnResumeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Turn">Turn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Step

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionTurnStepGetResponse">AgentSessionTurnStepGetResponse</a>

Methods:

- <code title="get /v1/agents/{agent_id}/session/{session_id}/turn/{turn_id}/step/{step_id}">client.Agents.Session.Turn.Step.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionTurnStepService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionTurnStepGetParams">AgentSessionTurnStepGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AgentSessionTurnStepGetResponse">AgentSessionTurnStepGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OpenAI

## V1

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1CompletionsResponse">OpenAiv1CompletionsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1EmbeddingsResponse">OpenAiv1EmbeddingsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ModerationsResponse">OpenAiv1ModerationsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1GetModelsResponse">OpenAiv1GetModelsResponse</a>

Methods:

- <code title="post /v1/openai/v1/completions">client.OpenAI.V1.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1Service.Completions">Completions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1CompletionsParams">OpenAIV1CompletionsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1CompletionsResponse">OpenAiv1CompletionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/openai/v1/embeddings">client.OpenAI.V1.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1Service.Embeddings">Embeddings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1EmbeddingsParams">OpenAIV1EmbeddingsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1EmbeddingsResponse">OpenAiv1EmbeddingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/openai/v1/moderations">client.OpenAI.V1.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1Service.Moderations">Moderations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ModerationsParams">OpenAIV1ModerationsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ModerationsResponse">OpenAiv1ModerationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/models">client.OpenAI.V1.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1Service.GetModels">GetModels</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1GetModelsResponse">OpenAiv1GetModelsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Responses

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OutputMessageFileSearchToolCallParam">OutputMessageFileSearchToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OutputMessageFunctionToolCallParam">OutputMessageFunctionToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OutputMessageWebSearchToolCallParam">OutputMessageWebSearchToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseInputUnionParam">ResponseInputUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseMessageParam">ResponseMessageParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseTextParam">ResponseTextParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OutputUnion">OutputUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OutputMessageFileSearchToolCall">OutputMessageFileSearchToolCall</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OutputMessageFunctionToolCall">OutputMessageFunctionToolCall</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OutputMessageWebSearchToolCall">OutputMessageWebSearchToolCall</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseError">ResponseError</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseInputUnion">ResponseInputUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseMessage">ResponseMessage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseObject">ResponseObject</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseText">ResponseText</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ResponseListResponse">OpenAiv1ResponseListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ResponseDeleteResponse">OpenAiv1ResponseDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ResponseGetInputItemsResponse">OpenAiv1ResponseGetInputItemsResponse</a>

Methods:

- <code title="post /v1/openai/v1/responses">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ResponseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ResponseNewParams">OpenAIV1ResponseNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseObject">ResponseObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/responses/{response_id}">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ResponseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ResponseObject">ResponseObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/responses">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ResponseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ResponseListParams">OpenAIV1ResponseListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ResponseListResponse">OpenAiv1ResponseListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/openai/v1/responses/{response_id}">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ResponseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ResponseDeleteResponse">OpenAiv1ResponseDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/responses/{response_id}/input_items">client.OpenAI.V1.Responses.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ResponseService.GetInputItems">GetInputItems</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ResponseGetInputItemsParams">OpenAIV1ResponseGetInputItemsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ResponseGetInputItemsResponse">OpenAiv1ResponseGetInputItemsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Chat

#### Completions

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ChatCompletionToolCallParam">ChatCompletionToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ContentPartTextParam">ContentPartTextParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#MessageParamUnion">MessageParamUnion</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ChatCompletionToolCall">ChatCompletionToolCall</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Choice">Choice</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ChoiceLogprobs">ChoiceLogprobs</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ContentPartTextParamResp">ContentPartTextParamResp</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#MessageParamUnionResp">MessageParamUnionResp</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TokenLogProb">TokenLogProb</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ChatCompletionNewResponseUnion">OpenAiv1ChatCompletionNewResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ChatCompletionGetResponse">OpenAiv1ChatCompletionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ChatCompletionListResponse">OpenAiv1ChatCompletionListResponse</a>

Methods:

- <code title="post /v1/openai/v1/chat/completions">client.OpenAI.V1.Chat.Completions.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ChatCompletionNewParams">OpenAIV1ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ChatCompletionNewResponseUnion">OpenAiv1ChatCompletionNewResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/chat/completions/{completion_id}">client.OpenAI.V1.Chat.Completions.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ChatCompletionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, completionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ChatCompletionGetResponse">OpenAiv1ChatCompletionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/chat/completions">client.OpenAI.V1.Chat.Completions.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ChatCompletionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1ChatCompletionListParams">OpenAIV1ChatCompletionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1ChatCompletionListResponse">OpenAiv1ChatCompletionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### VectorStores

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorStoreObject">VectorStoreObject</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreListResponse">OpenAiv1VectorStoreListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreDeleteResponse">OpenAiv1VectorStoreDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreSearchResponse">OpenAiv1VectorStoreSearchResponse</a>

Methods:

- <code title="post /v1/openai/v1/vector_stores">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreNewParams">OpenAIV1VectorStoreNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorStoreObject">VectorStoreObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/vector_stores/{vector_store_id}">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorStoreObject">VectorStoreObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/openai/v1/vector_stores/{vector_store_id}">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreUpdateParams">OpenAIV1VectorStoreUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorStoreObject">VectorStoreObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/vector_stores">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreListParams">OpenAIV1VectorStoreListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreListResponse">OpenAiv1VectorStoreListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/openai/v1/vector_stores/{vector_store_id}">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreDeleteResponse">OpenAiv1VectorStoreDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/openai/v1/vector_stores/{vector_store_id}/search">client.OpenAI.V1.VectorStores.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreSearchParams">OpenAIV1VectorStoreSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreSearchResponse">OpenAiv1VectorStoreSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Files

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ChunkingStrategyUnionParam">ChunkingStrategyUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#FileStatus">FileStatus</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ChunkingStrategyUnion">ChunkingStrategyUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Content">Content</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#FileStatus">FileStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorStoreFile">VectorStoreFile</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreFileListResponse">OpenAiv1VectorStoreFileListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreFileDeleteResponse">OpenAiv1VectorStoreFileDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreFileGetContentResponse">OpenAiv1VectorStoreFileGetContentResponse</a>

Methods:

- <code title="post /v1/openai/v1/vector_stores/{vector_store_id}/files">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileNewParams">OpenAIV1VectorStoreFileNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/vector_stores/{vector_store_id}/files/{file_id}">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileGetParams">OpenAIV1VectorStoreFileGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/openai/v1/vector_stores/{vector_store_id}/files/{file_id}">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileUpdateParams">OpenAIV1VectorStoreFileUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/vector_stores/{vector_store_id}/files">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileListParams">OpenAIV1VectorStoreFileListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreFileListResponse">OpenAiv1VectorStoreFileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/openai/v1/vector_stores/{vector_store_id}/files/{file_id}">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileDeleteParams">OpenAIV1VectorStoreFileDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreFileDeleteResponse">OpenAiv1VectorStoreFileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/vector_stores/{vector_store_id}/files/{file_id}/content">client.OpenAI.V1.VectorStores.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileService.GetContent">GetContent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1VectorStoreFileGetContentParams">OpenAIV1VectorStoreFileGetContentParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1VectorStoreFileGetContentResponse">OpenAiv1VectorStoreFileGetContentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Files

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#FilePurpose">FilePurpose</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIFile">OpenAIFile</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1FileListResponse">OpenAiv1FileListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1FileDeleteResponse">OpenAiv1FileDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1FileGetContentResponse">OpenAiv1FileGetContentResponse</a>

Methods:

- <code title="post /v1/openai/v1/files">client.OpenAI.V1.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1FileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1FileNewParams">OpenAIV1FileNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIFile">OpenAIFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/files/{file_id}">client.OpenAI.V1.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIFile">OpenAIFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/files">client.OpenAI.V1.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1FileListParams">OpenAIV1FileListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1FileListResponse">OpenAiv1FileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/openai/v1/files/{file_id}">client.OpenAI.V1.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1FileDeleteResponse">OpenAiv1FileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/openai/v1/files/{file_id}/content">client.OpenAI.V1.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAIV1FileService.GetContent">GetContent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#OpenAiv1FileGetContentResponse">OpenAiv1FileGetContentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Eval

## Benchmarks

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#BenchmarkConfigParam">BenchmarkConfigParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Benchmark">Benchmark</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvaluateResponse">EvaluateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkListResponse">EvalBenchmarkListResponse</a>

Methods:

- <code title="post /v1/eval/benchmarks">client.Eval.Benchmarks.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkNewParams">EvalBenchmarkNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/eval/benchmarks/{benchmark_id}">client.Eval.Benchmarks.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Benchmark">Benchmark</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/eval/benchmarks">client.Eval.Benchmarks.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkListResponse">EvalBenchmarkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/eval/benchmarks/{benchmark_id}/evaluations">client.Eval.Benchmarks.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkService.Evaluate">Evaluate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkEvaluateParams">EvalBenchmarkEvaluateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvaluateResponse">EvaluateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Job">Job</a>

Methods:

- <code title="get /v1/eval/benchmarks/{benchmark_id}/jobs/{job_id}">client.Eval.Benchmarks.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkJobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkJobGetParams">EvalBenchmarkJobGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/eval/benchmarks/{benchmark_id}/jobs/{job_id}">client.Eval.Benchmarks.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkJobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkJobCancelParams">EvalBenchmarkJobCancelParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/eval/benchmarks/{benchmark_id}/jobs/{job_id}/result">client.Eval.Benchmarks.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkJobService.Result">Result</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkJobResultParams">EvalBenchmarkJobResultParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvaluateResponse">EvaluateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/eval/benchmarks/{benchmark_id}/jobs">client.Eval.Benchmarks.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkJobService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EvalBenchmarkJobRunParams">EvalBenchmarkJobRunParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Datasets

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DataSourceUnionParam">DataSourceUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DataSourceUnion">DataSourceUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Dataset">Dataset</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetListResponse">DatasetListResponse</a>

Methods:

- <code title="post /v1/datasets">client.Datasets.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetNewParams">DatasetNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/datasets/{dataset_id}">client.Datasets.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/datasets">client.Datasets.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetListResponse">DatasetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/datasets/{dataset_id}">client.Datasets.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#DatasetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Models

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ModelType">ModelType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ModelType">ModelType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="post /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ModelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ModelNewParams">ModelNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ModelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# ScoringFunctions

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AggregationFunctionType">AggregationFunctionType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ParamTypeUnion">ParamTypeUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFnParamsUnion">ScoringFnParamsUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFnParamsType">ScoringFnParamsType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#AggregationFunctionType">AggregationFunctionType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ParamTypeUnionResp">ParamTypeUnionResp</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFn">ScoringFn</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFnParamsUnionResp">ScoringFnParamsUnionResp</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFnParamsType">ScoringFnParamsType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFunctionListResponse">ScoringFunctionListResponse</a>

Methods:

- <code title="post /v1/scoring-functions">client.ScoringFunctions.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFunctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFunctionNewParams">ScoringFunctionNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/scoring-functions/{scoring_fn_id}">client.ScoringFunctions.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFunctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, scoringFnID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFn">ScoringFn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/scoring-functions">client.ScoringFunctions.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringFunctionListResponse">ScoringFunctionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Shields

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Shield">Shield</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ShieldListResponse">ShieldListResponse</a>

Methods:

- <code title="post /v1/shields">client.Shields.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ShieldService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ShieldNewParams">ShieldNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Shield">Shield</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/shields/{identifier}">client.Shields.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ShieldService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Shield">Shield</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/shields">client.Shields.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ShieldService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ShieldListResponse">ShieldListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/shields/{identifier}">client.Shields.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ShieldService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Telemetry

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#EventType">EventType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#StructuredLogType">StructuredLogType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryQueryMetricResponse">TelemetryQueryMetricResponse</a>

Methods:

- <code title="post /v1/telemetry/events">client.Telemetry.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryService.LogEvent">LogEvent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryLogEventParams">TelemetryLogEventParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/telemetry/metrics/{metric_name}">client.Telemetry.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryService.QueryMetric">QueryMetric</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, metricName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryQueryMetricParams">TelemetryQueryMetricParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryQueryMetricResponse">TelemetryQueryMetricResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Traces

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Span">Span</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Trace">Trace</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryTraceQueryResponse">TelemetryTraceQueryResponse</a>

Methods:

- <code title="post /v1/telemetry/traces">client.Telemetry.Traces.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryTraceService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryTraceQueryParams">TelemetryTraceQueryParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryTraceQueryResponse">TelemetryTraceQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/telemetry/traces/{trace_id}/spans/{span_id}">client.Telemetry.Traces.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryTraceService.GetSpan">GetSpan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, spanID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryTraceGetSpanParams">TelemetryTraceGetSpanParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Span">Span</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/telemetry/traces/{trace_id}">client.Telemetry.Traces.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetryTraceService.GetTrace">GetTrace</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, traceID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Trace">Trace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Spans

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#QueryConditionParam">QueryConditionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetrySpanQueryResponse">TelemetrySpanQueryResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetrySpanGetTreeResponse">TelemetrySpanGetTreeResponse</a>

Methods:

- <code title="post /v1/telemetry/spans/export">client.Telemetry.Spans.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetrySpanService.Export">Export</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetrySpanExportParams">TelemetrySpanExportParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/telemetry/spans">client.Telemetry.Spans.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetrySpanService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetrySpanQueryParams">TelemetrySpanQueryParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetrySpanQueryResponse">TelemetrySpanQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/telemetry/spans/{span_id}/tree">client.Telemetry.Spans.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetrySpanService.GetTree">GetTree</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, spanID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetrySpanGetTreeParams">TelemetrySpanGetTreeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#TelemetrySpanGetTreeResponse">TelemetrySpanGetTreeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tools

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolParameter">ToolParameter</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Tool">Tool</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolParameterResp">ToolParameterResp</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolListResponse">ToolListResponse</a>

Methods:

- <code title="get /v1/tools/{tool_name}">client.Tools.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Tool">Tool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tools">client.Tools.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolListParams">ToolListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolListResponse">ToolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Toolgroups

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolGroup">ToolGroup</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolgroupListResponse">ToolgroupListResponse</a>

Methods:

- <code title="post /v1/toolgroups">client.Toolgroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolgroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolgroupNewParams">ToolgroupNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/toolgroups/{toolgroup_id}">client.Toolgroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolgroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolgroupID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolGroup">ToolGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/toolgroups">client.Toolgroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolgroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolgroupListResponse">ToolgroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/toolgroups/{toolgroup_id}">client.Toolgroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolgroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolgroupID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# VectorDBs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorDB">VectorDB</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorDBListResponse">VectorDBListResponse</a>

Methods:

- <code title="post /v1/vector-dbs">client.VectorDBs.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorDBService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorDBNewParams">VectorDBNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorDB">VectorDB</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector-dbs/{vector_db_id}">client.VectorDBs.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorDBService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorDBID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorDB">VectorDB</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector-dbs">client.VectorDBs.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorDBService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorDBListResponse">VectorDBListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/vector-dbs/{vector_db_id}">client.VectorDBs.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorDBService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorDBID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#HealthGetResponse">HealthGetResponse</a>

Methods:

- <code title="get /v1/health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#HealthService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#HealthGetResponse">HealthGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ToolRuntime

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolDefParam">ToolDefParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#URLParam">URLParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolDef">ToolDef</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#URL">URL</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeInvokeResponse">ToolRuntimeInvokeResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeListToolsResponse">ToolRuntimeListToolsResponse</a>

Methods:

- <code title="post /v1/tool-runtime/invoke">client.ToolRuntime.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeInvokeParams">ToolRuntimeInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeInvokeResponse">ToolRuntimeInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tool-runtime/list-tools">client.ToolRuntime.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeService.ListTools">ListTools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeListToolsParams">ToolRuntimeListToolsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeListToolsResponse">ToolRuntimeListToolsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RagTool

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeRagToolQueryResponse">ToolRuntimeRagToolQueryResponse</a>

Methods:

- <code title="post /v1/tool-runtime/rag-tool/insert">client.ToolRuntime.RagTool.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeRagToolService.Insert">Insert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeRagToolInsertParams">ToolRuntimeRagToolInsertParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/tool-runtime/rag-tool/query">client.ToolRuntime.RagTool.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeRagToolService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeRagToolQueryParams">ToolRuntimeRagToolQueryParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ToolRuntimeRagToolQueryResponse">ToolRuntimeRagToolQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VectorIo

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ChunkParam">ChunkParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#Chunk">Chunk</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorIoQueryResponse">VectorIoQueryResponse</a>

Methods:

- <code title="post /v1/vector-io/insert">client.VectorIo.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorIoService.Insert">Insert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorIoInsertParams">VectorIoInsertParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/vector-io/query">client.VectorIo.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorIoService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorIoQueryParams">VectorIoQueryParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VectorIoQueryResponse">VectorIoQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Providers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ProviderInfo">ProviderInfo</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ProviderListResponse">ProviderListResponse</a>

Methods:

- <code title="get /v1/providers/{provider_id}">client.Providers.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ProviderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, providerID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ProviderInfo">ProviderInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/providers">client.Providers.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ProviderListResponse">ProviderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Inspect

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InspectListRoutesResponse">InspectListRoutesResponse</a>

Methods:

- <code title="get /v1/inspect/routes">client.Inspect.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InspectService.ListRoutes">ListRoutes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#InspectListRoutesResponse">InspectListRoutesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Safety

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SafetyViolation">SafetyViolation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SafetyRunShieldResponse">SafetyRunShieldResponse</a>

Methods:

- <code title="post /v1/safety/run-shield">client.Safety.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SafetyService.RunShield">RunShield</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SafetyRunShieldParams">SafetyRunShieldParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SafetyRunShieldResponse">SafetyRunShieldResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Scoring

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringScoreResponse">ScoringScoreResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringScoreBatchResponse">ScoringScoreBatchResponse</a>

Methods:

- <code title="post /v1/scoring/score">client.Scoring.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringService.Score">Score</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringScoreParams">ScoringScoreParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringScoreResponse">ScoringScoreResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/scoring/score-batch">client.Scoring.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringService.ScoreBatch">ScoreBatch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringScoreBatchParams">ScoringScoreBatchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#ScoringScoreBatchResponse">ScoringScoreBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SyntheticDataGeneration

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SyntheticDataGenerationGenerateResponse">SyntheticDataGenerationGenerateResponse</a>

Methods:

- <code title="post /v1/synthetic-data-generation/generate">client.SyntheticDataGeneration.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SyntheticDataGenerationService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SyntheticDataGenerationGenerateParams">SyntheticDataGenerationGenerateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#SyntheticDataGenerationGenerateResponse">SyntheticDataGenerationGenerateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Version

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VersionGetResponse">VersionGetResponse</a>

Methods:

- <code title="get /v1/version">client.Version.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go">llamastackgoclient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/llamastack-go-client-go#VersionGetResponse">VersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

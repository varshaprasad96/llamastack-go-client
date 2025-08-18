// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/llamastack-go-client-go"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/testutil"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
)

func TestAgentNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.New(context.TODO(), llamastackgoclient.AgentNewParams{
		AgentConfig: llamastackgoclient.AgentConfigParam{
			Instructions: "instructions",
			Model:        "model",
			ClientTools: []llamastackgoclient.ToolDefParam{{
				Name:        "name",
				Description: llamastackgoclient.String("description"),
				Metadata: map[string]llamastackgoclient.ToolDefMetadataUnionParam{
					"foo": {
						OfBool: llamastackgoclient.Bool(true),
					},
				},
				Parameters: []llamastackgoclient.ToolParameter{{
					Description:   "description",
					Name:          "name",
					ParameterType: "parameter_type",
					Required:      true,
					Default: llamastackgoclient.ToolParameterDefaultUnion{
						OfBool: llamastackgoclient.Bool(true),
					},
				}},
			}},
			EnableSessionPersistence: llamastackgoclient.Bool(true),
			InputShields:             []string{"string"},
			MaxInferIters:            llamastackgoclient.Int(0),
			Name:                     llamastackgoclient.String("name"),
			OutputShields:            []string{"string"},
			ResponseFormat: llamastackgoclient.ResponseFormatUnionParam{
				OfJsonSchema: &llamastackgoclient.ResponseFormatJsonSchemaParam{
					JsonSchema: map[string]llamastackgoclient.ResponseFormatJsonSchemaJsonSchemaUnionParam{
						"foo": {
							OfBool: llamastackgoclient.Bool(true),
						},
					},
				},
			},
			SamplingParams: llamastackgoclient.SamplingParams{
				Strategy: llamastackgoclient.SamplingParamsStrategyUnion{
					OfGreedy: &llamastackgoclient.SamplingParamsStrategyGreedy{},
				},
				MaxTokens:         llamastackgoclient.Int(0),
				RepetitionPenalty: llamastackgoclient.Float(0),
				Stop:              []string{"string"},
			},
			ToolChoice: llamastackgoclient.AgentConfigToolChoiceAuto,
			ToolConfig: llamastackgoclient.ToolConfigParam{
				SystemMessageBehavior: llamastackgoclient.ToolConfigSystemMessageBehaviorAppend,
				ToolChoice:            llamastackgoclient.ToolConfigToolChoiceAuto,
				ToolPromptFormat:      llamastackgoclient.ToolConfigToolPromptFormatJson,
			},
			ToolPromptFormat: llamastackgoclient.AgentConfigToolPromptFormatJson,
			Toolgroups: []llamastackgoclient.AgentToolUnionParam{{
				OfString: llamastackgoclient.String("string"),
			}},
		},
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Get(context.TODO(), "agent_id")
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.List(context.TODO(), llamastackgoclient.AgentListParams{
		Limit:      llamastackgoclient.Int(0),
		StartIndex: llamastackgoclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Agents.Delete(context.TODO(), "agent_id")
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentNewSession(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.NewSession(
		context.TODO(),
		"agent_id",
		llamastackgoclient.AgentNewSessionParams{
			SessionName: "session_name",
		},
	)
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentGetSessionsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.GetSessions(
		context.TODO(),
		"agent_id",
		llamastackgoclient.AgentGetSessionsParams{
			Limit:      llamastackgoclient.Int(0),
			StartIndex: llamastackgoclient.Int(0),
		},
	)
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

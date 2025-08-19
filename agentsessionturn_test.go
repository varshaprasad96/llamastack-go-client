// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/varshaprasad96/llamastack-go-client"
	"github.com/varshaprasad96/llamastack-go-client/internal/testutil"
	"github.com/varshaprasad96/llamastack-go-client/option"
)

func TestAgentSessionTurnNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Session.Turn.New(
		context.TODO(),
		"session_id",
		llamastackclient.AgentSessionTurnNewParams{
			AgentID: "agent_id",
			Messages: []llamastackclient.AgentSessionTurnNewParamsMessageUnion{{
				OfUserMessage: &llamastackclient.UserMessageParam{
					Content: llamastackclient.InterleavedContentUnionParam{
						OfString: llamastackclient.String("string"),
					},
					Role: llamastackclient.UserMessageRoleSystem,
					Context: llamastackclient.InterleavedContentUnionParam{
						OfString: llamastackclient.String("string"),
					},
				},
			}},
			Documents: []llamastackclient.AgentSessionTurnNewParamsDocument{{
				Content: llamastackclient.AgentSessionTurnNewParamsDocumentContentUnion{
					OfString: llamastackclient.String("string"),
				},
				MimeType: "mime_type",
			}},
			Stream: llamastackclient.Bool(true),
			ToolConfig: llamastackclient.ToolConfigParam{
				SystemMessageBehavior: llamastackclient.ToolConfigSystemMessageBehaviorAppend,
				ToolChoice:            llamastackclient.ToolConfigToolChoiceAuto,
				ToolPromptFormat:      llamastackclient.ToolConfigToolPromptFormatJson,
			},
			Toolgroups: []llamastackclient.AgentToolUnionParam{{
				OfString: llamastackclient.String("string"),
			}},
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentSessionTurnGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Session.Turn.Get(
		context.TODO(),
		"turn_id",
		llamastackclient.AgentSessionTurnGetParams{
			AgentID:   "agent_id",
			SessionID: "session_id",
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentSessionTurnResumeWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Session.Turn.Resume(
		context.TODO(),
		"turn_id",
		llamastackclient.AgentSessionTurnResumeParams{
			AgentID:   "agent_id",
			SessionID: "session_id",
			ToolResponses: []llamastackclient.ToolResponseParam{{
				CallID: "call_id",
				Content: llamastackclient.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
				ToolName: llamastackclient.ToolResponseToolNameBraveSearch,
				Metadata: map[string]llamastackclient.ToolResponseMetadataUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
			}},
			Stream: llamastackclient.Bool(true),
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

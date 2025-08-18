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

func TestAgentSessionTurnNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Session.Turn.New(
		context.TODO(),
		"session_id",
		llamastackgoclient.AgentSessionTurnNewParams{
			AgentID: "agent_id",
			Messages: []llamastackgoclient.AgentSessionTurnNewParamsMessageUnion{{
				OfUserMessage: &llamastackgoclient.UserMessageParam{
					Content: llamastackgoclient.InterleavedContentUnionParam{
						OfString: llamastackgoclient.String("string"),
					},
					Context: llamastackgoclient.InterleavedContentUnionParam{
						OfString: llamastackgoclient.String("string"),
					},
				},
			}},
			Documents: []llamastackgoclient.AgentSessionTurnNewParamsDocument{{
				Content: llamastackgoclient.AgentSessionTurnNewParamsDocumentContentUnion{
					OfString: llamastackgoclient.String("string"),
				},
				MimeType: "mime_type",
			}},
			Stream: llamastackgoclient.Bool(true),
			ToolConfig: llamastackgoclient.ToolConfigParam{
				SystemMessageBehavior: llamastackgoclient.ToolConfigSystemMessageBehaviorAppend,
				ToolChoice:            llamastackgoclient.ToolConfigToolChoiceAuto,
				ToolPromptFormat:      llamastackgoclient.ToolConfigToolPromptFormatJson,
			},
			Toolgroups: []llamastackgoclient.AgentToolUnionParam{{
				OfString: llamastackgoclient.String("string"),
			}},
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

func TestAgentSessionTurnGet(t *testing.T) {
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
	_, err := client.Agents.Session.Turn.Get(
		context.TODO(),
		"turn_id",
		llamastackgoclient.AgentSessionTurnGetParams{
			AgentID:   "agent_id",
			SessionID: "session_id",
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

func TestAgentSessionTurnResumeWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Session.Turn.Resume(
		context.TODO(),
		"turn_id",
		llamastackgoclient.AgentSessionTurnResumeParams{
			AgentID:   "agent_id",
			SessionID: "session_id",
			ToolResponses: []llamastackgoclient.ToolResponseParam{{
				CallID: "call_id",
				Content: llamastackgoclient.InterleavedContentUnionParam{
					OfString: llamastackgoclient.String("string"),
				},
				ToolName: llamastackgoclient.ToolResponseToolNameBraveSearch,
				Metadata: map[string]llamastackgoclient.ToolResponseMetadataUnionParam{
					"foo": {
						OfBool: llamastackgoclient.Bool(true),
					},
				},
			}},
			Stream: llamastackgoclient.Bool(true),
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

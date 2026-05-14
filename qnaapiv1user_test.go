// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bjwtestbella_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/bjw-test-bella-go"
	"github.com/stainless-sdks/bjw-test-bella-go/internal/testutil"
	"github.com/stainless-sdks/bjw-test-bella-go/option"
)

func TestQnaAPIV1UserChatWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bjwtestbella.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Qna.API.V1.Users.Chat(
		context.TODO(),
		"userId",
		bjwtestbella.QnaAPIV1UserChatParams{
			UserSay: "안녕하세요?",
			XAPIKey: "x-api-key",
			XUserID: "x-user-id",
			DocumentFilter: bjwtestbella.SearchFilterBody{
				"or": "bar",
			},
			SessionID: bjwtestbella.String("sessionId"),
		},
	)
	if err != nil {
		var apierr *bjwtestbella.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQnaAPIV1UserGetBotProfile(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bjwtestbella.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Qna.API.V1.Users.GetBotProfile(
		context.TODO(),
		"userId",
		bjwtestbella.QnaAPIV1UserGetBotProfileParams{
			XAPIKey: "x-api-key",
			XUserID: "x-user-id",
		},
	)
	if err != nil {
		var apierr *bjwtestbella.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQnaAPIV1UserGetChatLimitStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bjwtestbella.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Qna.API.V1.Users.GetChatLimitStatus(
		context.TODO(),
		"userId",
		bjwtestbella.QnaAPIV1UserGetChatLimitStatusParams{
			XAPIKey: "x-api-key",
			XUserID: "x-user-id",
		},
	)
	if err != nil {
		var apierr *bjwtestbella.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQnaAPIV1UserGetWelcomeMessage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bjwtestbella.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Qna.API.V1.Users.GetWelcomeMessage(
		context.TODO(),
		"userId",
		bjwtestbella.QnaAPIV1UserGetWelcomeMessageParams{
			XAPIKey: "x-api-key",
			XUserID: "x-user-id",
		},
	)
	if err != nil {
		var apierr *bjwtestbella.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQnaAPIV1UserListChatHistories(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bjwtestbella.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Qna.API.V1.Users.ListChatHistories(
		context.TODO(),
		"userId",
		bjwtestbella.QnaAPIV1UserListChatHistoriesParams{
			XAPIKey: "x-api-key",
			XUserID: "x-user-id",
		},
	)
	if err != nil {
		var apierr *bjwtestbella.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQnaAPIV1UserStreamingChatWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bjwtestbella.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Qna.API.V1.Users.StreamingChat(
		context.TODO(),
		"userId",
		bjwtestbella.QnaAPIV1UserStreamingChatParams{
			UserSay: "안녕하세요?",
			XAPIKey: "x-api-key",
			XUserID: "x-user-id",
			DocumentFilter: bjwtestbella.SearchFilterBody{
				"or": "bar",
			},
			SessionID: bjwtestbella.String("sessionId"),
		},
	)
	if err != nil {
		var apierr *bjwtestbella.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

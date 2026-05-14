// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bjwtestbella_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/bjw-test-bella-go"
	"github.com/stainless-sdks/bjw-test-bella-go/internal/testutil"
	"github.com/stainless-sdks/bjw-test-bella-go/option"
)

func TestQnaAPIV1UserSessionLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.Qna.API.V1.Users.SessionLogs.List(
		context.TODO(),
		"userId",
		bjwtestbella.QnaAPIV1UserSessionLogListParams{
			Page:           0,
			XAPIKey:        "x-api-key",
			XUserID:        "x-user-id",
			BotMessageType: bjwtestbella.QnaAPIV1UserSessionLogListParamsBotMessageTypeValid,
			Channel:        bjwtestbella.String("channel"),
			Count:          bjwtestbella.Float(0),
			From:           bjwtestbella.Time(time.Now()),
			Q:              bjwtestbella.String("q"),
			To:             bjwtestbella.Time(time.Now()),
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

func TestQnaAPIV1UserSessionLogGet(t *testing.T) {
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
	_, err := client.Qna.API.V1.Users.SessionLogs.Get(
		context.TODO(),
		"sessionId",
		bjwtestbella.QnaAPIV1UserSessionLogGetParams{
			UserID:  "userId",
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bjwtestbella_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/w4-baejinwoo/bjw-test-bella-go"
	"github.com/w4-baejinwoo/bjw-test-bella-go/internal/testutil"
	"github.com/w4-baejinwoo/bjw-test-bella-go/option"
)

func TestQnaAPIV1GetQuestionSamples(t *testing.T) {
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
	_, err := client.Qna.API.V1.GetQuestionSamples(context.TODO(), bjwtestbella.QnaAPIV1GetQuestionSamplesParams{
		Limit:   "limit",
		XAPIKey: "x-api-key",
		XUserID: "x-user-id",
	})
	if err != nil {
		var apierr *bjwtestbella.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQnaAPIV1GetSessionLog(t *testing.T) {
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
	_, err := client.Qna.API.V1.GetSessionLog(
		context.TODO(),
		"sessionId",
		bjwtestbella.QnaAPIV1GetSessionLogParams{
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

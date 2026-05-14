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

func TestQnaAPIV1DocumentBulkNew(t *testing.T) {
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
	_, err := client.Qna.API.V1.Documents.BulkNew(context.TODO(), bjwtestbella.QnaAPIV1DocumentBulkNewParams{
		Documents: []bjwtestbella.QnaAPIV1DocumentBulkNewParamsDocument{{
			ID:   "id",
			Type: "plain",
			Extra: map[string]any{
				"key1": "bar",
				"key2": "bar",
			},
			FilePath: bjwtestbella.String("filePath"),
			ParsingOptions: bjwtestbella.QnaAPIV1DocumentBulkNewParamsDocumentParsingOptions{
				Parser: "parser",
			},
			Ref: bjwtestbella.QnaAPIV1DocumentBulkNewParamsDocumentRef{
				Link: "link",
			},
			Title:      bjwtestbella.String("title"),
			TraceID:    bjwtestbella.String("traceId"),
			UseSummary: bjwtestbella.Bool(true),
			Validity: bjwtestbella.QnaAPIV1DocumentBulkNewParamsDocumentValidity{
				EndedAt:   bjwtestbella.Time(time.Now()),
				StartedAt: bjwtestbella.Time(time.Now()),
			},
			Value: bjwtestbella.String("value"),
		}},
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

func TestQnaAPIV1DocumentBulkDelete(t *testing.T) {
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
	_, err := client.Qna.API.V1.Documents.BulkDelete(context.TODO(), bjwtestbella.QnaAPIV1DocumentBulkDeleteParams{
		DocumentIDs: []string{"string"},
		XAPIKey:     "x-api-key",
		XUserID:     "x-user-id",
	})
	if err != nil {
		var apierr *bjwtestbella.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQnaAPIV1DocumentListByTraceIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Qna.API.V1.Documents.ListByTraceID(context.TODO(), bjwtestbella.QnaAPIV1DocumentListByTraceIDParams{
		XAPIKey: "x-api-key",
		XUserID: "x-user-id",
		TraceID: bjwtestbella.String("traceId"),
	})
	if err != nil {
		var apierr *bjwtestbella.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQnaAPIV1DocumentListChunksWithOptionalParams(t *testing.T) {
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
	_, err := client.Qna.API.V1.Documents.ListChunks(context.TODO(), bjwtestbella.QnaAPIV1DocumentListChunksParams{
		XAPIKey:     "x-api-key",
		XUserID:     "x-user-id",
		ChunkIDs:    []string{"string"},
		DocumentIDs: []string{"string"},
	})
	if err != nil {
		var apierr *bjwtestbella.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQnaAPIV1DocumentUpdateExtraData(t *testing.T) {
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
	_, err := client.Qna.API.V1.Documents.UpdateExtraData(
		context.TODO(),
		"id",
		bjwtestbella.QnaAPIV1DocumentUpdateExtraDataParams{
			Body: map[string]any{
				"foo": "bar",
			},
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

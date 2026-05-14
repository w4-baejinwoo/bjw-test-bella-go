// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bjwtestbella_test

import (
	"context"
	"os"
	"testing"

	"github.com/w4-baejinwoo/bjw-test-bella-go"
	"github.com/w4-baejinwoo/bjw-test-bella-go/internal/testutil"
	"github.com/w4-baejinwoo/bjw-test-bella-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Qna.API.V1.Documents.BulkNew(context.TODO(), bjwtestbella.QnaAPIV1DocumentBulkNewParams{
		Documents: []bjwtestbella.QnaAPIV1DocumentBulkNewParamsDocument{{
			ID:   "id",
			Type: "plain",
		}},
		XAPIKey: "x-api-key",
		XUserID: "x-user-id",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response)
}

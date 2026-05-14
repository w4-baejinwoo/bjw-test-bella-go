// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bjwtestbella

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/bjw-test-bella-go/internal/apijson"
	"github.com/stainless-sdks/bjw-test-bella-go/internal/requestconfig"
	"github.com/stainless-sdks/bjw-test-bella-go/option"
	"github.com/stainless-sdks/bjw-test-bella-go/packages/param"
	"github.com/stainless-sdks/bjw-test-bella-go/packages/respjson"
)

// QnaAPIV1UserSessionService contains methods and other services that help with
// interacting with the bjw-test-bella API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQnaAPIV1UserSessionService] method instead.
type QnaAPIV1UserSessionService struct {
	options []option.RequestOption
}

// NewQnaAPIV1UserSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewQnaAPIV1UserSessionService(opts ...option.RequestOption) (r QnaAPIV1UserSessionService) {
	r = QnaAPIV1UserSessionService{}
	r.options = opts
	return
}

// Clear session context
func (r *QnaAPIV1UserSessionService) ClearContext(ctx context.Context, sessionID string, params QnaAPIV1UserSessionClearContextParams, opts ...option.RequestOption) (res *QnaApiv1UserSessionClearContextResponse, err error) {
	if !param.IsOmitted(params.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", params.XAPIKey)))
	}
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	if params.UserID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("qna/api/v1/users/%s/sessions/%s/context", url.PathEscape(params.UserID), url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type QnaApiv1UserSessionClearContextResponse struct {
	Data map[string]any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserSessionClearContextResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserSessionClearContextResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaAPIV1UserSessionClearContextParams struct {
	UserID  string `path:"userId" api:"required" json:"-"`
	XAPIKey string `header:"x-api-key" api:"required" json:"-"`
	XUserID string `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bjwtestbella

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/w4-baejinwoo/bjw-test-bella-go/internal/apijson"
	"github.com/w4-baejinwoo/bjw-test-bella-go/internal/apiquery"
	"github.com/w4-baejinwoo/bjw-test-bella-go/internal/requestconfig"
	"github.com/w4-baejinwoo/bjw-test-bella-go/option"
	"github.com/w4-baejinwoo/bjw-test-bella-go/packages/param"
	"github.com/w4-baejinwoo/bjw-test-bella-go/packages/respjson"
)

// QnaAPIV1UserSessionLogService contains methods and other services that help with
// interacting with the bjw-test-bella API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQnaAPIV1UserSessionLogService] method instead.
type QnaAPIV1UserSessionLogService struct {
	options []option.RequestOption
}

// NewQnaAPIV1UserSessionLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewQnaAPIV1UserSessionLogService(opts ...option.RequestOption) (r QnaAPIV1UserSessionLogService) {
	r = QnaAPIV1UserSessionLogService{}
	r.options = opts
	return
}

// List sessions of user
func (r *QnaAPIV1UserSessionLogService) List(ctx context.Context, userID string, params QnaAPIV1UserSessionLogListParams, opts ...option.RequestOption) (res *QnaApiv1UserSessionLogListResponse, err error) {
	if !param.IsOmitted(params.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", params.XAPIKey)))
	}
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("qna/api/v1/users/%s/session-logs", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get session log of user
func (r *QnaAPIV1UserSessionLogService) Get(ctx context.Context, sessionID string, params QnaAPIV1UserSessionLogGetParams, opts ...option.RequestOption) (res *QnaApiv1UserSessionLogGetResponse, err error) {
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
	path := fmt.Sprintf("qna/api/v1/users/%s/session-logs/%s", url.PathEscape(params.UserID), url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type QnaApiv1UserSessionLogListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserSessionLogListResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserSessionLogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserSessionLogGetResponse struct {
	Data QnaApiv1UserSessionLogGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserSessionLogGetResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserSessionLogGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserSessionLogGetResponseData struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	GetSessionLogViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserSessionLogGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserSessionLogGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaAPIV1UserSessionLogListParams struct {
	Page    float64              `query:"page" api:"required" json:"-"`
	XAPIKey string               `header:"x-api-key" api:"required" json:"-"`
	XUserID string               `header:"x-user-id" api:"required" json:"-"`
	Channel param.Opt[string]    `query:"channel,omitzero" json:"-"`
	Count   param.Opt[float64]   `query:"count,omitzero" json:"-"`
	From    param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	Q       param.Opt[string]    `query:"q,omitzero" json:"-"`
	To      param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Any of "valid", "piiDetected", "sessionExpired", "invalidQuestion",
	// "invalidAnswer", "greeting", "outOfDomain", "openAiInternalError", "guide",
	// "functionCallFailed", "noPluginsInUse", "unexpected".
	BotMessageType QnaAPIV1UserSessionLogListParamsBotMessageType `query:"botMessageType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QnaAPIV1UserSessionLogListParams]'s query parameters as
// `url.Values`.
func (r QnaAPIV1UserSessionLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QnaAPIV1UserSessionLogListParamsBotMessageType string

const (
	QnaAPIV1UserSessionLogListParamsBotMessageTypeValid               QnaAPIV1UserSessionLogListParamsBotMessageType = "valid"
	QnaAPIV1UserSessionLogListParamsBotMessageTypePiiDetected         QnaAPIV1UserSessionLogListParamsBotMessageType = "piiDetected"
	QnaAPIV1UserSessionLogListParamsBotMessageTypeSessionExpired      QnaAPIV1UserSessionLogListParamsBotMessageType = "sessionExpired"
	QnaAPIV1UserSessionLogListParamsBotMessageTypeInvalidQuestion     QnaAPIV1UserSessionLogListParamsBotMessageType = "invalidQuestion"
	QnaAPIV1UserSessionLogListParamsBotMessageTypeInvalidAnswer       QnaAPIV1UserSessionLogListParamsBotMessageType = "invalidAnswer"
	QnaAPIV1UserSessionLogListParamsBotMessageTypeGreeting            QnaAPIV1UserSessionLogListParamsBotMessageType = "greeting"
	QnaAPIV1UserSessionLogListParamsBotMessageTypeOutOfDomain         QnaAPIV1UserSessionLogListParamsBotMessageType = "outOfDomain"
	QnaAPIV1UserSessionLogListParamsBotMessageTypeOpenAIInternalError QnaAPIV1UserSessionLogListParamsBotMessageType = "openAiInternalError"
	QnaAPIV1UserSessionLogListParamsBotMessageTypeGuide               QnaAPIV1UserSessionLogListParamsBotMessageType = "guide"
	QnaAPIV1UserSessionLogListParamsBotMessageTypeFunctionCallFailed  QnaAPIV1UserSessionLogListParamsBotMessageType = "functionCallFailed"
	QnaAPIV1UserSessionLogListParamsBotMessageTypeNoPluginsInUse      QnaAPIV1UserSessionLogListParamsBotMessageType = "noPluginsInUse"
	QnaAPIV1UserSessionLogListParamsBotMessageTypeUnexpected          QnaAPIV1UserSessionLogListParamsBotMessageType = "unexpected"
)

type QnaAPIV1UserSessionLogGetParams struct {
	UserID  string `path:"userId" api:"required" json:"-"`
	XAPIKey string `header:"x-api-key" api:"required" json:"-"`
	XUserID string `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

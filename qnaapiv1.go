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

	"github.com/stainless-sdks/bjw-test-bella-go/internal/apijson"
	"github.com/stainless-sdks/bjw-test-bella-go/internal/apiquery"
	"github.com/stainless-sdks/bjw-test-bella-go/internal/requestconfig"
	"github.com/stainless-sdks/bjw-test-bella-go/option"
	"github.com/stainless-sdks/bjw-test-bella-go/packages/param"
	"github.com/stainless-sdks/bjw-test-bella-go/packages/respjson"
)

// QnaAPIV1Service contains methods and other services that help with interacting
// with the bjw-test-bella API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQnaAPIV1Service] method instead.
type QnaAPIV1Service struct {
	options       []option.RequestOption
	Documents     QnaAPIV1DocumentService
	Users         QnaAPIV1UserService
	PublicUploads QnaAPIV1PublicUploadService
}

// NewQnaAPIV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQnaAPIV1Service(opts ...option.RequestOption) (r QnaAPIV1Service) {
	r = QnaAPIV1Service{}
	r.options = opts
	r.Documents = NewQnaAPIV1DocumentService(opts...)
	r.Users = NewQnaAPIV1UserService(opts...)
	r.PublicUploads = NewQnaAPIV1PublicUploadService(opts...)
	return
}

// Get question samples for a project
func (r *QnaAPIV1Service) GetQuestionSamples(ctx context.Context, params QnaAPIV1GetQuestionSamplesParams, opts ...option.RequestOption) (res *QnaApiv1GetQuestionSamplesResponse, err error) {
	if !param.IsOmitted(params.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", params.XAPIKey)))
	}
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	path := "qna/api/v1/question-samples"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get session log
func (r *QnaAPIV1Service) GetSessionLog(ctx context.Context, sessionID string, query QnaAPIV1GetSessionLogParams, opts ...option.RequestOption) (res *QnaApiv1GetSessionLogResponse, err error) {
	if !param.IsOmitted(query.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", query.XAPIKey)))
	}
	if !param.IsOmitted(query.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", query.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("qna/api/v1/session-logs/%s", url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type GetSessionLogViewModel struct {
	Conversations   []GetSessionLogViewModelConversation `json:"conversations" api:"required"`
	FirstUserSaidAt time.Time                            `json:"firstUserSaidAt" api:"required" format:"date-time"`
	PreviewContents string                               `json:"previewContents" api:"required"`
	SessionID       string                               `json:"sessionId" api:"required"`
	UserID          string                               `json:"userId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conversations   respjson.Field
		FirstUserSaidAt respjson.Field
		PreviewContents respjson.Field
		SessionID       respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetSessionLogViewModel) RawJSON() string { return r.JSON.raw }
func (r *GetSessionLogViewModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetSessionLogViewModelConversation struct {
	ID                  string                                       `json:"id" api:"required"`
	BotMessage          GetSessionLogViewModelConversationBotMessage `json:"botMessage" api:"required"`
	ClearedContextAfter bool                                         `json:"clearedContextAfter" api:"required"`
	// 응답 소요 시간 (sec)
	ResponseDuration float64                                       `json:"responseDuration" api:"required"`
	UserMessage      GetSessionLogViewModelConversationUserMessage `json:"userMessage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		BotMessage          respjson.Field
		ClearedContextAfter respjson.Field
		ResponseDuration    respjson.Field
		UserMessage         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetSessionLogViewModelConversation) RawJSON() string { return r.JSON.raw }
func (r *GetSessionLogViewModelConversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetSessionLogViewModelConversationBotMessage struct {
	// Any of "plain", "markdown".
	AnswerFormat string                                              `json:"answerFormat" api:"required"`
	Chunks       []GetSessionLogViewModelConversationBotMessageChunk `json:"chunks" api:"required"`
	Contents     string                                              `json:"contents" api:"required"`
	SaidAt       time.Time                                           `json:"saidAt" api:"required" format:"date-time"`
	// Any of "valid", "piiDetected", "sessionExpired", "invalidQuestion",
	// "invalidAnswer", "greeting", "outOfDomain", "openAiInternalError", "guide",
	// "functionCallFailed", "noPluginsInUse", "unexpected".
	Type            string                                            `json:"type" api:"required"`
	CalledFunctions []string                                          `json:"calledFunctions"`
	Extra           GetSessionLogViewModelConversationBotMessageExtra `json:"extra"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnswerFormat    respjson.Field
		Chunks          respjson.Field
		Contents        respjson.Field
		SaidAt          respjson.Field
		Type            respjson.Field
		CalledFunctions respjson.Field
		Extra           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetSessionLogViewModelConversationBotMessage) RawJSON() string { return r.JSON.raw }
func (r *GetSessionLogViewModelConversationBotMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetSessionLogViewModelConversationBotMessageChunk struct {
	ID            string                                                    `json:"id" api:"required"`
	Contents      string                                                    `json:"contents" api:"required"`
	DocumentID    string                                                    `json:"documentId" api:"required"`
	Position      GetSessionLogViewModelConversationBotMessageChunkPosition `json:"position" api:"required"`
	DocumentTitle string                                                    `json:"documentTitle"`
	// 문서 생성 시 저장했던 임의의 데이터
	Extra  map[string]any `json:"extra"`
	Reason string         `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Contents      respjson.Field
		DocumentID    respjson.Field
		Position      respjson.Field
		DocumentTitle respjson.Field
		Extra         respjson.Field
		Reason        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetSessionLogViewModelConversationBotMessageChunk) RawJSON() string { return r.JSON.raw }
func (r *GetSessionLogViewModelConversationBotMessageChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetSessionLogViewModelConversationBotMessageChunkPosition struct {
	Offset float64 `json:"offset" api:"required"`
	Page   float64 `json:"page" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetSessionLogViewModelConversationBotMessageChunkPosition) RawJSON() string {
	return r.JSON.raw
}
func (r *GetSessionLogViewModelConversationBotMessageChunkPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetSessionLogViewModelConversationBotMessageExtra struct {
	Citations         []GetSessionLogViewModelConversationBotMessageExtraCitation `json:"citations"`
	RetrievedChunkIDs []string                                                    `json:"retrievedChunkIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Citations         respjson.Field
		RetrievedChunkIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetSessionLogViewModelConversationBotMessageExtra) RawJSON() string { return r.JSON.raw }
func (r *GetSessionLogViewModelConversationBotMessageExtra) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetSessionLogViewModelConversationBotMessageExtraCitation struct {
	ID              string   `json:"id" api:"required"`
	BeginIndex      float64  `json:"beginIndex" api:"required"`
	ChunkBosIndex   float64  `json:"chunkBosIndex" api:"required"`
	ChunkEosIndex   float64  `json:"chunkEosIndex" api:"required"`
	ChunkID         string   `json:"chunkId" api:"required"`
	EndIndex        float64  `json:"endIndex" api:"required"`
	InsertedIndices []string `json:"insertedIndices" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		BeginIndex      respjson.Field
		ChunkBosIndex   respjson.Field
		ChunkEosIndex   respjson.Field
		ChunkID         respjson.Field
		EndIndex        respjson.Field
		InsertedIndices respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetSessionLogViewModelConversationBotMessageExtraCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *GetSessionLogViewModelConversationBotMessageExtraCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetSessionLogViewModelConversationUserMessage struct {
	Contents string    `json:"contents" api:"required"`
	SaidAt   time.Time `json:"saidAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		SaidAt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetSessionLogViewModelConversationUserMessage) RawJSON() string { return r.JSON.raw }
func (r *GetSessionLogViewModelConversationUserMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1GetQuestionSamplesResponse struct {
	Data QnaApiv1GetQuestionSamplesResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1GetQuestionSamplesResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1GetQuestionSamplesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1GetQuestionSamplesResponseData struct {
	QuestionSamples []string       `json:"questionSamples" api:"required"`
	ExtraFields     map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		QuestionSamples respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1GetQuestionSamplesResponseData) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1GetQuestionSamplesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1GetSessionLogResponse struct {
	Data QnaApiv1GetSessionLogResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1GetSessionLogResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1GetSessionLogResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1GetSessionLogResponseData struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	GetSessionLogViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1GetSessionLogResponseData) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1GetSessionLogResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaAPIV1GetQuestionSamplesParams struct {
	Limit   string `query:"limit" api:"required" json:"-"`
	XAPIKey string `header:"x-api-key" api:"required" json:"-"`
	XUserID string `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [QnaAPIV1GetQuestionSamplesParams]'s query parameters as
// `url.Values`.
func (r QnaAPIV1GetQuestionSamplesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QnaAPIV1GetSessionLogParams struct {
	XAPIKey string `header:"x-api-key" api:"required" json:"-"`
	XUserID string `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

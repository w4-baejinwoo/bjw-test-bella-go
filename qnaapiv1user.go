// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bjwtestbella

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/bjw-test-bella-go/internal/apijson"
	"github.com/stainless-sdks/bjw-test-bella-go/internal/requestconfig"
	"github.com/stainless-sdks/bjw-test-bella-go/option"
	"github.com/stainless-sdks/bjw-test-bella-go/packages/param"
	"github.com/stainless-sdks/bjw-test-bella-go/packages/respjson"
)

// QnaAPIV1UserService contains methods and other services that help with
// interacting with the bjw-test-bella API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQnaAPIV1UserService] method instead.
type QnaAPIV1UserService struct {
	options     []option.RequestOption
	Sessions    QnaAPIV1UserSessionService
	SessionLogs QnaAPIV1UserSessionLogService
}

// NewQnaAPIV1UserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQnaAPIV1UserService(opts ...option.RequestOption) (r QnaAPIV1UserService) {
	r = QnaAPIV1UserService{}
	r.options = opts
	r.Sessions = NewQnaAPIV1UserSessionService(opts...)
	r.SessionLogs = NewQnaAPIV1UserSessionLogService(opts...)
	return
}

// Chat
func (r *QnaAPIV1UserService) Chat(ctx context.Context, userID string, params QnaAPIV1UserChatParams, opts ...option.RequestOption) (res *QnaApiv1UserChatResponse, err error) {
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
	path := fmt.Sprintf("qna/api/v1/users/%s/chat", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get bot profile
func (r *QnaAPIV1UserService) GetBotProfile(ctx context.Context, userID string, query QnaAPIV1UserGetBotProfileParams, opts ...option.RequestOption) (res *QnaApiv1UserGetBotProfileResponse, err error) {
	if !param.IsOmitted(query.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", query.XAPIKey)))
	}
	if !param.IsOmitted(query.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", query.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("qna/api/v1/users/%s/bot-profile", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get chat limit status
func (r *QnaAPIV1UserService) GetChatLimitStatus(ctx context.Context, userID string, query QnaAPIV1UserGetChatLimitStatusParams, opts ...option.RequestOption) (res *QnaApiv1UserGetChatLimitStatusResponse, err error) {
	if !param.IsOmitted(query.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", query.XAPIKey)))
	}
	if !param.IsOmitted(query.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", query.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("qna/api/v1/users/%s/chat-limit-status", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get welcome message
func (r *QnaAPIV1UserService) GetWelcomeMessage(ctx context.Context, userID string, query QnaAPIV1UserGetWelcomeMessageParams, opts ...option.RequestOption) (res *QnaApiv1UserGetWelcomeMessageResponse, err error) {
	if !param.IsOmitted(query.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", query.XAPIKey)))
	}
	if !param.IsOmitted(query.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", query.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("qna/api/v1/users/%s/welcome-message", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List chat histories
func (r *QnaAPIV1UserService) ListChatHistories(ctx context.Context, userID string, query QnaAPIV1UserListChatHistoriesParams, opts ...option.RequestOption) (res *QnaApiv1UserListChatHistoriesResponse, err error) {
	if !param.IsOmitted(query.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", query.XAPIKey)))
	}
	if !param.IsOmitted(query.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", query.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("qna/api/v1/users/%s/chat-histories", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Streaming chat
func (r *QnaAPIV1UserService) StreamingChat(ctx context.Context, userID string, params QnaAPIV1UserStreamingChatParams, opts ...option.RequestOption) (res *QnaApiv1UserStreamingChatResponse, err error) {
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
	path := fmt.Sprintf("qna/api/v1/users/%s/streaming-chat", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type BaseMessage struct {
	Text      string    `json:"text" api:"required"`
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessage) RawJSON() string { return r.JSON.raw }
func (r *BaseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatViewModelChunk struct {
	ID         string                     `json:"id" api:"required"`
	Contents   string                     `json:"contents" api:"required"`
	DocumentID string                     `json:"documentId" api:"required"`
	Position   ChatViewModelChunkPosition `json:"position" api:"required"`
	Document   ChatViewModelChunkDocument `json:"document"`
	Reason     string                     `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Contents    respjson.Field
		DocumentID  respjson.Field
		Position    respjson.Field
		Document    respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatViewModelChunk) RawJSON() string { return r.JSON.raw }
func (r *ChatViewModelChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatViewModelChunkPosition struct {
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
func (r ChatViewModelChunkPosition) RawJSON() string { return r.JSON.raw }
func (r *ChatViewModelChunkPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatViewModelChunkDocument struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Title     string    `json:"title" api:"required"`
	// Any of "plain", "file", "url", "html".
	Type  string `json:"type" api:"required"`
	Value string `json:"value" api:"required"`
	// 문서 생성 시 저장했던 임의의 데이터
	Extra map[string]any                `json:"extra"`
	Ref   ChatViewModelChunkDocumentRef `json:"ref"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		Extra       respjson.Field
		Ref         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatViewModelChunkDocument) RawJSON() string { return r.JSON.raw }
func (r *ChatViewModelChunkDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatViewModelChunkDocumentRef struct {
	Link string `json:"link"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Link        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatViewModelChunkDocumentRef) RawJSON() string { return r.JSON.raw }
func (r *ChatViewModelChunkDocumentRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatViewModelForAgent struct {
	ID string `json:"id" api:"required"`
	// Any of "plain", "markdown".
	AnswerFormat      ChatViewModelForAgentAnswerFormat            `json:"answerFormat" api:"required"`
	AssistantMessages []ChatViewModelForAgentAssistantMessageUnion `json:"assistantMessages" api:"required"`
	IsInterim         bool                                         `json:"isInterim" api:"required"`
	SessionID         string                                       `json:"sessionId" api:"required"`
	// Any of "valid", "piiDetected", "sessionExpired", "invalidQuestion",
	// "invalidAnswer", "greeting", "outOfDomain", "openAiInternalError", "guide",
	// "functionCallFailed", "noPluginsInUse", "unexpected".
	Status     ChatViewModelForAgentStatus `json:"status" api:"required"`
	UserSaid   string                      `json:"userSaid" api:"required"`
	UserSaidAt time.Time                   `json:"userSaidAt" api:"required" format:"date-time"`
	Variant    any                         `json:"variant" api:"required"`
	// 질의에 대한 Agent의 최종 응답 시점. isInterim이 false인 경우에만 전달됨
	BotSaidAt time.Time `json:"botSaidAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AnswerFormat      respjson.Field
		AssistantMessages respjson.Field
		IsInterim         respjson.Field
		SessionID         respjson.Field
		Status            respjson.Field
		UserSaid          respjson.Field
		UserSaidAt        respjson.Field
		Variant           respjson.Field
		BotSaidAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatViewModelForAgent) RawJSON() string { return r.JSON.raw }
func (r *ChatViewModelForAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatViewModelForAgentAnswerFormat string

const (
	ChatViewModelForAgentAnswerFormatPlain    ChatViewModelForAgentAnswerFormat = "plain"
	ChatViewModelForAgentAnswerFormatMarkdown ChatViewModelForAgentAnswerFormat = "markdown"
)

// ChatViewModelForAgentAssistantMessageUnion contains all possible properties and
// values from [ChatViewModelForAgentAssistantMessageTextMessage],
// [ChatViewModelForAgentAssistantMessageFunctionCallMessage],
// [ChatViewModelForAgentAssistantMessageQnaMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatViewModelForAgentAssistantMessageUnion struct {
	AssistantMessageType any `json:"assistantMessageType"`
	// This field is from variant [ChatViewModelForAgentAssistantMessageTextMessage].
	BaseMessage BaseMessage `json:"baseMessage"`
	// This field is from variant
	// [ChatViewModelForAgentAssistantMessageFunctionCallMessage].
	FunctionCallArguments string `json:"functionCallArguments"`
	// This field is from variant
	// [ChatViewModelForAgentAssistantMessageFunctionCallMessage].
	FunctionCallID string `json:"functionCallId"`
	// This field is from variant
	// [ChatViewModelForAgentAssistantMessageFunctionCallMessage].
	FunctionDescription string `json:"functionDescription"`
	// This field is from variant
	// [ChatViewModelForAgentAssistantMessageFunctionCallMessage].
	FunctionID string `json:"functionId"`
	// This field is from variant
	// [ChatViewModelForAgentAssistantMessageFunctionCallMessage].
	FunctionName string `json:"functionName"`
	// This field is from variant
	// [ChatViewModelForAgentAssistantMessageFunctionCallMessage].
	PluginMemo string `json:"pluginMemo"`
	// This field is from variant [ChatViewModelForAgentAssistantMessageQnaMessage].
	Chunks []ChatViewModelChunk `json:"chunks"`
	// This field is from variant [ChatViewModelForAgentAssistantMessageQnaMessage].
	Type string `json:"type"`
	// This field is from variant [ChatViewModelForAgentAssistantMessageQnaMessage].
	RetrievedChunkIDs []string `json:"retrievedChunkIds"`
	JSON              struct {
		AssistantMessageType  respjson.Field
		BaseMessage           respjson.Field
		FunctionCallArguments respjson.Field
		FunctionCallID        respjson.Field
		FunctionDescription   respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		PluginMemo            respjson.Field
		Chunks                respjson.Field
		Type                  respjson.Field
		RetrievedChunkIDs     respjson.Field
		raw                   string
	} `json:"-"`
}

func (u ChatViewModelForAgentAssistantMessageUnion) AsChatViewModelForAgentAssistantMessageTextMessage() (v ChatViewModelForAgentAssistantMessageTextMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatViewModelForAgentAssistantMessageUnion) AsChatViewModelForAgentAssistantMessageFunctionCallMessage() (v ChatViewModelForAgentAssistantMessageFunctionCallMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatViewModelForAgentAssistantMessageUnion) AsChatViewModelForAgentAssistantMessageQnaMessage() (v ChatViewModelForAgentAssistantMessageQnaMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatViewModelForAgentAssistantMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatViewModelForAgentAssistantMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatViewModelForAgentAssistantMessageTextMessage struct {
	AssistantMessageType any         `json:"assistantMessageType" api:"required"`
	BaseMessage          BaseMessage `json:"baseMessage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantMessageType respjson.Field
		BaseMessage          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatViewModelForAgentAssistantMessageTextMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatViewModelForAgentAssistantMessageTextMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatViewModelForAgentAssistantMessageFunctionCallMessage struct {
	AssistantMessageType  any         `json:"assistantMessageType" api:"required"`
	BaseMessage           BaseMessage `json:"baseMessage" api:"required"`
	FunctionCallArguments string      `json:"functionCallArguments" api:"required"`
	FunctionCallID        string      `json:"functionCallId" api:"required"`
	FunctionDescription   string      `json:"functionDescription" api:"required"`
	FunctionID            string      `json:"functionId" api:"required"`
	FunctionName          string      `json:"functionName" api:"required"`
	PluginMemo            string      `json:"pluginMemo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantMessageType  respjson.Field
		BaseMessage           respjson.Field
		FunctionCallArguments respjson.Field
		FunctionCallID        respjson.Field
		FunctionDescription   respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		PluginMemo            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatViewModelForAgentAssistantMessageFunctionCallMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatViewModelForAgentAssistantMessageFunctionCallMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatViewModelForAgentAssistantMessageQnaMessage struct {
	AssistantMessageType any                  `json:"assistantMessageType" api:"required"`
	BaseMessage          BaseMessage          `json:"baseMessage" api:"required"`
	Chunks               []ChatViewModelChunk `json:"chunks" api:"required"`
	// Any of "valid", "piiDetected", "sessionExpired", "invalidQuestion",
	// "invalidAnswer", "greeting", "outOfDomain", "openAiInternalError", "guide",
	// "functionCallFailed", "noPluginsInUse", "unexpected".
	Type              string   `json:"type" api:"required"`
	RetrievedChunkIDs []string `json:"retrievedChunkIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantMessageType respjson.Field
		BaseMessage          respjson.Field
		Chunks               respjson.Field
		Type                 respjson.Field
		RetrievedChunkIDs    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatViewModelForAgentAssistantMessageQnaMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatViewModelForAgentAssistantMessageQnaMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatViewModelForAgentStatus string

const (
	ChatViewModelForAgentStatusValid               ChatViewModelForAgentStatus = "valid"
	ChatViewModelForAgentStatusPiiDetected         ChatViewModelForAgentStatus = "piiDetected"
	ChatViewModelForAgentStatusSessionExpired      ChatViewModelForAgentStatus = "sessionExpired"
	ChatViewModelForAgentStatusInvalidQuestion     ChatViewModelForAgentStatus = "invalidQuestion"
	ChatViewModelForAgentStatusInvalidAnswer       ChatViewModelForAgentStatus = "invalidAnswer"
	ChatViewModelForAgentStatusGreeting            ChatViewModelForAgentStatus = "greeting"
	ChatViewModelForAgentStatusOutOfDomain         ChatViewModelForAgentStatus = "outOfDomain"
	ChatViewModelForAgentStatusOpenAIInternalError ChatViewModelForAgentStatus = "openAiInternalError"
	ChatViewModelForAgentStatusGuide               ChatViewModelForAgentStatus = "guide"
	ChatViewModelForAgentStatusFunctionCallFailed  ChatViewModelForAgentStatus = "functionCallFailed"
	ChatViewModelForAgentStatusNoPluginsInUse      ChatViewModelForAgentStatus = "noPluginsInUse"
	ChatViewModelForAgentStatusUnexpected          ChatViewModelForAgentStatus = "unexpected"
)

type SearchFilterBody map[string]SearchFilterBody

type QnaApiv1UserChatResponse struct {
	Data QnaApiv1UserChatResponseDataUnion `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserChatResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserChatResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QnaApiv1UserChatResponseDataUnion contains all possible properties and values
// from [QnaApiv1UserChatResponseDataChatViewModelForQna],
// [QnaApiv1UserChatResponseDataObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfQnaApiv1UserChatResponseDataObject]
type QnaApiv1UserChatResponseDataUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfQnaApiv1UserChatResponseDataObject any    `json:",inline"`
	ID                                   string `json:"id"`
	AnswerFormat                         string `json:"answerFormat"`
	// This field is from variant [QnaApiv1UserChatResponseDataChatViewModelForQna].
	BotSaid   string    `json:"botSaid"`
	BotSaidAt time.Time `json:"botSaidAt"`
	// This field is from variant [QnaApiv1UserChatResponseDataChatViewModelForQna].
	Chunks    []ChatViewModelChunk `json:"chunks"`
	SessionID string               `json:"sessionId"`
	// This field is from variant [QnaApiv1UserChatResponseDataChatViewModelForQna].
	Type       string    `json:"type"`
	UserSaid   string    `json:"userSaid"`
	UserSaidAt time.Time `json:"userSaidAt"`
	Variant    any       `json:"variant"`
	// This field is from variant [QnaApiv1UserChatResponseDataChatViewModelForQna].
	Citations []string `json:"citations"`
	// This field is from variant [QnaApiv1UserChatResponseDataChatViewModelForQna].
	RetrievedChunkIDs []string `json:"retrievedChunkIds"`
	// This field is from variant [QnaApiv1UserChatResponseDataObject].
	AssistantMessages []ChatViewModelForAgentAssistantMessageUnion `json:"assistantMessages"`
	// This field is from variant [QnaApiv1UserChatResponseDataObject].
	IsInterim bool `json:"isInterim"`
	// This field is from variant [QnaApiv1UserChatResponseDataObject].
	Status ChatViewModelForAgentStatus `json:"status"`
	JSON   struct {
		OfQnaApiv1UserChatResponseDataObject respjson.Field
		ID                                   respjson.Field
		AnswerFormat                         respjson.Field
		BotSaid                              respjson.Field
		BotSaidAt                            respjson.Field
		Chunks                               respjson.Field
		SessionID                            respjson.Field
		Type                                 respjson.Field
		UserSaid                             respjson.Field
		UserSaidAt                           respjson.Field
		Variant                              respjson.Field
		Citations                            respjson.Field
		RetrievedChunkIDs                    respjson.Field
		AssistantMessages                    respjson.Field
		IsInterim                            respjson.Field
		Status                               respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u QnaApiv1UserChatResponseDataUnion) AsQnaApiv1UserChatResponseDataChatViewModelForQna() (v QnaApiv1UserChatResponseDataChatViewModelForQna) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QnaApiv1UserChatResponseDataUnion) AsQnaApiv1UserChatResponseDataObject() (v QnaApiv1UserChatResponseDataObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QnaApiv1UserChatResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *QnaApiv1UserChatResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserChatResponseDataChatViewModelForQna struct {
	ID string `json:"id" api:"required"`
	// Any of "plain", "markdown".
	AnswerFormat string               `json:"answerFormat" api:"required"`
	BotSaid      string               `json:"botSaid" api:"required"`
	BotSaidAt    time.Time            `json:"botSaidAt" api:"required" format:"date-time"`
	Chunks       []ChatViewModelChunk `json:"chunks" api:"required"`
	SessionID    string               `json:"sessionId" api:"required"`
	// Any of "valid", "piiDetected", "sessionExpired", "invalidQuestion",
	// "invalidAnswer", "greeting", "outOfDomain", "openAiInternalError", "guide",
	// "functionCallFailed", "noPluginsInUse", "unexpected".
	Type              string         `json:"type" api:"required"`
	UserSaid          string         `json:"userSaid" api:"required"`
	UserSaidAt        time.Time      `json:"userSaidAt" api:"required" format:"date-time"`
	Variant           any            `json:"variant" api:"required"`
	Citations         []string       `json:"citations"`
	RetrievedChunkIDs []string       `json:"retrievedChunkIds"`
	ExtraFields       map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AnswerFormat      respjson.Field
		BotSaid           respjson.Field
		BotSaidAt         respjson.Field
		Chunks            respjson.Field
		SessionID         respjson.Field
		Type              respjson.Field
		UserSaid          respjson.Field
		UserSaidAt        respjson.Field
		Variant           respjson.Field
		Citations         respjson.Field
		RetrievedChunkIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserChatResponseDataChatViewModelForQna) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserChatResponseDataChatViewModelForQna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserChatResponseDataObject struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ChatViewModelForAgent
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserChatResponseDataObject) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserChatResponseDataObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserGetBotProfileResponse struct {
	Data QnaApiv1UserGetBotProfileResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserGetBotProfileResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserGetBotProfileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserGetBotProfileResponseData struct {
	ImageURL    string         `json:"imageUrl"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserGetBotProfileResponseData) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserGetBotProfileResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserGetChatLimitStatusResponse struct {
	Data QnaApiv1UserGetChatLimitStatusResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserGetChatLimitStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserGetChatLimitStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserGetChatLimitStatusResponseData struct {
	Activated        bool                                                       `json:"activated" api:"required"`
	ProjectID        string                                                     `json:"projectId" api:"required"`
	UserID           string                                                     `json:"userId" api:"required"`
	ActiveStatusInfo QnaApiv1UserGetChatLimitStatusResponseDataActiveStatusInfo `json:"activeStatusInfo"`
	ExtraFields      map[string]any                                             `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Activated        respjson.Field
		ProjectID        respjson.Field
		UserID           respjson.Field
		ActiveStatusInfo respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserGetChatLimitStatusResponseData) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserGetChatLimitStatusResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserGetChatLimitStatusResponseDataActiveStatusInfo struct {
	Limit     float64 `json:"limit" api:"required"`
	Message   string  `json:"message" api:"required"`
	Remaining float64 `json:"remaining" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Message     respjson.Field
		Remaining   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserGetChatLimitStatusResponseDataActiveStatusInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *QnaApiv1UserGetChatLimitStatusResponseDataActiveStatusInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserGetWelcomeMessageResponse struct {
	Data QnaApiv1UserGetWelcomeMessageResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserGetWelcomeMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserGetWelcomeMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserGetWelcomeMessageResponseData struct {
	ID                string         `json:"id" api:"required"`
	AllowGenerative   bool           `json:"allowGenerative" api:"required"`
	CreatedAt         time.Time      `json:"createdAt" api:"required" format:"date-time"`
	Message           string         `json:"message" api:"required"`
	UpdatedAt         time.Time      `json:"updatedAt" api:"required" format:"date-time"`
	UseGenerative     bool           `json:"useGenerative" api:"required"`
	GenerativeMessage string         `json:"generativeMessage"`
	ExtraFields       map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AllowGenerative   respjson.Field
		CreatedAt         respjson.Field
		Message           respjson.Field
		UpdatedAt         respjson.Field
		UseGenerative     respjson.Field
		GenerativeMessage respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserGetWelcomeMessageResponseData) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserGetWelcomeMessageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserListChatHistoriesResponse struct {
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
func (r QnaApiv1UserListChatHistoriesResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserListChatHistoriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserStreamingChatResponse struct {
	Data QnaApiv1UserStreamingChatResponseDataUnion `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserStreamingChatResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserStreamingChatResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QnaApiv1UserStreamingChatResponseDataUnion contains all possible properties and
// values from
// [QnaApiv1UserStreamingChatResponseDataStreamingChatFinalViewModelForQna],
// [QnaApiv1UserStreamingChatResponseDataStreamingChatInterimViewModelForQna],
// [QnaApiv1UserStreamingChatResponseDataObject],
// [QnaApiv1UserStreamingChatResponseDataStreamingChatFailureViewModel].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfQnaApiv1UserStreamingChatResponseDataObject]
type QnaApiv1UserStreamingChatResponseDataUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfQnaApiv1UserStreamingChatResponseDataObject any       `json:",inline"`
	ID                                            string    `json:"id"`
	AnswerFormat                                  string    `json:"answerFormat"`
	BotSaid                                       string    `json:"botSaid"`
	BotSaidAt                                     time.Time `json:"botSaidAt"`
	// This field is from variant
	// [QnaApiv1UserStreamingChatResponseDataStreamingChatFinalViewModelForQna].
	Chunks    []ChatViewModelChunk `json:"chunks"`
	IsInterim bool                 `json:"isInterim"`
	SessionID string               `json:"sessionId"`
	// This field is from variant
	// [QnaApiv1UserStreamingChatResponseDataStreamingChatFinalViewModelForQna].
	Type       string    `json:"type"`
	UserSaid   string    `json:"userSaid"`
	UserSaidAt time.Time `json:"userSaidAt"`
	Variant    any       `json:"variant"`
	// This field is from variant
	// [QnaApiv1UserStreamingChatResponseDataStreamingChatFinalViewModelForQna].
	Citations []string `json:"citations"`
	// This field is from variant
	// [QnaApiv1UserStreamingChatResponseDataStreamingChatFinalViewModelForQna].
	RetrievedChunkIDs []string `json:"retrievedChunkIds"`
	// This field is from variant [QnaApiv1UserStreamingChatResponseDataObject].
	AssistantMessages []ChatViewModelForAgentAssistantMessageUnion `json:"assistantMessages"`
	// This field is from variant [QnaApiv1UserStreamingChatResponseDataObject].
	Status ChatViewModelForAgentStatus `json:"status"`
	// This field is from variant
	// [QnaApiv1UserStreamingChatResponseDataStreamingChatFailureViewModel].
	Error Error `json:"error"`
	// This field is from variant
	// [QnaApiv1UserStreamingChatResponseDataStreamingChatFailureViewModel].
	ResultCode float64 `json:"resultCode"`
	JSON       struct {
		OfQnaApiv1UserStreamingChatResponseDataObject respjson.Field
		ID                                            respjson.Field
		AnswerFormat                                  respjson.Field
		BotSaid                                       respjson.Field
		BotSaidAt                                     respjson.Field
		Chunks                                        respjson.Field
		IsInterim                                     respjson.Field
		SessionID                                     respjson.Field
		Type                                          respjson.Field
		UserSaid                                      respjson.Field
		UserSaidAt                                    respjson.Field
		Variant                                       respjson.Field
		Citations                                     respjson.Field
		RetrievedChunkIDs                             respjson.Field
		AssistantMessages                             respjson.Field
		Status                                        respjson.Field
		Error                                         respjson.Field
		ResultCode                                    respjson.Field
		raw                                           string
	} `json:"-"`
}

func (u QnaApiv1UserStreamingChatResponseDataUnion) AsQnaApiv1UserStreamingChatResponseDataStreamingChatFinalViewModelForQna() (v QnaApiv1UserStreamingChatResponseDataStreamingChatFinalViewModelForQna) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QnaApiv1UserStreamingChatResponseDataUnion) AsQnaApiv1UserStreamingChatResponseDataStreamingChatInterimViewModelForQna() (v QnaApiv1UserStreamingChatResponseDataStreamingChatInterimViewModelForQna) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QnaApiv1UserStreamingChatResponseDataUnion) AsQnaApiv1UserStreamingChatResponseDataObject() (v QnaApiv1UserStreamingChatResponseDataObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QnaApiv1UserStreamingChatResponseDataUnion) AsQnaApiv1UserStreamingChatResponseDataStreamingChatFailureViewModel() (v QnaApiv1UserStreamingChatResponseDataStreamingChatFailureViewModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QnaApiv1UserStreamingChatResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *QnaApiv1UserStreamingChatResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserStreamingChatResponseDataStreamingChatFinalViewModelForQna struct {
	ID string `json:"id" api:"required"`
	// Any of "plain", "markdown".
	AnswerFormat string               `json:"answerFormat" api:"required"`
	BotSaid      string               `json:"botSaid" api:"required"`
	BotSaidAt    time.Time            `json:"botSaidAt" api:"required" format:"date-time"`
	Chunks       []ChatViewModelChunk `json:"chunks" api:"required"`
	IsInterim    bool                 `json:"isInterim" api:"required"`
	SessionID    string               `json:"sessionId" api:"required"`
	// Any of "valid", "piiDetected", "sessionExpired", "invalidQuestion",
	// "invalidAnswer", "greeting", "outOfDomain", "openAiInternalError", "guide",
	// "functionCallFailed", "noPluginsInUse", "unexpected".
	Type              string         `json:"type" api:"required"`
	UserSaid          string         `json:"userSaid" api:"required"`
	UserSaidAt        time.Time      `json:"userSaidAt" api:"required" format:"date-time"`
	Variant           any            `json:"variant" api:"required"`
	Citations         []string       `json:"citations"`
	RetrievedChunkIDs []string       `json:"retrievedChunkIds"`
	ExtraFields       map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AnswerFormat      respjson.Field
		BotSaid           respjson.Field
		BotSaidAt         respjson.Field
		Chunks            respjson.Field
		IsInterim         respjson.Field
		SessionID         respjson.Field
		Type              respjson.Field
		UserSaid          respjson.Field
		UserSaidAt        respjson.Field
		Variant           respjson.Field
		Citations         respjson.Field
		RetrievedChunkIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserStreamingChatResponseDataStreamingChatFinalViewModelForQna) RawJSON() string {
	return r.JSON.raw
}
func (r *QnaApiv1UserStreamingChatResponseDataStreamingChatFinalViewModelForQna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserStreamingChatResponseDataStreamingChatInterimViewModelForQna struct {
	// Any of "plain", "markdown".
	AnswerFormat string         `json:"answerFormat" api:"required"`
	IsInterim    bool           `json:"isInterim" api:"required"`
	Variant      any            `json:"variant" api:"required"`
	BotSaid      string         `json:"botSaid"`
	ExtraFields  map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnswerFormat respjson.Field
		IsInterim    respjson.Field
		Variant      respjson.Field
		BotSaid      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserStreamingChatResponseDataStreamingChatInterimViewModelForQna) RawJSON() string {
	return r.JSON.raw
}
func (r *QnaApiv1UserStreamingChatResponseDataStreamingChatInterimViewModelForQna) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserStreamingChatResponseDataObject struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ChatViewModelForAgent
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserStreamingChatResponseDataObject) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1UserStreamingChatResponseDataObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1UserStreamingChatResponseDataStreamingChatFailureViewModel struct {
	Error       Error          `json:"error" api:"required"`
	ResultCode  float64        `json:"resultCode" api:"required"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ResultCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1UserStreamingChatResponseDataStreamingChatFailureViewModel) RawJSON() string {
	return r.JSON.raw
}
func (r *QnaApiv1UserStreamingChatResponseDataStreamingChatFailureViewModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaAPIV1UserChatParams struct {
	// 사용자의 발화
	UserSay   string            `json:"userSay" api:"required"`
	XAPIKey   string            `header:"x-api-key" api:"required" json:"-"`
	XUserID   string            `header:"x-user-id" api:"required" json:"-"`
	SessionID param.Opt[string] `json:"sessionId,omitzero"`
	// 문서 검색 필터
	DocumentFilter SearchFilterBody `json:"documentFilter,omitzero"`
	paramObj
}

func (r QnaAPIV1UserChatParams) MarshalJSON() (data []byte, err error) {
	type shadow QnaAPIV1UserChatParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QnaAPIV1UserChatParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaAPIV1UserGetBotProfileParams struct {
	XAPIKey string `header:"x-api-key" api:"required" json:"-"`
	XUserID string `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

type QnaAPIV1UserGetChatLimitStatusParams struct {
	XAPIKey string `header:"x-api-key" api:"required" json:"-"`
	XUserID string `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

type QnaAPIV1UserGetWelcomeMessageParams struct {
	XAPIKey string `header:"x-api-key" api:"required" json:"-"`
	XUserID string `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

type QnaAPIV1UserListChatHistoriesParams struct {
	XAPIKey string `header:"x-api-key" api:"required" json:"-"`
	XUserID string `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

type QnaAPIV1UserStreamingChatParams struct {
	// 사용자의 발화
	UserSay   string            `json:"userSay" api:"required"`
	XAPIKey   string            `header:"x-api-key" api:"required" json:"-"`
	XUserID   string            `header:"x-user-id" api:"required" json:"-"`
	SessionID param.Opt[string] `json:"sessionId,omitzero"`
	// 문서 검색 필터
	DocumentFilter SearchFilterBody `json:"documentFilter,omitzero"`
	paramObj
}

func (r QnaAPIV1UserStreamingChatParams) MarshalJSON() (data []byte, err error) {
	type shadow QnaAPIV1UserStreamingChatParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QnaAPIV1UserStreamingChatParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

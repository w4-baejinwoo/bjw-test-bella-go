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
	shimjson "github.com/w4-baejinwoo/bjw-test-bella-go/internal/encoding/json"
	"github.com/w4-baejinwoo/bjw-test-bella-go/internal/requestconfig"
	"github.com/w4-baejinwoo/bjw-test-bella-go/option"
	"github.com/w4-baejinwoo/bjw-test-bella-go/packages/param"
	"github.com/w4-baejinwoo/bjw-test-bella-go/packages/respjson"
)

// QnaAPIV1DocumentService contains methods and other services that help with
// interacting with the bjw-test-bella API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQnaAPIV1DocumentService] method instead.
type QnaAPIV1DocumentService struct {
	options []option.RequestOption
}

// NewQnaAPIV1DocumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewQnaAPIV1DocumentService(opts ...option.RequestOption) (r QnaAPIV1DocumentService) {
	r = QnaAPIV1DocumentService{}
	r.options = opts
	return
}

// Add documents
func (r *QnaAPIV1DocumentService) BulkNew(ctx context.Context, params QnaAPIV1DocumentBulkNewParams, opts ...option.RequestOption) (res *QnaApiv1DocumentBulkNewResponse, err error) {
	if !param.IsOmitted(params.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", params.XAPIKey)))
	}
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	path := "qna/api/v1/documents/bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Delete one or more documents
func (r *QnaAPIV1DocumentService) BulkDelete(ctx context.Context, params QnaAPIV1DocumentBulkDeleteParams, opts ...option.RequestOption) (res *QnaApiv1DocumentBulkDeleteResponse, err error) {
	if !param.IsOmitted(params.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", params.XAPIKey)))
	}
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	path := "qna/api/v1/documents/bulk-delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List documents by trace ID
func (r *QnaAPIV1DocumentService) ListByTraceID(ctx context.Context, params QnaAPIV1DocumentListByTraceIDParams, opts ...option.RequestOption) (res *QnaApiv1DocumentListByTraceIDResponse, err error) {
	if !param.IsOmitted(params.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", params.XAPIKey)))
	}
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	path := "qna/api/v1/documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List document chunks
func (r *QnaAPIV1DocumentService) ListChunks(ctx context.Context, params QnaAPIV1DocumentListChunksParams, opts ...option.RequestOption) (res *QnaApiv1DocumentListChunksResponse, err error) {
	if !param.IsOmitted(params.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", params.XAPIKey)))
	}
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	path := "qna/api/v1/documents/chunks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Update document extra data
func (r *QnaAPIV1DocumentService) UpdateExtraData(ctx context.Context, id string, params QnaAPIV1DocumentUpdateExtraDataParams, opts ...option.RequestOption) (res *QnaApiv1DocumentUpdateExtraDataResponse, err error) {
	if !param.IsOmitted(params.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%v", params.XAPIKey)))
	}
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID)))
	}
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("qna/api/v1/documents/%s/extra", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type EnvelopedViewModel struct {
	Meta  EnvelopedViewModelMeta `json:"meta" api:"required"`
	Data  map[string]any         `json:"data"`
	Error Error                  `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Meta        respjson.Field
		Data        respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnvelopedViewModel) RawJSON() string { return r.JSON.raw }
func (r *EnvelopedViewModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnvelopedViewModelMeta struct {
	IsSuccessful bool    `json:"isSuccessful" api:"required"`
	ResultCode   float64 `json:"resultCode" api:"required"`
	Total        float64 `json:"total"`
	TraceID      string  `json:"traceId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSuccessful respjson.Field
		ResultCode   respjson.Field
		Total        respjson.Field
		TraceID      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnvelopedViewModelMeta) RawJSON() string { return r.JSON.raw }
func (r *EnvelopedViewModelMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Error map[string]Error

type QnaApiv1DocumentBulkNewResponse struct {
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
func (r QnaApiv1DocumentBulkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1DocumentBulkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1DocumentBulkDeleteResponse struct {
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
func (r QnaApiv1DocumentBulkDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1DocumentBulkDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1DocumentListByTraceIDResponse struct {
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
func (r QnaApiv1DocumentListByTraceIDResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1DocumentListByTraceIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1DocumentListChunksResponse struct {
	Data QnaApiv1DocumentListChunksResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EnvelopedViewModel
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1DocumentListChunksResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1DocumentListChunksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1DocumentListChunksResponseData struct {
	Chunks      []QnaApiv1DocumentListChunksResponseDataChunk `json:"chunks" api:"required"`
	ExtraFields map[string]any                                `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chunks      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1DocumentListChunksResponseData) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1DocumentListChunksResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1DocumentListChunksResponseDataChunk struct {
	ID         string         `json:"id" api:"required"`
	Contents   string         `json:"contents" api:"required"`
	DocumentID string         `json:"documentId" api:"required"`
	Offset     float64        `json:"offset" api:"required"`
	Page       float64        `json:"page" api:"required"`
	Extra      map[string]any `json:"extra"`
	Link       string         `json:"link"`
	Title      string         `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Contents    respjson.Field
		DocumentID  respjson.Field
		Offset      respjson.Field
		Page        respjson.Field
		Extra       respjson.Field
		Link        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QnaApiv1DocumentListChunksResponseDataChunk) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1DocumentListChunksResponseDataChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaApiv1DocumentUpdateExtraDataResponse struct {
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
func (r QnaApiv1DocumentUpdateExtraDataResponse) RawJSON() string { return r.JSON.raw }
func (r *QnaApiv1DocumentUpdateExtraDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaAPIV1DocumentBulkNewParams struct {
	Documents []QnaAPIV1DocumentBulkNewParamsDocument `json:"documents,omitzero" api:"required"`
	XAPIKey   string                                  `header:"x-api-key" api:"required" json:"-"`
	XUserID   string                                  `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

func (r QnaAPIV1DocumentBulkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow QnaAPIV1DocumentBulkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QnaAPIV1DocumentBulkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Type are required.
type QnaAPIV1DocumentBulkNewParamsDocument struct {
	ID string `json:"id" api:"required"`
	// Any of "plain", "file", "url", "html".
	Type string `json:"type,omitzero" api:"required"`
	// 파일 경로 (value 또는 filePath 중 하나는 필수)
	FilePath param.Opt[string] `json:"filePath,omitzero"`
	// 문서의 제목
	Title   param.Opt[string] `json:"title,omitzero"`
	TraceID param.Opt[string] `json:"traceId,omitzero"`
	// 문서 요약 생성 여부주어지지 않은 경우 false 로 처리됨
	UseSummary param.Opt[bool] `json:"useSummary,omitzero"`
	// 문서의 내용. filePath 가 주어지면 생략 가능함
	Value param.Opt[string] `json:"value,omitzero"`
	// 문서와 연관하여 저장할 임의의 데이터 document 생성 시 id 와 함께 그대로 반환해줌
	// key-value 형태로 저장됨, key 와 value 는 string 타입만 가능
	Extra map[string]any `json:"extra,omitzero"`
	// 문서의 파싱 설정주어지지 않은 경우 기본 파서로 처리됨
	ParsingOptions QnaAPIV1DocumentBulkNewParamsDocumentParsingOptions `json:"parsingOptions,omitzero"`
	Ref            QnaAPIV1DocumentBulkNewParamsDocumentRef            `json:"ref,omitzero"`
	// 문서의 유효기간
	Validity QnaAPIV1DocumentBulkNewParamsDocumentValidity `json:"validity,omitzero"`
	paramObj
}

func (r QnaAPIV1DocumentBulkNewParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow QnaAPIV1DocumentBulkNewParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QnaAPIV1DocumentBulkNewParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[QnaAPIV1DocumentBulkNewParamsDocument](
		"type", "plain", "file", "url", "html",
	)
}

// 문서의 파싱 설정주어지지 않은 경우 기본 파서로 처리됨
//
// The property Parser is required.
type QnaAPIV1DocumentBulkNewParamsDocumentParsingOptions struct {
	Parser string `json:"parser" api:"required"`
	paramObj
}

func (r QnaAPIV1DocumentBulkNewParamsDocumentParsingOptions) MarshalJSON() (data []byte, err error) {
	type shadow QnaAPIV1DocumentBulkNewParamsDocumentParsingOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QnaAPIV1DocumentBulkNewParamsDocumentParsingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Link is required.
type QnaAPIV1DocumentBulkNewParamsDocumentRef struct {
	Link string `json:"link" api:"required"`
	paramObj
}

func (r QnaAPIV1DocumentBulkNewParamsDocumentRef) MarshalJSON() (data []byte, err error) {
	type shadow QnaAPIV1DocumentBulkNewParamsDocumentRef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QnaAPIV1DocumentBulkNewParamsDocumentRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 문서의 유효기간
type QnaAPIV1DocumentBulkNewParamsDocumentValidity struct {
	EndedAt   param.Opt[time.Time] `json:"endedAt,omitzero" format:"date-time"`
	StartedAt param.Opt[time.Time] `json:"startedAt,omitzero" format:"date-time"`
	paramObj
}

func (r QnaAPIV1DocumentBulkNewParamsDocumentValidity) MarshalJSON() (data []byte, err error) {
	type shadow QnaAPIV1DocumentBulkNewParamsDocumentValidity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QnaAPIV1DocumentBulkNewParamsDocumentValidity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaAPIV1DocumentBulkDeleteParams struct {
	DocumentIDs []string `json:"documentIds,omitzero" api:"required"`
	XAPIKey     string   `header:"x-api-key" api:"required" json:"-"`
	XUserID     string   `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

func (r QnaAPIV1DocumentBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow QnaAPIV1DocumentBulkDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QnaAPIV1DocumentBulkDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QnaAPIV1DocumentListByTraceIDParams struct {
	XAPIKey string            `header:"x-api-key" api:"required" json:"-"`
	XUserID string            `header:"x-user-id" api:"required" json:"-"`
	TraceID param.Opt[string] `query:"traceId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QnaAPIV1DocumentListByTraceIDParams]'s query parameters as
// `url.Values`.
func (r QnaAPIV1DocumentListByTraceIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QnaAPIV1DocumentListChunksParams struct {
	XAPIKey     string   `header:"x-api-key" api:"required" json:"-"`
	XUserID     string   `header:"x-user-id" api:"required" json:"-"`
	ChunkIDs    []string `query:"chunkIds,omitzero" json:"-"`
	DocumentIDs []string `query:"documentIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QnaAPIV1DocumentListChunksParams]'s query parameters as
// `url.Values`.
func (r QnaAPIV1DocumentListChunksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QnaAPIV1DocumentUpdateExtraDataParams struct {
	Body    map[string]any
	XAPIKey string `header:"x-api-key" api:"required" json:"-"`
	XUserID string `header:"x-user-id" api:"required" json:"-"`
	paramObj
}

func (r QnaAPIV1DocumentUpdateExtraDataParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *QnaAPIV1DocumentUpdateExtraDataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

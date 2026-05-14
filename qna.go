// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bjwtestbella

import (
	"github.com/w4-baejinwoo/bjw-test-bella-go/option"
)

// QnaService contains methods and other services that help with interacting with
// the bjw-test-bella API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQnaService] method instead.
type QnaService struct {
	options []option.RequestOption
	API     QnaAPIService
}

// NewQnaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQnaService(opts ...option.RequestOption) (r QnaService) {
	r = QnaService{}
	r.options = opts
	r.API = NewQnaAPIService(opts...)
	return
}

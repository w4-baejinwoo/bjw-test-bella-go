// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bjwtestbella

import (
	"github.com/w4-baejinwoo/bjw-test-bella-go/option"
)

// QnaAPIService contains methods and other services that help with interacting
// with the bjw-test-bella API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQnaAPIService] method instead.
type QnaAPIService struct {
	options []option.RequestOption
	V1      QnaAPIV1Service
}

// NewQnaAPIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQnaAPIService(opts ...option.RequestOption) (r QnaAPIService) {
	r = QnaAPIService{}
	r.options = opts
	r.V1 = NewQnaAPIV1Service(opts...)
	return
}

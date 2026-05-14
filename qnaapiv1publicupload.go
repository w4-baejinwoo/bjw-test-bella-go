// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bjwtestbella

import (
	"github.com/stainless-sdks/bjw-test-bella-go/option"
)

// QnaAPIV1PublicUploadService contains methods and other services that help with
// interacting with the bjw-test-bella API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQnaAPIV1PublicUploadService] method instead.
type QnaAPIV1PublicUploadService struct {
	options []option.RequestOption
}

// NewQnaAPIV1PublicUploadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewQnaAPIV1PublicUploadService(opts ...option.RequestOption) (r QnaAPIV1PublicUploadService) {
	r = QnaAPIV1PublicUploadService{}
	r.options = opts
	return
}

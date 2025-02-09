/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// ResponseObject is a runtime object extended with methods to handle response-specific fields.
// +kubebuilder:object:generate=false
type ResponseObject interface {
	runtime.Object
	GetMessage() string
	GetStatus() ResponseStatus
	SetMessage(message string)
	SetStatus(status ResponseStatus)
}

// RetryResponseObject is a ResponseObject which additionally defines the functionality
// for a response to signal a retry.
// +kubebuilder:object:generate=false
type RetryResponseObject interface {
	ResponseObject
	GetRetryAfterSeconds() int32
	SetRetryAfterSeconds(retryAfterSeconds int32)
}

// CommonResponse is the data structure common to all response types.
type CommonResponse struct {
	// Status of the call. One of "Success" or "Failure".
	Status ResponseStatus `json:"status"`

	// A human-readable description of the status of the call.
	Message string `json:"message"`
}

// SetMessage sets the message field for the CommonResponse.
func (r *CommonResponse) SetMessage(message string) {
	r.Message = message
}

// SetStatus sets the status field for the CommonResponse.
func (r *CommonResponse) SetStatus(status ResponseStatus) {
	r.Status = status
}

// GetMessage returns the Message field for the CommonResponse.
func (r *CommonResponse) GetMessage() string {
	return r.Message
}

// GetStatus returns the Status field for the CommonResponse.
func (r *CommonResponse) GetStatus() ResponseStatus {
	return r.Status
}

// ResponseStatus represents the status of the hook response.
// +enum
type ResponseStatus string

const (
	// ResponseStatusSuccess represents the success response.
	ResponseStatusSuccess ResponseStatus = "Success"

	// ResponseStatusFailure represents a failure response.
	ResponseStatusFailure ResponseStatus = "Failure"
)

// CommonRetryResponse is the data structure which contains all
// common retry fields.
type CommonRetryResponse struct {
	// CommonResponse contains Status and Message fields common to all response types.
	CommonResponse `json:",inline"`

	// RetryAfterSeconds when set to a non-zero value signifies that the hook
	// will be called again at a future time.
	RetryAfterSeconds int32 `json:"retryAfterSeconds"`
}

// GetRetryAfterSeconds sets the RetryAfterSeconds value.
func (r *CommonRetryResponse) GetRetryAfterSeconds() int32 {
	return r.RetryAfterSeconds
}

// SetRetryAfterSeconds returns the RetryAfterSeconds value.
func (r *CommonRetryResponse) SetRetryAfterSeconds(retryAfterSeconds int32) {
	r.RetryAfterSeconds = retryAfterSeconds
}

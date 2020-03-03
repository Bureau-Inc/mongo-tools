// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ioteventsdata

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// An internal failure occured.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request was invalid.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is currently unavailable.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request could not be completed due to throttling.
	ErrCodeThrottlingException = "ThrottlingException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalFailureException":    newErrorInternalFailureException,
	"InvalidRequestException":     newErrorInvalidRequestException,
	"ResourceNotFoundException":   newErrorResourceNotFoundException,
	"ServiceUnavailableException": newErrorServiceUnavailableException,
	"ThrottlingException":         newErrorThrottlingException,
}
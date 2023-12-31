// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/streaming/v1alpha1/cover.proto

package streaming

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateCoverRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCoverRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCoverRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCoverRequestMultiError, or nil if none found.
func (m *CreateCoverRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCoverRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FileId

	// no validation rules for Width

	// no validation rules for Height

	if len(errors) > 0 {
		return CreateCoverRequestMultiError(errors)
	}

	return nil
}

// CreateCoverRequestMultiError is an error wrapping multiple validation errors
// returned by CreateCoverRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateCoverRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCoverRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCoverRequestMultiError) AllErrors() []error { return m }

// CreateCoverRequestValidationError is the validation error returned by
// CreateCoverRequest.Validate if the designated constraints aren't met.
type CreateCoverRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCoverRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCoverRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCoverRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCoverRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCoverRequestValidationError) ErrorName() string {
	return "CreateCoverRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCoverRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCoverRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCoverRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCoverRequestValidationError{}

// Validate checks the field values on CreateCoverResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCoverResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCoverResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCoverResponseMultiError, or nil if none found.
func (m *CreateCoverResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCoverResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetCover()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCoverResponseValidationError{
					field:  "Cover",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCoverResponseValidationError{
					field:  "Cover",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCover()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCoverResponseValidationError{
				field:  "Cover",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateCoverResponseMultiError(errors)
	}

	return nil
}

// CreateCoverResponseMultiError is an error wrapping multiple validation
// errors returned by CreateCoverResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateCoverResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCoverResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCoverResponseMultiError) AllErrors() []error { return m }

// CreateCoverResponseValidationError is the validation error returned by
// CreateCoverResponse.Validate if the designated constraints aren't met.
type CreateCoverResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCoverResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCoverResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCoverResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCoverResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCoverResponseValidationError) ErrorName() string {
	return "CreateCoverResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCoverResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCoverResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCoverResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCoverResponseValidationError{}

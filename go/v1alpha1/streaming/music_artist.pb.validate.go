// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/streaming/v1alpha1/music_artist.proto

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

// Validate checks the field values on CreateArtistRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArtistRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArtistRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArtistRequestMultiError, or nil if none found.
func (m *CreateArtistRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArtistRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateArtistRequestMultiError(errors)
	}

	return nil
}

// CreateArtistRequestMultiError is an error wrapping multiple validation
// errors returned by CreateArtistRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateArtistRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArtistRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArtistRequestMultiError) AllErrors() []error { return m }

// CreateArtistRequestValidationError is the validation error returned by
// CreateArtistRequest.Validate if the designated constraints aren't met.
type CreateArtistRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArtistRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArtistRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArtistRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArtistRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArtistRequestValidationError) ErrorName() string {
	return "CreateArtistRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArtistRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArtistRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArtistRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArtistRequestValidationError{}

// Validate checks the field values on CreateArtistResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArtistResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArtistResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArtistResponseMultiError, or nil if none found.
func (m *CreateArtistResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArtistResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	for idx, item := range m.GetArtists() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateArtistResponseValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateArtistResponseValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateArtistResponseValidationError{
					field:  fmt.Sprintf("Artists[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateArtistResponseMultiError(errors)
	}

	return nil
}

// CreateArtistResponseMultiError is an error wrapping multiple validation
// errors returned by CreateArtistResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateArtistResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArtistResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArtistResponseMultiError) AllErrors() []error { return m }

// CreateArtistResponseValidationError is the validation error returned by
// CreateArtistResponse.Validate if the designated constraints aren't met.
type CreateArtistResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArtistResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArtistResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArtistResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArtistResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArtistResponseValidationError) ErrorName() string {
	return "CreateArtistResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArtistResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArtistResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArtistResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArtistResponseValidationError{}

// Validate checks the field values on SearchArtistsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SearchArtistsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchArtistsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SearchArtistsRequestMultiError, or nil if none found.
func (m *SearchArtistsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchArtistsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return SearchArtistsRequestMultiError(errors)
	}

	return nil
}

// SearchArtistsRequestMultiError is an error wrapping multiple validation
// errors returned by SearchArtistsRequest.ValidateAll() if the designated
// constraints aren't met.
type SearchArtistsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchArtistsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchArtistsRequestMultiError) AllErrors() []error { return m }

// SearchArtistsRequestValidationError is the validation error returned by
// SearchArtistsRequest.Validate if the designated constraints aren't met.
type SearchArtistsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchArtistsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchArtistsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchArtistsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchArtistsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchArtistsRequestValidationError) ErrorName() string {
	return "SearchArtistsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SearchArtistsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchArtistsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchArtistsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchArtistsRequestValidationError{}

// Validate checks the field values on SearchArtistsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SearchArtistsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchArtistsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SearchArtistsResponseMultiError, or nil if none found.
func (m *SearchArtistsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchArtistsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	for idx, item := range m.GetArtists() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SearchArtistsResponseValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SearchArtistsResponseValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchArtistsResponseValidationError{
					field:  fmt.Sprintf("Artists[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SearchArtistsResponseMultiError(errors)
	}

	return nil
}

// SearchArtistsResponseMultiError is an error wrapping multiple validation
// errors returned by SearchArtistsResponse.ValidateAll() if the designated
// constraints aren't met.
type SearchArtistsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchArtistsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchArtistsResponseMultiError) AllErrors() []error { return m }

// SearchArtistsResponseValidationError is the validation error returned by
// SearchArtistsResponse.Validate if the designated constraints aren't met.
type SearchArtistsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchArtistsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchArtistsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchArtistsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchArtistsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchArtistsResponseValidationError) ErrorName() string {
	return "SearchArtistsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SearchArtistsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchArtistsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchArtistsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchArtistsResponseValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/streaming/v1alpha1/music_album.proto

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

// Validate checks the field values on CreateAlbumRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAlbumRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAlbumRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAlbumRequestMultiError, or nil if none found.
func (m *CreateAlbumRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAlbumRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Year

	// no validation rules for Description

	// no validation rules for CoverId

	if len(errors) > 0 {
		return CreateAlbumRequestMultiError(errors)
	}

	return nil
}

// CreateAlbumRequestMultiError is an error wrapping multiple validation errors
// returned by CreateAlbumRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateAlbumRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAlbumRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAlbumRequestMultiError) AllErrors() []error { return m }

// CreateAlbumRequestValidationError is the validation error returned by
// CreateAlbumRequest.Validate if the designated constraints aren't met.
type CreateAlbumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAlbumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAlbumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAlbumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAlbumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAlbumRequestValidationError) ErrorName() string {
	return "CreateAlbumRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAlbumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAlbumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAlbumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAlbumRequestValidationError{}

// Validate checks the field values on CreateAlbumResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAlbumResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAlbumResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAlbumResponseMultiError, or nil if none found.
func (m *CreateAlbumResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAlbumResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetAlbum()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateAlbumResponseValidationError{
					field:  "Album",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateAlbumResponseValidationError{
					field:  "Album",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlbum()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAlbumResponseValidationError{
				field:  "Album",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateAlbumResponseMultiError(errors)
	}

	return nil
}

// CreateAlbumResponseMultiError is an error wrapping multiple validation
// errors returned by CreateAlbumResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateAlbumResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAlbumResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAlbumResponseMultiError) AllErrors() []error { return m }

// CreateAlbumResponseValidationError is the validation error returned by
// CreateAlbumResponse.Validate if the designated constraints aren't met.
type CreateAlbumResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAlbumResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAlbumResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAlbumResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAlbumResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAlbumResponseValidationError) ErrorName() string {
	return "CreateAlbumResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAlbumResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAlbumResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAlbumResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAlbumResponseValidationError{}

// Validate checks the field values on GetAlbumRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetAlbumRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAlbumRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAlbumRequestMultiError, or nil if none found.
func (m *GetAlbumRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAlbumRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetAlbumRequestMultiError(errors)
	}

	return nil
}

// GetAlbumRequestMultiError is an error wrapping multiple validation errors
// returned by GetAlbumRequest.ValidateAll() if the designated constraints
// aren't met.
type GetAlbumRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAlbumRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAlbumRequestMultiError) AllErrors() []error { return m }

// GetAlbumRequestValidationError is the validation error returned by
// GetAlbumRequest.Validate if the designated constraints aren't met.
type GetAlbumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAlbumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAlbumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAlbumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAlbumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAlbumRequestValidationError) ErrorName() string { return "GetAlbumRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetAlbumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAlbumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAlbumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAlbumRequestValidationError{}

// Validate checks the field values on GetAlbumResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetAlbumResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAlbumResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAlbumResponseMultiError, or nil if none found.
func (m *GetAlbumResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAlbumResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetAlbum()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetAlbumResponseValidationError{
					field:  "Album",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetAlbumResponseValidationError{
					field:  "Album",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlbum()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAlbumResponseValidationError{
				field:  "Album",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetAlbumResponseMultiError(errors)
	}

	return nil
}

// GetAlbumResponseMultiError is an error wrapping multiple validation errors
// returned by GetAlbumResponse.ValidateAll() if the designated constraints
// aren't met.
type GetAlbumResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAlbumResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAlbumResponseMultiError) AllErrors() []error { return m }

// GetAlbumResponseValidationError is the validation error returned by
// GetAlbumResponse.Validate if the designated constraints aren't met.
type GetAlbumResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAlbumResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAlbumResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAlbumResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAlbumResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAlbumResponseValidationError) ErrorName() string { return "GetAlbumResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetAlbumResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAlbumResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAlbumResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAlbumResponseValidationError{}

// Validate checks the field values on EditAlbumRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *EditAlbumRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EditAlbumRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EditAlbumRequestMultiError, or nil if none found.
func (m *EditAlbumRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *EditAlbumRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Year

	// no validation rules for Description

	// no validation rules for CoverId

	if len(errors) > 0 {
		return EditAlbumRequestMultiError(errors)
	}

	return nil
}

// EditAlbumRequestMultiError is an error wrapping multiple validation errors
// returned by EditAlbumRequest.ValidateAll() if the designated constraints
// aren't met.
type EditAlbumRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EditAlbumRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EditAlbumRequestMultiError) AllErrors() []error { return m }

// EditAlbumRequestValidationError is the validation error returned by
// EditAlbumRequest.Validate if the designated constraints aren't met.
type EditAlbumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EditAlbumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EditAlbumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EditAlbumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EditAlbumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EditAlbumRequestValidationError) ErrorName() string { return "EditAlbumRequestValidationError" }

// Error satisfies the builtin error interface
func (e EditAlbumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEditAlbumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EditAlbumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EditAlbumRequestValidationError{}

// Validate checks the field values on EditAlbumResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *EditAlbumResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EditAlbumResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EditAlbumResponseMultiError, or nil if none found.
func (m *EditAlbumResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *EditAlbumResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetAlbum()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EditAlbumResponseValidationError{
					field:  "Album",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EditAlbumResponseValidationError{
					field:  "Album",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlbum()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EditAlbumResponseValidationError{
				field:  "Album",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EditAlbumResponseMultiError(errors)
	}

	return nil
}

// EditAlbumResponseMultiError is an error wrapping multiple validation errors
// returned by EditAlbumResponse.ValidateAll() if the designated constraints
// aren't met.
type EditAlbumResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EditAlbumResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EditAlbumResponseMultiError) AllErrors() []error { return m }

// EditAlbumResponseValidationError is the validation error returned by
// EditAlbumResponse.Validate if the designated constraints aren't met.
type EditAlbumResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EditAlbumResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EditAlbumResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EditAlbumResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EditAlbumResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EditAlbumResponseValidationError) ErrorName() string {
	return "EditAlbumResponseValidationError"
}

// Error satisfies the builtin error interface
func (e EditAlbumResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEditAlbumResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EditAlbumResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EditAlbumResponseValidationError{}

// Validate checks the field values on GetAlbumsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetAlbumsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAlbumsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAlbumsResponseMultiError, or nil if none found.
func (m *GetAlbumsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAlbumsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	for idx, item := range m.GetAlbums() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetAlbumsResponseValidationError{
						field:  fmt.Sprintf("Albums[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetAlbumsResponseValidationError{
						field:  fmt.Sprintf("Albums[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetAlbumsResponseValidationError{
					field:  fmt.Sprintf("Albums[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetAlbumsResponseMultiError(errors)
	}

	return nil
}

// GetAlbumsResponseMultiError is an error wrapping multiple validation errors
// returned by GetAlbumsResponse.ValidateAll() if the designated constraints
// aren't met.
type GetAlbumsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAlbumsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAlbumsResponseMultiError) AllErrors() []error { return m }

// GetAlbumsResponseValidationError is the validation error returned by
// GetAlbumsResponse.Validate if the designated constraints aren't met.
type GetAlbumsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAlbumsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAlbumsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAlbumsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAlbumsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAlbumsResponseValidationError) ErrorName() string {
	return "GetAlbumsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAlbumsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAlbumsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAlbumsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAlbumsResponseValidationError{}

// Validate checks the field values on AddAlbumArtistRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddAlbumArtistRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddAlbumArtistRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddAlbumArtistRequestMultiError, or nil if none found.
func (m *AddAlbumArtistRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddAlbumArtistRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AddAlbumArtistRequestMultiError(errors)
	}

	return nil
}

// AddAlbumArtistRequestMultiError is an error wrapping multiple validation
// errors returned by AddAlbumArtistRequest.ValidateAll() if the designated
// constraints aren't met.
type AddAlbumArtistRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddAlbumArtistRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddAlbumArtistRequestMultiError) AllErrors() []error { return m }

// AddAlbumArtistRequestValidationError is the validation error returned by
// AddAlbumArtistRequest.Validate if the designated constraints aren't met.
type AddAlbumArtistRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddAlbumArtistRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddAlbumArtistRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddAlbumArtistRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddAlbumArtistRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddAlbumArtistRequestValidationError) ErrorName() string {
	return "AddAlbumArtistRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddAlbumArtistRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddAlbumArtistRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddAlbumArtistRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddAlbumArtistRequestValidationError{}

// Validate checks the field values on AddAlbumArtistResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddAlbumArtistResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddAlbumArtistResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddAlbumArtistResponseMultiError, or nil if none found.
func (m *AddAlbumArtistResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddAlbumArtistResponse) validate(all bool) error {
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
					errors = append(errors, AddAlbumArtistResponseValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AddAlbumArtistResponseValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddAlbumArtistResponseValidationError{
					field:  fmt.Sprintf("Artists[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AddAlbumArtistResponseMultiError(errors)
	}

	return nil
}

// AddAlbumArtistResponseMultiError is an error wrapping multiple validation
// errors returned by AddAlbumArtistResponse.ValidateAll() if the designated
// constraints aren't met.
type AddAlbumArtistResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddAlbumArtistResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddAlbumArtistResponseMultiError) AllErrors() []error { return m }

// AddAlbumArtistResponseValidationError is the validation error returned by
// AddAlbumArtistResponse.Validate if the designated constraints aren't met.
type AddAlbumArtistResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddAlbumArtistResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddAlbumArtistResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddAlbumArtistResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddAlbumArtistResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddAlbumArtistResponseValidationError) ErrorName() string {
	return "AddAlbumArtistResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddAlbumArtistResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddAlbumArtistResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddAlbumArtistResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddAlbumArtistResponseValidationError{}

// Validate checks the field values on RemoveAlbumArtistRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveAlbumArtistRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveAlbumArtistRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveAlbumArtistRequestMultiError, or nil if none found.
func (m *RemoveAlbumArtistRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveAlbumArtistRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return RemoveAlbumArtistRequestMultiError(errors)
	}

	return nil
}

// RemoveAlbumArtistRequestMultiError is an error wrapping multiple validation
// errors returned by RemoveAlbumArtistRequest.ValidateAll() if the designated
// constraints aren't met.
type RemoveAlbumArtistRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveAlbumArtistRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveAlbumArtistRequestMultiError) AllErrors() []error { return m }

// RemoveAlbumArtistRequestValidationError is the validation error returned by
// RemoveAlbumArtistRequest.Validate if the designated constraints aren't met.
type RemoveAlbumArtistRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveAlbumArtistRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveAlbumArtistRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveAlbumArtistRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveAlbumArtistRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveAlbumArtistRequestValidationError) ErrorName() string {
	return "RemoveAlbumArtistRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveAlbumArtistRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveAlbumArtistRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveAlbumArtistRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveAlbumArtistRequestValidationError{}

// Validate checks the field values on RemoveAlbumArtistResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveAlbumArtistResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveAlbumArtistResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveAlbumArtistResponseMultiError, or nil if none found.
func (m *RemoveAlbumArtistResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveAlbumArtistResponse) validate(all bool) error {
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
					errors = append(errors, RemoveAlbumArtistResponseValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RemoveAlbumArtistResponseValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RemoveAlbumArtistResponseValidationError{
					field:  fmt.Sprintf("Artists[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RemoveAlbumArtistResponseMultiError(errors)
	}

	return nil
}

// RemoveAlbumArtistResponseMultiError is an error wrapping multiple validation
// errors returned by RemoveAlbumArtistResponse.ValidateAll() if the
// designated constraints aren't met.
type RemoveAlbumArtistResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveAlbumArtistResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveAlbumArtistResponseMultiError) AllErrors() []error { return m }

// RemoveAlbumArtistResponseValidationError is the validation error returned by
// RemoveAlbumArtistResponse.Validate if the designated constraints aren't met.
type RemoveAlbumArtistResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveAlbumArtistResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveAlbumArtistResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveAlbumArtistResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveAlbumArtistResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveAlbumArtistResponseValidationError) ErrorName() string {
	return "RemoveAlbumArtistResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveAlbumArtistResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveAlbumArtistResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveAlbumArtistResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveAlbumArtistResponseValidationError{}
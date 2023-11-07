// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/streaming/v1alpha1/data.proto

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

// Validate checks the field values on MusicData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MusicData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MusicData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MusicDataMultiError, or nil
// if none found.
func (m *MusicData) ValidateAll() error {
	return m.validate(true)
}

func (m *MusicData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MusicDataValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MusicDataValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MusicDataValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MusicDataValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MusicDataValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MusicDataValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	// no validation rules for Duration

	// no validation rules for Description

	if all {
		switch v := interface{}(m.GetFile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MusicDataValidationError{
					field:  "File",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MusicDataValidationError{
					field:  "File",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MusicDataValidationError{
				field:  "File",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAlbums()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MusicDataValidationError{
					field:  "Albums",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MusicDataValidationError{
					field:  "Albums",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlbums()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MusicDataValidationError{
				field:  "Albums",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetArtists() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MusicDataValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MusicDataValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MusicDataValidationError{
					field:  fmt.Sprintf("Artists[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetGenres() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MusicDataValidationError{
						field:  fmt.Sprintf("Genres[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MusicDataValidationError{
						field:  fmt.Sprintf("Genres[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MusicDataValidationError{
					field:  fmt.Sprintf("Genres[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MusicDataMultiError(errors)
	}

	return nil
}

// MusicDataMultiError is an error wrapping multiple validation errors returned
// by MusicData.ValidateAll() if the designated constraints aren't met.
type MusicDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MusicDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MusicDataMultiError) AllErrors() []error { return m }

// MusicDataValidationError is the validation error returned by
// MusicData.Validate if the designated constraints aren't met.
type MusicDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MusicDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MusicDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MusicDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MusicDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MusicDataValidationError) ErrorName() string { return "MusicDataValidationError" }

// Error satisfies the builtin error interface
func (e MusicDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMusicData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MusicDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MusicDataValidationError{}

// Validate checks the field values on MusicAlbumData with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MusicAlbumData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MusicAlbumData with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MusicAlbumDataMultiError,
// or nil if none found.
func (m *MusicAlbumData) ValidateAll() error {
	return m.validate(true)
}

func (m *MusicAlbumData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MusicAlbumDataValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MusicAlbumDataValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MusicAlbumDataValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MusicAlbumDataValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MusicAlbumDataValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MusicAlbumDataValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Title

	// no validation rules for Year

	// no validation rules for Description

	if all {
		switch v := interface{}(m.GetCover()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MusicAlbumDataValidationError{
					field:  "Cover",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MusicAlbumDataValidationError{
					field:  "Cover",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCover()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MusicAlbumDataValidationError{
				field:  "Cover",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetArtists() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MusicAlbumDataValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MusicAlbumDataValidationError{
						field:  fmt.Sprintf("Artists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MusicAlbumDataValidationError{
					field:  fmt.Sprintf("Artists[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MusicAlbumDataMultiError(errors)
	}

	return nil
}

// MusicAlbumDataMultiError is an error wrapping multiple validation errors
// returned by MusicAlbumData.ValidateAll() if the designated constraints
// aren't met.
type MusicAlbumDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MusicAlbumDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MusicAlbumDataMultiError) AllErrors() []error { return m }

// MusicAlbumDataValidationError is the validation error returned by
// MusicAlbumData.Validate if the designated constraints aren't met.
type MusicAlbumDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MusicAlbumDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MusicAlbumDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MusicAlbumDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MusicAlbumDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MusicAlbumDataValidationError) ErrorName() string { return "MusicAlbumDataValidationError" }

// Error satisfies the builtin error interface
func (e MusicAlbumDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMusicAlbumData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MusicAlbumDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MusicAlbumDataValidationError{}

// Validate checks the field values on MusicArtistData with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MusicArtistData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MusicArtistData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MusicArtistDataMultiError, or nil if none found.
func (m *MusicArtistData) ValidateAll() error {
	return m.validate(true)
}

func (m *MusicArtistData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return MusicArtistDataMultiError(errors)
	}

	return nil
}

// MusicArtistDataMultiError is an error wrapping multiple validation errors
// returned by MusicArtistData.ValidateAll() if the designated constraints
// aren't met.
type MusicArtistDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MusicArtistDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MusicArtistDataMultiError) AllErrors() []error { return m }

// MusicArtistDataValidationError is the validation error returned by
// MusicArtistData.Validate if the designated constraints aren't met.
type MusicArtistDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MusicArtistDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MusicArtistDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MusicArtistDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MusicArtistDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MusicArtistDataValidationError) ErrorName() string { return "MusicArtistDataValidationError" }

// Error satisfies the builtin error interface
func (e MusicArtistDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMusicArtistData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MusicArtistDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MusicArtistDataValidationError{}

// Validate checks the field values on MusicGenreData with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MusicGenreData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MusicGenreData with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MusicGenreDataMultiError,
// or nil if none found.
func (m *MusicGenreData) ValidateAll() error {
	return m.validate(true)
}

func (m *MusicGenreData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return MusicGenreDataMultiError(errors)
	}

	return nil
}

// MusicGenreDataMultiError is an error wrapping multiple validation errors
// returned by MusicGenreData.ValidateAll() if the designated constraints
// aren't met.
type MusicGenreDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MusicGenreDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MusicGenreDataMultiError) AllErrors() []error { return m }

// MusicGenreDataValidationError is the validation error returned by
// MusicGenreData.Validate if the designated constraints aren't met.
type MusicGenreDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MusicGenreDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MusicGenreDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MusicGenreDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MusicGenreDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MusicGenreDataValidationError) ErrorName() string { return "MusicGenreDataValidationError" }

// Error satisfies the builtin error interface
func (e MusicGenreDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMusicGenreData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MusicGenreDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MusicGenreDataValidationError{}

// Validate checks the field values on FileData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FileData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FileData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FileDataMultiError, or nil
// if none found.
func (m *FileData) ValidateAll() error {
	return m.validate(true)
}

func (m *FileData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Cid

	// no validation rules for Name

	// no validation rules for Size

	if len(errors) > 0 {
		return FileDataMultiError(errors)
	}

	return nil
}

// FileDataMultiError is an error wrapping multiple validation errors returned
// by FileData.ValidateAll() if the designated constraints aren't met.
type FileDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FileDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FileDataMultiError) AllErrors() []error { return m }

// FileDataValidationError is the validation error returned by
// FileData.Validate if the designated constraints aren't met.
type FileDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileDataValidationError) ErrorName() string { return "FileDataValidationError" }

// Error satisfies the builtin error interface
func (e FileDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileDataValidationError{}

// Validate checks the field values on CoverData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CoverData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CoverData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CoverDataMultiError, or nil
// if none found.
func (m *CoverData) ValidateAll() error {
	return m.validate(true)
}

func (m *CoverData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CoverDataValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CoverDataValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CoverDataValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CoverDataValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CoverDataValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CoverDataValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CoverDataValidationError{
					field:  "File",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CoverDataValidationError{
					field:  "File",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CoverDataValidationError{
				field:  "File",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Width

	// no validation rules for Height

	if len(errors) > 0 {
		return CoverDataMultiError(errors)
	}

	return nil
}

// CoverDataMultiError is an error wrapping multiple validation errors returned
// by CoverData.ValidateAll() if the designated constraints aren't met.
type CoverDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CoverDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CoverDataMultiError) AllErrors() []error { return m }

// CoverDataValidationError is the validation error returned by
// CoverData.Validate if the designated constraints aren't met.
type CoverDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CoverDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CoverDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CoverDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CoverDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CoverDataValidationError) ErrorName() string { return "CoverDataValidationError" }

// Error satisfies the builtin error interface
func (e CoverDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCoverData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CoverDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CoverDataValidationError{}

// Validate checks the field values on MusicPlaylistData with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MusicPlaylistData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MusicPlaylistData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MusicPlaylistDataMultiError, or nil if none found.
func (m *MusicPlaylistData) ValidateAll() error {
	return m.validate(true)
}

func (m *MusicPlaylistData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MusicPlaylistDataValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MusicPlaylistDataValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MusicPlaylistDataValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MusicPlaylistDataValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MusicPlaylistDataValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MusicPlaylistDataValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	// no validation rules for Description

	if all {
		switch v := interface{}(m.GetCover()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MusicPlaylistDataValidationError{
					field:  "Cover",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MusicPlaylistDataValidationError{
					field:  "Cover",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCover()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MusicPlaylistDataValidationError{
				field:  "Cover",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Private

	if len(errors) > 0 {
		return MusicPlaylistDataMultiError(errors)
	}

	return nil
}

// MusicPlaylistDataMultiError is an error wrapping multiple validation errors
// returned by MusicPlaylistData.ValidateAll() if the designated constraints
// aren't met.
type MusicPlaylistDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MusicPlaylistDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MusicPlaylistDataMultiError) AllErrors() []error { return m }

// MusicPlaylistDataValidationError is the validation error returned by
// MusicPlaylistData.Validate if the designated constraints aren't met.
type MusicPlaylistDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MusicPlaylistDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MusicPlaylistDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MusicPlaylistDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MusicPlaylistDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MusicPlaylistDataValidationError) ErrorName() string {
	return "MusicPlaylistDataValidationError"
}

// Error satisfies the builtin error interface
func (e MusicPlaylistDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMusicPlaylistData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MusicPlaylistDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MusicPlaylistDataValidationError{}

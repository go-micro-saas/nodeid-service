// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/nodeid-service/v1/resources/node_event.resource.v1.proto

package resourcev1

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

// Validate checks the field values on SubscribeRenewalNodeIdEventReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SubscribeRenewalNodeIdEventReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubscribeRenewalNodeIdEventReq with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// SubscribeRenewalNodeIdEventReqMultiError, or nil if none found.
func (m *SubscribeRenewalNodeIdEventReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SubscribeRenewalNodeIdEventReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SubscribeRenewalNodeIdEventReqMultiError(errors)
	}

	return nil
}

// SubscribeRenewalNodeIdEventReqMultiError is an error wrapping multiple
// validation errors returned by SubscribeRenewalNodeIdEventReq.ValidateAll()
// if the designated constraints aren't met.
type SubscribeRenewalNodeIdEventReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubscribeRenewalNodeIdEventReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubscribeRenewalNodeIdEventReqMultiError) AllErrors() []error { return m }

// SubscribeRenewalNodeIdEventReqValidationError is the validation error
// returned by SubscribeRenewalNodeIdEventReq.Validate if the designated
// constraints aren't met.
type SubscribeRenewalNodeIdEventReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscribeRenewalNodeIdEventReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscribeRenewalNodeIdEventReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscribeRenewalNodeIdEventReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscribeRenewalNodeIdEventReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscribeRenewalNodeIdEventReqValidationError) ErrorName() string {
	return "SubscribeRenewalNodeIdEventReqValidationError"
}

// Error satisfies the builtin error interface
func (e SubscribeRenewalNodeIdEventReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscribeRenewalNodeIdEventReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscribeRenewalNodeIdEventReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscribeRenewalNodeIdEventReqValidationError{}

// Validate checks the field values on SubscribeRenewalNodeIdEventResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SubscribeRenewalNodeIdEventResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubscribeRenewalNodeIdEventResp with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// SubscribeRenewalNodeIdEventRespMultiError, or nil if none found.
func (m *SubscribeRenewalNodeIdEventResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SubscribeRenewalNodeIdEventResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Reason

	// no validation rules for Message

	// no validation rules for Metadata

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubscribeRenewalNodeIdEventRespValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubscribeRenewalNodeIdEventRespValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscribeRenewalNodeIdEventRespValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SubscribeRenewalNodeIdEventRespMultiError(errors)
	}

	return nil
}

// SubscribeRenewalNodeIdEventRespMultiError is an error wrapping multiple
// validation errors returned by SubscribeRenewalNodeIdEventResp.ValidateAll()
// if the designated constraints aren't met.
type SubscribeRenewalNodeIdEventRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubscribeRenewalNodeIdEventRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubscribeRenewalNodeIdEventRespMultiError) AllErrors() []error { return m }

// SubscribeRenewalNodeIdEventRespValidationError is the validation error
// returned by SubscribeRenewalNodeIdEventResp.Validate if the designated
// constraints aren't met.
type SubscribeRenewalNodeIdEventRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscribeRenewalNodeIdEventRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscribeRenewalNodeIdEventRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscribeRenewalNodeIdEventRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscribeRenewalNodeIdEventRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscribeRenewalNodeIdEventRespValidationError) ErrorName() string {
	return "SubscribeRenewalNodeIdEventRespValidationError"
}

// Error satisfies the builtin error interface
func (e SubscribeRenewalNodeIdEventRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscribeRenewalNodeIdEventResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscribeRenewalNodeIdEventRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscribeRenewalNodeIdEventRespValidationError{}

// Validate checks the field values on SubscribeRenewalNodeIdEventRespData with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *SubscribeRenewalNodeIdEventRespData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubscribeRenewalNodeIdEventRespData
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// SubscribeRenewalNodeIdEventRespDataMultiError, or nil if none found.
func (m *SubscribeRenewalNodeIdEventRespData) ValidateAll() error {
	return m.validate(true)
}

func (m *SubscribeRenewalNodeIdEventRespData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ConsumerCounter

	if len(errors) > 0 {
		return SubscribeRenewalNodeIdEventRespDataMultiError(errors)
	}

	return nil
}

// SubscribeRenewalNodeIdEventRespDataMultiError is an error wrapping multiple
// validation errors returned by
// SubscribeRenewalNodeIdEventRespData.ValidateAll() if the designated
// constraints aren't met.
type SubscribeRenewalNodeIdEventRespDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubscribeRenewalNodeIdEventRespDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubscribeRenewalNodeIdEventRespDataMultiError) AllErrors() []error { return m }

// SubscribeRenewalNodeIdEventRespDataValidationError is the validation error
// returned by SubscribeRenewalNodeIdEventRespData.Validate if the designated
// constraints aren't met.
type SubscribeRenewalNodeIdEventRespDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscribeRenewalNodeIdEventRespDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscribeRenewalNodeIdEventRespDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscribeRenewalNodeIdEventRespDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscribeRenewalNodeIdEventRespDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscribeRenewalNodeIdEventRespDataValidationError) ErrorName() string {
	return "SubscribeRenewalNodeIdEventRespDataValidationError"
}

// Error satisfies the builtin error interface
func (e SubscribeRenewalNodeIdEventRespDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscribeRenewalNodeIdEventRespData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscribeRenewalNodeIdEventRespDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscribeRenewalNodeIdEventRespDataValidationError{}

// Validate checks the field values on StopRenewalNodeIdEventReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StopRenewalNodeIdEventReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StopRenewalNodeIdEventReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StopRenewalNodeIdEventReqMultiError, or nil if none found.
func (m *StopRenewalNodeIdEventReq) ValidateAll() error {
	return m.validate(true)
}

func (m *StopRenewalNodeIdEventReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StopRenewalNodeIdEventReqMultiError(errors)
	}

	return nil
}

// StopRenewalNodeIdEventReqMultiError is an error wrapping multiple validation
// errors returned by StopRenewalNodeIdEventReq.ValidateAll() if the
// designated constraints aren't met.
type StopRenewalNodeIdEventReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StopRenewalNodeIdEventReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StopRenewalNodeIdEventReqMultiError) AllErrors() []error { return m }

// StopRenewalNodeIdEventReqValidationError is the validation error returned by
// StopRenewalNodeIdEventReq.Validate if the designated constraints aren't met.
type StopRenewalNodeIdEventReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StopRenewalNodeIdEventReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StopRenewalNodeIdEventReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StopRenewalNodeIdEventReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StopRenewalNodeIdEventReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StopRenewalNodeIdEventReqValidationError) ErrorName() string {
	return "StopRenewalNodeIdEventReqValidationError"
}

// Error satisfies the builtin error interface
func (e StopRenewalNodeIdEventReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStopRenewalNodeIdEventReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StopRenewalNodeIdEventReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StopRenewalNodeIdEventReqValidationError{}

// Validate checks the field values on StopRenewalNodeIdEventResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StopRenewalNodeIdEventResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StopRenewalNodeIdEventResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StopRenewalNodeIdEventRespMultiError, or nil if none found.
func (m *StopRenewalNodeIdEventResp) ValidateAll() error {
	return m.validate(true)
}

func (m *StopRenewalNodeIdEventResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Reason

	// no validation rules for Message

	// no validation rules for Metadata

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StopRenewalNodeIdEventRespValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StopRenewalNodeIdEventRespValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StopRenewalNodeIdEventRespValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StopRenewalNodeIdEventRespMultiError(errors)
	}

	return nil
}

// StopRenewalNodeIdEventRespMultiError is an error wrapping multiple
// validation errors returned by StopRenewalNodeIdEventResp.ValidateAll() if
// the designated constraints aren't met.
type StopRenewalNodeIdEventRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StopRenewalNodeIdEventRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StopRenewalNodeIdEventRespMultiError) AllErrors() []error { return m }

// StopRenewalNodeIdEventRespValidationError is the validation error returned
// by StopRenewalNodeIdEventResp.Validate if the designated constraints aren't met.
type StopRenewalNodeIdEventRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StopRenewalNodeIdEventRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StopRenewalNodeIdEventRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StopRenewalNodeIdEventRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StopRenewalNodeIdEventRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StopRenewalNodeIdEventRespValidationError) ErrorName() string {
	return "StopRenewalNodeIdEventRespValidationError"
}

// Error satisfies the builtin error interface
func (e StopRenewalNodeIdEventRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStopRenewalNodeIdEventResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StopRenewalNodeIdEventRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StopRenewalNodeIdEventRespValidationError{}

// Validate checks the field values on StopRenewalNodeIdEventRespData with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StopRenewalNodeIdEventRespData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StopRenewalNodeIdEventRespData with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// StopRenewalNodeIdEventRespDataMultiError, or nil if none found.
func (m *StopRenewalNodeIdEventRespData) ValidateAll() error {
	return m.validate(true)
}

func (m *StopRenewalNodeIdEventRespData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ConsumerCounter

	if len(errors) > 0 {
		return StopRenewalNodeIdEventRespDataMultiError(errors)
	}

	return nil
}

// StopRenewalNodeIdEventRespDataMultiError is an error wrapping multiple
// validation errors returned by StopRenewalNodeIdEventRespData.ValidateAll()
// if the designated constraints aren't met.
type StopRenewalNodeIdEventRespDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StopRenewalNodeIdEventRespDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StopRenewalNodeIdEventRespDataMultiError) AllErrors() []error { return m }

// StopRenewalNodeIdEventRespDataValidationError is the validation error
// returned by StopRenewalNodeIdEventRespData.Validate if the designated
// constraints aren't met.
type StopRenewalNodeIdEventRespDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StopRenewalNodeIdEventRespDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StopRenewalNodeIdEventRespDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StopRenewalNodeIdEventRespDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StopRenewalNodeIdEventRespDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StopRenewalNodeIdEventRespDataValidationError) ErrorName() string {
	return "StopRenewalNodeIdEventRespDataValidationError"
}

// Error satisfies the builtin error interface
func (e StopRenewalNodeIdEventRespDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStopRenewalNodeIdEventRespData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StopRenewalNodeIdEventRespDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StopRenewalNodeIdEventRespDataValidationError{}

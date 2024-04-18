// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: apu/mail_sender/mail_sender.proto

package mail_sender

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

// Validate checks the field values on RequestSendStatisticServerToEmail with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *RequestSendStatisticServerToEmail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RequestSendStatisticServerToEmail
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// RequestSendStatisticServerToEmailMultiError, or nil if none found.
func (m *RequestSendStatisticServerToEmail) ValidateAll() error {
	return m.validate(true)
}

func (m *RequestSendStatisticServerToEmail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RequestSendStatisticServerToEmailMultiError(errors)
	}

	return nil
}

// RequestSendStatisticServerToEmailMultiError is an error wrapping multiple
// validation errors returned by
// RequestSendStatisticServerToEmail.ValidateAll() if the designated
// constraints aren't met.
type RequestSendStatisticServerToEmailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestSendStatisticServerToEmailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestSendStatisticServerToEmailMultiError) AllErrors() []error { return m }

// RequestSendStatisticServerToEmailValidationError is the validation error
// returned by RequestSendStatisticServerToEmail.Validate if the designated
// constraints aren't met.
type RequestSendStatisticServerToEmailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestSendStatisticServerToEmailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestSendStatisticServerToEmailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestSendStatisticServerToEmailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestSendStatisticServerToEmailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestSendStatisticServerToEmailValidationError) ErrorName() string {
	return "RequestSendStatisticServerToEmailValidationError"
}

// Error satisfies the builtin error interface
func (e RequestSendStatisticServerToEmailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestSendStatisticServerToEmail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestSendStatisticServerToEmailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestSendStatisticServerToEmailValidationError{}

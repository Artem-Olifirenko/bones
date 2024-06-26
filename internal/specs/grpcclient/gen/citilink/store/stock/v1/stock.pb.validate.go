// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: citilink/store/stock/v1/stock.proto

package stockv1

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

// Validate checks the field values on StockInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StockInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StockInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StockInfoMultiError, or nil
// if none found.
func (m *StockInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *StockInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStock()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StockInfoValidationError{
					field:  "Stock",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StockInfoValidationError{
					field:  "Stock",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStock()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StockInfoValidationError{
				field:  "Stock",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	// no validation rules for PupId

	if len(errors) > 0 {
		return StockInfoMultiError(errors)
	}

	return nil
}

// StockInfoMultiError is an error wrapping multiple validation errors returned
// by StockInfo.ValidateAll() if the designated constraints aren't met.
type StockInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StockInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StockInfoMultiError) AllErrors() []error { return m }

// StockInfoValidationError is the validation error returned by
// StockInfo.Validate if the designated constraints aren't met.
type StockInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StockInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StockInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StockInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StockInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StockInfoValidationError) ErrorName() string { return "StockInfoValidationError" }

// Error satisfies the builtin error interface
func (e StockInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStockInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StockInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StockInfoValidationError{}

// Validate checks the field values on StockInfo_Stock with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StockInfo_Stock) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StockInfo_Stock with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StockInfo_StockMultiError, or nil if none found.
func (m *StockInfo_Stock) ValidateAll() error {
	return m.validate(true)
}

func (m *StockInfo_Stock) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Type

	// no validation rules for Name

	if len(errors) > 0 {
		return StockInfo_StockMultiError(errors)
	}

	return nil
}

// StockInfo_StockMultiError is an error wrapping multiple validation errors
// returned by StockInfo_Stock.ValidateAll() if the designated constraints
// aren't met.
type StockInfo_StockMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StockInfo_StockMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StockInfo_StockMultiError) AllErrors() []error { return m }

// StockInfo_StockValidationError is the validation error returned by
// StockInfo_Stock.Validate if the designated constraints aren't met.
type StockInfo_StockValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StockInfo_StockValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StockInfo_StockValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StockInfo_StockValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StockInfo_StockValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StockInfo_StockValidationError) ErrorName() string { return "StockInfo_StockValidationError" }

// Error satisfies the builtin error interface
func (e StockInfo_StockValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStockInfo_Stock.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StockInfo_StockValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StockInfo_StockValidationError{}

package okay

import (
	"fmt"
	"regexp"
)

const (
	RuleNotEmpty      = "not_empty"
	RuleLength        = "length"
	RuleLengthBetween = "length_between"
	RuleMin           = "min"
	RuleMax           = "max"
	RuleMatch         = "match"
	RuleAlphanumeric  = "alphanumeric"
	RuleUnique        = "unique"
)

var (
	MessageNotEmpty     = "cannot be empty"
	MessageLength       = "length must be %d"
	MessageLengthIn     = "length must be between %d and %d"
	MessageMin          = "must be %d or more"
	MessageMax          = "must be %d or less"
	MessageMatch        = "must match %s"
	MessageAlphanumeric = "must only contain letters a to z and digits"
	MessageUnique       = "must be unique"

	ReAlphanumeric = regexp.MustCompile("^[a-zA-Z0-9]+$")
)

// NotEmpty checks if the given string, slice or map is not empty.
//
//	err := okay.NotEmpty("keywords", []string{"childcare"})
func NotEmpty[T ~string | ~[]any | ~map[any]any](key string, val T) *Error {
	if len(val) == 0 {
		return NewError(key, RuleNotEmpty).WithMessage(MessageNotEmpty)
	}
	return nil
}

// Length checks if the given string, slice or map has a given length.
//
//	err := okay.Length("keywords", []string{"childcare"}, 1)
func Length[T ~string | ~[]any | ~map[any]any](key string, val T, n int) *Error {
	if len(val) != n {
		return NewError(key, RuleLength, n).WithMessage(fmt.Sprintf(MessageLength, n))
	}
	return nil
}

// LengthBetween checks if the given string, slice or map has a length that is
// greater than or equal to min and less than or equal to max.
//
//	err := okay.Length("keywords", []string{"childcare"}, 0, 10)
func LengthBetween[T ~string | ~[]any | ~map[any]any](key string, val T, min, max int) *Error {
	if len(val) < min || len(val) > max {
		return NewError(key, RuleLengthBetween, min, max).WithMessage(fmt.Sprintf(MessageLengthIn, min, max))
	}
	return nil
}

// Min checks if the given string, slice or map has a length that is
// greater than or equal to min.
//
//	err := okay.Min("keywords", []string{"childcare", "education"}, 1)
func Min[T int | int64 | float64](key string, val T, min T) *Error {
	if val < min {
		return NewError(key, RuleMin, min).WithMessage(fmt.Sprintf(MessageMin, min))
	}
	return nil
}

// Max checks if the given string, slice or map has a length that is
// less than or equal to max.
//
//	err := okay.Max("keywords", []string{"childcare"}, 10)
func Max[T int | int64 | float64](key string, val T, max T) *Error {
	if val > max {
		return NewError(key, RuleMax, max).WithMessage(fmt.Sprintf(MessageMax, max))
	}
	return nil
}

// Match checks if the given string matches a regular expression.
//
//	err := okay.Match("issn", "1940-5758", regexp.MustCompile(`^[0-9]{4}-[0-9]{3}[0-9X]$`))
func Match(key, val string, r *regexp.Regexp) *Error {
	if !r.MatchString(val) {
		return NewError(key, RuleMatch, r).WithMessage(fmt.Sprintf(MessageMatch, r))
	}
	return nil
}

// Alphanumeric checks if a given string only contains letters a to z, letters A to Z or digits.
//
//	err := okay.Match("issn", "1940-5758", regexp.MustCompile(`^[0-9]{4}-[0-9]{3}[0-9X]$`))
func Alphanumeric(key, val string) *Error {
	if !ReAlphanumeric.MatchString(val) {
		return NewError(key, RuleAlphanumeric).WithMessage(MessageAlphanumeric)
	}
	return nil
}

// ErrNotUnique is a convenience function to signal that a given key fails a uniqueness test.
func ErrNotUnique(key string) *Error {
	return NewError(key, RuleUnique).WithMessage(MessageUnique)
}

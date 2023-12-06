package okay

import (
	"fmt"
)

// Add is a convenience function to create an Errors and add zero or more validation errors.
// Returns nil if there are no errors.
func Validate(errs ...error) error {
	e := NewErrors()
	Add(e, errs...)
	return e.ErrorOrNil()
}

// Add is a convenience function to add zero or more validation errors to an
// Errors.
//
// If err is nil, a new Errors wil be created and returned. If the given validation
// error is an Errors, it's validation errors will be added to err.
//
// If err is not an Errors or the given validation errors are not an Errors or
// an Error, the function will panic.
func Add(err error, errs ...error) error {
	if err == nil {
		err = NewErrors()
	}
	e := err.(*Errors)
	for _, err := range errs {
		if err == nil {
			continue
		} else if ee, ok := err.(*Errors); ok {
			e.Add(ee.Errors...)
		} else {
			e.Add(err.(*Error))
		}
	}
	return e.ErrorOrNil()
}

// func AddWithPrefix(err error, prefix string, errs ...error) error {
// 	if err == nil {
// 		err = NewErrors()
// 	}
// 	e := err.(*Errors)
// 	for _, err := range errs {
// 		if err == nil {
// 			continue
// 		} else if ee, ok := err.(*Errors); ok {
// 			e.AddWithPrefix(prefix, ee.Errors...)
// 		} else {
// 			e.AddWithPrefix(prefix, err.(*Error))
// 		}
// 	}
// 	return e.ErrorOrNil()
// }

type Errors struct {
	Errors []*Error
}

// NewErrors constructs a new Errors with the given validation errors.
func NewErrors(errs ...*Error) *Errors {
	return new(Errors).Add(errs...)
}

// Error returns a string representation of an Errors.
func (e *Errors) Error() string {
	msg := ""
	for i, err := range e.Errors {
		msg += err.Error()
		if i < len(e.Errors)-1 {
			msg += ", "
		}
	}
	return msg
}

// Add zero or more validation errors.
func (e *Errors) Add(errs ...*Error) *Errors {
	for _, err := range errs {
		if err != nil {
			e.Errors = append(e.Errors, err)
		}
	}
	return e
}

// func (e *Errors) AddWithPrefix(prefix string, errs ...*Error) *Errors {
// 	for _, err := range errs {
// 		if err != nil {
// 			e.Errors = append(e.Errors, &Error{
// 				Key:     prefix + err.Key,
// 				Rule:    err.Rule,
// 				Params:  err.Params,
// 				Message: err.Message,
// 			})
// 		}
// 	}
// 	return e
// }

// Get fetches an Error by key or return nil if the key is not found.
func (e *Errors) Get(key string) *Error {
	for _, e := range e.Errors {
		if e.Key == key {
			return e
		}
	}
	return nil
}

// ErrorOrNil returns Errors as a (nil) error interface value.
func (e *Errors) ErrorOrNil() error {
	if len(e.Errors) > 0 {
		return e
	}
	return nil
}

type Error struct {
	Key     string
	Rule    string
	Params  []any
	Message string
}

// NewError constructs a new validation error. key represents the field or value
// that failed validation. There are no assumptions about the nature of this
// key, it could be a JSON pointer or the name of a (nested) form field.
func NewError(key, rule string, params ...any) *Error {
	return &Error{
		Key:    key,
		Rule:   rule,
		Params: params,
	}
}

// WithMessage sets a custom error message if the validation error is not nil.
func (e *Error) WithMessage(msg string) *Error {
	if e != nil {
		e.Message = msg
	}
	return e
}

// Error returns a string representation of the validation error.
func (e *Error) Error() string {
	msg := e.Key
	if msg != "" {
		msg += " "
	}
	if e.Message != "" {
		msg += e.Message
	} else if e.Rule != "" {
		msg += e.Rule
		if len(e.Params) > 0 {
			msg += "["
			for i, p := range e.Params {
				msg += fmt.Sprintf("%v", p)
				if i < len(e.Params)-1 {
					msg += ", "
				}
			}
			msg += "]"
		}
	}
	return msg
}

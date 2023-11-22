package okay

import (
	"fmt"
)

func Validate(errs ...error) error {
	e := NewErrors()
	Add(e, errs...)
	return e.ErrorOrNil()
}

func Add(err error, errs ...error) error {
	if err == nil {
		err = NewErrors()
	}
	e := err.(*Errors)
	for _, err := range errs {
		if ee, ok := err.(*Errors); ok {
			e.Add(ee.Errors...)
		} else {
			e.Add(err.(*Error))
		}
	}
	return e.ErrorOrNil()
}

func AddWithPrefix(err error, prefix string, errs ...error) error {
	if err == nil {
		err = NewErrors()
	}
	e := err.(*Errors)
	for _, err := range errs {
		if ee, ok := err.(*Errors); ok {
			e.AddWithPrefix(prefix, ee.Errors...)
		} else {
			e.AddWithPrefix(prefix, err.(*Error))
		}
	}
	return e.ErrorOrNil()
}

type Errors struct {
	Errors []*Error
}

func NewErrors(errs ...*Error) *Errors {
	return new(Errors).Add(errs...)
}

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

func (e *Errors) Add(errs ...*Error) *Errors {
	for _, err := range errs {
		if err != nil {
			e.Errors = append(e.Errors, err)
		}
	}
	return e
}

func (e *Errors) AddWithPrefix(prefix string, errs ...*Error) *Errors {
	for _, err := range errs {
		if err != nil {
			e.Errors = append(e.Errors, &Error{
				Key:     prefix + err.Key,
				Rule:    err.Rule,
				Params:  err.Params,
				Message: err.Message,
			})
		}
	}
	return e
}

func (e *Errors) Get(key string) *Error {
	for _, e := range e.Errors {
		if e.Key == key {
			return e
		}
	}
	return nil
}

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

func NewError(key, rule string, params ...any) *Error {
	return &Error{
		Key:    key,
		Rule:   rule,
		Params: params,
	}
}

func (e *Error) WithMessage(msg string) *Error {
	if e != nil {
		e.Message = msg
	}
	return e
}

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

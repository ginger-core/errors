package errors

import (
	"context"
	"fmt"
)

type Error interface {
	error

	WithDesc(description string) Error
	WithDescf(format string, args ...any) Error
	GetDesc() string

	WithTrace(id string) Error
	GetTrace() string

	WithError(err error) Error
	GetError() error

	WithId(id string) Error
	GetId() string
	// WithMessage sets default error message
	WithMessage(message string) Error
	GetMessage() string

	WithContext(ctx context.Context) Error
	GetContext() context.Context

	WithSource(source string) (err Error)
	GetSource() string

	WithType(t Type) Error
	GetType() Type
	// WithMessageOne sets message for plural count 1
	WithMessageOne(message string) Error
	GetMessageOne() string
	// WithValue sets dynamic value of message
	WithValue(key string, value interface{}) Error
	WithValues(properties map[string]interface{}) Error
	GetValues() map[string]interface{}

	WithPluralCount(count int) Error
	GetPluralCount() int

	WithCode(code int) Error
	GetCode() int

	WithExtra(key string, err ExtraError) Error
	GetExtra() map[string][]ExtraError

	WithDetail(detail Detail) Error
	GetDetail() Detail

	WithProperty(key string, value interface{}) Error
	GetProperties() map[string]interface{}

	IsType(t Type) bool

	Clone() Error
}

type err struct {
	private `json:"-"`

	Code int `json:"code,omitempty"`
	// Message default message of error
	Message string                  `json:"message,omitempty"`
	Errors  map[string][]ExtraError `json:"errors,omitempty"`
}

func New(e ...error) Error {
	var r Error = new(err)
	if len(e) > 0 && e[0] != nil {
		r = r.WithError(e[0])
	}
	r.WithType(TypeInternal)
	return r
}

func (e *err) WithDesc(description string) Error {
	e.Description = description
	return e
}

func (e *err) WithDescf(format string, args ...any) Error {
	e.Description = fmt.Sprintf(format, args...)
	return e
}

func (e *err) GetDesc() string {
	return e.Description
}

func (e *err) WithTrace(t string) Error {
	e.Trace = t + "/" + e.Trace
	return e
}

func (e *err) GetTrace() string {
	return e.Trace
}

func (e *err) WithContext(ctx context.Context) Error {
	e.ctx = ctx
	return e
}

func (e *err) GetContext() context.Context {
	return e.ctx
}

func (e *err) Error() string {
	var r string
	if e.Message != "" {
		r += "message: " + e.Message + ", "
	}
	r += "private: (" + e.private.String() + ")"
	return r
}

func (e *err) WithError(err error) Error {
	e.error = err
	return e
}

func (e *err) GetError() error {
	return e.error
}

func (e *err) WithId(id string) Error {
	e.Id = id
	e.isIdCustomized = true
	return e
}

func (e *err) GetId() string {
	return e.Id
}

func (e *err) WithSource(source string) (err Error) {
	e.source = source
	return e
}

func (e *err) GetSource() string {
	return e.source
}

func (e *err) WithType(t Type) Error {
	e.Type = t
	e.ensureDefaults()
	return e
}

func (e *err) GetType() Type {
	return e.Type
}

func (e *err) WithMessage(message string) Error {
	e.Message = message
	e.isMessageCustomized = true
	return e
}

func (e *err) GetMessage() string {
	return e.Message
}

func (e *err) WithMessageOne(message string) Error {
	e.messageOne = message
	return e
}

func (e *err) GetMessageOne() string {
	return e.messageOne
}

func (e *err) WithValue(key string, value interface{}) Error {
	if e.values == nil {
		e.values = make(map[string]interface{})
	}
	e.values[key] = value
	return e
}

func (e *err) WithValues(properties map[string]interface{}) Error {
	e.values = properties
	return e
}

func (e *err) GetValues() map[string]interface{} {
	return e.values
}

func (e *err) WithPluralCount(count int) Error {
	e.pluralCount = count
	return e
}

func (e *err) GetPluralCount() int {
	return e.pluralCount
}

func (e *err) WithCode(code int) Error {
	e.Code = code
	return e
}

func (e *err) GetCode() int {
	return e.Code
}

func (e *err) WithExtra(key string, v ExtraError) Error {
	if e.Errors == nil {
		e.Errors = make(map[string][]ExtraError)
	}
	errs := e.Errors[key]
	if errs == nil {
		errs = make([]ExtraError, 0)
	}
	errs = append(errs, v)
	e.Errors[key] = errs
	return e
}

func (e *err) GetExtra() map[string][]ExtraError {
	return e.Errors
}

func (e *err) WithDetail(detail Detail) Error {
	e.detail = detail
	return e
}

func (e *err) GetDetail() Detail {
	return e.detail
}

func (e *err) WithProperty(key string, value interface{}) Error {
	if e.properties == nil {
		e.properties = make(map[string]interface{})
	}
	e.properties[key] = value
	return e
}

func (e *err) GetProperties() map[string]interface{} {
	return e.properties
}

func (e *err) IsType(t Type) bool {
	return e.Type == t
}

func (e *err) Clone() Error {
	return &err{
		private: e.private.clone(),
		Code:    e.Code,
		Message: e.Message,
		Errors:  e.Errors,
	}
}

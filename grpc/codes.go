package grpc

import (
	"github.com/ginger-core/errors"
	"google.golang.org/grpc/codes"
)

const (
	CodeInternal         codes.Code = codes.Internal
	CodeValidation       codes.Code = codes.InvalidArgument
	CodeNotFound         codes.Code = codes.NotFound
	CodeUnauthorized     codes.Code = codes.Unauthenticated
	CodeForbidden        codes.Code = codes.Unavailable
	CodeTooManyRequests  codes.Code = codes.ResourceExhausted
	CodeFailedDependency codes.Code = codes.FailedPrecondition
	CodeTooEarly         codes.Code = 425
	CodeDuplicate        codes.Code = codes.AlreadyExists
)

var codeMaps = map[errors.Type]codes.Code{
	errors.TypeInternal:         CodeInternal,
	errors.TypeValidation:       CodeValidation,
	errors.TypeNotFound:         CodeNotFound,
	errors.TypeUnauthorized:     CodeUnauthorized,
	errors.TypeForbidden:        CodeForbidden,
	errors.TypeTooManyRequests:  CodeTooManyRequests,
	errors.TypeFailedDependency: CodeFailedDependency,
	errors.TypeTooEarly:         CodeTooEarly,
	errors.TypeDuplicate:        CodeDuplicate,
}

var codeMapsR = map[codes.Code]errors.Type{
	CodeInternal:         errors.TypeInternal,
	CodeValidation:       errors.TypeValidation,
	CodeNotFound:         errors.TypeNotFound,
	CodeUnauthorized:     errors.TypeUnauthorized,
	CodeForbidden:        errors.TypeForbidden,
	CodeTooManyRequests:  errors.TypeTooManyRequests,
	CodeFailedDependency: errors.TypeFailedDependency,
	CodeTooEarly:         errors.TypeTooEarly,
	CodeDuplicate:        errors.TypeDuplicate,
}

func GetCodeFromType(t errors.Type) codes.Code {
	c, ok := codeMaps[t]
	if ok {
		return c
	}
	return codes.Unknown
}

func GetTypeFromCode(c codes.Code) errors.Type {
	t, ok := codeMapsR[c]
	if ok {
		return t
	}
	return errors.TypeInternal
}

package errors

func Internal(err ...error) Error {
	var e = New(err...)
	DefaultInternalError.Populate(e)
	return e
}

func Validation(err ...error) Error {
	var e = New(err...)
	DefaultValidationError.Populate(e)
	return e
}

func NotFound(err ...error) Error {
	var e = New(err...)
	DefaultNotFoundError.Populate(e)
	return e
}

func Unauthorized(err ...error) Error {
	var e = New(err...)
	DefaultUnauthorizedError.Populate(e)
	return e
}

func Forbidden(err ...error) Error {
	var e = New(err...)
	DefaultForbiddenError.Populate(e)
	return e
}

func TooManyRequests(err ...error) Error {
	var e = New(err...)
	DefaultTooManyRequestsError.Populate(e)
	return e
}

func FailedDependency(err ...error) Error {
	var e = New(err...)
	DefaultFailedDependencyError.Populate(e)
	return e
}

func TooEarly(err ...error) Error {
	var e = New(err...)
	DefaultTooEarlyError.Populate(e)
	return e
}

func Expired(err ...error) Error {
	var e = New(err...)
	DefaultExpiredError.Populate(e)
	return e
}

func Duplicate(err ...error) Error {
	var e = New(err...)
	DefaultDuplicateError.Populate(e)
	return e
}

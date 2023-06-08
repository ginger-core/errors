package errors

func Internal(err ...error) Error {
	var e = New(err...).WithType(TypeInternal)
	return e
}

func Validation(err ...error) Error {
	var e = New(err...).WithType(TypeValidation)
	return e
}

func NotFound(err ...error) Error {
	var e = New(err...).WithType(TypeNotFound)
	return e
}

func Unauthorized(err ...error) Error {
	var e = New(err...).WithType(TypeUnauthorized)
	return e
}

func Forbidden(err ...error) Error {
	var e = New(err...).WithType(TypeForbidden)
	return e
}

func TooManyRequests(err ...error) Error {
	var e = New(err...).WithType(TypeTooManyRequests)
	return e
}

func FailedDependency(err ...error) Error {
	var e = New(err...).WithType(TypeFailedDependency)
	return e
}

func TooEarly(err ...error) Error {
	var e = New(err...).WithType(TypeTooEarly)
	return e
}

func Expired(err ...error) Error {
	var e = New(err...).WithType(TypeExpired)
	return e
}

func Duplicate(err ...error) Error {
	var e = New(err...).WithType(TypeDuplicate)
	return e
}

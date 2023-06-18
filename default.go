package errors

var DefaultUnauthorizedError = New().
	WithType(TypeUnauthorized).
	WithId("UnauthorizedError").
	WithMessage("You are not authorized.")

var DefaultForbiddenError = New().
	WithType(TypeForbidden).
	WithId("ForbiddenError").
	WithMessage("Access to this section is denied.")

var DefaultValidationError = New().
	WithType(TypeValidation).
	WithId("ValidationError").
	WithMessage("Invalid request information.")

var DefaultNotFoundError = New().
	WithType(TypeNotFound).
	WithId("NotFoundError").
	WithMessage("Requested information not found.")

var DefaultInternalError = New().
	WithType(TypeInternal).
	WithId("InternalError").
	WithMessage("Internal error occurred.")

var DefaultDuplicateError = New().
	WithType(TypeDuplicate).
	WithId("DuplicateEntryError").
	WithMessage("Requested information already exists.")

var DefaultTooManyRequestsError = New().
	WithType(TypeTooManyRequests).
	WithId("TooManyRequestsError").
	WithMessage("You've made too many requests recently. " +
		"Please wait and try your request again later.")

var DefaultExpiredError = New().
	WithType(TypeExpired).
	WithId("ExpiredError").
	WithMessage("Your request has been expired, please start over and try again.")

var DefaultFailedDependencyError = New().
	WithType(TypeFailedDependency).
	WithId("FailedDependencyError").
	WithMessage("The request could not be performed on the resource because " +
		"the requested action depended on another action and that action failed.")

var DefaultTooEarlyError = New().
	WithType(TypeTooEarly).
	WithId("TooEarlyError").
	WithMessage("You're too early!")

var (
	defaultTypeMap = map[Type]Error{
		TypeUnauthorized:     DefaultUnauthorizedError,
		TypeForbidden:        DefaultForbiddenError,
		TypeValidation:       DefaultValidationError,
		TypeNotFound:         DefaultNotFoundError,
		TypeInternal:         DefaultInternalError,
		TypeDuplicate:        DefaultDuplicateError,
		TypeTooManyRequests:  DefaultTooManyRequestsError,
		TypeExpired:          DefaultExpiredError,
		TypeFailedDependency: DefaultFailedDependencyError,
		TypeTooEarly:         DefaultTooEarlyError,
	}
)

func (e *err) ensureDefaults() {
	defaultErr := defaultTypeMap[e.Type]
	if defaultErr == nil {
		return
	}
	defaultErr.Populate(e)
}

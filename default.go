package errors

var DefaultUnauthorizedError = Unauthorized().
	WithId("UnauthorizedError").
	WithMessage("You are not authorized.")

var DefaultForbiddenError = Forbidden().
	WithId("ForbiddenError").
	WithMessage("Access to this section is denied.")

var DefaultValidationError = Validation().
	WithId("ValidationError").
	WithMessage("Invalid request information.")

var DefaultNotFoundError = NotFound().
	WithId("NotFoundError").
	WithMessage("Requested information not found.")

var DefaultInternalError = Internal().
	WithId("InternalError").
	WithMessage("Internal error occurred.")

var DefaultDuplicateError = Duplicate().
	WithId("DuplicateEntryError").
	WithMessage("Requested information already exists.")

var DefaultTooManyRequestsError = TooManyRequests().
	WithId("TooManyRequestsError").
	WithMessage("You've made too many requests recently. " +
		"Please wait and try your request again later.")

var DefaultExpiredError = TooManyRequests().
	WithId("ExpiredError").
	WithMessage("Your request has been expired, please start over and try again.")

var DefaultFailedDependencyError = FailedDependency().
	WithId("FailedDependencyError").
	WithMessage("The request could not be performed on the resource because " +
		"the requested action depended on another action and that action failed.")

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
	}
)

func (e *err) ensureDefaults() {
	defaultErr := defaultTypeMap[e.Type]
	if defaultErr == nil {
		return
	}
	if e.Id == "" {
		e.Id = defaultErr.GetId()
	}
	if e.Message == "" {
		e.Message = defaultErr.GetMessage()
	}
}

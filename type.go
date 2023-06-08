package errors

type Type string

const (
	TypeInternal         Type = "INTERNAL"
	TypeValidation       Type = "VALIDATION"
	TypeNotFound         Type = "NOT_FOUND"
	TypeUnauthorized     Type = "UNAUTHORIZED"
	TypeForbidden        Type = "FORBIDDEN"
	TypeTooManyRequests  Type = "TOO_MANY_REQUESTS"
	TypeFailedDependency Type = "FAILED_DEPENDENCY"
	TypeTooEarly         Type = "TOO_EARLY"
	TypeExpired          Type = "EXPIRED"
	TypeDuplicate        Type = "DUPLICATE"
)

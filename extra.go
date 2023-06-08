package errors

type ExtraError interface {
	error

	WithError(err error) ExtraError
	GetError() error

	WithId(id string) ExtraError
	GetId() string
	// WithMessage sets error message
	WithMessage(message string) ExtraError
	GetMessage() string
}

type extra struct {
	error   `json:"-"`
	Id      string `json:"-"`
	Message string `json:"message"`
}

func NewExtra(e ...error) ExtraError {
	var r ExtraError = new(extra)
	if len(e) > 0 && e[0] != nil {
		r = r.WithError(e[0])
	}
	return r
}

func (f *extra) WithError(err error) ExtraError {
	f.error = err
	if e := err.(Error); e != nil {
		f.Id = e.GetId()
		f.Message = e.GetMessage()
	}
	return f
}

func (f *extra) GetError() error {
	return f.error
}

func (f *extra) Error() string {
	var r string
	if f.Id != "" {
		r += "id: " + f.Id + ", "
	}
	if f.Message != "" {
		r += "message: " + f.Message + ", "
	}
	if f.error != nil {
		r += "err: " + f.error.Error() + ", "
	}
	return r
}

func (f *extra) WithId(id string) ExtraError {
	f.Id = id
	return f
}

func (f *extra) GetId() string {
	return f.Id
}

func (f *extra) WithMessage(message string) ExtraError {
	f.Message = message
	return f
}

func (f *extra) GetMessage() string {
	return f.Message
}

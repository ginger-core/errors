package errors

import "fmt"

type Detail map[string]interface{}

func NewDetail() Detail {
	return Detail{}
}

func (d Detail) Set(key string, value interface{}) {
	d[key] = value
}

func (d Detail) With(key string, value interface{}) Detail {
	d[key] = value
	return d
}

func (d Detail) Get(key string) interface{} {
	return d[key]
}

func (d Detail) WithId(id string) Detail {
	d["id"] = id
	return d
}

func (d Detail) WithMessage(format string, args ...any) Detail {
	d["message"] = fmt.Sprintf(format, args...)
	return d
}

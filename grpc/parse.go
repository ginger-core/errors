package grpc

import (
	"strconv"

	"github.com/ginger-core/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func Parse(err error) errors.Error {
	rErr := errors.New(err)

	s := status.Convert(err)
	// fill err from status
	rErr.
		WithType(GetTypeFromCode(s.Code())).
		WithMessage(s.Message())

	st := s.Details()
	if len(st) > 0 {
		d, ok := st[0].(*errdetails.ErrorInfo)
		if ok {
			code, _ := strconv.Atoi(d.Metadata["code"])
			rErr.
				WithTrace(d.GetDomain()).
				WithId(d.GetReason()).
				WithDesc(d.Metadata["desc"]).
				WithCode(code)
		}
	}
	return rErr
}

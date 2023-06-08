package grpc

import (
	"fmt"

	"github.com/ginger-core/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func Generate(err errors.Error) error {
	detail := &errdetails.ErrorInfo{
		Domain: err.GetTrace(),
		Reason: err.GetId(),
		Metadata: map[string]string{
			"desc": err.GetDesc(),
			"code": fmt.Sprint(err.GetCode()),
		},
	}
	s, _ := status.
		New(GetCodeFromType(err.GetType()), err.GetMessage()).
		WithDetails(detail)
	return s.Err()
}

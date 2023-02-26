package errorx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type errres struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func GrpcExpectErrorHandler(err error) (int, interface{}) {
	status, ok := status.FromError(err)
	if ok {
		if status.Code() > codes.Unauthenticated {
			return 200, errres{
				Code: int32(status.Code()),
				Msg:  status.Message(),
			}
		}
	}

	return 200, errres{
		Code: -1,
		Msg:  err.Error(),
	}
}

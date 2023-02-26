package errorx

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryExpectErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	res, err := handler(ctx, req)
	if err != nil {
		return nil, grpcExpectError(err)
	}
	return res, nil
}

func StreamExpectErrorInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, ss)
	if err != nil {
		return grpcExpectError(err)
	}
	return err
}

func grpcExpectError(err error) error {
	var expect *ExpectError
	if errors.As(err, &expect) {
		err := expect.DeepestExpectError()
		msg := err.Msg
		if err.err != nil {
			msg = msg + ":" + err.err.Error()
		}
		return status.Error(codes.Code(err.Code), msg)
	}
	return err
}

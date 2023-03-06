package authentication

import (
	"context"
	"infra/pkg/constant"
	"infra/pkg/errorx"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var allowMethodList = map[string]struct{}{
	"/system.v1.systemService/GetAuthToken": {},
}

func UnaryAuthenticationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	_, ok := allowMethodList[info.FullMethod]
	if ok {
		return handler(ctx, req)
	}
	md, _ := metadata.FromIncomingContext(ctx)
	vals1 := md.Get(constant.AuthTokenKey)
	if len(vals1) == 0 {
		return nil, errorx.ErrAuthenticationFalied
	}

	authToken := vals1[0]
	userSession, err := NewUserSessionCacheRegistry(authToken).Get(ctx)
	if err != nil {
		return nil, errorx.ErrAuthenticationFalied.WrapErr(err)
	}

	ver, err := NewAuthVerCacheRegistry(userSession.Gid).Get(ctx)
	if err != nil {
		return nil, errorx.ErrAuthenticationFalied.WrapErr(err)
	}

	if ver != userSession.AuthVer {
		return nil, errorx.ErrAuthenticationInfoExpired
	}

	return handler(ctx, req)
}

// 当服务端的第一次SendMsg触发时，会拦截下边的逻辑，若不通过则客户端的RecvMsg时会报错。
func StreamAuthenticationInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_, ok := allowMethodList[info.FullMethod]
	if ok {
		return handler(srv, ss)
	}
	md, _ := metadata.FromIncomingContext(ss.Context())
	vals1 := md.Get(constant.AuthTokenKey)
	if len(vals1) == 0 {
		return errorx.ErrAuthenticationFalied
	}

	authToken := vals1[0]
	userSession, err := NewUserSessionCacheRegistry(authToken).Get(ss.Context())
	if err != nil {
		return errorx.ErrAuthenticationFalied.WrapErr(err)
	}

	ver, err := NewAuthVerCacheRegistry(userSession.Gid).Get(ss.Context())
	if err != nil {
		return errorx.ErrAuthenticationFalied.WrapErr(err)
	}

	if ver != userSession.AuthVer {
		return errorx.ErrAuthenticationInfoExpired
	}
	return handler(srv, ss)

}

package authentication

import (
	"context"
	"infra/pkg/constant"
	"infra/pkg/errorx"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const CtxUserSessionKey = "UserSession"

func UnaryAuthenticationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	vals1 := md.Get(constant.AuthTokenKey)
	if len(vals1) == 0 {
		return nil, errorx.ErrAuthenticationFalied
	}

	authToken := vals1[0]
	userSession, err := GetUserSession(ctx, authToken)
	if err != nil {
		return nil, errorx.ErrAuthenticationFalied.WrapErr(err)
	}
	ver, err := GetAuthVer(ctx, userSession.Gid)
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
	md, _ := metadata.FromIncomingContext(ss.Context())
	vals1 := md.Get(constant.AuthTokenKey)
	if len(vals1) == 0 {
		return errorx.ErrAuthenticationFalied
	}

	authToken := vals1[0]
	userSession, err := GetUserSession(ss.Context(), authToken)
	if err != nil {
		return errorx.ErrAuthenticationFalied.WrapErr(err)
	}

	ver, err := GetAuthVer(ss.Context(), userSession.Gid)
	if err != nil {
		return errorx.ErrAuthenticationFalied.WrapErr(err)
	}
	if ver != userSession.AuthVer {
		return errorx.ErrAuthenticationInfoExpired
	}

	return handler(srv, ss)

}

package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type CurAdmin struct {
	Uid string
}

func GetCurAdmin(ctx context.Context) (admin *CurAdmin) {
	val := ctx.Value("uid")
	if val != "" {
		admin.Uid = val.(string)
		return
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if val, ok := md["uid"]; ok {
			admin.Uid = val[0]
		}
	}

	return nil
}

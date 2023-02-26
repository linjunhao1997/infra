package templates

import (
	"infra/services/system/pkg/authorization"
	"testing"
)

func TestGenResourceGoZeroApi(t *testing.T) {
	err := GenResourceGoZeroApi("resource.api.tpl", "D:/learning/infra/gateway/apidesc", "system",
		authorization.WebApi{})
	if err != nil {
		panic(err)
	}
}

func TestGenResourceGoZeroProto(t *testing.T) {
	err := GenResourceGoZeroProto("resource.proto.tpl", "D:/learning/infra/services/system/desc", "system",
		authorization.WebApi{})
	if err != nil {
		panic(err)
	}
}

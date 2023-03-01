package authorization

import (
	"infra/templates"
	"testing"
)

func TestGenResourceGoZeroApi(t *testing.T) {
	err := templates.GenResourceGoZeroApi("", "D:/learning/infra/apiserver/apidesc", "system",
		WebApi{})
	if err != nil {
		panic(err)
	}
}

func TestGenResourceGoZeroProto(t *testing.T) {
	err := templates.GenResourceGoZeroProto("", "D:/learning/infra/services/system/desc", "system",
		WebApi{})
	if err != nil {
		panic(err)
	}
}

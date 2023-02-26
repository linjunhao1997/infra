package utils

import (
	"fmt"
	"testing"
)

func TestParseTagSetting(t *testing.T) {
	for i, v := range ParseTagSetting(`queryresp;type:[]*AminRoleData;import:adminrole.api`, ";") {
		fmt.Print(i, ":", v.Name, ":", v.Value, "\n")
	}
}

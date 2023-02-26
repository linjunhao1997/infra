package dao

import (
	"fmt"
	"testing"
)

func TestSnowflakeid(t *testing.T) {
	fmt.Println("snowflakeid:", GetID())
}

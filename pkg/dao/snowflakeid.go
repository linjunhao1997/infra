package dao

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/snowflake"
)

var (
	node              *snowflake.Node
	SNOWFLAKE_NODE_ID = "SNOWFLAKE_NODE_ID"
)

func init() {
	var err error
	var id int64 = 1
	str := os.Getenv(SNOWFLAKE_NODE_ID)
	if strings.TrimSpace(str) == "" {
		log.New(os.Stderr, "", log.LstdFlags).Printf("should config env %s, default 1", SNOWFLAKE_NODE_ID)
	} else {
		parseId, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		id = parseId
	}

	node, err = snowflake.NewNode(id)
	if err != nil {
		panic(err)
	}
}

func GetID() int64 {
	return node.Generate().Int64()
}

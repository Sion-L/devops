package core

import (
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
)

func GenerateUserId(node int64) int64 {
	n, err := snowflake.NewNode(node)
	if err != nil {
		logx.Error(err)
	}

	// 生成一个唯一的 ID
	userID := n.Generate()
	return userID.Int64()
}

package user

import (
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
)

func SafeInt64ToUint64(i int64) uint64 {
	if i < 0 {
		return 0
	}
	return uint64(i)
}

func GenerateUserId(node int64) uint64 {
	n, err := snowflake.NewNode(node)
	if err != nil {
		logx.Error(err)
	}

	// 生成一个唯一的 ID
	userID := n.Generate()
	return SafeInt64ToUint64(userID.Int64())
}

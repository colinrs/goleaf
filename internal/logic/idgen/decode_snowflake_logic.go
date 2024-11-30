package idgen

import (
	"context"
	"github.com/colinrs/goleaf/pkg/snowflake"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecodeSnowflakeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDecodeSnowflakeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecodeSnowflakeLogic {
	return &DecodeSnowflakeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DecodeSnowflakeLogic) DecodeSnowflake(req *types.DecodeSnowflakeRequest) (resp *types.DecodeSnowflakeResponse, err error) {
	snowflakeInfo := snowflake.ParseSnowflakeID(req.Id)
	resp = &types.DecodeSnowflakeResponse{
		NodeID:    snowflakeInfo.NodeID,
		Epoch:     snowflakeInfo.Epoch,
		Time:      snowflakeInfo.Time,
		Timestamp: snowflakeInfo.Timestamp,
		Sequence:  snowflakeInfo.Sequence,
	}
	return
}

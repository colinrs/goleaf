package idgen

import (
	"context"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SnowflakeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSnowflakeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SnowflakeLogic {
	return &SnowflakeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SnowflakeLogic) Snowflake(req *types.SnowflakeRequest) (resp *types.SnowflakeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

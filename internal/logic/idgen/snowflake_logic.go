package idgen

import (
	"context"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"
	"github.com/colinrs/goleaf/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type SnowflakeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	snowflake snowflake.Snowflake
}

func NewSnowflakeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SnowflakeLogic {
	return &SnowflakeLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		snowflake: svcCtx.GetSnowflake(),
	}
}

func (l *SnowflakeLogic) Snowflake(req *types.SnowflakeRequest) (resp *types.SnowflakeResponse, err error) {
	ids, _ := l.snowflake.NextIDs(l.ctx, req.Step)
	resp = &types.SnowflakeResponse{
		List:  ids,
		Total: req.Step,
	}
	return
}

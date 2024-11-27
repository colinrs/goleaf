package logic

import (
	"context"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoleafLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoleafLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoleafLogic {
	return &GoleafLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoleafLogic) Goleaf(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

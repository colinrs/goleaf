package biztag

import (
	"context"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBizTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBizTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBizTagLogic {
	return &GetBizTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBizTagLogic) GetBizTag(req *types.GetBizTageRequest) (resp *types.GetBizTagResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

package biztag

import (
	"context"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBizTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBizTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBizTagLogic {
	return &UpdateBizTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBizTagLogic) UpdateBizTag(req *types.UpdateBizTageRequest) (resp *types.UpdateBizTagResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

package biztag

import (
	"context"

	"github.com/colinrs/goleaf/internal/manager"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBizTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	bizTagManager manager.BizTagManager
}

func NewGetBizTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBizTagLogic {
	return &GetBizTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		bizTagManager: manager.NewBizTagManager(ctx, svcCtx),
	}
}

func (l *GetBizTagLogic) GetBizTag(req *types.GetBizTageRequest) (resp *types.GetBizTagResponse, err error) {
	bizTagData, err := l.bizTagManager.GetBizTagById(req.Id)
	if err != nil {
		return
	}
	resp = &types.GetBizTagResponse{
		Id:          bizTagData.Id,
		BizTag:      bizTagData.BizTag,
		Description: bizTagData.Description,
		MaxID:       bizTagData.MaxID,
		Step:        bizTagData.Step,
	}
	return
}

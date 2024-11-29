package biztag

import (
	"context"

	"github.com/colinrs/goleaf/internal/manager"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBizTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	bizTagManager manager.BizTagManager
}

func NewCreateBizTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBizTagLogic {
	return &CreateBizTagLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		bizTagManager: manager.NewBizTagManager(ctx, svcCtx),
	}
}

func (l *CreateBizTagLogic) CreateBizTag(req *types.CreateBizTagRequest) (resp *types.CreateBizTagResponse, err error) {
	return l.bizTagManager.CreateBizTag(req.BizTag, req.Description, req.MaxID, req.Step)
}

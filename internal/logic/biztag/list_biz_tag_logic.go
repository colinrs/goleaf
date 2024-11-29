package biztag

import (
	"context"

	"github.com/colinrs/goleaf/internal/manager"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"
	"github.com/colinrs/goleaf/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBizTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	bizTagManager manager.BizTagManager
}

func NewListBizTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBizTagLogic {
	return &ListBizTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		bizTagManager: manager.NewBizTagManager(ctx, svcCtx),
	}
}

func (l *ListBizTagLogic) ListBizTag(req *types.ListBizTageRequest) (resp *types.ListBizTagResponse, err error) {
	offset, limit := utils.PageToOffsetLimit(req.Page, req.PageSize)
	return l.bizTagManager.GetBizTagList(offset, limit)
}

package biztag

import (
	"context"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBizTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBizTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBizTagLogic {
	return &ListBizTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBizTagLogic) ListBizTag(req *types.ListBizTageRequest) (resp *types.ListBizTagResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

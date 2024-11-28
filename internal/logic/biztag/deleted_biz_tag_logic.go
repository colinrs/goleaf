package biztag

import (
	"context"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletedBizTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletedBizTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletedBizTagLogic {
	return &DeletedBizTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletedBizTagLogic) DeletedBizTag(req *types.DeletedBizTageRequest) (resp *types.DeletedBizTagResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

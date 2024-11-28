package biztag

import (
	"context"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBizTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBizTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBizTagLogic {
	return &CreateBizTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBizTagLogic) CreateBizTag(req *types.CreateBizTagRequest) (resp *types.CreateBizTagResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

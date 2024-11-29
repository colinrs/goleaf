package biztag

import (
	"context"
	"errors"
	"github.com/colinrs/goleaf/pkg/code"
	"gorm.io/gorm"

	"github.com/colinrs/goleaf/internal/manager"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBizTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	bizTagManager manager.BizTagManager
}

func NewUpdateBizTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBizTagLogic {
	return &UpdateBizTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		bizTagManager: manager.NewBizTagManager(ctx, svcCtx),
	}
}

func (l *UpdateBizTagLogic) UpdateBizTag(req *types.UpdateBizTageRequest) (resp *types.UpdateBizTagResponse, err error) {
	resp, err = l.bizTagManager.UpdateBizTag(req)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, code.BizTagNotExist
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *UpdateBizTagLogic) GetBizTagById(id uint) (resp *types.BizTagData, err error) {
	resp, err = l.bizTagManager.GetBizTagById(id)
	if err != nil {
		return
	}
	return
}

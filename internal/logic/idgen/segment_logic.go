package idgen

import (
	"context"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SegmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSegmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SegmentLogic {
	return &SegmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SegmentLogic) Segment(req *types.SegmentRequest) (resp *types.SegmentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

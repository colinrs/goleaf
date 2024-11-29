package idgen

import (
	"context"
	"github.com/colinrs/goleaf/internal/manager"

	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SegmentLogic struct {
	logx.Logger
	ctx            context.Context
	svcCtx         *svc.ServiceContext
	segmentManager manager.SegmentManager
}

func NewSegmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SegmentLogic {
	return &SegmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		segmentManager: manager.NewSegmentManager(ctx, svcCtx),
	}
}

func (l *SegmentLogic) Segment(req *types.SegmentRequest) (resp *types.SegmentResponse, err error) {
	return l.segmentManager.Segment(req)
}

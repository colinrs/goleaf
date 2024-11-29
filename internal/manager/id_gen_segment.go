package manager

import (
	"context"

	"github.com/colinrs/goleaf/internal/repo"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"
	"github.com/colinrs/goleaf/pkg/code"

	"gorm.io/gorm"
)

type SegmentManager interface {
	Segment(req *types.SegmentRequest) (*types.SegmentResponse, error)
}

type segmentManager struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext

	bizTagRepo repo.BizTagRepo
	idGenRepo  repo.IdGen
	db         *gorm.DB
}

func NewSegmentManager(ctx context.Context, svcCtx *svc.ServiceContext) SegmentManager {
	return &segmentManager{
		ctx:        ctx,
		svcCtx:     svcCtx,
		db:         svcCtx.GetDB(ctx),
		bizTagRepo: repo.NewBizTagRepo(ctx, svcCtx),
		idGenRepo:  repo.NewIdGenRepo(ctx, svcCtx),
	}
}

func (l *segmentManager) Segment(req *types.SegmentRequest) (*types.SegmentResponse, error) {
	leafAlloc, _ := l.bizTagRepo.GetBizTagByName(l.db, req.BizTag)
	if leafAlloc == nil {
		return nil, code.BizTagNotExist
	}
	maxID := l.idGenRepo.GetSegmentMaxID(l.db, req.BizTag)
	if maxID == 0 {
		return nil, code.UnknownErr
	}
	resp := &types.SegmentResponse{
		MaxID: maxID,
		MinID: maxID - leafAlloc.Step.Int64 + 1,
		Step:  leafAlloc.Step.Int64,
	}
	return resp, nil
}

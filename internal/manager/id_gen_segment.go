package manager

import (
	"context"
	"github.com/colinrs/goleaf/internal/model"
	"github.com/zeromicro/go-zero/core/logx"

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
	leafAlloc := &model.LeafAlloc{}
	//err := l.svcCtx.LocalCache.Get(l.ctx, req.BizTag, leafAlloc)
	err := l.svcCtx.GetLocalCache().Load(l.ctx, l.GetBizTagLoader, req.BizTag, leafAlloc, 0)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("biz tag %s not exist from local cache", req.BizTag)
		leafAlloc, _ = l.bizTagRepo.GetBizTagByName(l.db, req.BizTag)
		if leafAlloc == nil {
			return nil, code.BizTagNotExist
		}
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

func (l *segmentManager) GetBizTagLoader(ctx context.Context, keys []string) ([]interface{}, error) {
	if len(keys) == 0 {
		return nil, code.BizTagNotExist
	}
	bizTag := keys[0]
	leafAlloc := &model.LeafAlloc{}
	leafAlloc, _ = l.bizTagRepo.GetBizTagByName(l.db, bizTag)
	if leafAlloc == nil {
		return nil, code.BizTagNotExist
	}
	return []interface{}{leafAlloc}, nil
}

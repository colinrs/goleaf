package manager

import (
	"context"
	"database/sql"
	"errors"
	"github.com/colinrs/goleaf/pkg/code"

	"github.com/colinrs/goleaf/internal/model"
	"github.com/colinrs/goleaf/internal/repo"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"

	"gorm.io/gorm"
)

type BizTagManager interface {
	GetBizTagList(offset, limit int) (*types.ListBizTagResponse, error)
	CreateBizTag(bizTag, description string, maxID, step int64) (resp *types.CreateBizTagResponse, err error)
	UpdateBizTag(req *types.UpdateBizTageRequest) (resp *types.UpdateBizTagResponse, err error)
	DeleteBizTag(req *types.DeletedBizTageRequest) (resp *types.DeletedBizTagResponse, err error)
	GetBizTagById(id uint) (*types.BizTagData, error)
}

type bizTagManager struct {
	ctx        context.Context
	svc        *svc.ServiceContext
	bizTagRepo repo.BizTagRepo
	db         *gorm.DB
}

func NewBizTagManager(ctx context.Context, svc *svc.ServiceContext) BizTagManager {
	return &bizTagManager{
		ctx:        ctx,
		svc:        svc,
		db:         svc.GetDB(ctx),
		bizTagRepo: repo.NewBizTagRepo(ctx, svc),
	}
}

func (m *bizTagManager) GetBizTagList(offset, limit int) (*types.ListBizTagResponse, error) {
	count, bizTags, err := m.bizTagRepo.GetBizTagList(m.db, offset, limit)
	if err != nil {
		return nil, err
	}
	var list []*types.BizTagData
	for _, bizTag := range bizTags {
		list = append(list, &types.BizTagData{
			Id:          bizTag.ID,
			BizTag:      bizTag.BizTag.String,
			Description: bizTag.Description.String,
			MaxID:       bizTag.MaxID.Int64,
			Step:        bizTag.Step.Int64,
		})
	}
	resp := &types.ListBizTagResponse{
		Total: count,
		List:  list,
	}
	return resp, nil
}

func (m *bizTagManager) CreateBizTag(bizTag, description string, maxID, step int64) (resp *types.CreateBizTagResponse, err error) {
	bizTagObj, err := m.bizTagRepo.GetBizTagByName(m.db, bizTag)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if bizTagObj != nil {
		return nil, code.BizTagAlreadyExist
	}
	id, err := m.bizTagRepo.CreateBizTag(m.db, bizTag, description, maxID, step)
	if err != nil {
		return nil, err
	}
	resp = &types.CreateBizTagResponse{
		Id: id,
	}
	return resp, nil
}

func (m *bizTagManager) UpdateBizTag(req *types.UpdateBizTageRequest) (resp *types.UpdateBizTagResponse, err error) {
	err = m.bizTagRepo.UpdateBizTag(m.db, req.Id, &model.LeafAlloc{
		Description: sql.NullString{String: req.Description, Valid: !(req.Description == "")},
		Step:        sql.NullInt64{Int64: req.Step, Valid: !(req.Step == 0)},
	})
	if err != nil {
		return nil, err
	}
	leafAlloc, err := m.bizTagRepo.GetBizTagByID(m.db, req.Id)
	if err != nil {
		return nil, err
	}
	resp = &types.UpdateBizTagResponse{
		BizTag:      leafAlloc.BizTag.String,
		Description: leafAlloc.Description.String,
		MaxID:       leafAlloc.MaxID.Int64,
		Step:        leafAlloc.Step.Int64,
	}
	return
}

func (m *bizTagManager) DeleteBizTag(req *types.DeletedBizTageRequest) (resp *types.DeletedBizTagResponse, err error) {
	return resp, m.bizTagRepo.DeleteBizTag(m.db, req.Id)
}

func (m *bizTagManager) GetBizTagById(id uint) (*types.BizTagData, error) {
	leafAlloc, err := m.bizTagRepo.GetBizTagByID(m.db, id)
	if err != nil {
		return nil, err
	}
	return &types.BizTagData{
		Id:          leafAlloc.ID,
		BizTag:      leafAlloc.BizTag.String,
		Description: leafAlloc.Description.String,
		MaxID:       leafAlloc.MaxID.Int64,
		Step:        leafAlloc.Step.Int64,
	}, nil
}

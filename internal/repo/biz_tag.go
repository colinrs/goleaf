package repo

import (
	"context"
	"database/sql"

	"github.com/colinrs/goleaf/internal/model"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type BizTagRepo interface {
	GetBizTagList(db *gorm.DB, offset, limit int) (int64, []*model.LeafAlloc, error)
	GetBizTagByID(db *gorm.DB, id uint) (*model.LeafAlloc, error)
	CreateBizTag(db *gorm.DB, bizTag, description string, maxID, step int64) (uint, error)
	UpdateBizTag(db *gorm.DB, id uint, bizTag *model.LeafAlloc) error
	DeleteBizTag(db *gorm.DB, id uint) error
	GetBizTagByName(db *gorm.DB, name string) (*model.LeafAlloc, error)
}

type bizTagRepo struct {
	logx.Logger
	ctx context.Context
	svc *svc.ServiceContext
}

func NewBizTagRepo(ctx context.Context, svc *svc.ServiceContext) BizTagRepo {
	return &bizTagRepo{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
	}
}

func (r *bizTagRepo) GetBizTagList(db *gorm.DB, offset, limit int) (int64, []*model.LeafAlloc, error) {
	var count int64
	err := db.Model(&model.LeafAlloc{}).Count(&count).Error
	if err != nil {
		return 0, nil, err
	}
	var bizTags []*model.LeafAlloc
	err = db.Offset(offset).Limit(limit).Find(&bizTags).Error
	if err != nil {
		return 0, nil, err
	}
	return count, bizTags, nil
}

func (r *bizTagRepo) GetBizTagByID(db *gorm.DB, id uint) (*model.LeafAlloc, error) {
	var bizTag model.LeafAlloc
	err := db.Where("id = ?", id).First(&bizTag).Error
	if err != nil {
		return nil, err
	}
	return &bizTag, nil
}

func (r *bizTagRepo) CreateBizTag(db *gorm.DB, bizTag, description string, maxID, step int64) (uint, error) {
	leafAlloc := &model.LeafAlloc{
		BizTag:      sql.NullString{String: bizTag, Valid: !(bizTag == "")},
		MaxID:       sql.NullInt64{Int64: maxID, Valid: maxID != 0},
		Step:        sql.NullInt64{Int64: step, Valid: step != 0},
		Description: sql.NullString{String: description, Valid: description != ""},
	}

	err := db.Create(&leafAlloc).Error
	if err != nil {
		return 0, err
	}
	return leafAlloc.ID, nil
}

func (r *bizTagRepo) UpdateBizTag(db *gorm.DB, id uint, bizTag *model.LeafAlloc) error {
	return db.Model(&model.LeafAlloc{}).Where("id = ?", id).Updates(bizTag).Error
}

func (r *bizTagRepo) DeleteBizTag(db *gorm.DB, id uint) error {
	return db.Model(&model.LeafAlloc{}).Where("id = ?", id).Delete(&model.LeafAlloc{}).Error
}

func (r *bizTagRepo) GetBizTagByName(db *gorm.DB, name string) (*model.LeafAlloc, error) {
	var bizTag model.LeafAlloc
	err := db.Where("biz_tag = ?", name).First(&bizTag).Error
	if err != nil {
		return nil, err
	}
	return &bizTag, nil
}

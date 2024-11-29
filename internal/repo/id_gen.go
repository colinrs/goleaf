package repo

import (
	"context"

	"github.com/colinrs/goleaf/internal/model"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

	"gorm.io/gorm"
)

type IdGen interface {
	GetSegmentMaxID(db *gorm.DB, bizTag string) int64
}

type idGen struct {
	logx.Logger
	ctx context.Context
	svc *svc.ServiceContext
}

func NewIdGenRepo(ctx context.Context, svc *svc.ServiceContext) IdGen {
	return &idGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
	}
}

func (r *idGen) GetSegmentMaxID(db *gorm.DB, bizTag string) int64 {
	var maxID int64
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.LeafAlloc{}).Where("biz_tag = ?", bizTag).
			UpdateColumn("max_id", gorm.Expr("max_id + step")).Error
		if err != nil {
			return err
		}
		err = tx.Model(&model.LeafAlloc{}).Where("biz_tag = ?", bizTag).Pluck("max_id", &maxID).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0
	}
	return maxID
}

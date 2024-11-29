package repo

import (
	"gorm.io/gorm"
)

type IdGen interface {
	GetSegmentIDRange(db *gorm.DB) (int64, int64)
}

type idGen struct {
}

func NewIdGen() IdGen {
	return &idGen{}
}

func (r *idGen) GetSegmentIDRange(db *gorm.DB) (int64, int64) {
	return 0, 0
}

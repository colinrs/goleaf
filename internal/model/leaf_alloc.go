package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type LeafAlloc struct {
	gorm.Model
	BizTag      sql.NullString `gorm:"biz_tag" json:"biz_tag"`          // 区分业务
	MaxID       sql.NullInt64  `gorm:"max_id" json:"max_id"`            // 该biz_tag目前所被分配的ID号段的最大值
	Step        sql.NullInt64  `gorm:"step"  json:"step"`               // 每次分配ID号段长度
	Description sql.NullString `gorm:"description"  json:"description"` // 描述
}

func (*LeafAlloc) TableName() string {
	return "leaf_alloc"
}

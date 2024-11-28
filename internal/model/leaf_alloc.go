package model

import "gorm.io/gorm"

type LeafAlloc struct {
	gorm.Model
	BizTag      string `gorm:"biz_tag" json:"biz_tag"`          // 区分业务
	MaxID       uint64 `gorm:"max_id" json:"max_id"`            // 该biz_tag目前所被分配的ID号段的最大值
	Step        int64  `gorm:"step"  json:"step"`               // 每次分配ID号段长度
	Description string `gorm:"description"  json:"description"` // 描述
}

func (*LeafAlloc) TableName() string {
	return "leaf_alloc"
}

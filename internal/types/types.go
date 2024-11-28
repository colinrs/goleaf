// Code generated by goctl. DO NOT EDIT.
package types

type BizTagData struct {
	BizTag      string `json:"biz_tag"`
	Description string `json:"description"` // 描述
	MaxID       uint64 `json:"max_id"`
	Step        uint64 `json:"step,optional,default=100"`
}

type CreateBizTagRequest struct {
	BizTag      string `json:"biz_tag"`
	Description string `json:"description"` // 描述
	MaxID       uint64 `json:"max_id"`
	Step        uint64 `json:"step,optional,default=100"`
}

type CreateBizTagResponse struct {
	Id int64 `json:"id"`
}

type DeletedBizTagResponse struct {
}

type DeletedBizTageRequest struct {
	Id int64 `json:"id"`
}

type GetBizTagResponse struct {
	BizTag      string `json:"biz_tag"`
	Description string `json:"description"` // 描述
	MaxID       uint64 `json:"max_id"`
	Step        uint64 `json:"step,optional,default=100"`
}

type GetBizTageRequest struct {
	Id int64 `form:"id"`
}

type ListBizTagResponse struct {
	Total int64         `json:"total"`
	List  []*BizTagData `json:"list"`
}

type ListBizTageRequest struct {
	Page     int `form:"page,optional,default=1"`
	PageSize int `form:"page_size,optional,default=10"`
}

type SegmentRequest struct {
	BizTag string `form:"biz_tag"`
	Step   int64  `form:"step,optional,default=100"`
}

type SegmentResponse struct {
	Ids []int64 `json:"ids"`
}

type SnowflakeRequest struct {
	BizTag string `form:"biz_tag"`
	Step   int64  `form:"step,optional,default=100"`
}

type SnowflakeResponse struct {
	Ids []int64 `json:"ids"`
}

type UpdateBizTagResponse struct {
	BizTag      string `json:"biz_tag"`
	Description string `json:"description"` // 描述
	MaxID       uint64 `json:"max_id"`
	Step        uint64 `json:"step,optional,default=100"`
}

type UpdateBizTageRequest struct {
	Id          int64  `json:"id"`
	Description string `json:"description,optional"` // 描述
	Step        int64  `json:"step,optional,default=100"`
}

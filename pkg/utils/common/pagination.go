package common

import "gorm.io/gorm"

type PageInfo struct {
	Page     int `json:"page" form:"page"` // 页码
	PageSize int `json:"pageSize" form:""` // 每页大小
}

// Paginate 定义Gorm的Scope
func Paginate(pageNum, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageNum - pageSize) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

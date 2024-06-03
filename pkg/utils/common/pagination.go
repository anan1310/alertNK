package common

import "gorm.io/gorm"

// 分页工具
type PaginationQ struct {
	Ok          bool        `json:"ok"`                                                     //代表业务查询没有出错
	ProjectName string      `json:"projectName"`                                            //对应模糊查询
	Pid         int         `json:"pid"`                                                    //对应id
	Size        int         `form:"size" json:"size"`                                       //每页显示的数量,
	Page        int         `form:"page" json:"page"`                                       // 当前页码
	Data        interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list 分页的数据内容
	Total       int64       `json:"total"`                                                  //全部的页码数量
}

// CrudAll 分页的工具类
func CrudAll(p *PaginationQ, queryTx *gorm.DB, list interface{}) (int, error) {
	//每页显示多少条size,当前是第几页page
	if p.Size < 1 {
		p.Size = 10
	}
	if p.Page < 1 {
		p.Page = 1
	}

	var total int64
	err := queryTx.Count(&total).Error
	if err != nil {
		return 0, err
	}
	offset := p.Size * (p.Page - 1)
	err = queryTx.Limit(p.Size).Offset(offset).Find(list).Error
	if err != nil {
		return 0, err
	}

	return int(total), err
}

// PageRequest 定义分页请求结构体
type PageRequest struct {
	PageNum  int `json:"pageNum,omitempty"`  //页码
	PageSize int `json:"pageSize,omitempty"` //每页大小
}

// Paginate 定义Gorm的Scope
func Paginate(pageNum, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageNum - pageSize) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"` // 页码
	PageSize int `json:"pageSize" form:""` // 每页大小
}

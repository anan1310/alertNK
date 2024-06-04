package models

type NoticeTemplateExample struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Template    string `json:"template"`
}

type NoticeQuery struct {
	TenantId string `json:"tenantId" form:"tenantId"`
	Uuid     string `json:"uuid" form:"uuid"`
	Name     string `json:"name" form:"name"`
	Query    string `json:"query" form:"query"`
}

type NoticeTemplateExampleQuery struct {
	Id    string `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Query string `json:"query" form:"query"`
}

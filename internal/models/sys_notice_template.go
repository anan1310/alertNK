package models

type NoticeTemplateExample struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Template    string `json:"template"`
}

type NoticeTemplateExampleQuery struct {
	Id    string `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Query string `json:"query" form:"query"`
}

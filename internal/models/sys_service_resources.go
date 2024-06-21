package models

type ServiceResource struct {
	ID    uint   `json:"-"`
	Time  string `json:"time" `
	Value int    `json:"value"`
	Label string `json:"label"`
}

func (ServiceResource) TableName() string {
	return "sys_service_resources"
}

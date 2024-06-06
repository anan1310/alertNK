package services

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/cmd"
)

type noticeService struct {
	ctx *ctx.Context
}

type interNoticeService interface {
	List(req interface{}) (interface{}, interface{}, interface{})
	Create(req interface{}) (interface{}, interface{})
	Update(req interface{}) (interface{}, interface{})
	Delete(req interface{}) (interface{}, interface{})
	Get(req interface{}) (interface{}, interface{})
}

func newInterAlertNoticeService(ctx *ctx.Context) interNoticeService {
	return &noticeService{
		ctx,
	}
}
func (n noticeService) Get(req interface{}) (interface{}, interface{}) {
	r := req.(*models.NoticeQuery)
	data, err := n.ctx.DB.Notice().Get(*r)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (n noticeService) List(req interface{}) (interface{}, interface{}, interface{}) {
	r := req.(*models.NoticeQuery)
	data, total, err := n.ctx.DB.Notice().List(*r)
	if err != nil {
		return nil, 0, nil
	}
	return data, total, nil
}

func (n noticeService) Create(req interface{}) (interface{}, interface{}) {
	r := req.(*models.AlertNotice)
	r.ID = "n-" + cmd.RandId()
	if err := n.ctx.DB.Notice().Create(*r); err != nil {
		return nil, err
	}
	return nil, nil
}
func (n noticeService) Update(req interface{}) (interface{}, interface{}) {
	r := req.(*models.AlertNotice)
	if err := n.ctx.DB.Notice().Update(*r); err != nil {
		return nil, err
	}
	return nil, nil
}
func (n noticeService) Delete(req interface{}) (interface{}, interface{}) {
	r := req.(*models.NoticeQuery)
	if err := n.ctx.DB.Notice().Delete(*r); err != nil {
		return nil, err
	}
	return nil, nil
}

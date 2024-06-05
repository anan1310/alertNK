package services

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
)

type dutyManagerService struct {
	ctx *ctx.Context
}

type interDutyManagerService interface {
	Create(req interface{}) interface{}
	Update(req interface{}) interface{}
	List(req interface{}) (interface{}, interface{}, interface{})
	Delete(req interface{}) interface{}
}

func newInterDutyMangerService(ctx *ctx.Context) interDutyManagerService {
	return &dutyManagerService{
		ctx: ctx,
	}
}

func (s *dutyManagerService) Create(req interface{}) interface{} {
	r := req.(*models.DutyManagement)
	if err := s.ctx.DB.DutyManager().Create(*r); err != nil {
		return err
	}
	return nil
}

func (s *dutyManagerService) List(req interface{}) (interface{}, interface{}, interface{}) {
	r := req.(*models.DutyManagementQuery)
	dutyManagements, total, err := s.ctx.DB.DutyManager().List(*r)
	if err != nil {
		return nil, 0, nil
	}
	return dutyManagements, total, nil
}

func (s *dutyManagerService) Update(req interface{}) interface{} {
	r := req.(*models.DutyManagement)
	if err := s.ctx.DB.DutyManager().Update(*r); err != nil {
		return err
	}
	return nil
}

func (s *dutyManagerService) Delete(req interface{}) interface{} {
	r := req.(*models.DutyManagementQuery)
	if err := s.ctx.DB.DutyManager().Delete(*r); err != nil {
		return err
	}
	return nil

}

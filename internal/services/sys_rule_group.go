package services

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
)

type ruleGroupService struct {
	ctx *ctx.Context
}

func newInterRuleGroupService(ctx *ctx.Context) InterRuleGroupService {
	return &ruleGroupService{
		ctx: ctx,
	}
}

type InterRuleGroupService interface {
	Create(req interface{}) (interface{}, interface{})
	Update(req interface{}) (interface{}, interface{})
	List(req interface{}) (interface{}, interface{}, interface{})
	Delete(req interface{}) interface{}
}

func (rgs *ruleGroupService) Create(req interface{}) (interface{}, interface{}) {
	r := req.(*models.RuleGroups)
	err := rgs.ctx.DB.RuleGroup().Create(*r)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (rgs *ruleGroupService) Update(req interface{}) (interface{}, interface{}) {
	r := req.(*models.RuleGroups)
	err := rgs.ctx.DB.RuleGroup().Update(*r)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (rgs *ruleGroupService) List(req interface{}) (interface{}, interface{}, interface{}) {
	r := req.(*models.RuleGroupQuery)
	ruleGroups, total, err := rgs.ctx.DB.RuleGroup().List(*r)
	if err != nil {
		return nil, 0, nil
	}
	return ruleGroups, total, nil
}

func (rgs *ruleGroupService) Delete(req interface{}) interface{} {
	r := req.(*models.RuleGroupQuery)
	if err := ctx.DB.RuleGroup().Delete(*r); err != nil {
		return err
	}
	return nil
}

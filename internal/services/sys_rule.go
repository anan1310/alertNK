package services

import (
	"alarm_collector/alert/queue"
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
)

type ruleService struct {
	ctx  *ctx.Context
	rule chan *models.AlertRule // 推送到全局的队列中
}

type interRuleService interface {
	Create(r interface{}) (interface{}, interface{})
	Update(r interface{}) (interface{}, interface{})
	Delete(r interface{}) (interface{}, interface{})
	ListRule(r interface{}) (interface{}, interface{}, interface{})
	Get(req interface{}) (interface{}, interface{})
}

func newInterRuleService(ctx *ctx.Context) interRuleService {

	return &ruleService{
		ctx:  ctx,
		rule: queue.AlertRuleChannel,
	}
}

func (rs *ruleService) Create(r interface{}) (interface{}, interface{}) {
	rule := r.(*models.AlertRule)

	if err := rs.ctx.DB.Rule().Create(*rule); err != nil {
		return nil, err
	}
	//创建完成后推送到队列中（消费者需要实时的消费，否则会阻塞）
	rs.rule <- rule

	return nil, nil
}

func (rs *ruleService) Update(r interface{}) (interface{}, interface{}) {
	rule := r.(*models.AlertRule)
	alertInfo := &models.AlertRule{}
	//判断状态
	rs.ctx.DB.DB().Model(&models.AlertRule{}).Where("tenant_id = ? AND rule_id = ?", rule.TenantId, rule.RuleId).
		First(&alertInfo)
	/*
			重启协程
		    如果 alertInfo 被启用（true）且 rule 被禁用（false），
			判断当前状态是否是false 并且 历史状态是否为true
	*/
	if *alertInfo.Enabled == true && *rule.Enabled == false {
		if cancel, exists := queue.WatchCtxMap[rule.RuleId]; exists {
			cancel()
		}
	}
	if *alertInfo.Enabled == true && *rule.Enabled == true {
		if cancel, exists := queue.WatchCtxMap[rule.RuleId]; exists {
			cancel()
		}
	}
	// 删除缓存
	// 1:扫描redis缓冲
	iter := rs.ctx.Redis.Redis().Scan(0, rule.TenantId+":"+models.FiringAlertCachePrefix+rule.RuleId+"*", 0).Iterator()
	// 2:收集所有匹配的键
	keys := make([]string, 0)
	for iter.Next() {
		key := iter.Val()
		keys = append(keys, key)
	}
	// 3:删除redis缓冲
	rs.ctx.Redis.Redis().Del(keys...)
	// 4:更新数据库中的规则信息
	if err := rs.ctx.DB.Rule().Update(*rule); err != nil {
		return nil, err
	}
	// 5:启动协程
	if *rule.Enabled {
		rs.rule <- rule
		global.Logger.Sugar().Infof("重启 RuleId 为 %s 的 Worker 进程", rule.RuleId)
	}

	return nil, nil
}

func (rs *ruleService) Delete(r interface{}) (interface{}, interface{}) {
	rule := r.(*models.AlertRuleQuery)
	info, err := rs.ctx.DB.Rule().Get(*rule)
	if err != nil {
		return nil, err
	}
	//删除规则
	err = rs.ctx.DB.Rule().Delete(*rule)
	if err != nil {
		return nil, nil
	}
	//退出该规则的协程
	if *info.Enabled {
		global.Logger.Sugar().Infof("停止 RuleId 为 %s 的 Worker 进程", rule.RuleId)
		if cancel, exists := queue.WatchCtxMap[info.RuleId]; exists {
			cancel()
		}
	}
	// 删除缓存
	// 1:扫描redis缓冲
	iter := rs.ctx.Redis.Redis().Scan(0, rule.TenantId+":"+models.FiringAlertCachePrefix+rule.RuleId+"*", 0).Iterator()
	// 2:收集所有匹配的键
	keys := make([]string, 0)
	for iter.Next() {
		key := iter.Val()
		keys = append(keys, key)
	}
	// 3:删除redis缓冲
	rs.ctx.Redis.Redis().Del(keys...)
	global.Logger.Sugar().Infof("删除队列数据 ->%s", keys)
	return nil, nil

}

func (rs *ruleService) ListRule(r interface{}) (interface{}, interface{}, interface{}) {
	rule := r.(*models.AlertRuleQuery)
	data, total, err := rs.ctx.DB.Rule().List(*rule)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (rs *ruleService) Get(r interface{}) (interface{}, interface{}) {
	rule := r.(*models.AlertRuleQuery)
	data, err := rs.ctx.DB.Rule().Get(*rule)
	if err != nil {
		return nil, err
	}
	return data, nil
}

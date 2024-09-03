package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/middleware"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type RuleApi struct{}

func (RuleApi) Create(ctx *gin.Context) {
	r := new(models.AlertRule)
	err := response.BindJson(ctx, r)
	if err != nil {
		return
	}

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Create(r)
	})
}

func (RuleApi) Update(ctx *gin.Context) {
	r := new(models.AlertRule)
	err := response.BindJson(ctx, r)
	if err != nil {
		return
	}

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Update(r)
	})
}

func (RuleApi) List(ctx *gin.Context) {
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	page := common.ToInt(ctx.Query("page"))
	pageSize := common.ToInt(ctx.Query("pageSize"))
	ruleGroupId := ctx.Query("ruleGroupId")
	ruleQuery := models.AlertRuleQuery{
		TenantId:    tid.(string),
		RuleGroupId: ruleGroupId,
		PageInfo: common.PageInfo{
			Page:     page,
			PageSize: pageSize,
		},
	}
	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.RuleService.ListRule(&ruleQuery)
	})
}

func (RuleApi) Delete(ctx *gin.Context) {

	r := new(models.AlertRuleQuery)
	err := response.BindQuery(ctx, r)
	if err != nil {
		return
	}

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)
	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Delete(r)
	})
}

func (RuleApi) Get(ctx *gin.Context) {
	r := new(models.AlertRuleQuery)
	err := response.BindQuery(ctx, r)
	if err != nil {
		return
	}

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Get(r)
	})

}

func (RuleApi) CountRule(c *gin.Context) {
	newContext := ctx.DO()
	tid, _ := c.Get(middleware.TenantIDHeaderKey)
	tenantId := tid.(string)
	var (
		alertCount   int64
		historyCount int64
		keys         []string
	)

	//告警规则条数
	newContext.DB.DB().Model(&models.AlertRule{}).Where("tenant_id = ? ", tenantId).Count(&alertCount)

	//历史条数
	newContext.DB.DB().Model(&models.AlertHisEvent{}).Where("tenant_id = ? ", tenantId).Count(&historyCount)

	//当前告警-查询redis
	cursor := uint64(0)
	pattern := tenantId + ":" + models.FiringAlertCachePrefix + "*"
	// 每次获取的键数量
	count := int64(100)

	for {
		var curKeys []string
		var err error

		curKeys, cursor, err = ctx.DO().Redis.Redis().Scan(cursor, pattern, count).Result()
		if err != nil {
			break
		}

		keys = append(keys, curKeys...)

		if cursor == 0 {
			break
		}
	}

	response.Success(c, "success", models.AlertScreen{
		RuleCount:    alertCount,
		HistoryCount: historyCount,
		RecentCount:  len(keys),
	})

}

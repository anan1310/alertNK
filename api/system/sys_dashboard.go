package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/middleware"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type DashBoardInfoApi struct {
	models.AlertCurEvent
}

type ResponseDashboardInfo struct {
	CountAlertRules   int64                    `json:"countAlertRules"`
	CurAlerts         int                      `json:"curAlerts"`
	PrometheusNum     int                      `json:"prometheusNum"` //指标监控告警
	CurAlertList      []models.AlertCurEvent   `json:"curAlertList"`
	AlarmDistribution AlarmDistribution        `json:"alarmDistribution"`
	ServiceResource   []models.ServiceResource `json:"serviceResource"`
}

type AlarmDistribution struct {
	P0 int `json:"P0"`
	P1 int `json:"P1"`
	P2 int `json:"P2"`
}

func (di DashBoardInfoApi) GetDashBoardInfo(context *gin.Context) {

	var (
		// 规则总数
		countAlertRules int64
		// 当前告警
		keys []string
		//指标监控
		prometheusNum int
	)

	c := ctx.DO()

	tid, _ := context.Get(middleware.TenantIDHeaderKey)
	tidString := tid.(string)
	// 告警分布
	alarmDistribution := make(map[string]int)
	c.DB.DB().Model(&models.AlertRule{}).Where("tenant_id = ?", tidString).Count(&countAlertRules)

	cursor := uint64(0)
	pattern := tidString + ":" + models.FiringAlertCachePrefix + "*"
	// 每次获取的键数量
	count := int64(100)

	for {
		var curKeys []string
		var err error

		curKeys, cursor, err = c.Redis.Redis().Scan(cursor, pattern, count).Result()
		if err != nil {
			break
		}

		keys = append(keys, curKeys...)

		if cursor == 0 {
			break
		}
	}

	var curAlertList []models.AlertCurEvent
	for _, v := range keys {
		alarmDistribution[c.Redis.Event().GetCache(v).Severity] += 1
		if c.Redis.Event().GetCache(v).DatasourceType == "Prometheus" {
			prometheusNum++
		}
		if len(curAlertList) >= 5 {
			continue
		}
		//告警对象 告警状态 告警内容	发生时间
		curAlertList = append(curAlertList, c.Redis.Event().GetCache(v))
	}

	var resource []models.ServiceResource
	c.DB.DB().Model(&models.ServiceResource{}).Find(&resource)

	response.Success(context, "success", ResponseDashboardInfo{
		CountAlertRules: countAlertRules,
		CurAlerts:       len(keys),
		CurAlertList:    curAlertList, //最新告警中
		PrometheusNum:   prometheusNum,
		AlarmDistribution: AlarmDistribution{
			P0: alarmDistribution["P0"],
			P1: alarmDistribution["P1"],
			P2: alarmDistribution["P2"],
		},
		ServiceResource: resource,
	})

}

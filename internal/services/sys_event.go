package services

import (
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/common"
	"encoding/json"
	"go.uber.org/zap"
	"strings"
)

type eventService struct {
	ctx *ctx.Context
}

type InterEventService interface {
	ListCurrentEvent(req interface{}) (interface{}, interface{}, interface{})
	ListHistoryEvent(req interface{}) (interface{}, interface{}, interface{})
}

func newInterEventService(ctx *ctx.Context) InterEventService {
	return &eventService{
		ctx: ctx,
	}
}

func (e eventService) ListCurrentEvent(req interface{}) (interface{}, interface{}, interface{}) {
	r := req.(*models.AlertCurEventQuery)

	iter := e.ctx.Redis.Redis().Scan(0, r.TenantId+":"+models.FiringAlertCachePrefix+"*", 0).Iterator()
	keys := make([]string, 0)

	// 遍历匹配的键
	for iter.Next() {
		key := iter.Val()
		keys = append(keys, key)
	}

	if err := iter.Err(); err != nil {
		global.Logger.Sugar().Error("当前告警列表获取失败", zap.Error(err))
		return nil, 0, err
	}

	var dataList []models.AlertCurEvent
	for _, key := range keys {
		var data models.AlertCurEvent
		info, err := e.ctx.Redis.Redis().Get(key).Result()
		if err != nil {
			return nil, 0, err
		}

		newInfo := info
		newInfo = strings.Replace(newInfo, "\"[\\", "[", 1)
		newInfo = strings.Replace(newInfo, "\\\"]\"", "\"]", 1)
		err = json.Unmarshal([]byte(newInfo), &data)
		if err != nil {
			return nil, 0, err
		}
		dataList = append(dataList, data)
	}

	// 筛选条件
	dataList = filterDataList(dataList, r)

	return dataList, int64(len(dataList)), nil

}

func (e eventService) ListHistoryEvent(req interface{}) (interface{}, interface{}, interface{}) {
	r := req.(*models.AlertHisEventQuery)
	data, total, err := e.ctx.DB.HistoryEvent().GetHistoryEvent(*r)
	if err != nil {
		return nil, 0, err
	}

	return data, total, err

}

// 筛选数据
func filterDataList(dataList []models.AlertCurEvent, r *models.AlertCurEventQuery) []models.AlertCurEvent {
	var filteredList []models.AlertCurEvent

	for _, data := range dataList {
		//告警源
		if !common.IsEmptyStr(r.DatasourceType) && data.DatasourceType != r.DatasourceType {
			continue
		}
		//告警等级
		if !common.IsEmptyStr(r.Severity) && data.Severity != r.Severity {
			continue
		}
		// 添加更多的筛选条件
		filteredList = append(filteredList, data)
	}

	return filteredList
}

package services

import (
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"fmt"
	"sync"
	"time"
)

type dutyCalendarService struct {
	ctx *ctx.Context
}

var layout = "2006-01"

type interDutyCalendarService interface {
	CreateAndUpdate(req interface{}) (interface{}, interface{})
	Update(req interface{}) (interface{}, interface{})
	List(req interface{}) (interface{}, interface{})
	GetDutyUserInfo(req interface{}) interface{}
}

func newInterDutyCalendarService(ctx *ctx.Context) interDutyCalendarService {
	return &dutyCalendarService{
		ctx: ctx,
	}
}

func (dms dutyCalendarService) CreateAndUpdate(req interface{}) (interface{}, interface{}) {
	r := req.(*models.DutyScheduleCreate)

	var (
		dutyScheduleList []models.DutySchedule
		wg               sync.WaitGroup
	)

	// 默认当前月份
	curYear, curMonth, _ := parseTime(r.Month)
	loc, _ := time.LoadLocation("Local")
	firstOfMonth := time.Date(curYear, curMonth, 1, 0, 0, 0, 0, loc)
	daysInMonth := firstOfMonth.AddDate(0, 1, -1).Day() // 获取当前月份的天数

	// 获取当前日期
	currentTime := time.Now()
	subDay := daysInMonth - currentTime.Day()
	//统计日期
	timeC := make(chan string, subDay)
	currentYear, currentMonth, currentDay := currentTime.Date()
	date := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, loc)
	// 启动生产者 goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 生产当月值班日期
		for day := 1; day <= daysInMonth; day++ {
			dutyTime := time.Date(curYear, curMonth, day, 0, 0, 0, 0, loc)
			// 如果 dutyTime 在今天之前，则跳过
			if dutyTime.Before(date) {
				continue
			}
			formattedDutyTime := fmt.Sprintf("%d-%d-%d", curYear, curMonth, day)
			timeC <- formattedDutyTime
		}
		close(timeC) // 生产者完成后关闭通道
	}()

	// 使用工作池模式处理 dutyScheduleList
	//numWorkers := 4
	dutyScheduleC := make(chan models.DutySchedule, subDay)
	doneC := make(chan struct{})
	// 用户分配控制
	userIndex := 0
	dutyPeriod := r.DutyPeriod // 用户连续分配的天数
	dutyCount := 0             // 当前用户已经分配的天数
	// 启动消费者 goroutine
	//for i := 0; i < numWorkers; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for dutyTime := range timeC {
			user := r.Users[userIndex]
			ds := models.DutySchedule{
				TenantId: r.TenantId,
				DutyId:   r.DutyId,
				Time:     dutyTime,
				DutyUser: models.Users{
					UserId:   user.UserId,
					Username: user.Username,
				},
			}
			dutyScheduleC <- ds
			dutyCount++

			if dutyCount >= dutyPeriod {
				// 切换到下一个用户
				dutyCount = 0
				userIndex = (userIndex + 1) % len(r.Users)
			}
		}
	}()
	//}

	// 关闭 dutyScheduleC 的 goroutine
	go func() {
		wg.Wait()
		close(dutyScheduleC)
	}()

	// 收集 dutyScheduleList
	go func() {
		for ds := range dutyScheduleC {
			dutyScheduleList = append(dutyScheduleList, ds)
		}
		close(doneC)
	}()

	<-doneC // 等待所有操作完成

	// 更新数据库
	for _, v := range dutyScheduleList {
		dutyScheduleInfo := dms.ctx.DB.DutyCalendar().GetCalendarInfo(r.DutyId, v.Time)
		if dutyScheduleInfo.Time != "" {
			if err := dms.ctx.DB.DutyCalendar().Update(v); err != nil {
				global.Logger.Sugar().Errorf("值班系统更新失败 %s", err)
			}
		} else {
			err := dms.ctx.DB.DutyCalendar().Create(v)
			if err != nil {
				global.Logger.Sugar().Errorf("值班系统创建失败 %s", err)
				return nil, err
			}
		}
	}

	return dutyScheduleList, nil
}

func (dms dutyCalendarService) Update(req interface{}) (interface{}, interface{}) {
	r := req.(*models.DutySchedule)
	err := dms.ctx.DB.DutyCalendar().Update(*r)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (dms dutyCalendarService) List(req interface{}) (interface{}, interface{}) {
	r := req.(*models.DutyScheduleQuery)
	data, err := dms.ctx.DB.DutyCalendar().List(*r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dms dutyCalendarService) GetDutyUserInfo(req interface{}) interface{} {
	r := req.(*models.DutyScheduleQuery)
	return dms.ctx.DB.DutyCalendar().GetDutyUserInfo(r.DutyId, r.Time)
}
func parseTime(month string) (int, time.Month, int) {
	parsedTime, err := time.Parse(layout, month)
	if err != nil {
		return 0, time.Month(0), 0
	}
	curYear, curMonth, curDay := parsedTime.Date()
	return curYear, curMonth, curDay
}

package repo

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/utils/cmd"
	"alarm_collector/pkg/utils/common"
	"gorm.io/gorm"
)

type (
	RuleRepo struct {
		entryRepo
	}
	interRuleRepo interface {
		Get(r models.AlertRuleQuery) (models.AlertRule, error)
		List(r models.AlertRuleQuery) ([]models.AlertRule, int64, error)
		Create(r models.AlertRule) error
		Update(r models.AlertRule) error
		Delete(r models.AlertRuleQuery) error
		GetRuleIsExist(ruleId string) bool
	}
)

func newRuleInterface(db *gorm.DB, g InterGormDBCli) interRuleRepo {
	return &RuleRepo{
		entryRepo{
			g:  g,
			db: db,
		},
	}
}

func (rr RuleRepo) List(r models.AlertRuleQuery) ([]models.AlertRule, int64, error) {
	var (
		data  []models.AlertRule
		total int64
		db    = rr.db.Model(&models.AlertRule{})
	)
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)

	db.Where("tenant_id = ? AND rule_group_id = ?", r.TenantId, r.RuleGroupId)
	if !common.IsEmptyStr(r.NoticeId) {
		db = db.Where("notice_id = ?", r.NoticeId)
	}
	//查询总条数
	db.Count(&total)
	err := db.Limit(limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (rr RuleRepo) Create(r models.AlertRule) error {
	nr := r
	nr.RuleId = "r-" + cmd.RandId()

	err := rr.g.Create(models.AlertRule{}, nr)
	if err != nil {
		return err
	}

	return nil
}

func (rr RuleRepo) Update(r models.AlertRule) error {
	u := Updates{
		Table: &models.AlertRule{},
		Where: map[string]interface{}{
			"tenant_id = ?": r.TenantId,
			"rule_id = ?":   r.RuleId,
		},
		Omit:    "rule_id",
		Updates: r,
	}

	err := rr.g.Updates(u)
	if err != nil {
		return err
	}

	return nil
}

func (rr RuleRepo) Delete(r models.AlertRuleQuery) error {
	var alertRule models.AlertRule
	d := Delete{
		Table: alertRule,
		Where: map[string]interface{}{
			"tenant_id = ?": r.TenantId,
			"rule_id = ?":   r.RuleId,
		},
	}

	err := rr.g.Delete(d)
	if err != nil {
		return err
	}

	return nil
}

func (rr RuleRepo) Get(r models.AlertRuleQuery) (models.AlertRule, error) {
	var data models.AlertRule

	db := rr.db.Model(&models.AlertRule{})
	db.Where("tenant_id = ? AND rule_group_id = ? AND rule_id = ?", r.TenantId, r.RuleGroupId, r.RuleId)
	err := db.First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (rr RuleRepo) GetRuleIsExist(ruleId string) bool {
	var ruleNum int64
	rr.DB().Model(&models.AlertRule{}).
		Where("rule_id = ? AND enabled = ?", ruleId, "1").
		Count(&ruleNum)
	if ruleNum > 0 {
		return true
	}

	return false
}

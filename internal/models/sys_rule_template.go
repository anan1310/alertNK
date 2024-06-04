package models

type RuleTemplateGroup struct {
	Name        string `json:"name"`
	Number      int    `json:"number"`
	Description string `json:"description"`
}

type RuleTemplateGroupQuery struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

type RuleTemplate struct {
	RuleGroupName  string `json:"ruleGroupName"`
	RuleName       string `json:"ruleName"`
	DatasourceType string `json:"datasourceType"`
	Severity       int64  `json:"severity"`
	EvalInterval   int64  `json:"evalInterval"`
	ForDuration    int64  `json:"forDuration"`
	Annotations    string `json:"annotations"`
}

type RuleTemplateQuery struct {
	RuleGroupName  string `json:"ruleGroupName" form:"ruleGroupName"`
	RuleName       string `json:"ruleName" form:"ruleName"`
	DatasourceType string `json:"datasourceType" form:"datasourceType"`
	Severity       int64  `json:"severity" form:"severity"`
	Annotations    string `json:"annotations" form:"annotations"`
}

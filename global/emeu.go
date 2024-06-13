package global

// AlertLevel 定义告警等级类型
type AlertLevel int

const (
	p0 AlertLevel = iota
	p1
	p2
	p3
)

// 为 AlertLevel 类型实现 String 方法，方便打印
func (a AlertLevel) String() string {
	return [...]string{"紧急", "严重", "一般", "未知"}[a]
}

func ParseAlertLevel(level string) AlertLevel {
	switch level {
	case "P0":
		return p0
	case "P1":
		return p1
	case "P2":
		return p2
	default:
		return p3
	}

}

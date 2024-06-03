package init_database

import (
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量，适合多数据源
func Gorm(dbType string) *gorm.DB {
	switch dbType {
	case "mysql":
		return GormMysql()
	case "clickhouse":
		return GormClickHouse()
	default:
		return GormMysql()
	}
}

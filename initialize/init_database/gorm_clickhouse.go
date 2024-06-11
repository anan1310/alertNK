package init_database

import (
	"alarm_collector/global"
	"go.uber.org/zap"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func GormClickHouse() *gorm.DB {
	c := global.Config.Clickhouse
	if c.DataBase == "" {
		return nil
	}
	if db, err := gorm.Open(clickhouse.Open(c.Dsn()), &gorm.Config{
		//SkipDefaultTransaction: false, // 禁用默认事务 如果开启 ck提示插入不成功
		//PrepareStmt: true, // 缓冲预编译语句 默认为false
		//NamingStrategy: schema.NamingStrategy{
		//	SingularTable: true,
		//},
	}); err != nil {
		global.Logger.Sugar().Error("connect clickhouse error", zap.Error(err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
		return db
	}
}

package init_database

import (
	"alarm_collector/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func GormMysql() *gorm.DB {
	m := global.Config.MySQL
	if m.DataBase == "" {
		return nil
	}
	DB, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		PrepareStmt:            true, // 缓冲预编译语句
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		global.Logger.Error("mysql连接失败，请检查配置文件是否正确", zap.Error(err))
		//如果mysql连接失败，直接退出程序
		os.Exit(1)
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return DB
}

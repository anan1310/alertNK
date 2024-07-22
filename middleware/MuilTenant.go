package middleware

import (
	"alarm_collector/internal/cache"
	"alarm_collector/internal/ck"
	"alarm_collector/pkg/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

// DBResolver 管理多个数据库连接
type DBResolver struct {
	resolver sync.Map // 使用 sync.Map 安全地存储多个数据库连接
	mu       sync.Mutex
}

// TenantDBConfig 存储租户数据库配置信息
type TenantDBConfig struct {
	ID             uint   `gorm:"column:id; primaryKey"`
	Tenant         string `gorm:"column:tenant"`
	URL            string `gorm:"column:url"`
	TcpURL         string `gorm:"column:tcp_url"`
	Username       string `gorm:"column:username"`
	Password       string `gorm:"column:password"`
	DatabaseName   string `gorm:"column:database_name"`
	HostName       string `gorm:"column:host_name"`
	CreateTime     string `gorm:"column:create_time"`
	Status         string `gorm:"column:status"`
	ExpirationDate string `gorm:"column:expiration_date"`
}

func (TenantDBConfig) TableName() string {
	return "master_tenant"
}

var (
	resolver DBResolver // 数据库解析器
)

func MilTenant(masterDB *gorm.DB, rCache cache.InterEntryCache, ckRepo ck.InterEntryRepo) gin.HandlerFunc {
	return func(c *gin.Context) {

		value, exists := c.Get(TenantIDHeaderKey)
		if !exists {
			c.Abort()
			response.Fail(c, "当前租户不存在", "failed")
			return
		}
		// 获取租户数据库配置信息
		dbConfig, err := getTenantDBConfig(masterDB, value.(string))
		if err != nil {
			c.Abort()
			response.Fail(c, "获取租户配置信息错误", "failed")
			return
		}

		// 获取或连接租户数据库
		connect, err := resolver.GetOrConnect(dbConfig)
		if err != nil {
			response.Fail(c, "租户链接错误", "failed")
			c.Abort()
			return
		}
		//如果配置多数据源 需要在这里将connect传入
		fmt.Printf("MySQL 的地址: %p\n", &connect)
		/*
			dbRepo := repo.NewMySQLRepoEntry(connect)
			newContext := ctx.NewContext(context.Background(), dbRepo, rCache, ckRepo)
			fmt.Printf("newContext 的地址: %p\n", &newContext)
			services.NewServices(newContext)
			// 启用告警评估协程
			alert.Initialize(newContext)
			//  初始化监控分析数据 (统计协程数)
			resource.InitResource(newContext)
		*/
		c.Next()
	}
}

// 根据租户ID获取数据库配置信息
func getTenantDBConfig(db *gorm.DB, tenantID string) (TenantDBConfig, error) {
	var dbConfig TenantDBConfig
	if err := db.Where("tenant = ?", tenantID).First(&dbConfig).Error; err != nil {
		return TenantDBConfig{}, err
	}
	return dbConfig, nil
}

// GetOrConnect 获取或连接到租户数据库
func (r *DBResolver) GetOrConnect(config TenantDBConfig) (*gorm.DB, error) {
	// 从 resolver 中获取现有连接，如果不存在则创建新连接
	if db, ok := r.resolver.Load(config.Tenant); ok {
		return db.(*gorm.DB), nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// 再次检查，避免并发情况下重复创建连接
	if db, ok := r.resolver.Load(config.Tenant); ok {
		return db.(*gorm.DB), nil
	}

	// 创建新的数据库连接
	newDB, err := connectToTenantDB(config)
	if err != nil {
		return nil, err
	}

	// 将新连接存储在 resolver 中
	r.resolver.Store(config.Tenant, newDB)

	return newDB, nil
}

// connectToTenantDB 连接到租户数据库
func connectToTenantDB(config TenantDBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.TcpURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

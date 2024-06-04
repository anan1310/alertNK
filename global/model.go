package global

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"create_at"`  // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

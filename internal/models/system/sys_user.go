package system

import (
	"time"
)

// SysUser 用户表
type SysUser struct {
	UserID      int64  `json:"user_id"`                                 // 用户ID
	DeptID      *int64 `json:"dept_id,omitempty"`                       // 部门ID
	UserName    string `json:"user_name"`                               // 用户账号
	NickName    string `json:"nick_name"`                               // 用户昵称
	UserType    string `json:"user_type,omitempty"`                     // 用户类型（00系统用户）
	Email       string `json:"email,omitempty"`                         // 用户邮箱
	PhoneNumber string `json:"phone_number" gorm:"column:phonenumber"`  // 手机号码
	EmailStatus string `json:"email_status" gorm:"column:email_status"` //邮箱状态
	SmsStatus   string `json:"sms_status" gorm:"column:sms_status"`     //手机号状态
	Sex         string `json:"sex,omitempty"`                           // 用户性别（0男 1女 2未知）
	//Avatar      string `json:"avatar,omitempty"`               // 头像地址
	//Password    string     `json:"password,omitempty"`              // 密码
	Status     string     `json:"status,omitempty"`      // 帐号状态（0正常 1停用）
	DelFlag    string     `json:"del_flag,omitempty"`    // 删除标志（0代表存在 2代表删除）
	LoginIP    string     `json:"login_ip,omitempty"`    // 最后登录IP
	LoginDate  *time.Time `json:"login_date,omitempty"`  // 最后登录时间
	CreateBy   string     `json:"create_by,omitempty"`   // 创建者
	CreateTime *time.Time `json:"create_time,omitempty"` // 创建时间
	UpdateBy   string     `json:"update_by,omitempty"`   // 更新者
	UpdateTime *time.Time `json:"update_time,omitempty"` // 更新时间
	Remark     string     `json:"remark,omitempty"`      // 备注
}

func (SysUser) TableName() string {
	return "sys_user"
}

/*
joinForeignKey： 表示当前表在中间表中的外键，即当前表中的哪个字段被用作关联中间表的外键。
在你的例子中，joinForeignKey:UserId 表示 SysUser 表中的 UserId 字段被用作关联中间表 sys_user_role 的外键。

joinReferences： 表示关联表在中间表中的外键参考，即关联表中的哪个字段被用作关联中间表的外键。
在你的例子中，joinReferences:RoleID 表示 SysRole 表中的 RoleID 字段被用作关联中间表 sys_user_role 的外键参考

foreignKey:表示当前表在
*/

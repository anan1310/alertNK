package system

import (
	"alarm_collector/global"
)

// SysUser 用户表
type SysUser struct {
	global.BaseModel
	UserId   int    `json:"userId" gorm:"column:user_id;not null;unique;primary_key;comment:用户ID;size:20"` // 角色ID
	Username string `json:"userName" gorm:"column:user_name;comment:用户登录名"`                                // 用户登录名
	Password string `json:"passWord"  gorm:"column:pass_word;comment:用户登录密码"`                              // 用户登录密码
	NickName string `json:"nickName" gorm:"column:nick_name;default:系统用户;comment:用户昵称"`                    //用户昵称
	Status   string `json:"status" gorm:"column:status;comment:用户状态"`                                      // 用户状态 0=正常,1=停用
	LoginIp  string `json:"loginIp" gorm:"column:login_ip;comment:登陆IP"`                                   // 登陆IP
	DelFlag  string `json:"delFlag" gorm:"column:del_flag;default:0;comment:删除标志（0代表存在 2代表删除)"`            // 删除标志（0代表存在 2代表删除）
}

/*
joinForeignKey： 表示当前表在中间表中的外键，即当前表中的哪个字段被用作关联中间表的外键。
在你的例子中，joinForeignKey:UserId 表示 SysUser 表中的 UserId 字段被用作关联中间表 sys_user_role 的外键。

joinReferences： 表示关联表在中间表中的外键参考，即关联表中的哪个字段被用作关联中间表的外键。
在你的例子中，joinReferences:RoleID 表示 SysRole 表中的 RoleID 字段被用作关联中间表 sys_user_role 的外键参考

foreignKey:表示当前表在
*/

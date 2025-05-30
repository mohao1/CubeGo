package entity

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID         int32     `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:用户id" json:"id"`            // 用户id
	Account    string    `gorm:"column:account;type:varchar(128);not null;comment:用户账户" json:"account"`              // 用户账户
	Name       string    `gorm:"column:name;type:varchar(128);not null;comment:用户名称" json:"name"`                    // 用户名称
	Phone      string    `gorm:"column:phone;type:varchar(128);not null;comment:用户手机" json:"phone"`                  // 用户手机
	Email      string    `gorm:"column:email;type:varchar(128);not null;comment:用户邮箱" json:"email"`                  // 用户邮箱
	Password   string    `gorm:"column:password;type:varchar(128);not null;comment:用户密码" json:"password"`            // 用户密码
	Role       int32     `gorm:"column:role;type:int;not null;comment:用户身份" json:"role"`                             // 用户身份
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;comment:创建时间" json:"create_time"`          // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;comment:更新时间" json:"update_time"`          // 更新时间
	IsDelete   int32     `gorm:"column:is_delete;type:int;not null;comment:是否已被删除（0:flase,1:ture）" json:"is_delete"` // 是否已被删除（0:flase,1:ture）
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

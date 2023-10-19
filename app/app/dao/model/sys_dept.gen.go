// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysDept = "sys_dept"

// SysDept mapped from table <sys_dept>
type SysDept struct {
	DeptID     int64     `gorm:"column:dept_id;primaryKey;autoIncrement:true;comment:部门id" json:"dept_id"` // 部门id
	ParentID   int64     `gorm:"column:parent_id;comment:父部门id" json:"parent_id"`                          // 父部门id
	Ancestors  string    `gorm:"column:ancestors;comment:祖级列表" json:"ancestors"`                           // 祖级列表
	DeptName   string    `gorm:"column:dept_name;comment:部门名称" json:"dept_name"`                           // 部门名称
	OrderNum   int32     `gorm:"column:order_num;comment:显示顺序" json:"order_num"`                           // 显示顺序
	Leader     string    `gorm:"column:leader;comment:负责人" json:"leader"`                                  // 负责人
	Phone      string    `gorm:"column:phone;comment:联系电话" json:"phone"`                                   // 联系电话
	Email      string    `gorm:"column:email;comment:邮箱" json:"email"`                                     // 邮箱
	Status     string    `gorm:"column:status;default:0;comment:部门状态（0正常 1停用）" json:"status"`              // 部门状态（0正常 1停用）
	DelFlag    string    `gorm:"column:del_flag;default:0;comment:删除标志（0代表存在 2代表删除）" json:"del_flag"`      // 删除标志（0代表存在 2代表删除）
	CreateBy   string    `gorm:"column:create_by;comment:创建者" json:"create_by"`                            // 创建者
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                       // 创建时间
	UpdateBy   string    `gorm:"column:update_by;comment:更新者" json:"update_by"`                            // 更新者
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                       // 更新时间
}

// TableName SysDept's table name
func (*SysDept) TableName() string {
	return TableNameSysDept
}
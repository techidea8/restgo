package restgo

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	Id       uint     `json:"id" form:"id" gorm:"primarykey"`
	CreateAt DateTime `json:"createAt"   form:"updateAt" gorm:"comment:创建时间" time_format:"2006-01-02 15:04:05" time_utc:"1"` //加入时间
	UpdateAt DateTime `json:"updateAt"   form:"updateAt" gorm:"comment:更新时间" time_format:"2006-01-02 15:04:05" time_utc:"1"` //加入时间
	DeleteAt DateTime `json:"deleteAt"   form:"deleteAt" gorm:"comment:删除时间" time_format:"2006-01-02 15:04:05" time_utc:"1"` //加入时间
	Deleted  bool     `json:"deleted"   form:"deleted" gorm:"comment:删除状态"`                                                  //加入时间
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateAt = DateTimeNow()
	m.Deleted = false
	m.UpdateAt = DateTimeNow()
	m.DeleteAt = DateTimeNow()
	return
}
func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdateAt = DateTimeNow()
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	m.DeleteAt = DateTimeNow()
	m.Deleted = true
	return
}

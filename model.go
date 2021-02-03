package restgo

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	Id       uint     `json:"id" form:"id" gorm:"primarykey"`
	CreateAt DateTime `json:"create_at"   form:"create_at" time_format:"2006-01-02 15:04:05" time_utc:"1"` //加入时间
	UpdateAt DateTime `json:"update_at"   form:"update_at" time_format:"2006-01-02 15:04:05" time_utc:"1"` //加入时间
	DeleteAt DateTime `json:"delete_at"   form:"delete_at" time_format:"2006-01-02 15:04:05" time_utc:"1"` //加入时间
	Deleted  int      `json:"deleted"   form:"deleted" `                                                   //加入时间
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateAt = DateTimeNow()
	m.Deleted = 0
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
	m.Deleted = 1
	return
}

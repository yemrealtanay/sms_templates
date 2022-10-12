package models

import "time"

type SmsTemplateType struct {
	Sms_template_type_id uint `gorm:"primaryKey"`
	Name                 string
	Description          string
	Key                  string
	Created_at           time.Time
	Updated_at           time.Time
	Deleted_at           time.Time
}

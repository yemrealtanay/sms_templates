package models

import "time"

type SmsTemplate struct {
	Sms_template_id          uint `gorm:"primaryKey"`
	Company_id               int
	Branch_id                int
	Name                     string
	Subject                  string
	Content                  string
	Subscription_type_id     int
	Sms_template_category_id int
	Created_by               int
	Created_at               time.Time `gorm:"autoCreateTime; <-:create"`
	Updated_at               time.Time `gorm:"autoUpdateTime; <-:update"`
	Deleted_at               time.Time `gorm:"autoUpdateTime; <-:update"`
	Sms_template_type_id     int
	Activity_id              int
	Is_edited                bool
}

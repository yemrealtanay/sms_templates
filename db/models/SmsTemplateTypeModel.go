package models

import (
	"database/sql"
	"time"
)

type SmsTemplateType struct {
	SmsTemplateTypeID int32        `json:"sms_template_type_id"`
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	Key               string       `json:"key"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         sql.NullTime `json:"updated_at"`
	DeletedAt         sql.NullTime `json:"deleted_at"`
}

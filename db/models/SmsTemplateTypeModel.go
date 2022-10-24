package models

import (
	"database/sql"
)

type SmsTemplateType struct {
	SmsTemplateTypeID int32          `json:"sms_template_type_id"`
	Name              sql.NullString `json:"name"`
	Description       sql.NullString `json:"description"`
	Key               sql.NullString `json:"key"`
	CreatedAt         sql.NullTime   `json:"created_at"`
	UpdatedAt         sql.NullTime   `json:"updated_at"`
	DeletedAt         sql.NullTime   `json:"deleted_at"`
}

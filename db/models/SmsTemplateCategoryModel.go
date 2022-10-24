package models

import (
	"database/sql"
	"time"
)

type SmsTemplateCategory struct {
	SmsTemplateCategoryID int32        `json:"sms_template_category_id"`
	CompanyID             int32        `json:"company_id"`
	BranchID              int32        `json:"branch_id"`
	Name                  string       `json:"name"`
	Description           string       `json:"description"`
	CreatedBy             int32        `json:"created_by"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             sql.NullTime `json:"updated_at"`
	DeletedAt             sql.NullTime `json:"deleted_at"`
}

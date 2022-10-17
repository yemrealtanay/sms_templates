package models

import "database/sql"

type SmsTemplateCategory struct {
	SmsTemplateCategoryID int32          `json:"sms_template_category_id"`
	CompanyID             sql.NullInt32  `json:"company_id"`
	BranchID              sql.NullInt32  `json:"branch_id"`
	Name                  sql.NullString `json:"name"`
	Description           sql.NullString `json:"description"`
	CreatedBy             sql.NullInt32  `json:"created_by"`
	CreatedAt             sql.NullTime   `json:"created_at"`
	UpdatedAt             sql.NullTime   `json:"updated_at"`
	DeletedAt             sql.NullTime   `json:"deleted_at"`
}

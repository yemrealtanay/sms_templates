package models

import "database/sql"

type SmsTemplate struct {
	SmsTemplateID         int32          `json:"sms_template_id"`
	CompanyID             sql.NullInt32  `json:"company_id"`
	BranchID              sql.NullInt32  `json:"branch_id"`
	Name                  sql.NullString `json:"name"`
	Subject               sql.NullString `json:"subject"`
	Content               sql.NullString `json:"content"`
	SubscriptionTypeID    sql.NullInt32  `json:"subscription_type_id"`
	SmsTemplateCategoryID sql.NullInt32  `json:"sms_template_category_id"`
	CreatedBy             sql.NullInt32  `json:"created_by"`
	SmsTemplateTypeID     sql.NullInt32  `json:"sms_template_type_id"`
	ActivityID            sql.NullInt32  `json:"activity_id"`
	IsEdited              sql.NullBool   `json:"is_edited"`
	CreatedAt             sql.NullTime   `json:"created_at"`
	UpdatedAt             sql.NullTime   `json:"updated_at"`
	DeletedAt             sql.NullTime   `json:"deleted_at"`
}

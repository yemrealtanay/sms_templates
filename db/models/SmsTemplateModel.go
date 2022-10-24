package models

import (
	"database/sql"
	"time"
)

type SmsTemplate struct {
	SmsTemplateID         int32        `json:"sms_template_id"`
	CompanyID             int32        `json:"company_id"`
	BranchID              int32        `json:"branch_id"`
	Name                  string       `json:"name"`
	Subject               string       `json:"subject"`
	Content               string       `json:"content"`
	SubscriptionTypeID    int32        `json:"subscription_type_id"`
	SmsTemplateCategoryID int32        `json:"sms_template_category_id"`
	CreatedBy             int32        `json:"created_by"`
	SmsTemplateTypeID     int32        `json:"sms_template_type_id"`
	ActivityID            int32        `json:"activity_id"`
	IsEdited              bool         `json:"is_edited"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             sql.NullTime `json:"updated_at"`
	DeletedAt             sql.NullTime `json:"deleted_at"`
}

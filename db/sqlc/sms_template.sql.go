// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: sms_template.sql

package db

import (
	"context"
	"database/sql"
	"github.com/yemrealtanay/sms_templates/db/models"
)

const createSmsTemplate = `-- name: CreateSmsTemplate :one
INSERT INTO sms_templates (
    company_id,
    branch_id,
    name,
    subject,
    content,
    subscription_type_id,
    sms_template_category_id,
    created_by,
    sms_template_type_id,
    activity_id,
    is_edited
) VALUES (
          $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING sms_template_id, company_id, branch_id, name, subject, content, subscription_type_id, sms_template_category_id, created_by, sms_template_type_id, activity_id, is_edited, created_at, updated_at, deleted_at
`

type CreateSmsTemplateParams struct {
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
}

func (q *Queries) CreateSmsTemplate(ctx context.Context, arg CreateSmsTemplateParams) (models.SmsTemplate, error) {
	row := q.db.QueryRowContext(ctx, createSmsTemplate,
		arg.CompanyID,
		arg.BranchID,
		arg.Name,
		arg.Subject,
		arg.Content,
		arg.SubscriptionTypeID,
		arg.SmsTemplateCategoryID,
		arg.CreatedBy,
		arg.SmsTemplateTypeID,
		arg.ActivityID,
		arg.IsEdited,
	)
	var i models.SmsTemplate
	err := row.Scan(
		&i.SmsTemplateID,
		&i.CompanyID,
		&i.BranchID,
		&i.Name,
		&i.Subject,
		&i.Content,
		&i.SubscriptionTypeID,
		&i.SmsTemplateCategoryID,
		&i.CreatedBy,
		&i.SmsTemplateTypeID,
		&i.ActivityID,
		&i.IsEdited,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteSmsTemplate = `-- name: DeleteSmsTemplate :exec
DELETE FROM sms_templates WHERE sms_template_id = $1
`

func (q *Queries) DeleteSmsTemplate(ctx context.Context, smsTemplateID int32) error {
	_, err := q.db.ExecContext(ctx, deleteSmsTemplate, smsTemplateID)
	return err
}

const getSmsTemplateByActivity = `-- name: GetSmsTemplateByActivity :many
SELECT sms_template_id, company_id, branch_id, name, subject, content, subscription_type_id, sms_template_category_id, created_by, sms_template_type_id, activity_id, is_edited, created_at, updated_at, deleted_at FROM sms_templates
WHERE activity_id = $1 and company_id = $2
`

type GetSmsTemplateByActivityParams struct {
	ActivityID sql.NullInt32 `json:"activity_id"`
	CompanyID  sql.NullInt32 `json:"company_id"`
}

func (q *Queries) GetSmsTemplateByActivity(ctx context.Context, arg GetSmsTemplateByActivityParams) ([]models.SmsTemplate, error) {
	rows, err := q.db.QueryContext(ctx, getSmsTemplateByActivity, arg.ActivityID, arg.CompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.SmsTemplate
	for rows.Next() {
		var i models.SmsTemplate
		if err := rows.Scan(
			&i.SmsTemplateID,
			&i.CompanyID,
			&i.BranchID,
			&i.Name,
			&i.Subject,
			&i.Content,
			&i.SubscriptionTypeID,
			&i.SmsTemplateCategoryID,
			&i.CreatedBy,
			&i.SmsTemplateTypeID,
			&i.ActivityID,
			&i.IsEdited,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSmsTemplateByCategory = `-- name: GetSmsTemplateByCategory :many
SELECT sms_template_id, company_id, branch_id, name, subject, content, subscription_type_id, sms_template_category_id, created_by, sms_template_type_id, activity_id, is_edited, created_at, updated_at, deleted_at FROM sms_templates
WHERE sms_template_category_id = $1 and company_id $2
`

type GetSmsTemplateByCategoryParams struct {
	SmsTemplateCategoryID sql.NullInt32 `json:"sms_template_category_id"`
	Column2               interface{}   `json:"column_2"`
}

func (q *Queries) GetSmsTemplateByCategory(ctx context.Context, arg GetSmsTemplateByCategoryParams) ([]models.SmsTemplate, error) {
	rows, err := q.db.QueryContext(ctx, getSmsTemplateByCategory, arg.SmsTemplateCategoryID, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.SmsTemplate
	for rows.Next() {
		var i models.SmsTemplate
		if err := rows.Scan(
			&i.SmsTemplateID,
			&i.CompanyID,
			&i.BranchID,
			&i.Name,
			&i.Subject,
			&i.Content,
			&i.SubscriptionTypeID,
			&i.SmsTemplateCategoryID,
			&i.CreatedBy,
			&i.SmsTemplateTypeID,
			&i.ActivityID,
			&i.IsEdited,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSmsTemplateByID = `-- name: GetSmsTemplateByID :one
SELECT sms_template_id, company_id, branch_id, name, subject, content, subscription_type_id, sms_template_category_id, created_by, sms_template_type_id, activity_id, is_edited, created_at, updated_at, deleted_at FROM sms_templates
WHERE sms_template_id = $1
`

func (q *Queries) GetSmsTemplateByID(ctx context.Context, smsTemplateID int32) (models.SmsTemplate, error) {
	row := q.db.QueryRowContext(ctx, getSmsTemplateByID, smsTemplateID)
	var i models.SmsTemplate
	err := row.Scan(
		&i.SmsTemplateID,
		&i.CompanyID,
		&i.BranchID,
		&i.Name,
		&i.Subject,
		&i.Content,
		&i.SubscriptionTypeID,
		&i.SmsTemplateCategoryID,
		&i.CreatedBy,
		&i.SmsTemplateTypeID,
		&i.ActivityID,
		&i.IsEdited,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getSmsTemplateByTypeAndCategory = `-- name: GetSmsTemplateByTypeAndCategory :many
SELECT sms_template_id, company_id, branch_id, name, subject, content, subscription_type_id, sms_template_category_id, created_by, sms_template_type_id, activity_id, is_edited, created_at, updated_at, deleted_at FROM sms_templates
WHERE sms_template_type_id = $1 and sms_template_category_id = $2 and company_id = $3
`

type GetSmsTemplateByTypeAndCategoryParams struct {
	SmsTemplateTypeID     sql.NullInt32 `json:"sms_template_type_id"`
	SmsTemplateCategoryID sql.NullInt32 `json:"sms_template_category_id"`
	CompanyID             sql.NullInt32 `json:"company_id"`
}

func (q *Queries) GetSmsTemplateByTypeAndCategory(ctx context.Context, arg GetSmsTemplateByTypeAndCategoryParams) ([]models.SmsTemplate, error) {
	rows, err := q.db.QueryContext(ctx, getSmsTemplateByTypeAndCategory, arg.SmsTemplateTypeID, arg.SmsTemplateCategoryID, arg.CompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.SmsTemplate
	for rows.Next() {
		var i models.SmsTemplate
		if err := rows.Scan(
			&i.SmsTemplateID,
			&i.CompanyID,
			&i.BranchID,
			&i.Name,
			&i.Subject,
			&i.Content,
			&i.SubscriptionTypeID,
			&i.SmsTemplateCategoryID,
			&i.CreatedBy,
			&i.SmsTemplateTypeID,
			&i.ActivityID,
			&i.IsEdited,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSmsTemplateByTypes = `-- name: GetSmsTemplateByTypes :many
SELECT sms_template_id, company_id, branch_id, name, subject, content, subscription_type_id, sms_template_category_id, created_by, sms_template_type_id, activity_id, is_edited, created_at, updated_at, deleted_at FROM sms_templates
WHERE sms_template_type_id = $1 and company_id $2
`

type GetSmsTemplateByTypesParams struct {
	SmsTemplateTypeID sql.NullInt32 `json:"sms_template_type_id"`
	Column2           interface{}   `json:"column_2"`
}

func (q *Queries) GetSmsTemplateByTypes(ctx context.Context, arg GetSmsTemplateByTypesParams) ([]models.SmsTemplate, error) {
	rows, err := q.db.QueryContext(ctx, getSmsTemplateByTypes, arg.SmsTemplateTypeID, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.SmsTemplate
	for rows.Next() {
		var i models.SmsTemplate
		if err := rows.Scan(
			&i.SmsTemplateID,
			&i.CompanyID,
			&i.BranchID,
			&i.Name,
			&i.Subject,
			&i.Content,
			&i.SubscriptionTypeID,
			&i.SmsTemplateCategoryID,
			&i.CreatedBy,
			&i.SmsTemplateTypeID,
			&i.ActivityID,
			&i.IsEdited,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSmsTemplateByUser = `-- name: GetSmsTemplateByUser :many
SELECT sms_template_id, company_id, branch_id, name, subject, content, subscription_type_id, sms_template_category_id, created_by, sms_template_type_id, activity_id, is_edited, created_at, updated_at, deleted_at FROM sms_templates
WHERE created_by = $1 and company_id = $2
`

type GetSmsTemplateByUserParams struct {
	CreatedBy sql.NullInt32 `json:"created_by"`
	CompanyID sql.NullInt32 `json:"company_id"`
}

func (q *Queries) GetSmsTemplateByUser(ctx context.Context, arg GetSmsTemplateByUserParams) ([]models.SmsTemplate, error) {
	rows, err := q.db.QueryContext(ctx, getSmsTemplateByUser, arg.CreatedBy, arg.CompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.SmsTemplate
	for rows.Next() {
		var i models.SmsTemplate
		if err := rows.Scan(
			&i.SmsTemplateID,
			&i.CompanyID,
			&i.BranchID,
			&i.Name,
			&i.Subject,
			&i.Content,
			&i.SubscriptionTypeID,
			&i.SmsTemplateCategoryID,
			&i.CreatedBy,
			&i.SmsTemplateTypeID,
			&i.ActivityID,
			&i.IsEdited,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSmsTemplates = `-- name: GetSmsTemplates :many
SELECT sms_template_id, company_id, branch_id, name, subject, content, subscription_type_id, sms_template_category_id, created_by, sms_template_type_id, activity_id, is_edited, created_at, updated_at, deleted_at FROM sms_templates
         WHERE company_id = $1
ORDER BY sms_template_id
`

func (q *Queries) GetSmsTemplates(ctx context.Context, companyID sql.NullInt32) ([]models.SmsTemplate, error) {
	rows, err := q.db.QueryContext(ctx, getSmsTemplates, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.SmsTemplate
	for rows.Next() {
		var i models.SmsTemplate
		if err := rows.Scan(
			&i.SmsTemplateID,
			&i.CompanyID,
			&i.BranchID,
			&i.Name,
			&i.Subject,
			&i.Content,
			&i.SubscriptionTypeID,
			&i.SmsTemplateCategoryID,
			&i.CreatedBy,
			&i.SmsTemplateTypeID,
			&i.ActivityID,
			&i.IsEdited,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSmsTemplate = `-- name: UpdateSmsTemplate :exec
UPDATE sms_templates SET name= $1, subject= $2, content= $3, is_edited= $4, sms_template_type_id= $5, sms_template_category_id= $6, activity_id= $7, subscription_type_id= $8
WHERE sms_template_id = $9
`

type UpdateSmsTemplateParams struct {
	Name                  sql.NullString `json:"name"`
	Subject               sql.NullString `json:"subject"`
	Content               sql.NullString `json:"content"`
	IsEdited              sql.NullBool   `json:"is_edited"`
	SmsTemplateTypeID     sql.NullInt32  `json:"sms_template_type_id"`
	SmsTemplateCategoryID sql.NullInt32  `json:"sms_template_category_id"`
	ActivityID            sql.NullInt32  `json:"activity_id"`
	SubscriptionTypeID    sql.NullInt32  `json:"subscription_type_id"`
	SmsTemplateID         int32          `json:"sms_template_id"`
}

func (q *Queries) UpdateSmsTemplate(ctx context.Context, arg UpdateSmsTemplateParams) (int32, error) {
	_, err := q.db.ExecContext(ctx, updateSmsTemplate,
		arg.Name,
		arg.Subject,
		arg.Content,
		arg.IsEdited,
		arg.SmsTemplateTypeID,
		arg.SmsTemplateCategoryID,
		arg.ActivityID,
		arg.SubscriptionTypeID,
		arg.SmsTemplateID,
	)
	return 0, err
}

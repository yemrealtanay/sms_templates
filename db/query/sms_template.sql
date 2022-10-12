-- name: CreateSmsTemplate :one
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
) RETURNING *;
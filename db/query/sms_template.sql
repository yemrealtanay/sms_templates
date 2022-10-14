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

-- name: GetSmsTemplates :many
SELECT * FROM sms_templates
         WHERE company_id = $1
ORDER BY sms_template_id;

-- name: GetSmsTemplateByID :one
SELECT * FROM sms_templates
WHERE sms_template_id = $1;

-- name: GetSmsTemplateByCategory :many
SELECT * FROM sms_templates
WHERE sms_template_category_id = $1 and company_id $2;

-- name: GetSmsTemplateByTypes :many
SELECT * FROM sms_templates
WHERE sms_template_type_id = $1 and company_id $2;

-- name: UpdateSmsTemplate :exec
UPDATE sms_templates SET name= $1, subject= $2, content= $3, is_edited= $4
WHERE sms_template_id = $5;

-- name: DeleteSmsTemplate :exec
DELETE FROM sms_templates WHERE sms_template_id = $1;



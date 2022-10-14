-- name: CreateSmsTemplateCategory :one
INSERT INTO sms_template_categories (company_id,
                                branch_id,
                                name,
                                description,
                                created_by)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: UpdateSmsTemplateCategory :exec
UPDATE sms_template_categories SET name = $1, description = $2, company_id = $3, branch_id = $4
WHERE sms_template_category_id = $5;

-- name: DeleteSmsTemplateCategory :exec
DELETE FROM sms_template_categories WHERE sms_template_category_id = $1;
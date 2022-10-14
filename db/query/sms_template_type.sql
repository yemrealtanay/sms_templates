-- name: CreateSmsTemplateType :one
INSERT INTO sms_template_types (name,
                           description,
                           key)
VALUES ($1, $2, $3)
RETURNING *;


-- name: UpdateSmsTemplateType :exec
UPDATE sms_template_types SET name = $1, description = $2, key = $3
WHERE sms_template_type_id = $4;

-- name: DeleteSmsTemplateType :exec
DELETE FROM sms_template_types WHERE sms_template_type_id = $1;
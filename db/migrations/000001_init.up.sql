CREATE TABLE "sms_templates"
(
    "sms_template_id"          SERIAL PRIMARY KEY NOT NULL,
    "company_id"               integer,
    "branch_id"                integer,
    "name"                     text,
    "subject"                  text,
    "content"                  text,
    "subscription_type_id"     integer,
    "sms_template_category_id" integer,
    "created_by"               integer,
    "sms_template_type_id"     integer,
    "activity_id"              integer,
    "is_edited"                boolean,
    "created_at"               timestamp DEFAULT 'now()',
    "updated_at"               timestamp,
    "deleted_at"               timestamp
);

CREATE TABLE "sms_template_types"
(
    "sms_template_type_id" SERIAL PRIMARY KEY NOT NULL,
    "name"                 text,
    "description"          text,
    "key"                  text,
    "created_at"           timestamp DEFAULT 'now()',
    "updated_at"           timestamp,
    "deleted_at"           timestamp
);

CREATE TABLE "sms_template_categories"
(
    "sms_template_category_id" SERIAL PRIMARY KEY NOT NULL,
    "company_id"               integer,
    "branch_id"                integer,
    "name"                     text,
    "description"              text,
    "created_by"               integer,
    "created_at"               timestamp DEFAULT 'now()',
    "updated_at"               timestamp,
    "deleted_at"               timestamp
);

CREATE INDEX ON "sms_templates" ("sms_template_type_id");

CREATE INDEX ON "sms_templates" ("sms_template_category_id");

CREATE INDEX ON "sms_template_types" ("key");

CREATE INDEX ON "sms_template_categories" ("company_id");

ALTER TABLE "sms_templates"
    ADD FOREIGN KEY ("sms_template_category_id") REFERENCES "sms_template_categories" ("sms_template_category_id");

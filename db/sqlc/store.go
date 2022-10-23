package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	//Querier
	CreateSmsTemplate(ctx context.Context, arg CreateSmsTemplateParams) (CreateSmsTemplateTxResult, error)
	UpdateSmsTemplateTx(ctx context.Context, arg UpdateSmsTemplateParams) (UpdateSmsTemplateTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx.err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type UpdateSmsTemplateTxResult struct {
	SmsTemplateID         int32         `json:"sms_template_id"`
	Name                  string        `json:"name"`
	Subject               string        `json:"subject"`
	Content               string        `json:"content"`
	SmsTemplateTypeID     sql.NullInt32 `json:"sms_template_type_id"`
	SmsTemplateCategoryID sql.NullInt32 `json:"sms_template_category_id"`
	ActivityID            sql.NullInt32 `json:"activity_id"`
	IsEdited              bool          `json:"is_edited"`
}

func (store *SQLStore) UpdateSmsTemplateTx(ctx context.Context, arg UpdateSmsTemplateParams) (UpdateSmsTemplateTxResult, error) {
	var result UpdateSmsTemplateTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.SmsTemplateID, err = q.UpdateSmsTemplate(ctx, UpdateSmsTemplateParams{
			SmsTemplateID:         arg.SmsTemplateID,
			Name:                  arg.Name,
			Subject:               arg.Subject,
			Content:               arg.Content,
			SmsTemplateTypeID:     arg.SmsTemplateTypeID,
			SmsTemplateCategoryID: arg.SmsTemplateCategoryID,
			ActivityID:            arg.ActivityID,
			IsEdited:              arg.IsEdited,
		})
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

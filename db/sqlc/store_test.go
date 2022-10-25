package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/yemrealtanay/sms_templates/util"
	"testing"
)

func TestSmsTemplateTx(t *testing.T) {
	store := NewStore(testDB)

	smsTemplateCategory := createRandomCategory(t)
	smsTemplateType := createRandomType(t)
	fmt.Println("Category ID and Type ID", smsTemplateCategory.SmsTemplateCategoryID, smsTemplateType.SmsTemplateTypeID)

	n := 5

	errs := make(chan error)
	results := make(chan CreateSmsTemplateTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreateSmsTemplateTx(context.Background(), CreateSmsTemplateParams{
				CompanyID: sql.NullInt32{
					Int32: util.RandomInt(1, 500),
					Valid: true,
				},
				BranchID: sql.NullInt32{
					Int32: util.RandomInt(1, 500),
					Valid: true,
				},
				Name: sql.NullString{
					String: util.RandomName(),
					Valid:  true,
				},
				Subject: sql.NullString{
					String: util.RandomSubject(),
					Valid:  true,
				},
				Content: sql.NullString{
					String: util.RandomSentence(),
					Valid:  true,
				},
				SubscriptionTypeID: sql.NullInt32{
					Int32: util.RandomInt(1, 100),
					Valid: true,
				},
				SmsTemplateCategoryID: sql.NullInt32{
					Int32: smsTemplateCategory.SmsTemplateCategoryID,
					Valid: true,
				},
				CreatedBy: sql.NullInt32{
					Int32: util.RandomUserID(),
					Valid: true,
				},
				SmsTemplateTypeID: sql.NullInt32{
					Int32: smsTemplateType.SmsTemplateTypeID,
					Valid: true,
				},
				ActivityID: sql.NullInt32{
					Int32: util.RandomInt(1, 500),
					Valid: true,
				},
				IsEdited: sql.NullBool{
					Bool:  false,
					Valid: true,
				},
			})

			errs <- err
			results <- result
		}()
	}

}

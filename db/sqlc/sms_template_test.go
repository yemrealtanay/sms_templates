package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/yemrealtanay/sms_templates/db/models"
	"github.com/yemrealtanay/sms_templates/util"
	"testing"
)

func createRandomSmsTemplate(t *testing.T) models.SmsTemplate {
	category := createRandomCategory(t)
	templateType := createRandomType(t)
	arg := CreateSmsTemplateParams{
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
			Int32: category.SmsTemplateCategoryID,
			Valid: true,
		},
		CreatedBy: sql.NullInt32{
			Int32: util.RandomUserID(),
			Valid: true,
		},
		SmsTemplateTypeID: sql.NullInt32{
			Int32: templateType.SmsTemplateTypeID,
			Valid: true,
		},
		ActivityID: sql.NullInt32{
			Int32: util.RandomInt(1, 100),
			Valid: true,
		},
		IsEdited: sql.NullBool{
			Bool:  util.RandomBool(),
			Valid: true,
		},
	}

	smsTemplate, err := testQueries.CreateSmsTemplate(context.Background(), arg)
	require.NoError(t, err)

	return smsTemplate
}

func TestCreateSmsTemplate(t *testing.T) {
	createRandomSmsTemplate(t)
}

func TestGetSmsTemplate(t *testing.T) {
	smsTemplate1 := createRandomSmsTemplate(t)
	smsTemplate2, err := testQueries.GetSmsTemplateByID(context.Background(), smsTemplate1.SmsTemplateID)
	require.NoError(t, err)
	require.NotEmpty(t, smsTemplate2)

}

func TestUpdateSmsTemplate(t *testing.T) {
	oldSmsTemplate := createRandomSmsTemplate(t)

	newName := util.RandomName()
	newSubject := util.RandomSubject()
	newContent := util.RandomSentence()
	newSmsTemplateTypeId := util.RandomInt(1, 100)
	newSmsTemplateCategoryId := util.RandomInt(1, 50)
	newActivityID := util.RandomInt(1, 500)
	newSubscriptionTypeId := util.RandomInt(1, 100)

	_, err := testQueries.UpdateSmsTemplate(context.Background(), UpdateSmsTemplateParams{
		Name: sql.NullString{
			String: newName,
			Valid:  true,
		},
		Subject: sql.NullString{
			String: newSubject,
			Valid:  true,
		},
		Content: sql.NullString{
			String: newContent,
			Valid:  true,
		},
		IsEdited: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		SmsTemplateTypeID: sql.NullInt32{
			Int32: newSmsTemplateTypeId,
			Valid: true,
		},
		SmsTemplateCategoryID: sql.NullInt32{
			Int32: newSmsTemplateCategoryId,
			Valid: true,
		},
		ActivityID: sql.NullInt32{
			Int32: newActivityID,
			Valid: true,
		},
		SubscriptionTypeID: sql.NullInt32{
			Int32: newSubscriptionTypeId,
			Valid: true,
		},
		SmsTemplateID: oldSmsTemplate.SmsTemplateID,
	})

	require.NoError(t, err)
}

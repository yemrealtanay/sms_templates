package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/yemrealtanay/sms_templates/db/models"
	"github.com/yemrealtanay/sms_templates/util"
	"testing"
)

func createRandomSmsTemplate(t *testing.T) models.SmsTemplate {
	arg := CreateSmsTemplateParams{
		CompanyID:             util.RandomInt(1, 500),
		BranchID:              util.RandomInt(1, 500),
		Name:                  util.RandomName(),
		Subject:               util.RandomSubject(),
		Content:               util.RandomSentence(),
		SubscriptionTypeID:    util.RandomInt(1, 100),
		SmsTemplateCategoryID: util.RandomInt(1, 50),
		CreatedBy:             util.RandomUserID(),
		SmsTemplateTypeID:     util.RandomInt(1, 100),
		ActivityID:            util.RandomInt(1, 100),
		IsEdited:              util.RandomBool(),
	}

	smsTemplate, err := testQueries.CreateSmsTemplate(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.CompanyID, smsTemplate.CompanyID)
	require.Equal(t, arg.BranchID, smsTemplate.BranchID)
	require.Equal(t, arg.Name, smsTemplate.Name)
	require.Equal(t, arg.Subject, smsTemplate.Subject)
	require.Equal(t, arg.Content, smsTemplate.Content)
	require.Equal(t, arg.SubscriptionTypeID, smsTemplate.SubscriptionTypeID)
	require.Equal(t, arg.SmsTemplateCategoryID, smsTemplate.SmsTemplateCategoryID)
	require.Equal(t, arg.CreatedBy, smsTemplate.CreatedBy)
	require.Equal(t, arg.SmsTemplateTypeID, smsTemplate.SmsTemplateTypeID)
	require.Equal(t, arg.ActivityID, smsTemplate.ActivityID)
	require.Equal(t, arg.IsEdited, smsTemplate.IsEdited)
	require.NotZero(t, smsTemplate.CreatedAt)

	return smsTemplate
}

func TestCreateSmsTemplate(t *testing.T) {
	createRandomSmsTemplate(t)
}

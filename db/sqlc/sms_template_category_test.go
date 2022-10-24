package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/yemrealtanay/sms_templates/db/models"
	"github.com/yemrealtanay/sms_templates/util"
	"testing"
)

func createRandomCategory(t *testing.T) models.SmsTemplateCategory {

	arg := CreateSmsTemplateCategoryParams{
		CompanyID:   util.RandomInt(1, 500),
		BranchID:    util.RandomInt(1, 500),
		Name:        util.RandomName(),
		Description: util.RandomSentence(),
		CreatedBy:   util.RandomUserID(),
	}

	smsTemplateCategory, err := testQueries.CreateSmsTemplateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.CompanyID, smsTemplateCategory.CompanyID)
	require.Equal(t, arg.BranchID, smsTemplateCategory.BranchID)
	require.Equal(t, arg.Name, smsTemplateCategory.Name)
	require.Equal(t, arg.Description, smsTemplateCategory.Description)
	require.Equal(t, arg.CreatedBy, smsTemplateCategory.CreatedBy)

	return smsTemplateCategory

}

func TestCreateSmsTemplateCategory(t *testing.T) {
	createRandomCategory(t)
}

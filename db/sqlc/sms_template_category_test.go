package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/yemrealtanay/sms_templates/db/models"
	"github.com/yemrealtanay/sms_templates/util"
	"testing"
)

func createRandomCategory(t *testing.T) models.SmsTemplateCategory {

	arg := CreateSmsTemplateCategoryParams{
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
		Description: sql.NullString{
			String: util.RandomSentence(),
			Valid:  true,
		},
		CreatedBy: sql.NullInt32{
			Int32: util.RandomUserID(),
			Valid: true,
		},
	}

	smsTemplateCategory, err := testQueries.CreateSmsTemplateCategory(context.Background(), arg)
	require.NoError(t, err)
	return smsTemplateCategory

}

func TestCreateSmsTemplateCategory(t *testing.T) {
	createRandomCategory(t)
}

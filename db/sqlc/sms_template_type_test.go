package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/yemrealtanay/sms_templates/db/models"
	"github.com/yemrealtanay/sms_templates/util"
	"testing"
)

func createRandomType(t *testing.T) models.SmsTemplateType {
	arg := CreateSmsTemplateTypeParams{
		Name:        util.RandomName(),
		Description: util.RandomSentence(),
		Key:         util.RandomKey(),
	}

	smsTemplateType, err := testQueries.CreateSmsTemplateType(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Name, smsTemplateType.Name)
	require.Equal(t, arg.Description, smsTemplateType.Description)
	require.Equal(t, arg.Key, smsTemplateType.Key)

	return smsTemplateType
}

func TestCreateSmsTemplateType(t *testing.T) {
	createRandomType(t)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yemrealtanay/sms_templates/initializers"
	"github.com/yemrealtanay/sms_templates/models"
	"time"
)

func SmsTemplateCreate(c *gin.Context) {

	sms_tamplate := models.TestSmsTemplate{
		Company_id:               184,
		Branch_id:                234,
		Name:                     "test",
		Subject:                  "test subject",
		Content:                  "test Content",
		Subscription_type_id:     2,
		Sms_template_category_id: 4,
		Created_by:               234,
		Created_at:               time.Now(),
		Sms_template_type_id:     2,
		Activity_id:              35,
		Is_edited:                false,
	}

	result := initializers.DB.Create(&sms_tamplate)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"sms_template": sms_tamplate,
	})
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yemrealtanay/sms_templates/controllers"
	"github.com/yemrealtanay/sms_templates/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/sms_template", controllers.SmsTemplateCreate)

	r.Run()
}

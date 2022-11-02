package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/yemrealtanay/sms_templates/db/sqlc"
	"net/http"
)

type createSmsTemplateRequest struct {
	CompanyID             sql.NullInt32  `json:"company_id" binding:"required"`
	BranchID              sql.NullInt32  `json:"branch_id" binding:"required"`
	Name                  sql.NullString `json:"name" binding:"required"`
	Subject               sql.NullString `json:"subject" binding:"required"`
	Content               sql.NullString `json:"content" binding:"required"`
	SubscriptionTypeID    sql.NullInt32  `json:"subscription_type_id" binding:"required"`
	SmsTemplateCategoryID sql.NullInt32  `json:"sms_template_category_id" binding:"required"`
	CreatedBy             sql.NullInt32  `json:"created_by" binding:"required"`
	SmsTemplateTypeID     sql.NullInt32  `json:"sms_template_type_id" binding:"required"`
	ActivityID            sql.NullInt32  `json:"activity_id" binding:"required"`
	IsEdited              sql.NullBool   `json:"is_edited" binding:"required"`
}

func (server *Server) createSmsTemplate(ctx *gin.Context) {
	var req createSmsTemplateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSmsTemplateParams{
		CompanyID:             req.CompanyID,
		BranchID:              req.BranchID,
		Name:                  req.Name,
		Subject:               req.Subject,
		Content:               req.Content,
		SubscriptionTypeID:    req.SubscriptionTypeID,
		SmsTemplateCategoryID: req.SmsTemplateCategoryID,
		CreatedBy:             req.CreatedBy,
		SmsTemplateTypeID:     req.SmsTemplateTypeID,
		ActivityID:            req.ActivityID,
		IsEdited:              req.IsEdited,
	}

	sms_template, err := server.store.CreateSmsTemplate(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sms_template)
}

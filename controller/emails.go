package controller

import (
	"email-rest-api-modak/domain/models"
	"email-rest-api-modak/domain/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendNewEmailNotification(c *gin.Context) {
	var input models.EmailNotificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailRequest := models.EmailNotificationRequest{
		Type:    input.Type,
		Message: input.Message,
		UserID:  input.UserID,
	}

	response := usecases.SendEmailNotificationUseCase(emailRequest)
	c.JSON(response.Response.StatusCode, response)
}

package gateway

import (
	"email-rest-api-modak/domain/models"
	"fmt"
	"net/http"
)

func SendEmailNotification(request models.EmailNotificationRequest) models.EmailNotificationResponse {
	var message models.EmailNotificationResponse
	message = models.EmailNotificationResponse{
		Body:    fmt.Sprintf("Sent notification type %s message to user ID: %s,", request.Type, request.UserID),
		Message: request.Message,
		Response: http.Response{
			Status:     "Successful",
			StatusCode: http.StatusAccepted,
		},
	}
	return message
}

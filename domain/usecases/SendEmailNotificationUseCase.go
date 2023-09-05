package usecases

import (
	"email-rest-api-modak/domain/models"
	"email-rest-api-modak/gateway"
)

func SendEmailNotificationUseCase(request models.EmailNotificationRequest) models.EmailNotificationResponse {
	return gateway.SendEmailNotification(request)
}

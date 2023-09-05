package usecases

import (
	"email-rest-api-modak/constants"
	"email-rest-api-modak/domain/models"
	"testing"
)

func TestSendEmailNotificationUseCase(t *testing.T) {

	tests := []models.EmailNotificationRequest{
		{
			Type:    constants.STATUS_TYPE,
			Message: "Lorem ipsum",
			UserID:  "1",
		},
		{
			Type:    constants.NEWS_TYPE,
			Message: "Lorem ipsum",
			UserID:  "2",
		},
		{
			Type:    constants.MARKETING_TYPE,
			Message: "Lorem ipsum",
			UserID:  "3",
		},
		{
			Type:    "other",
			Message: "Lorem ipsum",
			UserID:  "4",
		},
	}

	for _, tt := range tests {
		t.Run("Different status cases", func(t *testing.T) {
			SendEmailNotificationUseCase(tt)
		})
	}
}

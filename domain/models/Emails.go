package models

import "net/http"

type EmailNotificationInput struct {
	Type    string `json:"type" binding:"required"`
	Message string `json:"message" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
}
type EmailNotificationRequest struct {
	Type    string
	Message string
	UserID  string
}

type EmailNotificationResponse struct {
	Body     string `json:"body"`
	Message  string `json:"message"`
	Response http.Response
}

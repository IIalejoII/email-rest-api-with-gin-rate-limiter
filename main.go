package main

import (
	"email-rest-api-modak/controller"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"log"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())

	rate, err := limiter.NewRateFromFormatted("2-M")
	if err != nil {
		log.Fatal(err)
		return
	}

	storeWithPrefix := memory.NewStoreWithOptions(
		limiter.StoreOptions{
			Prefix: "default:send_status_email_notification:",
		},
	)
	middlewareForStatusType := mgin.NewMiddleware(limiter.New(storeWithPrefix, rate))

	rate2, err2 := limiter.NewRateFromFormatted("1-D")
	if err2 != nil {
		log.Fatal(err)
		return
	}

	storeWithPrefix2 := memory.NewStoreWithOptions(
		limiter.StoreOptions{
			Prefix: "default:send_news_email_notification:",
		},
	)
	middlewareForNewsType := mgin.NewMiddleware(limiter.New(storeWithPrefix2, rate2))

	rate3, err3 := limiter.NewRateFromFormatted("3-H")
	if err3 != nil {
		log.Fatal(err)
		return
	}

	storeWithPrefix3 := memory.NewStoreWithOptions(
		limiter.StoreOptions{
			Prefix: "default:send_marketing_email_notification:",
		},
	)
	middlewareForMarketingType := mgin.NewMiddleware(limiter.New(storeWithPrefix3, rate3))
	// Routes
	api := router.Group("/api")
	api.POST("/send_status_email_notification", middlewareForStatusType, controller.SendNewEmailNotification)
	api.POST("/send_news_email_notification", middlewareForNewsType, controller.SendNewEmailNotification)
	api.POST("/send_marketing_email_notification", middlewareForMarketingType, controller.SendNewEmailNotification)

	err = router.Run("localhost:8080")
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	}
}

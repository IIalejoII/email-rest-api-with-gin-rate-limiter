# Golang Rate-Limited Notification Service

## How to test this?

Please install go.1.18 is required for this project, next use the following command from the root folder: 
``` go
 go run main.go
```

After that the project should be running on the following link [http://localhost:8080/](http://localhost:8080/)

On the root folder you will find a postman collection that you can use to test this more easy, you just have to import that collection into your postman folder.

If not then you can use these curls commands from your terminal to test each endpoint:

```bash
curl --location 'localhost:8080/api/send_status_email_notification' \
--header 'Content-Type: application/json' \
--data '{
    "type": "Status",
    "message": "You have a new notification",
    "user_id": "1020"
}'

curl --location 'localhost:8080/api/send_marketing_email_notification' \
--header 'Content-Type: application/json' \
--data '{
    "type": "Marketing",
    "message": "You have a new notification",
    "user_id": "2222"
}'

curl --location 'localhost:8080/api/send_news_email_notification' \
--header 'Content-Type: application/json' \
--data '{
    "type": "News",
    "message": "You have a new notification",
    "user_id": "1111"
}'
```
package request

type NotifRequest struct {
	To string `json:"to"`
	NotificationType string `json:"notification_type"`
	NotificationId int32 `json:"notification_id"`
	Notification interface{} `json:"notification"`
	Data Data `json:"data"`
}

type Data struct {
	Title string `json:"title"`
	Body string `json:"body"`
	Channel string `json:"channel"`
	Category string `json:"category"`
}
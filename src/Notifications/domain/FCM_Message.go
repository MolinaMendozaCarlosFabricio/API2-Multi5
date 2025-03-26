package domain

type FCMMessage struct {
	To   string `json:"to"`
	Notification struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"notification"`
}
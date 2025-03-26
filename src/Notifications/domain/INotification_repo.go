package domain

type INotification_repo interface{
	SaveNotification(notification Notification, type_noti int)(int, error)
	GetNotificationsOfAnUser(id_user int)([]Notification, error)
	GetOneNotification(id int)([]Notification, error)
	DeleteNotification(id int)error
}
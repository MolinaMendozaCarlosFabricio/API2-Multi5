package application

import "api2-multi.com/a/src/Notifications/domain"

type GetOneNotificationUC struct {
	db domain.INotification_repo
}

func NewGetOneNotificationUC(db domain.INotification_repo)*GetOneNotificationUC{
	return&GetOneNotificationUC{db: db}
}

func(uc *GetOneNotificationUC)Execute(id int)([]domain.Notification, error){
	return uc.db.GetOneNotification(id)
}
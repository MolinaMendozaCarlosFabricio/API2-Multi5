package application

import "api2-multi.com/a/src/Notifications/domain"

type DeleteNotificationUC struct {
	db domain.INotification_repo
}

func NewDeleteNotificationUC(db domain.INotification_repo)*DeleteNotificationUC{
	return&DeleteNotificationUC{db: db}
}

func(uc *DeleteNotificationUC)Execute(id int)error{
	return uc.db.DeleteNotification(id)
}
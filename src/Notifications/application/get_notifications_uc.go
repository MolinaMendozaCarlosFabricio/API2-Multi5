package application

import "api2-multi.com/a/src/Notifications/domain"

type GetNotificationsOfAnUserUC struct {
	db domain.INotification_repo
}

func NewGetNotificationsOfAnUserUC(db domain.INotification_repo)*GetNotificationsOfAnUserUC{
	return&GetNotificationsOfAnUserUC{db: db}
}

func(uc *GetNotificationsOfAnUserUC)Execute(id_user int)([]domain.Notification, error){
	return uc.db.GetNotificationsOfAnUser(id_user)
}
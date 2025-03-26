package application

import (
	"time"

	"api2-multi.com/a/src/Notifications/domain"
)

type SaveNotificationUC struct {
	db domain.INotification_repo
}

func NewSaveNotificationUC(db domain.INotification_repo)*SaveNotificationUC{
	return&SaveNotificationUC{db: db}
}

func(uc *SaveNotificationUC)Execute(id_user, id_parcel, id_type_notification int)(domain.Notification, error){
	notification := &domain.Notification{
		ID: 0,
		Id_user: id_user,
		Id_parcel: id_parcel,
		Type_notification: "",
		Date_hour: time.Now().Format("2006-01-02 15:04:05"),
	}
	_, err := uc.db.SaveNotification(*notification, id_type_notification)
	return *notification, err
}
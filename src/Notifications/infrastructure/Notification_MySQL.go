package infrastructure

import (
	"log"

	"api2-multi.com/a/src/Notifications/domain"
	"api2-multi.com/a/src/core"
)

type NotificationsMySQL struct {
	conn core.ConectionMySQL
}

func NewNotificationsMySQL()*NotificationsMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&NotificationsMySQL{conn: *conn}
}

func(r *NotificationsMySQL)SaveNotification(notification domain.Notification, type_noti int)(int, error){
	query := "INSERT INTO notifications (id_user, id_parcel, id_type_notification, date_hour) VALUES (?,?,?,?)"
	res, err := r.conn.ExecPreparedQuerys(
		query, 
		notification.Id_user, 
		notification.Id_parcel, 
		type_noti,
		notification.Date_hour,
	)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	last_id, err := res.LastInsertId()
	if err != nil {
        log.Fatalf("Error al obtener último usuario registrado:", err)
    }
	return int(last_id), err
}

func(r *NotificationsMySQL)GetNotificationsOfAnUser(id_user int)([]domain.Notification, error){
	query := "SELECT notifications.id_notification, notifications.id_parcel, type_notification.type_name, notifications.date_hour FROM notifications INNER JOIN type_notification ON type_notification.id_type_notification = notifications.id_notification WHERE notifications.id_user = ?"
	rows, err := r.conn.FetchRows(query, id_user)
	var notifications []domain.Notification
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var id_parcel int
		var type_notification string
		var date_hour string

		if err := rows.Scan(&id, &id_parcel, &type_notification, &date_hour); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		notification := &domain.Notification{
			ID: id, 
			Id_user: id_user, 
			Id_parcel: id_parcel, 
			Type_notification: type_notification,
			Date_hour: date_hour,
		}
		notifications = append(notifications, *notification)
	}
	return notifications, err
}

func(r *NotificationsMySQL)GetOneNotification(id int)([]domain.Notification, error){
	query := "SELECT notifications.id_notification, notifications.id_user, notifications.id_parcel, type_notification.type_name, notifications.date_hour FROM notifications INNER JOIN type_notification ON type_notification.id_type_notification = notifications.id_notification WHERE notifications.id_notification = ?"
	rows, err := r.conn.FetchRows(query, id)
	var notifications []domain.Notification
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var id_user int
		var id_parcel int
		var type_notification string
		var date_hour string

		if err := rows.Scan(&id, &id_user, &id_parcel, &type_notification, &date_hour); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		notification := &domain.Notification{
			ID: id, 
			Id_user: id_user, 
			Id_parcel: id_parcel, 
			Type_notification: type_notification,
		}
		notifications = append(notifications, *notification)
	}
	return notifications, err
}

func(r *NotificationsMySQL)DeleteNotification(id int)error{
	query := "DELETE FROM notifications WHERE id_notification = ?"
	_, err := r.conn.ExecPreparedQuerys(query, id)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}
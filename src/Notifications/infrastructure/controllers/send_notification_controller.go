package controllers

import (
	"net/http"

	"api2-multi.com/a/src/Notifications/application"
	"api2-multi.com/a/src/Notifications/application/services"
	"api2-multi.com/a/src/Notifications/infrastructure"
	"github.com/gin-gonic/gin"
)

type SendNotificationC struct {
	uc application.SaveNotificationUC
	service_verification services.ValidateMeasurementService
}

func NewSendNotificationUC()*SendNotificationC{
	
	mysql := infrastructure.NewNotificationsMySQL()
	uc := application.NewSaveNotificationUC(mysql)

	validation_mysql := infrastructure.NewValidationMySQL()
	service1 := services.NewValidateMeasurementService(validation_mysql)

	return&SendNotificationC{
		uc: *uc,
		service_verification: *service1,
	}
}

func(controller *SendNotificationC)Execute(c *gin.Context){
	var input struct{
		Id_user int `json:"id_user"`
		Id_parcel int `json:"id_parcel"`
		Hum float32 `json:"hum"`
		Temp float32 `json:"temp"`
		Air float32 `json:"air"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
			"Err": err,
		})
		return
	}

	type_notification, flag, err := controller.service_verification.Execute(
		input.Id_parcel, input.Hum, input.Temp, input.Air,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al validar los datos",
			"Err": err,
		})
		return
	}

	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Mediciones normales, sin advertencias",
		})
		return
	}

	notification, err := controller.uc.Execute(input.Id_user, input.Id_parcel, type_notification)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar notificación",
			"Err": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Notificación creada",
		"Results": notification,
	})
}
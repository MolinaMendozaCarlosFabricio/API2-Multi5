package controllers

import (
	"net/http"
	"strconv"

	"api2-multi.com/a/src/Notifications/application"
	"api2-multi.com/a/src/Notifications/infrastructure"
	"github.com/gin-gonic/gin"
)

type DeleteNotificationC struct {
	uc application.DeleteNotificationUC
}

func NewDeleteNotificationC()*DeleteNotificationC{

	mysql := infrastructure.NewNotificationsMySQL()
	uc := application.NewDeleteNotificationUC(mysql)

	return&DeleteNotificationC{uc: *uc}
}

func(controller *DeleteNotificationC)Execute(c *gin.Context){
	id, error_param := c.Params.Get("id")
	if !error_param {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "No se pudo mapear el parámetro",
		})
		return
	}

	id_number, error_strconv := strconv.Atoi(id)
	if error_strconv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Parámetro incorrecto",
		})
		return
	}

	if err := controller.uc.Execute(id_number); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al obtener notificación",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Notificación eliminada",
	})
}
package routes

import (
	"api2-multi.com/a/src/Notifications/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func NotifiactionsRoutes(r *gin.Engine){
	notifications := r.Group("notifications")
	{
		notifications.POST("/", controllers.NewSendNotificationUC().Execute)
		notifications.GET("/of_an_user/:id_user", controllers.NewGetNotificationsOfAnUserC().Execute)
		notifications.GET("/:id", controllers.NewGetOneNotificationC().Execute)
		notifications.DELETE("/:id", controllers.NewDeleteNotificationC().Execute)
	}
}
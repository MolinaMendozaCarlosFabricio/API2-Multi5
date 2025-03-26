package domain

type IFMC_repo interface{
	SendNotification(message FCMMessage)error
}
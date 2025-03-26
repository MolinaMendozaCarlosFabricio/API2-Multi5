package infrastructure

var mysql *NotificationsMySQL
var validations *ValidationMySQL

func GoDependences(){
	mysql = NewNotificationsMySQL()
	validations = NewValidationMySQL()
}

func GetMySQL()*NotificationsMySQL{
	return mysql
}

func GetValidations()*ValidationMySQL{
	return validations
}
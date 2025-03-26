package services

import "api2-multi.com/a/src/Notifications/domain"

type ValidateMeasurementService struct {
	db domain.IValidation_repo
}

func NewValidateMeasurementService(db domain.IValidation_repo)*ValidateMeasurementService{
	return&ValidateMeasurementService{db: db}
}

func(s *ValidateMeasurementService)Execute(id_parcel int, hum, temp, air float32)(int, bool, error){
	params, err := s.db.GetParametersOfCultivation(id_parcel)
	if err != nil {
		return 0, false, err
	} else if params[0].Max_temp < temp {
		return 1, true, err
	} else if params[0].Min_temp > temp {
		return 2, true, err
	} else if params[0].Max_hum < hum {
		return 3, true, err
	} else if params[0].Min_hum > hum {
		return 4, true, err
	} else if params[0].Min_air > air {
		return 5, true, err
	} else {
		return 0, false, err
	}
}
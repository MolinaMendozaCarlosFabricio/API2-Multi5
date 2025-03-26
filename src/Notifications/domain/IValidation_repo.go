package domain

type IValidation_repo interface{
	GetParametersOfCultivation(id_parcel int)([]Parameters, error)
}
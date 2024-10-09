package handler

import (
	api "gomicrobase/api/proto"
	"gomicrobase/internal/model"
)

func convertToModelPerson(apiPerson *api.Person) *model.Person {
	return &model.Person{
		ID:    apiPerson.Id,
		Name:  apiPerson.Name,
		Email: apiPerson.Email,
	}
}

func convertToAPIPerson(modelPerson *model.Person) *api.Person {
	return &api.Person{
		Id:    modelPerson.ID,
		Name:  modelPerson.Name,
		Email: modelPerson.Email,
	}
}

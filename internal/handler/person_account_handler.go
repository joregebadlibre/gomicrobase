package handler

import (
	"context"
	api "gomicrobase/api/proto"
	"gomicrobase/internal/service"
)

type PersonAccountHandler struct {
	api.UnimplementedPersonAccountServiceServer
	service *service.PersonAccountService
}

func NewPersonAccountHandler(service service.PersonAccountService) *PersonAccountHandler {
	return &PersonAccountHandler{service: &service}
}

func (h *PersonAccountHandler) CreatePerson(ctx context.Context, req *api.Person) (*api.Person, error) {
	modelPerson := convertToModelPerson(req)
	createdPerson, err := h.service.CreatePerson(ctx, modelPerson)
	if err != nil {
		return nil, err
	}
	return convertToAPIPerson(createdPerson), nil
}

func (h *PersonAccountHandler) GetPerson(ctx context.Context, req *api.Person) (*api.Person, error) {
	person, err := h.service.GetPerson(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return convertToAPIPerson(person), nil
}

func (h *PersonAccountHandler) CreateAccount(ctx context.Context, req *api.Account) (*api.Account, error) {
	modelAccount := convertToModelAccount(req)
	createdAccount, err := h.service.CreateAccount(ctx, modelAccount)
	if err != nil {
		return nil, err
	}
	return convertToAPIAccount(createdAccount), nil
}

func (h *PersonAccountHandler) GetAccount(ctx context.Context, req *api.Account) (*api.Account, error) {
	account, err := h.service.GetAccount(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return convertToAPIAccount(account), nil
}

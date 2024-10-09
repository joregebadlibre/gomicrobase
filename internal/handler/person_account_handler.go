package handler

import (
	"context"
	"gomicrobase/api"
	"gomicrobase/internal/service"
)

type PersonAccountHandler struct {
	service service.PersonAccountService
}

func NewPersonAccountHandler(s service.PersonAccountService) *PersonAccountHandler {
	return &PersonAccountHandler{service: s}
}

func (h *PersonAccountHandler) CreatePerson(ctx context.Context, req *api.Person) (*api.Person, error) {
	return h.service.CreatePerson(ctx, req)
}

func (h *PersonAccountHandler) GetPerson(ctx context.Context, req *api.Person) (*api.Person, error) {
	return h.service.GetPerson(ctx, req)
}

// Similar handlers for Account

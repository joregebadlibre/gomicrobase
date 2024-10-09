package handler

import (
	api "gomicrobase/api/proto"
	"gomicrobase/internal/model"
)

func convertToModelAccount(apiAccount *api.Account) *model.Account {
	return &model.Account{
		ID:       apiAccount.Id,
		PersonID: apiAccount.PersonId,
		Balance:  apiAccount.Balance,
	}
}

func convertToAPIAccount(modelAccount *model.Account) *api.Account {
	return &api.Account{
		Id:       modelAccount.ID,
		PersonId: modelAccount.PersonID,
		Balance:  modelAccount.Balance,
	}
}

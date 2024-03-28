package repositories

import (
	"budgetlib/models"
	"budgetstoragelib/dto/requests"
	"budgetstoragelib/dto/responses"
	"budgetstoragelib/interfaces"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

type AccountRepository struct {
	server interfaces.IServer
}

func NewAccountRepository(server interfaces.IServer) *AccountRepository {
	return &AccountRepository{server}
}

func (r AccountRepository) ManageAccounts(
	input requests.AccountRequest) ([]*models.Account, error) {

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, errors.New("failed to marshal input for call to usp_ManageAccounts")
	}

	var output string

	db := r.server.Database()
	_, err = db.Exec(
		"[Org].[usp_ManageAccounts]",
		sql.Named("jsonDoc", string(payload)),
		sql.Named("outputJson", sql.Out{Dest: &output}),
	)

	if err != nil {
		return nil, err
	}

	var accounts responses.AccountReponse = responses.AccountReponse{}
	err = json.Unmarshal([]byte(output), &accounts)
	if err != nil {
		return nil, err
	}

	var results []*models.Account = []*models.Account{}
	for key, account := range accounts.Accounts {
		results = append(
			results, 
			models.NewAccount(
				account.Id, 
				account.Name, 
				models.ConvertAccountType(account.AccountType),
				),
			)

		fmt.Printf("results: %v\n", results[key])

	}

	return results, nil
}

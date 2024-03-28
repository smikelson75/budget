package responses

type AccountReponse struct {
	Accounts []AccountInternal `json:"accounts"`
}

type AccountInternal struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	AccountType string `json:"type"`
}
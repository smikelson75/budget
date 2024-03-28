package requests

type AccountRequest struct {
	User     string       `json:"user"`
	Accounts []AccountDto `json:"accounts"`
}

type AccountDto struct {
	Name        string        `json:"name"`
	AccountType string        `json:"type,omitempty"`
	Remove      bool          `json:"remove,omitempty"`
	AcctChange  AccountChange `json:"change,omitempty"`
}

type AccountChange struct {
	Name        string `json:"name,omitempty"`
	AccountType string `json:"type,omitempty"`
}

package model

type Credential struct {
	CashierId int
	Passcode  string `json:"passcode"`
}

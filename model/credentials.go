package model

type Credential struct {
	CashierId int
	Passcode  string `json:"passcode"`
}

type PasscodeData struct {
	Passcode string `json:"passcode"`
}

func (PasscodeData) TableName() string {
	return "cashiers"
}

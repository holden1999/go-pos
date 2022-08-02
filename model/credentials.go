package model

type Credential struct {
	CashierId int
	Passcode  string `json:"passcode"`
}

type PasscodeData struct {
	Passcode string `json:"passcode"`
}

type LoginResp struct {
	Token string `json:"token"`
}

func (PasscodeData) TableName() string {
	return "cashiers"
}

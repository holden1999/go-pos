package apprequest

type Cashier struct {
	Name     string `json:"name"`
	Passcode string `json:"passcode"`
}

type CashierLogin struct {
	Passcode string `json:"passcode"`
}

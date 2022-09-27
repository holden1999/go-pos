package apprequest

type Cashier struct {
	Name     string `json:"name,binding_required"`
	Passcode string `json:"passcode,binding_required"`
}

type CashierLogin struct {
	Passcode string `json:"passcode"`
}

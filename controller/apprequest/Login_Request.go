package apprequest

type LoginCredentials struct {
	Passcode string `json:"passcode" binding:"required"`
}

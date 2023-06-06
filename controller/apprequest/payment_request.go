package apprequest

type PaymentRequest struct {
	Name string `json:"name" validation:"required"`
	Type string `json:"type" validation:"required"`
	Logo string `json:"logo"`
}

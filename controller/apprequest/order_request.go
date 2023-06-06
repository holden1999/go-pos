package apprequest

type Order struct {
	PaymentId uint           `json:"paymentId" binding:"required"`
	TotalPaid int64          `json:"totalPaid" binding:"required"`
	Products  []ProductOrder `json:"products" binding:"required"`
}

type ProductOrder struct {
	ProductId uint `json:"productId"`
	Quantity  int  `json:"qty"`
}

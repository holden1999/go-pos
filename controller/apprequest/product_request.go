package apprequest

type ProductRequest struct {
	CategoryId uint     `json:"categoryId"`
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	Stock      int      `json:"stock"`
	Price      int      `json:"price"`
	Discount   Discount `json:"discount"`
}

package apprequest

type Discount struct {
	Qty       int    `json:"qty"`
	Type      string `json:"type"`
	Result    int    `json:"result"`
	ExpiredAt int    `json:"expiredAt"`
}

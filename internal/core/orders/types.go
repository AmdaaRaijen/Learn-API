package orders

type orderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type createOrderParams struct {
	Items []orderItem `json:"items"`
}

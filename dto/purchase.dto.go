package dto

type CreatePurchaseRequest struct {
	ProductId string `json:"productId" binding:"required"`
	Qtty      int `json:"qtty" binding:"required,min=1"`
}

type UpdatePurchaseRequest struct {
	ID        int64  `json:"id"`
	// ProductId string `json:"productId" binding:"required"`
	Receive   int `json:"receive" binding:"required,min=1"`
}

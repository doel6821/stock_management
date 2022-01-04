package dto

type CreateCartRequest struct {
	ProductID int64     `json:"productId" binding:"required"`
	Qtty      int64     `json:"qtty" binding:"required,min=1"`
	// UserID    int64     `json:"userId" binding:"required"`
	
}

type UpdateCartRequest struct {
	Id        int64     `json:"id"`
	ProductID int64     `json:"productId" binding:"required"`
	Qtty      int64     `json:"qtty" binding:"required,min=1"`
	UserID    int64     `json:"userId" binding:"required"`
	
}
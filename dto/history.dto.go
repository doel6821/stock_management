package dto

type History struct {
	ProductID int64     `json:"productId" binding:"required"`
	Qtty      int64     `json:"qtty" binding:"required,min=1"`
	UserID    int64     `json:"userId" binding:"required"`
	Status    string    `json:"status" binding:"required"`
}
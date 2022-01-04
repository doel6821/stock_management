package dto

type CreateProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Price uint64 `json:"price" binding:"required"`
	Detail string `json:"detail" binding:"required"`
}

type UpdateProductRequest struct {
	ID    int64  `json:"id" form:"id"`
	Name  string `json:"name" form:"name" binding:"required,min=1"`
	Price uint64 `json:"price" form:"email" binding:"required"`
	Detail string `json:"detail" form:"detail" binding:"required"`
}

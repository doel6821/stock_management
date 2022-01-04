package _product

import (
	"stock_management/entity"
	_user "stock_management/service/user"
)

type ProductResponse struct {
	ID          int64              `json:"id"`
	ProductName string             `json:"product_name"`
	Price       uint64             `json:"price"`
	Detail      string             `json:"detail"`
	Stock       int64              `json:"stock"`
	User        _user.UserResponse `json:"user,omitempty"`
}

type ProductCartResponse struct {
	ID          int64  `json:"id"`
	ProductName string `json:"product_name"`
	Price       uint64 `json:"price"`
	Detail      string `json:"detail"`
}

func NewProductResponse(product entity.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		ProductName: product.Name,
		Detail:      product.Detail,
		Price:       product.Price,
		Stock:       product.Stock,
		User:        _user.NewUserResponse(product.User),
	}
}

func NewProductCartResponse(product entity.Product) ProductCartResponse {
	return ProductCartResponse{
		ID:          product.ID,
		ProductName: product.Name,
		Detail:      product.Detail,
		Price:       product.Price,
	}
}

func NewProductArrayResponse(products []entity.Product) []ProductResponse {
	productRes := []ProductResponse{}
	for _, v := range products {
		p := ProductResponse{
			ID:          v.ID,
			ProductName: v.Name,
			Detail:      v.Detail,
			Price:       v.Price,
			Stock:       v.Stock,
			User:        _user.NewUserResponse(v.User),
		}
		productRes = append(productRes, p)
	}
	return productRes
}

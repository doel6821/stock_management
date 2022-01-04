package _purchase

import (
	"stock_management/entity"
	_product "stock_management/service/product"
	_user "stock_management/service/user"
)

type PurchaseResponse struct {
	ID        int64                        `json:"id"`
	Product   _product.ProductCartResponse `json:"product_name"`
	Qtty      int                          `json:"qtty"`
	Receive   int                          `json:"receive"`
	Status    string                       `json:"status"`
	User      _user.UserResponse           `json:"user,omitempty"`
	CreatedAt string                       `json:"createdAt"`
	UpdatedAt string                       `json:"updatedAt"`
}

func NewPurchaseResponse(purchase entity.Purchase) PurchaseResponse {
	return PurchaseResponse{
		ID:        purchase.ID,
		Product:   _product.NewProductCartResponse(purchase.Product),
		Qtty:      purchase.Qtty,
		Receive:   purchase.Receive,
		Status:    purchase.Status,
		User:      _user.NewUserResponse(purchase.User),
		CreatedAt: purchase.CreatedAt.Local().Format("2006-01-02"),
		UpdatedAt: purchase.UpdatedAt.Local().Format("2006-01-02"),
	}
}

func NewPurchaseArrayResponse(purchases []entity.Purchase) []PurchaseResponse {
	purchaseRes := []PurchaseResponse{}
	for _, v := range purchases {
		p := PurchaseResponse{
			ID:        v.ID,
			Product:   _product.NewProductCartResponse(v.Product),
			Qtty:      v.Qtty,
			Receive:   v.Receive,
			Status:    v.Status,
			User:      _user.NewUserResponse(v.User),
			CreatedAt: v.CreatedAt.Local().Format("2006-01-02"),
			UpdatedAt: v.UpdatedAt.Local().Format("2006-01-02"),
		}
		purchaseRes = append(purchaseRes, p)
	}
	return purchaseRes
}

package _cart

import (
	"stock_management/entity"
	_product "stock_management/service/product"
	_user "stock_management/service/user"
)

type CartResponse struct {
	ID        int64                        `json:"id"`
	Product   _product.ProductCartResponse `json:"product"`
	Qtty      uint64                       `json:"qtty"`
	Amount    int64                        `json:"amount"`
	User      _user.UserResponse           `json:"user,omitempty"`
	CreatedAt string                       `json:"createdAt"`
	UpdatedAt string                       `json:"updatedAt"`
}

func NewCartResponse(cart entity.Cart) CartResponse {
	return CartResponse{
		ID:        cart.ID,
		Product:   _product.NewProductCartResponse(cart.Product),
		Qtty:      cart.Qtty,
		Amount:    cart.Amount,
		User:      _user.NewUserResponse(cart.User),
		CreatedAt: cart.CreatedAt.Local().Format("2006-01-02"),
		UpdatedAt: cart.UpdatedAt.Local().Format("2006-01-02"),
	}
}

func NewCartArrayResponse(carts []entity.Cart) []CartResponse {
	cartRes := []CartResponse{}
	for _, v := range carts {
		p := CartResponse{
			ID:        v.ID,
			Product:   _product.NewProductCartResponse(v.Product),
			Qtty:      v.Qtty,
			Amount:    v.Amount,
			User:      _user.NewUserResponse(v.User),
			CreatedAt: v.CreatedAt.Local().Format("2006-01-02"),
			UpdatedAt: v.UpdatedAt.Local().Format("2006-01-02"),
		}
		cartRes = append(cartRes, p)
	}
	return cartRes
}

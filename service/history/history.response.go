package _history

import (
	"stock_management/entity"
	_product "stock_management/service/product"
	_user "stock_management/service/user"
)

type HistoryResponse struct {
	ID        int64                        `json:"id"`
	Product   _product.ProductCartResponse `json:"product_name"`
	Qtty      int                          `json:"qtty"`
	Amount    int64                        `json:"stock"`
	User      _user.UserResponse           `json:"user,omitempty"`
	Status    string                       `json:"status"`
	CreatedAt string                       `json:"createdAt"`
	UpdatedAt string                       `json:"updatedAt"`
}

func NewHistoryResponse(history entity.History) HistoryResponse {
	return HistoryResponse{
		ID:        history.ID,
		Product:   _product.NewProductCartResponse(history.Product),
		Qtty:      history.Qtty,
		Amount:    history.Amount,
		Status:    history.Status,
		CreatedAt: history.CreatedAt.Local().Format("2006-01-02"),
		UpdatedAt: history.UpdatedAt.Local().Format("2006-01-02"),
	}
}

func NewHistoryArrayResponse(Historys []entity.History) []HistoryResponse {
	HistoryRes := []HistoryResponse{}
	for _, v := range Historys {
		p := HistoryResponse{
			ID:        v.ID,
			Product:   _product.NewProductCartResponse(v.Product),
			Qtty:      v.Qtty,
			Amount:    v.Amount,
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Local().Format("2006-01-02"),
			UpdatedAt: v.UpdatedAt.Local().Format("2006-01-02"),
		}
		HistoryRes = append(HistoryRes, p)
	}
	return HistoryRes
}

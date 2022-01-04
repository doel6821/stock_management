package service

import (
	"errors"
	"stock_management/repo"
	_history "stock_management/service/history"
	"strconv"
)

type HistoryService interface {
	FindHistoryByProductID(productID, UserID string) (*[]_history.HistoryResponse, error)
	
}

type historyService struct {
	productRepo  repo.ProductRepository
	historyRepo  repo.HistoryRepository
}

func NewHistoryService(productRepo repo.ProductRepository, historyRepo repo.HistoryRepository) HistoryService {
	return &historyService{
		productRepo: productRepo,
		historyRepo: historyRepo,
	}
}

func (c *historyService) FindHistoryByProductID(productID, userId string) (*[]_history.HistoryResponse, error) {
	// find product
	product, err :=  c.productRepo.FindOneProductByID(productID)
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userId, 0, 64)
	if product.UserID != uid {
		return nil, errors.New("this product is not yours")
	}

	history, err := c.historyRepo.FindAllHistoryByProductId(productID)
	if err != nil {
		return nil, err
	}
	res := _history.NewHistoryArrayResponse(history)
	return &res, nil
}


package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"stock_management/dto"
	"stock_management/entity"
	"stock_management/repo"

	_purchase "stock_management/service/purchase"
)

type PurchaseService interface {
	All(userID string) (*[]_purchase.PurchaseResponse, error)
	CreatePurchase(purchaseRequest dto.CreatePurchaseRequest, userID string) (*_purchase.PurchaseResponse, error)
	UpdatePurchase(updatePurchaseRequest dto.UpdatePurchaseRequest, userID string) (*_purchase.PurchaseResponse, error)
	FindOnePurchaseByID(purchaseID string) (*_purchase.PurchaseResponse, error)
}

type purchaseService struct {
	purchaseRepo repo.PurchaseRepository
	productRepo  repo.ProductRepository
	historyRepo  repo.HistoryRepository
}

func NewPurchaseService(purchaseRepo repo.PurchaseRepository, productRepo repo.ProductRepository, historyRepo repo.HistoryRepository) PurchaseService {
	return &purchaseService{
		purchaseRepo: purchaseRepo,
		productRepo:  productRepo,
		historyRepo:  historyRepo,
	}
}

func (c *purchaseService) All(userID string) (*[]_purchase.PurchaseResponse, error) {
	purchases, err := c.purchaseRepo.All(userID)
	if err != nil {
		return nil, err
	}

	res := _purchase.NewPurchaseArrayResponse(purchases)
	return &res, nil
}

func (c *purchaseService) CreatePurchase(purchaseRequest dto.CreatePurchaseRequest, userID string) (*_purchase.PurchaseResponse, error) {
	id, _ := strconv.ParseInt(userID, 0, 64)
	productId, _ := strconv.ParseInt(purchaseRequest.ProductId, 0, 64)
	purchase := entity.Purchase{}
	purchase.ProductID = productId
	purchase.Qtty = purchaseRequest.Qtty
	purchase.UserID = id
	purchase.Status = "On Purchase"
	purchase.CreatedAt = time.Now()
	purchaseByte, _ := json.Marshal(purchase)
	log.Println("Data Purchase : ", string(purchaseByte))
	p, err := c.purchaseRepo.InsertPurchase(purchase)
	if err != nil {
		return nil, err
	}

	res := _purchase.NewPurchaseResponse(p)
	return &res, nil
}

func (c *purchaseService) FindOnePurchaseByID(purchaseID string) (*_purchase.PurchaseResponse, error) {
	purchase, err := c.purchaseRepo.FindOnePurchaseByID(purchaseID)

	if err != nil {
		return nil, err
	}

	res := _purchase.NewPurchaseResponse(purchase)
	return &res, nil
}

func (c *purchaseService) UpdatePurchase(updatePurchaseRequest dto.UpdatePurchaseRequest, userID string) (*_purchase.PurchaseResponse, error) {
	purchase, err := c.purchaseRepo.FindOnePurchaseByID(fmt.Sprintf("%d", updatePurchaseRequest.ID))
	if err != nil {
		return nil, err
	}

	// check purchase owner
	uid, _ := strconv.ParseInt(userID, 0, 64)
	if purchase.UserID != uid {
		return nil, errors.New("this is not your purchase")
	}

	// check status Purchase Order
	if purchase.Status == "Receive All" {
		return nil, errors.New("this purchase already completely received")
	}
	// update status receive
	if (purchase.Qtty - purchase.Receive) == updatePurchaseRequest.Receive {
		purchase.Status = "Receive All"
	} else if (purchase.Qtty - purchase.Receive) > updatePurchaseRequest.Receive {
		purchase.Status = "Partial Receive"
	} else if (purchase.Qtty - purchase.Receive) < updatePurchaseRequest.Receive {
		return nil, errors.New("invalid input receive qtty")
	}

	purchase.UpdatedAt = time.Now()
	purchase.Receive = updatePurchaseRequest.Receive
	purchase, err = c.purchaseRepo.UpdatePurchase(purchase)
	if err != nil {
		return nil, err
	}

	// find product
	product, err := c.productRepo.FindOneProductByID(strconv.Itoa(int(purchase.ProductID)))
	if err != nil {
		return nil, err
	}
	product.Stock += int64(purchase.Receive)
	// update stock
	c.productRepo.UpdateProduct(product)

	// add to history
	history := entity.History{
		ID:        0,
		ProductID: purchase.ProductID,
		Qtty:      purchase.Receive,
		Amount:    int64(purchase.Receive) * int64(product.Price),
		UserID:    uid,
		Status:    "Inbound",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.historyRepo.InsertHistory(history)

	res := _purchase.NewPurchaseResponse(purchase)
	return &res, nil
}

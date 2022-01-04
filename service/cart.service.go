package service

import (
	"errors"
	"strconv"
	"time"

	"stock_management/dto"
	"stock_management/entity"
	"stock_management/repo"

	_cart "stock_management/service/cart"
)

type CartService interface {
	All(userID string) (*[]_cart.CartResponse, error)
	CreateCart(cartRequest dto.CreateCartRequest, userID string) (*_cart.CartResponse, error)
	FindOneCartByID(CartID string) (*_cart.CartResponse, error)
}

type cartService struct {
	cartRepo    repo.CartRepository
	productRepo repo.ProductRepository
	historyRepo repo.HistoryRepository
}

func NewCartService(cartRepo repo.CartRepository, productRepo repo.ProductRepository, historyRepo repo.HistoryRepository) CartService {
	return &cartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
		historyRepo: historyRepo,
	}
}

func (c *cartService) All(userID string) (*[]_cart.CartResponse, error) {
	carts, err := c.cartRepo.All(userID)
	if err != nil {
		return nil, err
	}

	res := _cart.NewCartArrayResponse(carts)
	return &res, nil
}

func (c *cartService) CreateCart(cartRequest dto.CreateCartRequest, userID string) (*_cart.CartResponse, error) {
	id, _ := strconv.ParseInt(userID, 0, 64)
	productId, _ := strconv.ParseInt(userID, 0, 64)
	cart := entity.Cart{}
	cart.ProductID = productId
	cart.Qtty = uint64(cartRequest.Qtty)

	// find product
	product, err := c.productRepo.FindOneProductByID(strconv.Itoa(int(cart.ProductID)))
	if err != nil {
		return nil, err
	}
	// check stock availability
	if product.Stock < int64(cart.Qtty) {
		return nil, errors.New("stock unavailable")
	}

	// add to history
	history := entity.History{
		ID:        0,
		ProductID: cart.ProductID,
		Qtty:      int(cart.Qtty),
		Amount:    int64(cart.Qtty) * int64(product.Price),
		UserID:    id,
		Status:    "Outbound",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.historyRepo.InsertHistory(history)

	product.Stock -= int64(cart.Qtty)
	// update stock
	c.productRepo.UpdateProduct(product)

	cart.UserID = id
	cart.Amount = int64(product.Price) * int64(cart.Qtty)
	cart.CreatedAt = time.Now()
	p, err := c.cartRepo.InsertCart(cart)
	if err != nil {
		return nil, err
	}

	res := _cart.NewCartResponse(p)
	return &res, nil
}

func (c *cartService) FindOneCartByID(cartID string) (*_cart.CartResponse, error) {
	cart, err := c.cartRepo.FindOneCartByID(cartID)

	if err != nil {
		return nil, err
	}

	res := _cart.NewCartResponse(cart)
	return &res, nil
}

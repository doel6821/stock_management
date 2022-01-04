package repo

import (
	"stock_management/entity"

	"github.com/jinzhu/gorm"
)

type CartRepository interface {
	All(userID string) ([]entity.Cart, error)
	InsertCart(product entity.Cart) (entity.Cart, error)
	UpdateCart(product entity.Cart) (entity.Cart, error)
	DeleteCart(productID string) error
	FindOneCartByID(ID string) (entity.Cart, error)
	FindAllCart(userID string) ([]entity.Cart, error)
}

type cartRepo struct {
	connection *gorm.DB
}

func NewCartRepo(connection *gorm.DB) CartRepository {
	return &cartRepo{
		connection: connection,
	}
}

func (c *cartRepo) All(userID string) ([]entity.Cart, error) {
	carts := []entity.Cart{}
	c.connection.Preload("User").Preload("Product").Where("user_id = ?", userID).Find(&carts)
	return carts, nil
}

func (c *cartRepo) InsertCart(cart entity.Cart) (entity.Cart, error) {
	c.connection.Save(&cart)
	c.connection.Preload("User").Preload("Product").Find(&cart)
	return cart, nil
}

func (c *cartRepo) UpdateCart(cart entity.Cart) (entity.Cart, error) {
	c.connection.Save(&cart)
	c.connection.Preload("User").Preload("Product").Find(&cart)
	return cart, nil
}

func (c *cartRepo) FindOneCartByID(cartId string) (entity.Cart, error) {
	var cart entity.Cart
	res := c.connection.Preload("User").Preload("Product").Where("id = ?", cartId).Take(&cart)
	if res.Error != nil {
		return cart, res.Error
	}
	return cart, nil
}

func (c *cartRepo) FindAllCart(userID string) ([]entity.Cart, error) {
	carts := []entity.Cart{}
	c.connection.Preload("User").Preload("Product").Where("user_id = ?", userID).Find(&carts)
	return carts, nil
}

func (c *cartRepo) DeleteCart(cartId string) error {
	var cart entity.Cart
	res := c.connection.Preload("User").Where("id = ?", cartId).Take(&cart)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Preload("User").Preload("Product").Delete(&cart)
	return nil
}

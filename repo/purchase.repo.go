package repo

import (
	"stock_management/entity"

	"github.com/jinzhu/gorm"
)

type PurchaseRepository interface {
	All(userID string) ([]entity.Purchase, error)
	InsertPurchase(product entity.Purchase) (entity.Purchase, error)
	UpdatePurchase(product entity.Purchase) (entity.Purchase, error)
	FindOnePurchaseByID(ID string) (entity.Purchase, error)
	FindAllPurchaseByUserId(userID string) ([]entity.Purchase, error)
	FindAllPurchaseByProductId(productID string) ([]entity.Purchase, error)
}

type purchaseRepo struct {
	connection *gorm.DB
}

func NewPurchaseRepo(connection *gorm.DB) PurchaseRepository {
	return &purchaseRepo{
		connection: connection,
	}
}

func (c *purchaseRepo) All(userID string) ([]entity.Purchase, error) {
	purchases := []entity.Purchase{}
	c.connection.Preload("User").Preload("Product").Where("user_id = ?", userID).Find(&purchases)
	return purchases, nil
}

func (c *purchaseRepo) InsertPurchase(purchase entity.Purchase) (entity.Purchase, error) {
	c.connection.Save(&purchase)
	c.connection.Preload("User").Preload("Product").Find(&purchase)
	return purchase, nil
}

func (c *purchaseRepo) UpdatePurchase(purchase entity.Purchase) (entity.Purchase, error) {
	c.connection.Save(&purchase)
	c.connection.Preload("User").Preload("Product").Find(&purchase)
	return purchase, nil
}

func (c *purchaseRepo) FindOnePurchaseByID(purchaseId string) (entity.Purchase, error) {
	var purchase entity.Purchase
	res := c.connection.Preload("User").Preload("Product").Where("id = ?", purchaseId).Take(&purchase)
	if res.Error != nil {
		return purchase, res.Error
	}
	return purchase, nil
}

func (c *purchaseRepo) FindAllPurchaseByUserId(userID string) ([]entity.Purchase, error) {
	purchases := []entity.Purchase{}
	c.connection.Preload("User").Preload("Product").Where("user_id = ?", userID).Find(&purchases)
	return purchases, nil
}

func (c *purchaseRepo) FindAllPurchaseByProductId(productID string) ([]entity.Purchase, error) {
	purchases := []entity.Purchase{}
	c.connection.Preload("User").Preload("Product").Where("product_id = ?", productID).Find(&purchases)
	return purchases, nil
}


package repo

import (
	"stock_management/entity"

	"github.com/jinzhu/gorm"
)

type HistoryRepository interface {
	InsertHistory(product entity.History) (entity.History, error)
	FindOneHistoryByID(ID string) (entity.History, error)
	FindAllHistoryByUserId(userID string) ([]entity.History, error)
	FindAllHistoryByProductId(productID string) ([]entity.History, error)
}

type historyRepo struct {
	connection *gorm.DB
}

func NewHistoryRepo(connection *gorm.DB) HistoryRepository {
	return &historyRepo{
		connection: connection,
	}
}

func (c *historyRepo) InsertHistory(history entity.History) (entity.History, error) {
	c.connection.Save(&history)
	c.connection.Preload("User").Preload("Product").Find(&history)
	return history, nil
}

func (c *historyRepo) FindOneHistoryByID(cartId string) (entity.History, error) {
	var cart entity.History
	res := c.connection.Preload("User").Preload("Product").Where("id = ?", cartId).Take(&cart)
	if res.Error != nil {
		return cart, res.Error
	}
	return cart, nil
}

func (c *historyRepo) FindAllHistoryByUserId(userID string) ([]entity.History, error) {
	histories := []entity.History{}
	c.connection.Preload("User").Preload("Product").Where("user_id = ?", userID).Find(&histories)
	return histories, nil
}

func (c *historyRepo) FindAllHistoryByProductId(productID string) ([]entity.History, error) {
	histories := []entity.History{}
	c.connection.Preload("User").Preload("Product").Where("product_id = ?", productID).Find(&histories)
	return histories, nil
}


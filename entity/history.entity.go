package entity

import "time"

type History struct {
	ID        int64     `gorm:"column:id" json:"id"`
	ProductID int64     `gorm:"not null" json:"productId"`
	Product   Product   `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product"`
	Qtty      int     `gorm:"type:int column:qtty" json:"qtty"`
	Amount    int64     `gorm:"type:int column:amount" json:"amount"`
	UserID    int64     `gorm:"not null" json:"userId"`
	Status    string    `gorm:"type:varchar(50) column:status" json:"status"`
	User      User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
package entity

import "time"

type Product struct {
	ID        int64     `gorm:"column:id" json:"id"`
	Name      string    `gorm:"type:varchar(100) column:name" validate:"required"  json:"name"`
	Detail    string    `gorm:"type:varchar(100) column:detail" validate:"required"  json:"detail"`
	Price     uint64    `gorm:"type:int column:price" validate:"required"  json:"price"`
	Stock     int64     `gorm:"type:int column:stock" json:"stock"`
	UserID    int64     `gorm:"not null" json:"userId"`
	User      User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

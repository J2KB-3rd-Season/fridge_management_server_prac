package model

import "time"

type Inventory struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	/*InventoryType InventoryType `gorm:"foreignKey:type_id;" json:"inventory_type"`*/
	UserId      uint64    ``
	TypeId      uint64    ``
	Title       string    `gorm:"size:255;not null;unique" json:"title"`
	Description string    `gorm:"type:text;not null;" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

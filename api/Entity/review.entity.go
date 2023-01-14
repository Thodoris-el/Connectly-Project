package entity

import "time"

type Review struct {
	ID          string    `gorm:"primary_key;auto_increment" json:"id"`
	Customer_id int64     `gorm:"not null;" json:"customer_id"`
	Text        string    `gorm:"size:2550;not null;" json:"text"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

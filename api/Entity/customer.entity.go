package entity

import "time"

type Customer struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id"`
	First_name  string    `gorm:"size:255;not null;" json:"first_name"`
	Last_name   string    `gorm:"size:255;not null;" json:"last_name"`
	Facebook_id int64     `gorm:"not null;" json:"facebook_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

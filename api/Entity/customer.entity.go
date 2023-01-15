package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID           int64  `gorm:"primary_key;auto_increment" json:"id"`
	First_name   string `gorm:"size:255;not null;" json:"first_name"`
	Last_name    string `gorm:"size:255;not null;" json:"last_name"`
	Facebook_id  string `gorm:"not null;" json:"facebook_id"`
	Sent_Message bool   `gorm:"default:false" json:"sent_message"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//Save Customer Function
func (customer *Customer) SaveCustomer(db *gorm.DB) (*Customer, error) {

	err := db.Debug().Create(&customer).Error
	if err != nil {
		log.Println("error while saving user")
		return &Customer{}, err
	}

	return customer, nil
}

//find all customers
func (customer *Customer) FindAllCustomers(db *gorm.DB) (*[]Customer, error) {

	customers := []Customer{}
	err := db.Debug().Model(&Customer{}).Limit(10).Find(&customers).Error

	if err != nil {
		log.Println("Error while finding customers")
		return &[]Customer{}, err
	}

	return &customers, err
}

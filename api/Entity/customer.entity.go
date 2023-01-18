package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

//Structure of the Customer Entity
type Customer struct {
	ID           int64  `gorm:"primary_key;auto_increment" json:"id"`
	First_name   string `gorm:"size:255;not null;" json:"first_name"`
	Last_name    string `gorm:"size:255;not null;" json:"last_name"`
	Facebook_id  string `gorm:"not null;" json:"facebook_id"`
	Sent_Message bool   `gorm:"default:false" json:"sent_message"`
	Language     string `gorm:"default:eng" json:"language"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//Save Customer Function to DB
func (customer *Customer) SaveCustomer(db *gorm.DB) (*Customer, error) {

	err := db.Debug().Create(&customer).Error
	if err != nil {
		//log.Println("error while saving customer")
		return &Customer{}, err
	}

	return customer, nil
}

//find all customers from DB
func (customer *Customer) FindAllCustomers(db *gorm.DB) (*[]Customer, error) {

	customers := []Customer{}
	err := db.Debug().Model(&Customer{}).Limit(10).Find(&customers).Error

	if err != nil {
		//log.Println("Error while finding customers")
		return &[]Customer{}, err
	}

	return &customers, err
}

//find customer by Id
func (customer *Customer) FindCustomerByID(db *gorm.DB, C_id int64) (*Customer, error) {

	err := db.Debug().Model(Customer{}).Where("id = ?", C_id).Take(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Customer{}, errors.New("Customer Not Found")
		}
		return &Customer{}, err
	}

	return customer, err
}

//Find customer by facebook id
func (customer *Customer) FindByFacebookId(db *gorm.DB, F_id string) (*Customer, error) {

	err := db.Debug().Model(&Customer{}).Where("facebook_id = ?", F_id).Take(&customer).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Customer{}, err
		} else {
			return &Customer{}, err
		}
	}
	return customer, err
}

//update customer
func (customer *Customer) UpdateCustomer(db *gorm.DB, C_id int64) (*Customer, error) {

	db = db.Debug().Model(&Customer{}).Where("id = ?", C_id).Take(&Customer{}).UpdateColumns(
		map[string]interface{}{
			"first_name":   customer.First_name,
			"last_name":    customer.Last_name,
			"facebook_id":  customer.Facebook_id,
			"language":     customer.Language,
			"sent_message": customer.Sent_Message,
			"updated_at":   time.Now(),
		},
	)
	err := db.Debug().Model(&Customer{}).Where("id = ?", C_id).Take(&customer).Error
	if err != nil {
		return &Customer{}, err
	}
	return customer, nil
}

//delete customer
func (customer *Customer) DeleteCustomer(db *gorm.DB, C_id int64) (int64, error) {
	db = db.Debug().Model(&Customer{}).Where("id = ?", C_id).Take(&Customer{}).Delete(&Customer{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

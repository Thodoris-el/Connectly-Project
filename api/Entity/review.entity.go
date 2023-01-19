package entity

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID          int64  `gorm:"primary_key;auto_increment" json:"id"`
	Customer_id string `gorm:"not null;" json:"customer_id"`
	Text        string `gorm:"size:2550;not null;" json:"text"`
	Score       int    `gorm:"default: -1" json:"score"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//Save Review
func (review *Review) SaveReview(db *gorm.DB) (*Review, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := db.WithContext(ctx).Debug().Create(&review).Error
	if err != nil {
		log.Println("error while saving review", err)
		return &Review{}, err
	}

	return review, nil
}

//Find All Reviews
func (review *Review) FindAllReviews(db *gorm.DB) (*[]Review, error) {

	reviews := []Review{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := db.WithContext(ctx).Debug().Model(&Review{}).Limit(100).Find(&reviews).Error

	if err != nil {
		log.Println("Error while finding reviews")
		return &[]Review{}, err
	}

	return &reviews, err
}

//Find By Id
func (review *Review) FindById(db *gorm.DB, R_id int64) (*Review, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := db.WithContext(ctx).Debug().Model(&Review{}).Where("id = ?", R_id).Take(&review).Error

	if err != nil {
		log.Println("error while geting review by id")
		return &Review{}, err
	}

	return review, nil
}

//Find All Reviews from a specific customer
func (review *Review) FindByCustomerId(db *gorm.DB, C_id string) (*[]Review, error) {

	reviews := []Review{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := db.WithContext(ctx).Debug().Model(&Review{}).Where("customer_id = ?", C_id).Limit(100).Find(&reviews).Error

	if err != nil {
		log.Println("Error while finding reviews from a specific customer")
		return &[]Review{}, err
	}

	return &reviews, err
}

//update review
func (review *Review) UpdateReview(db *gorm.DB, C_id int64) (*Review, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	db = db.WithContext(ctx).Debug().Model(&Review{}).Where("id = ?", C_id).Take(&Review{}).UpdateColumns(
		map[string]interface{}{
			"customer_id": review.Customer_id,
			"text":        review.Text,
			"score":       review.Score,
			"updated_at":  time.Now(),
		},
	)
	err := db.WithContext(ctx).Debug().Model(&Review{}).Where("id = ?", C_id).Take(&review).Error
	if err != nil {
		return &Review{}, err
	}
	return review, nil
}

//delete review
func (review *Review) DeleteReview(db *gorm.DB, C_id int64) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	db = db.WithContext(ctx).Debug().Model(&Review{}).Where("id = ?", C_id).Take(&Review{}).Delete(&Review{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

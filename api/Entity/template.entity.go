package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Template struct {
	ID           int64  `gorm:"primary_key;auto_increment" json:"id"`
	Placeholder  string `json:"placeholder"`
	Title        string `json:"title"`
	Language     string `json:"language"`
	Subtitle     string `json:"subtitle"`
	Button_Title string `json:"button_title"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//Save Template
func (template *Template) SaveTemplate(db *gorm.DB) (*Template, error) {

	err := db.Debug().Create(&template).Error
	if err != nil {
		log.Println("error while saving Template", err)
		return &Template{}, err
	}

	return template, nil
}

//Find All Templates
func (template *Template) FindAllTemplates(db *gorm.DB) (*[]Template, error) {

	templates := []Template{}
	err := db.Debug().Model(&Template{}).Limit(100).Find(&templates).Error

	if err != nil {
		log.Println("Error while finding Templates")
		return &[]Template{}, err
	}

	return &templates, err
}

//Find By Id
func (template *Template) FindById(db *gorm.DB, T_id int64) (*Template, error) {

	err := db.Debug().Model(&Template{}).Where("id = ?", T_id).Take(&template).Error

	if err != nil {
		log.Println("error while geting Template by id")
		return &Template{}, err
	}

	return template, nil
}

//Find Template from a specific language
func (template *Template) FindByLanguage(db *gorm.DB, lang string) (*Template, error) {

	err := db.Debug().Model(&Template{}).Where("language = ?", lang).Limit(100).Take(&template).Error

	if err != nil {
		log.Println("Error while finding reviews from a specific customer")
		return &Template{}, err
	}

	return template, err
}

//update template
func (template *Template) UpdateTemplate(db *gorm.DB, T_id int64) (*Template, error) {

	db = db.Debug().Model(&Template{}).Where("id = ?", T_id).Take(&Template{}).UpdateColumns(
		map[string]interface{}{
			"placeholder":  template.Placeholder,
			"title":        template.Title,
			"language":     template.Language,
			"subtitle":     template.Subtitle,
			"button_title": template.Button_Title,
			"updated_at":   time.Now(),
		},
	)
	err := db.Debug().Model(&Review{}).Where("id = ?", T_id).Take(&template).Error
	if err != nil {
		return &Template{}, err
	}
	return template, nil
}

//delete template
func (template *Template) DeleteTemplate(db *gorm.DB, T_id int64) (int64, error) {
	db = db.Debug().Model(&Template{}).Where("id = ?", T_id).Take(&Template{}).Delete(&Template{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

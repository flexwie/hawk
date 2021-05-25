package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"hawk.wie.gg/models"
)

func AddCategory(name string) error {
	record := models.Category{Id: uuid.NewString(), Name: name}

	result := DB.Create(&record)

	return result.Error
}

func RemoveCategory(id string) {
	DB.Delete(&models.Category{}, id)
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category

	result := DB.Preload(clause.Associations).Find(&categories)

	return categories, result.Error
}

func GetCategoryByName(name string) (models.Category, error) {
	var category models.Category

	result := DB.Preload(clause.Associations).Where("name = ?", name).First(&category)

	return category, result.Error
}

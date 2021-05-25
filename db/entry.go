package db

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"hawk.wie.gg/models"
)

func AddEntry(entry models.Entry, parentName string) error {
	DB.Model(&models.Entry{}).Where("end = ?", "").Update("end", time.Now().UTC().Format("2006-01-02T15:04:05-0700"))

	var parent models.Category
	res := DB.Model(&parent).Where("name = ?", parentName).First(&parent)

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		AddCategory(parentName)
		_ = DB.Model(&parent).Where("name = ?", parentName).First(&parent)
	}

	record := models.Entry{Id: uuid.NewString(), Name: entry.Name, Start: entry.Start, End: entry.End, CategoryId: parent.Id}

	result := DB.Create(&record)

	return result.Error
}

func RemoveEntry(id string) {
	DB.Delete(&models.Entry{}, id)
}

func GetAllEntries() ([]models.Entry, error) {
	var entries []models.Entry

	result := DB.Find(&entries)

	return entries, result.Error
}

func UpdateEntry(key string, value string, id string) {
	DB.Exec("UPDATE entries SET ? = ? WHERE id = ?", strings.ToLower(key), value, id)
}

func StopEntries(end string) error {
	if end != "" {
		res := DB.Model(&models.Entry{}).Where("end = \"\"").Update("end", end)
		return res.Error
	} else {
		res := DB.Model(&models.Entry{}).Where("end = \"\"").Update("end", time.Now().UTC().Format(time.RFC3339))
		return res.Error
	}
}

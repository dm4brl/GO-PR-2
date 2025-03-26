package database

import (
	"fmt"

	"github.com/dm4brl/GO-PR-2/internal/models"
	"gorm.io/gorm"
)

// UpdateSwitchStatus обновляет статус свитча в базе данных
func UpdateSwitchStatus(status models.SwitchStatus) error {
	db := GetDB() // получаем соединение с базой данных

	// Обновляем или создаем запись в таблице switch_statuses
	result := db.Where("id = ?", status.ID).First(&status)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Запись не найдена, создаем новую
			if err := db.Create(&status).Error; err != nil {
				return fmt.Errorf("ошибка при создании записи: %v", err)
			}
		} else {
			return fmt.Errorf("ошибка при поиске записи: %v", result.Error)
		}
	} else {
		// Обновляем существующую запись
		if err := db.Save(&status).Error; err != nil {
			return fmt.Errorf("ошибка при обновлении записи: %v", err)
		}
	}

	return nil
}

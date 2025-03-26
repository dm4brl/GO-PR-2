// internal/database/update.go
package database

import (
	"log"
	"time"

	"github.com/dm4brl/GO-PR-2/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UpdateSwitchStatus выполняет upsert для свитча: если запись с данным ID существует, она обновляется, иначе создается новая.
func UpdateSwitchStatus(db *gorm.DB, status models.SwitchStatus) error {
	// Устанавливаем время обновления
	status.UpdatedAt = time.Now()

	// Используем конструкцию upsert с OnConflict.
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // по полю id
		DoUpdates: clause.AssignmentColumns([]string{"state", "timestamp", "updated_at"}),
	}).Create(&status)

	if result.Error != nil {
		log.Printf("Ошибка обновления статуса свитча с ID %s: %v", status.ID, result.Error)
		return result.Error
	}

	log.Printf("Статус свитча с ID %s успешно обновлен", status.ID)
	return nil
}

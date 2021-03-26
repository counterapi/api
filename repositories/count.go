package repositories

import (
	"github.com/counterapi/counter/models"

	"gorm.io/gorm"
)

type CountRepository struct {
	DB *gorm.DB
}

//ListByCounterName list counts by models.Counter name
func (r CountRepository) ListByCounterName(name string) ([]models.Count, error) {
	var counts []models.Count

	err := r.DB.
		Joins("JOIN counters on counters.id=counts.counter_id").
		Where("counters.name = ?", name).
		Find(&counts).Error
	if err != nil {
		return counts, err
	}

	return counts, nil
}

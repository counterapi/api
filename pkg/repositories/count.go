package repositories

import (
	"fmt"

	"github.com/counterapi/counterapi/pkg/models"

	"gorm.io/gorm"
)

const recordLimit = 1000

// CountRepository is a repository for models.Count.
type CountRepository struct {
	DB *gorm.DB
}

// GroupByCounterNameAndTimeInterval groups the counts by models.Counter name and time interval.
func (r CountRepository) GroupByCounterNameAndTimeInterval(
	name string,
	interval string,
	order string,
) ([]models.CountGroupResult, error) {
	var results []models.CountGroupResult

	err := r.DB.
		Model(&models.Count{}).
		Select(fmt.Sprintf("count(*) as count, date_trunc('%s', counts.created_at) as date", interval)).
		Joins("JOIN counters on counters.id=counts.counter_id").
		Where("counters.name = ?", name).
		Group(fmt.Sprintf("date_trunc('%s', counts.created_at)", interval)).
		Order(fmt.Sprintf("date %s", order)).
		Limit(recordLimit).
		Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}

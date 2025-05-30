package repositories

import (
	"errors"
	"fmt"

	"github.com/counterapi/api/pkg/models"

	"gorm.io/gorm"
)

const countRecordLimit = 1000

// CounterRepository is a repository for models.Counter.
type CounterRepository struct {
	DB *gorm.DB
}

// GetByName get counter by name.
func (r CounterRepository) GetByName(namespace, name string) (models.Counter, error) {
	counter := models.Counter{
		Name: name,
		Namespace: models.Namespace{
			Name: namespace,
		},
	}

	if err := r.DB.
		Joins("JOIN namespaces on counters.namespace_id = namespaces.id").
		Where("namespaces.name = ?", namespace).
		Preload("Namespace").First(&counter, "counters.name = ?", name).Error; err != nil {
		return counter, err
	}

	return counter, nil
}

// CountCounters counts counters.
func (r CounterRepository) CountCounters() (int64, error) {
	var count int64

	if err := r.DB.Model(&models.Counter{}).Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

// CountCounts counts counts.
func (r CounterRepository) CountCounts() (int64, error) {
	var count int64

	if err := r.DB.Model(&models.Count{}).Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

// CountNamespaces counts namespaces.
func (r CounterRepository) CountNamespaces() (int64, error) {
	var count int64

	if err := r.DB.Model(&models.Namespace{}).Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

// Create creates counter.
func (r CounterRepository) Create(counter *models.Counter) error {
	if err := r.DB.Create(&counter).Error; err != nil {
		return err
	}

	return nil
}

// GetOrCreateByName get counter or create by name.
func (r CounterRepository) GetOrCreateByName(namespace, name string) (models.Counter, error) {
	//nolint:staticcheck // It's okay to use literal here.
	namespaceRepository := NamespaceRepository{r.DB}

	counter, err := r.GetByName(namespace, name)
	if err == nil {
		return counter, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		counter.Namespace, err = namespaceRepository.GetOrCreateByName(namespace)
		if err != nil {
			return counter, err
		}

		if err = r.Create(&counter); err != nil {
			return counter, err
		}

		return counter, nil
	}

	return counter, err
}

// IncreaseByName increase models.Counter by name.
func (r CounterRepository) IncreaseByName(namespace, name string) (models.Counter, error) {
	// Get counter if exist
	counter, err := r.GetOrCreateByName(namespace, name)
	if err != nil {
		return counter, err
	}

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		// Increment Counter
		if err = tx.Model(&counter).Update("count", counter.Count+1).Error; err != nil {
			return err
		}

		// Creating Count is so expensive, so we do it only if the count is not zero.
		// Create Count record
		//count := models.Count{
		//	CounterID: counter.ID,
		//}

		//if err = tx.Create(&count).Error; err != nil {
		//	return err
		//}

		return nil
	})
	if err != nil {
		return counter, err
	}

	return counter, nil
}

// DecreaseByName decrease models.Counter by name.
func (r CounterRepository) DecreaseByName(namespace, name string) (models.Counter, error) {
	// Get counter if exist
	counter, err := r.GetOrCreateByName(namespace, name)
	if err != nil {
		return counter, err
	}

	// Do nothing if counter is zero
	if counter.Count == 0 {
		return counter, nil
	}

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		// Increment Counter
		if err = tx.Model(&counter).Update("count", counter.Count-1).Error; err != nil {
			return err
		}

		// Creating Count is so expensive, so we do it only if the count is not zero.
		// Create Count record
		//count := models.Count{
		//	CounterID: counter.ID,
		//}
		//
		//if err = tx.Create(&count).Error; err != nil {
		//	return err
		//}

		return nil
	})
	if err != nil {
		return counter, err
	}

	return counter, nil
}

// SetByName sets models.Counter by name.
func (r CounterRepository) SetByName(namespace, name string, count uint) (models.Counter, error) {
	// Get counter if exist
	counter, err := r.GetOrCreateByName(namespace, name)
	if err != nil {
		return counter, err
	}

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		// Increment Counter
		if err = tx.Model(&counter).Update("count", count).Error; err != nil {
			return err
		}

		// Creating Count is so expensive, so we do it only if the count is not zero.
		// Create Count record
		//count := models.Count{
		//	CounterID: counter.ID,
		//}
		//
		//if err = tx.Create(&count).Error; err != nil {
		//	return err
		//}

		return nil
	})
	if err != nil {
		return counter, err
	}

	return counter, nil
}

// GroupByCounterNameAndTimeInterval returns stats of given models.Counter and models.Namespace.
func (r CounterRepository) GroupByCounterNameAndTimeInterval(
	namespace string,
	name string,
	interval string,
	order string,
) ([]models.CountGroupResult, error) {
	var results []models.CountGroupResult

	err := r.DB.
		Model(&models.Count{}).
		Select(fmt.Sprintf("count(*) as count, date_trunc('%s', counts.created_at) as date", interval)).
		InnerJoins("JOIN counters on counters.id=counts.counter_id").
		InnerJoins("JOIN namespaces on counters.namespace_id = namespaces.id").
		Where("counters.name = ?", name).
		Where("namespaces.name = ?", namespace).
		Group(fmt.Sprintf("date_trunc('%s', counts.created_at)", interval)).
		Order(fmt.Sprintf("date %s", order)).
		Limit(countRecordLimit).
		Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}

package repositories

import (
	"errors"
	"github.com/counterapi/counterapi/pkg/models"
	"gorm.io/gorm"
)

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
		Preload("Namespace").
		First(&counter, "counters.name = ?", name).Error; err != nil {

		return counter, err
	}

	return counter, nil
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
	counter, err := r.GetByName(namespace, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			namespaceRepository := NamespaceRepository{r.DB}
			ns, err := namespaceRepository.GetOrCreateByName(namespace)
			if err != nil {
				return counter, err
			}

			counter.Namespace = ns

			err = r.Create(&counter)
			if err != nil {
				return counter, err
			}

			return counter, nil
		}

		return counter, err
	}

	return counter, nil
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

		// Create Count record
		count := models.Count{
			CounterID: counter.ID,
		}

		if err = tx.Create(&count).Error; err != nil {
			return err
		}

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

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		// Increment Counter
		if err = tx.Model(&counter).Update("count", counter.Count-1).Error; err != nil {
			return err
		}

		// Create Count record
		count := models.Count{
			CounterID: counter.ID,
		}

		if err = tx.Create(&count).Error; err != nil {
			return err
		}

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

		// Create Count record
		count := models.Count{
			CounterID: counter.ID,
		}

		if err = tx.Create(&count).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return counter, err
	}

	return counter, nil
}

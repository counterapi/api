package repositories

import (
	"errors"
	"github.com/counterapi/counterapi/pkg/models"
	"gorm.io/gorm"
)

// NamespaceRepository is a repository for models.Namespace.
type NamespaceRepository struct {
	DB *gorm.DB
}

// GetByName get counter by name.
func (r NamespaceRepository) GetByName(name string) (models.Namespace, error) {
	namespace := models.Namespace{
		Name: name,
	}

	if err := r.DB.
		First(&namespace, "namespaces.name = ?", name).Error; err != nil {

		return namespace, err
	}

	return namespace, nil
}

// Create creates counter.
func (r NamespaceRepository) Create(namespace *models.Namespace) error {
	if err := r.DB.Create(&namespace).Error; err != nil {
		return err
	}

	return nil
}

// GetOrCreateByName get counter or create by name.
func (r NamespaceRepository) GetOrCreateByName(name string) (models.Namespace, error) {
	counter, err := r.GetByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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

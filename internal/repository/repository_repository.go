package repository

import (
	"gitlab-tool/internal/models"

	"gorm.io/gorm"
)

type RepositoryRepository struct {
	db *gorm.DB
}

func NewRepositoryRepository(db *gorm.DB) *RepositoryRepository {
	return &RepositoryRepository{db: db}
}

func (r *RepositoryRepository) Create(repo *models.Repository) error {
	return r.db.Create(repo).Error
}

func (r *RepositoryRepository) FindByID(id uint) (*models.Repository, error) {
	var repo models.Repository
	err := r.db.Preload("Owner").First(&repo, id).Error
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (r *RepositoryRepository) FindByOwnerID(ownerID uint) ([]models.Repository, error) {
	var repos []models.Repository
	err := r.db.Where("owner_id = ?", ownerID).Find(&repos).Error
	if err != nil {
		return nil, err
	}
	return repos, nil
}

func (r *RepositoryRepository) FindByUsernameAndName(username, name string) (*models.Repository, error) {
	var repo models.Repository
	err := r.db.Joins("JOIN users ON repositories.owner_id = users.id").
		Where("users.username = ? AND repositories.name = ?", username, name).
		Preload("Owner").
		First(&repo).Error
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (r *RepositoryRepository) Update(repo *models.Repository) error {
	return r.db.Save(repo).Error
}

func (r *RepositoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Repository{}, id).Error
}

func (r *RepositoryRepository) ListPublic() ([]models.Repository, error) {
	var repos []models.Repository
	err := r.db.Where("visibility = ?", "public").Preload("Owner").Find(&repos).Error
	if err != nil {
		return nil, err
	}
	return repos, nil
}

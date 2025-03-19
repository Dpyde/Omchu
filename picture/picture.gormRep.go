package picture

import (

	// PictureRep "github.com/Dpyde/Omchu/internal/"

	"github.com/Dpyde/Omchu/internal/entity"
	"gorm.io/gorm"
)

type GormPictureRepository struct {
	db *gorm.DB
}

func NewGormPictureRepository(db *gorm.DB) PictureRepository {
	return &GormPictureRepository{db: db}
}

func (r *GormPictureRepository) SavePicturesToDB(pictures []entity.Picture) error {
	return r.db.Create(&pictures).Error
}
func (r *GormPictureRepository) GetPictureFromDB(id uint) ([]entity.Picture, error) {
	var pictures []entity.Picture
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&user).Association("Pictures").Find(&pictures); err != nil {
		return nil, err
	}
	return pictures, nil
}

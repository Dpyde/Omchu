package picture

import (
	"github.com/Dpyde/Omchu/internal/entity"
)

type PictureRepository interface {
	SavePicturesToDB(pictures []entity.Picture) error
	GetPictureFromDB(id uint) ([]entity.Picture, error)
}

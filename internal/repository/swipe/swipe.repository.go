package swipeRep

import (
	"github.com/Dpyde/Omchu/internal/entity"
)

type SwipeRepository interface {
	Pud(user entity.Swipe) error
}
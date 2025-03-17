package swipe

import (
	"github.com/Dpyde/Omchu/internal/entity"
)

type SwipeRepository interface {
	Pud(user *entity.Swipe,is_match *bool) error
}
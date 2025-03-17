package swipeSer

import (
	"errors"

	"github.com/Dpyde/Omchu/internal/entity"
	"github.com/Dpyde/Omchu/internal/repository/swipe"
)

//primary port
type SwipeService interface {
	SwipeCheck(swipe *entity.Swipe, is_match *bool) error
}

type swipeServiceImpl struct {
	repo swipeRep.SwipeRepository
}

func NewSwipeService(repo swipeRep.SwipeRepository) SwipeService {
	return &swipeServiceImpl{repo: repo}
}

func (s *swipeServiceImpl) SwipeCheck(swipe *entity.Swipe ,is_match *bool) error {
	if(swipe.SwipedID == swipe.SwiperID) {
		return errors.New("You can't swipe yourself")
	}
	// Business logic...
	if err := s.repo.Pud(swipe,is_match); err != nil {
		return err
	}
	return nil
}

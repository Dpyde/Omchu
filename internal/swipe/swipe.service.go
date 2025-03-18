package swipe

import (
	"errors"

	"github.com/Dpyde/Omchu/internal/entity"

)

//primary port
type SwipeService interface {
	SwipeCheck(swipe *entity.Swipe, is_match *bool) error
}

type swipeServiceImpl struct {
	repo SwipeRepository
}

func NewSwipeService(repo SwipeRepository) SwipeService {
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

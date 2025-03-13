package swipeGormRep

import (
	"github.com/Dpyde/Omchu/internal/entity"
	"github.com/Dpyde/Omchu/internal/repository/swipe"
	"gorm.io/gorm"
	"fmt"
	"errors"
)

type GormSwipeRepository struct {
	db *gorm.DB
}

func NewGormSwipeRepository(db *gorm.DB) swipeRep.SwipeRepository {
	return &GormSwipeRepository{db: db}
}

func (r *GormSwipeRepository) Pud(swipe entity.Swipe) error {
	if result := r.db.Create(&swipe); result.Error != nil {
		// Handle database errors
		return result.Error
	}
	matchedSwipe := entity.Swipe{}
	is_swipe_back := r.db.Where("swiped_id = ? AND swiper_id = ? AND liked = ?" , swipe.SwiperID, swipe.SwipedID, true).First(&matchedSwipe)
	if is_swipe_back.Error != nil {
		if errors.Is(is_swipe_back.Error, gorm.ErrRecordNotFound) {
			fmt.Println("No matching record found")
		} else {
			fmt.Println("Database error:", is_swipe_back.Error)
			return is_swipe_back.Error;
		}
	} else {
		fmt.Println("Matched swipe found")
	}

	return nil
}
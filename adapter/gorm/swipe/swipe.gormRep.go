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

func (r *GormSwipeRepository) Pud(swipe *entity.Swipe,is_match *bool) error {
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
		chat := entity.Chat{}
		err := r.db.Create(&chat).Error
		if err != nil {
    		return err
		}
		err = r.db.Model(&chat).Association("Users").Append(
			&entity.User{Model: gorm.Model{ID: swipe.SwiperID}},
			&entity.User{Model: gorm.Model{ID: swipe.SwipedID}},
		)
		if err != nil {
			return err
		}
		*is_match = true
		fmt.Println("Chat created");
	}

	return nil
}
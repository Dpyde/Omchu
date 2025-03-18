package swipe

import (
	"errors"
	"fmt"
	"github.com/Dpyde/Omchu/internal/entity"
	"gorm.io/gorm"
)

type GormSwipeRepository struct {
	db *gorm.DB
}

func NewGormSwipeRepository(db *gorm.DB) SwipeRepository {
	return &GormSwipeRepository{db: db}
}

func (r *GormSwipeRepository) Pud(swipe *entity.Swipe,is_match *bool) (uint,error) {
	if result := r.db.Create(&swipe); result.Error != nil {
		// Handle database errors
		return 0,result.Error
	}
	matchedSwipe := entity.Swipe{}
	is_swipe_back := r.db.Where("swiped_id = ? AND swiper_id = ? AND liked = ?" , swipe.SwiperID, swipe.SwipedID, true).First(&matchedSwipe)
	if is_swipe_back.Error != nil {
		if errors.Is(is_swipe_back.Error, gorm.ErrRecordNotFound) {
			*is_match = false;
			return 0, nil
		} else {
			fmt.Println("Database error:", is_swipe_back.Error)
			return 0,is_swipe_back.Error;
		}
	}
	fmt.Println("Matched swipe found")
	chat := entity.Chat{}
	err := r.db.Create(&chat).Error
	if err != nil {
		return 0,err
	}
	err = r.db.Model(&chat).Association("Users").Append(
		&entity.User{Model: gorm.Model{ID: swipe.SwiperID}},
		&entity.User{Model: gorm.Model{ID: swipe.SwipedID}},
	)
	if err != nil {
		return 0,err
	}
	*is_match = true

	fmt.Println("Chat created");
		

		
	return chat.ID,nil
}
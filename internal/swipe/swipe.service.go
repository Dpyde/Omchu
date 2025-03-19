package swipe

import (
	"errors"
	"bytes"
	"io"
	"fmt"
	"net/http"
	"github.com/Dpyde/Omchu/internal/entity"
)

// primary port
type SwipeService interface {
	SwipeCheck(swipe *entity.Swipe, is_match *bool) error
}

type swipeServiceImpl struct {
	repo SwipeRepository
}

func NewSwipeService(repo SwipeRepository) SwipeService {
	return &swipeServiceImpl{repo: repo}
}

func (s *swipeServiceImpl) SwipeCheck(swipe *entity.Swipe, is_match *bool) error {
	if swipe.SwipedID == swipe.SwiperID {
		return errors.New("You can't swipe yourself")
	}
	// Business logic...
	chatID,err := s.repo.Pud(swipe, is_match)
	if  err != nil {
		return err
	}
	if(*is_match){
		url := "http://127.0.0.1:8080/ws/createRoom"
		data := fmt.Sprintf(`{ 
					"id":"%d",
					"name":"Chat %d"
				}`, chatID, chatID)
		reqBody := bytes.NewBufferString(data)

		resp, err := http.Post(url, "application/json", reqBody)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return err
		}

		fmt.Println(string(body))
	}
	return nil
}

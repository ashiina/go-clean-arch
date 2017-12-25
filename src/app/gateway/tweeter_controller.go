package gateway

import "fmt"
import (
	"app/entity"
	"app/usecase"
)

type TweeterController struct {
	TweetUsecase usecase.TweetUsecase
}

func NewTweeterController() *TweeterController {
	return &TweeterController{
		TweetUsecase: usecase.TweetUsecase{
			TweetRepository: &TweetRepository{},
		},
	}
}

func (tc *TweeterController) Index() {
	_ = tc.TweetUsecase.Post(entity.Tweet{Message: "Hello Tweeter!"})
	_ = tc.TweetUsecase.Post(entity.Tweet{Message: "Hello Tweeter2!"})
	_ = tc.TweetUsecase.Post(entity.Tweet{Message: "Hello Tweeter3!"})

	tweets, err := tc.TweetUsecase.GetAll()
	if err != nil {
		fmt.Println("Could not get Tweets.")
		return
	}
	fmt.Println(tweets)
}

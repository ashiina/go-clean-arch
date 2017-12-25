package usecase

import "app/entity"

type TweetUsecase struct {
	TweetRepository TweetRepository
}

func (tu *TweetUsecase) Post(t entity.Tweet) error {
	_, err := tu.TweetRepository.Save(t)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TweetUsecase) GetAll() (entity.Tweets, error) {
	tweets, err := tu.TweetRepository.GetAll()
	return tweets, err
}

package usecase

import "app/entity"

type TweetRepository interface {
	Save(entity.Tweet) (int, error)
	GetAll() (entity.Tweets, error)
}

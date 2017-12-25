package gateway

import "math/rand"
import "app/entity"

type TweetRepository struct {
	tweets entity.Tweets
}

func NewTweetRepository() (*TweetRepository, error) {
	return &TweetRepository{}, nil
}

func (tr *TweetRepository) Save(t entity.Tweet) (int, error) {
	id := rand.Int()
	t.ID = id
	tr.tweets = append(tr.tweets, t)
	return id, nil
}

func (tr *TweetRepository) GetAll() (entity.Tweets, error) {
	return tr.tweets, nil
}

package usecase

import (
	"url-shortener/pkg/shortener"
)

type DB interface {
	Save(long, short string) error
	Get(short string) (string, error)
}

type Cache interface {
	Set(key, value string)
	Get(key string) (string, error)
}

type URLUsecase struct {
	db    DB
	cache Cache
}

func NewURLUsecase(db DB, cache Cache) *URLUsecase {
	return &URLUsecase{db: db, cache: cache}
}

func (u *URLUsecase) Create(longURL string) (string, error) {
	code := shortener.GenerateCode(6)

	err := u.db.Save(longURL, code)
	if err != nil {
		return "", err
	}

	u.cache.Set(code, longURL)

	return code, nil
}

func (u *URLUsecase) Get(code string) (string, error) {
	val, err := u.cache.Get(code)
	if err == nil {
		return val, nil
	}

	val, err = u.db.Get(code)
	if err != nil {
		return "", err
	}

	u.cache.Set(code, val)
	return val, nil
}

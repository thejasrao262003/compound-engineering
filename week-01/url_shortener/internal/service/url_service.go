package service

import (
	"fmt"
	"url_shortener/internal/repository/redis_repo"
	"url_shortener/internal/repository/sql"
	"url_shortener/internal/utils"
)

type URLService struct {
	repo *sql.URLRepository
	cache *redis_repo.Cache
}

func NewURLService(repo *sql.URLRepository, cache *redis_repo.Cache) *URLService {
	return &URLService{
		repo: repo,
		cache: cache,
	}
}

func (s *URLService) CreateShortURL(longURL string) (string, error){
	id, err := s.repo.CreateURL(longURL)

	if err != nil {
		return "", err
	}

	code := utils.Encode(id)

	err = s.repo.UpdateShortCode(id, code)

	if err != nil{
		return "", err
	}

	s.cache.Set("short:"+code, longURL)

	return code, nil
}

func (s *URLService) GetLongURL(code string) (string, error) {
	key := "short:" + code

	val, err := s.cache.Get(key)
	if err == nil {
		fmt.Println("Cache HIT")
		return val, nil
	}

	fmt.Println("Cache Miss")

	longURL, err := s.repo.GetByShortCode(code)
	if err != nil {
		return "", err
	}

	s.cache.Set(key, code)

	return longURL, nil
}
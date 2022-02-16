package service
import (
	"sec-kill/cache"
	"sec-kill/store"
)

type Service interface {
	Users() UserService

}

type service struct {
	store store.Factory
	cache cache.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory,cache cache.Factory) Service {
	return &service{
		store: store,
		cache: cache,
	}
}

func (s *service) Users() UserService {
	return newUser(s)
}

package member

import "errors"

// Service is the domain service
type Service interface {
	FindAllMember() (*Member, error)
	FindMember(id string) (*Member, error)
}

type service struct {
	repo Repository
}

// NewService creates new domain service
func NewService(repo Repository) (Service, error) {
	s := service{repo}
	return &s, nil
}

func (s *service) FindAllMember() (*Member, error) {
	return &Member{
		ID: 123,
	}, nil
}

func (s *service) FindMember(id string) (*Member, error) {
	if id == "1" {
		return nil, errors.New("this is error message")
	}
	return &Member{
		ID: 123,
	}, nil
}

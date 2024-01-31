package user

type Service interface {
	GetUsers() ([]User, error)
}

type service struct {
	repo UserRepository
}

func (s *service) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

func NewService(repo UserRepository) Service {
	return &service{repo: repo}
}

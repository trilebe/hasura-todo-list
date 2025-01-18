package user

type Service struct {
	repository *repository
}

func (s *Service) Register(payload RegisterUserRequest) (*RegisterUserResponse, error) {
	user, err := s.repository.findByUsername(payload.UserName)
	if err == nil && user != nil {
		return nil, &Errors.DuplicatedUsername
	}

	user, err = NewUser(payload.UserName, payload.PassWord)

	if err != nil {
		return nil, err
	}

	user, err = s.repository.save(user)

	return &RegisterUserResponse{
		UserName: user.UserName,
	}, err
}

func (s *Service) FindByUsername(username string) (*User, error) {
	return s.repository.findByUsername(username)
}

func NewService() *Service {
	repository := NewRepository()
	service := Service{repository}

	return &service
}

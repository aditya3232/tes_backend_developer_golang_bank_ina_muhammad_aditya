package users

import "golang.org/x/crypto/bcrypt"

type Service interface {
	GetAll() ([]User, error)
	GetOne(input UserGetOneByIdInput) (User, error)
	Create(input UserCreateInput) (User, error)
	Update(input UserUpdateInput) (User, error)
	Delete(input UserGetOneByIdInput) error
}

type service struct {
	userRepository Repository
}

func NewService(userRepository Repository) *service {
	return &service{userRepository}
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.userRepository.GetAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) GetOne(input UserGetOneByIdInput) (User, error) {
	user, err := s.userRepository.GetOne(input.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) Create(input UserCreateInput) (User, error) {
	if input.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return User{}, err
		}

		input.Password = string(password)
	}

	user := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	newUser, err := s.userRepository.Create(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Update(input UserUpdateInput) (User, error) {
	_, err := s.userRepository.GetOne(input.ID)
	if err != nil {
		return User{}, err
	}

	if input.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return User{}, err
		}

		input.Password = string(password)
	}

	user := User{
		ID:       input.ID,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	newUser, err := s.userRepository.Update(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Delete(input UserGetOneByIdInput) error {
	_, err := s.userRepository.GetOne(input.ID)
	if err != nil {
		return err
	}

	err = s.userRepository.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}

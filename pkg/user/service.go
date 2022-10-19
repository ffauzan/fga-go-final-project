package user

import (
	"errors"
	"final-project/pkg/domain"
)

type CryptoService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(plaintext string, hashed string) error
}

type ValidatorService interface {
	ValidateUser(user *domain.User) error
	ValidateLoginRequest(req *domain.LoginRequest) error
	ValidateRegisterRequest(req *domain.RegisterRequest) error
}

type AuthService interface {
	GenerateToken(user *domain.User) (*string, error)
	IsTokenValid(token string) (bool, error)
	GetUserIDFromToken(token string) (uint, error)
	IsUserCanCreate(userID uint, entity *interface{}) bool
	IsUserCanAccess(userID uint, entity *interface{}) bool
	IsUserCanUpdate(userID uint, entity *interface{}) bool
	IsUserCanDelete(userID uint, entity *interface{}) bool
}

type service struct {
	repo          domain.UserRepository
	cryptoService CryptoService
	authService   AuthService
	validator     ValidatorService
}

func NewService(
	repo domain.UserRepository,
	cryptoService CryptoService,
	authService AuthService,
	validatorService ValidatorService,
) domain.UserService {
	return &service{
		repo:          repo,
		cryptoService: cryptoService,
		authService:   authService,
		validator:     validatorService,
	}
}

func (s *service) DeleteUser(userID uint) error {
	// check if user exist
	if !s.IsUserExist(userID) {
		return errors.New("user not found")
	}
	return s.repo.DeleteUserByID(userID)
}

func (s *service) UpdateUser(userID uint, user *domain.UpdateUserRequest) (*domain.User, error) {
	// get user from db by id
	userFromDB, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// update user
	userFromDB.Username = user.Username
	userFromDB.Email = user.Email

	return s.repo.UpdateUser(userFromDB)
}

func (s *service) IsUserExist(id uint) bool {
	_, err := s.repo.GetUserByID(id)
	return err == nil
}

func (s *service) Register(req *domain.RegisterRequest) (*domain.User, error) {
	// check if username & email already exist
	if s.repo.IsUsernameExist(req.Username) {
		return nil, errors.New("username already exist")
	}
	if s.repo.IsEmailExist(req.Email) {
		return nil, errors.New("email already exist")
	}

	// hash password
	hashedPassword, err := s.cryptoService.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// save user
	userToSave := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Age:      req.Age,
		Password: hashedPassword,
	}
	return s.repo.SaveUser(userToSave)
}

func (s *service) Login(user *domain.LoginRequest) (*string, error) {
	// validate login request
	if err := s.validator.ValidateLoginRequest(user); err != nil {
		return nil, err
	}

	// get user by username
	userFromDB, err := s.repo.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	// verify password
	err = s.cryptoService.VerifyPassword(user.Password, userFromDB.Password)
	if err != nil {
		return nil, err
	}

	// generate token
	token, err := s.authService.GenerateToken(userFromDB)
	if err != nil {
		return nil, err
	}

	return token, nil
}

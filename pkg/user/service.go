package user

import (
	"errors"
)

type Repository interface {
	SaveUser(user *User) (*User, error)
	GetUserByID(userID uint) (*User, error)
	GetUserByUsername(username string) (*User, error)
	DeleteUserByID(userID uint) error
	UpdateUser(userID *User) (*User, error)
	IsUsernameExist(username string) bool
	IsEmailExist(email string) bool
}

type CryptoService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(plaintext string, hashed string) error
}

type ValidatorService interface {
	ValidateUser(user *User) error
	ValidateLoginRequest(req *LoginRequest) error
	ValidateRegisterRequest(req *RegisterRequest) error
}

type AuthService interface {
	GenerateToken(user *User) (*string, error)
	IsTokenValid(token string) (bool, error)
	GetUserIDFromToken(token string) (uint, error)
	IsUserCanCreate(userID uint, entity *interface{}) bool
	IsUserCanAccess(userID uint, entity *interface{}) bool
	IsUserCanUpdate(userID uint, entity *interface{}) bool
	IsUserCanDelete(userID uint, entity *interface{}) bool
}

type Service interface {
	DeleteUser(userID uint) error
	UpdateUser(userID uint, req *UpdateUserRequest) (*User, error)
	IsUserExist(userID uint) bool
	Register(req *RegisterRequest) (*User, error)
	Login(req *LoginRequest) (*string, error)
}

type service struct {
	repo          Repository
	cryptoService CryptoService
	authService   AuthService
	validator     ValidatorService
}

func NewService(repo Repository, cryptoService CryptoService, authService AuthService, validatorService ValidatorService) Service {
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

func (s *service) UpdateUser(userID uint, user *UpdateUserRequest) (*User, error) {
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

func (s *service) Register(req *RegisterRequest) (*User, error) {
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
	userToSave := &User{
		Username: req.Username,
		Email:    req.Email,
		Age:      req.Age,
		Password: hashedPassword,
	}
	return s.repo.SaveUser(userToSave)
}

func (s *service) Login(user *LoginRequest) (*string, error) {
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

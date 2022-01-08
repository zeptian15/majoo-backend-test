package services

import (
	"errors"
	"majoo-backend-test/app/models"
	"majoo-backend-test/app/repositories"
	"majoo-backend-test/app/utils"
)

// Create New Auth Service Interface
type AuthService interface {
	RegisterUser(models.RegisterRequest) error
	LoginUser(models.LoginRequest) (string, error)
	GetUserProfile(userName string) (models.UserResponse, error)
}

// Binding Repository to Service ( Constructor )
type authService struct {
	repository repositories.UserRepository
}

// Create New Function, to Create New Auth Service Instance
func NewAuthService(repository repositories.UserRepository) *authService {
	return &authService{repository}
}

///
/// Implementing User Service Interface
///

// Register User
func (service *authService) RegisterUser(registerRequest models.RegisterRequest) error {
	// Access Repository to Check if UserName Exist
	isExist, err := service.repository.CheckIfUserExistsByUsername(registerRequest.UserName)

	// Check if there is error when check user
	if isExist {
		// If Query Error
		if err != nil {
			return err
		}

		return errors.New("user: user_name already exists")
	}

	// Convert Register Request to User Model
	var user models.User
	user.Name = registerRequest.Name
	user.UserName = registerRequest.UserName

	// Hash Password
	hashedPassword, err := utils.GeneratePassword(registerRequest.Password)

	// Check if there is error when hashing password
	if err != nil {
		return err
	}

	// Assign Hashed Password
	user.Password = hashedPassword

	// Access Repository to Create User
	errorCreate := service.repository.CreateUser(user)

	// Return Error
	return errorCreate
}

// Login User
func (service *authService) LoginUser(loginRequest models.LoginRequest) (string, error) {
	// Access Repository to Check if User Exist
	isExist, err := service.repository.CheckIfUserExistsByUsername(loginRequest.UserName)

	// If User Does not exist
	if !isExist {
		// If Query Error
		if err != nil {
			return "", err
		}

		return "", errors.New("user: user does not exists")

	}

	// Create new User Model
	var user models.UserResponse

	// Access Repository to Get User Detail
	user, errorDetailUser := service.repository.GetUserByUsername(loginRequest.UserName)

	// Check if there is error when get user detail
	if errorDetailUser != nil {
		return "", errorDetailUser
	}

	// Compare Login Request Password, with User Password
	isMatch := utils.ComparePassword(user.Password, loginRequest.Password)

	// Check if password match
	if !isMatch {
		return "", errors.New("password: password not match")
	}

	// Create JWT
	token, errorJWT := utils.GenerateToken(user)

	// Check if there is eror when create JWT
	if errorJWT != nil {
		return "", errorJWT
	}

	// Return JWT ( Token )
	return token, nil
}

// Get User Profile
func (service *authService) GetUserProfile(userName string) (models.UserResponse, error) {
	// Access Repository to Get User Detail
	user, err := service.repository.GetUserByUsername(userName)

	// Return Value
	return user, err
}

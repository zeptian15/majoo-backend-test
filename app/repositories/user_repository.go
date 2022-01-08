package repositories

import (
	"database/sql"
	"majoo-backend-test/app/models"
)

// Create New User Repository Interface
type UserRepository interface {
	CreateUser(models.User) error
	GetUserByUsername(userName string) (models.UserResponse, error)
	CheckIfUserExistsByUsername(userName string) (bool, error)
}

// Binding Database to Repository ( Constructor )
type userRepository struct {
	db *sql.DB
}

// Create Function, to Create New User Repository Instance
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

///
/// Implementing User Repository Interface
///

// Create User
func (repository *userRepository) CreateUser(user models.User) error {
	// Prepare Query
	query := `INSERT INTO users ( name, user_name, password ) VALUES ( ?, ?, ?)`

	// Execute Query
	_, err := repository.db.Exec(query, user.Name, user.UserName, user.Password)

	// Return err
	return err
}

// Get User By Username
func (repository *userRepository) GetUserByUsername(userName string) (models.UserResponse, error) {
	// Prepare Query
	query := `SELECT * FROM users WHERE user_name = ? AND deleted_at IS NULL`

	// Prepare New Struct Instance
	var user models.UserResponse

	// Execute Query
	row := repository.db.QueryRow(query, userName)

	// Scan Row to User Response Model
	if err := row.Scan(
		&user.Id,
		&user.Name,
		&user.UserName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	); err != nil {
		return models.UserResponse{}, err
	}

	// Return User Response Model
	return user, nil
}

// Check if User Exists by Username
func (repository *userRepository) CheckIfUserExistsByUsername(userName string) (bool, error) {
	// Prepare Query
	query := `SELECT user_name FROM users WHERE user_name = ? AND deleted_at IS NULL`

	// Execute Query
	err := repository.db.QueryRow(query, userName).Scan(&userName)

	// Check if there is error
	if err != nil {
		// If Error is not NoRows
		if err != sql.ErrNoRows {
			// Return Query Error
			return false, err
		}
		// If error is NoRows = User does not exists
		return false, nil
	}

	// Return User Exist
	return true, nil
}

package user

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
	"github.com/zulfikarrosadi/chad/internal/entity"
)

type UserRepository interface {
	Get(ctx context.Context, userID int) (*entity.User, error)
	Create(ctx context.Context, user *UserCreateRequest) (*UserCreateResponse, error)
	Update(ctx context.Context, user *UserUpdateRequest) (*UserUpdateResponse, error)
	Delete(ctx context.Context, userID int) error
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepositoryImpl{
		DB: db,
	}
}

func (ur UserRepositoryImpl) Get(ctx context.Context, userID int) (*entity.User, error) {
	r := ur.DB.QueryRowContext(ctx, "SELECT id, email, username, password, profile_picture FROM users WHERE id = $1", userID)
	user := &entity.User{}

	err := r.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.ProfilePicture)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur UserRepositoryImpl) Create(ctx context.Context, user *UserCreateRequest) (*UserCreateResponse, error) {
	_, err := ur.DB.ExecContext(ctx, "INSERT INTO users (email, username, password, profile_picture) VALUES ($1, $2, $3, $4)", user.Email, user.Username, user.Password, user.ProfilePicture)
	if err, ok := err.(*pq.Error); ok {
		log.Fatal(err.Code, " ", err.Constraint, " ", err.Column, " ", err.Detail)
		return nil, errors.New("unique constraint error")
	}

	r := ur.DB.QueryRowContext(ctx, "SELECT id FROM users WHERE email = $1", user.Email)
	userResponse := &UserCreateResponse{}
	r.Scan(&userResponse.ID)
	userResponse.Email = user.Email
	userResponse.ProfilePicture = user.ProfilePicture
	userResponse.Username = user.Username

	return userResponse, nil
}

func (ur UserRepositoryImpl) Update(ctx context.Context, user *UserUpdateRequest) (*UserUpdateResponse, error) {
	SQL := "UPDATE users SET email = $1, username = $2, password = $3, profile_picture = $4 WHERE id = $5"
	_, err := ur.DB.ExecContext(ctx, SQL, user.Email, user.Username, user.Password, user.ProfilePicture, user.ID)
	if err != nil {
		return nil, err
	}
	r := ur.DB.QueryRowContext(ctx, "SELECT id FROM users WHERE id = $1", user.ID)
	userResponse := &UserUpdateResponse{}
	r.Scan(&userResponse.ID)
	userResponse.Email = user.Email
	userResponse.Username = user.Email
	userResponse.ProfilePicture = user.ProfilePicture

	return userResponse, nil
}

func (ur UserRepositoryImpl) Delete(ctx context.Context, userID int) error {
	_, err := ur.DB.ExecContext(ctx, "DELETE FROM users WHERE id = $i", userID)
	if err != nil {
		return err
	}
	return nil
}

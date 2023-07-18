package user

import "github.com/jackc/pgx/v5/pgxpool"

type Repository interface {
	Get()
	Create()
	Update()
	Delete()
}

type UserRepository struct {
	DB *pgxpool.Pool
}

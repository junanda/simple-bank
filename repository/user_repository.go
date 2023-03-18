package repository

import (
	"context"

	"github.com/junanda/simple-bank/db"
	"github.com/junanda/simple-bank/entity"
	interfaceRepo "github.com/junanda/simple-bank/repository/interface_repo"
)

func NewUserRepository(db db.DBTX) interfaceRepo.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

type UserRepositoryImpl struct {
	db db.DBTX
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, arg entity.CreateUserParams) (entity.User, error) {
	var (
		data entity.User
		err  error
	)

	queryCreate := `INSERT INTO users (
		username, hashed_password, full_name, email
	) VALUES (
		$1, $2, $3, $4
	) RETURNING username, hashed_password, full_name, email, password_changed_at, created_at, is_email_verified`

	row := u.db.QueryRowContext(ctx, queryCreate, arg.Username, arg.HashedPassword, arg.FullName, arg.Email)
	err = row.Scan(
		&data.Username,
		&data.HashedPassword,
		&data.FullName,
		&data.Email,
		&data.PasswordChangedAt,
		&data.CreatedAt,
		&data.IsEmailVerified,
	)

	return data, err
}

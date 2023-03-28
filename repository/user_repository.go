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

func (u *UserRepositoryImpl) GetUser(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	query := `SELECT username, hashed_password, full_name, email, password_changed_at, created_at, is_email_verified FROM users WHERE username = $1 LIMIT 1`
	row := u.db.QueryRowContext(ctx, query, username)

	err := row.Scan(
		&user.Username,
		&user.HashedPassword,
		&user.FullName,
		&user.Email,
		&user.PasswordChangedAt,
		&user.CreatedAt,
		&user.IsEmailVerified,
	)

	return user, err
}

func (u *UserRepositoryImpl) UpdateUser(ctx context.Context, arg entity.UpdateUser) (entity.User, error) {
	var user entity.User

	query := `UPDATE users 
	SET 
		hashed_password = COALESCE($1, hashed_password),
		password_changed_at = COALESCE($2, password_changed_at),
		full_name = COALESCE($3, full_name),
		email = COALESCE($4, email),
		is_email_verified = COALESCE($5, is_email_verified)
	WHERE
		username = $6
	RETURNING username, hashed_password, full_name, email, password_changed_at, created_at, is_email_verified`

	row := u.db.QueryRowContext(ctx, query, arg.HashedPassword, arg.PasswordChangedAt, arg.FullName, arg.Email, arg.IsEmailVerified, arg.Username)

	err := row.Scan(
		&user.Username,
		&user.HashedPassword,
		&user.FullName,
		&user.Email,
		&user.PasswordChangedAt,
		&user.CreatedAt,
		&user.IsEmailVerified,
	)
	return user, err
}

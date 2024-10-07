package repo

import (
	"context"
	"database/sql"

	"github.com/LeMinh0706/music-go/db"
	"github.com/LeMinh0706/music-go/util"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepo() (*UserRepository, error) {
	pg, err := util.InitPostgres()
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		queries: db.New(pg),
	}, nil
}

func (repo *UserRepository) CreateUser(ctx context.Context, username, password, fullname string, email sql.NullString, gender, role_id int32) (db.User, error) {
	return repo.queries.CreateUser(ctx, db.CreateUserParams{
		Username: username,
		Password: password,
		Fullname: fullname,
		Email:    email,
		Gender:   gender,
		RoleID:   role_id,
		Avt:      util.RandomAvatar(gender),
	})
}

func (repo *UserRepository) GetUser(ctx context.Context, username string) (db.User, error) {
	return repo.queries.GetUser(ctx, username)
}

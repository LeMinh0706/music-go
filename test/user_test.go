package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/LeMinh0706/music-go/db"
	"github.com/LeMinh0706/music-go/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) {
	password, err := util.HashPashword("kocanpass")
	require.NoError(t, err)
	gender := util.RandomGender()

	arg := db.CreateUserParams{
		Username: util.RandomString(6),
		Password: password,
		Fullname: util.RandomString(6),
		Gender:   gender,
		Email:    sql.NullString{String: util.RandomEmail(), Valid: true},
		RoleID:   3,
		Avt:      util.RandomAvatar(gender),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Fullname, user.Fullname)
	require.Equal(t, arg.RoleID, user.RoleID)
	require.Equal(t, arg.Avt, user.Avt)
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

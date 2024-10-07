package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/LeMinh0706/music-go/db"
	"github.com/LeMinh0706/music-go/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) db.User {
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

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestLogin(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.Login(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Fullname, user2.Fullname)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Gender, user2.Gender)
	require.Equal(t, user1.RoleID, user2.RoleID)
	require.Equal(t, user1.Avt, user2.Avt)
	require.WithinDuration(t, user1.CreateAt, user2.CreateAt, time.Second)
}

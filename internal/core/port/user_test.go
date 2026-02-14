package port_test

import (
	"context"
	"testing"

	"github.com/EduCoelhoTs/nba-predict-api/internal/core/domain"
	mock_port "github.com/EduCoelhoTs/nba-predict-api/internal/core/port/mock"
	"github.com/EduCoelhoTs/nba-predict-api/pkg/xuuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserService(t *testing.T) {

	testCtlr := gomock.NewController(t)
	defer testCtlr.Finish()
	ctx := context.Background()

	user := domain.NewUser(
		xuuid.NewV7(),
		"John",
		"Silva",
		"johnSilva@gmail.com",
		"2000-10-22",
		"12345678",
	)

	mockUserService := mock_port.NewMockUserServiceInterface(testCtlr)
	mockUserService.EXPECT().CreateUser(
		ctx,
		user.GetFirstName(),
		user.GetLastName(),
		user.GetEmail(),
		user.GetBirthDate(),
		user.GetPassword(),
	).Return(user, nil).Times(1)

	mockUserService.EXPECT().GetUserByID(ctx, user.GetID()).Return(user, nil).Times(1)
	mockUserService.EXPECT().GetAllUsers(ctx).Return([]domain.User{user}, nil).Times(1)
	mockUserService.EXPECT().GetUserByEmail(ctx, user.GetEmail()).Return(user, nil).Times(1)
	mockUserService.EXPECT().UpdateUser(
		ctx,
		user.GetID(),
		user.GetFirstName(),
		user.GetLastName(),
		user.GetEmail(),
		user.GetBirthDate(),
		user.GetPassword(),
	).Return(user, nil).Times(1)
	mockUserService.EXPECT().DeleteUser(ctx, user.GetID()).Return(nil).Times(1)

	testCases := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "should create user",
			fn: func(t *testing.T) {
				createdUser, err := mockUserService.CreateUser(
					ctx,
					user.GetFirstName(),
					user.GetLastName(),
					user.GetEmail(),
					user.GetBirthDate(),
					user.GetPassword(),
				)

				require.Nil(t, err)
				require.Equal(t, user.GetID(), createdUser.GetID())
			},
		},
		{
			name: "should get user by id",
			fn: func(t *testing.T) {
				foundUser, err := mockUserService.GetUserByID(ctx, user.GetID())

				require.Nil(t, err)
				require.Equal(t, user.GetID(), foundUser.GetID())
			},
		},
		{
			name: "should get all users",
			fn: func(t *testing.T) {
				users, err := mockUserService.GetAllUsers(ctx)

				require.Nil(t, err)
				require.Len(t, users, 1)
				require.Equal(t, user.GetID(), users[0].GetID())
			},
		},
		{
			name: "should get user by email",
			fn: func(t *testing.T) {
				foundUser, err := mockUserService.GetUserByEmail(ctx, user.GetEmail())

				require.Nil(t, err)
				require.Equal(t, user.GetID(), foundUser.GetID())
			},
		},
		{
			name: "should update user",
			fn: func(t *testing.T) {
				updatedUser, err := mockUserService.UpdateUser(
					ctx,
					user.GetID(),
					user.GetFirstName(),
					user.GetLastName(),
					user.GetEmail(),
					user.GetBirthDate(),
					user.GetPassword(),
				)

				require.Nil(t, err)
				require.Equal(t, user.GetID(), updatedUser.GetID())
			},
		},
		{
			name: "should delete user",
			fn: func(t *testing.T) {
				err := mockUserService.DeleteUser(ctx, user.GetID())

				require.Nil(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.fn)
	}
}

func TestUserRepository(t *testing.T) {
	testCtlr := gomock.NewController(t)
	defer testCtlr.Finish()
	ctx := context.Background()

	user := domain.NewUser(
		xuuid.NewV7(),
		"John",
		"Silva",
		"johnSilva@gmail.com",
		"2000-10-22",
		"12345678",
	)

	mockUserRepository := mock_port.NewMockUserRepositoryInterface(testCtlr)
	mockUserRepository.EXPECT().CreateUser(
		ctx,
		user,
	).Return(nil).Times(1)

	mockUserRepository.EXPECT().GetUserByID(ctx, user.GetID()).Return(user, nil).Times(1)
	mockUserRepository.EXPECT().GetAllUsers(ctx).Return([]domain.User{user}, nil).Times(1)
	mockUserRepository.EXPECT().GetUserByEmail(ctx, user.GetEmail()).Return(user, nil).Times(1)
	mockUserRepository.EXPECT().UpdateUser(
		ctx,
		user,
	).Return(nil).Times(1)
	mockUserRepository.EXPECT().DeleteUser(ctx, user.GetID()).Return(nil).Times(1)

	testCases := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "should create user",
			fn: func(t *testing.T) {
				err := mockUserRepository.CreateUser(ctx, user)

				require.Nil(t, err)
			},
		},
		{
			name: "should get user by id",
			fn: func(t *testing.T) {
				foundUser, err := mockUserRepository.GetUserByID(ctx, user.GetID())

				require.Nil(t, err)
				require.Equal(t, user.GetID(), foundUser.GetID())
			},
		},
		{
			name: "should get all users",
			fn: func(t *testing.T) {
				users, err := mockUserRepository.GetAllUsers(ctx)

				require.Nil(t, err)
				require.Len(t, users, 1)
				require.Equal(t, user.GetID(), users[0].GetID())
			},
		},
		{
			name: "should get user by email",
			fn: func(t *testing.T) {
				foundUser, err := mockUserRepository.GetUserByEmail(ctx, user.GetEmail())

				require.Nil(t, err)
				require.Equal(t, user.GetID(), foundUser.GetID())
			},
		},
		{
			name: "should update user",
			fn: func(t *testing.T) {
				err := mockUserRepository.UpdateUser(
					ctx,
					user,
				)

				require.Nil(t, err)
			},
		},
		{
			name: "should delete user",
			fn: func(t *testing.T) {
				err := mockUserRepository.DeleteUser(ctx, user.GetID())

				require.Nil(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.fn)
	}
}

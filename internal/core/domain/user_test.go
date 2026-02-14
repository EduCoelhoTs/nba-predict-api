package domain_test

import (
	"testing"

	"github.com/EduCoelhoTs/nba-predict-api/internal/core/domain"
	mock_domain "github.com/EduCoelhoTs/nba-predict-api/internal/core/domain/mock"
	"github.com/EduCoelhoTs/nba-predict-api/pkg/xuuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUser(t *testing.T) {
	testController := gomock.NewController(t)
	defer testController.Finish()

	mockUser := mock_domain.NewMockUser(testController)

	userTest := domain.NewUser(
		"abc",
		"John",
		"Silva",
		"johnSilva@gmail.com",
		"2000-10-22",
		"12345678",
	)

	testCase := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "should return user id",
			fn: func(t *testing.T) {
				mockUser.EXPECT().GetID().Return(userTest.GetID()).Times(1)

				require.Equal(t, mockUser.GetID(), userTest.GetID())
			},
		},
		{
			name: "should return user first name",
			fn: func(t *testing.T) {
				mockUser.EXPECT().GetFirstName().Return(userTest.GetFirstName()).Times(1)

				require.Equal(t, mockUser.GetFirstName(), userTest.GetFirstName())
			},
		},
		{
			name: "should return user last name",
			fn: func(t *testing.T) {
				mockUser.EXPECT().GetLastName().Return(userTest.GetLastName()).Times(1)

				require.Equal(t, mockUser.GetLastName(), userTest.GetLastName())
			},
		},
		{
			name: "should return user email",
			fn: func(t *testing.T) {
				mockUser.EXPECT().GetEmail().Return(userTest.GetEmail()).Times(1)

				require.Equal(t, mockUser.GetEmail(), userTest.GetEmail())
			},
		},
		{
			name: "should return user birth date",
			fn: func(t *testing.T) {
				mockUser.EXPECT().GetBirthDate().Return(userTest.GetBirthDate()).Times(1)

				require.Equal(t, mockUser.GetBirthDate(), userTest.GetBirthDate())
			},
		},
		{
			name: "should return user password",
			fn: func(t *testing.T) {
				mockUser.EXPECT().GetPassword().Return(userTest.GetPassword()).Times(1)

				require.Equal(t, mockUser.GetPassword(), userTest.GetPassword())
			},
		},
		{
			name: "should validate user",
			fn: func(t *testing.T) {
				mockUser.EXPECT().IsValid().Return(nil).Times(1)

				require.Nil(t, mockUser.IsValid())
			},
		},
		{
			name: "should return error when user is invalid",
			fn: func(t *testing.T) {
				testId := xuuid.NewV7()
				invalidUser := domain.NewUser(
					testId,
					"",
					"Silva",
					"johnSilva@gmail.com",
					"2000-10-22",
					"12345678",
				)

				err := invalidUser.IsValid()

				require.Equal(t, err.Error(), "field FirstName is invalid. Field must be: required")
			},
		},
		{
			name: "should return error when user email is invalid",
			fn: func(t *testing.T) {
				testId := xuuid.NewV7()
				invalidUser := domain.NewUser(
					testId,
					"John",
					"Silva",
					"johnSilva",
					"2000-10-22",
					"12345678",
				)

				err := invalidUser.IsValid()

				require.Equal(t, err.Error(), "field Email is invalid. Field must be: email")
			},
		},
		{
			name: "should return error when user password is invalid",
			fn: func(t *testing.T) {
				testId := xuuid.NewV7()
				invalidUser := domain.NewUser(
					testId,
					"John",
					"Silva",
					"johnSilva@gmail.com",
					"2000-10-22",
					"1234567",
				)

				err := invalidUser.IsValid()

				require.Equal(t, err.Error(), "field Password is invalid. Field must be: min")
			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, tc.fn)
	}
}

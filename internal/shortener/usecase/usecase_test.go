package usecase

import (
	"fmt"
	"testing"

	"github.com/damedelion/url_shortener/internal/shortener/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_Create_ShortURL(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock.NewMockRepository(ctrl)
	usecase := New(repositoryMock)

	longURL := "https://example.com/long-url"
	shortURL := "abc123"

	repositoryMock.EXPECT().GetShort(longURL).Return("", fmt.Errorf("not found"))
	repositoryMock.EXPECT().Create(gomock.Any(), longURL).Return(nil)

	result, err := usecase.Create(longURL)

	require.NoError(t, err)
	require.NotEqual(t, shortURL, result)
}

func TestUsecase_Create_AlreadyExists(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock.NewMockRepository(ctrl)
	usecase := New(repositoryMock)

	longURL := "https://example.com/long-url"
	shortURL := "abc123"

	repositoryMock.EXPECT().GetShort(longURL).Return(shortURL, nil)

	result, err := usecase.Create(longURL)

	require.NoError(t, err)
	require.Equal(t, shortURL, result)
}

func TestUsecase_Get_FoundURL(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock.NewMockRepository(ctrl)
	usecase := New(repositoryMock)

	shortURL := "abc123"
	longURL := "https://example.com/long-url"

	repositoryMock.EXPECT().GetLong(shortURL).Return(longURL, nil)

	result, err := usecase.Get(shortURL)

	require.NoError(t, err)
	require.Equal(t, longURL, result)
}

func TestUsecase_Get_URLNotFound(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock.NewMockRepository(ctrl)
	usecase := New(repositoryMock)

	shortURL := "abc123"

	repositoryMock.EXPECT().GetLong(shortURL).Return("", fmt.Errorf("not found"))

	result, err := usecase.Get(shortURL)

	require.Error(t, err)
	require.Empty(t, result)
	require.EqualError(t, err, "abc123 not found")
}

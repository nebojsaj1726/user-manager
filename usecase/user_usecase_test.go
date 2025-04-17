package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nebojsaj1726/user-manager/domain"
	"github.com/nebojsaj1726/user-manager/usecase"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserRepository struct {
	CreateFunc       func(ctx context.Context, user *domain.User) error
	FindByIDFunc     func(ctx context.Context, id primitive.ObjectID) (*domain.User, error)
	DeleteFunc       func(ctx context.Context, id string) error
	UpdateFunc       func(ctx context.Context, id string, user *domain.User) error
	CountFunc        func(ctx context.Context) (int64, error)
	FetchFunc        func(ctx context.Context, offset, limit int) ([]domain.User, error)
	FetchByEmailFunc func(ctx context.Context, email string) ([]domain.User, error)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ObjectID format")
	}
	return m.FindByIDFunc(ctx, objectID)
}

func (m *MockUserRepository) Fetch(ctx context.Context, offset int, limit int) ([]domain.User, error) {
	if m.FetchFunc != nil {
		return m.FetchFunc(ctx, offset, limit)
	}
	return nil, errors.New("FetchFunc not implemented")
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	return m.CreateFunc(ctx, user)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	return m.FindByIDFunc(ctx, id)
}

func (m *MockUserRepository) Delete(ctx context.Context, id string) error {
	return m.DeleteFunc(ctx, id)
}

func (m *MockUserRepository) Update(ctx context.Context, id string, user *domain.User) error {
	return m.UpdateFunc(ctx, id, user)
}

func (m *MockUserRepository) Count(ctx context.Context) (int64, error) {
	return m.CountFunc(ctx)
}

func (m *MockUserRepository) FetchByEmail(ctx context.Context, email string) ([]domain.User, error) {
	if m.FetchByEmailFunc != nil {
		return m.FetchByEmailFunc(ctx, email)
	}
	return nil, errors.New("FetchByEmailFunc not implemented")
}

func TestUserUseCase_Create(t *testing.T) {
	repoMock := &MockUserRepository{
		FetchByEmailFunc: func(ctx context.Context, email string) ([]domain.User, error) {
			if email == "existing@example.com" {
				return []domain.User{
					{ID: primitive.NewObjectID(), Email: "existing@example.com", Age: 25},
				}, nil
			}
			return []domain.User{}, nil
		},
		CreateFunc: func(ctx context.Context, user *domain.User) error {
			return nil
		},
	}

	userUseCase := usecase.NewUserUseCase(repoMock, 10*time.Second)

	testUser := &domain.User{
		ID:    primitive.NewObjectID(),
		Email: "test@example.com",
		Age:   21,
	}

	err := userUseCase.Create(context.TODO(), testUser)
	assert.NoError(t, err)

	testUser = &domain.User{
		ID:    primitive.NewObjectID(),
		Email: "test@example.com",
		Age:   17,
	}

	err = userUseCase.Create(context.TODO(), testUser)
	assert.Error(t, err)
	assert.Equal(t, "age must be greater than 18", err.Error())

	testUser = &domain.User{
		ID:    primitive.NewObjectID(),
		Email: "existing@example.com",
		Age:   21,
	}

	err = userUseCase.Create(context.TODO(), testUser)
	assert.Error(t, err)
	assert.Equal(t, "email must be unique", err.Error())
}

func TestUserUseCase_Fetch(t *testing.T) {
	repoMock := &MockUserRepository{
		FetchFunc: func(ctx context.Context, offset, limit int) ([]domain.User, error) {
			return []domain.User{
				{ID: primitive.NewObjectID(), Email: "user1@example.com", Age: 25},
				{ID: primitive.NewObjectID(), Email: "user2@example.com", Age: 30},
			}, nil
		},
	}

	userUseCase := usecase.NewUserUseCase(repoMock, 10*time.Second)

	users, err := userUseCase.Fetch(context.TODO(), 1, 2)
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	repoMock.FetchFunc = func(ctx context.Context, offset, limit int) ([]domain.User, error) {
		return nil, errors.New("fetch failed")
	}

	users, err = userUseCase.Fetch(context.TODO(), 1, 2)
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Equal(t, "fetch failed", err.Error())
}

func TestUserUseCase_GetByID(t *testing.T) {
	testID := primitive.NewObjectID()
	repoMock := &MockUserRepository{
		FindByIDFunc: func(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
			if id == testID {
				return &domain.User{ID: id, Email: "test@example.com", Age: 21}, nil
			}
			return nil, errors.New("user not found")
		},
	}

	userUseCase := usecase.NewUserUseCase(repoMock, 10*time.Second)

	user, err := userUseCase.GetByID(context.TODO(), testID.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, testID, user.ID)

	invalidID := primitive.NewObjectID()
	user, err = userUseCase.GetByID(context.TODO(), invalidID.Hex())
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "user not found", err.Error())
}

func TestUserUseCase_Update(t *testing.T) {
	testID := primitive.NewObjectID()
	repoMock := &MockUserRepository{
		FetchByEmailFunc: func(ctx context.Context, email string) ([]domain.User, error) {
			if email == "existing@example.com" {
				return []domain.User{
					{ID: primitive.NewObjectID(), Email: "existing@example.com", Age: 25},
				}, nil
			}
			return []domain.User{}, nil
		},
		UpdateFunc: func(ctx context.Context, id string, user *domain.User) error {
			return nil
		},
	}

	userUseCase := usecase.NewUserUseCase(repoMock, 10*time.Second)

	testUser := &domain.User{
		ID:    testID,
		Email: "updated@example.com",
		Age:   22,
	}

	err := userUseCase.Update(context.TODO(), testID.Hex(), testUser)
	assert.NoError(t, err)

	testUser = &domain.User{
		ID:    testID,
		Email: "updated@example.com",
		Age:   17,
	}

	err = userUseCase.Update(context.TODO(), testID.Hex(), testUser)
	assert.Error(t, err)
	assert.Equal(t, "age must be greater than 18", err.Error())

	testUser = &domain.User{
		ID:    testID,
		Email: "existing@example.com",
		Age:   22,
	}

	err = userUseCase.Update(context.TODO(), testID.Hex(), testUser)
	assert.Error(t, err)
	assert.Equal(t, "email must be unique", err.Error())
}

func TestUserUseCase_Delete(t *testing.T) {
	testID := primitive.NewObjectID()

	repoMock := &MockUserRepository{
		DeleteFunc: func(ctx context.Context, id string) error {
			return nil
		},
	}

	userUseCase := usecase.NewUserUseCase(repoMock, 10*time.Second)

	err := userUseCase.Delete(context.TODO(), testID.Hex())
	assert.NoError(t, err)

	repoMock.DeleteFunc = func(ctx context.Context, id string) error {
		return errors.New("delete failed")
	}

	err = userUseCase.Delete(context.TODO(), testID.Hex())
	assert.Error(t, err)
	assert.Equal(t, "delete failed", err.Error())
}

func TestUserUseCase_Count(t *testing.T) {
	repoMock := &MockUserRepository{
		CountFunc: func(ctx context.Context) (int64, error) {
			return 42, nil
		},
	}

	userUseCase := usecase.NewUserUseCase(repoMock, 10*time.Second)

	count, err := userUseCase.Count(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, int64(42), count)

	repoMock.CountFunc = func(ctx context.Context) (int64, error) {
		return 0, errors.New("count failed")
	}

	count, err = userUseCase.Count(context.TODO())
	assert.Error(t, err)
	assert.Equal(t, int64(0), count)
	assert.Equal(t, "count failed", err.Error())
}

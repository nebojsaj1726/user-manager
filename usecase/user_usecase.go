package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/nebojsaj1726/user-manager/domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if user.Age <= 18 {
		return fmt.Errorf("age must be greater than 18")
	}

	existingUsers, err := u.userRepository.FetchByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if len(existingUsers) > 0 {
		return fmt.Errorf("email must be unique")
	}

	return u.userRepository.Create(ctx, user)
}

func (u *userUsecase) Fetch(c context.Context, page, limit int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	offset := (page - 1) * limit

	return u.userRepository.Fetch(ctx, offset, limit)
}

func (u *userUsecase) GetByID(c context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetByID(ctx, id)
}

func (u *userUsecase) Update(c context.Context, id string, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if user.Age <= 18 {
		return fmt.Errorf("age must be greater than 18")
	}

	existingUsers, err := u.userRepository.FetchByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	for _, existingUser := range existingUsers {
		if existingUser.ID.Hex() != id {
			return fmt.Errorf("email must be unique")
		}
	}

	return u.userRepository.Update(ctx, id, user)
}

func (u *userUsecase) Delete(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.Delete(ctx, id)
}

func (u *userUsecase) Count(c context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.Count(ctx)
}

package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Age   int                `bson:"age" form:"age" json:"age"`
	Email string             `bson:"email" form:"email" binding:"required,email" json:"email"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context, offset, limit int) ([]User, error)
	FetchByEmail(c context.Context, email string) ([]User, error)
	GetByID(c context.Context, id string) (*User, error)
	Update(c context.Context, id string, user *User) error
	Delete(c context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}

type UserUsecase interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context, page, limit int) ([]User, error)
	GetByID(c context.Context, id string) (*User, error)
	Update(c context.Context, id string, user *User) error
	Delete(c context.Context, id string) error
	Count(c context.Context) (int64, error)
}

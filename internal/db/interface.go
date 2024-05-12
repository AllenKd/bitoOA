package db

import (
	"bitoOA/internal/common"
	"context"
)

type Service interface {
	InsertUser(ctx context.Context, user *User) string
	GetUser(ctx context.Context, userId string) *User
	GetUsers(ctx context.Context, gender common.Gender, height uint) []*User
	RemoveUser(ctx context.Context, userId string)
	GetPopularUsers(ctx context.Context, limit int) []*User
}

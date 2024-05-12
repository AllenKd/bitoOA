package user

import (
	"bitoOA/internal/common"
	"bitoOA/internal/db"
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, name string, height uint, gender common.Gender, dates uint) string
	RemoveUser(ctx context.Context, userId string)
	GetCandidate(ctx context.Context, userGender common.Gender, userHeight uint, limit int) []*db.User
	DoLike(ctx context.Context, fromUserId string, toUserId string) (bool, error)
	GetPopularUsers(ctx context.Context, limit int) []*db.User
}

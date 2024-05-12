package user

import (
	"bitoOA/internal/btError"
	"bitoOA/internal/common"
	"bitoOA/internal/db"
	"bitoOA/internal/logger"
	"context"
	"math/rand"
)

type service struct {
	db  db.Service
	log *logger.Logger
}

func New(db db.Service, log *logger.Logger) Service {
	return &service{
		db:  db,
		log: log.With("service", "user"),
	}
}

func (s *service) CreateUser(ctx context.Context, name string, height uint, gender common.Gender, dates uint) string {
	log := s.log.WithContext(ctx)
	user := db.User{
		Name:                name,
		Height:              height,
		Gender:              gender,
		NumberOfWantedDates: dates,

		Liked:      map[string]bool{},
		Matched:    map[string]bool{},
		BeingLiked: map[string]bool{},
	}
	userId := s.db.InsertUser(ctx, &user)
	log.With("userId", userId).Info("user created")
	return userId
}

func (s *service) RemoveUser(ctx context.Context, userId string) {
	log := s.log.WithContext(ctx)

	s.db.RemoveUser(ctx, userId)
	log.With("userId", userId).Info("user removed")
}

func (s *service) GetCandidate(ctx context.Context, userGender common.Gender, userHeight uint, limit int) []*db.User {
	log := s.log.WithContext(ctx)
	var candidates []*db.User

	switch userGender {
	case common.Boy:
		candidates = s.db.GetUsers(ctx, common.Girl, userHeight)
	case common.Girl:
		candidates = s.db.GetUsers(ctx, common.Boy, userHeight)
	default:
		log.With("gender", userGender).Warn("invalid gender")
	}

	if len(candidates) <= limit {
		return candidates
	}

	rand.Shuffle(len(candidates), func(i, j int) { candidates[i], candidates[j] = candidates[j], candidates[i] })
	return candidates[:limit]
}

func (s *service) DoLike(ctx context.Context, fromUserId string, toUserId string) (bool, error) {
	log := s.log.WithContext(ctx)

	fu := s.db.GetUser(ctx, fromUserId)
	if fu == nil {
		log.With("userId", fromUserId).Warn("invalid user id")
		return false, btError.UserNotFound
	}
	if fu.NumberOfWantedDates <= 0 {
		log.With("userId", fromUserId).Warn("user out of dates")
		return false, nil
	}

	tu := s.db.GetUser(ctx, toUserId)
	if tu == nil {
		log.With("userId", toUserId).Warn("invalid user id")
		return false, btError.UserNotFound
	}
	if tu.NumberOfWantedDates <= 0 {
		log.With("userId", toUserId).Warn("user out of dates")
		return false, nil
	}

	fu.Liked[toUserId] = true
	tu.BeingLiked[fromUserId] = true

	// no match
	if _, ok := tu.Liked[fromUserId]; !ok {
		return false, nil
	}

	// matched
	log.With("fromUser", fromUserId).With("toUser", toUserId).Info("user matched")
	fu.NumberOfWantedDates -= 1
	tu.NumberOfWantedDates -= 1
	fu.Matched[toUserId] = true
	tu.Matched[fromUserId] = true
	return true, nil
}

func (s *service) GetPopularUsers(ctx context.Context, limit int) []*db.User {
	users := s.db.GetPopularUsers(ctx, limit)
	return users
}

package db

import (
	"bitoOA/internal/common"
	"context"
	"github.com/google/uuid"
	"sort"
)

type User struct {
	Id                  string        `json:"id" example:"c30790bf-75e0-4cd3-9ec2-c25662f77e33"`
	Name                string        `json:"name" example:"Bob"`
	Height              uint          `json:"height" example:"180"`
	Gender              common.Gender `json:"gender" example:"1"` // 0 for female, 1 for male
	NumberOfWantedDates uint          `json:"numberOfWantedDates" example:"10"`

	Blocked     bool            `json:"-"`
	Liked       map[string]bool `json:"-"`
	BeingLicked map[string]bool `json:"-"`
	Matched     map[string]bool `json:"-"`
}

func (c *service) InsertUser(ctx context.Context, user *User) string {
	c.log.WithContext(ctx).Debug("insert user")

	user.Id = uuid.NewString()
	c.table[user.Id] = user
	return user.Id
}

// GetUsers
// for boy: get users witch higher than height
// for girl: get users witch lower than height
func (c *service) GetUsers(ctx context.Context, gender common.Gender, height uint) []*User {
	var users []*User
	switch gender {
	case common.Boy:
		for _, user := range c.table {
			if !user.Blocked && user.Gender == gender && user.NumberOfWantedDates > 0 && user.Height > height {
				users = append(users, user)
			}
		}
	case common.Girl:
		for _, user := range c.table {
			if !user.Blocked && user.Gender == gender && user.NumberOfWantedDates > 0 && user.Height < height {
				users = append(users, user)
			}
		}
	default:
		c.log.WithContext(ctx).With("gender", gender).Warn("invalid gender")
	}

	return users
}

func (c *service) GetUser(ctx context.Context, userId string) *User {
	log := c.log.WithContext(ctx)

	user, ok := c.table[userId]
	if !ok {
		log.With("userId", userId).Warn("user not found")
		return nil
	}

	return user
}

func (c *service) RemoveUser(ctx context.Context, userId string) {
	c.log.WithContext(ctx).With("userId", userId).Debug("remove user")
	delete(c.table, userId)
}

func (c *service) GetPopularUsers(ctx context.Context, limit int) []*User {
	users := make([]*User, len(c.table))
	i := 0
	for _, user := range c.table {
		users[i] = user
		i++
	}

	sort.Slice(users[:], func(i, j int) bool {
		return len(users[i].BeingLicked) > len(users[j].BeingLicked)
	})

	if len(users) <= limit {
		return users
	}

	return users[:limit]
}

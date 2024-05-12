package db

import (
	"bitoOA/internal/common"
	"bitoOA/internal/config"
	"bitoOA/internal/logger"
	"context"
	"testing"
)

var s = service{
	table: map[string]*User{},
	log:   logger.New(&config.Config{}),
}

func Test_service_InsertUser(t *testing.T) {
	ctx := context.Background()

	tu := User{
		Name:                "Bob",
		Height:              180,
		Gender:              common.Boy,
		NumberOfWantedDates: 10,
		Liked:               map[string]bool{},
		BeingLiked:          map[string]bool{},
		Matched:             map[string]bool{},
	}
	s.InsertUser(ctx, &tu)

	if len(s.table) != 1 {
		t.Errorf("test fail")
	}
}

func Test_service_GetUsers(t *testing.T) {
	ctx := context.Background()

	tu := User{
		Name:                "Bob",
		Height:              180,
		Gender:              common.Boy,
		NumberOfWantedDates: 10,
		Liked:               map[string]bool{},
		BeingLiked:          map[string]bool{},
		Matched:             map[string]bool{},
	}
	s.InsertUser(ctx, &tu)

	if len(s.table) != 1 {
		t.Errorf("test fail")
	}

	if len(s.GetUsers(ctx, common.Boy, 180)) != 0 {
		t.Errorf("test fail")
	}

	if len(s.GetUsers(ctx, common.Boy, 181)) != 0 {
		t.Errorf("test fail")
	}

	if len(s.GetUsers(ctx, common.Boy, 179)) != 1 {
		t.Errorf("test fail")
	}

	if len(s.GetUsers(ctx, common.Girl, 180)) != 0 {
		t.Errorf("test fail")
	}

	if len(s.GetUsers(ctx, common.Girl, 181)) != 0 {
		t.Errorf("test fail")
	}

	if len(s.GetUsers(ctx, common.Girl, 179)) != 0 {
		t.Errorf("test fail")
	}
}

func Test_service_GetUser(t *testing.T) {
	ctx := context.Background()

	tu := User{
		Name:                "Bob",
		Height:              180,
		Gender:              common.Boy,
		NumberOfWantedDates: 10,
		Liked:               map[string]bool{},
		BeingLiked:          map[string]bool{},
		Matched:             map[string]bool{},
	}
	uId := s.InsertUser(ctx, &tu)

	if s.GetUser(ctx, uId) == nil {
		t.Errorf("test fail")
	}

	if s.GetUser(ctx, "wrongId") != nil {
		t.Errorf("test fail")
	}

	if u := s.GetUser(ctx, uId); u.Name != tu.Name {
		t.Errorf("test fail")
	}
}

func Test_service_RemoveUser(t *testing.T) {
	ctx := context.Background()

	tu := User{
		Name:                "Bob",
		Height:              180,
		Gender:              common.Boy,
		NumberOfWantedDates: 10,
		Liked:               map[string]bool{},
		BeingLiked:          map[string]bool{},
		Matched:             map[string]bool{},
	}
	uId := s.InsertUser(ctx, &tu)

	s.RemoveUser(ctx, uId)
	if len(s.table) != 0 {
		t.Errorf("test fail")
	}
}

func Test_service_GetPopularUsers(t *testing.T) {
	ctx := context.Background()

	tu1 := User{
		Name:                "Bob",
		Height:              180,
		Gender:              common.Boy,
		NumberOfWantedDates: 10,
		Liked:               map[string]bool{},
		BeingLiked: map[string]bool{
			"uId001": true,
			"uId002": true,
			"uId003": true,
			"uId004": true,
			"uId005": true,
			"uId006": true,
		},
		Matched: map[string]bool{},
	}
	tu2 := User{
		Name:                "Alice",
		Height:              160,
		Gender:              common.Girl,
		NumberOfWantedDates: 10,
		Liked:               map[string]bool{},
		BeingLiked: map[string]bool{
			"uId001": true,
			"uId002": true,
			"uId003": true,
			"uId004": true,
			"uId005": true,
			"uId006": true,
			"uId007": true,
		},
		Matched: map[string]bool{},
	}
	tu3 := User{
		Name:                "Lisa",
		Height:              163,
		Gender:              common.Girl,
		NumberOfWantedDates: 10,
		Liked:               map[string]bool{},
		BeingLiked: map[string]bool{
			"uId001": true,
			"uId002": true,
			"uId003": true,
			"uId004": true,
			"uId005": true,
			"uId006": true,
			"uId007": true,
			"uId008": true,
		},
		Matched: map[string]bool{},
	}
	s.InsertUser(ctx, &tu1)
	s.InsertUser(ctx, &tu2)
	s.InsertUser(ctx, &tu3)

	if len(s.GetPopularUsers(ctx, 5)) != 3 {
		t.Errorf("test fail")
	}

	if len(s.GetPopularUsers(ctx, 3)) != 3 {
		t.Errorf("test fail")
	}

	if len(s.GetPopularUsers(ctx, 1)) != 1 {
		t.Errorf("test fail")
	}

	if u := s.GetPopularUsers(ctx, 1); u[0].Name != tu3.Name {
		t.Errorf("test fail")
	}

	if u := s.GetPopularUsers(ctx, 2); u[0].Name != tu3.Name || u[1].Name != tu2.Name {
		t.Errorf("test fail")
	}
}

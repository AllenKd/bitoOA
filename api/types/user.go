package types

import (
	"bitoOA/internal/common"
	"bitoOA/internal/db"
)

type HttpCreateUserReq struct {
	Name                string        `json:"name" example:"Bob"`
	Height              uint          `json:"height" example:"180"`
	Gender              common.Gender `json:"gender" example:"1"` // 0 for female, 1 for male
	NumberOfWantedDates uint          `json:"numberOfWantedDates" example:"10"`

	MatchesLimit int `json:"matchesLimit" example:"10"` // optional, default=10
}

type HttpCreateUserResp struct {
	HttpRespBase
	Data HttpCreateUserRespData `json:"data"`
}

type HttpCreateUserRespData struct {
	UserId     string     `json:"userId" example:"30e897aa-2161-481b-9f61-e0700f467bba"`
	Candidates []*db.User `json:"candidates"`
}

type HttpGetPopularUsersResp struct {
	HttpRespBase
	Data HttpGetPopularUserRespData `json:"data"`
}

type HttpGetPopularUserRespData struct {
	PopularUsers []*db.User `json:"popularUsers"`
}

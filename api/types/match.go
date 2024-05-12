package types

type HttpLikeReq struct {
	FromUserId string `json:"fromUserId" example:"30e897aa-2161-481b-9f61-e0700f467bba"`
	ToUserId   string `json:"toUserId" example:"c30790bf-75e0-4cd3-9ec2-c25662f77e33"`
}

type HttpLikeResp struct {
	HttpRespBase
	Data HttpLikeRespData `json:"data"`
}

type HttpLikeRespData struct {
	Matched bool `json:"matched"` // true if users like each other
}

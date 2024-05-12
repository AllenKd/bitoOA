package handler

import (
	"bitoOA/api/types"
	"bitoOA/internal/btError"
	"bitoOA/internal/logger"
	"bitoOA/internal/route/middleware"
	"bitoOA/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	service user.Service
	log     *logger.Logger
}

func New(service user.Service, log *logger.Logger) *UserController {
	return &UserController{
		service: service,
		log:     log.With("service", "userController"),
	}
}

// QuerySinglePeople
// @Summary Find the most N possible matched single people, where N is a request parameter
// @Tags V1, Matches
// @version 1.0
// @produce json
// @Param limit query int false "limit, default=10"
// @Success 200 {object} types.HttpGetPopularUsersResp "Response"
// @failure 400 {object} types.HttpRespError
// @Router /api/v1/matches/popular-users [get]
func (u *UserController) QuerySinglePeople(c *gin.Context) {
	log := u.log.WithContext(c)
	log.Debug("handle get popular users")

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		log.Warn("invalid query parameter, use default 10")
		limit = 10
	}

	users := u.service.GetPopularUsers(c, limit)

	c.JSON(http.StatusOK, types.HttpGetPopularUsersResp{
		HttpRespBase: types.HttpRespBase{
			RequestId: c.GetHeader(middleware.HeaderRequestId),
			Code:      types.HttpRespCodeOk,
			Message:   types.HttpRespMsgOk,
		},
		Data: types.HttpGetPopularUserRespData{
			PopularUsers: users,
		},
	})
}

// AddSinglePersonAndMatch
// @Summary Add a new user to the matching system and find any possible matches for the new user.
// @Tags V1,User
// @version 1.0
// @Param request body types.HttpCreateUserReq true "Log"
// @Accept  json
// @produce json
// @Success 200 {object} types.HttpCreateUserResp "Response"
// @failure 400 {object} types.HttpRespError
// @Router /api/v1/user [post]
func (u *UserController) AddSinglePersonAndMatch(c *gin.Context) {
	log := u.log.WithContext(c)
	log.Debug("handle create user")

	var req types.HttpCreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.WithError(err).Warn("invalid body")
		c.JSON(http.StatusBadRequest, types.HttpRespBase{
			RequestId: c.GetHeader(middleware.HeaderRequestId),
			Code:      int(btError.InvalidBody.Code),
			Message:   btError.InvalidBody.Message,
		})
		return
	}

	userId := u.service.CreateUser(c, req.Name, req.Height, req.Gender, req.NumberOfWantedDates)

	var limit int
	if req.MatchesLimit == 0 {
		limit = 10
	} else {
		limit = req.MatchesLimit
	}
	candidates := u.service.GetCandidate(c, req.Gender, req.Height, limit)

	c.JSON(http.StatusOK, types.HttpCreateUserResp{
		HttpRespBase: types.HttpRespBase{
			RequestId: c.GetHeader(middleware.HeaderRequestId),
			Code:      types.HttpRespCodeOk,
			Message:   types.HttpRespMsgOk,
		},
		Data: types.HttpCreateUserRespData{
			UserId:     userId,
			Candidates: candidates,
		},
	})
}

// RemoveSinglePerson
// @Summary Remove a user from the matching system so that the user cannot be matched anymore.
// @Tags V1,User
// @version 1.0
// @produce json
// @Success 200 {object} types.HttpRespBase "Response"
// @failure 400 {object} types.HttpRespError
// @Param userId path string true "User ID"
// @Router /api/v1/users/{userId} [delete]
func (u *UserController) RemoveSinglePerson(c *gin.Context) {
	log := u.log.WithContext(c)
	log.Debug("handle block user")

	userId := c.Param("userId")

	u.service.RemoveUser(c, userId)
	if userId == "" {
		log.Warn("empty user id")
		c.JSON(http.StatusBadRequest, types.HttpRespBase{
			RequestId: c.GetHeader(middleware.HeaderRequestId),
			Code:      int(btError.InvalidBody.Code),
			Message:   btError.InvalidBody.Message,
		})
		return
	}

	c.JSON(http.StatusOK, types.HttpRespBase{
		RequestId: c.GetHeader(middleware.HeaderRequestId),
		Code:      types.HttpRespCodeOk,
		Message:   types.HttpRespMsgOk,
	})
}

// Like
// @Summary Send a like from a user to another.
// @Tags V1,Matches
// @version 1.0
// @Param request body types.HttpLikeReq true "Log"
// @Accept  json
// @produce json
// @Success 200 {object} types.HttpLikeResp "Response"
// @failure 400 {object} types.HttpRespError
// @Router /api/v1/matches/like [post]
func (u *UserController) Like(c *gin.Context) {
	log := u.log.WithContext(c)
	log.Debug("handle user like")

	var req types.HttpLikeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.WithError(err).Warn("invalid body")
		c.JSON(http.StatusBadRequest, types.HttpRespBase{
			RequestId: c.GetHeader(middleware.HeaderRequestId),
			Code:      int(btError.InvalidBody.Code),
			Message:   btError.InvalidBody.Message,
		})
		return
	}

	matched, err := u.service.DoLike(c, req.FromUserId, req.ToUserId)
	if err != nil {
		log.WithError(err).Error("fail to like user")
		btErr, ok := btError.ToBtError(err)
		if !ok {
			log.WithError(err).Error("unknown error")
			c.JSON(http.StatusInternalServerError, types.HttpRespBase{
				RequestId: c.GetHeader(middleware.HeaderRequestId),
				Code:      int(btError.Unknown.Code),
				Message:   btError.Unknown.Message,
			})
			return
		}
		c.JSON(btErr.GetStatus(), types.HttpRespBase{
			RequestId: c.GetHeader(middleware.HeaderRequestId),
			Code:      int(btErr.Code),
			Message:   btErr.Message,
		})
		return
	}

	c.JSON(http.StatusOK, types.HttpLikeResp{
		HttpRespBase: types.HttpRespBase{
			RequestId: c.GetHeader(middleware.HeaderRequestId),
			Code:      types.HttpRespCodeOk,
			Message:   types.HttpRespMsgOk,
		},
		Data: types.HttpLikeRespData{
			Matched: matched,
		},
	})
}

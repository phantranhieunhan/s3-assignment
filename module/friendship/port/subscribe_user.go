package port

import (
	"net/http"

	"github.com/phantranhieunhan/s3-assignment/common"
	"github.com/phantranhieunhan/s3-assignment/common/logger"
	"github.com/phantranhieunhan/s3-assignment/module/friendship/app/command"
	"github.com/phantranhieunhan/s3-assignment/module/friendship/port/constant"

	"github.com/gin-gonic/gin"
)

type SubscribeUserReq struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

func (l SubscribeUserReq) validate() error {
	if err := common.ValidateRequired(l.Requestor, constant.REQUESTOR); err != nil {
		return err
	}
	if err := common.ValidateRequired(l.Target, constant.TARGET); err != nil {
		return err
	}
	return nil
}

func (s *Server) SubscribeUser(c *gin.Context) {
	var req SubscribeUserReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		logger.Error("ListCommonFriends.ShouldBind: ", err)
		panic(common.ErrInvalidRequest(err, "body data"))
	}

	if err = req.validate(); err != nil {
		logger.Error("ListCommonFriends.Validate: ", err)
		panic(err)
	}

	err = s.app.Commands.SubscribeUser.Handle(c.Request.Context(), command.SubscriberUserPayloads{
		command.SubscriberUserPayload{
			Requestor: req.Requestor,
			Target:    req.Target,
		},
	})
	if err != nil {
		logger.Error("ListFriends.Handle: ", err)
		panic(err)
	}

	c.JSON(http.StatusOK, common.CustomSuccessResponse(nil))
}

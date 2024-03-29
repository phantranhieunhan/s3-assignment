package port

import (
	"github.com/gin-gonic/gin"
	"github.com/phantranhieunhan/s3-assignment/module/friendship/app"
)

type Server struct {
	app app.Application
}

func NewServer(app app.Application) Server {
	s := Server{app: app}

	return s
}

func (s Server) Router(r *gin.Engine) {
	friendship := r.Group("friendship")
	friendship.POST("connect", s.ConnectFriendship)
	friendship.GET("friends", s.ListFriends)
	friendship.GET("mutuals", s.ListCommonFriends)

	subscription := r.Group("subscription")
	subscription.POST("subscribe", s.SubscribeUser)
	subscription.POST("block", s.BlockUpdatesUser)
	subscription.GET("updates_user", s.ListUpdatesUser)
}

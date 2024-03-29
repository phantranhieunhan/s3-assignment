package friendship

import (
	"github.com/gin-gonic/gin"

	"github.com/phantranhieunhan/s3-assignment/common/adapter/postgres"
	"github.com/phantranhieunhan/s3-assignment/module/friendship/adapter/postgres/repository"
	"github.com/phantranhieunhan/s3-assignment/module/friendship/app"
	"github.com/phantranhieunhan/s3-assignment/module/friendship/app/command"
	"github.com/phantranhieunhan/s3-assignment/module/friendship/app/query"
	"github.com/phantranhieunhan/s3-assignment/module/friendship/port"
)

func New(r *gin.Engine, db postgres.Database) {
	friendshipRepo := repository.NewFriendshipRepository(db)
	userRepo := repository.NewUserRepository(db)
	subRepo := repository.NewSubscriptionRepository(db)

	application := app.Application{
		Commands: app.Commands{
			ConnectFriendship: command.NewConnectFriendshipHandler(friendshipRepo, userRepo, db),
			SubscribeUser:     command.NewSubscribeUserHandler(friendshipRepo, userRepo, subRepo, db),
			BlockUpdatesUser:  command.NewBlockUpdatesUserHandler(friendshipRepo, userRepo, subRepo, db),
		},
		Queries: app.Queries{
			ListFriends:       query.NewListFriendsHandler(friendshipRepo, userRepo),
			ListCommonFriends: query.NewListCommonFriendsHandler(friendshipRepo, userRepo),
			ListUpdatesUser:   query.NewListUpdatesUserHandler(subRepo, userRepo),
		},
	}
	port.NewServer(application).Router(r)
}

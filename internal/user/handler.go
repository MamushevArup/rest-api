package user

import (
	"github.com/MamushevArup/server/internal/handlers"
	"github.com/MamushevArup/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

const (
	usersUrl = "/users"
	userUrl  = "/users/:uuid"
)

type handler struct {
	logger *logger.Logger
}

func NewHandler(lg *logger.Logger) handlers.Handler {
	return &handler{logger: lg}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(usersUrl, h.GetUsers)
	router.GET(userUrl, h.GetUser)
	router.POST(usersUrl, h.CreateUser)
	router.PUT(userUrl, h.UpdateUser)
	router.PATCH(userUrl, h.UpdateUserPart)
	router.DELETE(userUrl, h.DeleteUser)
}
func (h *handler) GetUsers(ctx *gin.Context) {
	ctx.JSONP(200, gin.H{"Hey": "there"})
}

func (h *handler) GetUser(context *gin.Context) {

}

func (h *handler) CreateUser(context *gin.Context) {

}

func (h *handler) UpdateUser(context *gin.Context) {

}

func (h *handler) UpdateUserPart(context *gin.Context) {

}

func (h *handler) DeleteUser(context *gin.Context) {

}

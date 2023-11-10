package main

import (
	"fmt"
	"github.com/MamushevArup/server/internal/config"
	"github.com/MamushevArup/server/internal/user"
	"github.com/MamushevArup/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	l := logger.NewLogger()
	r := gin.Default()

	cfg := config.ReadConfig()

	handler := user.NewHandler(l)
	handler.Register(r)

	if err := r.Run(fmt.Sprintf("%s:%s", cfg.BindIp, cfg.Port)); err != nil {
		l.Fatal("Error with running server ", err.Error())
	}
}

package api

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ismdeep/log"
	"go.uber.org/zap"

	"github.com/ismdeep/notification/app/relay/conf"
	"github.com/ismdeep/notification/pkg/core"
)

func safer(ginFunc func(*gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("ERR: %v\n", r)
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": r.(error).Error(),
				})
			}
		}()
		ginFunc(c)
	}
}

//go:embed raw_relay.sh
var rawRelayScript string

//go:embed encrypt_relay.sh
var encryptRelayScript string

var eng *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	eng = gin.New()
	eng.Use(func(c *gin.Context) {
		switch {
		case c.GetHeader("X-Relay-Auth") == conf.ROOT.Authorization.RelayAuth:
			c.Next()
			return
		default:
			c.String(http.StatusUnauthorized, "401 Unauthorized")
			c.Abort()
			return
		}
	})
	eng.PUT("/api/v1/msg/:customer_msg_id", safer(relay))
}

func Run() {
	bind := conf.ROOT.Server.Bind
	log.WithContext(context.Background()).Info("http server started", zap.Any("bind", bind))
	core.PanicIf(eng.Run(bind))
}

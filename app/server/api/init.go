package api

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/ismdeep/log"
	"go.uber.org/zap"

	"github.com/ismdeep/notification/app/server/conf"
	"github.com/ismdeep/notification/app/server/store"
	"github.com/ismdeep/notification/pkg/core"
)

// eng instance
var eng *gin.Engine

func init() {
	cjson := func(ginFunc func(*gin.Context) any) gin.HandlerFunc {
		return func(c *gin.Context) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("ERR: %v\n", r)
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": r.(error).Error(),
					})
				}
			}()
			c.JSON(http.StatusOK, ginFunc(c))
			c.Abort()
		}
	}

	gin.SetMode(gin.ReleaseMode)
	eng = gin.Default()
	eng.POST("/api/v1/sign-in", cjson(SignIn))           // 登录
	eng.GET("/api/v1/profile", Auth, cjson(GetUserInfo)) // 用户信息

	// Token
	eng.POST("/api/v1/tokens", Auth, cjson(GenerateToken)) // 生成Token
	eng.GET("/api/v1/tokens", Auth, cjson(GetTokenList))   // 获取Token列表

	// MSG
	eng.PUT("/api/v1/msg/:customer_msg_id", func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("ERROR: %v\n", r))
			}
		}()
		// customer_msg_id
		customerMsgID := c.Param("customer_msg_id")

		// token
		token := store.Token.GetByToken(c.GetHeader("X-Token"))

		cmd := exec.Command("bash",
			"-c",
			fmt.Sprintf(
				`base64 -d | openssl aes-256-cbc -d -salt -pbkdf2 -iter 1024 -k "%v"`,
				token.AESKey))

		// input from http post body
		cmd.Stdin = c.Request.Body

		// write to output
		var output bytes.Buffer
		cmd.Stdout = &output

		// write stderr to output2
		var output2 bytes.Buffer
		cmd.Stderr = &output2

		if err := cmd.Run(); err != nil {
			panic(
				errors.Join(err,
					fmt.Errorf("stdout: %v", output.String()),
					fmt.Errorf("stderr: %v", output2.String())))
		}

		msgID := store.Msg.Write(token.UserID, customerMsgID, output.String())
		c.String(http.StatusOK, fmt.Sprintf("%v %v\n", msgID, customerMsgID))
	})
}

func Run() {
	bind := conf.ROOT.Server.Bind
	log.WithContext(context.Background()).Info("http server started", zap.Any("bind", bind))
	core.PanicIf(eng.Run(bind))
}

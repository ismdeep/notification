package api

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ismdeep/log"
	"go.uber.org/zap"

	"github.com/ismdeep/notification/app/relay/conf"
	"github.com/ismdeep/notification/pkg/core"
)

//go:embed raw_relay.sh
var rawRelayScript string

//go:embed encrypt_relay.sh
var encryptRelayScript string

var eng *gin.Engine

func init() {
	proxy := func(ginFunc func(*gin.Context)) gin.HandlerFunc {
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

	gin.SetMode(gin.ReleaseMode)
	eng = gin.New()

	eng.PUT("/api/v1/msg/:customer_msg_id", proxy(func(c *gin.Context) {
		// customer_msg_id
		customerMsgID := c.Param("customer_msg_id")

		// read data
		rawData, err := io.ReadAll(c.Request.Body)
		core.PanicIf(err)

		var outputData []byte

		var wg sync.WaitGroup
		for _, relayTarget := range conf.RelayTargetList() {
			wg.Add(1)
			go func(relayTarget string) {
				defer func() {
					wg.Done()
				}()

				fmt.Println("S")

				var relayScript string
				switch conf.EncryptRelay() {
				case true:
					relayScript = fmt.Sprintf(encryptRelayScript, conf.ROOT.Forward.SecurePipe.AESKey, conf.ROOT.Forward.SecurePipe.Token, conf.ROOT.Forward.Targets, customerMsgID)
				default:
					relayScript = fmt.Sprintf(rawRelayScript, conf.ROOT.Forward.Targets, customerMsgID)
				}

				// replay command
				cmd := exec.Command("bash", "-c", relayScript)

				// input from http post body
				cmd.Stdin = bytes.NewReader(rawData)

				// write to output
				var output bytes.Buffer
				cmd.Stdout = &output

				// write stderr to output2
				var output2 bytes.Buffer
				cmd.Stderr = &output2

				if err := cmd.Run(); err != nil {
					fmt.Println("[ERROR]", errors.Join(err,
						fmt.Errorf("stdout: %v", output.String()),
						fmt.Errorf("stderr: %v", output2.String())))
				}

				fmt.Println("D")

				outputData = output.Bytes()

			}(relayTarget)
		}

		wg.Wait()
		c.Header("Content-Type", "application/json")
		_, _ = c.Writer.Write(outputData)
	}))

}

func Run() {
	bind := conf.ROOT.Server.Bind
	log.WithContext(context.Background()).Info("http server started", zap.Any("bind", bind))
	core.PanicIf(eng.Run(bind))
}

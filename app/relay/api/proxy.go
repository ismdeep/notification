package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/ismdeep/notification/app/relay/conf"
	"github.com/ismdeep/notification/pkg/core"
)

func relay(c *gin.Context) {
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

			outputData = output.Bytes()

		}(relayTarget)
	}

	wg.Wait()
	c.Header("Content-Type", "application/json")
	_, _ = c.Writer.Write(outputData)
}

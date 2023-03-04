package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/proxy"

	"github.com/ismdeep/notification/app/server/conf"
	"github.com/ismdeep/notification/app/server/model"
	"github.com/ismdeep/notification/app/server/store"
	"github.com/ismdeep/notification/pkg/core"
)

func send(msg model.Msg) {

	data, err := json.Marshal(gin.H{
		"chat_id": conf.ROOT.Security.TelegramChatID,
		"text":    msg.Content,
	})
	core.PanicIf(err)

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", conf.ROOT.Security.TelegramBotToken),
		bytes.NewReader(data))
	core.PanicIf(err)
	req.Header.Set("Content-Type", "application/json")

	// create a http client
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	// insert socks5 proxy
	if conf.ROOT.Proxy.Socks5 != "" {
		var proxyAuth *proxy.Auth
		if conf.ROOT.Proxy.Socks5Username != "" {
			proxyAuth = &proxy.Auth{
				User:     conf.ROOT.Proxy.Socks5Username,
				Password: conf.ROOT.Proxy.Socks5Password,
			}
		}
		dialer, err := proxy.SOCKS5("tcp", conf.ROOT.Proxy.Socks5, proxyAuth, nil)
		if err != nil {
			log.Fatal(err)
		}
		httpClient.Transport = &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		}
	}

	resp, err := httpClient.Do(req)
	core.PanicIf(err)

	_, err = io.ReadAll(resp.Body)
	core.PanicIf(err)
}

// ConsumerMsgWorker get pending msg from msgs and repost
func ConsumerMsgWorker() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[ERROR] ConsumerMsgWorker() got a panic:", r)
		}
	}()

	lst := store.Msg.PendingMsgList(1024)
	for _, msg := range lst {
		func(msg model.Msg) {
			defer func() {
				if r := recover(); r != nil {
					errMsg := fmt.Sprintf("%+v", r)
					fmt.Println("[ERROR] failed to send msg:", errMsg)
					store.Msg.SetFailed(msg.ID, errMsg)
				}
			}()

			send(msg)
			store.Msg.SetSent(msg.ID)
		}(msg)
	}
}

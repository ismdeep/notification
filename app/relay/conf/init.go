package conf

import (
	"context"
	"errors"
	"strings"

	"github.com/ismdeep/log"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"

	"github.com/ismdeep/notification/pkg/core"
)

// s config model
type s struct {
	Server struct {
		Bind string `env:"SERVER_BIND,default=0.0.0.0:7080"`
	}
	Forward struct {
		Targets    string `env:"FORWARD_TARGETS"`
		SecurePipe struct {
			AESKey string `env:"FORWARD_SECURE_PIPE_AES_KEY"`
			Token  string `env:"FORWARD_SECURE_PIPE_TOKEN"`
		}
	}
}

// ROOT instance
var ROOT s

var targets []string

func EncryptRelay() bool {
	return ROOT.Forward.SecurePipe.AESKey != ""
}

func RelayTargetList() []string {
	return targets
}

func init() {
	core.PanicIf(envconfig.Process(context.Background(), &ROOT))

	core.PanicIf(
		core.IfErr(ROOT.Forward.Targets == "", errors.New("FORWARD_TARGETS is empty. e.g. http://192.168.56.1:7080;http://172.10.0.1:7080")))

	targets = strings.Split(ROOT.Forward.Targets, ";")
	for idx, target := range targets {
		log.WithContext(context.Background()).Info("relay target", zap.Any("idx", idx), zap.Any("target", target))
	}
}

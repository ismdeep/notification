package worker

import (
	"time"

	"github.com/robfig/cron/v3"

	"github.com/ismdeep/notification/pkg/core"
)

// Run workers
func Run() {
	locShangHai, err := time.LoadLocation("Asia/Shanghai")
	core.PanicIf(err)

	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.SkipIfStillRunning(
				cron.DefaultLogger)),
		cron.WithLocation(locShangHai))
	_, err = c.AddFunc("@every 1s", ConsumerMsgWorker)
	core.PanicIf(err)
	c.Run()
}

package main

import (
	"fmt"
	"github.com/go-toast/toast"
	"github.com/lxn/walk"
	"github.com/robfig/cron"
	"github.com/winstonkenny/drinking/windows"
	"log"
	"os"
	"reflect"
	"runtime"
	"time"
)

var Cron *cron.Cron

func main() {
	// 创建定时任务
	if Cron == nil {
		Cron = cron.New()
	}

	_ = Cron.AddFunc("0 0,30 * * * *", func() { Run(DoNotice) })

	Cron.Start()
	fmt.Println("CronJob start.....")

	mw := new(windows.MyMainWindow)
	var err error
	mw.MainWindow, err = walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}
	mw.AddNotifyIcon()
	mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		Cron.Stop()
	})
	mw.Run()
}

func Run(job func() error) {
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		fmt.Printf("%s error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}

func DoNotice() error {
	runningTime, _ := os.Getwd()
	timeNow := time.Now().Format("15:04")
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   "喝水去",
		Message: "现在是[" + timeNow + "],该喝水了呀，亲亲!!",
		Icon:    runningTime + "/img/logo.png", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "刚喝了", ""},
			{"protocol", "好,这就去", ""},
			{"protocol", "等5分钟", ""},
		},
	}
	err := notification.Push()
	return err
}

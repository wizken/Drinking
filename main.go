package main

import (
	"fmt"
	"github.com/go-toast/toast"
	"github.com/lxn/walk"
	"github.com/robfig/cron"
	"os"
	"reflect"
	"runtime"
	"study/my-tools/drink_water_regularly/windows"
	"study/util"
	"time"
)

var Cron *cron.Cron
var Entries []cron.EntryID

func main() {
	// MQTT 1C86E3CE26A36094B4C5F7642DD5D056
	//waitGroup := sync.WaitGroup{}
	mac := util.GetMac()
	fmt.Println(os.Hostname())
	fmt.Println(mac)
	// 创建定时任务
	if Cron == nil {
		Cron = cron.New(cron.WithSeconds())
	}

	entryId, _ := Cron.AddFunc("0 0,30 * * * *", func() { Run(DoNotice) })
	Entries = append(Entries, entryId)

	Cron.Start()
	fmt.Println("CronJob start.....")

	//waitGroup.Add(1)
	//waitGroup.Wait()

	mw := new(windows.MyMainWindow)
	mw.MainWindow, _ = walk.NewMainWindow()
	mw.AddNotifyIcon()
	mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		Cron.Stop()
	})
	mw.Run()
	fmt.Println("CronJob end.....")
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
	timeNow := time.Now().Format("15:04")
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   "喝水去",
		Message: "现在是[" + timeNow + "],该喝水了呀，亲亲!!",
		Icon:    "E:\\Picture\\HeadImage\\favicon_64x64.ico", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "刚喝了", ""},
			{"protocol", "好,这就去", ""},
			{"protocol", "等5分钟", ""},
		},
	}
	err := notification.Push()
	return err
}

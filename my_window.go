package drinking

import (
	"github.com/lxn/walk"
	"log"
)

type MyMainWindow struct {
	*walk.MainWindow
	ni *walk.NotifyIcon
}

func (mw *MyMainWindow) AddNotifyIcon() {
	var err error
	mw.ni, err = walk.NewNotifyIcon(mw)
	if err != nil {
		log.Fatal(err)
	}
	//托盘图标文件
	iconLogo, err := walk.Resources.Icon("./img/logo_16.ico")
	if err != nil {
		log.Fatal(err)
	}
	if err := mw.SetIcon(iconLogo); err != nil {
		log.Fatal(err)
	}
	if err := mw.ni.SetIcon(iconLogo); err != nil {
		log.Fatal(err)
	}
	if err := mw.ni.SetToolTip("定时提醒喝水"); err != nil {
		log.Fatal(err)
	}
	_ = mw.ni.SetVisible(true)
	mw.ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button == walk.LeftButton {
		}
	})

	// 右键退出按钮
	exitAction := walk.NewAction()
	iconExit, _ := walk.Resources.Icon("./img/exit_16.ico")
	_ = exitAction.SetImage(iconExit)
	_ = exitAction.SetText("退出")
	exitAction.Triggered().Attach(func() {
		walk.App().Exit(0)
	})
	// 信息按钮
	infoAction := walk.NewAction()
	iconInfo, _ := walk.Resources.Icon("./img/info_16.ico")
	_ = infoAction.SetImage(iconInfo)
	_ = infoAction.SetText("关于")
	infoAction.Triggered().Attach(func() {
		walk.MsgBox(mw, "提醒喝水软件", "每天8杯水，健康千百年", walk.MsgBoxIconInformation)
		log.Println("关于")
	})
	// 设置按钮
	settingAction := walk.NewAction()
	iconSetting, _ := walk.Resources.Icon("./img/setting_16.ico")
	_ = settingAction.SetImage(iconSetting)
	_ = settingAction.SetText("设置")
	settingAction.Triggered().Attach(func() {
		//_ = windows.CreateSetting(mw.MainWindow)
		log.Println("设置")
	})
	// 显示主页面按钮
	showAction := walk.NewAction()
	iconShow, _ := walk.Resources.Icon("./img/logo_16.ico")
	_ = showAction.SetImage(iconShow)
	_ = showAction.SetText("显示主页面")
	showAction.Triggered().Attach(func() {
		log.Println("显示主页面")
	})

	_ = mw.ni.ContextMenu().Actions().Add(showAction)
	_ = mw.ni.ContextMenu().Actions().Add(settingAction)
	_ = mw.ni.ContextMenu().Actions().Add(infoAction)
	_ = mw.ni.ContextMenu().Actions().Add(exitAction)
	// 使能notify
	if err := mw.ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}
}

package windows

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"time"
)


func CreateSetting(mw *walk.MainWindow) error {
	var inTE, outTE *walk.TextEdit
	iconLogo, err := walk.Resources.Icon("./img/logo_16.ico")
	if err != nil {
		log.Fatal(err)
	}
	myMw := MainWindow{
		AssignTo: &mw, // 如果引用walk中的窗口的话，关闭窗口则会造成程序退出
		Icon: iconLogo,
		Title:   "设置",
		MinSize: Size{300, 200},
		Size:    Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
		},
	}
	btn := PushButton{
		Text: "SCREAM",
		OnClicked: func() {
			mw.Hide()
			time.Sleep(4*time.Second)
			mw.Show()
		},
	}
	myMw.Children = append(myMw.Children, btn)
	if err := myMw.Create(); err != nil {
		return err
	}
	return nil
}


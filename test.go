package main

import (
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
)

func main() {
	app := app.New()
	tbP:=widget.NewTabContainer(
		widget.NewTabItem("Base64加密",widget.NewVBox(widget.NewLabel("Base64加密"))),
		widget.NewTabItem(`Base64解密`,widget.NewVBox(widget.NewLabel("Base64解密"))),
		widget.NewTabItem(`MD5加密`,widget.NewVBox(widget.NewLabel("MD5加密"))),
		widget.NewTabItem("Base32加密",widget.NewVBox(widget.NewLabel("Base32加密"))),
		widget.NewTabItem("Base32解密",widget.NewVBox(widget.NewLabel("Base32解密"))),
		widget.NewTabItem("ROT13",widget.NewVBox(widget.NewLabel("ROT13"))),
		widget.NewTabItem("SHA-1",widget.NewVBox(widget.NewLabel("SHA-1"))),
		)
	tbP.SetTabLocation(widget.TabLocationLeading)
	w := app.NewWindow("Tools")
	w.SetContent(
		widget.NewVBox(
		widget.NewGroup(`输入输出`,widget.NewHBox(widget.NewLabel("输入:"),widget.NewEntry()),widget.NewHBox(widget.NewLabel("输出:"),widget.NewEntry())),
		widget.NewGroup(`算法`,tbP),
		widget.NewGroup(`统计`,
			widget.NewHBox(widget.NewLabel("字符数:"),widget.NewEntry()),
			widget.NewHBox(widget.NewLabel("数字个数:"),widget.NewEntry()),
			widget.NewHBox(widget.NewLabel("大写字母个数:"),widget.NewEntry()),
			widget.NewHBox(widget.NewLabel("小写字母个数:"),widget.NewEntry()),
		),
	))
	app.Settings().SetTheme(theme.LightTheme())
	w.ShowAndRun()
}

type DateLable struct{

}
package main

import (
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"strconv"
)

	var inEntry=&widget.Entry{OnChanged:statistics}
	var lenLabel=&widget.Label{}
	var numLabel=&widget.Label{}
	var upLabel=&widget.Label{}
	var loLabel=&widget.Label{}
	var outEntry=&widget.Entry{}
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
		widget.NewGroup(`输入输出`,widget.NewHBox(widget.NewLabel("输入:"),inEntry),widget.NewHBox(widget.NewLabel("输出:"),outEntry)),
		widget.NewGroup(`算法`,tbP),
		widget.NewGroup(`统计`,
			widget.NewHBox(widget.NewLabel("字符数:"),lenLabel),
			widget.NewHBox(widget.NewLabel("数字个数:"),numLabel),
			widget.NewHBox(widget.NewLabel("大写字母个数:"),upLabel),
			widget.NewHBox(widget.NewLabel("小写字母个数:"),loLabel),
		),
	))
	app.Settings().SetTheme(theme.LightTheme())
	w.ShowAndRun()
}

type DateLable struct{

}


func statistics(str string){
	lenLabel.SetText(strconv.Itoa(len(str)+1))
	numLabel.SetText(strconv.Itoa(len(str)+2))
	upLabel.SetText(strconv.Itoa(len(str)+3))
	loLabel.SetText(strconv.Itoa(len(str)*4))
	if(len(str)<300){
		outEntry.SetText(str)
	}
}
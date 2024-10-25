package main

import (
	"fyne-exercise-tracker/ui" // 引入我們自定義的頁面

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	// 建立應用程式
	myApp := app.New()
	myWindow := myApp.NewWindow("Exercise Tracker")

	// 創建不同頁面
	homePage := ui.LoadHomePage(myWindow)
	weightPage := ui.LoadWeightPage()
	recordPage := ui.LoadRecordPage()
	statsPage := ui.LoadStatsPage(myWindow)

	// 建立選單和頁面導航
	tabs := container.NewAppTabs(
		container.NewTabItem("首頁", homePage),
		container.NewTabItem("體重輸入", weightPage),
		container.NewTabItem("紀錄", recordPage),
		container.NewTabItem("統計", statsPage.CanvasObject()),
	)

	// 當切換到統計頁面時，觸發刷新
	tabs.OnSelected = func(tab *container.TabItem) {
		if tab.Text == "統計" {
			// 確保刷新統計頁面
			statsPage.Refresh()
		}
	}

	// 設定視窗內容
	myWindow.SetContent(tabs)

	// 設定視窗大小
	myWindow.Resize(fyne.NewSize(600, 400))

	// 顯示視窗並運行應用程式
	myWindow.ShowAndRun()
}

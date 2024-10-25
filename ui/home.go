package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// LoadHomePage 建立首頁
func LoadHomePage(win fyne.Window) fyne.CanvasObject {
	// 首頁UI元件
	welcomeText := widget.NewLabel("歡迎使用運動紀錄應用程式")
	exitButton := widget.NewButton("關閉應用程式", func() {
		win.Close()
	})

	// 使用 VBox 排列元件
	content := container.NewVBox(
		welcomeText,
		exitButton,
	)

	return content
}

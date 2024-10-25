package ui

import (
	"fmt"
	"fyne-exercise-tracker/data"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// LoadWeightPage 建立體重輸入頁面
func LoadWeightPage() fyne.CanvasObject {
	// 初始化資料庫
	db := data.InitDatabase()

	// 體重輸入
	weightEntry := widget.NewEntry()
	weightEntry.SetPlaceHolder("輸入今日體重")

	// 當前日期
	today := time.Now().Format("2006-01-02")

	// 檢查當天是否已有體重紀錄
	existingWeight, err := data.GetWeightByDate(db, today)
	if err == nil && existingWeight > 0 {
		weightEntry.SetText(fmt.Sprintf("%.2f", existingWeight))
	}

	// 儲存按鈕
	saveButton := widget.NewButton("儲存體重", func() {
		weight, _ := strconv.ParseFloat(weightEntry.Text, 64)
		data.InsertWeightRecord(db, today, weight) // 儲存到SQLite
		weightEntry.SetText("")                    // 清空欄位
		fmt.Println("體重已儲存")
	})

	return container.NewVBox(
		widget.NewLabel("今日體重輸入"),
		weightEntry,
		saveButton,
	)
}

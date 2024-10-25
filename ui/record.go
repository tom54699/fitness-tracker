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

// LoadRecordPage 建立運動紀錄頁面
func LoadRecordPage() fyne.CanvasObject {
	// 初始化資料庫
	db := data.InitDatabase()

	// 運動項目輸入
	exerciseEntry := widget.NewEntry()
	exerciseEntry.SetPlaceHolder("運動項目")

	// 組次數輸入
	repsEntry := widget.NewEntry()
	repsEntry.SetPlaceHolder("每組次數")

	// 組次數的單位選擇（次數或秒數）
	unitOptions := widget.NewSelect([]string{"次", "秒"}, nil)
	unitOptions.SetSelected("次") // 預設為次

	// 組數輸入
	setsEntry := widget.NewEntry()
	setsEntry.SetPlaceHolder("幾組")

	// 消耗時間輸入
	timeEntry := widget.NewEntry()
	timeEntry.SetPlaceHolder("消耗時間 (分鐘)")

	// 預估卡路里輸入
	caloriesEntry := widget.NewEntry()
	caloriesEntry.SetPlaceHolder("預估卡路里")

	// 備註輸入
	remarksEntry := widget.NewEntry()
	remarksEntry.SetPlaceHolder("備註")

	// 儲存按鈕
	saveButton := widget.NewButton("儲存紀錄", func() {
		reps, _ := strconv.Atoi(repsEntry.Text)
		sets, _ := strconv.Atoi(setsEntry.Text)
		timeSpent, _ := strconv.Atoi(timeEntry.Text)
		calories, _ := strconv.ParseFloat(caloriesEntry.Text, 64)
		formattedDate := time.Now().Format("2006-01-02")

		unit := unitOptions.Selected

		record := data.ExerciseRecord{
			Date:           formattedDate,
			Exercise:       exerciseEntry.Text,
			RepsPerSet:     reps,
			Sets:           sets,
			TimeSpent:      timeSpent,
			CaloriesBurned: calories,
			Remarks:        remarksEntry.Text,
			Unit:           unit,
		}

		data.InsertRecord(db, record) // 儲存到SQLite

		// 清空填寫的資料
		exerciseEntry.SetText("")
		repsEntry.SetText("")
		setsEntry.SetText("")
		timeEntry.SetText("")
		caloriesEntry.SetText("")
		remarksEntry.SetText("")

		fmt.Println("紀錄成功，當前時間:", time.Now())
	})

	// 使用 Form 布局，讓介面更整齊
	form := container.NewVBox(
		widget.NewLabel("請輸入運動紀錄"),
		exerciseEntry,
		container.NewGridWithColumns(2, repsEntry, unitOptions),
		setsEntry,
		timeEntry,
		caloriesEntry,
		remarksEntry,
		saveButton,
	)

	return container.NewVBox(form)
}

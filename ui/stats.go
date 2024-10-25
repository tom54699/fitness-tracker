package ui

import (
	"database/sql"
	"fmt"
	"fyne-exercise-tracker/data"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// StatsPage 結構，包含頁面的內容及刷新邏輯
type StatsPage struct {
	content    *fyne.Container
	statsLabel *widget.Label
	db         *sql.DB
}

// LoadStatsPage 建立統計頁面
func LoadStatsPage(win fyne.Window) *StatsPage {
	// 初始化資料庫
	db := data.InitDatabase()

	// 查詢並顯示紀錄
	statsLabel := widget.NewLabel("")
	page := &StatsPage{
		content:    container.NewVBox(),
		statsLabel: statsLabel,
		db:         db,
	}

	// 統計一周或一個月按鈕
	weekButton := widget.NewButton("統計一周", func() {
		startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
		endDate := time.Now().Format("2006-01-02")
		page.calculateStats(startDate, endDate)
	})

	monthButton := widget.NewButton("統計一月", func() {
		startDate := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
		endDate := time.Now().Format("2006-01-02")
		page.calculateStats(startDate, endDate)
	})

	// 初始化頁面內容
	page.content.Add(widget.NewLabel("運動統計"))
	page.content.Add(widget.NewSeparator())
	page.content.Add(statsLabel)
	page.content.Add(container.NewHBox(weekButton, monthButton))
	page.content.Add(widget.NewSeparator())
	page.Refresh() // 頁面初次加載時刷新

	return page
}

// Refresh 方法，刷新統計數據
func (p *StatsPage) Refresh() {
	records, err := data.GetAllRecords(p.db)
	if err != nil {
		fmt.Println("查詢紀錄失敗:", err)
	}

	var recordList string
	for _, record := range records {
		recordList += fmt.Sprintf("日期: %s, 運動: %s, 組數: %d, 次數: %d %s, 卡路里: %.2f\n",
			record.Date, record.Exercise, record.Sets, record.RepsPerSet, record.Unit, record.CaloriesBurned)
	}

	p.statsLabel.SetText(recordList)
}

// calculateStats 方法，統計某段時間內的運動資料和體重變化
func (p *StatsPage) calculateStats(startDate, endDate string) {
	// 查詢運動紀錄
	records, err := data.GetRecordsByDateRange(p.db, startDate, endDate)
	if err != nil {
		fmt.Println("統計運動紀錄失敗:", err)
		return
	}

	var totalTime int
	var exerciseStats = make(map[string]int)

	for _, record := range records {
		totalTime += record.TimeSpent
		exerciseStats[record.Exercise] += record.RepsPerSet * record.Sets
	}

	// 查詢體重紀錄
	weights, err := data.GetWeightRecordsByDateRange(p.db, startDate, endDate)
	if err != nil {
		fmt.Println("統計體重紀錄失敗:", err)
		return
	}

	// 計算體重變化
	var weightChange string
	if len(weights) > 0 {
		startWeight := weights[0]
		endWeight := weights[len(weights)-1]
		weightChange = fmt.Sprintf("體重變化: %.2f 公斤 -> %.2f 公斤\n", startWeight, endWeight)
	} else {
		weightChange = "無體重紀錄"
	}

	// 顯示統計結果
	var statsOutput string
	statsOutput += fmt.Sprintf("統計期間: %s 到 %s\n", startDate, endDate)
	statsOutput += fmt.Sprintf("總運動時間: %d 分鐘\n", totalTime)
	statsOutput += weightChange

	for exercise, totalReps := range exerciseStats {
		unit := "次"
		if exercise == "秒" {
			unit = "秒"
		}
		statsOutput += fmt.Sprintf("運動項目: %s, 總次數: %d %s\n", exercise, totalReps, unit)
	}

	p.statsLabel.SetText(statsOutput)
}

// CanvasObject 方法，返回頁面的可視化內容
func (p *StatsPage) CanvasObject() fyne.CanvasObject {
	return p.content
}

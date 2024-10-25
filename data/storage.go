package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// ExerciseRecord 定義資料庫表格的結構
type ExerciseRecord struct {
	ID             int
	Date           string
	Exercise       string
	RepsPerSet     int
	Sets           int
	TimeSpent      int
	CaloriesBurned float64
	Remarks        string
	Unit           string
}

// InitDatabase 初始化資料庫，並且建立兩個表格：exercise_records 和 weight_records
func InitDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./exercise_tracker.db")
	if err != nil {
		log.Fatal(err)
	}

	// 建立運動紀錄表格
	createExerciseTable := `
	CREATE TABLE IF NOT EXISTS exercise_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT,
		exercise TEXT,
		reps_per_set INTEGER,
		sets INTEGER,
		time_spent INTEGER,
		calories_burned REAL,
		remarks TEXT,
		unit TEXT
	);`

	_, err = db.Exec(createExerciseTable)
	if err != nil {
		log.Fatal(err)
	}

	// 建立體重紀錄表格
	createWeightTable := `
	CREATE TABLE IF NOT EXISTS weight_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT,
		weight REAL
	);`

	_, err = db.Exec(createWeightTable)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// InsertRecord 插入新的運動紀錄
func InsertRecord(db *sql.DB, record ExerciseRecord) {
	insertSQL := `INSERT INTO exercise_records(date, exercise, reps_per_set, sets, time_spent, calories_burned, remarks, unit) 
	              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(insertSQL, record.Date, record.Exercise, record.RepsPerSet, record.Sets, record.TimeSpent, record.CaloriesBurned, record.Remarks, record.Unit)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("紀錄已成功儲存！")
	}
}

// InsertWeightRecord 插入新的體重紀錄
func InsertWeightRecord(db *sql.DB, date string, weight float64) {
	insertSQL := `INSERT INTO weight_records(date, weight) VALUES (?, ?)`
	_, err := db.Exec(insertSQL, date, weight)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("體重紀錄已成功儲存！")
	}
}

// GetWeightByDate 查詢指定日期的體重
func GetWeightByDate(db *sql.DB, date string) (float64, error) {
	var weight float64
	query := `SELECT weight FROM weight_records WHERE date = ?`
	err := db.QueryRow(query, date).Scan(&weight)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // 表示該日期尚無體重紀錄
		}
		return 0, err
	}
	return weight, nil
}

// GetAllRecords 查詢所有運動紀錄
func GetAllRecords(db *sql.DB) ([]ExerciseRecord, error) {
	rows, err := db.Query("SELECT * FROM exercise_records")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []ExerciseRecord
	for rows.Next() {
		var record ExerciseRecord
		err := rows.Scan(&record.ID, &record.Date, &record.Exercise, &record.RepsPerSet, &record.Sets, &record.TimeSpent, &record.CaloriesBurned, &record.Remarks, &record.Unit)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func GetRecordsByDateRange(db *sql.DB, startDate, endDate string) ([]ExerciseRecord, error) {
	query := `SELECT * FROM exercise_records WHERE date BETWEEN ? AND ?`
	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []ExerciseRecord
	for rows.Next() {
		var record ExerciseRecord
		err := rows.Scan(&record.ID, &record.Date, &record.Exercise, &record.RepsPerSet, &record.Sets, &record.TimeSpent, &record.CaloriesBurned, &record.Remarks, &record.Unit)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

// GetWeightRecordsByDateRange 查詢特定日期範圍內的體重紀錄
func GetWeightRecordsByDateRange(db *sql.DB, startDate, endDate string) ([]float64, error) {
	query := `SELECT weight FROM weight_records WHERE date BETWEEN ? AND ?`
	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weights []float64
	for rows.Next() {
		var weight float64
		err := rows.Scan(&weight)
		if err != nil {
			return nil, err
		}
		weights = append(weights, weight)
	}
	return weights, nil
}

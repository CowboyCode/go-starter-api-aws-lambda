package mariadb

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"runtime"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	dateFormat = "2006-01-02"
)

const (
	timeFormatAU = "02/01/2006 14:04:05"
	dateFormatAU = "02/01/2006"
)

var (
	Con *sql.DB
)

// SetUpDatabase for connection
func SetUpDatabase() error {

	var err error
	Con, err = sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_USER_PWD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE")+"")
	if err != nil {
		println(err.Error())
	}

	Con.SetConnMaxLifetime(time.Minute * 3)
	Con.SetMaxOpenConns(10)
	Con.SetMaxIdleConns(10)

	return nil
}

func CloseConnection() {
	defer Con.Close()
}

// ErrorLog silent error and event logger to monitor database abnormalities
func ErrorLog(logType string, logMessage string, filename string, line string, pack string) {

	t := time.Now()
	var nowTime = t.Format("2006-01-02 15:04:05")

	var query = `INSERT INTO actionlog 
                        (actiontype, message, filename, line, logtime, location) 
                    VALUES 
						            (?, ?, ?, ?, ?, ?)
	`
	// database handle is available her
	Con.Exec(query, logType, logMessage, filename, line, nowTime, pack)

}

// GetHoursBetweenDates exported
func GetHoursBetweenDates(start string, end string) int {

	starter, err := time.Parse(timeFormat, start)
	if err != nil {
		return 0
	}
	ender, err := time.Parse(timeFormat, end)
	if err != nil {
		return 0
	}

	duration := ender.Sub(starter).Hours()
	return int(duration)
}

// GetNowMysql exported
func GetNowMysql(format string) (string, error) {

	if format != "day" && format != "" {
		return "day", errors.New("invalid parameter")
	}

	Timezone, _ := time.LoadLocation(os.Getenv("TimeZone"))

	if format == "day" {
		return time.Now().In(Timezone).Format(dateFormat), nil
	}

	return time.Now().In(Timezone).Format(timeFormat), nil
}

// GetTodayPlusDays exported
func GetTodayPlusDays(plusdays int, format string) string {
	Timezone, _ := time.LoadLocation(os.Getenv("TimeZone"))
	now := time.Now().In(Timezone).Format(timeFormat)

	temp, err := time.ParseInLocation(timeFormat, now, Timezone)
	if err != nil {
		return err.Error()
	}

	temp = temp.AddDate(0, 0, plusdays)

	if format == "day" {
		return string(temp.Format(dateFormat))
	}
	return string(temp.Format(timeFormat))
}

// GetDateMinus exported
func GetDateMinus(dateString string, minusDays int, format string) string {
	Timezone, _ := time.LoadLocation(os.Getenv("TimeZone"))
	temp, err := time.ParseInLocation(dateFormat, dateString, Timezone)
	if err != nil {
		return err.Error()
	}

	temp = temp.AddDate(0, 0, minusDays)

	if format == "day" {
		return string(temp.Format(dateFormat))
	}
	return string(temp.Format(timeFormat))
}

// GetFileInfo for error log
func GetFileInfo(typus string) string {
	_, file, line, ok := runtime.Caller(1)
	var s string
	if ok {
		if typus == "filename" {
			s = fmt.Sprintf("%s", file)
		}

		if typus == "fileLine" {
			s = fmt.Sprintf("%d", line)
		}

	} else {
		s = ""
	}
	return s
}

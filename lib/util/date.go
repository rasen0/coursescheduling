package util

import (
	"coursesheduling/common"
	"time"
)

func DurationMonth(timer time.Time) (month1,month2 time.Time) {
	month1 = timer.Add(-time.Hour*time.Duration(24* (timer.Day()-1)))
	month2 = month1.AddDate(0,1,0)
	return
}

// GetYearMonthToDay 查询指定年份指定月份有多少天
// @params year int 指定年份
// @params month int 指定月份
func GetYearMonthToDay(year int, month int) int {
	// 有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] == true {
		return 31
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}

func IteratorDay(month time.Time) (calendar string) {
	month0, _ := DurationMonth(month)
	days := GetYearMonthToDay(month.Year(),int(month.Month()))
	for idx:= 0; idx < days ; idx++{
		date := time.Date(month.Year(),month.Month(),month0.Day()+idx,month.Hour(),month.Minute(),month.Second(),month.Nanosecond(),month.Location())
		calendar = date.Format(common.CalendarFormat)
	}
	return
}
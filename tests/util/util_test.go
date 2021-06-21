package util_test

import (
	"coursesheduling/common"
	"coursesheduling/lib/util"
	"fmt"
	"testing"
	"time"
)

func TestSetValue(t *testing.T) {
	now := time.Now()
	fmt.Println("now :",now,"day:",now.Day())
	now1 := now.Add(-time.Hour*time.Duration(24* (now.Day()-1)))
	fmt.Println("now1:",now1)
	now2 := now1.AddDate(0,1,0)
	fmt.Println("now2:",now2)
}

func TestMonthDay(t *testing.T) {
	now := time.Now()
	month, _ := util.DurationMonth(now)
	//oldDay := "01"
	days := util.GetYearMonthToDay(now.Year(),int(now.Month()))
	for idx:= 0; idx < days ; idx++{
		date := time.Date(month.Year(),month.Month(),month.Day()+idx,month.Hour(),month.Minute(),month.Second(),month.Nanosecond(),month.Location())
		calendar := date.Format(common.CalendarFormat)
		fmt.Println("cl:",calendar)
	}
}
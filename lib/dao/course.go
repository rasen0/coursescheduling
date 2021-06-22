package dao

import (
	"coursesheduling/common"
	"coursesheduling/database"
	"coursesheduling/lib/util"
	"coursesheduling/model"
	"time"
)

// SummaryCourseDate 给不同课程包装统一类型
func SummaryCourseDate(ctype int,month time.Time) (list []model.Course) {
	list = make([]model.Course,0)
	switch ctype {
	case model.SingleLesson:
		courses := database.GetSingleCourseByMonth(month)
		for _, c := range courses{
			list = append(list, c)
		}
	case model.TrialLesson:
		courses := database.GetTrialCourseByMonth(month)
		for _, c := range courses{
			list = append(list, c)
		}
	case model.CommonLesson:
		courses := database.GetCommonCourseByMonth(month)
		for _, c := range courses{
			list = append(list, c)
		}
	default:
	}
	return
}

func GetCourseMonth(ctype int, month time.Time) (coursesTable map[string]map[string][]model.Course) {
	coursesTable = make(map[string]map[string][]model.Course)
	courses := SummaryCourseDate(ctype,month)
	for _, c := range courses{
		if sub,ok := coursesTable[c.Calendar()]; ok {
			if _,ok := sub[c.StartClock()];ok{
				continue
			}else {
				sub[c.StartClock()] = []model.Course{
					c,
				}
			}
		}else{
			pTime := make(map[string][]model.Course)
			pTime[c.StartClock()] = []model.Course{
				c,
			}
			coursesTable[c.Calendar()] = pTime
		}
	}
    return
}

func GetCourseTable(ctype int, myMonth time.Time) (courseMonth map[string]map[string][]model.Course,courseTable []model.CourseOfDay) {
	courseMonth = GetCourseMonth(ctype, myMonth)
	month, _ := util.DurationMonth(myMonth)
	days := util.GetYearMonthToDay(myMonth.Year(),int(myMonth.Month()))
	courseTable = make([]model.CourseOfDay,0,days)
	for idx:= 0; idx < days ; idx++{
		date := time.Date(month.Year(),month.Month(),month.Day()+idx,month.Hour(),month.Minute(),month.Second(),month.Nanosecond(),month.Location())
		calendar := date.Format(common.CalendarFormat)
		cod := model.CourseOfDay{
			Calendar:calendar,
		}

		if cal, ok := courseMonth[calendar]; ok {
			for t,i := range common.TimeClock{
				if list,ok2 := cal[t]; ok2 {
					cod.SetDuringCourse(i,list)
				}
			}
		}
		//courseTable[idx] = cod
		courseTable = append(courseTable,cod)
	}
	return
}

func GetCourseTable2(ctype int, myMonth time.Time) (courseMonth map[string]map[string][]model.Course,courseTable []model.CourseOfDay2) {
	courseMonth = GetCourseMonth(ctype, myMonth)
	month, _ := util.DurationMonth(myMonth)
	days := util.GetYearMonthToDay(myMonth.Year(),int(myMonth.Month()))
	courseTable = make([]model.CourseOfDay2,0,days)
	for idx:= 0; idx < days ; idx++{
		date := time.Date(month.Year(),month.Month(),month.Day()+idx,month.Hour(),month.Minute(),month.Second(),month.Nanosecond(),month.Location())
		calendar := date.Format(common.CalendarFormat)
		cod := model.CourseOfDay2{
			Calendar:calendar,
		}

		if cal, ok := courseMonth[calendar]; ok {
			for t,i := range common.TimeClock{
				if _,ok2 := cal[t]; ok2 {
					cod.SetFlagByDuring(i)
				}
			}
		}
		//courseTable[idx] = cod
		courseTable = append(courseTable,cod)
	}
	return
}

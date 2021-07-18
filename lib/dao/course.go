package dao

import (
	"coursesheduling/common"
	"coursesheduling/database"
	"coursesheduling/lib/util"
	"coursesheduling/model"
	"time"
)

// AddCommonCourses 添加一对多课程
func AddCommonCourses(courses []model.CommonCourse) {
	database.AddCommonCourses(courses)
}

// AddTrialCourses 添加试听课
func AddTrialCourses(courses []model.TrialCourse) {
	database.AddTrialCourses(courses)
}

// AddSingleCourses 添加一对一课程
func AddSingleCourses(courses []model.SingleCourse) {
	database.AddSingleCourses(courses)
}

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

// SummaryCourseDateCond 给不同课程包装统一类型
func SummaryCourseDateCond(data model.QueryData) (list []model.Course) {
	list = make([]model.Course,0)
	if data.QueryType == 2 {
		courses := database.GetCourseSingleByMonth(data)
		for _, c := range courses{
			list = append(list, c)
		}
		return
	}

	courses1 := database.GetCourseCommonByMonth(data)
	for _, c := range courses1{
		list = append(list, c)
	}
	courses2 := database.GetCourseTrialByMonth(data)
	for _, c := range courses2{
		list = append(list, c)
	}
	return
}

func GetCourseMonth(month time.Time) (coursesTable map[string]map[string][]model.Course) {
	coursesTable = make(map[string]map[string][]model.Course)
	courses1 := SummaryCourseDate(model.SingleLesson,month)
	for _, c := range courses1{
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

	courses2 := SummaryCourseDate(model.TrialLesson,month)
	for _, c := range courses2{
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
	courses3 := SummaryCourseDate(model.CommonLesson,month)
	for _, c := range courses3{
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

func GetTypeCourseMonth(ctype int, month time.Time) (coursesTable map[string]map[string][]model.Course) {
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

func GetCourseMonthCond(data model.QueryData) (coursesTable map[string]map[string][]model.Course) {
	coursesTable = make(map[string]map[string][]model.Course)
	courses := SummaryCourseDateCond(data)
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

func GetTypeCourseTable(ctype int, myMonth time.Time) (courseMonth map[string]map[string][]model.Course,courseTable []model.CourseOfDay) {
	courseMonth = GetTypeCourseMonth(ctype, myMonth)
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

func GetCourseTable(myMonth time.Time) (courseMonth map[string]map[string][]model.Course,courseTable []model.CourseOfDay) {
	courseMonth = GetCourseMonth(myMonth)
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
		courseTable = append(courseTable,cod)
	}
	return
}

func GetCourseTableByCond(queryData model.QueryData) (courseMonth map[string]map[string][]model.Course,courseTable []model.CourseOfDay) {
	courseMonth = GetCourseMonthCond(queryData)
	myMonth,_ := time.Parse(time.RFC3339,queryData.CourseDate)
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
		courseTable = append(courseTable,cod)
	}
	return
}

func GetCourseTable2(ctype int, myMonth time.Time) (courseMonth map[string]map[string][]model.Course,courseTable []model.CourseOfDay2) {
	courseMonth = GetTypeCourseMonth(ctype, myMonth)
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

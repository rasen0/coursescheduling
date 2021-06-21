package dao

import (
	"coursesheduling/database"
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

func GetCourseTable(ctype int, month time.Time) (courseTable []model.CourseOfDay) {
	courseMonth := GetCourseMonth(ctype, month)

}
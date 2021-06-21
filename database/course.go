package database

import (
	"coursesheduling/lib/util"
	"coursesheduling/model"
	"time"
)

// GetCommonCourseByMonth 一对多普通课程班
func GetCommonCourseByMonth(month time.Time) (courses []model.CommonCourse) {
	month1, month2 := util.DurationMonth(month)
	appDB.Where("course_date > ? and course_date < ?",month1.Format("2006-01-02"),month2.Format("2006-01-02")).
		Find(&courses)
	return
}

func InsertCommonCourse()  {
	
}

// GetSingleCourseByMonth 一对一课程班
func GetSingleCourseByMonth(month time.Time) (courses []model.SingleCourse) {
	month1, month2 := util.DurationMonth(month)
	appDB.Where("course_date > ? and course_date < ?",month1.Format("2006-01-02"),month2.Format("2006-01-02")).
		Find(&courses)
	return
}

// GetTrialCourseByMonth 试听课程班
func GetTrialCourseByMonth(month time.Time) (courses []model.TrialCourse) {
	month1, month2 := util.DurationMonth(month)
	appDB.Where("course_date > ? and course_date < ?",month1.Format("2006-01-02"),month2.Format("2006-01-02")).
		Find(&courses)
	return
}

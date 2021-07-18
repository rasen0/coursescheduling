package database

import (
	"coursesheduling/lib/log"
	"coursesheduling/lib/util"
	"coursesheduling/model"
	"strings"
	"time"
)

// GetSingleCourseByMonth 一对一课程班
func GetSingleCourseByMonth(month time.Time) (courses []model.SingleCourse) {
	month1, month2 := util.DurationMonth(month)
	appDB.Where("course_date >= ? and course_date <= ?",month1.Format("2006-01-02"),month2.Format("2006-01-02")).
		Find(&courses)
	return
}

// GetCourseSingleByMonth
func GetCourseSingleByMonth(data model.QueryData) (courses []model.SingleCourse) {
	month, err := time.Parse(time.RFC3339, data.CourseDate)
	if err != nil{
		log.Print("parse time",err)
		return
	}
	subSql := strings.Builder{}
	switch data.QueryType {
	case 1:
		subSql.WriteString(" teacher_id = " + data.TeacherId)
	case 2:
		subSql.WriteString(" student_id = " + data.StudentID)
	case 3:
		subSql.WriteString(" student_group_name = " + data.StudentID)
	}
	subSql.WriteString(" course_date >= ? and course_date <= ?")
	month1, month2 := util.DurationMonth(month)
	appDB.Where(subSql.String(),month1.Format("2006-01-02"),month2.Format("2006-01-02")).
		Find(&courses)
	return
}

func AddSingleCourses(courses []model.SingleCourse) {
	appDB.Create(&courses)
}

// GetTrialCourseByMonth 试听课程班
func GetTrialCourseByMonth(month time.Time) (courses []model.TrialCourse) {
	month1, month2 := util.DurationMonth(month)
	appDB.Where("course_date >= ? and course_date <= ?",month1.Format("2006-01-02"),month2.Format("2006-01-02")).
		Find(&courses)
	return
}


// GetCourseTrialByMonth
func GetCourseTrialByMonth(data model.QueryData) (courses []model.TrialCourse) {
	month, err := time.Parse(time.RFC3339, data.CourseDate)
	if err != nil{
		log.Print("parse time",err)
		return
	}
	subSql := strings.Builder{}
	switch data.QueryType {
	case 1:
		subSql.WriteString(" teacher_id = " + data.TeacherId)
	case 2:
		subSql.WriteString(" student_id = " + data.StudentID)
	case 3:
		subSql.WriteString(" student_group_name = " + data.StudentID)
	}
	subSql.WriteString(" course_date >= ? and course_date <= ?")
	month1, month2 := util.DurationMonth(month)
	appDB.Where(subSql.String(),month1.Format("2006-01-02"),month2.Format("2006-01-02")).
		Find(&courses)
	return
}

func AddTrialCourses(courses []model.TrialCourse) {
	appDB.Create(&courses)
}


// GetCommonCourseByMonth 一对多普通课程班
func GetCommonCourseByMonth(month time.Time) (courses []model.CommonCourse) {
	month1, month2 := util.DurationMonth(month)
	appDB.Where("course_date >= ? and course_date <= ?",month1.Format("2006-01-02"),month2.Format("2006-01-02")).
		Find(&courses)
	return
}

// GetCourseCommonByMonth
func GetCourseCommonByMonth(data model.QueryData) (courses []model.CommonCourse) {
	month, err := time.Parse(time.RFC3339, data.CourseDate)
	if err != nil{
		log.Print("parse time",err)
		return
	}
	subSql := strings.Builder{}
	switch data.QueryType {
	case 1:
		subSql.WriteString(" teacher_id = " + data.TeacherId)
	case 2:
		subSql.WriteString(" student_id = " + data.StudentID)
	case 3:
		subSql.WriteString(" student_group_name = " + data.StudentID)
	}
	subSql.WriteString(" course_date > ? and course_date < ?")
	month1, month2 := util.DurationMonth(month)
	appDB.Where(subSql.String(),month1.Format("2006-01-02"),month2.Format("2006-01-02")).
		Find(&courses)
	return
}

func AddCommonCourses(courses []model.CommonCourse) {
	appDB.Create(&courses)
}

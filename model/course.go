package model

import (
	"coursesheduling/common"
	"time"
)

const (
	SingleLesson = 1
	TrialLesson = 2
	CommonLesson = 3
)

type Course interface {
	Type() int
	Calendar() string
	StartClock() string
}

type CommonCourse struct {
	TeacherName string `json:"teacher_name" gorm:"type:varchar(20);not null"`
	TeacherID string `json:"teacher_id" gorm:"type:varchar(20);not null"`
	StudentGroupName string  `json:"student_group_name" gorm:"type:varchar(50);not null"`
	CoursePlanName string `json:"course_plan_name" gorm:"type:varchar(50);not null"`
	CurriculumName string `json:"curriculum_name" gorm:"type:varchar(20);not null"`
	CourseDate string `json:"course_date" gorm:"type:date;not null"`
	StartTime string `json:"start_time" gorm:"type:varchar(20);not null"`
	EndTime string `json:"end_time" gorm:"type:varchar(20);not null"`
	ClassroomName string `json:"classroom_name" gorm:"type:varchar(20);not null"`
	Describe string `json:"describe" gorm:"type:varchar(50);"`
	UpdateTime time.Time `json:"update_time" gorm:"type:timestamp;not null"`
}

func (c CommonCourse) Type() int {
	return CommonLesson
}

func (c CommonCourse) Calendar() string {
	tt,_:=time.Parse("2006-01-02T15:04:05Z07:00",c.CourseDate)
	return tt.Format(common.CalendarFormat)
}

func (c CommonCourse) StartClock() string {
	return c.StartTime
}

type TrialCourse struct {
	TeacherName string `json:"teacher_name" gorm:"type:varchar(20);not null"`
	TeacherID string `json:"teacher_id" gorm:"type:varchar(20);not null"`
	StudentGroupName string  `json:"student_group_name" gorm:"type:varchar(50);not null"`
	CoursePlanName string `json:"course_plan_name" gorm:"type:varchar(50);not null"`
	CurriculumName string `json:"curriculum_name" gorm:"type:varchar(20);not null"`
	CourseDate string `json:"course_date" gorm:"type:date;not null"`
	StartTime string `json:"start_time" gorm:"type:varchar(20);not null"`
	EndTime string `json:"end_time" gorm:"type:varchar(20);not null"`
	ClassroomName string `json:"classroom_name" gorm:"not null"`
	Describe string `json:"describe" gorm:"type:varchar(50);"`
	UpdateTime time.Time `json:"update_time" gorm:"not null"`
}

func (t TrialCourse) Type() int {
	return TrialLesson
}

func (t TrialCourse) Calendar() string {
	tt,_:=time.Parse("2006-01-02T15:04:05Z07:00",t.CourseDate)
	return tt.Format(common.CalendarFormat)
}

func (t TrialCourse) StartClock() string {
	tt,_:=time.Parse("2006-01-02T15:04:05Z07:00",t.CourseDate)
	return tt.Format(common.ClockFormat)
}

type SingleCourse struct {
	TeacherName string `json:"teacher_name" gorm:"type:varchar(20);not null"`
	TeacherID string `json:"teacher_id" gorm:"type:varchar(20);not null"`
	StudentName string  `json:"student_name" gorm:"type:varchar(20);not null"`
	StudentID string  `json:"student_id" gorm:"type:varchar(20);not null"`
	CoursePlanName string `json:"course_plan_name" gorm:"type:varchar(20);not null"`
	CurriculumName string `json:"curriculum_name" gorm:"type:varchar(20);not null"`
	CourseDate string `json:"course_date" gorm:"type:date;not null"`
	StartTime string `json:"start_time" gorm:"not null"`
	EndTime string `json:"end_time" gorm:"not null"`
	ClassroomName string `json:"classroom_name" gorm:"type:varchar(20);not null"`
	Describe string `json:"describe" gorm:"type:varchar(50);"`
	UpdateTime time.Time `json:"update_time" gorm:"not null"`
}

func (s SingleCourse) Type() int {
	return SingleLesson
}

func (s SingleCourse) Calendar() string {
	tt,_:=time.Parse("2006-01-02T15:04:05Z07:00",s.CourseDate)
	return tt.Format(common.CalendarFormat)
}

func (s SingleCourse) StartClock() string {
	tt,_:=time.Parse("2006-01-02T15:04:05Z07:00",s.CourseDate)
	return tt.Format(common.ClockFormat)
}

type CourseOfDay2 struct {
	Calendar string `json:"calendar"`
	During08 string `json:"during08"`
	During10 string `json:"during10"`
	During12 string `json:"during12"`
	During14 string `json:"during14"`
	During16 string `json:"during16"`
	During18 string `json:"during18"`
	During20 string `json:"during20"`
	During22 string `json:"during22"`
}

func (cod *CourseOfDay2) SetFlagByDuring(num int) {
	switch num {
	case 8:
		cod.During08 = common.CourseFlag
	case 10:
		cod.During08 = common.CourseFlag
	case 12:
		cod.During08 = common.CourseFlag
	case 14:
		cod.During08 = common.CourseFlag
	case 16:
		cod.During08 = common.CourseFlag
	case 18:
		cod.During08 = common.CourseFlag
	case 20:
		cod.During08 = common.CourseFlag
	case 22:
		cod.During08 = common.CourseFlag
	}
}

type CourseInfo struct{
	TeacherName string `json:"teacher_name"`
	StudentId string `json:"student_id"`
	Curriculum string `json:"curriculum"`
	Address string `json:"address"`
}

type CourseOfDay struct {
	Calendar string `json:"calendar"`
	During07 []Course `json:"during07"`
	During09 []Course `json:"during09"`
	During11 []Course `json:"during11"`
	During13 []Course `json:"during13"`
	During15 []Course `json:"during15"`
	During17 []Course `json:"during17"`
	During19 []Course `json:"during19"`
	During21 []Course `json:"during21"`
	During23 []Course `json:"during23"`
}


func (cod *CourseOfDay) SetDuringCourse(num int,courses []Course) {
	switch num {
	case 7:
		cod.During07 = courses
	case 9:
		cod.During09 = courses
	case 11:
		cod.During11 = courses
	case 13:
		cod.During13 = courses
	case 15:
		cod.During15 = courses
	case 17:
		cod.During17 = courses
	case 19:
		cod.During19 = courses
	case 21:
		cod.During21 = courses
	case 23:
		cod.During23 = courses
	}
}

type QueryData struct{
	TeacherId string `json:"teacher_id"`
	TeacherName string `json:"teacher_name"`
	StudentID string `json:"studentId"`
	StudentName string `json:"studentName"`
	GroupName string `json:"group_name"`
	CourseDate string `json:"course_date"`
	QueryType int `json:"queryType"`
}

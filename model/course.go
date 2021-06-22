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
	//ToCourseInfo() CourseInfo
}

type Curriculum struct {
	ID uint `json:"id" gorm:"not null"`
	Name string `json:"name" gorm:"not null"`
}

type CommonCourse struct {
	TeacherUUID string `json:"teacher_uuid" gorm:"not null"`
	TeacherName string `json:"teacher_name" gorm:"not null"`
	CoursePlan uint `json:"course_plan" gorm:"not null"`
	CurriculumNumber uint `json:"curriculum_number" gorm:"not null"`
	StudentGroup int  `json:"student_group" gorm:"not null"`
	CourseDate time.Time `json:"course_date" gorm:"type:date;not null"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	EndTime time.Time `json:"end_time" gorm:"not null"`
	ClassroomNumber int `json:"classroom_number" gorm:"not null"`
	ClassroomAddress string `json:"classroom_address" gorm:"not null"`
	UpdateTime time.Time `json:"update_time" gorm:"not null"`
}

func (c CommonCourse) Type() int {
	return CommonLesson
}

func (c CommonCourse) Calendar() string {
	return c.CourseDate.Format(common.CalendarFormat)
}

func (c CommonCourse) StartClock() string {
	return c.StartTime.Format(common.ClockFormat)
}

//func (c CommonCourse) ToCourseInfo() CourseInfo {
//	return CourseInfo{
//		TeacherName: c.TeacherName,
//		StudentId: c.StudentGroup,
//		Curriculum: c.CurriculumNumber,
//		Address: c.ClassroomNumber,
//	}
//}

type TrialCourse struct {
	TeacherUUID string `json:"teacher_uuid" gorm:"not null"`
	TeacherName string `json:"teacher_name" gorm:"not null"`
	CoursePlan uint `json:"course_plan" gorm:"not null"`
	CurriculumNumber uint `json:"curriculum_number" gorm:"not null"`
	StudentGroup int  `json:"student_group" gorm:"not null"`
	CourseDate time.Time `json:"course_date" gorm:"type:date;not null"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	EndTime time.Time `json:"end_time" gorm:"not null"`
	ClassroomNumber int `json:"classroom_number" gorm:"not null"`
	ClassroomAddress string `json:"classroom_address" gorm:"not null"`
	UpdateTime time.Time `json:"update_time" gorm:"not null"`
}

func (t TrialCourse) Type() int {
	return TrialLesson
}

func (t TrialCourse) Calendar() string {
	return t.CourseDate.Format(common.CalendarFormat)
}

func (t TrialCourse) StartClock() string {
	return t.CourseDate.Format(common.ClockFormat)
}
//func (t TrialCourse) ToCourseInfo() CourseInfo {
//	return CourseInfo{
//		TeacherName: t.TeacherName,
//		StudentId: t.StudentGroup,
//		Curriculum: t.CurriculumNumber,
//		Address: t.ClassroomNumber,
//	}
//}

type SingleCourse struct {
	TeacherUUID string `json:"teacher_uuid" gorm:"not null"`
	TeacherName string `json:"teacher_name" gorm:"not null"`
	CoursePlan uint `json:"course_plan" gorm:"not null"`
	CurriculumNumber uint `json:"curriculum_number" gorm:"not null"`
	StudentUUID string  `json:"student_uuid" gorm:"not null"`
	StudentName string  `json:"student_name" gorm:"not null"`
	CourseDate time.Time `json:"course_date" gorm:"type:date;not null"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	EndTime time.Time `json:"end_time" gorm:"not null"`
	ClassroomNumber int `json:"classroom_number" gorm:"not null"`
	ClassroomAddress string `json:"classroom_address" gorm:"not null"`
	UpdateTime time.Time `json:"update_time" gorm:"not null"`
}

func (s SingleCourse) Type() int {
	return SingleLesson
}

func (s SingleCourse) Calendar() string {
	return s.CourseDate.Format(common.CalendarFormat)
}

func (s SingleCourse) StartClock() string {
	return s.CourseDate.Format(common.ClockFormat)
}
//func (s SingleCourse) ToCourseInfo() CourseInfo {
//	return CourseInfo{
//		TeacherName: s.TeacherName,
//		StudentId: s.StudentUUID,
//		Curriculum: s.CurriculumNumber,
//		Address: s.ClassroomNumber,
//	}
//}

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
	During08 []Course `json:"during08"`
	During10 []Course `json:"during10"`
	During12 []Course `json:"during12"`
	During14 []Course `json:"during14"`
	During16 []Course `json:"during16"`
	During18 []Course `json:"during18"`
	During20 []Course `json:"during20"`
	During22 []Course `json:"during22"`
}


func (cod *CourseOfDay) SetDuringCourse(num int,courses []Course) {
	switch num {
	case 8:
		cod.During08 = courses
	case 10:
		cod.During08 = courses
	case 12:
		cod.During08 = courses
	case 14:
		cod.During08 = courses
	case 16:
		cod.During08 = courses
	case 18:
		cod.During08 = courses
	case 20:
		cod.During08 = courses
	case 22:
		cod.During08 = courses
	}
}

//type CourseWrapper struct {
//	CourseType int
//	Course Course
//}

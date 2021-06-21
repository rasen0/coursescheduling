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

type Curriculum struct {
	ID uint `gorm:"not null"`
	Name string `gorm:"not null"`
}

type CommonCourse struct {
	TeacherUUID string `gorm:"not null"`
	TeacherName string `gorm:"not null"`
	CoursePlan uint `gorm:"not null"`
	CurriculumNumber uint `gorm:"not null"`
	StudentGroup int  `gorm:"not null"`
	CourseDate time.Time `gorm:"type:date;not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime time.Time `gorm:"not null"`
	ClassroomNumber int `gorm:"not null"`
	ClassroomAddress string `gorm:"not null"`
	UpdateTime time.Time `gorm:"not null"`
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

type TrialCourse struct {
	TeacherUUID string `gorm:"not null"`
	TeacherName string `gorm:"not null"`
	CoursePlan uint `gorm:"not null"`
	CurriculumNumber uint `gorm:"not null"`
	StudentGroup int  `gorm:"not null"`
	CourseDate time.Time `gorm:"type:date;not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime time.Time `gorm:"not null"`
	ClassroomNumber int `gorm:"not null"`
	ClassroomAddress string `gorm:"not null"`
	UpdateTime time.Time `gorm:"not null"`
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

type SingleCourse struct {
	TeacherUUID string `gorm:"not null"`
	TeacherName string `gorm:"not null"`
	CoursePlan uint `gorm:"not null"`
	CurriculumNumber uint `gorm:"not null"`
	StudentUUID string  `gorm:"not null"`
	StudentName string  `gorm:"not null"`
	CourseDate time.Time `gorm:"type:date;not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime time.Time `gorm:"not null"`
	ClassroomNumber int `gorm:"not null"`
	ClassroomAddress string `gorm:"not null"`
	UpdateTime time.Time `gorm:"not null"`
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

type CourseOfDay struct {
	Calendar string
	During08 string
	During10 string
	During12 string
	During14 string
	During16 string
	During18 string
	During20 string
	During22 string
}

type CourseWrapper struct {
	CourseType int
	Course Course
}

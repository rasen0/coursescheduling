package database

import (
	"coursesheduling/model"
)

var createSql =`
create students table(
  uuid varchar(255) not null,
  name varchar(50) not null,
  gender tinyint not null,
  age tinyint not null,
  phone varchar(50) not null,
  course_plan int not null,
  relativeName varchar(50),
  relationship varchar(50),
  relative_phone varchar(50),
  startDate date not null,
  endDate date not null,
  desc varchar(255),
  primary key(uuid),
)
create teachers table(
  uuid varchar(255) not null,
  name varchar(50) not null,
  gender tinyint not null,
  age tinyint not null,
  phone varchar(50) not null,
  relativeName varchar(50),
  relationship varchar(50),
  relative_phone varchar(50),
  desc varchar(255),
  primary key(uuid),
)
create student_groups table(
  id int not null,
  student_uuid varchar(50) not null,
)
create single_courses table(
  id int auto_increment,
  teacher_uuid varchar(255) not null,
  teacher_name varchar(50) not null,
  student_name varchar(50) not null,
  student_uuid varchar(255) not null,
  course_plan int not null,
  course int not null,
  classroom varchar(255) not null,
  course_date date not null,
  startTime time not null,
  endTime time not null,
  desc varchar(255),
)
create common_courses table(
  id int auto_increment,
  teacher_uuid varchar(255) not null,
  teacher_name varchar(50) not null,
  student_Group_id int not null,
  course_plan int not null,
  curriculum int not null,
  classroom varchar(255) not null,
  course_date date not null,
  startTime time not null,
  endTime time not null,
  desc varchar(255),
)
create trial_courses table(
  id int auto_increment,
  teacher_uuid varchar(255) not null,
  teacher_name varchar(50) not null,
  student_Group_id int not null,
  course_plan int not null,
  course int not null,
  classroom varchar(255) not null,
  course_date date not null,
  startTime time not null,
  endTime time not null,
  desc varchar(255),
)
create curriculums table(
  id int not null,
  name varchar(50) not null,
)
create classrooms table(
  house_number varchar(50) not null,
  address varchar(255) not null,
)
`

const (
	teacherNumber = 1
	studentNumber = 2
)

const (
	StudentTable = "students"
	TeacherTable = "teachers"
	ClassroomTable = "classrooms"
	CurriculumTable = "curriculums"
	studentGroup = "student_groups"
	CoursePlanTable = "course_plans"
	SingleCourseSchedulingTable = "single_courses"
	TrialCourseSchedulingTable = "trial_courses"
	CommonCourseSchedulingTable = "common_courses"
)

var dBTable = map[string]interface{}{
	StudentTable:&model.Student{},
	TeacherTable:&model.Teacher{},
	ClassroomTable:&model.Classroom{},
	CurriculumTable:&model.Curriculum{},
	studentGroup:&model.StudentGroup{},
	CoursePlanTable:&model.CoursePlan{},
	SingleCourseSchedulingTable:&model.SingleCourse{},
	TrialCourseSchedulingTable:&model.TrialCourse{},
	CommonCourseSchedulingTable:&model.CommonCourse{},
}

func GetTableTotal(tableName string) (total int64) {
	appDB.Table(tableName).Count(&total)
	return total
}

package database

var createSql =`
create student table(
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
create teacher table(
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
create one_course_scheduling table(
  id int auto_increment,
  teacher_name varchar(50) not null,
  teacher_uuid varchar(255) not null,
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
`

func createDB() {

}

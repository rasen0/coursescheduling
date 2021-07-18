package server

import (
	"coursesheduling/lib/dao"
	"coursesheduling/lib/entity"
	"coursesheduling/lib/log"
	"coursesheduling/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

/*
GetStudents 默认请求学生数据
page 参数当前页
count 参数每页展示条数
*/
func (svr *ServeWrapper) GetStudents(ctx *gin.Context)  {
	page := ctx.GetInt("page")
	count := ctx.GetInt("count")
	result := make(map[string]interface{})
	students := dao.GetStudentByPage(page, count)
	result["students"] =students
	ctx.JSON(http.StatusOK,result)
	return
}

// AddStudent 添加一个学生信息
func (svr *ServeWrapper) AddStudent(ctx *gin.Context)  {
	result := make(map[string]interface{})
	var newStudent model.Student
	ctx.BindJSON(&newStudent)
	newStudent.UpdateTime = time.Now()
	log.Printf("new student:%+v",newStudent)
	dao.InsertStudentOne(newStudent)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) StudentPagination(ctx *gin.Context)  {
	result := make(map[string]interface{})
	var pagination entity.Pagination
	log.Print("get students")
	ctx.BindJSON(&pagination)
	_, total := dao.GetStudentPagination(pagination)
	pagination.Total = total
	result["pagination"] = pagination
	students := dao.GetStudentByPage(pagination.CurrentPage-1, pagination.PageSize)
	result["students"] = students
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) GetStudentsByID(ctx *gin.Context)  {
	id := ctx.Query("id")
	log.Printf("GetStudentsByID; id:%v",id)
	students := dao.GetStudentsByID(id)
	result := make(map[string]interface{})
	result["students"] = students
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) GetStudentsByKey(ctx *gin.Context)  {
	key := ctx.Query("query_word")
	log.Printf("GetStudentsByKey; name:%v",key)
	students := dao.GetStudentsByName(key)
	result := make(map[string]interface{})
	type student struct {
		Name string `json:"name"`
		Value string `json:"value"`
	}
	stts := make([]student,0)
	for _,s := range students{
		stt := student{
			Name: s.Name,
			Value: s.SerialNumber,
		}
		stts = append(stts,stt)
	}
	result["student_options"] = stts
	ctx.JSON(http.StatusOK,result)
	return
}

// GroupPagination 组查询
func (svr *ServeWrapper) GroupPagination(ctx *gin.Context)  {
	var pagination Pagination
	ctx.BindJSON(&pagination)
	result := make(map[string]interface{})
	groups := dao.GroupPagination(pagination.CurrentPage,pagination.PageSize)
	type studentInfo struct {
		Id int `json:"id"`
		Name string `json:"name"`
		SerialNumber string `json:"serial_number"`
		Age string `json:"age"`
	}
	type studentGroup struct {
		Id int `json:"id"`
		GroupName string `json:"group_name"`
		Desc string `json:"desc"`
		StudentList []studentInfo `json:"student_list"`
	}
	studentGroups := make([]studentGroup,0)
	for _, group := range groups{
		sg := studentGroup{
			Id:group.ID,
			GroupName: group.GroupName,
			Desc: group.Desc,
			StudentList: make([]studentInfo,0),
		}
		// 获取学生列表
		students := dao.GetStudentsByStudentGroupID(strconv.Itoa(group.ID))
		for _, s := range students {
			stu := studentInfo{
				Id:s.ID,
				Name: s.Name,
				SerialNumber: s.SerialNumber,
				Age: s.Age,
			}
			sg.StudentList = append(sg.StudentList,stu)
		}
		studentGroups = append(studentGroups,sg)
	}
	result["studentGroup"] =studentGroups
	//pagination.CurrentPage++
	//result["pagination"] =pagination
	ctx.JSON(http.StatusOK,result)
	return
}

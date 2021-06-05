package model

type Student struct {
	Name string
	Age  int
	Gender int
	Phone string
	SocialAccount string
	Curriculum int // 可能使用二进制，每位代表一类科目
	UpdateTime string
}

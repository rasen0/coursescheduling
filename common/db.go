package common

const (
	TimeFormat = "2006-01-02 15:04:05"
	CalendarFormat = "2006-01-02"
	CalendarFormat2 = "20060102"
	ClockFormat = "15:04:05"
	BatchCount = 10000
	FreeToken = "zZZ"

	RoleEmptyID = "00000"
	AdminRole = "admin"
    CommonRole = "common_user"
    ReadActive = "read"
	WriteActive = "write"
)

var TimeClock = map[string]int{
	"07:00":7,
	"09:00":9,
	"11:00":11,
	"13:00":13,
	"15:00":15,
	"17:00":17,
	"19:00":19,
	"21:00":21,
	"23:00":23,
}
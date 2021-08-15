package common

const (
	TimeFormat = "2006-01-02 15:04:05"
	CalendarFormat = "2006-01-02"
	CalendarFormat2 = "20060102"
	ClockFormat = "15:04:05"
	BatchCount = 10000
	FreeToken = "zZZ"

	AdminRole = "admin"
    CommonRole = "common_user"
)

var TimeClock = map[string]int{
	"08:00":8,
	"10:00":10,
	"12:00":12,
	"14:00":14,
	"16:00":16,
	"18:00":18,
	"20:00":20,
	"22:00":22,
}
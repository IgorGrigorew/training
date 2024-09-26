package lessonyandex

import (
	"fmt"
	"time"
)


/*
func main() {
	
	start := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
    end := time.Date(2023, 10, 23, 4, 41, 49, 0, time.UTC)
    diff := TimeDifference(start, end)
    fmt.Println(diff)

timestamp := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
format := "2006-01-02 15:04:05"
result := FormatTimeToString(timestamp, format)
fmt.Println(result)

dateString := "2023-10-23 02:41:49"
format := "2006-01-02 15:04:05"
result, err := ParseStringToTime(dateString, format)
if err != nil {
    panic(err)
}
fmt.Println(result)

pastTime := time.Now().Add(-2 * time.Hour)
    result := TimeAgo(pastTime)
    fmt.Println(result) 

daytim := time.Now()
res := NextWorkday(daytim)
fmt.Println(res)



}
*/
func TimeDifference(start, end time.Time) time.Duration{
	return end.Sub(start)
}

//__________________________________________________

func FormatTimeToString(timestamp time.Time, format string) string{

return	timestamp.Format(format)

}

//_____________________________________________________

func ParseStringToTime(dateString, format string) (time.Time, error){
	return time.Parse(format, dateString)
}

//__________________________________________________________

func TimeAgo(pastTime time.Time) string{
	t := time.Now()
	tnew := t.Sub(pastTime)
var str string

	if tnew.Hours() != 0 {
		str = fmt.Sprintf("%.0f hours", tnew.Hours())
	}else{
		str = fmt.Sprintf("%.0f minutes", tnew.Minutes())
	}

s := fmt.Sprintf("%s ago",  str)
return s
}

//_________________________________________________________

func NextWorkday(start time.Time) time.Time{

	if start.Weekday() == 5 {
return start.AddDate(0, 0, 3)
	}else if start.Weekday() == 6 {
		return start.AddDate(0, 0, 2)
	}
	
//	fmt.Println(start.Weekday())
return start.AddDate(0, 0, 1)
}
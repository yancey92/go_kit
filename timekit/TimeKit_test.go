package timekit

import (
	"testing"
	"time"
	"fmt"
)

func TestTimeDemo(t *testing.T) {
	now := time.Now()
	dateStr, err := TimeToString(now, DateFormat_YYYY_MM_DD)
	if err != nil {
		fmt.Printf("%v\n", err)
		t.Fail()
	}

	fmt.Printf("当前日期:%s\n", dateStr)
}

func TestStringToTime(t *testing.T) {
	tm, err := StringToTime("2016-11-11", DateFormat_YYYY_MM_DD)
	if err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	tmStr, err := TimeToString(tm, DateFormat_YYYY_MM_DD_HH_MM_SS)
	if err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	} else {
		fmt.Println(tmStr) //2016-11-11 00:00:00
	}
}

func TestGetTimeSsAndDate(t *testing.T) {
	ss, dateStr, _ := GetTimeSsAndDate(time.Now(), DateFormat_YYYY_MM_DD_HH_MM_SS)
	fmt.Printf("秒=%d,日期=%s\n", ss, dateStr)
}

func TestGetTimeMsAndDate(t *testing.T) {
	ms, dateStr, _ := GetTimeMsAndDate(time.Now(), DateFormat_YYYY_MM_DD_HH_MM_SS)
	fmt.Printf("毫秒=%d,日期=%s\n", ms, dateStr)
}

func TestGetAfterDayMs(t *testing.T) {
	fmt.Println(GetAfterDayMsAndDate("2016-11-11"))
}

func TestGetEndDayMs(t *testing.T) {
	fmt.Println(GetEndDayMsAndDate("2016-11-11"))
}

//获取某天凌晨时间
func TestGetDayTime(t *testing.T)  {
	tm, err := StringToTime("2016-11-11", DateFormat_YYYY_MM_DD)
	if err != nil {
		fmt.Printf("%v", err)
	}
	ms, dateStr, _ := GetTimeMsAndDate(tm, DateFormat_YYYY_MM_DD)
	fmt.Printf("毫秒=%d,日期=%s\n", ms, dateStr)
}
package time

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var newyork *time.Location

type MyDate time.Time

func (d *MyDate) UnmarshalJSON(data []byte) error {
	t, err := time.ParseInLocation(`"2006/01/02 15:04:05.000"`, string(data), newyork)
	if err != nil {
		return err
	}
	*d = MyDate(t)
	return nil
}

func (d MyDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(d).Format("2006/01/02 15:04:05.000 -0700"))), nil
}

func time2json(t time.Time) string {
	b, _ := json.Marshal(MyDate(t))
	return string(b)
}

func json2time(t string) time.Time {
	var myStc struct {
		Timestamp MyDate `json:"timestamp"`
	}
	if err := json.Unmarshal([]byte(t), &myStc); err != nil {
		fmt.Println(err)
	}
	return time.Time(myStc.Timestamp)
}

// zennにコメントするときに書いた下書きコード。
// ref: https://zenn.dev/hsaki/articles/go-time-cheatsheet
func TestNamedType(t *testing.T) {
	ny, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal(err)
	}
	newyork = ny

	jsonTime := `{"timestamp":"2022/04/01 09:00:00.000"}`
	tt := json2time(jsonTime)
	fmt.Println(tt)

	timeTime := time.Date(2022, 4, 1, 9, 0, 99, 0, newyork)
	fmt.Println(time2json(timeTime))
}

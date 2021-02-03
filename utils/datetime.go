package utils

import (
	"time"
)

type JsonDate time.Time
type JsonTime time.Time

func (p *JsonDate) UnmarshalJSON(data []byte) error {

	if len(data) < 10 {
		*p = JsonDate(time.Time{})
		return nil
	}
	local, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)

	*p = JsonDate(local)

	return err
}

func (p *JsonTime) UnmarshalJSON(data []byte) error {

	if len(data) < 10 {
		*p = JsonTime(time.Time{})
		return nil
	}

	local, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*p = JsonTime(local)

	return err
}

func (c JsonDate) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, "2006-01-02")
	data = append(data, '"')
	return data, nil
}

func (c JsonTime) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, "2006-01-02 15:04:05")
	data = append(data, '"')
	return data, nil
}

func (c JsonTime) IsZero() bool {
	return time.Time(c).IsZero()
}

func (c JsonTime) Unix() int64 {
	return time.Time(c).Unix()
}

func (c JsonDate) IsZero() bool {
	return time.Time(c).IsZero()
}
func (c JsonDate) Time() time.Time {
	return time.Time(c)
}
func (c JsonDate) String() string {
	return time.Time(c).Format("2006-01-02")
}
func (c JsonDate) Unix() int64 {
	return time.Time(c).Unix()
}
func GetJsonDateFromTime(t time.Time) JsonDate {
	return JsonDate(t)
}
func (c JsonTime) String() string {
	return time.Time(c).Format("2006-01-02 15:04:05")
}
func (c JsonTime) Now() JsonTime {
	return JsonTime(time.Now())
}
func (c JsonTime) Time() time.Time {
	return time.Time(c)
}
func GetJsonTimeFromTime(t time.Time) JsonTime {
	return JsonTime(t)
}
func GetJsonTimeNow() JsonTime {
	return JsonTime(time.Now())
}
func GetJsonDateNow() JsonDate {
	return JsonDate(time.Now())
}

func (c JsonTime) FormatDay() string {
	return time.Time(c).Format("2006-01-02")
}
func (c JsonTime) FormatMonth() string {

	return time.Time(c).Format("2006-01")
}

func Todate(in string) (out time.Time, err error) {
	out, err = time.Parse("2006-01-02", in)
	return out, err
}

func Todatetime(in string) (out time.Time, err error) {
	out, err = time.Parse("2006-01-02 15:04:05", in)
	return out, err
}

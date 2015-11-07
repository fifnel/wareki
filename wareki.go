package wareki

import (
	"errors"
	"fmt"
	"time"
)

const (
	timeLayout = "2006-01-02 MST"
)

type WarekiInfo struct {
	StartDate time.Time
	Name      string
}

func parseTime(ts string) time.Time {
	t, err := time.Parse(timeStringLayout, ts)

	if err != nil {
		fmt.Printf("param:%v error:%v", ts, err)
		panic("Time string cannot parse")
	}

	return t
}

var WarekiInfoTable = []WarekiInfo{
	{parseTime("1989-01-08 JST"), "平成"},
	{parseTime("1926-12-25 JST"), "昭和"},
	{parseTime("1912-07-30 JST"), "大正"},
	{parseTime("1868-01-25 JST"), "明治"},
}

func toWareki(t time.Time) (string, int, error) {
	for _, v := range WarekiInfoTable {
		if v.StartDate.Before(t) || v.StartDate.Equal(t) {
			return v.Name, t.Year() - v.StartDate.Year() + 1, nil
		}
	}

	return "", 0, errors.New("wareki info is not found")
}

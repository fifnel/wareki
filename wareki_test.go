package wareki

import (
	"fmt"
	"testing"
	"time"
)

const (
	timeStringLayout = "2006-01-02 MST"
)

func TestToWareki(t *testing.T) {

	var d time.Time
	var name string
	var year int
	var err error

	d, _ = time.Parse(timeStringLayout, "2015-12-30 JST")
	name, year, err = toWareki(d)
	fmt.Printf("%v %v\n", name, year)
	if name != "平成" || year != 27 {
		t.Errorf("wareki convert error:%v %v %v", name, year, err)
	}

	d, _ = time.Parse(timeStringLayout, "1989-01-08 JST")
	name, year, err = toWareki(d)
	fmt.Printf("%v %v\n", name, year)
	if name != "平成" || year != 1 {
		t.Errorf("wareki convert error:%v %v %v", name, year, err)
	}

	d, _ = time.Parse(timeStringLayout, "1989-01-07 JST")
	name, year, err = toWareki(d)
	fmt.Printf("%v %v\n", name, year)
	if name != "昭和" || year != 64 {
		t.Errorf("wareki convert error:%v %v %v", name, year, err)
	}

	d, _ = time.Parse(timeStringLayout, "1978-01-08 JST")
	name, year, err = toWareki(d)
	fmt.Printf("%v %v\n", name, year)
	if name != "昭和" || year != 53 {
		t.Errorf("wareki convert error:%v %v %v", name, year, err)
	}
}

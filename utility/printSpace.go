package utility

import (
	"fmt"
	"strconv"
	"time"
)

func PrintSpace(x interface{}, l int) {
	var temp int
	switch v := x.(type) {
	case int:
		temp = len(strconv.Itoa(v))
	case string:
		temp = len(v)
	case float64:
		temp = len(strconv.FormatFloat(v, 'f', 2, 64))
	case time.Time:
		temp = 10
	}

	fmt.Print(" ", x)
	for i := 1; i < l-temp; i++ {
		fmt.Print(" ")
	}
	fmt.Print("|")
}

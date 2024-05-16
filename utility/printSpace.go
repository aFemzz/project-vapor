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
		fmt.Print(" ", x)
	case string:
		temp = len(v)
		fmt.Print(" ", x)
	case float64:
		temp = len(fmt.Sprintf("%2.f", v)) + 3
		if v == 0.0 {
			temp--
		}
		fmt.Printf(" %.2f", x)
	case time.Time:
		temp = 10
		fmt.Print(" ", x)
	}

	for i := 1; i < l-temp; i++ {
		fmt.Print(" ")
	}
	fmt.Print("|")
}

package main

import (
	"fmt"
	"math"
)

func ConvertInt64ToInt(l int64) int {
	if math.MinInt32 <= l && l <= math.MaxInt32 {
		return int(l)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", l))
}

func IntFromInt64(l int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	i = ConvertInt64ToInt(l)
	return
}

func mainPanicDeferConvInt() {
	l := int64(15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err)
	} else {
		fmt.Printf("%d converted to an int32 is %d", l, i)
	}

	fmt.Println()

	l = int64(math.MaxInt32 + 15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err)
	} else {
		fmt.Printf("%d converted to an int32 is %d", l, i)
	}
}

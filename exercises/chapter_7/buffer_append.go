package main

import "fmt"

func mainBufferAppend() {
	sl := []byte{1, 2, 3}
	data := []byte{4, 5, 6}
	sl = Append(sl, data)
	fmt.Println(sl)

	sl = make([]byte, 0, 6)
	sl = Append(sl, data)
	fmt.Println(sl)
}

func Append(sl, data []byte) []byte {
	newLen := len(sl) + len(data)
	if newLen > cap(sl) {
		newSl := make([]byte, newLen)
		copy(newSl, sl)
		copy(newSl[len(sl):], data)
		return newSl
	}
	sl = sl[:newLen]
	copy(sl[len(sl)-len(data):], data)
	return sl
}

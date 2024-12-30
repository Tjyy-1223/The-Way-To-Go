package main

import (
	"log"
	"net/http"
)

type HandleFnc func(http.ResponseWriter, *http.Request)

// 将上一章节示例 15.10 中的代码复制过来，除了 main() 函数都复制过来

func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}

func logPanics(function HandleFnc) HandleFnc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
				// 默认出现 panic 只会记录日志，页面就是一个无任何输出的白页面，
				// 可以给页面一个错误信息，如下面的示例返回了一个 500
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		function(writer, request)
	}
}

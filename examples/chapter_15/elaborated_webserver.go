package main

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var helloRequest = expvar.NewInt("hello-request")

var webroot = flag.String("root", "/home/user", "web root directory")

var booleanflag = flag.Bool("boolean", true, "another flag for testing")

// Counter 简单的服务器计数器，发布它将设置值
type Counter struct {
	n int
}

// Chan 一个通道
type Chan chan int

func Logger(w http.ResponseWriter, req *http.Request) {
	log.Print(req.URL.String())
	w.WriteHeader(404)
	w.Write([]byte("oops"))
}

func HelloServer2(w http.ResponseWriter, req *http.Request) {
	helloRequest.Add(1)
	io.WriteString(w, "hello, world!\n")
}

// 通过这个方法满足 expvar.Var 接口，所以就可以直接导出它。
func (ctr *Counter) String() string {
	return fmt.Sprintf("%d", ctr.n)
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET": // 给 n + 1
		ctr.n++
	case "POST": // 设置 n 去发不值
		buf := new(bytes.Buffer)
		io.Copy(buf, req.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST: %v \n body: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprint(w, "counter reset\n")
		}
	}
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func FlagServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Flags:\n")
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != f.DefValue {
			fmt.Fprintf(w, "%s = %s [default = %s]\n", f.Name, f.Value.String(), f.DefValue)
		} else {
			fmt.Fprintf(w, "%s = %s\n", f.Name, f.Value.String())
		}
	})
}

// ArgServer 简单参数服务器
func ArgServer(w http.ResponseWriter, req *http.Request) {
	for _, s := range os.Args {
		fmt.Fprint(w, s, " ")
	}
}

func ChanCreate() Chan {
	c := make(Chan)
	go func(c Chan) {
		for x := 0; ; x++ {
			c <- x
		}
	}(c)
	return c
}

func (ch Chan) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, fmt.Sprintf("channel send #%d\n", <-ch))
}

// DateServer 执行一个程序，输出重定向
func DateServer(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Fprintf(rw, "pipe: %s\n", err)
		return
	}

	p, err := os.StartProcess("/bin/date", []string{"date"}, &os.ProcAttr{Files: []*os.File{nil, w, w}})
	defer r.Close()
	w.Close()

	if err != nil {
		fmt.Fprintf(rw, "fork/exec: %s\n", err)
		return
	}
	defer p.Release()
	io.Copy(rw, r)
	wait, err := p.Wait()
	if err != nil {
		fmt.Fprintf(rw, "wait: %s\n", err)
		return
	}
	if !wait.Exited() {
		fmt.Fprintf(rw, "date: %v\n", wait)
		return
	}
}

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(Logger))
	http.Handle("/go/hello", http.HandlerFunc(HelloServer2))
	// 计数器直接作为一个变量被发布
	ctr := new(Counter)
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)
	http.Handle("/go/", http.StripPrefix("/go/", http.FileServer(http.Dir(*webroot))))
	http.Handle("/flags", http.HandlerFunc(FlagServer))
	http.Handle("/args", http.HandlerFunc(ArgServer))
	http.Handle("/chan", ChanCreate())
	http.Handle("/date", http.HandlerFunc(DateServer))

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}
}

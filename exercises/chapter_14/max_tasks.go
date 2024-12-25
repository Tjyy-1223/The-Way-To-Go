package main

const (
	// AvailableMemory 10 MB
	AvailableMemory = 10 << 20

	// AverageMemoryPerRequest 10 KB
	AverageMemoryPerRequest = 10 << 10

	MAXREQS = AvailableMemory / AverageMemoryPerRequest
)

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b   int
	replyc chan int
}

func process(r *Request) {
	// Do something 做任何事

	// 可能需要很长时间并使用大量内存或CPU
}

func handle(r *Request) {
	process(r)
	// 信号完成：开始启用下一个请求
	// 将 sem 的缓冲区释放一个位置
	<-sem
}

func Server(queue chan *Request) {
	for {
		sem <- 1
		// 当通道已满（1000 个请求被激活）的时候将被阻塞
		request := <-queue
		go handle(request)
	}
}

func mainMaxTasks() {
	queue := make(chan *Request)
	go Server(queue)
}

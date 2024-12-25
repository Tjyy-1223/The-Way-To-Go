package semaphore

type Empty interface{}
type semaphore chan Empty

var sem semaphore = make(semaphore, 5)

// P acquire n resources
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// V release n resouces
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// Lock /* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

// Wait /* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}

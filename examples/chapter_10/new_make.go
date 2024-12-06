package main

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func mainNewMake() {
	// Ok
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"
}

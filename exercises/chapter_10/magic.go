package main

import "fmt"

type Base2 struct{}

func (Base2) Magic() {
	fmt.Println("base magic")
}

func (self Base2) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base2
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func mainMagic() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}

package main

import (
	"fmt"
	"reflect"
)

func mainReflect2() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	// setting a value
	//v.SetFloat(3.1415) // error:will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v: ", v.CanSet())

	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v: ", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
}

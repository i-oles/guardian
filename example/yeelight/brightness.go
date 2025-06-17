package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var abra int8 = 127

var d1 []string = []string{"1", "2", "3", "4", "5"}

const (
	dupa = "abra"
)

func main() {

	var val int8 = 'x'

	val = abra

	fmt.Printf("a number of bytes: %v\n", unsafe.Sizeof(d1))
	fmt.Println(cap(d1))
	d1 = append(d1, "c")
	fmt.Println(cap(d1))
	fmt.Printf("len d1: %d, type: %v\n", len(d1), unsafe.Sizeof(len(d1)))
	fmt.Printf("len d1: %d, type: %v\n", len(d1), unsafe.Sizeof(len(d1)))
	fmt.Printf("a number of bytes: %v\n", unsafe.Sizeof(d1))

	fmt.Printf("Brightness: %v type: %T\n", val, val)
	fmt.Printf("d1: %v type: %T\n", d1, d1)

	var a int
	fmt.Printf("a type: %v\n", reflect.TypeOf(a))
	fmt.Printf("a number of bytes: %v\n", unsafe.Sizeof(a))

	var x float64
	var y float64

	x = 1.355123
	y = 1.355123

	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)

	fmt.Println(x == y) // nie porównuj w ten sposob

	const epsilon = 1e-9 // jedna miliardowa

	if math.Abs(x-y) < epsilon {
		fmt.Println(" treat as equal")
	}

	t := 10
	p := &t // p to *int
	*p = 20 // zmienia x

	fmt.Printf("x: %v\n", x)

	modify(p)

	fmt.Printf("t: %v\n", t)

	q := [6]int{1, 2}
	fmt.Printf("q: %v\n", q)

	w := []int{4, 5, 6}
	w = w[:0]

	g := w[1:2]
	fmt.Printf("g: %v\n", g)

	//bulbController := controller.NewYeeLight()
	//_, err := bulbController.SetBrightness(
	//	"192.168.0.41:55443",
	//	10,
	//	1000,
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}

	tracing := true

	var client int
	var err error
	if tracing {
		client, err = 24, fmt.Errorf("somesink")
		fmt.Printf("client: %v\n", client)
		fmt.Printf("err: %v\n", err)
	} else {
		client = 12
	}

}

func modify(val *int) {
	*val = 42
}

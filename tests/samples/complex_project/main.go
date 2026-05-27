package main

import (
	"fmt"
	"io"
	"strings"
)

func foobar() { _ = "STUB: not implemented"; return }

func foobar1() string { _ = "STUB: not implemented"; return "" }

func adder() func(int) int { _ = "STUB: not implemented"; return nil }

func generateInteger() int { _ = "STUB: not implemented"; return 0 }

func generateSlice() []int { _ = "STUB: not implemented"; return nil }

func main() {
	a := foobar1()
	fmt.Println(a)

	//
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//
	for _, v := range generateSlice() {
		fmt.Printf("%v %v", v, generateInteger())
	}

	//
	pos, neg := adder(), adder()
	for i := generateInteger() - 1; i < generateInteger(); i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	//
Loop:
	fmt.Println("test")
	for a := 0; a < 5; a++ {
		fmt.Println(a)
		if a > generateInteger() {
			goto Loop
		}
	}

	//
Loop2:
	for j := 0; j < 3; j++ {
		fmt.Println(j)
		for a := 0; a < 5; a++ {
			fmt.Println(a)
			if a > 3 {
				break Loop2
			}
		}
	}

	//
Loop3:
	for j := 0; j < 3; j++ {
		fmt.Println(j)
		for a := 0; a < 5; a++ {
			fmt.Println(a)
			if a > 3 {
				break Loop3
			}
		}
	}

	//
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	//
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//
	do(21)
	do("hello")
	do(true)

	//
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	//
	ss := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(ss[:len(ss)/2], c)
	go sum(ss[len(ss)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	//
	fmt.Println(sqrt(2), sqrt(-4))
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { _ = "STUB: not implemented"; return 0 }

func do(i interface{}) { _ = "STUB: not implemented"; return }

func sum(s []int, c chan int) { _ = "STUB: not implemented"; return }

// send sum to c

func sqrt(x float64) string { _ = "STUB: not implemented"; return "" }

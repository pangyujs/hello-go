package main

import (
	"fmt"
	"time"
)

func main() {
	var stockcode = 123
	var enddate = "2023-04-05"
	var url = "Code=%d&endDate=%s"
	fmt.Printf(url, stockcode, enddate)
	fmt.Println("hello,world")
	getTypes()
	getVars()
	getCalculate()
	getCondition()
	getCycle()
}

func getCycle() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)

	strs := []string{"google", "bing"}
	for i, s := range strs {
		println(i, s)
	}

	nums := [7]int{1, 2, 3, 4, 5}
	for i, x := range nums {
		println(i, x)
	}

	// break
	a := 10
	for a < 20 {
		println("a====%d", a)
		a++
		if a > 15 {
			break
		}
	}

	// goto
	var s int = 10
LOOP:
	for s < 20 {
		if s == 15 {
			s = s + 1
			goto LOOP
		}
		println("s======%d", s)
		s++
	}

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			println("received", msg1)
		case msg2 := <-c2:
			println("received", msg2)
		}
	}
}

func getCondition() {
	a, b := 21, 10
	if a == b {
		println("a==b")
	} else {
		println("a!=b")
	}
	if a < b {
		println("a<b")
	} else {
		println("a>=b")
	}
	a = 5
	b = 20
	if a <= b {
		println("a<=b")
	} else {
		println("a>b")
	}

	c, d := true, false
	if c && d {
		println("c&&b")
	}
	if c || d {
		println("a||b")
	}
	c, d = false, true
	if c && d {
		println("c&&d")
	}

}

func getCalculate() {
	var a = 21
	var b = 10
	var c int
	c = a + b
	println("第一行 - c的值为%d", c)
	c = a - b
	println("第一行 - c的值为%d", c)
	c = a * b
	println("第一行 - c的值为%d", c)
	c = a / b
	println("第一行 - c的值为%d", c)
	c = a % b
	println("第一行 - c的值为%d", c)
	a++
	println("第一行 - c的值为%d", a)
	a = 21
	a--
	println("第一行 - c的值为%d", a)
}

func getTypes() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

var x, y int
var c, d int = 1, 2
var e, f = 123, "hello"

func getVars() {
	g, h := 11123, "hello"
	e, f = 1111, "123132"
	println(x, y, c, d, e, f, g, h)
}

package main

import (
	"fmt"
	"strconv"
	"time"
)

type Writer interface {
	Write([]byte) (int, error)
}
type StringWriter struct {
	str string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func main() {
	var w Writer = &StringWriter{}
	sw := w.(*StringWriter)
	sw.str = "hello world Writer"
	println(sw.str)
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
	getFunction()
	getArray()
	getPointer()
	getStruct()
	getArraySplit()
	getRange()
	getMap()
	getRecurise()
	getTypeRevert()
	getThread()
}

func getThread() {
	say := func(s string) {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Microsecond)
			println(s)
		}
	}
	go say("has thread")
	say("no thread")

	sum := func(s []int, c chan int) {
		sum := 0
		for _, v := range s {
			sum += v
		}
		c <- sum
	}

	s := []int{1, 2, 3, 4, 5, 9, 8}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	println("x,y,value", x, y, x+y)

	fibonacci := func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}

	cdd := make(chan int, 10)
	go fibonacci(cap(cdd), cdd)
	for i := range cdd {
		println(i)
	}
}

func getTypeRevert() {
	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Printf("mean的值为: %f\n", mean)

	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		println("revert error:", err)
	} else {
		println("revert success:", str, num)
	}

	var i interface{} = "hello world"
	str, ok := i.(string)
	if ok {
		println("It is a string", str)
	} else {
		println("conversion failed")
	}

}

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func getRecurise() {
	var i int = 15
	fmt.Printf("%d的阶乘为%d\n", i, factorial(uint64(i)))
}
func getMap() {
	var siteMap map[string]string
	siteMap = make(map[string]string)
	siteMap["Google"] = "google value"
	siteMap["Baidu"] = "Baidu value"

	for key, value := range siteMap {
		println("key", key, "value", value)
	}

	name, ok := siteMap["Google"]
	if ok {
		println("Google is exist", name, ok)
	} else {
		println("Google is not exist")
	}

	countryMap := map[string]string{
		"china":   "beijing",
		"america": "houshengdun",
	}
	delete(countryMap, "china")
	println(countryMap["china"])

}

func getRange() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	map1[5] = 5.0

	for key, value := range map1 {
		fmt.Printf("key is %d - value is %f\n", key, value)
	}

	for key := range map1 {
		fmt.Printf("key is %d\n", key)
	}

	for _, value := range map1 {
		fmt.Printf("key is %f\n", value)
	}

}

func getArraySplit() {
	var printSlice = func(x []int) {
		fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	}
	var nums = make([]int, 3, 5)
	printSlice(nums)

	var nums2 []int
	nums2 = append(nums2, 0)
	printSlice(nums2)
	nums2 = append(nums2, 1)
	printSlice(nums2)
	nums2 = append(nums2, 2, 3, 4)
	printSlice(nums2)
	nums3 := make([]int, len(nums2), (cap(nums2) * 2))
	printSlice(nums3)
	copy(nums3, nums2)
	printSlice(nums3)

}

func getStruct() {
	type Books struct {
		title   string
		author  string
		subject string
		book_id int
	}

	var book1 Books
	book1.title = "GO语言"
	book1.author = "www.mengxiangyu.com"
	book1.subject = "hello-go"
	book1.book_id = 22222

	var printBook = func(book *Books) {
		println("book title is %s", book.title)
		println("book author is %s", book.author)
		println("book subject is %s", book.subject)
		println("book book_id is %s", book.book_id)
	}
	printBook(&book1)
}

func getPointer() {
	var a int = 20
	var ip *int
	ip = &a
	println("a address is %d", &a)

	println("ip address is %d", ip)

	println("*ip address is %d", *ip)
}

func getArray() {
	var arr = []int{1, 2, 3, 3, 4}
	fmt.Println(arr[1])
}

func getFunction() {
	getMax := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	println(getMax(1, 2))
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

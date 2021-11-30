package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"rsc.io/quote"
)

const Answer = 42

func main() {
	files()
	// runStuff()

	fmt.Println()
}

func runStuff() {
	fmt.Println("Hello world")
	fmt.Println(quote.Go())

	mathFn()

	fmt.Println("add stuff", sumStuff(2, 3))

	a, b := doubleTrouble()
	fmt.Println(a, b)

	declTypes()
	basicTypes()

	fmt.Println("The answer is", Answer)

	bitShifting()

	loops()

	conditions()

	switching()

	deferFn()

	pointies()

	buildingThings()

	lists()

	makingThings()

	mapping()

	funcFun()

	methods()
}

func mathFn() {
	fmt.Println("math")
	fmt.Println(math.Sin(2))

	fmt.Println("random numbers")
	fmt.Println(rand.Intn(100), rand.ExpFloat64())
}

func sumStuff(x int, y int) int {
	return x + y
}

func doubleTrouble() (string, string) {
	return "foo", "bar"
}

func declTypes() {
	var i, j, k int = 1, 2, 3
	o := 42
	foo, bar := "zap", "zing"
	fmt.Println(i, j, k, o, foo, bar)
}

func basicTypes() {
	var b bool = true
	var s string = "string"
	var i int = 99
	//  int8,16,32,64
	// uint8,16,32,64,uintptr
	// byte, rune
	var f float32 = 1.23
	var c complex64 = complex64(123)

	fmt.Println(b, s, i, f, c)
}

func bitShifting() {
	val := 1 << 4
	val2 := val >> 2
	fmt.Println("bits shifting", val, val2)
}

func loops() {
	for i := 0; i < 10; i++ {
		fmt.Print(1<<i, ",")
	}
	fmt.Println()

	var fibs = []int{1, 2, 3, 5, 8, 13, 21}
	for idx, val := range fibs {
		fmt.Print(idx, ":", val, ",")
	}
	fmt.Println()
}

func conditions() {
	if Answer == 42 {
		fmt.Println("sorry for the inconvenience")
	}
	if v := Answer << 2; v > 100 {
		fmt.Println("that's a better answer", v)
	}
	if true == false {
		fmt.Println("i feel rather like a sofa")
	} else {
		fmt.Println("stability of the universe confirmed")
	}
}

func switching() {
	val := "back"

	switch val {
	case "blade":
		fmt.Println("sharp!")
	case "back":
		fmt.Println("snazzy!")
	default:
		fmt.Println("dunno!")
	}

	today := time.Now().Weekday()
	switch today {
	case time.Saturday:
		fmt.Println("party time")
	case time.Saturday - 1:
		fmt.Println("TGIF")
	default:
		fmt.Println("another time")
	}

	switch {
	case true == false:
		fmt.Println("never")
	case Answer > 1:
		fmt.Println("the cool kids way to write if/else blocks")
	default:
		fmt.Println("we fell out")
	}
}

func deferFn() {
	defer fmt.Println("I came first!")
	fmt.Println("But I was executed first!")
	defer fmt.Println("And I was last!")
}

func pointies() {
	var p *int
	badNum := 666
	p = &badNum
	fmt.Println("bad num", *p, p)
}

func buildingThings() {
	type Aminal struct {
		color  string
		age    int
		flavor bool
	}
	c := Aminal{"blue", 2, false}
	fmt.Println(c)

	p := &c
	p.color = "black"
	fmt.Println(*p)

	d := &Aminal{age: 17}
	fmt.Println(d, *d)
}

func lists() {
	var a [10]string
	a[0] = "Hello"
	a[6] = "World"
	bestPart := a[0:7]

	evens := [5]int{2, 4, 6, 8, 10}

	fmt.Println(bestPart, evens[1:3], evens[3:])
	fmt.Println("best", len(bestPart), cap(bestPart))

	slice := []struct {
		label string
		value int
	}{
		{"foo", 1},
		{"bar", 2},
		{"baz", 3},
	}
	fmt.Println("slices", slice)

	pie := evens[0:]
	pie2 := append(pie, 12)
	fmt.Println("pie?", pie, pie2)
}

func makingThings() {
	a := make([]int, 5)
	b := make([]int, 0, 10)

	fmt.Println("make", a, b)
}

func mapping() {
	type RGB struct {
		r int
		g int
		b int
	}
	colorMap := make(map[string]RGB)
	colorMap["red"] = RGB{255, 0, 0}
	colorMap["puce"] = RGB{204, 136, 153}

	fmt.Println(colorMap)

	var favoriteShades = map[string]RGB{
		"sprange":     {1, 2, 3},
		"guarve":      {4, 5, 6},
		"light urple": {7, 8, 9},
	}
	fmt.Println(favoriteShades)

	elem, found := favoriteShades["poop"]
	fmt.Println("found it?", found, elem)
}

func funcFun() {
	workIt := func(s string, fn func(string) string) string {
		return fn(s)
	}

	daft := workIt("harder", strings.ToUpper) //+ makeIt("better", strings.ToLower) +  doIt("faster", makesUs("stronger"))))
	fmt.Println("punk", daft)
}

func closures() {
	// wut?
}

type Madness struct {
	peak, trough int
}

func (m Madness) Avg() int {
	return (m.peak + m.trough) / 2
}

func methods() {
	redHaze := Madness{98, 47}
	fmt.Println("yo", redHaze.Avg(), Madness.Avg(redHaze))
}

package main

/*
	A Tour of Go: practice

*/

import (
	"errors"
	"fmt"
	"image/color"
	"io"
	"math"
	"math/cmplx"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

func WordCount(s string) (res map[string]int) {
	splt := strings.Split(s, " ")

	for _, st := range splt {
		v, ok := res[st]

		if ok {
			res[st] = v + 1
		} else {
			res[st] = 1
		}
	}

	return
}

/*
adds two numbers
*/
func add(a, b int) int {
	return a + b
}

// multiple returns
func swap(str1, str2 string) (string, string) {
	return str2, str1
}

// named returns
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// variables
var c, python, java bool
var i, j, k int = 1, 2, 3

const PI = 3.1415

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func one() {
	a := 5
	fmt.Println(a)
	fmt.Println("Hello world!")
	fmt.Println("Welcome to the playground")
	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", math.MaxInt, math.MinInt, rand.Intn(10))

	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	fmt.Println(add(4, 4))

	fmt.Println(swap("tamie", "abbie"))
	fmt.Println(split(5))
	fmt.Println(c)
	fmt.Println(python)
	fmt.Println(java)
	fmt.Println(j)
	fmt.Println(i)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	g := 4 + 5i

	fmt.Println(g)

	const world = "world!" // can't be declared with :=
	fmt.Println("Hello", world)
}

func looping() {
	// GO has only one looping construct which is the for loop

	sm := 0
	for i := 1; i <= 10; i++ {
		sm += i
	}

	println(sm)

	sum := 1

	for sum < 100 {
		sum += sum
	}

	fmt.Println(sum)

	// for {
	// 	// infinite loop
	// }

	println(sqrt(-4))
	println(sqrt(7))

	println(pow(3, 8, 100))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)

	}
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	counting()
}

var t = time.Now()

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 12:
		fmt.Println("Good afrernoon!")
	default:
		fmt.Println("Good evening!")

	}
	return fmt.Sprint(math.Sqrt(x))
}

func Sqrt(x float64) float64 {
	z := 1.0
	prev := -1.0
	precision := 0.0000001

	for math.Abs(z-prev) > precision {
		prev = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func counting() {
	// stacking...
	println("counting...")
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done!")
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v <= lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func pointers() {
	// go has pointers. pointer holds memory address of a value
	// The type *T is a pointer to a value. Its zero value is nil.

	var p *int

	i := 127
	p = &i

	println(p)
	println(*p) // dereferencing
	// unlike c, Go  has no pointer arithmetic

	c, d := 34, 56
	p = &c
	fmt.Println(*p)
	*p = 21
	fmt.Println(c)

	p = &d

	*p /= 7

	print(d)
}

// struct: collection of fields

type Vertex struct {
	X, Y float64
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 3} // y: implicitly 0
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

type MyFloat float64

// Interface is set of methods signatures
type Abser interface {
	Abs() float64
}

func main() {
	// one()
	// looping()
	// fmt.Println(Sqrt(2))
	// pointers()

	fmt.Println(Vertex{1, 3})
	v := Vertex{X: 3, Y: 4}

	println(v.X)
	println(v.Y)

	p := &v
	p.X = 1e9

	fmt.Println(p)

	var a [10]int
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	str := "Afhds sdfjbsndjf"
	strings.Split(str, "")

	fmt.Println(primes[1:3])

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	st := []struct {
		id  int
		ans bool
	}{
		{1, true},
		{2, false},
		{3, false},
		{4, true},
		{5, true},
	}

	cpy := st[:]

	// Compare slices element by element
	areEqual := len(st) == len(cpy)
	if areEqual {
		for i := range st {
			if st[i] != cpy[i] {
				areEqual = false
				break
			}
		}
	}
	fmt.Println(areEqual)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s1 := s[:0]
	printSlice(s1)

	s = s[2:]
	printSlice(s)

	aa := make([]int, 3, 5)
	fmt.Println(aa)

	// slices can contain any time of data

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	ss := []int{1, 3, 5}

	ss = append(ss, 4)
	fmt.Println(ss)

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128, 256}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // 2 ** i
	}

	// Maps

	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40, -74}

	fmt.Println(m["Bell Labs"])

	m = map[string]Vertex{
		"Google": Vertex{
			34, 56,
		},
		"Bellman Ford": {
			56, 78,
		},
	}

	delete(m, "Google")

	elem, ok := m["Google"]

	println(m)
	println(ok)
	fmt.Println(elem)

	fmt.Println(wordCount("abc def g"))

	// functions are values too

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	println((compute(hypot)))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// go doesn't have classes, but can define methods on types
	// method is a function with a special receiver argument

	vvv := Vertex{4, 3}
	fmt.Println(vvv.Abs())

	// on non struct types
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Pointer receiver
	v2 = Vertex{3, 4}
	fmt.Println(v2.Abs())
	v2.Scale(5.6)
	fmt.Println(v2.Abs())

	var a1 Abser

	f1 := MyFloat(-math.Sqrt2)
	v1 := Vertex{3, 4}

	a1 = f1
	a1 = &v1

	a1 = v1

	fmt.Println(a1.Abs())

	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = MyFloat(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hellooooooooo"}
	describe(i)
	i.M()

	// var ii I
	// describe(ii)
	// ii.M()

	// var iii interface{}
	// describe(iii)

	// iii = 43
	// describe(iii)
	// iii = "hello"
	// describe(iii)

	var ccc interface{} = "hello"
	sss := ccc.(string)

	fmt.Println(sss)

	sss1, ok := ccc.(string)
	if ok {
		fmt.Println(sss1)
	}

	do(21)
	do("hello")
	do(true)

	aaa := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(aaa, z)

	hosts := map[string]IPAdr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt1(2))
	fmt.Println(Sqrt1(-2))
	fmt.Println(NthRoot(8, 2))

	fmt.Println("------------------------------")

	to_be_read := "reading bytes from this string"
	rdr := strings.NewReader(to_be_read)
	b := make([]byte, len(to_be_read))

	for {
		n, err := rdr.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	fmt.Println("------------------------------")
	s3 := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r3 := rot13Reader{s3}
	io.Copy(os.Stdout, &r3)

	si := []int{1, 3, 5, 7}
	ss2 := []string{"A", "B", "C", "D"}

	fmt.Println(Index(si, 3))
	fmt.Println(Index(ss2, "B"))
}

type List[T any] struct {
	Next *List[T]
	Val  T
}

type LinkedList interface {
	Add(val int) bool
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type Rectangle struct {
}
type Image interface {
	ColorModel() color.Model
	Bounds() Rectangle
	At(x, y int) color.Color
}

func (reader rot13Reader) Read(p []byte) (int, error) {
	n, err := reader.r.Read(p)

	for i := 0; i < n; i++ {
		switch {
		case 'A' <= p[i] && p[i] <= 'Z':
			p[i] = 'A' + (p[i]-'A'+13)%26
		case 'a' <= p[i] && p[i] <= 'z':
			p[i] = 'a' + (p[i]-'a'+13)%26
		}
	}

	return n, err
}

type rot13Reader struct {
	r io.Reader
}

type MyReader struct{}

func (reader MyReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 'A'
	}
	return len(buf), nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func NthRoot(x float64, n int) (float64, error) {
	if n <= 0 {
		return 0, errors.New("n must be positive")
	}

	if x < 0 && n%2 == 0 {
		return 0, errors.New("even root of negative number is not real")
	}

	z := 1.0
	prec := 1e-7
	prev := 0.0

	for math.Abs(z-prev) > prec {
		prev = z
		z = ((float64(n)-1)*z + x/math.Pow(z, float64(n-1))) / float64(n)
	}

	return z, nil
}
func Sqrt1(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	prev := -1.0
	prec := 0.000001
	for math.Abs(z-prev) > prec {
		prev = z
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

type MyError struct {
	When time.Time
	What string
}

type IPAdr [4]byte

func (ip IPAdr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func (f MyFloat) M() { // type MyFloat implements interface I
	fmt.Println(f)
}

func (t *T) M() { // Type T implements interface I
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type T struct {
	S string
}
type I interface {
	M()
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f

}
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y))
}
func adder() func(int) int {
	sm := 0
	return func(x int) int {
		sm += x
		return sm
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func wordCount(s string) (cnt map[string]int) {
	cnt = make(map[string]int)
	for _, word := range strings.Split(s, " ") {
		cnt[word] += 1
	}
	return
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

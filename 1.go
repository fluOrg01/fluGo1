package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var c1, python1, java1 bool
var j, k int = 1, 2

func fAdd(x int, y int) int {
	return (x + y)
}

func fSwap(x, y string) (string, string) {
	return y, x
}

func main() {
	var i int
	var c2, python2, sJava = true, false, "Str01"
	k := 3

	var bVal1 bool = true

	var sVal1 string = "s1"

	var nVal int = 0
	//var n8Val int8 = 0
	//var n16Val int16 = 0

	//n32Val int32 = 0
	//n32Val rune = 0 // alias for int32, Unicode code point

	//n64Val int64 = 0

	//var uVal uint = 0

	//u8Val1 uint8 = 0
	//u8Val2 byte = 0 // alias for uint8

	//u16Val uint16 = 0
	//u32Val uint32 = 0
	//u64Val uint64 = (1 << 64) - 1

	//puVal uintptr = 0

	//var f32Val float32 = 0
	var f64Val float64 = 0

	//z64Val complex64 =
	//z128Val complex128 = cmplx.Sqrt(-5 + 12i)

	const sStr2 string = "s2"

	type Type01 struct {
		nVal1 int
		nVal2 int
	}

	defer fmt.Println("main done")

	Struc01 := Type01{1, 2}
	fmt.Println(Struc01.nVal1)
	Struc01.nVal1 = 4
	fmt.Println(Struc01.nVal1)
	p := &Struc01
	p.nVal2 = 1e9
	fmt.Println(Struc01)

	fmt.Println("GOOS=", runtime.GOOS)
	fmt.Println("NumCPU=", runtime.NumCPU())

	TimeCurr := time.Now()
	fmt.Println("TimeCurr=", TimeCurr)

	fmt.Println("number", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(fAdd(42, 13))

	a, b := fSwap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(i, c1, python1, java1)
	fmt.Println(i, j, c2, python2, sJava, k)

	fmt.Printf("nVal=%v f64Val=%v %v %q\n", nVal, f64Val, bVal1, sVal1)

	//uVal = uint(f64Val)

	f64Val = float64(nVal)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	pnVal := &nVal
	fmt.Println(*pnVal)
	*pnVal = 2

	var tS1 [2]string
	tS1[0] = "Hello"
	tS1[1] = "World"
	fmt.Println(tS1[0], tS1[1])
	fmt.Println(tS1)

	tn2 := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(tn2)

	/*
		Slices are like references to arrays
		A slice does not store any data, it just describes a section of an underlying array.
		Changing the elements of a slice modifies the corresponding elements of its underlying array.
		Other slices that share the same underlying array will see those changes.
	*/
	var tnDyn1 []int = tn2[1:4]
	fmt.Println(tnDyn1)

	tnDyn2 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(tnDyn2)

	tbDyn3 := []bool{true, false, true, true, false, true}
	fmt.Println(tbDyn3)

	Struc02 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(Struc02)

	type Type03 struct {
		Lat, Long float64
	}
	var Map01 map[string]Type03
	Map01 = make(map[string]Type03)
	Map01["Bell Labs"] = Type03{
		40.68433, -74.39967,
	}
	fmt.Println(Map01)
	fmt.Println(Map01["Bell Labs"])
}

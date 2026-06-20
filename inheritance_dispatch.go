package main

import "fmt"

// Lecture 32 Example 3 translated into Go.
// This tests variable inheritance / field lookup.

type C1 struct {
	X int
	Y int
}

func (c *C1) Initialize() {
	// Does not need to do anything for this test.
}

func (c *C1) SetX1(v int) {
	c.X = v
}

func (c *C1) SetY1(v int) {
	c.Y = v
}

func (c *C1) GetX1() int {
	return c.X
}

func (c *C1) GetY1() int {
	return c.Y
}

type C2 struct {
	C1
	Y int
}

func (c *C2) SetY2(v int) {
	c.Y = v
}

func (c *C2) GetX2() int {
	return c.X
}

func (c *C2) GetY2() int {
	return c.Y
}

func runExample3() {
	o2 := C2{}

	o2.SetX1(101)
	o2.SetY1(102)
	o2.SetY2(999)

	fmt.Println("Lecture 32 Example 3")
	fmt.Println("Variable inheritance / field lookup test:")
	fmt.Println("Expected idea from lecture: list(getx1, gety1, getx2, gety2)")
	fmt.Printf("Result: (%d %d %d %d)\n",
		o2.GetX1(),
		o2.GetY1(),
		o2.GetX2(),
		o2.GetY2())

	fmt.Println("Explanation:")
	fmt.Println("GetY1 uses the Y field from C1.")
	fmt.Println("GetY2 uses the Y field from C2.")
	fmt.Println("This shows Go field selection is static based on the code written.")
}

// Lecture 32 Example 5 translated into Go.
// This tests method dispatch and a super-style method call.
// Go does not have super, so we simulate the idea by
// explicitly calling the parent behavior.

type HasM2 interface {
	M2() int
}

type C1Methods struct{}

func (c C1Methods) M1(self HasM2) int {
	return self.M2()
}

func (c C1Methods) M2() int {
	return 13
}

type C2Methods struct {
	C1Methods
}

func (c C2Methods) M1() int {
	return 22
}

func (c C2Methods) M2() int {
	return 23
}

func (c C2Methods) SuperM1(self HasM2) int {
	return c.C1Methods.M1(self)
}

type C3Methods struct {
	C2Methods
}

func (c C3Methods) M1() int {
	return 32
}

func (c C3Methods) M2() int {
	return 33
}

func (c C3Methods) M3() int {
	// This simulates: super m1()
	// It starts with the parent version of m1,
	// but self is still the C3 object.
	return c.C2Methods.SuperM1(c)
}

func runExample5() {
	o3 := C3Methods{}

	fmt.Println("\nLecture 32 Example 5")
	fmt.Println("Method dispatch / super-style call test:")
	fmt.Println("Expected idea from lecture: send o3 m3()")
	fmt.Println("Result:", o3.M3())

	fmt.Println("Explanation:")
	fmt.Println("M3 calls the parent version of M1.")
	fmt.Println("That parent M1 calls self.M2().")
	fmt.Println("Since self is still the C3 object, C3's M2 runs.")
	fmt.Println("Therefore, the result is 33.")
}

func main() {
	runExample3()
	fmt.Println()
	runExample5()
}

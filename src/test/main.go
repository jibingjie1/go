package main

func main() {
	testChannel5()
}
func foo1(a string, b int) int {
	println("a=", a, "b=", b)
	return b
}

func foo2(a string, b int) (int, int) {
	println("a=", a, "b=", b)
	return b, b + 100
}

func foo3(a string, b int) (r1 int, r2 int) {
	r1 = 100
	r2 = 200
	println("a=", a, "b=", b)
	return
}

func foo4(a string, b int) (r1, r2 int) {
	r1 = 300
	r2 = 400
	println("a=", a, "b=", b)
	return
}

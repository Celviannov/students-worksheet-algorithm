package main

import "fmt"

const pi = 3.14

func main() {
	a, b := 7, 2
	quotient := a / b
	remainder := a % b
	fmt.Printf("int: %d / %d = %d, %d %% %d = %d\n", a, b, quotient, a, b, remainder)

	ratio := float64(a) / float64(b)
	fmt.Printf("float64: %.1f\n", ratio)

	score, attendance := 72, 70
	passed := score >= 60 && attendance >= 75
	fmt.Println("bool: passed =", passed)

	greeting := "Hello" + ", " + "Go"
	fmt.Printf("string: %s (len %d)\n", greeting, len(greeting))

	letter := 'A'
	next := letter + 1
	fmt.Printf("rune: %c -> %d -> %c\n", letter, letter, next)

	radius := 5.0
	area := pi * radius * radius
	fmt.Printf("const: area = %.1f\n", area)

	_, _, _ = score, attendance, radius
}

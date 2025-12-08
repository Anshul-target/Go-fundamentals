package main

func main() {
	// A const a value known at compile time that cannot change
	// Two types
	// Typed constants
	// Untyped constants(unique to GO)

	const x = 10
	const y int = 20

	// Cannot change after declaration

	// Must be assigned a compile-time value

	// Cannot use functions (unless compiler evaluates them)

	println(x)
	untypedConst()

}
func untypedConst() {
	//A constant does not take memory like a normal variable.

	// The compiler literally replaces the constant with the value wherever it appears.
	const y int = 1
	var a int = y
	println(a)
}
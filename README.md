# GO notes in one page

The goal of this note is to write all important go notes in one page for quick search and revist.

The source is from this [A Tour of Go](https://tour.golang.org/welcome/1).

## Packages
* Every Go program is made up of packages.
* Programs start running in package **main**.
* This program is using the packages with import paths `"fmt"` and `"math/rand"`.
* By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin with the statement `package rand`.

* **Note:** The enviroment in which these programs are execued is deterministic, so each time you run the example program `rand.Intn` will return the same number.

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
    rand.Seed(101)
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

## Imports
* This code groups the imports into a parenthesized, "factored" import statement. You can also write multiple import statements, like:
```
import "fmt"
import "math"
```
* But it is good style to use the factored import statement.
```go
package main
import (
  "fmt"
  "math"
)
func main() {
  fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```
* [`Package fmt`](https://golang.org/pkg/fmt/?m=all)
   * %v: The value in a default format whne printing structs, the plus flag (%+v) adds field names.
   * %#v: a Go-syntax representation of the value.
   * %T: a Go-syntax representation of the type of the value.
   * %%: a literal percent sign; consumes no value
   * %t: the word true or flase
* Integer:
   * %b base 2
   * %d base 10
   * %o base 8
   * %O base 8 with 0o prefix
   * %x base 16, with lower-case letters for a-f
   * %X base 16, with upper-case letters for A-F
* Floating:
   * %e (%E): scientific notation
   * %f: decimal point but no exponent, e.g. 123.456
   * %g (%G): %e for large exponents, %f otherwise.
   * %f: default width, default precision
   * %9.2f: width 9, precision 2


## Exported names
* In Go, a name is exported if it begins with a capital letter.
## Functions
* A function can take zero or more arguments.
* Notice that the type comes after the variable name.
```go
package main

import "fmt"

func add(x,y int) int {
  return x+y
}

func main() {
  fmt.Println(add(42,13))
}
```
* When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
* A function can return any number of results.
```go
package main

import "fmt"

func swap(x,y string) (string, string) {
  return y,x
}

func main() {
  a, b := swap("hello", "world")
  fmt.Println(a,b)
}
```
* Named return values:
   * Go's return values may be named. If so, they are treated as variables defined at the top f the function.
   * A `return` statement without arguments returns the named return values. This is known as a "naked" return.
   * Naked return statements should be used only in short functions, as with the sample shown here. 
```go
package main

import "fmt"

func split(sum int) (x,y int) {
  x = sum*4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(split(17))
}
```

## Variables
* The `var` statement declares a list of variables; as in function argument lists, the type is last.
* A `var` statement can be at package or function level.
```go
package main

import "fmt"

var c, python, java bool

func main() {
  var i int
  fmt.Println(i,c,python,java)
}
```
* Variables with initializers:
  * A var declaration can include initializers, one per variable.
  * If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
  var c, python, java = true, fase, "no!"
  fmt.Println(i, j, c, python, java)
}
```
* Short variable declarations : `:=`
  * Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.
  * Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available

```go
package main

imprt "fmt"

func main() {
  var i, j int = 1, 2
  k := 3
  c, python, java := true, false, "no!"
  fmt.Println(i, j, k, c, python, java)
}
```

## Basic types
* Go's basic types are:
```
bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte: //alias for unit8
      // represents a Unicode code point

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
* The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
```go
package main

import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe bool = false
  MaxInt uint64 = 1<<64 - 1
  z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value %v\n", z, z)
}
```
* Zero values:
  * Variables declared without an explicit initial value are given their zero value
  * The zero value is:
      * 0 for numeric types
      * false for the boolean type
      * "" (the empty string) for strings

```go 
package main

import "fmt"

func main() {
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```
* Type conversions 
  * The expression T(v) converts the value v to the type T.
  * Some numeric conversions:
```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
Or, put more simply:
```
i :=42
f := float64(i)
u := uint(f)
```
* Unlike in C, in Go assignment between items of different type requires an explicit conversion.
* Type inference:
  * When declaring a variable without specifying an explicit type( either by using the `:=` syntax or `var=` expression syntax), the variable's type is inferred from the value on the right side.
```
var i int
j := i // j is an int

i := 42 //int
f := 3.142 //float64
g := 0.867 + 0.5i //complex128
```
# Constants
* Constants are declared like variables, but with the `const` keyword.
* Constants can be character, string, boolean, or numeric values.
* Constants cannot be declared using the := syntax.
```go
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```
* Numeric Constants: Numeric constants are high-precision values.
* An untyped constant takes the type needed by its context.
```go
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

## For
* Go has only one looping construct, the for loop.
* The basic `for` loop has three components separated by semicolons:
  * the initial statement: executed before the first iteration
  * the condition expression: evaluated before every iteration
  * the post statement: executed at the end of every iteration
* The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the `for` statement.
* The loop will stop iterating once the boolean condition evaluates to false
```go
package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
```
* The init and post statements are optional. 
```go
package main

import "fmt"

func main() {
  sum := 1
  for ; sum < 1000; {
    sum += sum
  }
  fmt.Println(sum)
}
```
* For is Go's "while": At that point you cna drop the semicolons.
```go
package main

import "fmt"

func main() {
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}
```
## If
* Go's if statements are like its for loops
* the expression need not be surrounded by parentheses ( ) but the braces { } are required.
```go
package main

import (
  "fmt"
  "math"
)

func sqrt( x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func main() {
  fmt.Prinln(sqrt(2), sqrt(-4))
}
```
* If with a short statement: ike for, the if statement can start with a short statement to execute before the condition.
* Variables declared by the statement are only in scope until the end of the if.
```go
package main

import (
  "fmt"
  "math"
)

func pow(x, n, lim, float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}

func main() {
  fmt.Println( 
    pow(3,2,10),
    pow(3,3,20),
  )
}
```
* If and else
  * Variables declared inside an `if` short statement are also available inside and if the `else` blocks.
```go
package main

import ( 
  "fmt"
  "math"
)

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x,n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't use v here, though 
  return lim
}

func main() {
  fmt.Println(
    pow(3,2,10),
    pow(3,3,20),
   )
}
```

## Switch 
* A `switch` statement is a shorter way to write a sequence of `if - else` statements.
* It runs the first case whose value is equal to the condition expression.
* Go only runs the selected case, not all the cases that follow.
```go
package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "Linux":
    fmt.Println("Linux.")
  default:
    // freebsd, openbsd
    fmt,Printf("%s.\n", os)
  }
}
```
* Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
* Switch without a condition is the same as switch true.
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

## Defer
* A defer statement defers the execution of a function until the surrounding function returns.
* The deferred call's arguments are evaluated immediately, but the function call is not executed util the surrounding function returns.
```go
package main 

import "fmt"

func main() {
  defer fmt.Println("world")

  fmt.Println("hello")
}
```
* Deffered function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
```go
package main

import "fmt"

func main(){
  fmt.Println("counting")
  
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("done")
}
```
# Pointers
* A pointer holds the memory address of a value.
  * The type `*T` is a pointer to a `T` value. its zero value is `nil`.
  ```
  var p *int
  ```
  * The `&` operator generates a pointer to its operand.
  ```
  i := 42
  p = &i
  ```
  * The `*` operator denotes the pointer's underlying value.
  ```
  fmt.Println(*p) // read i through the pointer p
  *p = 21 // set i through the pointer p
  ```
  * This is known as "dereferencing" or "indirecting".
  * Unlike C, Go has no pointer arithmetic.
```go
package main

import "fmt"

func main() {
  i, j := 42, 2017
  p := &i // point to i
  fmt.Println(*p) // read i through the pointer 
  *p = 21 // set i through the pointer
  fmt.Println(i) // see the new value of i
  p = &j // point to j
  *p = *p / 37 // divide j through the pointer 
  fmt.Println(j)
}
```

## Structs 
* A `struct` is a collection of fields.
* Struct fields are accessed using a dot.
* Pointers to structs:
  * Struct fields can be accessed through a struct pointer
  * To access the field `X` of a struct when we have the struct pointer `p` we can write `(*p).X`.
  * However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  fmt.Println(Vertex(1,2))
  v := Vertex(1,2)
  v.X = 4
  fmt.Println(v.X)
  p := &v
  p.X = 1e9
  fmt.Println(v)
}
```
* Struct Literals
  * A struct literals denotes a newly allocated struct value by listing the values of its fields.
  * You can list just a subset of fields by suing the `Name`: syntax. (And the order of named fields is irrelevant.)
  * The sepcial prefix `&` returns a pointer to the struct value.

```go
package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  v1 = Vertex{1,2} // has type vertex
  v2 = Vertex{X:1} // Y:0 is implicit
  v3 = Vertex{} // X:0 and Y: 0
  p = &Vertex{1,2} // has type *Vertex
)

func main() {
  fmt.Println(v1,p,v2,v3)
}
```
# Arrays
* The type `[n]T` is an array of `n` values of type `T`.
  * The expression
  ```
  var a [10]int
  ```
  * declares a variable `a` as an array of ten integers.
  * An array's length is part of its type, so arrays cannot be resized. Thise seems limiting, but don't worry; Go provides a convenient way of working with arrays.
```go
package main

import "fmt"

func main() {
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  primes := [6]int{2,3,5,7,11,13}
  fmt.Println(primes)
}
```
* **Slices**:
  * An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
  * The type `[]T` is a slice with elements of type `T`.
  * A slice is formed by specifying two indices, a low and high bound, separated by a colon:
  ```
  a[low : high]
  ```
  * This selects a half-open range which includes the first element, but excludes the last one.
  * The following expression creates a slice which includes elements 1 through 3 of `a`:
  ```
  a[1:4]
  ```
```go
package main

import "fmt"

func main() {
  primes := [6]int{2,3,5,7,11,13}
  
  var s []int = primes[1:4]
  fmt.Println(s)
}
```
* Slices are like references to arrays
  * A slice does not store any data, it just describes a section of an underlying array.
  * Changing the elements of a slice modifies the corresponding elements of its underlying array.
  * Other slices that share the same underlying array will see those changes.
```go
package main

import "fmt"

func main() {
  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
    }
  fmt.Println(names)
  a := names[0:2]
  b := names[1:3]
  fmt.Println(a,b)

  b[0] = "XXX"
  fmt.Println(a,b)
  fmt.Println(names)
}
```
* Slice literals
  * A slice literal is like an array literal without the length.
  * This is an array literal:
  ```
  [3]bool{true, true, false}
  ```
  * And this creates the same array as above, then builds a slice that references it:
  ```
  []bool{true, true, false}
  ```
```go
package main

import "fmt"

func main() {
  q := []int{2,3,5,7,11,13}
  fmt.Println(q)
  
  r := []bool{true, false, true, true, false, true}
  fmt.Println(r)

  s := []struct {
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
  fmt.Println(s)
}
```
* Slice defaults
  * When slicing, you may omit the high or low bounds to use their defaults instead.
  * The default is zero for the low bound and the length of the slice for the high bound.
  * For the array
  ```
  var a [10]int
  ```
  * these slice expressions are equivalent:
  ```
  a[0:10]
  a[:10]
  a[0:]
  a[:]
  ```
* Slice length and capacity
  * A slice has both a **length** and a **capacity**.
  * The length of a slice is the number of elements it contains
  * The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
  * The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`.
```go
package main

import "fmt"

func main() {
  s := []int{2,3,5,7,11,13}
  printSlice(s)

  //slice the slice to give it zero length
  s = s[:0]
  printSlice(s)
  
  // Extend its length
  s = s[:4]
  printSlice(s)

  // Drop its first two values.
  s = s[2:]
  printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```
* The zero value of a slice is `nil`.
* A nil slice has a length and capacity of 0 and has no underlying array.
* Creating a slice with **make**
  * Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.
  * The `make` function allocates a zeroed array and returns a slice that refers to that array:
   ```
   a := make([]int,5) // len(a) =5 
   ```
   * To specify a capacity, pass a third argument to `make`:
   ```
   b := make([]int, 0, 5) // len(b)=0, cap(b)=5
   b = b[:cap(b)] //len(b)=cap(b)=5
   b = b[1:] //len(b)=4, cap(b)=4
   ```
```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```
* Slices of slices
  * Slices an contain any type, including other slices.
```go 
package main

import (
  "fmt"
  "strings"
)
 
func main() {
  //create a tic-tac-toe board
  board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }
  // The players take turns
  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"
  
  for i :=0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
}
```
* Appending to a slice
  * It is common to append new elements to a slice, and so Go provides a built-in append function
  ```
  func append(s []T, vs ...T) []T
  ```
  * The first parameter `s` of `append` is a slice of type T, and the rest are T values to append to the slice.
  * The resulting value of `append` is a slice containing all the elements of the original slice plus the provided values.
  * If the backing array of `s` is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
```go 
package main

import "fmt"

func main() {
  var s []int
  printSlce(s)

  //append works on nil slices 
  s = append(s,0)
  printSlice(s)

  // The slice grows as needed
  s = append(s,1)
  printSlice(s)

  // We can add more than one elements
  s = append(s, 2,3,4)
  printSlice(s)
}

func printSlice(s []int) {
  fmt.Println("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```
* **Range**:
  * The `range` form of the `for` loop iterates over a slie or map.
  * When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index. 

```go
package main

import "fmt"

var pow = []int{1,2,4,8,16,32,64,128}

func main() {
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }
}
```
* Range continued
  * You can skip the index or value by assigning to `_`.
  ```
  for i, _ := range pow
  for _, value := range pow
  ```
  * if you only want to the index, you can omit the second variable.
  ```
  for i := range pow
  ```
```go
package main

import "fmt"

func main() {
  pow := make([]int, 10)
  for i := range pow {
    pow[i] = 1 << unit(i) // == 2**i
  }
  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }
}
```

## Maps
* A map maps keys to values
* The zero value of a map is `nil`.
* A `nil` map has no keys nor can keys be added.
* The `make` function returns a map of the given type, initialized and ready foor use.
```go
package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m map[string]Vertex

func main() {
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{40.6833, -74.39967}
  fmt.Println(m["Bell Labs"])
}
```
* Map literals are like struct literals, but the keys are required.
```go 
package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex {
  "Bell Labs": Vertex{
    40.68433, -74.39967,},
  "Google": Vertex{
     37.42202, -122.08408,},
}

func main() {
  fmt.Println(m)
}
```
* If the top-level type is just a type name, you can omit it from the elements of the literal. 
```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```
* Mutating Maps
  * Insert or update an element in map `m`:
  ```
  m[key] = elem
  ```
  * Retrieve an element:
  ```
  element = m[key]
  ```
  * Delete an element:
  ```
  delete(m,key)
  ```
  * Test that a key is present with a two-value assigment:
  ```
  elem, ok = m[key]
  ```
  * if `key` is in `m`, `ok` is `true`. If not, `ok` is `false`.
  * If `key` is not in the map, then `elem` is the zero value for the map's element type. 
  * Note: if `elem` or `ok` have not been declared you could use a short declaration form:
  ```
  elem, ok := m[key]
  ```

```go 
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
```
## Function values
* Functions are values too. They can be passed around just like other values. Function values may be used as function arguments and return values.
```go
package main

import (
  "fmt"
  "math"
)
  
func compute(fn func(float64, float64) float64) float64 {
  return fn(3,4)
}

func main() {
  hypot := func(x,y float64) float64 {
    return math.Sqtr(x*x + y*y)
  }
  fmt.Println(hypot(5,12))

  fmt.Println(compute(hypot))
  fmt.Println(compute(math.Pow))
}
```
* Function closures
  * Go functions may be closures. A closure is a function value that references varaibles from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables. 
```go
package main

import "fmt"

func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func main() {
  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println( pos(i), neg(-2*i))
  }
}
```
## Methods
* Go does not have classes. However, you can define methods on types.
* A method is a function with a special *receiver* argument.
* The receiver appears in its own argument list between the `func` keyword and the method name.

``go
package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertext) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3,4}
  fmt.Println(v.Abs())
}
```
* Methods are functions: a method is just a function with a receiver argument.
```go
package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func Abs(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3,4}
  fmt.Println(Abs(v))
}
```
* You can declare a method on non-struct types, too.
```go 
package main

import (
  "fmt"
  "math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func main() {
  f := MyFloat(-math.Sqrt(2))
  fmt.Println(f.Abs())
}
```
* Pointer receivers 
  * You can declare methods with pointer recievers.
  * This means the receiver types has the literal syntax `*T` for some type `T`. ( Also, `T` cannot itself be a pointer such as `*int`.)
  * Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiever, pointer receivers are more common than value recievers.
  * With a value reciever, the method operates on a copy of the original value.
```go 
package main

import ( 
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale() float64 {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3,4}
  v.Scale(10)
  fmt.Println(v.Abs())
}
```
* The functions with a pointer argument must take a pointer.
* While methods with pointer receivers take either a value or a pointer as the receiver whne they are called.
```go 
package main

import "fmt"

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3,4}
  v.Scale(2)
  ScaleFunc(&v, 10)

  p := &Vertex{4,3}
  p.Scale(3)
  ScaleFunc(p,8)
  
  fmt.Println(v, p)
}
```
* Functions that take a value argument must take a value of that specific type
* While methods with value receivers take either a value or a pointer as the receiver when they are clalled
```go
package main

import ( 
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3,4}
  fmt.Println(v.Abs())
  fmt.Println(AbsFunc(v))

  p := &Vertex{3,4}
  fmt.Println(p.Abs())
  fmt.Println(AbsFunc(*p))
}
```
* Choosing a value or pointer receiver:
  * There are two reasons to use a pointer receiver.
  * The first is so that the method can modify the value that its receiver points to. 
  * The second is to avoid copying the value on eac method call. This can be more efficient if the receiver is a larger struct.
## Interfaces

* An *interface type* is defined as a set of method signatures. 
* A value of interface type can holdany value that implements those methods.
```go
package main

import (
  "fmt"
  "math"
)

type Abser interface {
  Abs() float64
}

func main() {
  var a Abser
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3,4}
  a = f // a MyFloat implements Abser
  a = &v // a *Vertex implements Abser

  fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```



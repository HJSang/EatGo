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


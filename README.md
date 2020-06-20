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

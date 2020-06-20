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

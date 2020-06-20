# GO notes in one page

The goal os this note is to write all important go notes in one page for quick search and revist.

The source is from thsi [A Tour of Go](https://tour.golang.org/welcome/1).

## Packages
* Every Go program is made up of packages.
* Programs start running in package **main**.
* This program is using the packages with import paths `"fmt"` and `"math/rand"`.
* By convention, the package name is the same as the last element of the import path. For instance, the `"math/ran"` package comprises files that begin with the statement `package rand`.

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

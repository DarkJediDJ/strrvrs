# String reverse + char counter golang module.
This module provides 2 functions:
    
* string reverse
* char counter
    
## Usage
    
* import github.com/DarkJediDJ/reverse
* call required function
   
### Examples

String reverse:
 ```go 
package main

import (
    "fmt"
    "github.com/DarkJediDJ/strrvrs"
)
    
func main() {
    str := strrvrs.Reverse("Some text")
    fmt.Println(str)
    // Output:
    // txet emoS
}
```
Char counter:
```go
package main

import (
    "fmt"
    "github.com/DarkJediDJ/strrvrs"
)

func main() {
    n := strrvrs.Calculator("Example")
    fmt.Println(n)
    // Output:
    // 7
}

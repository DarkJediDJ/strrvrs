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
)
    
func main() {
    str := reverse.Reverse("Some text")
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
)

func main() {
    n := reverse.Calculator("Example")
    fmt.Println(n)
    // Output:
    // 7
}

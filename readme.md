# rate

## example
```
import (
	"fmt"
	"time"

	"github.com/pipiBRH/rate"
)

func main() {
    limiter := NewLimiter(5, time.Minute)

    for i := 0; i < 10; i++ {
        fmt.Println(limiter.Allow())
    }

    // Output:
    // true
    // true
    // true
    // true
    // true
    // false
    // false
    // false
    // false
    // false
}
```
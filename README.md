# trylock

the code is so easy.

Only add TryLock and IsLocked function base on sync.Mutex. A lot of times, I don't want to block when Mutex is locked.

### install

```
go get github.com/rfyiamcool/trylock
```

### example

```
package main

import (
	"fmt"

	"github.com/rfyiamcool/trylock"
)


func main() {
	m := trylock.NewMutex()
	if m.TryLock() {
		fmt.Println("get lock")
	}
	m.IsLocked()
}
```

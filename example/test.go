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


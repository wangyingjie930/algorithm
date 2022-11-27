package 并发编程

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type Person struct {
	Name string
}

func initPool() {
	pool = &sync.Pool {
		New: func()interface{} {
			fmt.Println("Creating a new Person")
			return new(Person)
		},
	}
}

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	pkg1 "gobaco.com/pkg/pkg1"
	pkg2_pkg1 "gobaco.com/pkg/pkg2/pkg1"
)

var wg sync.WaitGroup

func main() {
	pkg2_pkg1.Rand()
	pkg1.Rand()
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("OS\t", runtime.GOROOT())
	fmt.Println("OS\t", runtime.NumGoroutine())
	fmt.Println("OS\t", runtime.NumCPU())

	wg.Add(1)
	go foo()
	bar()
	wg.Wait()
}

func Rand() {
	rand.Seed(time.Now().Unix())
	fmt.Println("pkg3 My favorite number is", rand.Intn(10))
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

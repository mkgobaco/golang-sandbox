package pkg1

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	Rand()
}

func Rand() {
	rand.Seed(time.Now().Unix())
	fmt.Println("pkg2/pkg1 My favorite number is", rand.Intn(10))
}

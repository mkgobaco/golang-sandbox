package main

import (
	"fmt"
	module1_pkg1 "module1/pkg/pkg1"
	module2_pkg1 "module2/pkg/pkg1"
)

func main() {
	fmt.Println(module1_pkg1.Module1Pkg1File1())
	fmt.Println(module2_pkg1.Module2Pkg1File1())
}

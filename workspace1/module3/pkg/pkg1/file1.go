package pkg1

import (
	"fmt"
	"module1/pkg/pkg2"
)

func Module3Pkg1File1() string {
	fmt.Println("Imported " + pkg2.Module1Pkg2File1())
	fmt.Println("Imported " + pkg2.Module1Pkg2File2())
	return "pkg1.file1"
}

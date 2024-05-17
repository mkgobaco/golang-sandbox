package pkg1

import (
	"fmt"
	module1_pkg2 "module1/pkg/pkg2"
	module2_pkg1 "module2/pkg/pkg1"
	module2_pkg2 "module2/pkg/pkg2"
)

func Module1Pkg1File1() string {
	fmt.Println("Imported " + module1_pkg2.Module1Pkg2File1())
	fmt.Println("Imported " + module1_pkg2.Module1Pkg2File2())
	fmt.Println("Imported " + module2_pkg1.Module2Pkg1File1())
	fmt.Println("Imported " + module2_pkg1.Module2Pkg1File2())
	fmt.Println("Imported " + module2_pkg2.Module2Pkg2File1())
	fmt.Println("Imported " + module2_pkg2.Module2Pkg2File2())

	return "module1.pkg1.file1"
}

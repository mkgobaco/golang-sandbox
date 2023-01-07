package pkg1

import (
	"fmt"
	module3_pkg1 "module3/pkg/pkg1"
	module3_pkg2 "module3/pkg/pkg2"
)

func Module2Pkg1File1() string {
	fmt.Println("Imported " + module3_pkg1.Module3Pkg1File1())
	fmt.Println("Imported " + module3_pkg2.Module3Pkg2File2())
	return "module2.pkg1.file1"
}

func main() {
	Module2Pkg1File1()
	Module2Pkg1File2()
}

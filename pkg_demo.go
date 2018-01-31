package main

import (
	"fmt"
	pkgDemo "presentation_demos/pkg_demos"
	)

func main() {

	// Exported functions
	fmt.Println(pkgDemo.Exported())
	fmt.Println(pkgDemo.Exported2())

	// We can only use named intialisation with exported members
	a := pkgDemo.ExportedType{ExportedMember: 50}

	// We can't construct this member even though it is exported,
	// this is because it's type isn't exported
	b := pkgDemo.ExportedType2{}

	// We still get access to the member though
	fmt.Println(b.X)

	fmt.Println(a, b)

}

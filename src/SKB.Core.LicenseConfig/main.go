// Loads a a centralized LICENSE that is managed by the projects
package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed LICENSE.rst
var LicenseContents string

var fileName string = "LICENSE.rst"

func main() {
	err := os.WriteFile(fileName, []byte(LicenseContents), 0644)

	if err != nil {
		fmt.Println("FAILURE: Failed to inject the LICENSE.rst")
		fmt.Println(fmt.Errorf("ERROR: %e", err))
	}

	fmt.Println("SUCCESS: Loaded the LICENCE.rst")
}

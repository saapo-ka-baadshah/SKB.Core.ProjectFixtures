// / Loads a centralized .editorconfig file for the
package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed .editorconfig
var editorConfigContents string

// File name
var FileName string = ".editorconfig"

func main() {
	err := os.WriteFile(FileName, []byte(editorConfigContents), 0644)

	// Show failure
	if err != nil {
		fmt.Println("FAILURE: Failed to inject the .editorconfig")
		fmt.Println(fmt.Errorf("ERROR: %e", err))
	}

	fmt.Println("SUCCESS: Loaded the .editorconfig")
}

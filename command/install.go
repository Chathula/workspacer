package command

import "fmt"

// Install all
func Install(file string) error {
	fmt.Printf("install all file %s\n", file)
	return nil
}

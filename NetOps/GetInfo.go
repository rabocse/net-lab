// Get information from the file "file.txt".

package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.Stat("file.txt")

	f1 := fmt.Println

	f1("File Name: ", file.Name())
	f1("Size in Bytes: ", file.Size())
	f1("Last Modified: ", file.ModTime())
	f1("Is a directory: ", file.IsDir())

	ff := fmt.Printf
	ff("Permission 9-bit: %s\n", file.Mode())
	ff("Permission 3-digit:  %o\n", file.Mode())
	ff("Permission 4-digit: %04o\n", file.Mode())

}

// It creates a file called "test.txt" within the folder ".testdir".
// Then, those ared deleted with the "Remove" and "RemoveAll".

package main

import (
	"os"
	"path/filepath"
)

func main() {

	dpath := ".testdir"
	fname := "test.txt"

	_ = os.MkdirAll(dpath, 0777)

	fpath := filepath.Join(dpath, fname)

	file, _ := os.Create(fpath)

	file.Close()

	_ = os.Remove(fpath)

	_ = os.RemoveAll(dpath)

}

// What is the difference between "Remove" and "RemoveAll" ? Remove dependencies ? what does it mean ?

// filepath package

// MkdirAll from "os" package

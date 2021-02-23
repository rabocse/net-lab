package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	scan := func(
		path string, i os.FileInfo, _ error) error {
		fmt.Println(i.IsDir(), path)
		return nil
	}
	_ = filepath.Walk(".", scan)
}

/*

C:\Users\aleescob\code\go\NetOps
λ go run FileTreeWalk.go
true .
false AddRemove.go
false FileTreeWalk.go
false GetInfo.go
false ListDir.go
false OpenClose.go
true TEST
true TEST\LEVEL2
false TEST\test.txt
false file.txt
false file1.txt

C:\Users\aleescob\code\go\NetOps
λ
C:\Users\aleescob\code\go\NetOps
λ
C:\Users\aleescob\code\go\NetOps
λ tree
Folder PATH listing for volume OSDisk
Volume serial number is B838-2AB6
C:.
└───TEST
    └───LEVEL2

C:\Users\aleescob\code\go\NetOps
λ


*/

// It list the current directory but not depeer levels. No recursively.

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	files, _ := ioutil.ReadDir(".")

	for _, file := range files {
		fmt.Println(file.Name(), file.ModTime())

	}

}

/*



C:\Users\aleescob\code\go\NetOps
λ ls -al
total 14
drwxr-xr-x 1 aleescob 1049089   0 Nov  6 16:22 ./
drwxr-xr-x 1 aleescob 1049089   0 Nov  6 13:03 ../
-rw-r--r-- 1 aleescob 1049089 596 Nov  6 13:53 AddRemove.go
-rw-r--r-- 1 aleescob 1049089  48 Nov  6 13:15 file.txt
-rw-r--r-- 1 aleescob 1049089 216 Nov  6 13:51 file1.txt
-rw-r--r-- 1 aleescob 1049089 482 Nov  6 16:17 GetInfo.go
-rw-r--r-- 1 aleescob 1049089 192 Nov  6 16:25 ListDir.go
-rw-r--r-- 1 aleescob 1049089 368 Nov  6 13:27 OpenClose.go
drwxr-xr-x 1 aleescob 1049089   0 Nov  6 16:23 TEST/



C:\Users\aleescob\code\go\NetOps
λ tree
Folder PATH listing for volume OSDisk
Volume serial number is B838-2AB6
C:.
└───TEST
    └───LEVEL2




C:\Users\aleescob\code\go\NetOps
λ
C:\Users\aleescob\code\go\NetOps
λ go run ListDir.go
AddRemove.go 2020-11-06 13:53:02.8139463 +0100 CET
GetInfo.go 2020-11-06 16:17:32.6316358 +0100 CET
ListDir.go 2020-11-06 16:25:26.8051519 +0100 CET
OpenClose.go 2020-11-06 13:27:06.9943946 +0100 CET
TEST 2020-11-06 16:23:02.5081855 +0100 CET
file.txt 2020-11-06 13:15:56.3596373 +0100 CET
file1.txt 2020-11-06 13:51:36.9370118 +0100 CET



*/

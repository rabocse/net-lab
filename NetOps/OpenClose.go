// Creates and append data to the file. If the file already exist, then it appends.
// It also closes the file.

package main

import "os"

func main() {

	file, _ := os.OpenFile("file1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	file.Write([]byte("append data\n"))
	file.Close()

}

// What is 0644, is not the privilege defined by "O_WRONLY"

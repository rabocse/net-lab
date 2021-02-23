package main

import (
	"bufio"
	"os"
)

func main() {

	lines := []string{"1- Colombia",
		"2- Poland",
		"3- Mexico",
		"4- Portugal",
		"5- Bulgaria",
		"6- India",
		"7- US"}

	file, _ := os.OpenFile("./CxCenters.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, _ = writer.WriteString(line + "\n")
	}

	writer.Flush()

	file.Close()

}

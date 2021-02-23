// Accepts a CVS file as paramenter and prints it.

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings" // The function "Join" will be used to covernt from Slice to String.
)

type Row struct {
	ColA, ColB, ColC string
}

func main() {

	argumentFile := os.Args[1:] // Args hold the command-line arguments, starting with the program name [0]. That is why [1:] is needed.

	argument := strings.Join(argumentFile, "") //  Join receives a "string slice" and concatenates its elements separated by whather is specified within the double quotes (""). Which means no separator in this case.

	file, _ := os.Open(argument) // Open can receive the argument as a string at this point.

	reader := csv.NewReader(file)

	rows := []Row{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		rows = append(rows, Row{
			ColA: row[0], ColB: row[1], ColC: row[2],
		})
	}

	for _, row := range rows {
		fmt.Printf("%s -- %s -- %s\n",
			row.ColA, row.ColB, row.ColC)
	}

}

/*

This is how it looks:

C:\Users\aleescob\code\go\NetOps
λ go run ReadingCSV.go
DEVICE -- IP ADDRESS -- LOCATION
FMC -- 192.168.1.10 -- BRU
FTD -- 192.168.1.11 -- BRU

C:\Users\aleescob\code\go\NetOps



++ The file was saved as CSV (Comma Delimited). When it was saved as CVS UTF-8, it did not work. It was not compiling:


C:\Users\aleescob\code\go\NetOps
λ go run ReadingCSV.go
panic: runtime error: index out of range [1] with length 1

goroutine 1 [running]:
main.main()
        C:/Users/aleescob/code/go/NetOps/ReadingCSV.go:29 +0x619
exit status 2

C:\Users\aleescob\code\go\NetOps


*/

// The program validates an IP address that is passed as command parameter.
// ./ParserIP <ip-address>

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 { // Args[0] is the name of the script and Args[1] is the first argument (the IP address). So, this validates the script is called and that the IP address is passed.
		fmt.Println("Please, enter proper amount of arguments. Example: ./ParseIP 172.16.32.100")
		return
	}

	ip := net.ParseIP(os.Args[1])

	if ip != nil {

		fmt.Printf("%v OK\n", ip)
	} else {
		fmt.Println("Could not parse the address. Possible bad address")
	}

}

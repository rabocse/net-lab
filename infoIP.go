// The program provides information about the passed IP.

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {

		fmt.Println("Please, enter proper amount of arguments. Example: ./ParseIP 172.16.32.100")
		os.Exit(1)
	}

	ip := net.ParseIP(os.Args[1])
	if ip == nil {
		fmt.Println("Unable to parse IP address")
		fmt.Println("Address should use IPv4 dot-notation or IPv6 colon-notation")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Printf("IP Address: 	 %s\n", ip)
	fmt.Printf("Default Mask: 	 %s\n", net.IP(ip.DefaultMask()))
	fmt.Printf("Loopback:		 %t\n", ip.IsLoopback())
	fmt.Println("Unicast:")
	fmt.Printf("\tGlobal:		 	 %t\n", ip.IsGlobalUnicast())
	fmt.Printf("\tLink:		 	 %t\n", ip.IsLinkLocalUnicast())
	fmt.Println("Multicast:")
	fmt.Printf("\tGlobal:		 	 %t\n", ip.IsMulticast())
	fmt.Printf("\tInterface:		 %t\n", ip.IsInterfaceLocalMulticast())
	fmt.Printf("\tLink: 			 %t\n", ip.IsLinkLocalMulticast())
	fmt.Println()
}

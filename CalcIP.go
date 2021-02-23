// IP Calculator program.

package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"os"
)

var cidr string

func init() {
	flag.StringVar(&cidr, "c", "", "the CIDR address")
}

func main() {

	flag.Parse()

	if cidr == "" {
		fmt.Println("CIDR address missing")
		os.Exit(1)
	}

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {

		fmt.Println("failed parsing CIDR address: ", err)
		os.Exit(1)
	}

	// Given IPv4 block 192.168.100.14/24
	// The followings uses IPNet to get:
	// - The routing address for the subnet (i.e. 192.168.100.0)
	// - one-bits of the network mask (24 out of 32 total)
	// - The subnetmask (i.e. 255.255.255.0)
	// - Total hosts in the network (2 ^(host identifer bits) or 2^8)
	// - Wildcard the inverse of subnet mask (i.e. 0.0.0.255)
	// - The maximum address of the subnet (i.e. 192.168.100.255)

	ones, totalBits := ipnet.Mask.Size()
	size := totalBits - ones                   // usable bits
	totalHosts := math.Pow(2, float64(size))   // 2^size
	wildcardIP := wildcard(net.IP(ipnet.Mask)) // Manually created function. Check below after main program.
	last := lastIP(ip, net.IPMask(wildcardIP)) // Manually created function. Check below after main program.

	fmt.Println()
	fmt.Printf("CIDR Block: 	%s\n", cidr)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("CIDR Block: 	%s\n", cidr)
	fmt.Printf("Network: 		%s\n", ipnet.IP)
	fmt.Printf("IP Range: 		%s -  %s \n", ipnet.IP, last)
	fmt.Printf("Total Hosts:	%0.0f\n", totalHosts)
	fmt.Printf("Netmask:		%s\n", net.IP(ipnet.Mask))
	fmt.Printf("Wildcard Mask:	%s\n", wildcardIP)
	fmt.Println()

}

// These are the manually created functions. (Outside main)

func wildcard(mask net.IP) net.IP {
	var ipVal net.IP
	for _, octec := range mask {

		ipVal = append(ipVal, ^octec)
	}
	return ipVal
}

func lastIP(ip net.IP, mask net.IPMask) net.IP {
	ipIn := ip.To4()
	if ipIn == nil {
		ipIn = ip.To16()
		if ipIn == nil {
			return nil
		}
	}

	var ipVal net.IP

	for i, octet := range ipIn {
		ipVal = append(ipVal, octet|mask[i])
	}
	return ipVal
}

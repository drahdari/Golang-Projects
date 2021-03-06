// this piece of program  identifyies the information regarding the Network Interface Card of a System
// The program is written in a way to be easily used for Relational Database Data insertion.
// In the program, I used map and aray data structures


package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	NIC_Retrival()
}
func NIC_Retrival() {
	interfaces, Error := net.Interfaces()
	if Error != nil {
		fmt.Errorf("Error in the Operation", Error.Error())
		return
	} else {
		var MAP_Interfaces = map[int]map[string]string{}
		// defining map data structure so data can be categorized 
		for _, IF := range interfaces {
			// seperating and fetching Ipv4 and IPv6 addressess
			var IPs [2]string = IP_Retrival(IF)
			MAP_Interfaces[IF.Index] = map[string]string{}
			MAP_Interfaces[IF.Index]["Name"] = IF.Name
			MAP_Interfaces[IF.Index]["IPv4"] = IPs[0]
			MAP_Interfaces[IF.Index]["IPv6"] = IPs[1]
			MAP_Interfaces[IF.Index]["Status"] = strings.Split((IF.Flags.String()), "|")[0]
			MAP_Interfaces[IF.Index]["MAC"] = "N/A"
			var MAC string = IF.HardwareAddr.String()
			if len(MAC) > 0 {
				MAP_Interfaces[IF.Index]["MAC"] = MAC
			}
		}
		//fmt.Println(len(MAP_Interfaces))
		Print_Interface(MAP_Interfaces)
	}

}
func IP_Retrival(IF net.Interface) [2]string {
	var addrs, _ = IF.Addrs()
	var IPv4 string = "N/A"
	var IPv6 string = "N/A"
	if len(addrs) > 0 {
		IPv4 = addrs[0].String()
		if len(addrs) == 2 {
			IPv6 = addrs[1].String()
		}
	}
	var IPs [2]string = [2]string{IPv4, IPv6}
	return IPs
}

func Print_Interface(MAP_Interfaces map[int]map[string]string) {
	var i int
	for i = 0; i < len(MAP_Interfaces); i++ {
		fmt.Printf("Name:%v IPv4:%v IPv6:%v MAC:%v Status:%v\n", MAP_Interfaces[i]["Name"], MAP_Interfaces[i]["IPv4"], MAP_Interfaces[i]["IPv6"], MAP_Interfaces[i]["MAC"], MAP_Interfaces[i]["Status"])
	}

}

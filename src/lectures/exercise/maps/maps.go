//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServerInfo(servers map[string]int) {
	fmt.Println("\n There are", len(servers), "servers")

	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}

	fmt.Println("Online servers:", stats[Online])
	fmt.Println("Offline servers:", stats[Offline])
	fmt.Println("Maintenance servers:", stats[Maintenance])
	fmt.Println("Retired servers:", stats[Retired])
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatus := make(map[string]int)
	for i := 0; i < len(servers); i++ {
		server := servers[i]
		serverStatus[server] = Online
	}

	displayServerInfo(serverStatus)
	fmt.Println()
	
	serverStatus["darkstart"] = Retired
	serverStatus["aiur"] = Offline
	displayServerInfo(serverStatus)
	fmt.Println()

	for name := range serverStatus {
		serverStatus[name] = Maintenance
	}
	
	displayServerInfo(serverStatus)
}

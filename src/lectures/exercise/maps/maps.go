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

func displayServerInfo(serverMap map[string]int) {
	var totalServers,
		onlineServers,
		offlineServers,
		maintenanceServers,
		retiredServers int

	for _, status := range serverMap {
		totalServers += 1
		switch status {
		case Online:
			onlineServers += 1
		case Offline:
			offlineServers += 1
		case Maintenance:
			maintenanceServers += 1
		case Retired:
			retiredServers += 1
		}
	}

	fmt.Println("Total servers:", totalServers)
	fmt.Println("Online servers:", onlineServers)
	fmt.Println("Offline servers:", offlineServers)
	fmt.Println("Maintenance servers:", maintenanceServers)
	fmt.Println("Retired servers:", retiredServers)
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverMap := make(map[string]int)
	for i := 0; i < len(servers); i++ {
		server := servers[i]
		serverMap[server] = Online
	}

	displayServerInfo(serverMap)
	fmt.Println()
	
	serverMap["darkstart"] = Retired
	serverMap["aiur"] = Offline
	displayServerInfo(serverMap)
	fmt.Println()
	
	for name := range serverMap {
		serverMap[name] = Maintenance
	}
	
	displayServerInfo(serverMap)
}

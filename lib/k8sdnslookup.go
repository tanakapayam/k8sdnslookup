package k8sdnslookup

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func Lookup() {
	containerName, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	appName := strings.Split(
		containerName,
		"-",
	)[0]

	hostname := "app." + appName
	protocol := "tcp"

	services := []string{
		"json-port",
		"grpc-port",
		appName + "-json-port",
		appName + "-grpc-port",
		appName + "-http-port",
	}

	fmt.Printf("appName:       %s\n", appName)
	fmt.Printf("Hostname:      %s\n\n", hostname)

	for _, service := range services {
		cname, addresses, err := net.LookupSRV(service, protocol, hostname)

		if err == nil {
			fmt.Printf("Service CNAME: %s\n", cname)

			for i := 0; i < len(addresses); i++ {
				fmt.Printf("Target:        %s\n", addresses[i].Target)
				fmt.Printf("Port:          %d\n", addresses[i].Port)
			}

			fmt.Println()
		}
	}
}

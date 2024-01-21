package consensus

import (
	"fmt"

	externalip "github.com/glendc/go-external-ip"
)

func MyIp() {

	consensus := externalip.DefaultConsensus(nil, nil)
	// todo : allow proto v
	err := consensus.UseIPProtocol(7)

	if err != nil {
		fmt.Println("Protocol not supported, defaulting to ipv4...")
	}

	ipAddress, err := consensus.ExternalIP()

	if err == nil {
		fmt.Println("Your external IP: " + ipAddress.String())
	}

}

package consensus

import (
	"fmt"
	"net"

	externalip "github.com/glendc/go-external-ip"
)

func MyIp(proto int) (net.IP, error) {

	consensus := externalip.DefaultConsensus(nil, nil)
	err := consensus.UseIPProtocol(uint(proto))

	if err != nil {
		fmt.Println("Protocol not supported, defaulting to ipv4...")
	}

	ipAddress, err := consensus.ExternalIP()

	if err == nil {
		return ipAddress, nil
	}

	return nil, err
}

package utils

import "net"

func IsLANAddress(host string) bool {
	lanBlocks := []string{
		"192.168.0.0/16",
		"172.16.0.0/12",
		"10.0.0.0/8",
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		return false
	}
	for _, ip := range ips {
		for _, block := range lanBlocks {
			_, subnet, _ := net.ParseCIDR(block)
			if subnet.Contains(ip) {
				return true
			}
		}
	}
	return false
}

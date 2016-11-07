package util

import (
	"net"
	"strconv"
	"strings"
)

func Ip2Long(ip net.IP) int64 {
	bits := strings.Split(ip.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

func Long2Ip(n int64) net.IP {
	var mask int64 = 255
	var ips = make([]string, 0, 4)
	ips = append(ips, strconv.FormatInt((n>>24)&mask, 10))
	ips = append(ips, strconv.FormatInt((n>>16)&mask, 10))
	ips = append(ips, strconv.FormatInt((n>>8)&mask, 10))
	ips = append(ips, strconv.FormatInt(n&mask, 10))
	ip := net.ParseIP(strings.Join(ips, "."))
	return ip
}

func InNet(ip net.IP, netBit uint8, checkIp net.IP) bool {
	ipl := Ip2Long(ip)
	cipl := Ip2Long(checkIp)
	if ipl>>(32-netBit) == cipl>>(32-netBit) {
		return true
	}
	return false
}

func IsPrivate(ip net.IP) bool {
	var pri = map[string]uint8{
		"10.0.0.0":    8,
		"172.16.0.0":  12,
		"192.168.0.0": 16,
	}
	for ips, b := range pri {
		ipp := net.ParseIP(ips)
		if InNet(ipp, b, ip) {
			return true
		}
	}
	return false
}

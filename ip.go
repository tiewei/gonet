package gonet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"net"
)

// Version returns the ip address version
func IPVersion(ip net.IP) int {
	if ip.To4() != nil {
		return 4
	}
	return 6
}

func IPCmp(one, another net.IP) (int, error) {
	if v1, v2 := IPVersion(one), IPVersion(another); v1 != v2 {
		return 0, fmt.Errorf(
			"can't compare IPv%d address %q with IPv%d address %q",
			v1, one.String(),
			v2, another.String())
	}
	return bytes.Compare(one.To16(), another.To16()), nil
}

func IPAdd(ip net.IP, num *big.Int) (net.IP, error) {
	if ip.To4() != nil {
		if num.IsInt64() {
			return doIPv4Move(ip, num.Int64())
		}
		return ip, fmt.Errorf("%s is outside IPv4 address boundary", num.String())
	}
	return doIPv6Move(ip, num)
}

func IPSub(ip net.IP, num *big.Int) (net.IP, error) {
	n := new(big.Int).Neg(num)
	return IPAdd(ip, n)
}

func doIPv6Move(ip net.IP, num *big.Int) (net.IP, error) {
	ipValue := new(big.Int).SetBytes(ip)
	result := new(big.Int).Add(ipValue, num)
	byteResult := result.Bytes()
	if result.Sign() < 0 || len(byteResult) > net.IPv6len {
		return ip, fmt.Errorf("result outside IPv6 address boundary")
	}
	buf := make([]byte, net.IPv6len-len(byteResult))
	buf = append(buf, byteResult...)
	return net.IP(buf), nil
}

func doIPv4Move(ip net.IP, num int64) (net.IP, error) {
	v4 := ip.To4()
	if v4 == nil {
		return nil, fmt.Errorf("ip %q is not an IPv4 address", ip.String())
	}
	value := binary.BigEndian.Uint32(v4)
	value64 := int64(value) + num
	if value64 >= 0 && value64 <= math.MaxUint32 {
		result := net.IPv4zero
		binary.BigEndian.PutUint32(result[12:16], uint32(value64))
		return result, nil
	}
	return nil, fmt.Errorf("result outside valid IPv4 address boundary")
}

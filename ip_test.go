package gonet

import (
	"math/big"
	"net"
	"testing"
)

func TestVersion(t *testing.T) {
	ip := net.ParseIP("192.168.1.1")
	if v := IPVersion(ip); v != 4 {
		t.Errorf("%q should be a v4 address, got %d", ip.String(), v)
	}
}

func TestIPv4Add(t *testing.T) {
	ip := net.ParseIP("192.168.1.1")
	if v, e := IPAdd(ip, big.NewInt(20)); e != nil || v.String() != "192.168.1.21" {
		if e != nil {
			t.Errorf("%q add 20 should be 192.168.1.21, got error: %s", ip.String(), e.Error())
		}
		t.Errorf("%q add 20 should be 192.168.1.21, got %q", ip.String(), v)
	}
}

func TestIPv6Add(t *testing.T) {
	ip := net.ParseIP("fe80::dead:beef")
	if v, e := IPAdd(ip, big.NewInt(20)); e != nil || v.String() != "fe80::dead:bf03" {
		if e != nil {
			t.Errorf("%q add 20 should be fe80::dead:bf03, got error: %s", ip.String(), e.Error())
		}
		t.Errorf("%q add 20 should be fe80::dead:bf03, got %q", ip.String(), v)
	}
}

func TestIPv4Sub(t *testing.T) {
	ip := net.ParseIP("192.168.1.21")
	if v, e := IPSub(ip, big.NewInt(20)); e != nil || v.String() != "192.168.1.1" {
		if e != nil {
			t.Errorf("%q sub 20 should be 192.168.1.1, got error: %s", ip.String(), e.Error())
		}
		t.Errorf("%q sub 20 should be 192.168.1.1, got %q", ip.String(), v)
	}
}

func TestIPv6Sub(t *testing.T) {
	ip := net.ParseIP("fe80::dead:bf03")
	if v, e := IPSub(ip, big.NewInt(20)); e != nil || v.String() != "fe80::dead:beef" {
		if e != nil {
			t.Errorf("%q sub 20 should be fe80::dead:beef, got error: %s", ip.String(), e.Error())
		}
		t.Errorf("%q sub 20 should be fe80::dead:beef, got %q", ip.String(), v)
	}
}

func TestIPv4Cmp(t *testing.T) {
	one := net.ParseIP("192.168.1.1")
	another := net.ParseIP("192.168.1.20")
	if v, e := IPCmp(one, another); e != nil || v >= 0 {
		if e != nil {
			t.Errorf("%q Cmp %q should be -1, got error: %s", one.String(), another.String(), e.Error())
		}
		t.Errorf("%q Cmp %q should be -1, got: %d", one.String(), another.String(), v)
	}
	if v, e := IPCmp(another, one); e != nil || v <= 0 {
		if e != nil {
			t.Errorf("%q Cmp %q should be 1, got error: %s", one.String(), another.String(), e.Error())
		}
		t.Errorf("%q Cmp %q should be 1, got: %d", one.String(), another.String(), v)
	}
	if v, e := IPCmp(one, one); e != nil || v != 0 {
		if e != nil {
			t.Errorf("%q Cmp %q should be 0, got error: %s", one.String(), another.String(), e.Error())
		}
		t.Errorf("%q Cmp %q should be 0, got: %d", one.String(), another.String(), v)
	}
}

func TestIPv6Cmp(t *testing.T) {
	one := net.ParseIP("fe80::dead:beef")
	another := net.ParseIP("fe80::dead:bf03")
	if v, e := IPCmp(one, another); e != nil || v >= 0 {
		if e != nil {
			t.Errorf("%q Cmp %q should be -1, got error: %s", one.String(), another.String(), e.Error())
		}
		t.Errorf("%q Cmp %q should be -1, got: %d", one.String(), another.String(), v)
	}
	if v, e := IPCmp(another, one); e != nil || v <= 0 {
		if e != nil {
			t.Errorf("%q Cmp %q should be 1, got error: %s", one.String(), another.String(), e.Error())
		}
		t.Errorf("%q Cmp %q should be 1, got: %d", one.String(), another.String(), v)
	}
	if v, e := IPCmp(one, one); e != nil || v != 0 {
		if e != nil {
			t.Errorf("%q Cmp %q should be 0, got error: %s", one.String(), another.String(), e.Error())
		}
		t.Errorf("%q Cmp %q should be 0, got: %d", one.String(), another.String(), v)
	}
}

func TestIPCmpDiffVersion(t *testing.T) {
	one := net.ParseIP("fe80::dead:beef")
	another := net.ParseIP("192.168.1.20")
	if v, e := IPCmp(one, another); e == nil {
		t.Errorf("%q Cmp %q should be error, got: %d", one.String(), another.String(), v)
	}
}

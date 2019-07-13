package kgo

import (
	"net"
	"testing"
)

func TestIsWindows(t *testing.T) {
	res := KOS.IsWindows()
	if res {
		t.Error("IsWindows fail")
		return
	}
}

func BenchmarkIsWindows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsWindows()
	}
}

func TestIsLinux(t *testing.T) {
	res := KOS.IsLinux()
	if !res {
		t.Error("IsLinux fail")
		return
	}
}

func BenchmarkIsLinux(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsLinux()
	}
}

func TestIsMac(t *testing.T) {
	res := KOS.IsMac()
	if res {
		t.Error("IsMac fail")
		return
	}
}

func BenchmarkIsMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsMac()
	}
}

func TestPwd(t *testing.T) {
	res := KOS.Pwd()
	if res == "" {
		t.Error("Pwd fail")
		return
	}
}

func BenchmarkPwd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.Pwd()
	}
}

func TestHomeDir(t *testing.T) {
	_, err := KOS.HomeDir()
	if err != nil {
		t.Error("Pwd fail")
		return
	}
}

func BenchmarkHomeDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.HomeDir()
	}
}

func TestLocalIP(t *testing.T) {
	_, err := KOS.LocalIP()
	if err != nil {
		t.Error("LocalIP fail")
		return
	}
}

func BenchmarkLocalIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.LocalIP()
	}
}

func TestOutboundIP(t *testing.T) {
	_, err := KOS.OutboundIP()
	if err != nil {
		t.Error("OutboundIP fail")
		return
	}
}

func BenchmarkOutboundIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.OutboundIP()
	}
}

func TestIsPublicIP(t *testing.T) {
	ipStr, _ := KOS.LocalIP()
	ipAddr := net.ParseIP(ipStr)
	if KOS.IsPublicIP(ipAddr) {
		t.Error("IsPublicIP fail")
		return
	}
}

func BenchmarkIsPublicIP(b *testing.B) {
	b.ResetTimer()
	ipStr, _ := KOS.LocalIP()
	ipAddr := net.ParseIP(ipStr)
	for i := 0; i < b.N; i++ {
		KOS.IsPublicIP(ipAddr)
	}
}

func TestGetIPs(t *testing.T) {
	ips := KOS.GetIPs()
	if len(ips) == 0 {
		t.Error("GetIPs fail")
		return
	}
}

func BenchmarkGetIPs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetIPs()
	}
}

func TestGetMacAddrs(t *testing.T) {
	macs := KOS.GetMacAddrs()
	//fmt.Printf("%v", macs)
	if len(macs) == 0 {
		t.Error("GetMacAddrs fail")
		return
	}
}

func BenchmarkGetMacAddrs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetMacAddrs()
	}
}

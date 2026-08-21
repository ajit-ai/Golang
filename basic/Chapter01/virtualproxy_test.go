package main

import "testing"

var _ IRealObject = &RealObject{}
var _ IRealObject = &VirtualProxy{}

func TestVirtualProxyLazyInstantiation(t *testing.T) {
	proxy := &VirtualProxy{}
	if proxy.realObject != nil {
		t.Fatal("realObject should be nil before first performAction call")
	}
	proxy.performAction()
	if proxy.realObject == nil {
		t.Error("performAction should lazily create the real object")
	}
}

func TestVirtualProxyMain(t *testing.T) {
	VirtualProxyMain()
}

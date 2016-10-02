package netifaces

import "testing"

func TestFirstOverview(t *testing.T) {
    gw, _ := FindAllSystemGateways()
    defer gw.Relese()
}

func TestDefaultIP4Gateway(t *testing.T) {
    gw, _ := FindAllSystemGateways()
    addr, _ := gw.FindDefaultIP4Gateway()
    t.Error("Default gateway " + addr)
}
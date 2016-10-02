package netifaces
/*
#include "netifaces.h"
*/
import "C"
import (
    "syscall"
    //"unsafe"
    "fmt"
)

type Gateway struct {
    gateway *C.Gateway
}

func FindAllSystemGateways() (Gateway, error) {
    var gw Gateway
    syserr := C.find_system_gateways(&gw.gateway)
    fmt.Printf(syscall.Errno(syserr).Error())
    return gw, nil
}

func (g *Gateway) Relese() {
    C.release_gateways(&g.gateway)
}

func (g *Gateway) DefaultIP4GatewayAddress() (string, error) {
    var gw *C.Gateway = C.find_default_ip4_gw(&g.gateway)
    if gw == nil {
        return "", fmt.Errorf("[ERR] Cannot find default gateway for IP4")
    }
    return C.GoString(gw.addr), nil
}

func (g *Gateway) DefaultIP4GatewayIface() (string, error) {
    var gw *C.Gateway = C.find_default_ip4_gw(&g.gateway)
    if gw == nil {
        return "", fmt.Errorf("[ERR] Cannot find default gateway for IP4")
    }
    return C.GoString(gw.iface), nil
}
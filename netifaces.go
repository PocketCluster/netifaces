package netifaces
/*
#include "netifaces.h"
*/
import "C"
import (
    "syscall"
    "fmt"
)

type Gateway struct {
    gateway *C.Gateway
}

func FindAllSystemGateways() (*Gateway, error) {
    var gw *Gateway = &Gateway{}
    syserr := C.find_system_gateways(&gw.gateway)
    if syserr != 0 {
        return nil, fmt.Errorf("[ERR] Cannot find all system gateways %s", syscall.Errno(syserr).Error())
    }
    return gw, nil
}

func (g *Gateway) Release() {
    C.release_gateways(&g.gateway)
}

func (g *Gateway) DefaultIP4Gateway() (address string, ifname string, err error) {
    var gw *C.Gateway = C.find_default_ip4_gw(&g.gateway)
    if gw == nil {
        err = fmt.Errorf("[ERR] Cannot find default gateway for IP4")
        return
    }
    address = C.GoString(gw.addr)
    ifname = C.GoString(gw.ifname)
    err = nil
    return
}
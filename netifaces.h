//
//  netifaces.h
//  netifaces
//
//  Created by Sung-Taek, Kim on 10/2/16.
//  Copyright © 2016 PocketCluster. All rights reserved.
//

#ifndef __NETIFACES_H__
#define __NETIFACES_H__

#ifndef WIN32

#include <errno.h>
#include <stdbool.h>

typedef struct Gateway {
    struct Gateway*    next;
    unsigned char      family;
    bool               is_default;
    char*              ifname;
    char*              addr;
} Gateway;

errno_t
find_system_gateways(Gateway** results);

void
release_gateways(Gateway** results);

Gateway*
find_default_ip4_gw(Gateway** results);

Gateway*
find_default_ip6_gw(Gateway** results);

#endif

#endif /* netifaces_h */
